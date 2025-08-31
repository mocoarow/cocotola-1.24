import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../models/word_problem.dart';
import '../../models/problem_display_config.dart';
import '../widgets/blank_widget.dart';
import '../widgets/hints_widget.dart';
import '../widgets/problem_content_widget.dart';
import '../../widgets/custom_keyboard.dart';
import 'dart:developer' as developer;

class ProblemDisplayScreen extends ConsumerStatefulWidget {
  final ProblemDisplayConfig config;
  final ProblemDisplayCallbacks callbacks;

  const ProblemDisplayScreen({
    super.key,
    required this.config,
    required this.callbacks,
  });

  @override
  ConsumerState<ProblemDisplayScreen> createState() =>
      _ProblemDisplayScreenState();
}

class _ProblemDisplayScreenState extends ConsumerState<ProblemDisplayScreen> {
  int _currentBlankIndex = 0;
  int _cursorPosition = 0;

  @override
  void initState() {
    super.initState();
    _currentBlankIndex = widget.config.initialBlankIndex;
    _cursorPosition = widget.config.initialCursorPosition;
  }

  @override
  void didUpdateWidget(ProblemDisplayScreen oldWidget) {
    super.didUpdateWidget(oldWidget);
    // 新しい問題に変わった場合、初期値をリセット
    if (oldWidget.config.currentIndex != widget.config.currentIndex) {
      _currentBlankIndex = widget.config.initialBlankIndex;
      _cursorPosition = widget.config.initialCursorPosition;
    }
  }

  @override
  Widget build(BuildContext context) {
    final currentProblem = widget.config.problems[widget.config.currentIndex];
    final requiredBlanks = currentProblem.blanks.length;

    developer.log(
        '[ProblemDisplayScreen] Current problem needs $requiredBlanks blanks, we have ${widget.config.answerControllers.length}');

    // コントローラーが初期化されていない場合は初期化を実行
    if (widget.config.answerControllers.isEmpty ||
        widget.config.answerControllers.length < requiredBlanks) {
      developer
          .log('[ProblemDisplayScreen] Controllers not ready, initializing...');
      widget.callbacks.onInitializeControllers(widget.config.problems);
      return const Scaffold(
        body: Center(child: CircularProgressIndicator()),
      );
    }

    final englishWords = currentProblem.english
        .replaceAll('.', ' .')
        .split(' ')
        .where((word) => word.isNotEmpty)
        .toList();

    // 複数の空欄のインデックスを取得
    final blankIndices = <int>[];
    for (int i = 0; i < englishWords.length; i++) {
      if (englishWords[i] == '___') {
        blankIndices.add(i);
      }
    }

    return Scaffold(
      appBar: AppBar(
        title: const Text('単語学習'),
      ),
      body: ProblemContentWidget(
        currentProblem: currentProblem,
        englishWords: englishWords,
        blankIndices: blankIndices,
        buildBlankWidget: (wordIndex, blankIndex, problem) => BlankWidget(
          wordIndex: wordIndex,
          blankIndex: blankIndex,
          problem: problem,
          controller: widget.config.answerControllers[blankIndex],
          focusNode: widget.config.answerFocusNodes[blankIndex],
          onChanged: (value) {
            developer.log('[ProblemDisplayScreen] Physical keyboard input for blank $blankIndex with value: "$value"');
            
            // 物理キーボード入力時の現在フォーカス状態を同期
            if (_currentBlankIndex != blankIndex) {
              developer.log('[ProblemDisplayScreen] Syncing focus from blank $_currentBlankIndex to $blankIndex');
              setState(() {
                _currentBlankIndex = blankIndex;
                _cursorPosition = widget.config.answerControllers[blankIndex].selection.end;
              });
              widget.callbacks.onBlankIndexChanged(blankIndex);
            } else {
              // 同じ空欄での物理キーボード入力時もカーソル位置を更新
              final newCursorPosition = widget.config.answerControllers[blankIndex].selection.end;
              developer.log('[ProblemDisplayScreen] Updating cursor position from $_cursorPosition to $newCursorPosition');
              setState(() {
                _cursorPosition = newCursorPosition;
              });
            }
            
            widget.callbacks.onAnswerChanged(blankIndex, value);
            // 物理キーボード入力時も自動チェック
            developer.log('[ProblemDisplayScreen] Calling auto-check for physical keyboard input');
            widget.callbacks.onAnswerChangedForAutoCheck(blankIndex, value);
          },
          onTap: () => _handleBlankTap(blankIndex),
        ),
        buildHintsSection: (problem) => HintsWidget(problem: problem),
        buildKeyboard: _buildKeyboard,
        buildActionButtons: _buildActionButtons,
      ),
    );
  }

  Widget _buildKeyboard(WordProblem problem) {
    if (problem.isCompleted) {
      return const SizedBox.shrink();
    }

    return CustomKeyboard(
      onKeyPressed: _handleKeyPressed,
      onDelete: _handleDeleteKey,
      onMoveLeft: _handleMoveLeft,
      onMoveRight: _handleMoveRight,
    );
  }

  Widget _buildActionButtons(WordProblem problem) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          ElevatedButton(
            onPressed: widget.callbacks.onShowAnswer,
            child: const Text('答えを見る'),
          ),
          if (!problem.isCompleted) ...[
            ElevatedButton(
              onPressed: widget.callbacks.onCheckAnswers,
              child: const Text('確認'),
            ),
          ],
          if (problem.isCompleted) ...[
            ElevatedButton(
              onPressed: widget.callbacks.onNextProblem,
              child: const Text('次へ'),
            ),
          ],
        ],
      ),
    );
  }

  void _handleBlankTap(int blankIndex) {
    final controller = widget.config.answerControllers[blankIndex];
    final cursorPos = controller.selection.isValid ? controller.selection.end : controller.text.length;
    developer.log('[ProblemDisplayScreen] _handleBlankTap: blank $blankIndex, cursor position: $cursorPos');
    
    setState(() {
      _currentBlankIndex = blankIndex;
      _cursorPosition = cursorPos;
    });
    widget.callbacks.onBlankTap(blankIndex);
    widget.callbacks.onBlankIndexChanged(blankIndex);
  }

  void _handleKeyPressed(String key) {
    developer.log('[ProblemDisplayScreen] _handleKeyPressed: "$key" at position $_cursorPosition in blank $_currentBlankIndex');
    if (_currentBlankIndex < widget.config.answerControllers.length) {
      final controller = widget.config.answerControllers[_currentBlankIndex];
      final text = controller.text;
      developer.log('[ProblemDisplayScreen] Current text: "$text", inserting at position $_cursorPosition');
      final newText = text.substring(0, _cursorPosition) +
          key +
          text.substring(_cursorPosition);
      controller.text = newText;
      _cursorPosition++;
      developer.log('[ProblemDisplayScreen] New text: "$newText", new cursor position: $_cursorPosition');
      controller.selection = TextSelection.fromPosition(
        TextPosition(offset: _cursorPosition),
      );

      // プロバイダーの状態を更新
      widget.callbacks.onAnswerChanged(_currentBlankIndex, newText);

      developer.log(
          '[ProblemDisplayScreen] _handleKeyPressed onAnswerChanged called for blank $_currentBlankIndex with input: $newText');
      // 自動チェック機能を呼び出し（親コンポーネントで処理）
      widget.callbacks.onAnswerChangedForAutoCheck(_currentBlankIndex, newText);
    }
  }

  void _handleDeleteKey() {
    if (_currentBlankIndex < widget.config.answerControllers.length &&
        _cursorPosition > 0) {
      final controller = widget.config.answerControllers[_currentBlankIndex];
      final text = controller.text;
      final newText = text.substring(0, _cursorPosition - 1) +
          text.substring(_cursorPosition);
      controller.text = newText;
      _cursorPosition--;
      controller.selection = TextSelection.fromPosition(
        TextPosition(offset: _cursorPosition),
      );

      // プロバイダーの状態を更新
      widget.callbacks.onAnswerChanged(_currentBlankIndex, newText);
    }
  }

  void _handleMoveLeft() {
    if (_cursorPosition > 0) {
      setState(() {
        _cursorPosition--;
      });
      if (_currentBlankIndex < widget.config.answerControllers.length) {
        widget.config.answerControllers[_currentBlankIndex].selection =
            TextSelection.fromPosition(
          TextPosition(offset: _cursorPosition),
        );
        widget.config.answerFocusNodes[_currentBlankIndex].requestFocus();
      }
    }
  }

  void _handleMoveRight() {
    if (_currentBlankIndex < widget.config.answerControllers.length) {
      final controller = widget.config.answerControllers[_currentBlankIndex];
      if (_cursorPosition < controller.text.length) {
        setState(() {
          _cursorPosition++;
        });
        controller.selection = TextSelection.fromPosition(
          TextPosition(offset: _cursorPosition),
        );
        widget.config.answerFocusNodes[_currentBlankIndex].requestFocus();
      }
    }
  }
}

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../models/word_problem.dart';
import '../../models/problem_display_config.dart';
import '../../providers/word_problem_provider.dart';
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
    _setInitialBlankIndex();
    _cursorPosition = 0;
  }

  @override
  void didUpdateWidget(ProblemDisplayScreen oldWidget) {
    super.didUpdateWidget(oldWidget);
    // 新しい問題に変わった場合、初期値をリセット
    if (oldWidget.config.currentIndex != widget.config.currentIndex) {
      _setInitialBlankIndex();
      _cursorPosition = 0;
    }
  }

  @override
  Widget build(BuildContext context) {
    // final currentProblem = widget.config.problems[widget.config.currentIndex];
    final currentProblem = widget.config.problem;
    final requiredBlanks = currentProblem.blanks.length;

    // コントローラーが初期化されていない場合は初期化を実行
    if (widget.config.answerControllers.isEmpty ||
        widget.config.answerControllers.length < requiredBlanks) {
      widget.callbacks.onInitializeControllers();
      return const Scaffold(
        body: Center(child: CircularProgressIndicator()),
      );
    }

    final englishWords = currentProblem.englishWords;
    final blankIndices = currentProblem.blankIndices;

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
          readOnly: false, // 問題表示時は編集可能
          onChanged: (value) {
            developer.log(
                '[ProblemDisplayScreen] Physical keyboard input for blank $blankIndex with value: "$value"');

            // 物理キーボード入力時の現在フォーカス状態を同期
            if (_currentBlankIndex != blankIndex) {
              developer.log(
                  '[ProblemDisplayScreen] Syncing focus from blank $_currentBlankIndex to $blankIndex');
              setState(() {
                _currentBlankIndex = blankIndex;
                _cursorPosition =
                    widget.config.answerControllers[blankIndex].selection.end;
              });
            } else {
              // 同じ空欄での物理キーボード入力時もカーソル位置を更新
              final newCursorPosition =
                  widget.config.answerControllers[blankIndex].selection.end;
              developer.log(
                  '[ProblemDisplayScreen] Updating cursor position from $_cursorPosition to $newCursorPosition');
              setState(() {
                _cursorPosition = newCursorPosition;
              });
            }

            widget.callbacks.onAnswerChanged(blankIndex, value);
            // 物理キーボード入力時も自動チェック
            developer.log(
                '[ProblemDisplayScreen] Calling auto-check for physical keyboard input');
            _checkAnswerAutomatically(blankIndex, value);
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
    final cursorPos = controller.selection.isValid
        ? controller.selection.end
        : controller.text.length;
    developer.log(
        '[ProblemDisplayScreen] _handleBlankTap: blank $blankIndex, cursor position: $cursorPos');

    setState(() {
      _currentBlankIndex = blankIndex;
      _cursorPosition = cursorPos;
    });
  }

  void _handleKeyPressed(String key) {
    developer.log(
        '[ProblemDisplayScreen] _handleKeyPressed: "$key" at position $_cursorPosition in blank $_currentBlankIndex');
    if (_currentBlankIndex < widget.config.answerControllers.length) {
      final controller = widget.config.answerControllers[_currentBlankIndex];
      final text = controller.text;
      developer.log(
          '[ProblemDisplayScreen] Current text: "$text", inserting at position $_cursorPosition');
      final newText = text.substring(0, _cursorPosition) +
          key +
          text.substring(_cursorPosition);
      controller.text = newText;
      _cursorPosition++;
      developer.log(
          '[ProblemDisplayScreen] New text: "$newText", new cursor position: $_cursorPosition');
      controller.selection = TextSelection.fromPosition(
        TextPosition(offset: _cursorPosition),
      );

      // プロバイダーの状態を更新
      widget.callbacks.onAnswerChanged(_currentBlankIndex, newText);

      developer.log(
          '[ProblemDisplayScreen] _handleKeyPressed onAnswerChanged called for blank $_currentBlankIndex with input: $newText');
      // 自動チェック機能を呼び出し
      _checkAnswerAutomatically(_currentBlankIndex, newText);
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

  void _checkAnswerAutomatically(int blankIndex, String value) {
    developer.log(
        '[ProblemDisplayScreen] _checkAnswerAutomatically called for blank $blankIndex with value "$value"');

    // final currentProblem = widget.config.problems[widget.config.currentIndex];
    final currentProblem = widget.config.problem;
    final blank = currentProblem.blanks[blankIndex];

    developer.log(
        '[ProblemDisplayScreen] Expected answer: "${blank.answer}", isAnswered: ${blank.isAnswered}');

    // 既に回答済みの場合は何もしない
    if (blank.isAnswered) {
      developer.log(
          '[ProblemDisplayScreen] Blank already answered, skipping auto-check');
      return;
    }

    final trimmedValue = value.trim();
    developer.log('[ProblemDisplayScreen] trimmedValue: "$trimmedValue"');
    if (trimmedValue.isNotEmpty) {
      final isCorrect =
          trimmedValue.toLowerCase() == blank.answer.toLowerCase();
      developer.log(
          '[ProblemDisplayScreen] Comparing "$trimmedValue" with "${blank.answer}", isCorrect: $isCorrect');

      if (isCorrect) {
        developer.log(
            '[ProblemDisplayScreen] Auto-check: Correct answer detected for blank $blankIndex');

        // プロバイダーで正解をマーク
        ref
            .read(wordProblemsProvider.notifier)
            .checkAnswer(widget.config.currentIndex, blankIndex, trimmedValue);

        developer.log(
            '[ProblemDisplayScreen] Answer marked as correct, checking completion...');

        // 正解時に次の未回答空欄にフォーカスを移動
        _moveToNextIncorrectBlank(blankIndex);

        // 全ての空欄が正解かチェック（自動で解説画面に遷移）
        _checkIfAllBlanksCompleted();
      } else {
        developer
            .log('[ProblemDisplayScreen] Answer not correct, continuing...');
      }
    } else {
      developer.log('[ProblemDisplayScreen] Empty value, skipping auto-check');
    }
  }

  void _moveToNextIncorrectBlank(int currentBlankIndex) {
    // 現在の問題をウィジェット設定から取得（ref.readを回避）
    final currentProblem = widget.config.problem;
    
    WidgetsBinding.instance.addPostFrameCallback((_) {
      // 次の未回答空欄を探す（現在の空欄の次から）
      int? nextIncorrectBlankIndex;

      // 現在の空欄より後ろを探す
      for (int i = currentBlankIndex + 1;
          i < currentProblem.blanks.length;
          i++) {
        if (!currentProblem.blanks[i].isAnswered ||
            !currentProblem.blanks[i].isCorrect) {
          nextIncorrectBlankIndex = i;
          break;
        }
      }

      // 後ろで見つからない場合、先頭から現在の空欄まで探す
      if (nextIncorrectBlankIndex == null) {
        for (int i = 0; i < currentBlankIndex; i++) {
          if (!currentProblem.blanks[i].isAnswered ||
              !currentProblem.blanks[i].isCorrect) {
            nextIncorrectBlankIndex = i;
            break;
          }
        }
      }

      // 未回答空欄が見つかった場合、フォーカスを移動
      if (nextIncorrectBlankIndex != null &&
          nextIncorrectBlankIndex < widget.config.answerFocusNodes.length) {
        developer.log(
            '[ProblemDisplayScreen] Moving focus to blank $nextIncorrectBlankIndex');
        setState(() {
          _currentBlankIndex = nextIncorrectBlankIndex!;
          _cursorPosition = widget
              .config.answerControllers[nextIncorrectBlankIndex].text.length;
        });
        widget.config.answerFocusNodes[nextIncorrectBlankIndex].requestFocus();
      }
    });
  }

  /// 最初の未回答空欄のインデックスを設定
  void _setInitialBlankIndex() {
    final currentProblem = widget.config.problem;
    
    // 最初の未回答空欄を見つける
    for (int i = 0; i < currentProblem.blanks.length; i++) {
      if (!currentProblem.blanks[i].isAnswered || !currentProblem.blanks[i].isCorrect) {
        _currentBlankIndex = i;
        developer.log('[ProblemDisplayScreen] Set initial blank index to $i');
        return;
      }
    }
    
    // 全て回答済みの場合は0にセット
    _currentBlankIndex = 0;
    developer.log('[ProblemDisplayScreen] All blanks answered, set blank index to 0');
  }

  void _checkIfAllBlanksCompleted() {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      // プロバイダーから最新の状態を取得
      final updatedProblems = ref.read(wordProblemsProvider);
      final updatedProblem = updatedProblems[widget.config.currentIndex];
      
      developer.log(
          '[ProblemDisplayScreen] Checking completion: isCompleted=${updatedProblem.isCompleted}');
      
      if (updatedProblem.isCompleted) {
        developer.log(
            '[ProblemDisplayScreen] All blanks completed, notifying parent');
        // 親コンポーネントに完了を通知
        widget.callbacks.onAllBlanksCompleted();
      }
    });
  }
}

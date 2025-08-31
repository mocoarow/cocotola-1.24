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
  ConsumerState<ProblemDisplayScreen> createState() => _ProblemDisplayScreenState();
}

class _ProblemDisplayScreenState extends ConsumerState<ProblemDisplayScreen> {
  @override
  Widget build(BuildContext context) {
    final currentProblem = widget.config.problems[widget.config.currentIndex];
    final requiredBlanks = currentProblem.blanks.length;
    
    developer.log(
        '[ProblemDisplayScreen] Current problem needs $requiredBlanks blanks, we have ${widget.config.answerControllers.length}');

    // コントローラーが初期化されていない場合は初期化を実行
    if (widget.config.answerControllers.isEmpty || widget.config.answerControllers.length < requiredBlanks) {
      developer.log('[ProblemDisplayScreen] Controllers not ready, initializing...');
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
          onChanged: (value) => widget.callbacks.onAnswerChanged(blankIndex, value),
          onTap: () => widget.callbacks.onBlankTap(blankIndex),
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
      onKeyPressed: widget.callbacks.onKeyPressed,
      onDelete: widget.callbacks.onDeleteKey,
      onMoveLeft: widget.callbacks.onMoveLeft,
      onMoveRight: widget.callbacks.onMoveRight,
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
}
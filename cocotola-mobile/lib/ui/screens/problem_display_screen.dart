import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../models/word_problem.dart';
import '../widgets/blank_widget.dart';
import '../widgets/hints_widget.dart';
import '../widgets/problem_content_widget.dart';
import '../../widgets/custom_keyboard.dart';
import 'dart:developer' as developer;

class ProblemDisplayScreen extends ConsumerWidget {
  final List<WordProblem> problems;
  final int currentIndex;
  final List<TextEditingController> answerControllers;
  final List<FocusNode> answerFocusNodes;
  final int currentBlankIndex;
  final int cursorPosition;
  final void Function(int blankIndex, String value) onAnswerChanged;
  final void Function(int blankIndex) onBlankTap;
  final VoidCallback onCheckAnswers;
  final VoidCallback onShowAnswer;
  final VoidCallback onNextProblem;
  final void Function(String) onKeyPressed;
  final VoidCallback onDeleteKey;
  final VoidCallback onMoveLeft;
  final VoidCallback onMoveRight;

  const ProblemDisplayScreen({
    super.key,
    required this.problems,
    required this.currentIndex,
    required this.answerControllers,
    required this.answerFocusNodes,
    required this.currentBlankIndex,
    required this.cursorPosition,
    required this.onAnswerChanged,
    required this.onBlankTap,
    required this.onCheckAnswers,
    required this.onShowAnswer,
    required this.onNextProblem,
    required this.onKeyPressed,
    required this.onDeleteKey,
    required this.onMoveLeft,
    required this.onMoveRight,
  });

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final currentProblem = problems[currentIndex];
    final requiredBlanks = currentProblem.blanks.length;
    
    developer.log(
        '[ProblemDisplayScreen] Current problem needs $requiredBlanks blanks, we have ${answerControllers.length}');

    // コントローラーが初期化されていない場合はローディング表示
    if (answerControllers.isEmpty || answerControllers.length < requiredBlanks) {
      developer.log('[ProblemDisplayScreen] Controllers not ready, showing loading...');
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
          controller: answerControllers[blankIndex],
          focusNode: answerFocusNodes[blankIndex],
          onChanged: (value) => onAnswerChanged(blankIndex, value),
          onTap: () => onBlankTap(blankIndex),
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
      onKeyPressed: onKeyPressed,
      onDelete: onDeleteKey,
      onMoveLeft: onMoveLeft,
      onMoveRight: onMoveRight,
    );
  }

  Widget _buildActionButtons(WordProblem problem) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          ElevatedButton(
            onPressed: onShowAnswer,
            child: const Text('答えを見る'),
          ),
          if (!problem.isCompleted) ...[
            ElevatedButton(
              onPressed: onCheckAnswers,
              child: const Text('確認'),
            ),
          ],
          if (problem.isCompleted) ...[
            ElevatedButton(
              onPressed: onNextProblem,
              child: const Text('次へ'),
            ),
          ],
        ],
      ),
    );
  }
}
import 'package:flutter/material.dart';
import '../../models/word_problem.dart';
import '../widgets/blank_widget.dart';
import '../widgets/hints_widget.dart';
import '../widgets/problem_content_widget.dart';
import 'dart:developer' as developer;

class AnswerDisplayScreen extends StatelessWidget {
  final WordProblem problem;
  final int currentIndex;
  final List<TextEditingController> answerControllers;
  final List<FocusNode> answerFocusNodes;
  final VoidCallback onNextProblem;

  const AnswerDisplayScreen({
    super.key,
    required this.problem,
    required this.currentIndex,
    required this.answerControllers,
    required this.answerFocusNodes,
    required this.onNextProblem,
  });

  @override
  Widget build(BuildContext context) {
    final currentProblem = problem;

    final englishWords = currentProblem.englishWords;
    final blankIndices = currentProblem.blankIndices;
    
    developer.log('[AnswerDisplayScreen] Building answer display for problem $currentIndex');
    developer.log('[AnswerDisplayScreen] Problem has ${currentProblem.blanks.length} blanks');
    developer.log('[AnswerDisplayScreen] Controllers count: ${answerControllers.length}');
    
    for (int i = 0; i < answerControllers.length && i < currentProblem.blanks.length; i++) {
      developer.log('[AnswerDisplayScreen] Controller $i text: "${answerControllers[i].text}"');
      developer.log('[AnswerDisplayScreen] Blank $i answer: "${currentProblem.blanks[i].answer}"');
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
          onChanged: (_) {}, // 解説表示時は変更を受け付けない
          onTap: () {}, // 解説表示時はタップを受け付けない
          readOnly: true, // 答え表示時は読み取り専用
        ),
        buildHintsSection: (problem) => HintsWidget(problem: problem),
        buildKeyboard: _buildKeyboard, // キーボード非表示
        buildActionButtons: _buildActionButtons, // 次へボタンのみ
      ),
    );
  }

  Widget _buildKeyboard(WordProblem problem) {
    // 解説表示時はキーボードを表示しない
    return const SizedBox.shrink();
  }

  Widget _buildActionButtons(WordProblem problem) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          ElevatedButton(
            onPressed: onNextProblem,
            child: const Text('次へ'),
          ),
        ],
      ),
    );
  }
}

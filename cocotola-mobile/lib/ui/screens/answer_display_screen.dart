import 'package:flutter/material.dart';
import '../../models/word_problem.dart';
import '../widgets/blank_widget.dart';
import '../widgets/hints_widget.dart';
import '../widgets/problem_content_widget.dart';

class AnswerDisplayScreen extends StatelessWidget {
  final List<WordProblem> problems;
  final int currentIndex;
  final List<TextEditingController> answerControllers;
  final List<FocusNode> answerFocusNodes;
  final VoidCallback onNextProblem;

  const AnswerDisplayScreen({
    super.key,
    required this.problems,
    required this.currentIndex,
    required this.answerControllers,
    required this.answerFocusNodes,
    required this.onNextProblem,
  });

  @override
  Widget build(BuildContext context) {
    final currentProblem = problems[currentIndex];
    
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
          controller: answerControllers[blankIndex],
          focusNode: answerFocusNodes[blankIndex],
          onChanged: (_) {}, // 解説表示時は変更を受け付けない
          onTap: () {}, // 解説表示時はタップを受け付けない
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
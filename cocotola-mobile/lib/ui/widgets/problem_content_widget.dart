import 'package:flutter/material.dart';
import '../../models/word_problem.dart';

class ProblemContentWidget extends StatelessWidget {
  final WordProblem currentProblem;
  final List<String> englishWords;
  final List<int> blankIndices;
  final Widget Function(int wordIndex, int blankIndex, WordProblem problem) buildBlankWidget;
  final Widget Function(WordProblem problem) buildHintsSection;
  final Widget Function(WordProblem problem) buildKeyboard;
  final Widget Function(WordProblem problem) buildActionButtons;

  const ProblemContentWidget({
    super.key,
    required this.currentProblem,
    required this.englishWords,
    required this.blankIndices,
    required this.buildBlankWidget,
    required this.buildHintsSection,
    required this.buildKeyboard,
    required this.buildActionButtons,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Padding(
          padding: const EdgeInsets.all(16.0),
          child: Text(
            currentProblem.japanese,
            style: Theme.of(context).textTheme.headlineSmall,
          ),
        ),
        Padding(
          padding: const EdgeInsets.all(16.0),
          child: Wrap(
            alignment: WrapAlignment.center,
            children: [
              for (int i = 0; i < englishWords.length; i++)
                if (blankIndices.contains(i))
                  buildBlankWidget(i, blankIndices.indexOf(i), currentProblem)
                else
                  Padding(
                    padding: const EdgeInsets.symmetric(
                        horizontal: 4.0, vertical: 2.0),
                    child: Text(
                      englishWords[i],
                      style: const TextStyle(fontSize: 16),
                    ),
                  ),
            ],
          ),
        ),
        buildHintsSection(currentProblem),
        const Spacer(),
        buildKeyboard(currentProblem),
        buildActionButtons(currentProblem),
      ],
    );
  }
}
import 'package:flutter/material.dart';
import '../../models/word_problem.dart';
import 'cefr_level_badge.dart';

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
        // 問題文をカード形式で表示
        Flexible(
          child: SingleChildScrollView(
            child: Column(
              children: [
                Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: Card(
                    elevation: 4,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(12),
                    ),
                    child: Padding(
                      padding: const EdgeInsets.all(20.0),
                      child: Column(
                        children: [
                          // 日本語文とCEFRレベル
                          Container(
                            width: double.infinity,
                            padding: const EdgeInsets.all(16.0),
                            decoration: BoxDecoration(
                              color: Colors.blue.shade50,
                              borderRadius: BorderRadius.circular(8),
                              border: Border.all(color: Colors.blue.shade200),
                            ),
                            child: Column(
                              children: [
                                // CEFRレベルバッジ
                                Align(
                                  alignment: Alignment.topRight,
                                  child: CefrLevelBadge(
                                    cefrLevel: currentProblem.cefrLevel,
                                  ),
                                ),
                                const SizedBox(height: 12),
                                // 日本語文
                                Text(
                                  currentProblem.japanese,
                                  style: Theme.of(context).textTheme.headlineSmall?.copyWith(
                                    color: Colors.blue.shade800,
                                    fontWeight: FontWeight.w600,
                                  ),
                                  textAlign: TextAlign.center,
                                ),
                              ],
                            ),
                          ),
                          const SizedBox(height: 20),
                          // 英語文（空欄問題）
                          Container(
                            width: double.infinity,
                            padding: const EdgeInsets.all(16.0),
                            decoration: BoxDecoration(
                              color: Colors.grey.shade50,
                              borderRadius: BorderRadius.circular(8),
                              border: Border.all(color: Colors.grey.shade300),
                            ),
                            child: Wrap(
                              alignment: WrapAlignment.start,
                              crossAxisAlignment: WrapCrossAlignment.center,
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
                                        style: const TextStyle(
                                          fontSize: 18,
                                          fontWeight: FontWeight.w500,
                                        ),
                                      ),
                                    ),
                              ],
                            ),
                          ),
                        ],
                      ),
                    ),
                  ),
                ),
                buildHintsSection(currentProblem),
              ],
            ),
          ),
        ),
        buildKeyboard(currentProblem),
        buildActionButtons(currentProblem),
      ],
    );
  }
}
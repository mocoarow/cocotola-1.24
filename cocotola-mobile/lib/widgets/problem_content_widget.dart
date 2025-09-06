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
    return Container(
      color: Colors.grey.shade100, // 背景を薄いグレーに
      child: Column(
        children: [
          // 問題文をカード形式で表示
          Flexible(
            child: SingleChildScrollView(
              child: Column(
                children: [
                  const SizedBox(height: 16),
                  Padding(
                    padding: const EdgeInsets.all(16.0),
                    child: Card(
                      elevation: 2,
                      color: Colors.white, // カードの背景を白に
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12),
                      ),
                      child: Column(
                        children: [
                          // 日本語文セクション
                          Padding(
                            padding: const EdgeInsets.all(20.0),
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
                                  style: Theme.of(context).textTheme.titleMedium?.copyWith(
                                    color: Theme.of(context).colorScheme.onSurface,
                                    fontWeight: FontWeight.w500,
                                  ),
                                  textAlign: TextAlign.center,
                                ),
                              ],
                            ),
                          ),
                          // 区切り線
                          Divider(
                            color: Theme.of(context).colorScheme.outline.withValues(alpha: 0.3),
                            thickness: 1,
                            height: 1,
                          ),
                          // 英語文セクション（空欄問題）
                          Padding(
                            padding: const EdgeInsets.all(20.0),
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
                                        style: Theme.of(context).textTheme.titleMedium?.copyWith(
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
                ],
              ),
            ),
          ),
          buildKeyboard(currentProblem),
          buildActionButtons(currentProblem),
        ],
      ),
    );
  }
}
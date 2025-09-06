import 'package:flutter/material.dart';
import '../models/memorization_problem.dart';
import 'cefr_level_badge.dart';

/// 暗記問題の解答表示ウィジェット
class MemorizationAnswerWidget extends StatelessWidget {
  final MemorizationProblem problem;
  final Function(bool wasCorrect)? onAnswer;

  const MemorizationAnswerWidget({
    super.key,
    required this.problem,
    this.onAnswer,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.grey.shade100,
      child: Column(
        children: [
          // メインコンテンツ（スクロール可能）
          Expanded(
            child: SingleChildScrollView(
              child: Column(
                children: [
                  const SizedBox(height: 16),
                  // 問題と答えの統合カード
                  Padding(
                    padding: const EdgeInsets.all(16.0),
                    child: Card(
                      elevation: 2,
                      color: Colors.white,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12),
                      ),
                      child: Column(
                        children: [
                          // 問題文セクション
                          Padding(
                            padding: const EdgeInsets.all(20.0),
                            child: Column(
                              children: [
                                // CEFRレベルバッジ
                                Align(
                                  alignment: Alignment.topRight,
                                  child: CefrLevelBadge(
                                    cefrLevel: problem.cefrLevel,
                                  ),
                                ),
                                const SizedBox(height: 12),
                                // 問題文
                                Text(
                                  problem.question,
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
                          // 答えセクション
                          Padding(
                            padding: const EdgeInsets.all(20.0),
                            child: Column(
                              children: [
                                Icon(
                                  Icons.lightbulb,
                                  size: 32,
                                  color: Theme.of(context).colorScheme.primary,
                                ),
                                const SizedBox(height: 12),
                                Text(
                                  '答え',
                                  style: Theme.of(context).textTheme.labelLarge?.copyWith(
                                    color: Theme.of(context).colorScheme.primary,
                                    fontWeight: FontWeight.w600,
                                  ),
                                ),
                                const SizedBox(height: 8),
                                Text(
                                  problem.answer,
                                  style: Theme.of(context).textTheme.titleLarge?.copyWith(
                                    color: Theme.of(context).colorScheme.onSurface,
                                    fontWeight: FontWeight.w600,
                                    height: 1.3,
                                  ),
                                  textAlign: TextAlign.center,
                                ),
                              ],
                            ),
                          ),
                        ],
                      ),
                    ),
                  ),
                  // 下部ボタン用の余白
                  const SizedBox(height: 100),
                ],
              ),
            ),
          ),
          // 理解度選択ボタン（画面下部に固定）
          SafeArea(
            child: Container(
              color: Colors.grey.shade100,
              child: _buildAnswerButtons(context),
            ),
          ),
        ],
      ),
    );
  }


  Widget _buildAnswerButtons(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        children: [
          Row(
            children: [
              Expanded(
                child: ElevatedButton.icon(
                  onPressed: () => onAnswer?.call(false),
                  icon: const Icon(Icons.close),
                  label: const Text('できなかった'),
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Theme.of(context).colorScheme.errorContainer,
                    foregroundColor: Theme.of(context).colorScheme.onErrorContainer,
                    padding: const EdgeInsets.symmetric(vertical: 12),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                ),
              ),
              const SizedBox(width: 16),
              Expanded(
                child: ElevatedButton.icon(
                  onPressed: () => onAnswer?.call(true),
                  icon: const Icon(Icons.check),
                  label: const Text('できた'),
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Theme.of(context).colorScheme.primary,
                    foregroundColor: Theme.of(context).colorScheme.onPrimary,
                    padding: const EdgeInsets.symmetric(vertical: 12),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
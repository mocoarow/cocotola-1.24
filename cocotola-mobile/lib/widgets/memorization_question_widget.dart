import 'package:flutter/material.dart';
import '../models/memorization_problem.dart';
import 'cefr_level_badge.dart';

/// 暗記問題の問題文表示ウィジェット
class MemorizationQuestionWidget extends StatelessWidget {
  final MemorizationProblem problem;

  const MemorizationQuestionWidget({
    super.key,
    required this.problem,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.grey.shade100, // 背景を薄いグレーに
      child: Column(
        children: [
          Flexible(
            child: SingleChildScrollView(
              child: Column(
                children: [
                  const SizedBox(height: 16),
                  // 問題文カード（シンプルな白いカード）
                  Padding(
                    padding: const EdgeInsets.all(16.0),
                    child: Card(
                      elevation: 2,
                      color: Colors.white, // カードの背景を白に
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12),
                      ),
                      child: Padding(
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
                    ),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }

}
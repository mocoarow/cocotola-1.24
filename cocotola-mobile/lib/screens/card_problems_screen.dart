import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/card_provider.dart';
import '../models/problem_base.dart';

class CardProblemsScreen extends ConsumerStatefulWidget {
  const CardProblemsScreen({super.key});

  @override
  ConsumerState<CardProblemsScreen> createState() => _CardProblemsScreenState();
}

class _CardProblemsScreenState extends ConsumerState<CardProblemsScreen> {
  @override
  void initState() {
    super.initState();
    // 画面初期化時にカード問題を取得
    WidgetsBinding.instance.addPostFrameCallback((_) {
      ref.read(cardProblemsNotifierProvider.notifier).fetchCards();
    });
  }

  @override
  Widget build(BuildContext context) {
    final cardProblemsAsync = ref.watch(cardProblemsNotifierProvider);

    return Scaffold(
      appBar: AppBar(
        title: const Text('カード問題'),
        actions: [
          IconButton(
            icon: const Icon(Icons.refresh),
            onPressed: () {
              ref.read(cardProblemsNotifierProvider.notifier).refresh();
            },
          ),
        ],
      ),
      body: cardProblemsAsync.when(
        loading: () => const Center(
          child: CircularProgressIndicator(),
        ),
        error: (error, stackTrace) => Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const Icon(
                Icons.error_outline,
                size: 64,
                color: Colors.red,
              ),
              const SizedBox(height: 16),
              Text(
                'エラーが発生しました',
                style: Theme.of(context).textTheme.headlineSmall,
              ),
              const SizedBox(height: 8),
              Text(
                error.toString(),
                style: Theme.of(context).textTheme.bodyMedium,
                textAlign: TextAlign.center,
              ),
              const SizedBox(height: 16),
              ElevatedButton(
                onPressed: () {
                  ref.read(cardProblemsNotifierProvider.notifier).refresh();
                },
                child: const Text('再試行'),
              ),
            ],
          ),
        ),
        data: (problems) => problems.isEmpty
            ? const Center(
                child: Text(
                  'カード問題がありません',
                  style: TextStyle(fontSize: 18),
                ),
              )
            : ListView.builder(
                padding: const EdgeInsets.all(16),
                itemCount: problems.length,
                itemBuilder: (context, index) {
                  final problem = problems[index];
                  return Card(
                    margin: const EdgeInsets.only(bottom: 16),
                    child: Padding(
                      padding: const EdgeInsets.all(16),
                      child: _buildProblemContent(problem),
                    ),
                  );
                },
              ),
      ),
    );
  }

  Widget _buildProblemContent(Problem problem) {
    switch (problem.type) {
      case ProblemType.word:
        final wordProblem = problem.wordProblem!;
        return Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                Chip(
                  label: Text(wordProblem.cefrLevel),
                  backgroundColor: Colors.blue.shade100,
                ),
                const Spacer(),
                Text(
                  problem.isCompleted ? '完了' : '未完了',
                  style: TextStyle(
                    color: problem.isCompleted ? Colors.green : Colors.grey,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ],
            ),
            const SizedBox(height: 8),
            Text(
              '日本語: ${wordProblem.japanese}',
              style: const TextStyle(fontSize: 16),
            ),
            const SizedBox(height: 4),
            Text(
              '英語: ${wordProblem.english}',
              style: const TextStyle(
                fontSize: 16,
                fontWeight: FontWeight.bold,
              ),
            ),
            if (wordProblem.blanks.isNotEmpty) ...[
              const SizedBox(height: 12),
              const Text(
                'ヒント:',
                style: TextStyle(
                  fontSize: 14,
                  fontWeight: FontWeight.bold,
                ),
              ),
              ...wordProblem.blanks.map((blank) => Padding(
                    padding: const EdgeInsets.only(left: 8, top: 4),
                    child: Text(
                      '• ${blank.hint} (答え: ${blank.answer})',
                      style: const TextStyle(fontSize: 14),
                    ),
                  )),
            ],
          ],
        );
      case ProblemType.memorization:
        // 暗記問題の場合の実装（必要に応じて）
        return const Text('暗記問題（未実装）');
    }
  }
}
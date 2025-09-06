import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/card_provider.dart';
import '../providers/problem_set_provider.dart';
import '../models/problem_base.dart';

class ApiDebugScreen extends ConsumerWidget {
  const ApiDebugScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('API Debug Screen'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Card Problems Provider のテスト
            const Text(
              'Card Problems Provider Test:',
              style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 8),
            ElevatedButton(
              onPressed: () async {
                try {
                  final problems = await ref.read(cardProblemsProvider.future);
                  if (context.mounted) {
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(
                        content: Text('取得成功: ${problems.length}問'),
                        backgroundColor: Colors.green,
                      ),
                    );
                  }
                } catch (e) {
                  if (context.mounted) {
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(
                        content: Text('エラー: $e'),
                        backgroundColor: Colors.red,
                      ),
                    );
                  }
                }
              },
              child: const Text('Fetch Card Problems'),
            ),
            const SizedBox(height: 16),
            
            // Card Problems の表示
            Consumer(
              builder: (context, ref, child) {
                final cardProblemsAsync = ref.watch(cardProblemsProvider);
                
                return cardProblemsAsync.when(
                  loading: () => const CircularProgressIndicator(),
                  error: (error, stack) => Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text('Error: $error'),
                      if (stack != null) Text('Stack: $stack'),
                    ],
                  ),
                  data: (problems) => Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text('問題数: ${problems.length}'),
                      if (problems.isNotEmpty) ...[
                        const SizedBox(height: 8),
                        const Text('最初の問題:'),
                        _buildProblemInfo(problems.first),
                      ],
                    ],
                  ),
                );
              },
            ),
            const SizedBox(height: 24),
            
            // Problem Sets Provider のテスト
            const Text(
              'Problem Sets Provider Test:',
              style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 8),
            ElevatedButton(
              onPressed: () async {
                try {
                  final problemSetsNotifier = ref.read(problemSetsProvider.notifier);
                  await problemSetsNotifier.addProblemSetFromApi();
                  if (context.mounted) {
                    ScaffoldMessenger.of(context).showSnackBar(
                      const SnackBar(
                        content: Text('問題セット追加成功'),
                        backgroundColor: Colors.green,
                      ),
                    );
                  }
                } catch (e) {
                  if (context.mounted) {
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(
                        content: Text('エラー: $e'),
                        backgroundColor: Colors.red,
                      ),
                    );
                  }
                }
              },
              child: const Text('Add API Problem Set'),
            ),
            const SizedBox(height: 16),
            
            // Problem Sets の表示
            Consumer(
              builder: (context, ref, child) {
                final problemSets = ref.watch(problemSetsProvider);
                
                return Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text('問題セット数: ${problemSets.length}'),
                    ...problemSets.map((set) => Padding(
                          padding: const EdgeInsets.only(top: 4),
                          child: Text('• ${set.title} (${set.problems.length}問)'),
                        )),
                  ],
                );
              },
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildProblemInfo(Problem problem) {
    switch (problem.type) {
      case ProblemType.word:
        final wordProblem = problem.wordProblem!;
        return Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text('Japanese: ${wordProblem.japanese}'),
            Text('English: ${wordProblem.english}'),
            Text('CEFR Level: ${wordProblem.cefrLevel}'),
            Text('Blanks: ${wordProblem.blanks.length}'),
            ...wordProblem.blanks.asMap().entries.map((entry) {
              final index = entry.key;
              final blank = entry.value;
              return Text('  Blank ${index + 1}: ${blank.answer} (${blank.hint})');
            }),
          ],
        );
      case ProblemType.memorization:
        return const Text('Memorization Problem');
    }
  }
}
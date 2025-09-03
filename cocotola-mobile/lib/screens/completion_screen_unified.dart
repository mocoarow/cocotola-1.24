import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../view_models/app_state_manager.dart';

class CompletionScreenUnified extends ConsumerWidget {
  const CompletionScreenUnified({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final appState = ref.watch(appStateProvider);
    final appStateManager = ref.read(appStateProvider.notifier);

    final problemSetTitle = appState.selectedProblemSet?.title ?? '問題セット';

    return Scaffold(
      appBar: AppBar(
        title: Text(problemSetTitle),
      ),
      body: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            // 完了アイコン
            Container(
              padding: const EdgeInsets.all(20),
              decoration: BoxDecoration(
                color: Colors.green.shade100,
                shape: BoxShape.circle,
              ),
              child: Icon(
                Icons.check_circle,
                size: 80,
                color: Colors.green.shade600,
              ),
            ),
            const SizedBox(height: 32),
            // 完了メッセージ
            Text(
              'お疲れ様でした！',
              style: Theme.of(context).textTheme.headlineMedium?.copyWith(
                fontWeight: FontWeight.bold,
                color: Colors.green.shade700,
              ),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 16),
            Text(
              '全問正解です！🎉',
              style: Theme.of(context).textTheme.headlineSmall?.copyWith(
                color: Colors.green.shade600,
              ),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 48),
            // ボタン群
            Column(
              children: [
                // もう一度挑戦ボタン
                SizedBox(
                  width: double.infinity,
                  child: ElevatedButton.icon(
                    onPressed: () => appStateManager.resetProblemSet(),
                    icon: const Icon(Icons.refresh),
                    label: const Text(
                      'もう一度挑戦',
                      style: TextStyle(fontSize: 18, fontWeight: FontWeight.w600),
                    ),
                    style: ElevatedButton.styleFrom(
                      padding: const EdgeInsets.symmetric(vertical: 16),
                      backgroundColor: Colors.blue.shade600,
                      foregroundColor: Colors.white,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12),
                      ),
                    ),
                  ),
                ),
                const SizedBox(height: 16),
                // 問題セット選択に戻るボタン
                SizedBox(
                  width: double.infinity,
                  child: OutlinedButton.icon(
                    onPressed: () => appStateManager.returnToProblemSelection(),
                    icon: const Icon(Icons.list),
                    label: const Text(
                      '問題セット選択に戻る',
                      style: TextStyle(fontSize: 16, fontWeight: FontWeight.w500),
                    ),
                    style: OutlinedButton.styleFrom(
                      padding: const EdgeInsets.symmetric(vertical: 16),
                      foregroundColor: Colors.blue.shade600,
                      side: BorderSide(color: Colors.blue.shade600),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12),
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
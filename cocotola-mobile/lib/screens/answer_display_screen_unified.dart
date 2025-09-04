import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../view_models/app_state_manager.dart';
import '../models/problem_base.dart';
import '../widgets/problem_content_widget.dart';
import '../widgets/blank_widget.dart';
import '../widgets/hints_widget.dart';
import '../widgets/memorization_answer_widget.dart';

class AnswerDisplayScreenUnified extends ConsumerWidget {
  const AnswerDisplayScreenUnified({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final appState = ref.watch(appStateProvider);
    final appStateManager = ref.read(appStateProvider.notifier);

    if (appState.selectedProblemSet == null || appState.currentProblem == null) {
      return Scaffold(
        appBar: AppBar(title: const Text('エラー')),
        body: const Center(child: Text('問題が見つかりません')),
      );
    }

    final currentProblem = appState.currentProblem!;

    // 問題タイプによって処理を分ける
    if (currentProblem.type == ProblemType.memorization) {
      return _buildMemorizationAnswer(context, appState, appStateManager);
    } else {
      return _buildWordAnswer(context, appState, appStateManager);
    }
  }

  /// 暗記問題の解答画面構築
  Widget _buildMemorizationAnswer(BuildContext context, AppState appState, AppStateManager appStateManager) {
    final currentProblem = appState.currentProblem!;
    final memorizationProblem = currentProblem.memorizationProblem!;
    
    return Scaffold(
      appBar: AppBar(
        title: Text(appState.selectedProblemSet!.title),
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () => appStateManager.returnToMenu(),
          tooltip: 'メニューに戻る',
        ),
      ),
      body: MemorizationAnswerWidget(
        problem: memorizationProblem,
        onAnswer: (wasCorrect) => appStateManager.handleMemorizationAnswer(wasCorrect),
      ),
    );
  }

  /// 英単語問題の解答画面構築
  Widget _buildWordAnswer(BuildContext context, AppState appState, AppStateManager appStateManager) {
    final currentProblem = appState.currentProblem!;
    final wordProblem = currentProblem.wordProblem!;
    final englishWords = wordProblem.englishWords;
    final blankIndices = wordProblem.blankIndices;

    return Scaffold(
      appBar: AppBar(
        title: Text(appState.selectedProblemSet!.title),
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () => appStateManager.returnToMenu(),
          tooltip: 'メニューに戻る',
        ),
      ),
      body: Column(
        children: [
          // 答え表示部分
          Expanded(
            child: ProblemContentWidget(
              currentProblem: wordProblem,
              englishWords: englishWords,
              blankIndices: blankIndices,
              buildBlankWidget: (wordIndex, blankIndex, problem) => BlankWidget(
                wordIndex: wordIndex,
                blankIndex: blankIndex,
                problem: problem,
                controller: appState.answerControllers[blankIndex],
                focusNode: appState.answerFocusNodes[blankIndex],
                readOnly: true, // 答え表示時は読み取り専用
                onChanged: (_) {}, // 読み取り専用なので何もしない
                onTap: () {}, // 読み取り専用なので何もしない
              ),
              buildHintsSection: (problem) => HintsWidget(problem: problem),
              buildKeyboard: (_) => const SizedBox.shrink(), // キーボード非表示
              buildActionButtons: (_) => const SizedBox.shrink(), // アクションボタンは下部に移動
            ),
          ),
          // アクションボタン（画面最下部に固定）
          _buildActionButtons(appStateManager),
        ],
      ),
    );
  }

  Widget _buildActionButtons(AppStateManager appStateManager) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: SizedBox(
        width: double.infinity,
        child: ElevatedButton(
          onPressed: () => appStateManager.moveToNextProblem(),
          style: ElevatedButton.styleFrom(
            padding: const EdgeInsets.symmetric(vertical: 16),
            backgroundColor: Colors.blue.shade600,
            foregroundColor: Colors.white,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(12),
            ),
          ),
          child: const Text(
            '次の問題へ',
            style: TextStyle(fontSize: 18, fontWeight: FontWeight.w600),
          ),
        ),
      ),
    );
  }
}
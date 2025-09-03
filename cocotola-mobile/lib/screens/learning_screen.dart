import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/problem_set_provider.dart';
import '../models/word_problem.dart';
import '../models/problem_display_config.dart';
import '../ui/screens/answer_display_screen.dart';
import '../ui/screens/problem_display_screen.dart';
import '../view_models/learning_view_model.dart';
import 'dart:developer' as developer;

class LearningScreen extends ConsumerWidget {
  final String problemSetId;
  
  const LearningScreen({
    super.key,
    required this.problemSetId,
  });

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final problemSetsNotifier = ref.read(problemSetsProvider.notifier);
    final problemSet = problemSetsNotifier.getProblemSetById(problemSetId);
    
    if (problemSet == null) {
      return Scaffold(
        appBar: AppBar(title: const Text('エラー')),
        body: const Center(child: Text('問題セットが見つかりません')),
      );
    }
    
    final problems = problemSet.problems;
    final viewState = ref.watch(learningViewModelProvider);
    final viewModel = ref.read(learningViewModelProvider.notifier);

    developer.log(
        '[LearningScreen] build called - problems count: ${problems.length}, currentIndex: ${viewState.currentIndex}, state: ${viewState.currentState}');

    if (problems.isEmpty) {
      developer.log('[LearningScreen] No problems available');
      return Scaffold(
        appBar: AppBar(title: Text(problemSet.title)),
        body: const Center(child: Text('問題がありません')),
      );
    }

    // 問題が初期化されていない場合は初期化
    if (viewState.problems.isEmpty) {
      WidgetsBinding.instance.addPostFrameCallback((_) {
        viewModel.initializeProblems(problems);
        viewModel.initializeControllersForCurrentProblem(null, problems);
      });
      return Scaffold(
        appBar: AppBar(title: Text(problemSet.title)),
        body: const Center(child: CircularProgressIndicator()),
      );
    }

    // コントローラーが初期化されていない場合は初期化
    if (viewState.answerControllers.isEmpty) {
      WidgetsBinding.instance.addPostFrameCallback((_) {
        viewModel.initializeControllersForCurrentProblem(null, viewState.problems);
      });
      return Scaffold(
        appBar: AppBar(title: Text(problemSet.title)),
        body: const Center(child: CircularProgressIndicator()),
      );
    }

    // 状態に基づいてUIを返す
    switch (viewState.currentState) {
      case LearningState.completed:
        return _buildCompletedScreen(problemSet.title, viewModel, context);

      case LearningState.answerDisplay:
        return _buildAnswerDisplayScreen(viewState.problems, viewState, viewModel, problemSet.title);

      case LearningState.problemDisplay:
        return _buildProblemDisplayScreen(viewState.problems, viewState, viewModel, problemSet.title);
    }
  }

  Widget _buildCompletedScreen(String problemSetTitle, LearningViewModel viewModel, BuildContext context) {
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
                    onPressed: () => viewModel.resetProblemSet(),
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
                    onPressed: () => Navigator.of(context).popUntil((route) => route.isFirst),
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

  Widget _buildAnswerDisplayScreen(List<WordProblem> problems,
      LearningViewState viewState, LearningViewModel viewModel, String problemSetTitle) {
    return AnswerDisplayScreen(
      problem: problems[viewState.currentIndex],
      currentIndex: viewState.currentIndex,
      answerControllers: viewState.answerControllers,
      answerFocusNodes: viewState.answerFocusNodes,
      onNextProblem: viewModel.transitionToNextProblem,
    );
  }

  Widget _buildProblemDisplayScreen(List<WordProblem> problems,
      LearningViewState viewState, LearningViewModel viewModel, String problemSetTitle) {
    return ProblemDisplayScreen(
      config: ProblemDisplayConfig(
        problem: problems[viewState.currentIndex],
        currentIndex: viewState.currentIndex,
        answerControllers: viewState.answerControllers,
        answerFocusNodes: viewState.answerFocusNodes,
      ),
      callbacks: ProblemDisplayCallbacks(
        onAnswerChanged: viewModel.handleAnswerChanged,
        onAllBlanksCompleted: viewModel.handleAllBlanksCompleted,
        onShowAnswer: viewModel.handleShowAnswer,
        onNextProblem: viewModel.transitionToAnswerDisplay,
        onInitializeControllers: () =>
            viewModel.initializeControllersForCurrentProblem(null, viewState.problems),
        getCurrentProblem: () => viewState.problems[viewState.currentIndex],
      ),
    );
  }
}

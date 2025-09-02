import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/word_problem_provider.dart';
import '../models/word_problem.dart';
import '../models/problem_display_config.dart';
import '../ui/screens/answer_display_screen.dart';
import '../ui/screens/problem_display_screen.dart';
import '../view_models/learning_view_model.dart';
import 'dart:developer' as developer;

class LearningScreen extends ConsumerWidget {
  const LearningScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final problems = ref.watch(wordProblemsProvider);
    final viewState = ref.watch(learningViewModelProvider);
    final viewModel = ref.read(learningViewModelProvider.notifier);

    developer.log(
        '[LearningScreen] build called - problems count: ${problems.length}, currentIndex: ${viewState.currentIndex}, state: ${viewState.currentState}');

    if (problems.isEmpty) {
      developer.log('[LearningScreen] No problems available');
      return const Center(child: Text('お疲れ様でした！'));
    }

    // コントローラーが初期化されていない場合は初期化
    if (viewState.answerControllers.isEmpty) {
      WidgetsBinding.instance.addPostFrameCallback((_) {
        viewModel.initializeControllersForCurrentProblem(null, problems);
      });
      return const Scaffold(
        body: Center(child: CircularProgressIndicator()),
      );
    }

    // 状態に基づいてUIを返す
    switch (viewState.currentState) {
      case LearningState.completed:
        return _buildCompletedScreen();

      case LearningState.answerDisplay:
        return _buildAnswerDisplayScreen(problems, viewState, viewModel);

      case LearningState.problemDisplay:
        return _buildProblemDisplayScreen(problems, viewState, viewModel);
    }
  }

  Widget _buildCompletedScreen() {
    return Scaffold(
      appBar: AppBar(
        title: const Text('単語学習'),
      ),
      body: const Center(
        child: Text(
          'お疲れ様でした！\n全問正解です！',
          textAlign: TextAlign.center,
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }

  Widget _buildAnswerDisplayScreen(List<WordProblem> problems,
      LearningViewState viewState, LearningViewModel viewModel) {
    return AnswerDisplayScreen(
      problem: problems[viewState.currentIndex],
      currentIndex: viewState.currentIndex,
      answerControllers: viewState.answerControllers,
      answerFocusNodes: viewState.answerFocusNodes,
      onNextProblem: viewModel.transitionToNextProblem,
    );
  }

  Widget _buildProblemDisplayScreen(List<WordProblem> problems,
      LearningViewState viewState, LearningViewModel viewModel) {
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
            viewModel.initializeControllersForCurrentProblem(null, problems),
      ),
    );
  }
}

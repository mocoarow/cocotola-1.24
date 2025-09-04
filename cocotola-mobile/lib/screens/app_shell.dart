import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../view_models/app_state_manager.dart';
import 'problem_set_selection_screen_unified.dart';
import 'learning_screen_unified.dart';
import 'answer_display_screen_unified.dart';
import 'completion_screen_unified.dart';

/// アプリ全体の状態に基づいて適切な画面を表示するシェル
class AppShell extends ConsumerWidget {
  const AppShell({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final appState = ref.watch(appStateProvider);

    switch (appState.learningState) {
      case LearningPhase.problemSelection:
        return const ProblemSetSelectionScreenUnified();
      
      case LearningPhase.problemDisplay:
        return const LearningScreenUnified();
      
      case LearningPhase.answerDisplay:
        return const AnswerDisplayScreenUnified();
      
      case LearningPhase.completed:
        return const CompletionScreenUnified();
    }
  }
}
import 'package:flutter/material.dart';
import 'word_problem.dart';

/// ProblemDisplayScreenの設定を管理するクラス
class ProblemDisplayConfig {
  final List<WordProblem> problems;
  final WordProblem problem;
  final int currentIndex;
  final List<TextEditingController> answerControllers;
  final List<FocusNode> answerFocusNodes;

  const ProblemDisplayConfig({
    required this.problems,
    required this.problem,
    required this.currentIndex,
    required this.answerControllers,
    required this.answerFocusNodes,
  });
}

/// ProblemDisplayScreenのコールバック関数を管理するクラス
class ProblemDisplayCallbacks {
  final void Function(int blankIndex, String value) onAnswerChanged;
  final VoidCallback onAllBlanksCompleted; // 全空欄完了通知
  final VoidCallback onShowAnswer;
  final VoidCallback onNextProblem;
  final void Function() onInitializeControllers;

  const ProblemDisplayCallbacks({
    required this.onAnswerChanged,
    required this.onAllBlanksCompleted,
    required this.onShowAnswer,
    required this.onNextProblem,
    required this.onInitializeControllers,
  });
}

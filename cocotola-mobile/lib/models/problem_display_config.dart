import 'package:flutter/material.dart';
import 'word_problem.dart';

/// ProblemDisplayScreenの設定を管理するクラス
class ProblemDisplayConfig {
  final List<WordProblem> problems;
  final int currentIndex;
  final List<TextEditingController> answerControllers;
  final List<FocusNode> answerFocusNodes;
  final int currentBlankIndex;
  final int cursorPosition;

  const ProblemDisplayConfig({
    required this.problems,
    required this.currentIndex,
    required this.answerControllers,
    required this.answerFocusNodes,
    required this.currentBlankIndex,
    required this.cursorPosition,
  });
}

/// ProblemDisplayScreenのコールバック関数を管理するクラス
class ProblemDisplayCallbacks {
  final void Function(int blankIndex, String value) onAnswerChanged;
  final void Function(int blankIndex) onBlankTap;
  final VoidCallback onCheckAnswers;
  final VoidCallback onShowAnswer;
  final VoidCallback onNextProblem;
  final void Function(String) onKeyPressed;
  final VoidCallback onDeleteKey;
  final VoidCallback onMoveLeft;
  final VoidCallback onMoveRight;
  final void Function(List<WordProblem>) onInitializeControllers;

  const ProblemDisplayCallbacks({
    required this.onAnswerChanged,
    required this.onBlankTap,
    required this.onCheckAnswers,
    required this.onShowAnswer,
    required this.onNextProblem,
    required this.onKeyPressed,
    required this.onDeleteKey,
    required this.onMoveLeft,
    required this.onMoveRight,
    required this.onInitializeControllers,
  });
}
import 'package:flutter/material.dart';
import 'word_problem.dart';

/// ProblemDisplayScreenの設定を管理するクラス
class ProblemDisplayConfig {
  final List<WordProblem> problems;
  final int currentIndex;
  final List<TextEditingController> answerControllers;
  final List<FocusNode> answerFocusNodes;
  final int initialBlankIndex;
  final int initialCursorPosition;

  const ProblemDisplayConfig({
    required this.problems,
    required this.currentIndex,
    required this.answerControllers,
    required this.answerFocusNodes,
    this.initialBlankIndex = 0,
    this.initialCursorPosition = 0,
  });
}

/// ProblemDisplayScreenのコールバック関数を管理するクラス
class ProblemDisplayCallbacks {
  final void Function(int blankIndex, String value) onAnswerChanged;
  final void Function(int blankIndex, String value) onAnswerChangedForAutoCheck; // 自動チェック用
  final void Function(int blankIndex) onBlankTap;
  final void Function(int blankIndex) onBlankIndexChanged; // 現在の空欄インデックス変更通知
  final VoidCallback onCheckAnswers;
  final VoidCallback onShowAnswer;
  final VoidCallback onNextProblem;
  final void Function(List<WordProblem>) onInitializeControllers;

  const ProblemDisplayCallbacks({
    required this.onAnswerChanged,
    required this.onAnswerChangedForAutoCheck,
    required this.onBlankTap,
    required this.onBlankIndexChanged,
    required this.onCheckAnswers,
    required this.onShowAnswer,
    required this.onNextProblem,
    required this.onInitializeControllers,
  });
}
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../models/word_problem.dart';
import 'problem_display_screen.dart';
import 'dart:developer' as developer;

class ProblemDisplayWrapper extends ConsumerStatefulWidget {
  final List<WordProblem> problems;
  final int currentIndex;
  final int currentBlankIndex;
  final int cursorPosition;
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
  final List<TextEditingController> answerControllers;
  final List<FocusNode> answerFocusNodes;

  const ProblemDisplayWrapper({
    super.key,
    required this.problems,
    required this.currentIndex,
    required this.currentBlankIndex,
    required this.cursorPosition,
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
    required this.answerControllers,
    required this.answerFocusNodes,
  });

  @override
  ConsumerState<ProblemDisplayWrapper> createState() => _ProblemDisplayWrapperState();
}

class _ProblemDisplayWrapperState extends ConsumerState<ProblemDisplayWrapper> {
  @override
  Widget build(BuildContext context) {
    // コントローラーが初期化されていない場合は初期化を実行
    final currentProblem = widget.problems[widget.currentIndex];
    final requiredBlanks = currentProblem.blanks.length;
    
    if (widget.answerControllers.isEmpty || widget.answerControllers.length < requiredBlanks) {
      developer.log('[ProblemDisplayWrapper] Controllers not ready, initializing...');
      widget.onInitializeControllers(widget.problems);
      return const Scaffold(
        body: Center(child: CircularProgressIndicator()),
      );
    }

    return ProblemDisplayScreen(
      problems: widget.problems,
      currentIndex: widget.currentIndex,
      answerControllers: widget.answerControllers,
      answerFocusNodes: widget.answerFocusNodes,
      currentBlankIndex: widget.currentBlankIndex,
      cursorPosition: widget.cursorPosition,
      onAnswerChanged: widget.onAnswerChanged,
      onBlankTap: widget.onBlankTap,
      onCheckAnswers: widget.onCheckAnswers,
      onShowAnswer: widget.onShowAnswer,
      onNextProblem: widget.onNextProblem,
      onKeyPressed: widget.onKeyPressed,
      onDeleteKey: widget.onDeleteKey,
      onMoveLeft: widget.onMoveLeft,
      onMoveRight: widget.onMoveRight,
    );
  }
}
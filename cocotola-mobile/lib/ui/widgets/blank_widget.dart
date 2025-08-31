import 'package:flutter/material.dart';
import '../../models/word_problem.dart';

class BlankWidget extends StatelessWidget {
  final int wordIndex;
  final int blankIndex;
  final WordProblem problem;
  final TextEditingController controller;
  final FocusNode focusNode;
  final void Function(String) onChanged;
  final VoidCallback onTap;

  const BlankWidget({
    super.key,
    required this.wordIndex,
    required this.blankIndex,
    required this.problem,
    required this.controller,
    required this.focusNode,
    required this.onChanged,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    final blank = problem.blanks[blankIndex];
    final inputWidth = (blank.answer.length * 20.0).clamp(100.0, 200.0);

    // プロバイダーの状態とコントローラーが同期していない場合は更新
    if (controller.text != blank.userInput) {
      controller.text = blank.userInput;
    }

    return Container(
      margin: const EdgeInsets.symmetric(horizontal: 4.0, vertical: 2.0),
      width: inputWidth,
      child: blank.isAnswered && blank.isCorrect
          ? Container(
              padding: const EdgeInsets.all(8.0),
              decoration: BoxDecoration(
                border: Border.all(color: Colors.green, width: 2),
                borderRadius: BorderRadius.circular(4),
                color: Colors.green.shade50,
              ),
              child: Text(
                blank.answer,
                textAlign: TextAlign.center,
                style: const TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.bold,
                  color: Colors.green,
                ),
              ),
            )
          : TextField(
              controller: controller,
              textAlign: TextAlign.center,
              decoration: InputDecoration(
                border: const OutlineInputBorder(),
                fillColor: blank.isAnswered && !blank.isCorrect
                    ? Colors.red.shade50
                    : null,
                filled: blank.isAnswered && !blank.isCorrect,
              ),
              focusNode: focusNode,
              enabled: !(blank.isAnswered && blank.isCorrect),
              onChanged: onChanged,
              onTap: onTap,
            ),
    );
  }
}
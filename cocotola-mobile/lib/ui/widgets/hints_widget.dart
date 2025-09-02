import 'package:flutter/material.dart';
import '../../models/word_problem.dart';

class HintsWidget extends StatelessWidget {
  final WordProblem problem;

  const HintsWidget({
    super.key,
    required this.problem,
  });

  @override
  Widget build(BuildContext context) {
    final correctBlanks = problem.blanks
        .where((blank) => blank.isAnswered && blank.isCorrect)
        .toList();

    if (correctBlanks.isEmpty) {
      return const SizedBox.shrink();
    }

    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        children: correctBlanks
            .map((blank) => Container(
                  margin: const EdgeInsets.only(bottom: 8.0),
                  padding: const EdgeInsets.all(12.0),
                  decoration: BoxDecoration(
                    color: Colors.green.shade50,
                    borderRadius: BorderRadius.circular(8),
                    border: Border.all(color: Colors.green.shade200),
                  ),
                  child: Column(
                    children: [
                      Text(
                        '正解: ${blank.answer}',
                        style: const TextStyle(
                          fontSize: 16,
                          fontWeight: FontWeight.bold,
                          color: Colors.green,
                        ),
                      ),
                      const SizedBox(height: 4),
                      Text(
                        'ヒント: ${blank.hint}',
                        style: const TextStyle(fontSize: 14),
                      ),
                    ],
                  ),
                ))
            .toList(),
      ),
    );
  }
}
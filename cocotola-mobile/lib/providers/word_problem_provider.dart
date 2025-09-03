import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../models/word_problem.dart';
import 'dart:developer' as developer;

final wordProblemsProvider =
    StateNotifierProvider<WordProblemsNotifier, List<WordProblem>>((ref) {
  return WordProblemsNotifier();
});

class WordProblemsNotifier extends StateNotifier<List<WordProblem>> {
  WordProblemsNotifier()
      : super([
          WordProblem(
            japanese: '私は毎日英語を勉強します。',
            english: 'I ___ English every day.',
            cefrLevel: 'A1',
            blanks: [
              BlankAnswer(
                answer: 'study',
                hint: '「勉強する」という意味の動詞です。',
              ),
            ],
          ),
          WordProblem(
            japanese: '彼女は新しい本を買いました。',
            english: 'She ___ a new book.',
            cefrLevel: 'A2',
            blanks: [
              BlankAnswer(
                answer: 'bought',
                hint: '「買う」の過去形です。',
              ),
            ],
          ),
          WordProblem(
            japanese: 'この問題は難しいです。',
            english: 'This problem is ___.',
            cefrLevel: 'B1',
            blanks: [
              BlankAnswer(
                answer: 'difficult',
                hint: '「難しい」という意味の形容詞です。',
              ),
            ],
          ),
          WordProblem(
            japanese: '私は彼女に図書館で会った。',
            english: 'I ___ her ___ the library.',
            cefrLevel: 'B1',
            blanks: [
              BlankAnswer(
                answer: 'met',
                hint: '「会う」の過去形です。',
              ),
              BlankAnswer(
                answer: 'at',
                hint: '場所を示す前置詞です。',
              ),
            ],
          ),
          WordProblem(
            japanese: '彼は英語を上手に話します。',
            english: 'He speaks English ___.',
            cefrLevel: 'A2',
            blanks: [
              BlankAnswer(
                answer: 'well',
                hint: '「上手に」という意味の副詞です。',
              ),
            ],
          ),
        ]);

  void checkAnswer(int problemIndex, int blankIndex, String userInput) {
    final problem = state[problemIndex];
    final blank = problem.blanks[blankIndex];

    final isCorrect =
        userInput.toLowerCase().trim() == blank.answer.toLowerCase().trim();

    final updatedBlank = blank.copyWith(
      userInput: userInput,
      isAnswered: isCorrect,
      isCorrect: isCorrect,
    );

    final updatedProblem = problem.updateBlank(blankIndex, updatedBlank);

    state = [
      for (int i = 0; i < state.length; i++)
        if (i == problemIndex) updatedProblem else state[i]
    ];
  }

  void markAsSkipped(int index) {
    state = [
      for (int i = 0; i < state.length; i++)
        if (i == index) state[i].copyWith(isSkipped: true) else state[i]
    ];
  }

  void updateUserInput(int problemIndex, int blankIndex, String input) {
    final problem = state[problemIndex];
    final blank = problem.blanks[blankIndex];

    if (blank.isAnswered && blank.isCorrect) {
      return; // 正解済みの場合は更新しない
    }

    final updatedBlank = blank.copyWith(userInput: input);
    final updatedProblem = problem.updateBlank(blankIndex, updatedBlank);

    state = [
      for (int i = 0; i < state.length; i++)
        if (i == problemIndex) updatedProblem else state[i]
    ];
  }

  void clearUserInputs(int problemIndex) {
    developer.log(
        '[WordProblemsProvider] clearUserInputs called for problem $problemIndex');
    final problem = state[problemIndex];
    developer.log(
        '[WordProblemsProvider] Before clear - blanks count: ${problem.blanks.length}');

    for (int i = 0; i < problem.blanks.length; i++) {
      developer.log(
          '[WordProblemsProvider] Blank[$i] before: userInput="${problem.blanks[i].userInput}", isAnswered=${problem.blanks[i].isAnswered}');
    }

    final clearedBlanks = problem.blanks
        .map((blank) => blank.copyWith(
              userInput: '',
              isAnswered: false,
              isCorrect: false,
            ))
        .toList();

    final updatedProblem = problem.copyWith(blanks: clearedBlanks);

    state = [
      for (int i = 0; i < state.length; i++)
        if (i == problemIndex) updatedProblem else state[i]
    ];

    developer.log('[WordProblemsProvider] After clear - state updated');
    for (int i = 0; i < state[problemIndex].blanks.length; i++) {
      developer.log(
          '[WordProblemsProvider] Blank[$i] after: userInput="${state[problemIndex].blanks[i].userInput}", isAnswered=${state[problemIndex].blanks[i].isAnswered}');
    }
  }
}

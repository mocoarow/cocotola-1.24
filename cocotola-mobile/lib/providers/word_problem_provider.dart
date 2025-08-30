import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../models/word_problem.dart';

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

    if (blank.isAnswered && blank.isCorrect) return; // 正解済みの場合は更新しない

    final updatedBlank = blank.copyWith(userInput: input);
    final updatedProblem = problem.updateBlank(blankIndex, updatedBlank);

    state = [
      for (int i = 0; i < state.length; i++)
        if (i == problemIndex) updatedProblem else state[i]
    ];
  }

  void clearUserInputs(int problemIndex) {
    final problem = state[problemIndex];
    
    final clearedBlanks = problem.blanks.map((blank) => blank.copyWith(
      userInput: '',
      isAnswered: false,
      isCorrect: false,
    )).toList();
    
    final updatedProblem = problem.copyWith(blanks: clearedBlanks);
    
    state = [
      for (int i = 0; i < state.length; i++)
        if (i == problemIndex) updatedProblem else state[i]
    ];
  }
}

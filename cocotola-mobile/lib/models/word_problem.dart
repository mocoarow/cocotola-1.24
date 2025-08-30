class BlankAnswer {
  final String answer;
  final String hint;
  bool isAnswered;
  bool isCorrect;
  String userInput;

  BlankAnswer({
    required this.answer,
    required this.hint,
    this.isAnswered = false,
    this.isCorrect = false,
    this.userInput = '',
  });

  BlankAnswer copyWith({
    String? answer,
    String? hint,
    bool? isAnswered,
    bool? isCorrect,
    String? userInput,
  }) {
    return BlankAnswer(
      answer: answer ?? this.answer,
      hint: hint ?? this.hint,
      isAnswered: isAnswered ?? this.isAnswered,
      isCorrect: isCorrect ?? this.isCorrect,
      userInput: userInput ?? this.userInput,
    );
  }
}

class WordProblem {
  final String japanese;
  final String english;
  final List<BlankAnswer> blanks;
  bool isSkipped;

  WordProblem({
    required this.japanese,
    required this.english,
    required this.blanks,
    this.isSkipped = false,
  });

  bool get isCompleted => blanks.every((blank) => blank.isAnswered && blank.isCorrect);
  bool get hasAnyCorrectAnswer => blanks.any((blank) => blank.isAnswered && blank.isCorrect);

  WordProblem copyWith({
    String? japanese,
    String? english,
    List<BlankAnswer>? blanks,
    bool? isSkipped,
  }) {
    return WordProblem(
      japanese: japanese ?? this.japanese,
      english: english ?? this.english,
      blanks: blanks ?? this.blanks.map((blank) => blank.copyWith()).toList(),
      isSkipped: isSkipped ?? this.isSkipped,
    );
  }

  WordProblem updateBlank(int index, BlankAnswer newBlank) {
    final updatedBlanks = List<BlankAnswer>.from(blanks);
    updatedBlanks[index] = newBlank;
    return copyWith(blanks: updatedBlanks);
  }
}

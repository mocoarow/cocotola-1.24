import 'word_problem.dart';
import 'memorization_problem.dart';

/// 問題の種類
enum ProblemType {
  word,        // 英単語穴埋め問題
  memorization // 暗記問題
}

/// 問題の基底クラス
abstract class ProblemBase {
  String get cefrLevel;
  bool get isCompleted;
  ProblemType get problemType;
}

/// WordProblemにProblemBaseを実装
extension WordProblemExtension on WordProblem {
  ProblemType get problemType => ProblemType.word;
}

/// MemorizationProblemにProblemBaseを実装  
extension MemorizationProblemExtension on MemorizationProblem {
  ProblemType get problemType => ProblemType.memorization;
}

/// 統合された問題クラス
class Problem {
  final WordProblem? wordProblem;
  final MemorizationProblem? memorizationProblem;
  
  const Problem.word(this.wordProblem) : memorizationProblem = null;
  const Problem.memorization(this.memorizationProblem) : wordProblem = null;
  
  ProblemType get type {
    if (wordProblem != null) return ProblemType.word;
    if (memorizationProblem != null) return ProblemType.memorization;
    throw Exception('Invalid problem state');
  }
  
  String get cefrLevel {
    switch (type) {
      case ProblemType.word:
        return wordProblem!.cefrLevel;
      case ProblemType.memorization:
        return memorizationProblem!.cefrLevel;
    }
  }
  
  bool get isCompleted {
    switch (type) {
      case ProblemType.word:
        return wordProblem!.isCompleted;
      case ProblemType.memorization:
        return memorizationProblem!.isCompleted;
    }
  }
  
  Problem copyWith({
    WordProblem? wordProblem,
    MemorizationProblem? memorizationProblem,
  }) {
    switch (type) {
      case ProblemType.word:
        return Problem.word(wordProblem ?? this.wordProblem!);
      case ProblemType.memorization:
        return Problem.memorization(memorizationProblem ?? this.memorizationProblem!);
    }
  }
}
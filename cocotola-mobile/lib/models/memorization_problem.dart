/// 暗記系問題のモデル
class MemorizationProblem {
  final String question; // 問題文
  final String answer; // 答え
  final String? hint; // ヒント（オプション）
  final String cefrLevel; // CEFRレベル
  bool isAnswered; // 回答済みかどうか
  bool isCorrect; // 正解かどうか（できた/できなかった）
  
  MemorizationProblem({
    required this.question,
    required this.answer,
    this.hint,
    this.cefrLevel = 'A1',
    this.isAnswered = false,
    this.isCorrect = false,
  });

  /// 問題が完了したかどうか
  bool get isCompleted => isAnswered;

  MemorizationProblem copyWith({
    String? question,
    String? answer,
    String? hint,
    String? cefrLevel,
    bool? isAnswered,
    bool? isCorrect,
  }) {
    return MemorizationProblem(
      question: question ?? this.question,
      answer: answer ?? this.answer,
      hint: hint ?? this.hint,
      cefrLevel: cefrLevel ?? this.cefrLevel,
      isAnswered: isAnswered ?? this.isAnswered,
      isCorrect: isCorrect ?? this.isCorrect,
    );
  }

  /// ユーザーが「できた」「できなかった」を選択
  MemorizationProblem markAsAnswered(bool wasCorrect) {
    return copyWith(
      isAnswered: true,
      isCorrect: wasCorrect,
    );
  }
}
import 'package:flutter/material.dart';

/// 暗記系問題のモデル
class MemorizationProblem {
  final String question; // 問題文
  final String answer; // 答え
  final String? hint; // ヒント（オプション）
  final String cefrLevel; // CEFRレベル
  final TextAlign questionAlignment; // 問題文のテキスト配置（中央寄せ・左寄せ）
  final TextAlign answerAlignment; // 答えのテキスト配置（中央寄せ・左寄せ）
  bool isAnswered; // 回答済みかどうか
  bool isCorrect; // 正解かどうか（できた/できなかった）
  
  MemorizationProblem({
    required this.question,
    required this.answer,
    this.hint,
    this.cefrLevel = 'A1',
    this.questionAlignment = TextAlign.center, // デフォルトは中央寄せ
    this.answerAlignment = TextAlign.center, // デフォルトは中央寄せ
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
    TextAlign? questionAlignment,
    TextAlign? answerAlignment,
    bool? isAnswered,
    bool? isCorrect,
  }) {
    return MemorizationProblem(
      question: question ?? this.question,
      answer: answer ?? this.answer,
      hint: hint ?? this.hint,
      cefrLevel: cefrLevel ?? this.cefrLevel,
      questionAlignment: questionAlignment ?? this.questionAlignment,
      answerAlignment: answerAlignment ?? this.answerAlignment,
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
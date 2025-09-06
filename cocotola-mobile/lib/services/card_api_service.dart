import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/problem_base.dart';
import '../models/word_problem.dart';

class CardApiService {
  static const String baseUrl = 'http://localhost:8000';
  static const String cardEndpoint = '/core/api/v1/card';

  /// カード問題をAPIから取得する
  Future<List<Problem>> fetchCards() async {
    try {
      final response = await http.get(
        Uri.parse('$baseUrl$cardEndpoint'),
        headers: {
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final data = json.decode(response.body);
        return _parseProblemsFromResponse(data);
      } else {
        throw Exception('Failed to fetch cards: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to fetch cards: $e');
    }
  }

  /// APIレスポンスからProblemリストに変換
  List<Problem> _parseProblemsFromResponse(dynamic data) {
    final List<Problem> problems = [];
    
    // APIレスポンスの構造に応じて適切にパースする
    if (data is Map<String, dynamic> && data.containsKey('results')) {
      final results = data['results'] as List;
      for (final card in results) {
        final problem = _parseCardToProblem(card);
        if (problem != null) {
          problems.add(problem);
        }
      }
    } else if (data is Map<String, dynamic> && data.containsKey('cards')) {
      final cards = data['cards'] as List;
      for (final card in cards) {
        final problem = _parseCardToProblem(card);
        if (problem != null) {
          problems.add(problem);
        }
      }
    } else if (data is List) {
      for (final card in data) {
        final problem = _parseCardToProblem(card);
        if (problem != null) {
          problems.add(problem);
        }
      }
    }
    
    return problems;
  }

  /// 1つのカードデータをProblemに変換
  Problem? _parseCardToProblem(dynamic cardData) {
    if (cardData is! Map<String, dynamic>) return null;

    try {
      // nameフィールド内のJSON文字列を解析
      final nameJson = cardData['name'] as String?;
      if (nameJson == null) return null;
      
      final nameData = json.decode(nameJson) as Map<String, dynamic>;
      
      // 英単語問題として処理
      return Problem.word(WordProblem(
        japanese: nameData['Japanese'] as String? ?? '',
        english: nameData['English'] as String? ?? '',
        cefrLevel: nameData['CEFRLevel'] as String? ?? 'A1',
        blanks: _parseBlanks(nameData['Blanks']),
      ));
      
    } catch (e) {
      // パースエラーの場合はnullを返す
      return null;
    }
  }

  /// blanks配列をBlankAnswerリストに変換
  List<BlankAnswer> _parseBlanks(dynamic blanksData) {
    if (blanksData is! List) return [];
    
    return blanksData.map<BlankAnswer>((blankData) {
      if (blankData is Map<String, dynamic>) {
        return BlankAnswer(
          answer: blankData['Answer'] as String? ?? blankData['answer'] as String? ?? '',
          hint: blankData['Hint'] as String? ?? blankData['hint'] as String? ?? '',
        );
      }
      return BlankAnswer(answer: '', hint: '');
    }).toList();
  }
}
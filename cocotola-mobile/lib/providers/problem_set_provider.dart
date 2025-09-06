import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../models/problem_set.dart';
import '../models/word_problem.dart';
import '../models/memorization_problem.dart';
import '../models/problem_base.dart';
import 'card_provider.dart';

final problemSetsProvider = StateNotifierProvider<ProblemSetsNotifier, List<ProblemSet>>((ref) {
  return ProblemSetsNotifier(ref);
});

class ProblemSetsNotifier extends StateNotifier<List<ProblemSet>> {
  final Ref ref;
  
  ProblemSetsNotifier(this.ref) : super(_generateProblemSets());
  
  static List<ProblemSet> _generateProblemSets() {
    return [
      ProblemSet(
        id: 'beginner-basics',
        title: '初心者向け基本文法',
        description: '日常会話で使用する基本的な英語表現を学習します。',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: '私は毎日英語を勉強します。',
            english: 'I ___ English every day.',
            cefrLevel: 'A1',
            blanks: [
              BlankAnswer(
                answer: 'study',
                hint: '「勉強する」という意味の動詞です。',
              ),
            ],
          )),
          Problem.word(WordProblem(
            japanese: '彼は英語を上手に話します。',
            english: 'He speaks English ___.',
            cefrLevel: 'A2',
            blanks: [
              BlankAnswer(
                answer: 'well',
                hint: '「上手に」という意味の副詞です。',
              ),
            ],
          )),
        ],
      ),
      ProblemSet(
        id: 'elementary-shopping',
        title: 'ショッピングの英語',
        description: '買い物や日常の購買活動で使用する英語表現を身につけます。',
        cefrLevel: 'A2',
        problems: [
          Problem.word(WordProblem(
            japanese: '彼女は新しい本を買いました。',
            english: 'She ___ a new book.',
            cefrLevel: 'A2',
            blanks: [
              BlankAnswer(
                answer: 'bought',
                hint: '「買う」の過去形です。',
              ),
            ],
          )),
        ],
      ),
      ProblemSet(
        id: 'intermediate-grammar',
        title: '中級文法チャレンジ',
        description: 'より複雑な文法構造と表現を学習し、スキルアップを目指します。',
        cefrLevel: 'B1',
        problems: [
          Problem.word(WordProblem(
            japanese: 'この問題は難しいです。',
            english: 'This problem is ___.',
            cefrLevel: 'B1',
            blanks: [
              BlankAnswer(
                answer: 'difficult',
                hint: '「難しい」という意味の形容詞です。',
              ),
            ],
          )),
          Problem.word(WordProblem(
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
          )),
        ],
      ),
      // 暗記問題のサンプル
      ProblemSet(
        id: 'memorization-vocabulary',
        title: '単語暗記',
        description: '基本的な英単語を暗記しましょう。',
        cefrLevel: 'A1',
        problems: [
          Problem.memorization(MemorizationProblem(
            question: 'apple',
            answer: 'りんご',
            hint: '赤い果物です',
            cefrLevel: 'A1',
          )),
          Problem.memorization(MemorizationProblem(
            question: 'book',
            answer: '本',
            hint: '読むものです',
            cefrLevel: 'A1',
          )),
          Problem.memorization(MemorizationProblem(
            question: 'water',
            answer: '水',
            hint: '飲み物です',
            cefrLevel: 'A1',
          )),
        ],
      ),
      ProblemSet(
        id: 'memorization-phrases',
        title: 'フレーズ暗記',
        description: '日常会話で使える基本フレーズを覚えましょう。',
        cefrLevel: 'A2',
        problems: [
          Problem.memorization(MemorizationProblem(
            question: 'Good morning',
            answer: 'おはようございます',
            cefrLevel: 'A1',
          )),
          Problem.memorization(MemorizationProblem(
            question: 'Thank you very much',
            answer: 'どうもありがとうございます',
            cefrLevel: 'A1',
          )),
          Problem.memorization(MemorizationProblem(
            question: 'How are you?',
            answer: '元気ですか？',
            cefrLevel: 'A1',
          )),
        ],
      ),
    ];
  }
  
  ProblemSet? getProblemSetById(String id) {
    try {
      return state.firstWhere((set) => set.id == id);
    } catch (e) {
      return null;
    }
  }

  /// APIから取得したカード問題を問題セットとして追加
  Future<void> addProblemSetFromApi() async {
    try {
      final cardProblems = await ref.read(cardProblemsProvider.future);
      
      if (cardProblems.isNotEmpty) {
        final apiProblemSet = ProblemSet(
          id: 'api-cards',
          title: 'APIカード問題セット',
          description: 'サーバーから取得したカード問題を学習します。',
          cefrLevel: 'Mixed', // 複数レベルが混在する可能性があるため
          problems: cardProblems,
        );
        
        // 既存の問題セットと結合（APIセットが既に存在する場合は更新）
        final updatedState = [...state];
        final existingIndex = updatedState.indexWhere((set) => set.id == 'api-cards');
        
        if (existingIndex != -1) {
          updatedState[existingIndex] = apiProblemSet;
        } else {
          updatedState.add(apiProblemSet);
        }
        
        state = updatedState;
      } else {
        throw Exception('APIから問題を取得できませんでした');
      }
    } catch (e) {
      // エラーを再スローして上位で処理
      rethrow;
    }
  }

  /// API問題セットを削除
  void removeApiProblemSet() {
    state = state.where((set) => set.id != 'api-cards').toList();
  }
}
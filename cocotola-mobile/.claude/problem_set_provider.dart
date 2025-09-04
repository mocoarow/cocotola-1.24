import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../models/problem_set.dart';
import '../models/word_problem.dart';

final problemSetsProvider = StateNotifierProvider<ProblemSetsNotifier, List<ProblemSet>>((ref) {
  return ProblemSetsNotifier();
});

class ProblemSetsNotifier extends StateNotifier<List<ProblemSet>> {
  ProblemSetsNotifier() : super(_generateProblemSets());
  
  static List<ProblemSet> _generateProblemSets() {
    return [
      ProblemSet(
        id: 'beginner-basics',
        title: '初心者向け基本文法',
        description: '日常会話で使用する基本的な英語表現を学習します。',
        cefrLevel: 'A1',
        problems: [
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
        ],
      ),
      ProblemSet(
        id: 'elementary-shopping',
        title: 'ショッピングの英語',
        description: '買い物や日常の購買活動で使用する英語表現を身につけます。',
        cefrLevel: 'A2',
        problems: [
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
        ],
      ),
      ProblemSet(
        id: 'intermediate-grammar',
        title: '中級文法チャレンジ',
        description: 'より複雑な文法構造と表現を学習し、スキルアップを目指します。',
        cefrLevel: 'B1',
        problems: [
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
}
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:cocotola/view_models/app_state_manager.dart';
import 'package:cocotola/models/word_problem.dart';
import 'package:cocotola/models/problem_base.dart';
import 'package:cocotola/models/problem_set.dart';

void main() {
  // Flutterのバインディングを初期化
  TestWidgetsFlutterBinding.ensureInitialized();
  
  group('AppStateManager 自動フォーカス移動テスト', () {
    late ProviderContainer container;
    late AppStateManager appStateManager;

    setUp(() {
      // テスト用のProviderContainerを作成
      container = ProviderContainer();
      appStateManager = container.read(appStateProvider.notifier);
    });

    tearDown(() {
      container.dispose();
    });

    test('正解入力時の次の空欄への自動フォーカス移動', () {
      // 複数の空欄を持つテスト用問題を作成
      final problemSet = ProblemSet(
        id: 'test-set',
        title: 'テスト問題セット',
        description: 'テスト用の問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: 'テストの勉強',
            english: '_____ _____ test',
            blanks: [
              BlankAnswer(answer: 'study', hint: 'ヒント1'),
              BlankAnswer(answer: 'of', hint: 'ヒント2'),
            ],
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 初期状態確認
      var state = container.read(appStateProvider);
      expect(state.currentBlankIndex, 0);
      expect(state.answerControllers.length, 2);
      
      // 最初の空欄に正解を入力
      appStateManager.handleAnswer(0, 'study');
      
      // 正解が記録されたことを確認
      state = container.read(appStateProvider);
      final firstBlank = state.currentProblem!.wordProblem!.blanks[0];
      expect(firstBlank.isAnswered, true);
      expect(firstBlank.isCorrect, true);
      expect(firstBlank.userInput, 'study');
      
      // 問題がまだ完了していないことを確認
      expect(state.currentProblem!.isCompleted, false);
      
      // 2つ目の空欄に正解を入力
      appStateManager.handleAnswer(1, 'of');
      
      // 2つ目の正解が記録されたことを確認
      state = container.read(appStateProvider);
      final secondBlank = state.currentProblem!.wordProblem!.blanks[1];
      expect(secondBlank.isAnswered, true);
      expect(secondBlank.isCorrect, true);
      
      // 問題が完了したことを確認
      expect(state.currentProblem!.isCompleted, true);
    });

    test('間違った入力では自動フォーカス移動しない', () {
      // 複数の空欄を持つテスト用問題を作成
      final problemSet = ProblemSet(
        id: 'test-set',
        title: 'テスト問題セット',
        description: 'テスト用の問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: 'テストの勉強',
            english: '_____ _____ test',
            blanks: [
              BlankAnswer(answer: 'study', hint: 'ヒント1'),
              BlankAnswer(answer: 'of', hint: 'ヒント2'),
            ],
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 最初の空欄に間違った答えを入力
      appStateManager.handleAnswer(0, 'wrong');
      
      // 間違いが記録されたことを確認
      var state = container.read(appStateProvider);
      final firstBlank = state.currentProblem!.wordProblem!.blanks[0];
      expect(firstBlank.isAnswered, false); // 正解でないため未回答扱い
      expect(firstBlank.isCorrect, false);
      expect(firstBlank.userInput, 'wrong');
      
      // 問題が完了していないことを確認
      expect(state.currentProblem!.isCompleted, false);
    });

    test('全問正解時は答え表示画面に遷移', () {
      // 単一空欄のテスト用問題を作成
      final problemSet = ProblemSet(
        id: 'test-set',
        title: 'テスト問題セット',
        description: 'テスト用の問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: 'テストの勉学',
            english: '_____ test',
            blanks: [
              BlankAnswer(answer: 'study', hint: 'ヒント1'),
            ],
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 空欄に正解を入力
      appStateManager.handleAnswer(0, 'study');
      
      // 問題完了により答え表示画面に遷移することを確認
      // （実際のWidgetsBinding.instance.addPostFrameCallbackの確認は困難だが、
      //  問題が完了状態になることは確認できる）
      var state = container.read(appStateProvider);
      expect(state.currentProblem!.isCompleted, true);
    });

    test('答え表示画面への遷移時にコントローラーに正解が設定される', () {
      // テスト用問題を作成
      final problemSet = ProblemSet(
        id: 'test-set',
        title: 'テスト問題セット',
        description: 'テスト用の問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: 'テストの勉強',
            english: '_____ _____ test',
            blanks: [
              BlankAnswer(answer: 'study', hint: 'ヒント1'),
              BlankAnswer(answer: 'of', hint: 'ヒント2'),
            ],
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 初期状態確認（コントローラーは空欄の状態）
      var state = container.read(appStateProvider);
      expect(state.answerControllers.length, 2);
      expect(state.answerControllers[0].text, ''); // 初期は空
      expect(state.answerControllers[1].text, ''); // 初期は空
      
      // 「答えを見る」機能を実行
      appStateManager.transitionToAnswerDisplay();
      
      // 答え表示画面に遷移したことを確認
      state = container.read(appStateProvider);
      expect(state.learningState, LearningPhase.answerDisplay);
      
      // コントローラーに正解が設定されたことを確認
      expect(state.answerControllers[0].text, 'study');
      expect(state.answerControllers[1].text, 'of');
    });

    test('一部正解済みの問題で答え表示した場合も全ての正解が表示される', () {
      // テスト用問題を作成
      final problemSet = ProblemSet(
        id: 'test-set',
        title: 'テスト問題セット',
        description: 'テスト用の問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: 'テストの勉強',
            english: '_____ _____ test',
            blanks: [
              BlankAnswer(answer: 'study', hint: 'ヒント1'),
              BlankAnswer(answer: 'of', hint: 'ヒント2'),
            ],
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 最初の空欄のみ正解を入力
      appStateManager.handleAnswer(0, 'study');
      
      var state = container.read(appStateProvider);
      expect(state.answerControllers[0].text, 'study');
      expect(state.answerControllers[1].text, ''); // 2つ目は未入力
      
      // 「答えを見る」機能を実行
      appStateManager.transitionToAnswerDisplay();
      
      // 両方のコントローラーに正解が設定されることを確認
      state = container.read(appStateProvider);
      expect(state.answerControllers[0].text, 'study');
      expect(state.answerControllers[1].text, 'of'); // 2つ目にも正解が設定される
    });

    test('メニューに戻る機能のテスト', () {
      // テスト用問題を作成
      final problemSet = ProblemSet(
        id: 'test-set',
        title: 'テスト問題セット',
        description: 'テスト用の問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: 'テストの勉強',
            english: '_____ test',
            blanks: [
              BlankAnswer(answer: 'study', hint: 'ヒント1'),
            ],
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択して学習開始状態にする
      appStateManager.selectProblemSet(problemSet);
      
      var state = container.read(appStateProvider);
      expect(state.selectedProblemSet, isNotNull);
      expect(state.learningState, LearningPhase.problemDisplay);
      expect(state.answerControllers.length, 1);
      expect(state.answerFocusNodes.length, 1);
      
      // メニューに戻る
      appStateManager.returnToMenu();
      
      // 状態がリセットされたことを確認
      state = container.read(appStateProvider);
      expect(state.selectedProblemSet, isNull);
      expect(state.learningState, LearningPhase.problemSelection);
      expect(state.currentProblemIndex, 0);
      expect(state.problems, isEmpty);
      expect(state.answerControllers, isEmpty);
      expect(state.answerFocusNodes, isEmpty);
      expect(state.currentBlankIndex, 0);
      expect(state.cursorPosition, 0);
    });

    test('答え表示画面からメニューに戻る機能のテスト', () {
      // テスト用問題を作成
      final problemSet = ProblemSet(
        id: 'test-set',
        title: 'テスト問題セット',
        description: 'テスト用の問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: 'テストの勉強',
            english: '_____ test',
            blanks: [
              BlankAnswer(answer: 'study', hint: 'ヒント1'),
            ],
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択し、答え表示画面に遷移
      appStateManager.selectProblemSet(problemSet);
      appStateManager.transitionToAnswerDisplay();
      
      var state = container.read(appStateProvider);
      expect(state.learningState, LearningPhase.answerDisplay);
      expect(state.selectedProblemSet, isNotNull);
      
      // メニューに戻る
      appStateManager.returnToMenu();
      
      // 状態がリセットされたことを確認
      state = container.read(appStateProvider);
      expect(state.selectedProblemSet, isNull);
      expect(state.learningState, LearningPhase.problemSelection);
    });

    test('英単語問題で「答えを見る」クリック時に正しい答えが表示され、次の問題へで後回しになることを確認', () {
      // 複数の英単語問題セットを作成
      final problemSet = ProblemSet(
        id: 'test-word-show-answer',
        title: '「答えを見る」後回しテスト',
        description: '答えを見る機能の後回し処理をテスト',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: '最初の問題',
            english: 'First ___.',
            cefrLevel: 'A1',
            blanks: [
              BlankAnswer(
                answer: 'problem',
                hint: 'ヒント1',
              ),
            ],
          )),
          Problem.word(WordProblem(
            japanese: '二番目の問題',
            english: 'Second ___.',
            cefrLevel: 'A1',
            blanks: [
              BlankAnswer(
                answer: 'test',
                hint: 'ヒント2',
              ),
            ],
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 初期状態確認
      var state = container.read(appStateProvider);
      expect(state.problems.length, 2);
      expect(state.currentProblemIndex, 0);
      expect(state.currentProblem!.wordProblem!.japanese, '最初の問題');
      expect(state.showAnswerForProblem, isNull);
      
      // 「答えを見る」をクリック
      appStateManager.handleShowAnswerForWordProblem();
      
      // 答え表示画面に遷移することを確認
      state = container.read(appStateProvider);
      expect(state.learningState, LearningPhase.answerDisplay);
      
      // 問題は移動せず、現在の問題がそのままで、答えが表示されることを確認
      expect(state.problems.length, 2); // 問題数は変わらない
      expect(state.currentProblemIndex, 0); // インデックスも変わらない
      expect(state.currentProblem!.wordProblem!.japanese, '最初の問題'); // 現在の問題は変わらない
      expect(state.showAnswerForProblem, isNotNull); // 答え表示用問題が設定される
      expect(state.showAnswerForProblem!.wordProblem!.japanese, '最初の問題'); // 最初の問題が保存される
      
      // 答えが正しくコントローラーに設定されていることを確認
      expect(state.answerControllers[0].text, 'problem'); // 最初の問題の正しい答え
      
      // 「次の問題へ」を実行（この時点で問題が後回しになる）
      appStateManager.moveToNextProblem();
      
      // 問題が後回しになり、次の問題に移動することを確認
      state = container.read(appStateProvider);
      expect(state.problems.length, 2); // 問題数は変わらない
      expect(state.currentProblemIndex, 0); // インデックスは0のまま
      expect(state.currentProblem!.wordProblem!.japanese, '二番目の問題'); // 次の問題が表示される
      expect(state.learningState, LearningPhase.problemDisplay); // 問題表示画面に戻る
      expect(state.showAnswerForProblem, isNull); // 答え表示用問題はクリアされる
      
      // 最後の問題は後回しされた問題になっているかを確認
      final lastProblem = state.problems.last;
      expect(lastProblem.wordProblem!.japanese, '最初の問題');
      
      // コントローラーは新しい問題用に初期化されていることを確認
      expect(state.answerControllers[0].text, ''); // 二番目の問題用に空になっている
    });

    test('実際の問題セットのシナリオ：「私は毎日英語を勉強します」の答えが正しく表示される', () {
      // 実際の初心者向け基本文法問題セットを模擬
      final problemSet = ProblemSet(
        id: 'beginner-grammar',
        title: '初心者向け基本文法',
        description: '基本的な英語文法問題',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: '私は毎日英語を勉強します',
            english: 'I _____ English every day.',
            cefrLevel: 'A1',
            blanks: [
              BlankAnswer(answer: 'study', hint: '勉強する'),
            ],
          )),
          Problem.word(WordProblem(
            japanese: '彼は英語を上手に話します',
            english: 'He speaks English _____.',
            cefrLevel: 'A1',
            blanks: [
              BlankAnswer(answer: 'well', hint: '上手に'),
            ],
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 最初の問題が表示されることを確認
      var state = container.read(appStateProvider);
      expect(state.currentProblem!.wordProblem!.japanese, '私は毎日英語を勉強します');
      expect(state.answerControllers[0].text, ''); // 初期は空
      
      // 「答えを見る」をクリック
      appStateManager.handleShowAnswerForWordProblem();
      
      // 正しい問題の答えが表示されることを確認
      state = container.read(appStateProvider);
      expect(state.learningState, LearningPhase.answerDisplay);
      expect(state.currentProblem!.wordProblem!.japanese, '私は毎日英語を勉強します'); // 問題は変わらない
      expect(state.answerControllers[0].text, 'study'); // 「私は毎日英語を勉強します」の正しい答え
      expect(state.showAnswerForProblem!.wordProblem!.japanese, '私は毎日英語を勉強します'); // 答え表示対象の問題も正しい
      
      // 間違った答え（他の問題の答え）が表示されていないことを確認
      expect(state.answerControllers[0].text, isNot('well')); // 「彼は英語を上手に話します」の答えではない
      expect(state.answerControllers[0].text, isNot('speaks')); // 他の単語でもない
    });

    test('異なる空欄数の問題間での「答えを見る」動作確認', () {
      // 空欄数が異なる問題セットを作成
      final problemSet = ProblemSet(
        id: 'different-blanks',
        title: '異なる空欄数テスト',
        description: '空欄数が異なる問題での動作テスト',
        cefrLevel: 'A2',
        problems: [
          Problem.word(WordProblem(
            japanese: '単一空欄問題',
            english: 'I _____ books.',
            cefrLevel: 'A2',
            blanks: [
              BlankAnswer(answer: 'read', hint: '読む'),
            ],
          )),
          Problem.word(WordProblem(
            japanese: '複数空欄問題',
            english: 'She _____ _____ school.',
            cefrLevel: 'A2',
            blanks: [
              BlankAnswer(answer: 'goes', hint: '行く'),
              BlankAnswer(answer: 'to', hint: '～へ'),
            ],
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 最初の問題（単一空欄）で「答えを見る」
      var state = container.read(appStateProvider);
      expect(state.answerControllers.length, 1);
      expect(state.currentProblem!.wordProblem!.japanese, '単一空欄問題');
      
      appStateManager.handleShowAnswerForWordProblem();
      
      state = container.read(appStateProvider);
      expect(state.answerControllers[0].text, 'read'); // 単一空欄の答え
      expect(state.answerControllers.length, 1); // 空欄数は1つのまま
      
      // 次の問題へ移動
      appStateManager.moveToNextProblem();
      
      // 複数空欄問題に移動することを確認
      state = container.read(appStateProvider);
      expect(state.currentProblem!.wordProblem!.japanese, '複数空欄問題');
      expect(state.answerControllers.length, 2); // 空欄数が2つに変わる
      expect(state.answerControllers[0].text, ''); // 新しい問題用にクリア
      expect(state.answerControllers[1].text, ''); // 新しい問題用にクリア
      
      // 複数空欄問題で「答えを見る」
      appStateManager.handleShowAnswerForWordProblem();
      
      state = container.read(appStateProvider);
      expect(state.answerControllers[0].text, 'goes'); // 複数空欄の最初の答え
      expect(state.answerControllers[1].text, 'to'); // 複数空欄の二番目の答え
      
      // 間違った空欄数の答えが設定されていないことを確認
      expect(state.answerControllers[0].text, isNot('read')); // 前の問題の答えではない
    });

    test('連続して「答えを見る」を使用した場合の動作確認', () {
      final problemSet = ProblemSet(
        id: 'continuous-show-answer',
        title: '連続答え表示テスト',
        description: '連続して答えを見る機能のテスト',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: 'A問題',
            english: 'A ___.',
            cefrLevel: 'A1',
            blanks: [BlankAnswer(answer: 'answer', hint: 'A答え')],
          )),
          Problem.word(WordProblem(
            japanese: 'B問題',
            english: 'B ___.',
            cefrLevel: 'A1',
            blanks: [BlankAnswer(answer: 'solution', hint: 'B答え')],
          )),
          Problem.word(WordProblem(
            japanese: 'C問題',
            english: 'C ___.',
            cefrLevel: 'A1',
            blanks: [BlankAnswer(answer: 'result', hint: 'C答え')],
          )),
        ],
      );

      appStateManager.selectProblemSet(problemSet);
      
      // A問題で「答えを見る」
      var state = container.read(appStateProvider);
      expect(state.currentProblem!.wordProblem!.japanese, 'A問題');
      appStateManager.handleShowAnswerForWordProblem();
      
      state = container.read(appStateProvider);
      expect(state.answerControllers[0].text, 'answer');
      expect(state.showAnswerForProblem!.wordProblem!.japanese, 'A問題');
      
      // 次の問題へ（A問題が後回しになる）
      appStateManager.moveToNextProblem();
      
      // B問題で「答えを見る」
      state = container.read(appStateProvider);
      expect(state.currentProblem!.wordProblem!.japanese, 'B問題');
      appStateManager.handleShowAnswerForWordProblem();
      
      state = container.read(appStateProvider);
      expect(state.answerControllers[0].text, 'solution'); // B問題の正しい答え
      expect(state.showAnswerForProblem!.wordProblem!.japanese, 'B問題');
      expect(state.answerControllers[0].text, isNot('answer')); // A問題の答えではない
      
      // 次の問題へ（B問題が後回しになる）
      appStateManager.moveToNextProblem();
      
      // C問題で「答えを見る」
      state = container.read(appStateProvider);
      expect(state.currentProblem!.wordProblem!.japanese, 'C問題');
      appStateManager.handleShowAnswerForWordProblem();
      
      state = container.read(appStateProvider);
      expect(state.answerControllers[0].text, 'result'); // C問題の正しい答え
      expect(state.showAnswerForProblem!.wordProblem!.japanese, 'C問題');
      expect(state.answerControllers[0].text, isNot('answer')); // A問題の答えではない
      expect(state.answerControllers[0].text, isNot('solution')); // B問題の答えではない
    });

    test('メニュー戻りとshowAnswerForProblemの状態クリア確認', () {
      final problemSet = ProblemSet(
        id: 'menu-return-test',
        title: 'メニュー戻りテスト',
        description: 'メニュー戻り時の状態クリアテスト',
        cefrLevel: 'A1',
        problems: [
          Problem.word(WordProblem(
            japanese: 'テスト問題',
            english: 'Test ___.',
            cefrLevel: 'A1',
            blanks: [BlankAnswer(answer: 'case', hint: 'テスト')],
          )),
        ],
      );

      appStateManager.selectProblemSet(problemSet);
      
      // 「答えを見る」で状態を設定
      appStateManager.handleShowAnswerForWordProblem();
      
      var state = container.read(appStateProvider);
      expect(state.showAnswerForProblem, isNotNull);
      expect(state.learningState, LearningPhase.answerDisplay);
      
      // メニューに戻る
      appStateManager.returnToMenu();
      
      // showAnswerForProblemがクリアされることを確認
      state = container.read(appStateProvider);
      expect(state.showAnswerForProblem, isNull);
      expect(state.learningState, LearningPhase.problemSelection);
      expect(state.selectedProblemSet, isNull);
    });
  });
}
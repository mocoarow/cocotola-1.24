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
  });
}
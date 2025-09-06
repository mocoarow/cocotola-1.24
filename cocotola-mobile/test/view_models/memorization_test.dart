import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:cocotola/view_models/app_state_manager.dart';
import 'package:cocotola/models/memorization_problem.dart';
import 'package:cocotola/models/problem_base.dart';
import 'package:cocotola/models/problem_set.dart';

void main() {
  // Flutterのバインディングを初期化
  TestWidgetsFlutterBinding.ensureInitialized();
  
  group('暗記問題のAppStateManagerテスト', () {
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

    test('暗記問題セットの選択', () {
      // 暗記問題セットを作成
      final problemSet = ProblemSet(
        id: 'test-memorization-set',
        title: '暗記問題テストセット',
        description: 'テスト用の暗記問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.memorization(MemorizationProblem(
            question: 'apple',
            answer: 'りんご',
            hint: '赤い果物',
            cefrLevel: 'A1',
          )),
          Problem.memorization(MemorizationProblem(
            question: 'book',
            answer: '本',
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 状態確認
      final state = container.read(appStateProvider);
      expect(state.selectedProblemSet, equals(problemSet));
      expect(state.currentProblemIndex, 0);
      expect(state.learningState, LearningPhase.problemDisplay);
      expect(state.problems.length, 2);
      expect(state.problems[0].type, ProblemType.memorization);
      
      // 暗記問題の場合、コントローラーは不要
      expect(state.answerControllers, isEmpty);
      expect(state.answerFocusNodes, isEmpty);
    });

    test('暗記問題の「できた」回答処理', () {
      // 暗記問題セットを作成
      final problemSet = ProblemSet(
        id: 'test-memorization-set',
        title: '暗記問題テストセット',
        description: 'テスト用の暗記問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.memorization(MemorizationProblem(
            question: 'apple',
            answer: 'りんご',
            hint: '赤い果物',
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 初期状態確認
      var state = container.read(appStateProvider);
      final memoProblem = state.currentProblem!.memorizationProblem!;
      expect(memoProblem.isAnswered, false);
      expect(memoProblem.isCorrect, false);
      
      // 「できた」と回答
      appStateManager.handleMemorizationAnswer(true);
      
      // 正解した問題はリストから削除されるため、問題数が減ることを確認
      state = container.read(appStateProvider);
      expect(state.problems.length, 0); // 1問から1問削除されて0問になる（完了状態）
      expect(state.learningState, LearningPhase.completed); // 全問完了
    });

    test('暗記問題の「できなかった」回答処理', () {
      // 暗記問題セットを作成（複数問題で「できなかった」動作を確認）
      final problemSet = ProblemSet(
        id: 'test-memorization-set',
        title: '暗記問題テストセット',
        description: 'テスト用の暗記問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.memorization(MemorizationProblem(
            question: 'book',
            answer: '本',
            cefrLevel: 'A1',
          )),
          Problem.memorization(MemorizationProblem(
            question: 'car',
            answer: '車',
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 最初の問題を確認
      var state = container.read(appStateProvider);
      final initialProblem = state.currentProblem!.memorizationProblem!;
      expect(initialProblem.question, 'book');
      
      // 「できなかった」と回答
      appStateManager.handleMemorizationAnswer(false);
      
      // 間違えた問題は後回しされ、次の問題に移ることを確認
      state = container.read(appStateProvider);
      expect(state.problems.length, 2); // 問題数は変わらない
      final currentProblem = state.currentProblem!.memorizationProblem!;
      expect(currentProblem.question, 'car'); // 次の問題になっている
    });

    test('複数の暗記問題での問題移動', () async {
      // 複数の暗記問題セットを作成
      final problemSet = ProblemSet(
        id: 'test-memorization-set',
        title: '暗記問題テストセット',
        description: 'テスト用の暗記問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.memorization(MemorizationProblem(
            question: 'apple',
            answer: 'りんご',
            cefrLevel: 'A1',
          )),
          Problem.memorization(MemorizationProblem(
            question: 'book',
            answer: '本',
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 初期状態確認
      var state = container.read(appStateProvider);
      expect(state.currentProblemIndex, 0);
      expect(state.currentProblem!.memorizationProblem!.question, 'apple');
      
      // 最初の問題に「できた」と回答
      appStateManager.handleMemorizationAnswer(true);
      
      // 正解した問題はリストから削除され、次の問題が表示されることを確認
      state = container.read(appStateProvider);
      expect(state.currentProblemIndex, 0); // インデックスは0のまま（リストから削除されるため）
      expect(state.problems.length, 1); // 問題数が減る
      expect(state.currentProblem!.memorizationProblem!.question, 'book'); // 次の問題
      expect(state.learningState, LearningPhase.problemDisplay);
    });

    test('全暗記問題完了時の状態', () async {
      // 単一の暗記問題セットを作成
      final problemSet = ProblemSet(
        id: 'test-memorization-set',
        title: '暗記問題テストセット',
        description: 'テスト用の暗記問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.memorization(MemorizationProblem(
            question: 'water',
            answer: '水',
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 問題に回答
      appStateManager.handleMemorizationAnswer(true);
      
      // 全問完了で完了状態になることを確認
      final state = container.read(appStateProvider);
      expect(state.learningState, LearningPhase.completed);
      expect(state.problems.length, 0); // 全問削除された
    });

    test('暗記問題での答え表示画面遷移', () {
      // 暗記問題セットを作成
      final problemSet = ProblemSet(
        id: 'test-memorization-set',
        title: '暗記問題テストセット',
        description: 'テスト用の暗記問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.memorization(MemorizationProblem(
            question: 'hello',
            answer: 'こんにちは',
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 初期状態確認
      var state = container.read(appStateProvider);
      expect(state.learningState, LearningPhase.problemDisplay);
      
      // 答え表示画面に遷移
      appStateManager.transitionToAnswerDisplay();
      
      // 答え表示状態になることを確認
      state = container.read(appStateProvider);
      expect(state.learningState, LearningPhase.answerDisplay);
    });

    test('メニューに戻る機能（暗記問題）', () {
      // 暗記問題セットを作成
      final problemSet = ProblemSet(
        id: 'test-memorization-set',
        title: '暗記問題テストセット',
        description: 'テスト用の暗記問題セット',
        cefrLevel: 'A1',
        problems: [
          Problem.memorization(MemorizationProblem(
            question: 'cat',
            answer: 'ねこ',
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // メニューに戻る
      appStateManager.returnToMenu();
      
      // 状態がリセットされたことを確認
      final state = container.read(appStateProvider);
      expect(state.selectedProblemSet, isNull);
      expect(state.learningState, LearningPhase.problemSelection);
      expect(state.currentProblemIndex, 0);
      expect(state.problems, isEmpty);
      expect(state.answerControllers, isEmpty);
      expect(state.answerFocusNodes, isEmpty);
    });

    test('暗記問題の回答処理でのアトミックな状態更新確認', () {
      // 複数の暗記問題セットを作成
      final problemSet = ProblemSet(
        id: 'test-atomic-update',
        title: 'アトミック更新テスト',
        description: '状態更新がアトミックに行われることをテスト',
        cefrLevel: 'A1',
        problems: [
          Problem.memorization(MemorizationProblem(
            question: 'first',
            answer: '最初',
            cefrLevel: 'A1',
          )),
          Problem.memorization(MemorizationProblem(
            question: 'second',
            answer: '二番目',
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 初期状態確認
      var state = container.read(appStateProvider);
      expect(state.currentProblem!.memorizationProblem!.question, 'first');
      expect(state.learningState, LearningPhase.problemDisplay);
      expect(state.problems.length, 2);
      
      // 最初の問題に「できなかった」と回答
      appStateManager.handleMemorizationAnswer(false);
      
      // 状態が一度に更新されていることを確認
      state = container.read(appStateProvider);
      expect(state.currentProblem!.memorizationProblem!.question, 'second'); // 次の問題
      expect(state.learningState, LearningPhase.problemDisplay); // 問題表示状態
      expect(state.problems.length, 2); // 問題数は変わらない
      expect(state.problems.last.memorizationProblem!.question, 'first'); // 最初の問題が後回し
    });

    test('暗記問題で正解時の即座の状態更新確認', () {
      // 単一の暗記問題セットを作成
      final problemSet = ProblemSet(
        id: 'test-immediate-completion',
        title: '即座の完了テスト',
        description: '正解時の即座の状態更新をテスト',
        cefrLevel: 'A1',
        problems: [
          Problem.memorization(MemorizationProblem(
            question: 'test',
            answer: 'テスト',
            cefrLevel: 'A1',
          )),
        ],
      );

      // 問題セットを選択
      appStateManager.selectProblemSet(problemSet);
      
      // 初期状態確認
      var state = container.read(appStateProvider);
      expect(state.problems.length, 1);
      expect(state.learningState, LearningPhase.problemDisplay);
      
      // 問題に「できた」と回答
      appStateManager.handleMemorizationAnswer(true);
      
      // 即座に完了状態になることを確認
      state = container.read(appStateProvider);
      expect(state.learningState, LearningPhase.completed);
      expect(state.problems.length, 0);
    });
  });
}
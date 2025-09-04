import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:cocotola/view_models/learning_view_model.dart';
import 'package:cocotola/models/word_problem.dart';

void main() {
  // Flutterのバインディングを初期化
  TestWidgetsFlutterBinding.ensureInitialized();
  
  group('LearningViewModel Tests', () {
    late ProviderContainer container;
    late LearningViewModel viewModel;

    setUp(() {
      // テスト用のProviderContainerを作成
      container = ProviderContainer();
      viewModel = container.read(learningViewModelProvider.notifier);
    });

    tearDown(() {
      container.dispose();
    });

    test('初期状態のテスト', () {
      final state = container.read(learningViewModelProvider);
      
      expect(state.currentIndex, 0);
      expect(state.currentState, LearningState.problemDisplay);
      expect(state.answerControllers, isEmpty);
      expect(state.answerFocusNodes, isEmpty);
    });

    test('handleAllBlanksCompleted - 全空欄完了時に答え表示状態に遷移', () {
      viewModel.handleAllBlanksCompleted();
      
      final state = container.read(learningViewModelProvider);
      expect(state.currentState, LearningState.answerDisplay);
    });

    test('transitionToAnswerDisplay - 答え表示画面への遷移', () {
      viewModel.transitionToAnswerDisplay();
      
      final state = container.read(learningViewModelProvider);
      expect(state.currentState, LearningState.answerDisplay);
    });

    test('initializeControllersForCurrentProblem - コントローラーの初期化', () {
      // テスト用の問題データを作成
      final problems = [
        WordProblem(
          japanese: 'テスト問題',
          english: 'This is ___ test.',
          cefrLevel: 'A1',
          blanks: [
            BlankAnswer(
              answer: 'a',
              hint: 'テストヒント',
              isAnswered: false,
              isCorrect: false,
              userInput: '',
            ),
          ],
        ),
      ];
      
      viewModel.initializeControllersForCurrentProblem(0, problems);
      
      final state = container.read(learningViewModelProvider);
      expect(state.answerControllers.length, 1);
      expect(state.answerFocusNodes.length, 1);
    });

    test('複数の空欄を持つ問題でのコントローラー初期化', () {
      final problems = [
        WordProblem(
          japanese: '複数空欄テスト',
          english: 'I ___ ___ school.',
          cefrLevel: 'A2',
          blanks: [
            BlankAnswer(answer: 'go', hint: 'ヒント1'),
            BlankAnswer(answer: 'to', hint: 'ヒント2'),
          ],
        ),
      ];
      
      viewModel.initializeControllersForCurrentProblem(0, problems);
      
      final state = container.read(learningViewModelProvider);
      expect(state.answerControllers.length, 2);
      expect(state.answerFocusNodes.length, 2);
    });

    test('コントローラーの初期化 - 既存のコントローラーが適切に破棄される', () {
      final firstProblems = [
        WordProblem(
          japanese: '最初の問題',
          english: 'First ___ problem.',
          cefrLevel: 'A1',
          blanks: [BlankAnswer(answer: 'test', hint: 'ヒント')],
        ),
      ];
      
      final secondProblems = [
        WordProblem(
          japanese: '2番目の問題',
          english: 'Second ___ ___ problem.',
          cefrLevel: 'B1',
          blanks: [
            BlankAnswer(answer: 'test', hint: 'ヒント1'),
            BlankAnswer(answer: 'case', hint: 'ヒント2'),
          ],
        ),
      ];
      
      // 最初の初期化
      viewModel.initializeControllersForCurrentProblem(0, firstProblems);
      final firstState = container.read(learningViewModelProvider);
      expect(firstState.answerControllers.length, 1);
      
      // 2番目の初期化（異なる数の空欄）
      viewModel.initializeControllersForCurrentProblem(0, secondProblems);
      final secondState = container.read(learningViewModelProvider);
      expect(secondState.answerControllers.length, 2);
    });

    test('状態のイミュータビリティテスト', () {
      final initialState = container.read(learningViewModelProvider);
      
      // 状態を変更
      viewModel.handleAllBlanksCompleted();
      
      final newState = container.read(learningViewModelProvider);
      
      // 元の状態は変更されていないことを確認
      expect(initialState.currentState, LearningState.problemDisplay);
      expect(newState.currentState, LearningState.answerDisplay);
      
      // 新しい状態オブジェクトが作成されていることを確認
      expect(identical(initialState, newState), false);
    });

    test('handleAnswerChanged - メソッドが呼び出せることを確認', () {
      // テスト用の問題データを初期化
      final problems = [
        WordProblem(
          japanese: 'テスト問題',
          english: 'This is ___ test.',
          cefrLevel: 'A1',
          blanks: [
            BlankAnswer(
              answer: 'a',
              hint: 'テストヒント',
            ),
          ],
        ),
      ];
      
      viewModel.initializeProblems(problems);
      
      // メソッドが例外をスローしないことを確認
      expect(() => viewModel.handleAnswerChanged(0, 'test'), returnsNormally);
    });

    test('handleShowAnswer - メソッドが正しく動作することを確認', () {
      // テスト用の問題データを初期化
      final problems = [
        WordProblem(
          japanese: 'テスト問題',
          english: 'This is ___ test.',
          cefrLevel: 'A1',
          blanks: [
            BlankAnswer(
              answer: 'a',
              hint: 'テストヒント',
            ),
          ],
        ),
      ];
      
      viewModel.initializeProblems(problems);
      
      // 初期状態を確認
      expect(container.read(learningViewModelProvider).currentState, 
             LearningState.problemDisplay);
      
      // handleShowAnswerを呼び出し
      viewModel.handleShowAnswer();
      
      // 答え表示状態に遷移することを確認
      expect(container.read(learningViewModelProvider).currentState, 
             LearningState.answerDisplay);
    });

    test('transitionToNextProblem - メソッドが呼び出せることを確認', () {
      // メソッドが例外をスローしないことを確認
      expect(() => viewModel.transitionToNextProblem(), returnsNormally);
    });
  });

  group('LearningViewState Tests', () {
    test('copyWith - 部分的な状態更新', () {
      const initialState = LearningViewState(
        currentIndex: 0,
        currentState: LearningState.problemDisplay,
        answerControllers: [],
        answerFocusNodes: [],
        problems: [],
      );

      // currentStateのみ変更
      final updatedState = initialState.copyWith(
        currentState: LearningState.answerDisplay,
      );

      expect(updatedState.currentState, LearningState.answerDisplay);
      expect(updatedState.currentIndex, 0); // 変更されていないことを確認
      expect(updatedState.answerControllers, isEmpty);
      expect(updatedState.answerFocusNodes, isEmpty);
    });

    test('copyWith - 複数フィールドの同時更新', () {
      const initialState = LearningViewState(
        currentIndex: 0,
        currentState: LearningState.problemDisplay,
        answerControllers: [],
        answerFocusNodes: [],
        problems: [],
      );

      final updatedState = initialState.copyWith(
        currentIndex: 1,
        currentState: LearningState.completed,
      );

      expect(updatedState.currentIndex, 1);
      expect(updatedState.currentState, LearningState.completed);
      expect(updatedState.answerControllers, isEmpty); // 変更されていないことを確認
      expect(updatedState.answerFocusNodes, isEmpty);
    });

    test('copyWith - nullを渡した場合は元の値を保持', () {
      const initialState = LearningViewState(
        currentIndex: 5,
        currentState: LearningState.answerDisplay,
        answerControllers: [],
        answerFocusNodes: [],
        problems: [],
      );

      final updatedState = initialState.copyWith(
        currentIndex: null,
        currentState: null,
      );

      // nullを渡したフィールドは元の値を保持
      expect(updatedState.currentIndex, 5);
      expect(updatedState.currentState, LearningState.answerDisplay);
      expect(updatedState.answerControllers, isEmpty);
      expect(updatedState.answerFocusNodes, isEmpty);
    });
  });

  group('LearningState Enum Tests', () {
    test('LearningState列挙値のテスト', () {
      expect(LearningState.problemDisplay.toString(), contains('problemDisplay'));
      expect(LearningState.answerDisplay.toString(), contains('answerDisplay'));
      expect(LearningState.completed.toString(), contains('completed'));
    });

    test('LearningState列挙値の順序性', () {
      final states = LearningState.values;
      expect(states.length, 3);
      expect(states.contains(LearningState.problemDisplay), true);
      expect(states.contains(LearningState.answerDisplay), true);
      expect(states.contains(LearningState.completed), true);
    });
  });
}
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../models/word_problem.dart';
import '../providers/word_problem_provider.dart';
import 'dart:developer' as developer;

enum LearningState {
  problemDisplay, // 問題表示状態
  answerDisplay, // 答え表示状態（解説表示）
  completed // 完了状態
}

class LearningViewState {
  final int currentIndex;
  final LearningState currentState;
  final List<TextEditingController> answerControllers;
  final List<FocusNode> answerFocusNodes;

  const LearningViewState({
    required this.currentIndex,
    required this.currentState,
    required this.answerControllers,
    required this.answerFocusNodes,
  });

  LearningViewState copyWith({
    int? currentIndex,
    LearningState? currentState,
    List<TextEditingController>? answerControllers,
    List<FocusNode>? answerFocusNodes,
  }) {
    return LearningViewState(
      currentIndex: currentIndex ?? this.currentIndex,
      currentState: currentState ?? this.currentState,
      answerControllers: answerControllers ?? this.answerControllers,
      answerFocusNodes: answerFocusNodes ?? this.answerFocusNodes,
    );
  }
}

class LearningViewModel extends StateNotifier<LearningViewState> {
  final Ref ref;

  LearningViewModel(this.ref)
      : super(const LearningViewState(
          currentIndex: 0,
          currentState: LearningState.problemDisplay,
          answerControllers: [],
          answerFocusNodes: [],
        ));

  /// ユーザーの入力が変更されたときの処理
  void handleAnswerChanged(int blankIndex, String value) {
    developer.log('[LearningViewModel] handleAnswerChanged: blankIndex=$blankIndex, value=$value');
    
    ref
        .read(wordProblemsProvider.notifier)
        .updateUserInput(state.currentIndex, blankIndex, value);
  }

  /// 全ての空欄が完了したときの処理
  void handleAllBlanksCompleted() {
    developer.log('[LearningViewModel] All blanks completed, automatically transitioning to answer display');
    
    transitionToAnswerDisplay();
  }

  /// 答えを表示するボタンが押されたときの処理
  void handleShowAnswer() {
    developer.log('[LearningViewModel] Show answer requested');
    
    ref.read(wordProblemsProvider.notifier).markAsSkipped(state.currentIndex);
    
    // 先に画面遷移してから正解を表示
    transitionToAnswerDisplay();
    
    // 画面遷移後に正解を設定
    WidgetsBinding.instance.addPostFrameCallback((_) {
      _fillAllBlanksWithAnswers();
    });
  }

  /// 答え表示画面への遷移
  void transitionToAnswerDisplay() {
    developer.log('[LearningViewModel] Transitioning to answer display');
    
    state = state.copyWith(currentState: LearningState.answerDisplay);
  }

  /// 次の問題への遷移処理
  void transitionToNextProblem() {
    developer.log('[LearningViewModel] transitionToNextProblem called, current index: ${state.currentIndex}');

    final problems = ref.read(wordProblemsProvider);
    final nextIndex = _findNextIncompleteProblem(problems);

    if (nextIndex == null) {
      // 全ての問題が完了した場合
      developer.log('[LearningViewModel] All problems completed, transitioning to completed state');
      state = state.copyWith(currentState: LearningState.completed);
      return;
    }

    // 次の未完了問題に遷移
    _moveToNextProblem(nextIndex, problems);
  }

  /// 次の未完了問題のインデックスを検索
  int? _findNextIncompleteProblem(List<WordProblem> problems) {
    // 現在の問題の次から検索
    for (int i = state.currentIndex + 1; i < problems.length; i++) {
      if (!problems[i].isCompleted) {
        return i;
      }
    }

    // 見つからない場合、先頭から現在の問題まで検索（循環）
    for (int i = 0; i < state.currentIndex; i++) {
      if (!problems[i].isCompleted) {
        return i;
      }
    }

    // 全ての問題が完了している場合
    return null;
  }

  /// 指定された問題インデックスに移動
  void _moveToNextProblem(int newIndex, List<WordProblem> problems) {
    developer.log('[LearningViewModel] Moving to problem $newIndex');

    // 新しい問題のユーザー入力をクリア
    ref.read(wordProblemsProvider.notifier).clearUserInputs(newIndex);

    // コントローラーを初期化
    initializeControllersForCurrentProblem(newIndex, problems);

    state = state.copyWith(
      currentIndex: newIndex,
      currentState: LearningState.problemDisplay,
    );
  }

  /// 現在の問題用のコントローラーを初期化
  void initializeControllersForCurrentProblem(int? problemIndex, List<WordProblem> problems) {
    final index = problemIndex ?? state.currentIndex;
    developer.log('[LearningViewModel] initializeControllersForCurrentProblem called for index: $index');

    if (problems.isNotEmpty && index < problems.length) {
      final currentProblem = problems[index];
      final maxBlanks = currentProblem.blanks.length;

      // 既存のコントローラーを破棄
      _disposeControllers();

      // 新しいコントローラーを作成
      final newControllers = List.generate(maxBlanks, (blankIndex) {
        final blank = currentProblem.blanks[blankIndex];
        // 正解済みの場合は答えを表示、そうでなければ空文字
        final initialText = blank.isCorrect ? blank.answer : blank.userInput;
        return TextEditingController(text: initialText);
      });

      final newFocusNodes = List.generate(maxBlanks, (index) => FocusNode());

      state = state.copyWith(
        answerControllers: newControllers,
        answerFocusNodes: newFocusNodes,
      );

      developer.log('[LearningViewModel] Initialized ${newControllers.length} controllers and ${newFocusNodes.length} focus nodes');
      
      // 最初の未回答空欄に自動フォーカス
      _setInitialFocus();
    }
  }

  /// 全ての空欄に正解を表示
  void _fillAllBlanksWithAnswers() {
    final problems = ref.read(wordProblemsProvider);
    if (problems.isNotEmpty && state.currentIndex < problems.length) {
      final currentProblem = problems[state.currentIndex];
      
      developer.log('[LearningViewModel] Filling all blanks with answers for problem ${state.currentIndex}');
      developer.log('[LearningViewModel] Problem has ${currentProblem.blanks.length} blanks');
      developer.log('[LearningViewModel] State has ${state.answerControllers.length} controllers');
      
      // プロバイダーの状態も正解として更新（緑色表示のため）
      for (int i = 0; i < currentProblem.blanks.length; i++) {
        final correctAnswer = currentProblem.blanks[i].answer;
        ref.read(wordProblemsProvider.notifier).checkAnswer(state.currentIndex, i, correctAnswer);
        developer.log('[LearningViewModel] Marked blank $i as correct with answer: "$correctAnswer"');
      }
      
      // コントローラーにも正解を設定
      for (int i = 0; i < currentProblem.blanks.length && i < state.answerControllers.length; i++) {
        final correctAnswer = currentProblem.blanks[i].answer;
        final oldText = state.answerControllers[i].text;
        state.answerControllers[i].text = correctAnswer;
        
        developer.log('[LearningViewModel] Set controller $i: "$oldText" -> "$correctAnswer"');
      }
    }
  }

  /// 最初の未回答空欄に自動フォーカスを設定
  void _setInitialFocus() {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      final problems = ref.read(wordProblemsProvider);
      if (problems.isNotEmpty && state.currentIndex < problems.length) {
        final currentProblem = problems[state.currentIndex];
        
        // 最初の未回答空欄を見つける
        int? firstUnAnsweredIndex;
        for (int i = 0; i < currentProblem.blanks.length; i++) {
          if (!currentProblem.blanks[i].isAnswered || !currentProblem.blanks[i].isCorrect) {
            firstUnAnsweredIndex = i;
            break;
          }
        }
        
        // 見つかった場合、そこにフォーカスを設定
        if (firstUnAnsweredIndex != null && 
            firstUnAnsweredIndex < state.answerFocusNodes.length) {
          developer.log('[LearningViewModel] Setting initial focus to blank $firstUnAnsweredIndex');
          state.answerFocusNodes[firstUnAnsweredIndex].requestFocus();
        }
      }
    });
  }

  /// コントローラーとフォーカスノードを破棄
  void _disposeControllers() {
    for (final controller in state.answerControllers) {
      controller.dispose();
    }
    for (final focusNode in state.answerFocusNodes) {
      focusNode.dispose();
    }
  }

  @override
  void dispose() {
    _disposeControllers();
    super.dispose();
  }
}

final learningViewModelProvider = StateNotifierProvider<LearningViewModel, LearningViewState>((ref) {
  return LearningViewModel(ref);
});
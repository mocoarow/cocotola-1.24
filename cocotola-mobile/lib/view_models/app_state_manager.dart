import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter/material.dart';
import '../models/word_problem.dart';
import '../models/problem_set.dart';
import 'dart:developer' as developer;

/// アプリ全体の状態を管理する統合ViewModel
class AppStateManager extends StateNotifier<AppState> {
  AppStateManager() : super(const AppState());

  /// 問題セットを選択
  void selectProblemSet(ProblemSet problemSet) {
    developer.log('[AppStateManager] Problem set selected: ${problemSet.title}');
    
    state = state.copyWith(
      selectedProblemSet: problemSet,
      currentProblemIndex: 0,
      learningState: LearningPhase.problemDisplay,
      problems: problemSet.problems.map((p) => p.copyWith()).toList(),
      answerControllers: [],
      answerFocusNodes: [],
      currentBlankIndex: 0, // カスタムキーボード用の空欄インデックス初期化
      cursorPosition: 0, // カーソル位置初期化
    );
    
    _initializeControllersForCurrentProblem();
  }

  /// ユーザーの回答を処理
  void handleAnswer(int blankIndex, String value) {
    if (state.selectedProblemSet == null) return;
    
    developer.log('[AppStateManager] Answer changed: blank=$blankIndex, value="$value"');
    
    final updatedProblems = List<WordProblem>.from(state.problems);
    final currentProblem = updatedProblems[state.currentProblemIndex];
    final blank = currentProblem.blanks[blankIndex];
    
    final isCorrect = value.toLowerCase().trim() == blank.answer.toLowerCase().trim();
    final updatedBlank = blank.copyWith(
      userInput: isCorrect ? blank.answer : value, // 正解時は元のanswerを表示
      isAnswered: isCorrect,
      isCorrect: isCorrect,
    );
    
    final updatedProblem = currentProblem.updateBlank(blankIndex, updatedBlank);
    updatedProblems[state.currentProblemIndex] = updatedProblem;
    
    state = state.copyWith(problems: updatedProblems);
    
    // 正解時にコントローラーのテキストも正しい表記に更新
    if (isCorrect && blankIndex < state.answerControllers.length) {
      state.answerControllers[blankIndex].text = blank.answer;
    }
    
    // 正解時の自動処理
    if (isCorrect && updatedProblem.isCompleted) {
      developer.log('[AppStateManager] Problem completed automatically');
      WidgetsBinding.instance.addPostFrameCallback((_) {
        transitionToAnswerDisplay();
      });
    }
  }

  /// 答え表示画面への遷移
  void transitionToAnswerDisplay() {
    developer.log('[AppStateManager] Transitioning to answer display');
    state = state.copyWith(learningState: LearningPhase.answerDisplay);
  }

  /// 次の問題への遷移
  void moveToNextProblem() {
    final nextIndex = _findNextIncompleteProblemIndex();
    
    if (nextIndex == null) {
      // 全問完了
      state = state.copyWith(learningState: LearningPhase.completed);
      return;
    }
    
    state = state.copyWith(
      currentProblemIndex: nextIndex,
      learningState: LearningPhase.problemDisplay,
      currentBlankIndex: 0, // 最初の空欄にリセット
      cursorPosition: 0, // カーソル位置もリセット
    );
    
    _initializeControllersForCurrentProblem();
  }

  /// 問題セットをリセット
  void resetProblemSet() {
    if (state.selectedProblemSet == null) return;
    
    // コントローラーを破棄
    _disposeControllers();
    
    // 問題をリセット
    final resetProblems = state.selectedProblemSet!.problems.map((problem) {
      final resetBlanks = problem.blanks.map((blank) => blank.copyWith(
        userInput: '',
        isAnswered: false,
        isCorrect: false,
      )).toList();
      
      return problem.copyWith(
        blanks: resetBlanks,
        isSkipped: false,
      );
    }).toList();
    
    state = state.copyWith(
      currentProblemIndex: 0,
      learningState: LearningPhase.problemDisplay,
      problems: resetProblems,
      answerControllers: [],
      answerFocusNodes: [],
    );
    
    _initializeControllersForCurrentProblem();
  }

  /// 問題セット選択画面に戻る
  void returnToProblemSelection() {
    developer.log('[AppStateManager] Returning to problem selection');
    
    // コントローラーを破棄
    _disposeControllers();
    
    state = const AppState(
      learningState: LearningPhase.problemSelection,
    );
  }

  /// フォーカスとカーソル位置を更新
  void updateFocusAndCursor(int blankIndex, int cursorPos) {
    developer.log('[AppStateManager] Updating focus: blank=$blankIndex, cursor=$cursorPos');
    
    state = state.copyWith(
      currentBlankIndex: blankIndex,
      cursorPosition: cursorPos,
    );
    
    // コントローラーのカーソル位置も同期
    if (blankIndex < state.answerControllers.length) {
      state.answerControllers[blankIndex].selection = 
          TextSelection.fromPosition(TextPosition(offset: cursorPos));
    }
  }

  /// カスタムキーボードでのキー入力処理
  void handleCustomKeyboardInput(String key) {
    developer.log('[AppStateManager] Custom keyboard input: "$key" at position ${state.cursorPosition} in blank ${state.currentBlankIndex}');
    
    if (state.currentBlankIndex < state.answerControllers.length) {
      final controller = state.answerControllers[state.currentBlankIndex];
      final text = controller.text;
      final newText = text.substring(0, state.cursorPosition) +
          key +
          text.substring(state.cursorPosition);
      
      controller.text = newText;
      final newCursorPos = state.cursorPosition + 1;
      
      // 状態を更新
      state = state.copyWith(cursorPosition: newCursorPos);
      
      // コントローラーのカーソル位置も同期
      controller.selection = TextSelection.fromPosition(
        TextPosition(offset: newCursorPos),
      );

      // 回答処理
      handleAnswer(state.currentBlankIndex, newText);
    }
  }

  /// カスタムキーボードでの削除処理
  void handleCustomKeyboardDelete() {
    if (state.currentBlankIndex < state.answerControllers.length && state.cursorPosition > 0) {
      final controller = state.answerControllers[state.currentBlankIndex];
      final text = controller.text;
      final newText = text.substring(0, state.cursorPosition - 1) +
          text.substring(state.cursorPosition);
      
      controller.text = newText;
      final newCursorPos = state.cursorPosition - 1;
      
      // 状態を更新
      state = state.copyWith(cursorPosition: newCursorPos);
      
      // コントローラーのカーソル位置も同期
      controller.selection = TextSelection.fromPosition(
        TextPosition(offset: newCursorPos),
      );

      // 回答処理
      handleAnswer(state.currentBlankIndex, newText);
    }
  }

  /// カーソル移動処理
  void moveCursor(int delta) {
    if (state.currentBlankIndex < state.answerControllers.length) {
      final controller = state.answerControllers[state.currentBlankIndex];
      final newPos = (state.cursorPosition + delta).clamp(0, controller.text.length);
      
      state = state.copyWith(cursorPosition: newPos);
      
      controller.selection = TextSelection.fromPosition(
        TextPosition(offset: newPos),
      );
    }
  }

  /// コントローラーの初期化
  void _initializeControllersForCurrentProblem() {
    if (state.problems.isEmpty || 
        state.currentProblemIndex >= state.problems.length) {
      return;
    }
    
    _disposeControllers();
    
    final currentProblem = state.problems[state.currentProblemIndex];
    final controllers = List.generate(
      currentProblem.blanks.length,
      (index) => TextEditingController(text: currentProblem.blanks[index].userInput),
    );
    
    final focusNodes = List.generate(
      currentProblem.blanks.length,
      (index) => FocusNode(),
    );
    
    state = state.copyWith(
      answerControllers: controllers,
      answerFocusNodes: focusNodes,
    );
    
    // 自動フォーカスを設定
    _setInitialFocus();
  }

  /// 最初の未回答空欄に自動フォーカスを設定
  void _setInitialFocus() {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      if (state.problems.isEmpty || 
          state.currentProblemIndex >= state.problems.length) {
        return;
      }
      
      final currentProblem = state.problems[state.currentProblemIndex];
      
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
        developer.log('[AppStateManager] Setting initial focus to blank $firstUnAnsweredIndex');
        
        // カスタムキーボード用の状態も更新
        state = state.copyWith(
          currentBlankIndex: firstUnAnsweredIndex,
          cursorPosition: 0,
        );
        
        state.answerFocusNodes[firstUnAnsweredIndex].requestFocus();
      }
    });
  }

  /// 次の未完了問題を検索
  int? _findNextIncompleteProblemIndex() {
    for (int i = state.currentProblemIndex + 1; i < state.problems.length; i++) {
      if (!state.problems[i].isCompleted) return i;
    }
    for (int i = 0; i < state.currentProblemIndex; i++) {
      if (!state.problems[i].isCompleted) return i;
    }
    return null;
  }

  /// コントローラーを破棄
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

enum LearningPhase {
  problemSelection,
  problemDisplay,
  answerDisplay,
  completed,
}

/// アプリ全体の状態
class AppState {
  final ProblemSet? selectedProblemSet;
  final int currentProblemIndex;
  final LearningPhase learningState;
  final List<WordProblem> problems;
  final List<TextEditingController> answerControllers;
  final List<FocusNode> answerFocusNodes;
  final int currentBlankIndex; // 現在フォーカスされている空欄
  final int cursorPosition; // カーソル位置

  const AppState({
    this.selectedProblemSet,
    this.currentProblemIndex = 0,
    this.learningState = LearningPhase.problemSelection,
    this.problems = const [],
    this.answerControllers = const [],
    this.answerFocusNodes = const [],
    this.currentBlankIndex = 0,
    this.cursorPosition = 0,
  });

  AppState copyWith({
    ProblemSet? selectedProblemSet,
    int? currentProblemIndex,
    LearningPhase? learningState,
    List<WordProblem>? problems,
    List<TextEditingController>? answerControllers,
    List<FocusNode>? answerFocusNodes,
    int? currentBlankIndex,
    int? cursorPosition,
  }) {
    return AppState(
      selectedProblemSet: selectedProblemSet ?? this.selectedProblemSet,
      currentProblemIndex: currentProblemIndex ?? this.currentProblemIndex,
      learningState: learningState ?? this.learningState,
      problems: problems ?? this.problems,
      answerControllers: answerControllers ?? this.answerControllers,
      answerFocusNodes: answerFocusNodes ?? this.answerFocusNodes,
      currentBlankIndex: currentBlankIndex ?? this.currentBlankIndex,
      cursorPosition: cursorPosition ?? this.cursorPosition,
    );
  }

  WordProblem? get currentProblem => 
      problems.isNotEmpty && currentProblemIndex < problems.length 
          ? problems[currentProblemIndex] 
          : null;
}

final appStateProvider = StateNotifierProvider<AppStateManager, AppState>((ref) {
  return AppStateManager();
});
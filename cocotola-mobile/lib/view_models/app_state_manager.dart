import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter/material.dart';
import '../models/word_problem.dart';
import '../models/problem_base.dart';
import '../models/problem_set.dart';
import 'dart:developer' as developer;

// copyWith でnullを明示的に設定するための定数
const Object _undefined = Object();

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
      problems: problemSet.problems,
      answerControllers: [],
      answerFocusNodes: [],
      currentBlankIndex: 0, // カスタムキーボード用の空欄インデックス初期化
      cursorPosition: 0, // カーソル位置初期化
    );
    
    _initializeControllersForCurrentProblem();
  }

  /// ユーザーの回答を処理（英単語問題用）
  void handleAnswer(int blankIndex, String value) {
    if (state.selectedProblemSet == null || state.currentProblem == null) return;
    if (state.currentProblem!.type != ProblemType.word) return;
    
    developer.log('[AppStateManager] Word problem answer changed: blank=$blankIndex, value="$value"');
    
    final updatedProblems = List<Problem>.from(state.problems);
    final currentWordProblem = updatedProblems[state.currentProblemIndex].wordProblem!;
    final blank = currentWordProblem.blanks[blankIndex];
    
    final isCorrect = value.toLowerCase().trim() == blank.answer.toLowerCase().trim();
    final updatedBlank = blank.copyWith(
      userInput: isCorrect ? blank.answer : value, // 正解時は元のanswerを表示
      isAnswered: isCorrect,
      isCorrect: isCorrect,
    );
    
    final updatedWordProblem = currentWordProblem.updateBlank(blankIndex, updatedBlank);
    updatedProblems[state.currentProblemIndex] = Problem.word(updatedWordProblem);
    
    state = state.copyWith(problems: updatedProblems);
    
    // 正解時にコントローラーのテキストも正しい表記に更新
    if (isCorrect && blankIndex < state.answerControllers.length) {
      state.answerControllers[blankIndex].text = blank.answer;
    }
    
    // 正解時の自動処理
    if (isCorrect) {
      if (updatedWordProblem.isCompleted) {
        developer.log('[AppStateManager] Problem completed automatically');
        WidgetsBinding.instance.addPostFrameCallback((_) {
          transitionToAnswerDisplay();
        });
      } else {
        // 次の未正解の空欄に自動フォーカス移動
        _moveToNextIncorrectBlank(updatedWordProblem);
      }
    }
  }

  /// 暗記問題の回答を処理（できた/できなかった）
  void handleMemorizationAnswer(bool wasCorrect) {
    if (state.selectedProblemSet == null || state.currentProblem == null) return;
    if (state.currentProblem!.type != ProblemType.memorization) return;
    
    developer.log('[AppStateManager] Memorization problem answered: $wasCorrect');
    
    final updatedProblems = List<Problem>.from(state.problems);
    final currentMemoProblem = updatedProblems[state.currentProblemIndex].memorizationProblem!;
    
    final updatedMemoProblem = currentMemoProblem.markAsAnswered(wasCorrect);
    updatedProblems[state.currentProblemIndex] = Problem.memorization(updatedMemoProblem);
    
    state = state.copyWith(problems: updatedProblems);
    
    // 次の問題に自動遷移
    WidgetsBinding.instance.addPostFrameCallback((_) {
      moveToNextProblem();
    });
  }

  /// 答え表示画面への遷移
  void transitionToAnswerDisplay() {
    developer.log('[AppStateManager] Transitioning to answer display');
    
    // 英単語問題の場合のみコントローラーに正解を設定
    if (state.currentProblem != null && 
        state.currentProblem!.type == ProblemType.word &&
        state.answerControllers.isNotEmpty) {
      final currentWordProblem = state.currentProblem!.wordProblem!;
      for (int i = 0; i < currentWordProblem.blanks.length && i < state.answerControllers.length; i++) {
        state.answerControllers[i].text = currentWordProblem.blanks[i].answer;
      }
      developer.log('[AppStateManager] Set correct answers to controllers for answer display');
    }
    
    state = state.copyWith(learningState: LearningPhase.answerDisplay);
  }

  /// 次の問題への遷移
  void moveToNextProblem() {
    // 暗記問題の場合は単純に次のインデックスに移動
    if (state.currentProblem?.type == ProblemType.memorization) {
      final nextIndex = state.currentProblemIndex + 1;
      
      if (nextIndex >= state.problems.length) {
        // 全問完了
        state = state.copyWith(learningState: LearningPhase.completed);
        return;
      }
      
      state = state.copyWith(
        currentProblemIndex: nextIndex,
        learningState: LearningPhase.problemDisplay,
      );
    } else {
      // 英単語問題の場合は未完了問題を探す
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
    }
    
    _initializeControllersForCurrentProblem();
  }

  /// メニュー（問題セット選択画面）に戻る
  void returnToMenu() {
    developer.log('[AppStateManager] Returning to menu');
    
    // コントローラーを破棄
    _disposeControllers();
    
    // 状態をリセット
    state = state.copyWith(
      selectedProblemSet: null,
      currentProblemIndex: 0,
      learningState: LearningPhase.problemSelection,
      problems: [],
      answerControllers: [],
      answerFocusNodes: [],
      currentBlankIndex: 0,
      cursorPosition: 0,
    );
  }

  /// 問題セットをリセット
  void resetProblemSet() {
    if (state.selectedProblemSet == null) return;
    
    // コントローラーを破棄
    _disposeControllers();
    
    // 問題をリセット
    final resetProblems = state.selectedProblemSet!.problems.map((problem) {
      if (problem.type == ProblemType.word) {
        final wordProblem = problem.wordProblem!;
        final resetBlanks = wordProblem.blanks.map((blank) => blank.copyWith(
          userInput: '',
          isAnswered: false,
          isCorrect: false,
        )).toList();
        
        final resetWordProblem = wordProblem.copyWith(
          blanks: resetBlanks,
          isSkipped: false,
        );
        return Problem.word(resetWordProblem);
      } else if (problem.type == ProblemType.memorization) {
        final memoProblem = problem.memorizationProblem!;
        final resetMemoProblem = memoProblem.copyWith(
          isAnswered: false,
          isCorrect: false,
        );
        return Problem.memorization(resetMemoProblem);
      } else {
        return problem;
      }
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

  /// コントローラーの初期化（英単語問題の場合のみ）
  void _initializeControllersForCurrentProblem() {
    if (state.problems.isEmpty || 
        state.currentProblemIndex >= state.problems.length) {
      return;
    }
    
    _disposeControllers();
    
    final currentProblem = state.problems[state.currentProblemIndex];
    
    // 英単語問題の場合のみコントローラーを初期化
    if (currentProblem.type == ProblemType.word) {
      final wordProblem = currentProblem.wordProblem!;
      final controllers = List.generate(
        wordProblem.blanks.length,
        (index) => TextEditingController(text: wordProblem.blanks[index].userInput),
      );
      
      final focusNodes = List.generate(
        wordProblem.blanks.length,
        (index) => FocusNode(),
      );
      
      state = state.copyWith(
        answerControllers: controllers,
        answerFocusNodes: focusNodes,
      );
      
      // 自動フォーカスを設定
      _setInitialFocus();
    } else {
      // 暗記問題の場合はコントローラーは不要
      state = state.copyWith(
        answerControllers: [],
        answerFocusNodes: [],
      );
    }
  }

  /// 最初の未回答空欄に自動フォーカスを設定（英単語問題の場合のみ）
  void _setInitialFocus() {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      if (state.problems.isEmpty || 
          state.currentProblemIndex >= state.problems.length) {
        return;
      }
      
      final currentProblem = state.problems[state.currentProblemIndex];
      
      // 英単語問題の場合のみフォーカス設定
      if (currentProblem.type == ProblemType.word) {
        final wordProblem = currentProblem.wordProblem!;
        
        // 最初の未回答空欄を見つける
        int? firstUnAnsweredIndex;
        for (int i = 0; i < wordProblem.blanks.length; i++) {
          if (!wordProblem.blanks[i].isAnswered || !wordProblem.blanks[i].isCorrect) {
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
      }
    });
  }

  /// 次の未正解空欄に自動フォーカス移動
  void _moveToNextIncorrectBlank(WordProblem problem) {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      // 次の未正解空欄を探す
      int? nextIncorrectBlankIndex;
      for (int i = 0; i < problem.blanks.length; i++) {
        if (!problem.blanks[i].isAnswered || !problem.blanks[i].isCorrect) {
          nextIncorrectBlankIndex = i;
          break;
        }
      }
      
      // 見つかった場合、フォーカス移動
      if (nextIncorrectBlankIndex != null && 
          nextIncorrectBlankIndex < state.answerFocusNodes.length) {
        developer.log('[AppStateManager] Moving focus to blank $nextIncorrectBlankIndex');
        
        // カスタムキーボード用の状態も更新
        state = state.copyWith(
          currentBlankIndex: nextIncorrectBlankIndex,
          cursorPosition: 0,
        );
        
        state.answerFocusNodes[nextIncorrectBlankIndex].requestFocus();
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
  final List<Problem> problems;
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
    Object? selectedProblemSet = _undefined,
    int? currentProblemIndex,
    LearningPhase? learningState,
    List<Problem>? problems,
    List<TextEditingController>? answerControllers,
    List<FocusNode>? answerFocusNodes,
    int? currentBlankIndex,
    int? cursorPosition,
  }) {
    return AppState(
      selectedProblemSet: selectedProblemSet == _undefined 
          ? this.selectedProblemSet 
          : selectedProblemSet as ProblemSet?,
      currentProblemIndex: currentProblemIndex ?? this.currentProblemIndex,
      learningState: learningState ?? this.learningState,
      problems: problems ?? this.problems,
      answerControllers: answerControllers ?? this.answerControllers,
      answerFocusNodes: answerFocusNodes ?? this.answerFocusNodes,
      currentBlankIndex: currentBlankIndex ?? this.currentBlankIndex,
      cursorPosition: cursorPosition ?? this.cursorPosition,
    );
  }

  Problem? get currentProblem => 
      problems.isNotEmpty && currentProblemIndex < problems.length 
          ? problems[currentProblemIndex] 
          : null;
}

final appStateProvider = StateNotifierProvider<AppStateManager, AppState>((ref) {
  return AppStateManager();
});
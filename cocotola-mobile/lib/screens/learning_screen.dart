import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/word_problem_provider.dart';
import '../models/word_problem.dart';
import '../models/problem_display_config.dart';
import '../ui/screens/answer_display_screen.dart';
import '../ui/screens/problem_display_screen.dart';
import 'dart:developer' as developer;

enum LearningState {
  problemDisplay, // 問題表示状態
  answerDisplay, // 答え表示状態（解説表示）
  completed // 完了状態
}

class LearningScreen extends ConsumerStatefulWidget {
  const LearningScreen({super.key});

  @override
  ConsumerState<LearningScreen> createState() => _LearningScreenState();
}

class _LearningScreenState extends ConsumerState<LearningScreen> {
  List<TextEditingController> _answerControllers = [];
  List<FocusNode> _answerFocusNodes = [];
  int _currentIndex = 0;
  LearningState _currentState = LearningState.problemDisplay;

  @override
  void initState() {
    super.initState();
    developer.log('[LearningScreen] initState called');
    // 初期化は最初のbuildで行う
  }

  void _initializeControllersForCurrentProblem(List<WordProblem> problems) {
    developer.log(
        '[LearningScreen] _initializeControllersForCurrentProblem called for currentIndex: $_currentIndex');
    if (problems.isNotEmpty && _currentIndex < problems.length) {
      final maxBlanks = problems[_currentIndex].blanks.length;
      developer
          .log('[LearningScreen] Max blanks for current problem: $maxBlanks');

      // 既存のコントローラーを破棄
      _disposeControllers();

      // 新しいコントローラーを作成
      final currentProblem = problems[_currentIndex];
      developer.log(
          '[LearningScreen] Current problem has ${currentProblem.blanks.length} blanks');

      _answerControllers = List.generate(maxBlanks, (index) {
        // stateのuserInputを確認し、正解済みでない場合は空文字にする
        final blank = currentProblem.blanks[index];
        developer.log(
            '[LearningScreen] Blank[$index] state: userInput="${blank.userInput}", isAnswered=${blank.isAnswered}, isCorrect=${blank.isCorrect}');

        // 正解済みの場合は答えを表示、そうでなければ空文字
        final initialText =
            (blank.isAnswered && blank.isCorrect) ? blank.answer : '';
        developer.log(
            '[LearningScreen] Controller[$index] will be initialized with: "$initialText"');
        return TextEditingController(text: initialText);
      });
      _answerFocusNodes = List.generate(maxBlanks, (index) => FocusNode());
      developer.log(
          '[LearningScreen] Controllers initialized: ${_answerControllers.length}');

      // setStateを使って再描画をトリガー
      WidgetsBinding.instance.addPostFrameCallback((_) {
        if (mounted) {
          setState(() {
            // コントローラーが初期化されたことを通知
          });
          if (_answerFocusNodes.isNotEmpty) {
            _answerFocusNodes[0].requestFocus();
          }
        }
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    final problems = ref.watch(wordProblemsProvider);
    developer.log(
        '[LearningScreen] build called - problems count: ${problems.length}, currentIndex: $_currentIndex, state: $_currentState');

    if (problems.isEmpty) {
      developer.log('[LearningScreen] No problems available');
      return const Center(child: Text('お疲れ様でした！'));
    }

    // 状態に基づいてUIを返す
    switch (_currentState) {
      case LearningState.completed:
        return _buildCompletedScreen();

      case LearningState.answerDisplay:
        return _buildAnswerDisplayScreen(problems);

      case LearningState.problemDisplay:
        return _buildProblemDisplayScreen(problems);
    }
  }

  Widget _buildCompletedScreen() {
    return Scaffold(
      appBar: AppBar(
        title: const Text('単語学習'),
      ),
      body: const Center(
        child: Text(
          'お疲れ様でした！\n全問正解です！',
          textAlign: TextAlign.center,
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }

  Widget _buildAnswerDisplayScreen(List<WordProblem> problems) {
    return AnswerDisplayScreen(
      problems: problems,
      currentIndex: _currentIndex,
      answerControllers: _answerControllers,
      answerFocusNodes: _answerFocusNodes,
      onNextProblem: _transitionToNextProblem,
    );
  }

  Widget _buildProblemDisplayScreen(List<WordProblem> problems) {
    return ProblemDisplayScreen(
      config: ProblemDisplayConfig(
        problems: problems,
        currentIndex: _currentIndex,
        answerControllers: _answerControllers,
        answerFocusNodes: _answerFocusNodes,
      ),
      callbacks: ProblemDisplayCallbacks(
        onAnswerChanged: (blankIndex, value) {
          ref
              .read(wordProblemsProvider.notifier)
              .updateUserInput(_currentIndex, blankIndex, value);
        },
        onAllBlanksCompleted: () {
          developer.log('[LearningScreen] All blanks completed, transitioning to answer display');
          setState(() {
            _currentState = LearningState.answerDisplay;
          });
        },
        onShowAnswer: () {
          ref.read(wordProblemsProvider.notifier).markAsSkipped(_currentIndex);
          _transitionToAnswerDisplay();
        },
        onNextProblem: _transitionToAnswerDisplay,
        onInitializeControllers: _initializeControllersForCurrentProblem,
      ),
    );
  }

  void _transitionToNextProblem() {
    developer.log('[LearningScreen] _transitionToNextProblem called, current index: $_currentIndex');

    final problems = ref.read(wordProblemsProvider);
    final nextIndex = _findNextIncompleteProblem(problems);

    if (nextIndex == null) {
      // 全ての問題が完了した場合
      developer.log('[LearningScreen] All problems completed, transitioning to completed state');
      setState(() {
        _currentState = LearningState.completed;
      });
      return;
    }

    // 次の未完了問題に遷移
    _moveToNextProblem(nextIndex);
  }

  /// 次の未完了問題のインデックスを検索
  int? _findNextIncompleteProblem(List<WordProblem> problems) {
    // 現在の問題の次から検索
    for (int i = _currentIndex + 1; i < problems.length; i++) {
      if (!problems[i].isCompleted) {
        return i;
      }
    }
    
    // 見つからない場合、先頭から現在の問題まで検索（循環）
    for (int i = 0; i < _currentIndex; i++) {
      if (!problems[i].isCompleted) {
        return i;
      }
    }
    
    // 全ての問題が完了している場合
    return null;
  }

  /// 指定された問題インデックスに移動
  void _moveToNextProblem(int newIndex) {
    developer.log('[LearningScreen] Moving to problem $newIndex');
    
    // 新しい問題のユーザー入力をクリア
    ref.read(wordProblemsProvider.notifier).clearUserInputs(newIndex);
    
    setState(() {
      _currentIndex = newIndex;
      _currentState = LearningState.problemDisplay;
    });
    
    developer.log('[LearningScreen] Moved to problem $_currentIndex');
    
    // 既存のコントローラーを破棄して新しい問題用に再初期化
    _disposeControllers();
  }

  void _transitionToAnswerDisplay() {
    developer.log('[LearningScreen] Transitioning to answer display state');
    setState(() {
      _currentState = LearningState.answerDisplay;
    });
  }

  void _disposeControllers() {
    developer.log(
        '[LearningScreen] _disposeControllers called, disposing ${_answerControllers.length} controllers');
    for (final controller in _answerControllers) {
      controller.dispose();
    }
    for (final focusNode in _answerFocusNodes) {
      focusNode.dispose();
    }
    _answerControllers.clear();
    _answerFocusNodes.clear();
    developer.log('[LearningScreen] Controllers cleared');
  }

  @override
  void dispose() {
    _disposeControllers();
    super.dispose();
  }
}

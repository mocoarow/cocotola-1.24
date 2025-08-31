import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/word_problem_provider.dart';
import '../widgets/custom_keyboard.dart';
import '../models/word_problem.dart';
import '../ui/widgets/blank_widget.dart';
import '../ui/widgets/hints_widget.dart';
import '../ui/widgets/problem_content_widget.dart';
import 'dart:developer' as developer;

enum LearningState {
  problemDisplay,  // 問題表示状態
  answerDisplay,   // 答え表示状態（解説表示）
  completed        // 完了状態
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
  int _currentBlankIndex = 0;
  int _cursorPosition = 0;
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
    final currentProblem = problems[_currentIndex];
    
    final englishWords = currentProblem.english
        .replaceAll('.', ' .')
        .split(' ')
        .where((word) => word.isNotEmpty)
        .toList();

    // 複数の空欄のインデックスを取得
    final blankIndices = <int>[];
    for (int i = 0; i < englishWords.length; i++) {
      if (englishWords[i] == '___') {
        blankIndices.add(i);
      }
    }
    
    return Scaffold(
      appBar: AppBar(
        title: const Text('単語学習'),
      ),
      body: ProblemContentWidget(
        currentProblem: currentProblem,
        englishWords: englishWords,
        blankIndices: blankIndices,
        buildBlankWidget: (wordIndex, blankIndex, problem) => BlankWidget(
          wordIndex: wordIndex,
          blankIndex: blankIndex,
          problem: problem,
          controller: _answerControllers[blankIndex],
          focusNode: _answerFocusNodes[blankIndex],
          onChanged: (value) {
            ref
                .read(wordProblemsProvider.notifier)
                .updateUserInput(_currentIndex, blankIndex, value);
            _checkAnswerAutomatically(blankIndex, value);
          },
          onTap: () {
            setState(() {
              _currentBlankIndex = blankIndex;
            });
          },
        ),
        buildHintsSection: (problem) => HintsWidget(problem: problem),
        buildKeyboard: _buildAnswerDisplayKeyboard, // キーボード非表示版
        buildActionButtons: _buildAnswerDisplayActionButtons, // 次へボタンのみ
      ),
    );
  }

  Widget _buildProblemDisplayScreen(List<WordProblem> problems) {
    final currentProblem = problems[_currentIndex];
    final requiredBlanks = currentProblem.blanks.length;
    developer.log(
        '[LearningScreen] Current problem needs $requiredBlanks blanks, we have ${_answerControllers.length}');

    // _answerControllersが初期化されていない場合は空の画面を返す
    if (_answerControllers.isEmpty ||
        _answerControllers.length < requiredBlanks) {
      developer.log('[LearningScreen] Controllers not ready, initializing...');
      _initializeControllersForCurrentProblem(problems);
      return const Scaffold(
        body: Center(child: CircularProgressIndicator()),
      );
    }

    final englishWords = currentProblem.english
        .replaceAll('.', ' .')
        .split(' ')
        .where((word) => word.isNotEmpty)
        .toList();

    // 複数の空欄のインデックスを取得
    final blankIndices = <int>[];
    for (int i = 0; i < englishWords.length; i++) {
      if (englishWords[i] == '___') {
        blankIndices.add(i);
      }
    }

    return Scaffold(
      appBar: AppBar(
        title: const Text('単語学習'),
      ),
      body: ProblemContentWidget(
        currentProblem: currentProblem,
        englishWords: englishWords,
        blankIndices: blankIndices,
        buildBlankWidget: (wordIndex, blankIndex, problem) => BlankWidget(
          wordIndex: wordIndex,
          blankIndex: blankIndex,
          problem: problem,
          controller: _answerControllers[blankIndex],
          focusNode: _answerFocusNodes[blankIndex],
          onChanged: (value) {
            ref
                .read(wordProblemsProvider.notifier)
                .updateUserInput(_currentIndex, blankIndex, value);
            _checkAnswerAutomatically(blankIndex, value);
          },
          onTap: () {
            setState(() {
              _currentBlankIndex = blankIndex;
            });
          },
        ),
        buildHintsSection: (problem) => HintsWidget(problem: problem),
        buildKeyboard: _buildKeyboard,
        buildActionButtons: _buildActionButtons,
      ),
    );
  }

  Widget _buildKeyboard(WordProblem problem) {
    if (problem.isCompleted) {
      return const SizedBox.shrink();
    }

    return CustomKeyboard(
      onKeyPressed: (key) {
        if (_currentBlankIndex < _answerControllers.length) {
          final controller = _answerControllers[_currentBlankIndex];
          final text = controller.text;
          final newText = text.substring(0, _cursorPosition) +
              key +
              text.substring(_cursorPosition);
          controller.text = newText;
          _cursorPosition++;
          controller.selection = TextSelection.fromPosition(
            TextPosition(offset: _cursorPosition),
          );
          
          // プロバイダーの状態を更新
          ref
              .read(wordProblemsProvider.notifier)
              .updateUserInput(_currentIndex, _currentBlankIndex, newText);
          
          // 自動チェック機能
          _checkAnswerAutomatically(_currentBlankIndex, newText);
        }
      },
      onDelete: () {
        if (_currentBlankIndex < _answerControllers.length &&
            _cursorPosition > 0) {
          final controller = _answerControllers[_currentBlankIndex];
          final text = controller.text;
          final newText = text.substring(0, _cursorPosition - 1) +
              text.substring(_cursorPosition);
          controller.text = newText;
          _cursorPosition--;
          controller.selection = TextSelection.fromPosition(
            TextPosition(offset: _cursorPosition),
          );
          
          // プロバイダーの状態を更新
          ref
              .read(wordProblemsProvider.notifier)
              .updateUserInput(_currentIndex, _currentBlankIndex, newText);
        }
      },
      onMoveLeft: () {
        if (_cursorPosition > 0) {
          _cursorPosition--;
          if (_currentBlankIndex < _answerControllers.length) {
            _answerControllers[_currentBlankIndex].selection =
                TextSelection.fromPosition(
              TextPosition(offset: _cursorPosition),
            );
            _answerFocusNodes[_currentBlankIndex].requestFocus();
          }
        }
      },
      onMoveRight: () {
        if (_currentBlankIndex < _answerControllers.length) {
          final controller = _answerControllers[_currentBlankIndex];
          if (_cursorPosition < controller.text.length) {
            _cursorPosition++;
            controller.selection = TextSelection.fromPosition(
              TextPosition(offset: _cursorPosition),
            );
            _answerFocusNodes[_currentBlankIndex].requestFocus();
          }
        }
      },
    );
  }

  Widget _buildActionButtons(WordProblem problem) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          ElevatedButton(
            onPressed: () {
              ref
                  .read(wordProblemsProvider.notifier)
                  .markAsSkipped(_currentIndex);
              _transitionToAnswerDisplay();
            },
            child: const Text('答えを見る'),
          ),
          if (!problem.isCompleted) ...[
            ElevatedButton(
              onPressed: _checkCurrentAnswers,
              child: const Text('確認'),
            ),
          ],
          if (problem.isCompleted) ...[
            ElevatedButton(
              onPressed: _transitionToAnswerDisplay,
              child: const Text('次へ'),
            ),
          ],
        ],
      ),
    );
  }

  Widget _buildAnswerDisplayKeyboard(WordProblem problem) {
    // 解説表示時はキーボードを表示しない
    return const SizedBox.shrink();
  }

  Widget _buildAnswerDisplayActionButtons(WordProblem problem) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          ElevatedButton(
            onPressed: _transitionToNextProblem,
            child: const Text('次へ'),
          ),
        ],
      ),
    );
  }

  void _checkCurrentAnswers() {
    final problem = ref.read(wordProblemsProvider)[_currentIndex];

    for (int i = 0; i < problem.blanks.length; i++) {
      final userInput = _answerControllers[i].text.trim();
      if (userInput.isNotEmpty && !problem.blanks[i].isAnswered) {
        ref
            .read(wordProblemsProvider.notifier)
            .checkAnswer(_currentIndex, i, userInput);
      }
    }
    
    // 全ての空欄が完了したかチェック
    _checkIfAllBlanksCompleted();
  }

  void _checkAnswerAutomatically(int blankIndex, String value) {
    final problem = ref.read(wordProblemsProvider)[_currentIndex];
    final blank = problem.blanks[blankIndex];
    
    // 既に回答済みの場合は何もしない
    if (blank.isAnswered) return;
    
    final trimmedValue = value.trim();
    if (trimmedValue.isNotEmpty) {
      final isCorrect = trimmedValue.toLowerCase() == blank.answer.toLowerCase();
      
      if (isCorrect) {
        developer.log('[LearningScreen] Auto-check: Correct answer detected for blank $blankIndex');
        ref
            .read(wordProblemsProvider.notifier)
            .checkAnswer(_currentIndex, blankIndex, trimmedValue);
        
        // 全ての空欄が正解かチェック
        _checkIfAllBlanksCompleted();
      }
    }
  }

  void _checkIfAllBlanksCompleted() {
    // 少し遅延を入れてからチェック（Riverpodの状態更新を待つため）
    WidgetsBinding.instance.addPostFrameCallback((_) {
      final updatedProblem = ref.read(wordProblemsProvider)[_currentIndex];
      
      if (updatedProblem.isCompleted) {
        developer.log('[LearningScreen] All blanks completed, transitioning to answer display');
        setState(() {
          _currentState = LearningState.answerDisplay;
        });
      }
    });
  }

  void _transitionToNextProblem() {
    developer.log(
        '[LearningScreen] _transitionToNextProblem called, current index: $_currentIndex');

    final problems = ref.read(wordProblemsProvider);
    final currentProblemIndex = _currentIndex;

    // 未完了の問題を探す（現在の問題の次から）
    int? nextIncompleteIndex;
    for (int i = 1; i < problems.length; i++) {
      final checkIndex = (currentProblemIndex + i) % problems.length;
      if (!problems[checkIndex].isCompleted) {
        nextIncompleteIndex = checkIndex;
        break;
      }
    }

    // 未完了の問題がない場合は、完了状態に遷移
    if (nextIncompleteIndex == null) {
      developer.log('[LearningScreen] All problems completed, transitioning to completed state');
      setState(() {
        _currentState = LearningState.completed;
      });
      return;
    }

    final newIndex = nextIncompleteIndex;

    // 新しい問題のユーザー入力をクリア
    ref.read(wordProblemsProvider.notifier).clearUserInputs(newIndex);
    developer.log('[LearningScreen] Cleared user inputs for new problem');

    setState(() {
      _currentIndex = newIndex;
      _currentBlankIndex = 0;
      _cursorPosition = 0;
      _currentState = LearningState.problemDisplay;  // 問題表示状態に遷移
    });

    developer.log('[LearningScreen] New index: $_currentIndex, state: $_currentState');

    // 既存のコントローラーを破棄して新しい問題用に再初期化をフォース
    _disposeControllers();
    
    // 次のbuildで新しいコントローラーが作成される
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

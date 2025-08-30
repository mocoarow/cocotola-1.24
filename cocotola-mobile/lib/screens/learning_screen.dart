import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/word_problem_provider.dart';
import '../widgets/custom_keyboard.dart';
import '../models/word_problem.dart';
import 'dart:developer' as developer;

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

  @override
  void initState() {
    super.initState();
    developer.log('[LearningScreen] initState called');
    // 初期化は最初のbuildで行う
  }

  void _initializeControllersForCurrentProblem(List<WordProblem> problems) {
    developer.log('[LearningScreen] _initializeControllersForCurrentProblem called for currentIndex: $_currentIndex');
    if (problems.isNotEmpty && _currentIndex < problems.length) {
      final maxBlanks = problems[_currentIndex].blanks.length;
      developer.log('[LearningScreen] Max blanks for current problem: $maxBlanks');
      
      // 既存のコントローラーを破棄
      _disposeControllers();
      
      // 新しいコントローラーを作成
      final currentProblem = problems[_currentIndex];
      _answerControllers = List.generate(maxBlanks, (index) {
        // stateのuserInputを確認し、正解済みでない場合は空文字にする
        final blank = currentProblem.blanks[index];
        final initialText = (blank.isAnswered && blank.isCorrect) ? blank.answer : '';
        developer.log('[LearningScreen] Controller[$index] initialized with: "$initialText"');
        return TextEditingController(text: initialText);
      });
      _answerFocusNodes = List.generate(maxBlanks, (index) => FocusNode());
      developer.log('[LearningScreen] Controllers initialized: ${_answerControllers.length}');
      
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
    developer.log('[LearningScreen] build called - problems count: ${problems.length}, currentIndex: $_currentIndex');
    developer.log('[LearningScreen] controllers length: ${_answerControllers.length}');
    
    if (problems.isEmpty) {
      developer.log('[LearningScreen] No problems available');
      return const Center(child: Text('お疲れ様でした！'));
    }

    final currentProblem = problems[_currentIndex];
    final requiredBlanks = currentProblem.blanks.length;
    developer.log('[LearningScreen] Current problem needs $requiredBlanks blanks, we have ${_answerControllers.length}');

    // _answerControllersが初期化されていない場合は空の画面を返す
    if (_answerControllers.isEmpty || _answerControllers.length < requiredBlanks) {
      developer.log('[LearningScreen] Controllers not ready, initializing...');
      _initializeControllersForCurrentProblem(problems);
      return const Scaffold(
        body: Center(child: CircularProgressIndicator()),
      );
    }

    // 全問完了かチェック
    final allCompleted = problems.every((problem) => problem.isCompleted);
    if (allCompleted) {
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

    developer.log('[LearningScreen] Rendering problem UI for currentIndex: $_currentIndex');
    
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
      body: Column(
        children: [
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: Text(
              currentProblem.japanese,
              style: Theme.of(context).textTheme.headlineSmall,
            ),
          ),
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: Wrap(
              alignment: WrapAlignment.center,
              children: [
                for (int i = 0; i < englishWords.length; i++)
                  if (blankIndices.contains(i))
                    _buildBlankWidget(i, blankIndices.indexOf(i), currentProblem)
                  else
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 4.0, vertical: 2.0),
                      child: Text(
                        englishWords[i],
                        style: const TextStyle(fontSize: 16),
                      ),
                    ),
              ],
            ),
          ),
          _buildHintsSection(currentProblem),
          const Spacer(),
          _buildKeyboard(currentProblem),
          _buildActionButtons(currentProblem),
        ],
      ),
    );
  }

  Widget _buildBlankWidget(int wordIndex, int blankIndex, WordProblem problem) {
    final blank = problem.blanks[blankIndex];
    final inputWidth = (blank.answer.length * 20.0).clamp(100.0, 200.0);
    
    return Container(
      margin: const EdgeInsets.symmetric(horizontal: 4.0, vertical: 2.0),
      width: inputWidth,
      child: blank.isAnswered && blank.isCorrect
          ? Container(
              padding: const EdgeInsets.all(8.0),
              decoration: BoxDecoration(
                border: Border.all(color: Colors.green, width: 2),
                borderRadius: BorderRadius.circular(4),
                color: Colors.green.shade50,
              ),
              child: Text(
                blank.answer,
                textAlign: TextAlign.center,
                style: const TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.bold,
                  color: Colors.green,
                ),
              ),
            )
          : TextField(
              controller: _answerControllers[blankIndex],
              textAlign: TextAlign.center,
              decoration: InputDecoration(
                border: const OutlineInputBorder(),
                fillColor: blank.isAnswered && !blank.isCorrect 
                    ? Colors.red.shade50 
                    : null,
                filled: blank.isAnswered && !blank.isCorrect,
              ),
              focusNode: _answerFocusNodes[blankIndex],
              enabled: !(blank.isAnswered && blank.isCorrect),
              onChanged: (value) {
                ref.read(wordProblemsProvider.notifier)
                    .updateUserInput(_currentIndex, blankIndex, value);
              },
              onTap: () {
                setState(() {
                  _currentBlankIndex = blankIndex;
                });
              },
            ),
    );
  }

  Widget _buildHintsSection(WordProblem problem) {
    final correctBlanks = problem.blanks.where((blank) => blank.isAnswered && blank.isCorrect).toList();
    
    if (correctBlanks.isEmpty) {
      return const SizedBox.shrink();
    }
    
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        children: correctBlanks.map((blank) => Container(
          margin: const EdgeInsets.only(bottom: 8.0),
          padding: const EdgeInsets.all(12.0),
          decoration: BoxDecoration(
            color: Colors.green.shade50,
            borderRadius: BorderRadius.circular(8),
            border: Border.all(color: Colors.green.shade200),
          ),
          child: Column(
            children: [
              Text(
                '正解: ${blank.answer}',
                style: const TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.bold,
                  color: Colors.green,
                ),
              ),
              const SizedBox(height: 4),
              Text(
                'ヒント: ${blank.hint}',
                style: const TextStyle(fontSize: 14),
              ),
            ],
          ),
        )).toList(),
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
          controller.text = text.substring(0, _cursorPosition) +
              key +
              text.substring(_cursorPosition);
          _cursorPosition++;
          controller.selection = TextSelection.fromPosition(
            TextPosition(offset: _cursorPosition),
          );
        }
      },
      onDelete: () {
        if (_currentBlankIndex < _answerControllers.length && _cursorPosition > 0) {
          final controller = _answerControllers[_currentBlankIndex];
          final text = controller.text;
          controller.text = text.substring(0, _cursorPosition - 1) +
              text.substring(_cursorPosition);
          _cursorPosition--;
          controller.selection = TextSelection.fromPosition(
            TextPosition(offset: _cursorPosition),
          );
        }
      },
      onMoveLeft: () {
        if (_cursorPosition > 0) {
          _cursorPosition--;
          if (_currentBlankIndex < _answerControllers.length) {
            _answerControllers[_currentBlankIndex].selection = TextSelection.fromPosition(
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
              ref.read(wordProblemsProvider.notifier).markAsSkipped(_currentIndex);
              _moveToNextProblem();
            },
            child: const Text('答えを見る'),
          ),
          if (!problem.isCompleted) ...[
            ElevatedButton(
              onPressed: _checkCurrentAnswers,
              child: const Text('確認'),
            ),
          ],
          if (problem.hasAnyCorrectAnswer) ...[
            ElevatedButton(
              onPressed: _moveToNextProblem,
              child: const Text('次へ'),
            ),
          ],
        ],
      ),
    );
  }

  void _checkCurrentAnswers() {
    final problem = ref.read(wordProblemsProvider)[_currentIndex];
    
    for (int i = 0; i < problem.blanks.length; i++) {
      final userInput = _answerControllers[i].text.trim();
      if (userInput.isNotEmpty && !problem.blanks[i].isAnswered) {
        ref.read(wordProblemsProvider.notifier)
            .checkAnswer(_currentIndex, i, userInput);
      }
    }
  }

  void _moveToNextProblem() {
    developer.log('[LearningScreen] _moveToNextProblem called, current index: $_currentIndex');
    
    final newIndex = (_currentIndex + 1) % ref.read(wordProblemsProvider).length;
    
    setState(() {
      _currentIndex = newIndex;
      _currentBlankIndex = 0;
      _cursorPosition = 0;
    });
    
    developer.log('[LearningScreen] New index: $_currentIndex');
    
    // 新しい問題のユーザー入力をクリア
    ref.read(wordProblemsProvider.notifier).clearUserInputs(newIndex);
    developer.log('[LearningScreen] Cleared user inputs for new problem');
    
    // コントローラーは次のbuildで自動的に再初期化される
  }

  void _disposeControllers() {
    developer.log('[LearningScreen] _disposeControllers called, disposing ${_answerControllers.length} controllers');
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

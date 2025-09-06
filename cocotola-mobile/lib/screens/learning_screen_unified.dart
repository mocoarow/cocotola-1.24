import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../view_models/app_state_manager.dart';
import '../models/problem_base.dart';
import '../widgets/problem_content_widget.dart';
import '../widgets/blank_widget.dart';
import '../widgets/hints_widget.dart';
import '../widgets/custom_keyboard.dart';
import '../widgets/memorization_question_widget.dart';
import '../widgets/memorization_answer_widget.dart';
import 'dart:developer' as developer;

class LearningScreenUnified extends ConsumerWidget {
  const LearningScreenUnified({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final appState = ref.watch(appStateProvider);
    final appStateManager = ref.read(appStateProvider.notifier);

    if (appState.selectedProblemSet == null || appState.currentProblem == null) {
      return Scaffold(
        appBar: AppBar(title: const Text('エラー')),
        body: const Center(child: Text('問題が見つかりません')),
      );
    }

    final currentProblem = appState.currentProblem!;

    // 問題タイプによって処理を分ける
    if (currentProblem.type == ProblemType.memorization) {
      return _buildMemorizationProblem(context, appState, appStateManager);
    } else {
      return _buildWordProblem(context, appState, appStateManager);
    }
  }

  /// 暗記問題の画面構築
  Widget _buildMemorizationProblem(BuildContext context, AppState appState, AppStateManager appStateManager) {
    final currentProblem = appState.currentProblem!;
    final memorizationProblem = currentProblem.memorizationProblem!;
    
    return Scaffold(
      appBar: AppBar(
        title: Text(appState.selectedProblemSet!.title),
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () => appStateManager.returnToMenu(),
          tooltip: 'メニューに戻る',
        ),
      ),
      body: appState.learningState == LearningPhase.problemDisplay
        ? MemorizationQuestionWidget(
            problem: memorizationProblem,
            onShowAnswer: () => appStateManager.transitionToAnswerDisplay(),
          )
        : MemorizationAnswerWidget(
            problem: memorizationProblem,
            onAnswer: (wasCorrect) => appStateManager.handleMemorizationAnswer(wasCorrect),
          ),
    );
  }

  /// 英単語問題の画面構築
  Widget _buildWordProblem(BuildContext context, AppState appState, AppStateManager appStateManager) {
    final currentProblem = appState.currentProblem!;
    final wordProblem = currentProblem.wordProblem!;
    final englishWords = wordProblem.englishWords;
    final blankIndices = wordProblem.blankIndices;

    // コントローラーが初期化されていない場合
    if (appState.answerControllers.isEmpty) {
      return Scaffold(
        appBar: AppBar(title: Text(appState.selectedProblemSet!.title)),
        body: const Center(child: CircularProgressIndicator()),
      );
    }

    return Scaffold(
      appBar: AppBar(
        title: Text(appState.selectedProblemSet!.title),
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () => appStateManager.returnToMenu(),
          tooltip: 'メニューに戻る',
        ),
      ),
      body: Column(
        children: [
          // 問題表示部分
          Expanded(
            child: ProblemContentWidget(
              currentProblem: wordProblem,
              englishWords: englishWords,
              blankIndices: blankIndices,
              buildBlankWidget: (wordIndex, blankIndex, problem) => BlankWidget(
                wordIndex: wordIndex,
                blankIndex: blankIndex,
                problem: problem,
                controller: appState.answerControllers[blankIndex],
                focusNode: appState.answerFocusNodes[blankIndex],
                readOnly: false,
                onChanged: (value) {
                  developer.log('[LearningScreenUnified] Physical keyboard input: blank=$blankIndex, value="$value"');
                  
                  // カーソル位置を同期
                  final cursorPos = appState.answerControllers[blankIndex].selection.end;
                  appStateManager.updateFocusAndCursor(blankIndex, cursorPos);
                  
                  // 状態管理を通じて回答を処理
                  appStateManager.handleAnswer(blankIndex, value);
                },
                onTap: () => _handleBlankTap(blankIndex, appStateManager, appState),
              ),
              buildHintsSection: (problem) => HintsWidget(problem: problem),
              buildKeyboard: (problem) => const SizedBox.shrink(), // キーボードは下部に移動
              buildActionButtons: (problem) => const SizedBox.shrink(), // アクションボタンも下部に移動
            ),
          ),
          // カスタムキーボード（画面下部に固定）
          _buildKeyboard(wordProblem, appStateManager),
          // アクションボタン（画面最下部に固定）
          _buildActionButtons(wordProblem, appStateManager),
        ],
      ),
    );
  }

  Widget _buildKeyboard(problem, AppStateManager appStateManager) {
    if (problem.isCompleted) {
      return const SizedBox.shrink();
    }

    return Consumer(
      builder: (context, ref, child) {
        final appState = ref.watch(appStateProvider);
        
        // 現在フォーカスされている空欄のテキストとカーソル位置を取得
        String currentText = '';
        int cursorPosition = 0;
        
        if (appState.currentBlankIndex < appState.answerControllers.length) {
          currentText = appState.answerControllers[appState.currentBlankIndex].text;
          cursorPosition = appState.cursorPosition;
        }
        
        return CustomKeyboard(
          onKeyPressed: (key) => appStateManager.handleCustomKeyboardInput(key),
          onDelete: () => appStateManager.handleCustomKeyboardDelete(),
          onMoveLeft: () => appStateManager.moveCursor(-1),
          onMoveRight: () => appStateManager.moveCursor(1),
          currentText: currentText,
          cursorPosition: cursorPosition,
          showCursorIndicator: true, // 実際のアプリでは表示
        );
      },
    );
  }

  Widget _buildActionButtons(problem, AppStateManager appStateManager) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          ElevatedButton(
            onPressed: () {
              // 問題を後回しにして答え表示画面へ
              appStateManager.handleShowAnswerForWordProblem();
            },
            child: const Text('答えを見る'),
          ),
          if (problem.isCompleted) ...[
            ElevatedButton(
              onPressed: () => appStateManager.transitionToAnswerDisplay(),
              child: const Text('次へ'),
            ),
          ],
        ],
      ),
    );
  }

  void _handleBlankTap(int blankIndex, AppStateManager appStateManager, AppState appState) {
    final controller = appState.answerControllers[blankIndex];
    final cursorPos = controller.selection.isValid
        ? controller.selection.end
        : controller.text.length;

    developer.log('[LearningScreenUnified] Blank tapped: $blankIndex, cursor: $cursorPos');
    appStateManager.updateFocusAndCursor(blankIndex, cursorPos);
  }
}
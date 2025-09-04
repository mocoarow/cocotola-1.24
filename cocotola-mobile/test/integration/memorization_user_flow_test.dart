import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:cocotola/main.dart' as app;
import 'package:cocotola/screens/app_shell.dart';
import 'package:cocotola/screens/problem_set_selection_screen_unified.dart';
import 'package:cocotola/screens/learning_screen_unified.dart';
import 'package:cocotola/widgets/memorization_question_widget.dart';
import 'package:cocotola/widgets/memorization_answer_widget.dart';
import 'package:cocotola/providers/problem_set_provider.dart';

void main() {
  group('暗記問題ユーザーフロー統合テスト', () {
    testWidgets('暗記問題セット選択から学習完了までのフロー', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // AppShellが表示され、問題セット選択画面が表示されることを確認
      expect(find.byType(AppShell), findsOneWidget);
      expect(find.byType(ProblemSetSelectionScreenUnified), findsOneWidget);

      // デバッグ: プロバイダーの状態を直接確認
      final container = ProviderScope.containerOf(tester.element(find.byType(ProblemSetSelectionScreenUnified)));
      final problemSets = container.read(problemSetsProvider);
      print('Debug: ProblemSets from provider: ${problemSets.length}');
      for (int i = 0; i < problemSets.length; i++) {
        print('Debug: ProblemSet[$i]: ${problemSets[i].title}');
        try {
          print('Debug: ProblemSet[$i] problemCount: ${problemSets[i].problemCount}');
          print('Debug: ProblemSet[$i] cefrLevel: ${problemSets[i].cefrLevel}');
        } catch (e) {
          print('Debug: Error accessing ProblemSet[$i] properties: $e');
        }
      }
      
      // デバッグ: 実際に表示されているカードのタイトルを確認
      final allCards = find.byType(Card);
      print('Debug: Found ${tester.widgetList(allCards).length} cards');
      
      // カード内のタイトルテキストを確認
      for (int i = 0; i < tester.widgetList(allCards).length; i++) {
        print('Debug: Card $i exists');
      }
      
      // 全てのテキストを確認
      final allTexts = find.byType(Text);
      for (final textWidget in tester.widgetList<Text>(allTexts)) {
        if (textWidget.data != null) {
          print('Debug: Found text: "${textWidget.data}"');
        }
      }

      // 暗記問題セット（単語暗記）を探す
      await tester.pumpAndSettle(const Duration(milliseconds: 300));
      final memorizationCard = find.text('単語暗記');
      expect(memorizationCard, findsOneWidget);

      // スクロールして要素を画面内に表示
      await tester.scrollUntilVisible(memorizationCard, 100.0);
      
      // 暗記問題セットをタップ
      await tester.tap(memorizationCard);
      await tester.pumpAndSettle();

      // 学習画面に遷移することを確認
      expect(find.byType(LearningScreenUnified), findsOneWidget);
      
      // 暗記問題の問題表示ウィジェットが表示されることを確認
      expect(find.byType(MemorizationQuestionWidget), findsOneWidget);
      
      // 問題文が表示されることを確認
      expect(find.text('apple'), findsOneWidget);
    });

    testWidgets('暗記問題の問題→答え表示→回答のフロー', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 暗記問題セット（単語暗記）を選択
      await tester.pumpAndSettle(const Duration(milliseconds: 300));
      final memorizationCard = find.text('単語暗記');
      await tester.scrollUntilVisible(memorizationCard, 100.0);
      await tester.tap(memorizationCard);
      await tester.pumpAndSettle();

      // 問題表示画面で「答えを見る」ボタンを押す
      final showAnswerButton = find.text('答えを見る');
      expect(showAnswerButton, findsOneWidget);
      await tester.tap(showAnswerButton);
      await tester.pumpAndSettle();

      // 答え表示ウィジェットが表示されることを確認
      expect(find.byType(MemorizationAnswerWidget), findsOneWidget);
      
      // 答えが表示されることを確認
      expect(find.text('りんご'), findsOneWidget);
      
      // 「できた」「できなかった」ボタンが表示されることを確認
      expect(find.text('できた'), findsOneWidget);
      expect(find.text('できなかった'), findsOneWidget);
    });

    testWidgets('「できた」を選択した場合の次問題遷移', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 暗記問題セット（単語暗記）を選択
      await tester.pumpAndSettle(const Duration(milliseconds: 300));
      final memorizationCard = find.text('単語暗記');
      await tester.scrollUntilVisible(memorizationCard, 100.0);
      await tester.tap(memorizationCard);
      await tester.pumpAndSettle();

      // 「答えを見る」ボタンを押す
      await tester.tap(find.text('答えを見る'));
      await tester.pumpAndSettle();

      // 「できた」ボタンを押す
      final dekitaButton = find.text('できた');
      await tester.tap(dekitaButton);
      await tester.pumpAndSettle();

      // 次の問題に遷移することを確認
      // （複数問題がある場合、次の問題が表示される）
      expect(find.byType(MemorizationQuestionWidget), findsOneWidget);
      
      // 2番目の問題が表示されることを確認
      expect(find.text('book'), findsOneWidget);
    });

    testWidgets('「できなかった」を選択した場合の次問題遷移', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 暗記問題セット（単語暗記）を選択
      await tester.pumpAndSettle(const Duration(milliseconds: 300));
      final memorizationCard = find.text('単語暗記');
      await tester.scrollUntilVisible(memorizationCard, 100.0);
      await tester.tap(memorizationCard);
      await tester.pumpAndSettle();

      // 「答えを見る」ボタンを押す
      await tester.tap(find.text('答えを見る'));
      await tester.pumpAndSettle();

      // 「できなかった」ボタンを押す
      final dekinakattaButton = find.text('できなかった');
      await tester.tap(dekinakattaButton);
      await tester.pumpAndSettle();

      // 次の問題に遷移することを確認
      expect(find.byType(MemorizationQuestionWidget), findsOneWidget);
      
      // 2番目の問題が表示されることを確認
      expect(find.text('book'), findsOneWidget);
    });

    testWidgets('暗記問題でのメニュー戻り機能', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 暗記問題セット（単語暗記）を選択
      await tester.pumpAndSettle(const Duration(milliseconds: 300));
      final memorizationCard = find.text('単語暗記');
      await tester.scrollUntilVisible(memorizationCard, 100.0);
      await tester.tap(memorizationCard);
      await tester.pumpAndSettle();

      // 学習画面が表示されることを確認
      expect(find.byType(MemorizationQuestionWidget), findsOneWidget);
      
      // AppBarの戻るボタンを押す
      final backButton = find.byIcon(Icons.arrow_back);
      expect(backButton, findsWidgets);
      await tester.tap(backButton.first);
      await tester.pumpAndSettle();

      // 問題セット選択画面に戻ったことを確認
      expect(find.byType(ProblemSetSelectionScreenUnified), findsOneWidget);
      expect(find.byType(Card), findsWidgets);
    });

    testWidgets('暗記問題の答え表示画面からメニュー戻り', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 暗記問題セット（単語暗記）を選択
      await tester.pumpAndSettle(const Duration(milliseconds: 300));
      final memorizationCard = find.text('単語暗記');
      await tester.scrollUntilVisible(memorizationCard, 100.0);
      await tester.tap(memorizationCard);
      await tester.pumpAndSettle();

      // 「答えを見る」ボタンを押して答え表示画面に移動
      final showAnswerButton = find.text('答えを見る');
      await tester.tap(showAnswerButton);
      await tester.pumpAndSettle();
      
      // 答え表示画面が表示されることを確認
      expect(find.byType(MemorizationAnswerWidget), findsOneWidget);
      
      // AppBarの戻るボタンを押す
      final backButton = find.byIcon(Icons.arrow_back);
      expect(backButton, findsWidgets);
      await tester.tap(backButton.first);
      await tester.pumpAndSettle();

      // 問題セット選択画面に戻ったことを確認
      expect(find.byType(ProblemSetSelectionScreenUnified), findsOneWidget);
      expect(find.byType(Card), findsWidgets);
    });

    testWidgets('フレーズ暗記問題セットのテスト', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // フレーズ暗記問題セットを探す
      await tester.pumpAndSettle(const Duration(milliseconds: 300));
      final phraseCard = find.text('フレーズ暗記');
      expect(phraseCard, findsOneWidget);

      // スクロールして要素を画面内に表示
      await tester.scrollUntilVisible(phraseCard, 100.0);
      
      // フレーズ暗記問題セットをタップ
      await tester.tap(phraseCard);
      await tester.pumpAndSettle();

      // 学習画面に遷移することを確認
      expect(find.byType(LearningScreenUnified), findsOneWidget);
      expect(find.byType(MemorizationQuestionWidget), findsOneWidget);
      
      // フレーズ問題が表示されることを確認
      expect(find.text('Good morning'), findsOneWidget);
    });
  });
}
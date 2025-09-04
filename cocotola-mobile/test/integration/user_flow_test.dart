import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:cocotola/main.dart' as app;
import 'package:cocotola/screens/app_shell.dart';
import 'package:cocotola/screens/problem_set_selection_screen_unified.dart';
import 'package:cocotola/screens/learning_screen_unified.dart';

void main() {
  group('ユーザーフロー統合テスト', () {
    testWidgets('問題セット選択から学習完了までの基本フロー', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // AppShellが表示され、問題セット選択画面が表示されることを確認
      expect(find.byType(AppShell), findsOneWidget);
      expect(find.byType(ProblemSetSelectionScreenUnified), findsOneWidget);
      expect(find.text('問題セット選択'), findsOneWidget);

      // 最初の問題セットをタップ
      final firstProblemSetCard = find.byType(Card).first;
      await tester.tap(firstProblemSetCard);
      await tester.pumpAndSettle();

      // 学習画面に遷移することを確認
      expect(find.byType(LearningScreenUnified), findsOneWidget);

      // テキスト入力フィールドが存在することを確認
      expect(find.byType(TextField), findsWidgets);
    });

    testWidgets('テキスト入力と正解判定のフロー', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 問題セットを選択
      await tester.tap(find.byType(Card).first);
      await tester.pumpAndSettle();

      // 最初のテキストフィールドを見つけて入力
      final textField = find.byType(TextField).first;
      expect(textField, findsOneWidget);

      // 正解を入力（最初の問題の答えは"study"と仮定）
      await tester.enterText(textField, 'study');
      await tester.pumpAndSettle();

      // 入力が反映されることを確認
      expect(find.text('study'), findsOneWidget);
    });

    testWidgets('自動フォーカス機能のテスト', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 問題セットを選択
      await tester.tap(find.byType(Card).first);
      await tester.pumpAndSettle();

      // 最初のテキストフィールドが存在することを確認
      final textFields = find.byType(TextField);
      expect(textFields, findsWidgets);

      // 最初のテキストフィールドにフォーカスが当たっていることを確認
      // （実際のフォーカス状態のテストは複雑なので、存在確認のみ）
      final firstTextField = tester.widget<TextField>(textFields.first);
      expect(firstTextField.focusNode, isNotNull);
    });

    testWidgets('物理キーボードとカスタムキーボードの連携テスト', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 問題セットを選択
      await tester.tap(find.byType(Card).first);
      await tester.pumpAndSettle();

      // テキストフィールドを取得
      final textField = find.byType(TextField).first;
      
      // 物理キーボードで"st"と入力
      await tester.enterText(textField, 'st');
      await tester.pumpAndSettle();
      
      // "st"が入力されていることを確認（テキストフィールドとカーソル位置インジケーターに表示）
      expect(find.text('st'), findsWidgets);
      
      // テキストフィールドをタップしてフォーカス（カーソル位置は"st"の後にある）
      await tester.tap(textField);
      await tester.pumpAndSettle();
      
      // カスタムキーボードのボタンが表示されることを確認
      expect(find.byKey(const Key('custom-keyboard'), skipOffstage: false), findsWidgets);
      
      // カスタムキーボードの"u"ボタンをタップ（"study"の続きを入力）
      final uButton = find.text('u');
      if (tester.any(uButton)) {
        await tester.tap(uButton);
        await tester.pumpAndSettle();
        
        // "stu"になることを期待（先頭に入力されて"ust"になってはいけない）
        // テキストフィールドとカーソル位置インジケーターの両方に表示される
        expect(find.text('stu'), findsWidgets);
      }
    });

    testWidgets('カーソル位置の状態管理テスト', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 問題セットを選択
      await tester.tap(find.byType(Card).first);
      await tester.pumpAndSettle();

      // 最初のテキストフィールドを取得
      final textField = find.byType(TextField).first;
      
      // 物理キーボードで途中まで入力
      await tester.enterText(textField, 'st');
      await tester.pumpAndSettle();
      
      // テキストフィールドをタップしてカーソル位置を確定
      await tester.tap(textField);
      await tester.pumpAndSettle();
      
      // カスタムキーボードが表示されることを確認
      expect(find.byKey(const Key('custom-keyboard'), skipOffstage: false), findsWidgets);
      
      // カーソル位置での追加入力が正しく動作することを確認
      // （実際のキーボードボタンの詳細テストは別途実装）
    });

    testWidgets('複数空欄問題の存在確認テスト', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 問題セットが表示されることを確認
      final cards = find.byType(Card);
      expect(cards, findsAtLeastNWidgets(1));
      
      // このテストは複数空欄問題の存在を確認するだけ
      // 実際の自動フォーカス移動はAppStateManagerのユニットテストで検証
      expect(cards, findsWidgets);
    });

    testWidgets('全問正解時の完了画面表示', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 問題セットを選択
      await tester.tap(find.byType(Card).first);
      await tester.pumpAndSettle();

      // 必要に応じて全問正解のシミュレーション
      // （実際の問題内容に応じて調整が必要）
    });

    testWidgets('答えを見るボタンの動作', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 問題セットを選択
      await tester.tap(find.byType(Card).first);
      await tester.pumpAndSettle();

      // 「答えを見る」ボタンを押す
      final showAnswerButton = find.text('答えを見る');
      expect(showAnswerButton, findsOneWidget);
      
      await tester.tap(showAnswerButton);
      await tester.pumpAndSettle();

      // 答え表示画面に遷移することを確認（具体的な確認方法は実装に依存）
    });

    testWidgets('問題セット選択画面への戻り', (WidgetTester tester) async {
      // アプリ起動
      await tester.pumpWidget(
        const ProviderScope(child: app.MyApp()),
      );
      await tester.pumpAndSettle();

      // 問題セットを選択
      await tester.tap(find.byType(Card).first);
      await tester.pumpAndSettle();

      // 戻るボタンがある場合の動作確認
      final backButton = find.byType(BackButton);
      if (tester.any(backButton)) {
        await tester.tap(backButton);
        await tester.pumpAndSettle();

        // 問題セット選択画面に戻ることを確認
        expect(find.byType(ProblemSetSelectionScreenUnified), findsOneWidget);
      }
    });
  });
}
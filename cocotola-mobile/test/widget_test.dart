// This is a basic Flutter widget test.
//
// To perform an interaction with a widget in your test, use the WidgetTester
// utility in the flutter_test package. For example, you can send tap and scroll
// gestures. You can also use WidgetTester to find child widgets in the widget
// tree, read text, and verify that the values of widget properties are correct.

import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:cocotola/main.dart';

void main() {
  testWidgets('Learning app basic smoke test', (WidgetTester tester) async {
    // Build our app with ProviderScope and trigger a frame.
    await tester.pumpWidget(
      const ProviderScope(
        child: MyApp(),
      ),
    );

    // Wait for any async initialization
    await tester.pumpAndSettle();

    // Verify that the app starts with learning screen elements
    // Look for common elements like AppBar or learning-related text
    expect(find.byType(AppBar), findsOneWidget);
    
    // The app should load without crashing
    expect(tester.takeException(), isNull);
  });
}

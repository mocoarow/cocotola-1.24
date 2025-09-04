# 開発ワークフロー - 回帰防止ガイド

## 🎯 目的

このガイドは、1つの修正が他の機能を破壊する回帰問題を防ぐためのワークフローを定義します。

## 📋 変更前チェックリスト

### 1. 現在の動作を記録

```shell
# 既存機能のテストを実行
flutter test

# 統合テストで主要フローを確認
flutter test test/integration/
```

### 2. 影響範囲の分析

- 変更する関数/クラスがどこから呼ばれているかを確認
- 依存関係を把握する
- 状態管理の影響を考慮する

## 🔧 変更中のプラクティス

### 1. 小さな変更を心がける

- 1つのPRで1つの機能に集中
- 大きな変更は複数のPRに分割

### 2. テストファーストアプローチ

- 既存のテストが通ることを確認
- 新機能のテストを先に書く
- 変更後もテストが通ることを確認

### 3. ログの活用

```dart
import 'dart:developer' as developer;

void someFunction() {
  developer.log('[ClassName] Method called with params: $params');
  // 実装
  developer.log('[ClassName] Method completed successfully');
}
```

## ✅ 変更後チェックリスト

### 1. 自動テスト実行

```bash
# 回帰テストスクリプト実行
./scripts/regression_test.sh
```

### 2. 手動テスト項目

- [ ] 問題セット選択が正常に動作する
- [ ] テキスト入力が正常に動作する
- [ ] 正解時の自動遷移が動作する
- [ ] 「答えを見る」ボタンが動作する
- [ ] 問題間の遷移が動作する
- [ ] 完了画面が正常に表示される
- [ ] 「もう一度挑戦」が動作する
- [ ] 戻るナビゲーションが動作する

### 3. パフォーマンステスト

- メモリリークがないかチェック
- UIの応答性を確認

## 🚨 回帰が発生した場合

### 1. 即座にロールバック

```shell
git revert <commit-hash>
```

### 2. 問題の分析

- どの変更が原因かを特定
- なぜテストで検出されなかったかを分析

### 3. テストの改善

- 見落とした部分のテストを追加
- 統合テストのカバレッジを向上

## 🛠️ ツール設定

### VS Code設定（推奨）

```json
{
  "editor.formatOnSave": true,
  "dart.flutterSdkPath": "/opt/flutter",
  "dart.runPubGetOnPubspecChanges": true,
  "dart.previewFlutterUiGuides": true
}
```

### Git Hooks（推奨）

```bash
# pre-commit hookを設定
echo '#!/bin/sh\nflutter analyze && flutter test' > .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
```

## 📊 品質メトリクス

定期的に以下を確認：

- テストカバレッジ率
- コード分析での警告数
- 実行時間（パフォーマンス回帰）

## 🎯 長期的改善

1. **状態管理の統一化**
   - AppStateManagerを使用して状態を一元管理
   - プロバイダーの依存関係を整理

2. **テストの充実**
   - ウィジェットテスト
   - 統合テスト
   - エンドツーエンドテスト

3. **CI/CDパイプライン**
   - 自動テスト実行
   - 自動デプロイ
   - 回帰検出アラート

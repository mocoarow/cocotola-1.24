#!/bin/bash

# 回帰テスト自動実行スクリプト

echo "🚀 回帰テスト開始..."

# 1. コード分析
echo "📝 コード分析実行中..."
flutter analyze
if [ $? -ne 0 ]; then
    echo "❌ コード分析でエラーが見つかりました"
    exit 1
fi

# 2. ユニットテスト実行
echo "🧪 ユニットテスト実行中..."
flutter test
if [ $? -ne 0 ]; then
    echo "❌ ユニットテストが失敗しました"
    exit 1
fi

# 3. 統合テスト実行
echo "🔗 統合テスト実行中..."
flutter test test/integration/
if [ $? -ne 0 ]; then
    echo "❌ 統合テストが失敗しました"
    exit 1
fi

# 4. ビルドテスト
echo "🏗️  ビルドテスト実行中..."
flutter build apk --debug
if [ $? -ne 0 ]; then
    echo "❌ ビルドが失敗しました"
    exit 1
fi

echo "✅ 全ての回帰テストが成功しました！"
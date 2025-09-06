import 'package:flutter/material.dart';

/// 軽量Markdownパーサー - 基本的な記法をサポート
class MarkdownTextParser {
  /// Markdownテキストを解析してTextSpanに変換
  static TextSpan parseMarkdown(String text, TextStyle? baseStyle) {
    if (text.isEmpty) {
      return TextSpan(text: text, style: baseStyle);
    }

    final List<TextSpan> spans = [];
    final RegExp markdownPattern = RegExp(
      r'(\*\*([^*]+)\*\*)|(\*([^*]+)\*)|(`([^`]+)`)|(\n)',
      multiLine: true,
    );

    int lastMatchEnd = 0;
    
    for (final match in markdownPattern.allMatches(text)) {
      // マッチ前のテキストを追加
      if (match.start > lastMatchEnd) {
        final beforeText = text.substring(lastMatchEnd, match.start);
        spans.add(TextSpan(text: beforeText, style: baseStyle));
      }

      // マッチしたMarkdown記法を処理
      if (match.group(1) != null) {
        // **太字**
        final boldText = match.group(2)!;
        spans.add(TextSpan(
          text: boldText,
          style: baseStyle?.copyWith(fontWeight: FontWeight.bold) ??
                 const TextStyle(fontWeight: FontWeight.bold),
        ));
      } else if (match.group(3) != null) {
        // *斜体*
        final italicText = match.group(4)!;
        spans.add(TextSpan(
          text: italicText,
          style: baseStyle?.copyWith(fontStyle: FontStyle.italic) ??
                 const TextStyle(fontStyle: FontStyle.italic),
        ));
      } else if (match.group(5) != null) {
        // `コード`
        final codeText = match.group(6)!;
        spans.add(TextSpan(
          text: codeText,
          style: baseStyle?.copyWith(
            fontFamily: 'Courier',
            backgroundColor: Colors.grey.shade200,
            fontWeight: FontWeight.w500,
          ) ?? const TextStyle(
            fontFamily: 'Courier',
            backgroundColor: Colors.grey,
            fontWeight: FontWeight.w500,
          ),
        ));
      } else if (match.group(7) != null) {
        // 改行
        spans.add(const TextSpan(text: '\n'));
      }

      lastMatchEnd = match.end;
    }

    // 最後の部分を追加
    if (lastMatchEnd < text.length) {
      final remainingText = text.substring(lastMatchEnd);
      spans.add(TextSpan(text: remainingText, style: baseStyle));
    }

    // spansが空の場合は元のテキストをそのまま返す
    if (spans.isEmpty) {
      return TextSpan(text: text, style: baseStyle);
    }

    return TextSpan(children: spans);
  }

  /// プレーンテキストからMarkdown記法を除去
  static String stripMarkdown(String text) {
    return text
        .replaceAll(RegExp(r'\*\*([^*]+)\*\*'), r'$1') // **太字** → 太字
        .replaceAll(RegExp(r'\*([^*]+)\*'), r'$1')     // *斜体* → 斜体
        .replaceAll(RegExp(r'`([^`]+)`'), r'$1');      // `コード` → コード
  }
}
import 'package:flutter/material.dart';

/// CEFRレベルを表示するバッジウィジェット
class CefrLevelBadge extends StatelessWidget {
  final String cefrLevel;
  final double? fontSize;
  final EdgeInsetsGeometry? padding;
  final bool showCefrPrefix;

  const CefrLevelBadge({
    super.key,
    required this.cefrLevel,
    this.fontSize = 12,
    this.padding,
    this.showCefrPrefix = true,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: padding ?? const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
      decoration: BoxDecoration(
        color: _getCefrLevelColor(cefrLevel),
        borderRadius: BorderRadius.circular(12),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withValues(alpha: 0.1),
            blurRadius: 2,
            offset: const Offset(0, 1),
          ),
        ],
      ),
      child: Text(
        showCefrPrefix ? 'CEFR $cefrLevel' : cefrLevel,
        style: TextStyle(
          color: Colors.white,
          fontSize: fontSize,
          fontWeight: FontWeight.bold,
        ),
      ),
    );
  }

  /// CEFRレベルに応じた色を返す
  Color _getCefrLevelColor(String cefrLevel) {
    switch (cefrLevel.toUpperCase()) {
      case 'A1':
        return Colors.green.shade600; // 初級（易しい）
      case 'A2':
        return Colors.lightGreen.shade600; // 初級上位
      case 'B1':
        return Colors.orange.shade600; // 中級
      case 'B2':
        return Colors.deepOrange.shade600; // 中級上位
      case 'C1':
        return Colors.red.shade600; // 上級
      case 'C2':
        return Colors.purple.shade600; // 最上級
      default:
        return Colors.grey.shade600; // その他
    }
  }

  /// CEFRレベルの説明を返す
  static String getCefrLevelDescription(String cefrLevel) {
    switch (cefrLevel.toUpperCase()) {
      case 'A1':
        return '初級 - 基本的な日常表現を理解し使用できる';
      case 'A2':
        return '初級上位 - 身近な話題について簡単なやりとりができる';
      case 'B1':
        return '中級 - 日常生活や仕事の基本的な内容を理解できる';
      case 'B2':
        return '中級上位 - 複雑な文章の要点を理解し、議論に参加できる';
      case 'C1':
        return '上級 - 複雑で長い文章を理解し、流暢に表現できる';
      case 'C2':
        return '最上級 - ほぼ全てを容易に理解し、自然に表現できる';
      default:
        return '不明なレベル';
    }
  }

  /// 大きいサイズのバッジを作成するファクトリーメソッド
  factory CefrLevelBadge.large({
    required String cefrLevel,
    bool showCefrPrefix = true,
  }) {
    return CefrLevelBadge(
      cefrLevel: cefrLevel,
      fontSize: 16,
      padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 6),
      showCefrPrefix: showCefrPrefix,
    );
  }

  /// 小さいサイズのバッジを作成するファクトリーメソッド
  factory CefrLevelBadge.small({
    required String cefrLevel,
    bool showCefrPrefix = false,
  }) {
    return CefrLevelBadge(
      cefrLevel: cefrLevel,
      fontSize: 10,
      padding: const EdgeInsets.symmetric(horizontal: 6, vertical: 2),
      showCefrPrefix: showCefrPrefix,
    );
  }
}
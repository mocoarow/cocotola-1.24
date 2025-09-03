import 'package:flutter/material.dart';

class CustomKeyboard extends StatelessWidget {
  final Function(String) onKeyPressed;
  final VoidCallback onDelete;
  final VoidCallback onMoveLeft;
  final VoidCallback onMoveRight;
  final String currentText;
  final int cursorPosition;
  final bool showCursorIndicator;

  const CustomKeyboard({
    super.key,
    required this.onKeyPressed,
    required this.onDelete,
    required this.onMoveLeft,
    required this.onMoveRight,
    this.currentText = '',
    this.cursorPosition = 0,
    this.showCursorIndicator = true,
  });

  @override
  Widget build(BuildContext context) {
    final screenWidth = MediaQuery.of(context).size.width;
    final screenHeight = MediaQuery.of(context).size.height;
    
    // 画面サイズに応じてボタンサイズを計算
    final keyWidth = (screenWidth - 80) / 10; // 余白を考慮して10個のキーに分割
    final keyHeight = screenHeight * 0.05; // 画面高さの5%
    final iconSize = keyHeight * 0.6; // ボタン高さの60%
    
    return Container(
      key: const Key('custom-keyboard'),
      padding: const EdgeInsets.all(8.0),
      child: Column(
        children: [
          // カーソル位置インジケーター（条件付き表示）
          if (showCursorIndicator) ...[
            _buildCursorIndicator(context),
            const SizedBox(height: 4),
          ],
          _buildRow(['Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P'], keyWidth, keyHeight),
          _buildRow(['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L'], keyWidth, keyHeight),
          _buildRow(['Z', 'X', 'C', 'V', 'B', 'N', 'M'], keyWidth, keyHeight),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              IconButton(
                icon: Icon(Icons.arrow_back, size: iconSize),
                onPressed: onMoveLeft,
                constraints: BoxConstraints(
                  minWidth: keyWidth,
                  minHeight: keyHeight,
                ),
              ),
              IconButton(
                icon: Icon(Icons.backspace, size: iconSize),
                onPressed: onDelete,
                constraints: BoxConstraints(
                  minWidth: keyWidth,
                  minHeight: keyHeight,
                ),
              ),
              IconButton(
                icon: Icon(Icons.arrow_forward, size: iconSize),
                onPressed: onMoveRight,
                constraints: BoxConstraints(
                  minWidth: keyWidth,
                  minHeight: keyHeight,
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildRow(List<String> keys, double keyWidth, double keyHeight) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: keys.map((key) => _buildKey(key, keyWidth, keyHeight)).toList(),
    );
  }

  Widget _buildKey(String key, double keyWidth, double keyHeight) {
    return Padding(
      padding: const EdgeInsets.all(2.0), // パディングを小さく
      child: SizedBox(
        width: keyWidth,
        height: keyHeight,
        child: ElevatedButton(
          onPressed: () => onKeyPressed(key),
          style: ElevatedButton.styleFrom(
            padding: EdgeInsets.zero, // 内部パディングを削除
            textStyle: TextStyle(
              fontSize: keyHeight * 0.4, // ボタン高さに応じてフォントサイズを調整
            ),
          ),
          child: Text(key),
        ),
      ),
    );
  }

  /// カーソル位置インジケーターを構築
  Widget _buildCursorIndicator(BuildContext context) {
    return Container(
      height: 32,
      decoration: BoxDecoration(
        color: Colors.grey[100],
        borderRadius: BorderRadius.circular(8),
        border: Border.all(color: Colors.grey[300]!),
      ),
      padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
      child: Row(
        children: [
          const Icon(Icons.edit, size: 14, color: Colors.grey),
          const SizedBox(width: 6),
          Expanded(
            child: _buildTextWithCursor(),
          ),
          Text(
            '$cursorPosition/${currentText.length}',
            style: TextStyle(
              fontSize: 10,
              color: Colors.grey[600],
            ),
          ),
        ],
      ),
    );
  }

  /// テキストとカーソルを表示するウィジェット
  Widget _buildTextWithCursor() {
    if (currentText.isEmpty) {
      return Row(
        children: [
          Container(
            width: 2,
            height: 16,
            decoration: BoxDecoration(
              color: Colors.blue,
              borderRadius: BorderRadius.circular(1),
            ),
          ),
          const Text('_', style: TextStyle(color: Colors.transparent)),
        ],
      );
    }

    final beforeCursor = currentText.substring(0, cursorPosition);
    final afterCursor = currentText.substring(cursorPosition);

    return Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        if (beforeCursor.isNotEmpty)
          Text(
            beforeCursor,
            style: const TextStyle(fontSize: 14),
          ),
        Container(
          width: 2,
          height: 16,
          decoration: BoxDecoration(
            color: Colors.blue,
            borderRadius: BorderRadius.circular(1),
          ),
        ),
        if (afterCursor.isNotEmpty)
          Text(
            afterCursor,
            style: const TextStyle(fontSize: 14),
          ),
      ],
    );
  }
}

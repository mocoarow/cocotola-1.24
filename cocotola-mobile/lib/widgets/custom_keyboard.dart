import 'package:flutter/material.dart';

class CustomKeyboard extends StatelessWidget {
  final Function(String) onKeyPressed;
  final VoidCallback onDelete;
  final VoidCallback onMoveLeft;
  final VoidCallback onMoveRight;

  const CustomKeyboard({
    super.key,
    required this.onKeyPressed,
    required this.onDelete,
    required this.onMoveLeft,
    required this.onMoveRight,
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
      padding: const EdgeInsets.all(8.0),
      child: Column(
        children: [
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
}

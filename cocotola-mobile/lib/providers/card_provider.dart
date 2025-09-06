import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../models/problem_base.dart';
import '../services/card_api_service.dart';

/// CardApiServiceのプロバイダー
final cardApiServiceProvider = Provider<CardApiService>((ref) {
  return CardApiService();
});

/// カード問題一覧を取得するプロバイダー
final cardProblemsProvider = FutureProvider<List<Problem>>((ref) async {
  final service = ref.read(cardApiServiceProvider);
  return service.fetchCards();
});

/// カード問題一覧の状態管理プロバイダー
final cardProblemsNotifierProvider = NotifierProvider<CardProblemsNotifier, AsyncValue<List<Problem>>>(
  CardProblemsNotifier.new,
);

class CardProblemsNotifier extends Notifier<AsyncValue<List<Problem>>> {
  @override
  AsyncValue<List<Problem>> build() {
    return const AsyncValue.loading();
  }

  /// カード問題を取得
  Future<void> fetchCards() async {
    state = const AsyncValue.loading();
    try {
      final service = ref.read(cardApiServiceProvider);
      final problems = await service.fetchCards();
      state = AsyncValue.data(problems);
    } catch (error, stackTrace) {
      state = AsyncValue.error(error, stackTrace);
    }
  }

  /// カード問題をリフレッシュ
  Future<void> refresh() async {
    await fetchCards();
  }
}
import 'problem_base.dart';

class ProblemSet {
  final String id;
  final String title;
  final String description;
  final String cefrLevel;
  final List<Problem> problems;
  final String iconPath;
  
  const ProblemSet({
    required this.id,
    required this.title,
    required this.description,
    required this.cefrLevel,
    required this.problems,
    this.iconPath = '',
  });
  
  int get problemCount => problems.length;
  
  ProblemSet copyWith({
    String? id,
    String? title,
    String? description,
    String? cefrLevel,
    List<Problem>? problems,
    String? iconPath,
  }) {
    return ProblemSet(
      id: id ?? this.id,
      title: title ?? this.title,
      description: description ?? this.description,
      cefrLevel: cefrLevel ?? this.cefrLevel,
      problems: problems ?? this.problems,
      iconPath: iconPath ?? this.iconPath,
    );
  }
}
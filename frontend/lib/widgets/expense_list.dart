import 'package:flutter/material.dart';
import '../models/expense.dart';
import 'expense_card.dart';

class ExpenseList extends StatelessWidget {
  final List<Expense> expenses;
  final Function(int) onDelete;

  ExpenseList({required this.expenses, required this.onDelete});

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemCount: expenses.length,
      itemBuilder: (context, index) {
        return ExpenseCard(
          expense: expenses[index],
          onDelete: () => onDelete(expenses[index].id),
        );
      },
    );
  }
}

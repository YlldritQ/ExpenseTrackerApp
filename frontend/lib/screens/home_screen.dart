import 'package:flutter/material.dart';
import '../models/expense.dart';
import '../../services/api_service.dart';
import '../../widgets/expense_list.dart';

class HomeScreen extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  final ApiService apiService = ApiService();
  List<Expense> expenses = [];
  bool loading = true;

  @override
  void initState() {
    super.initState();
    fetchExpenses();
  }

  Future<void> fetchExpenses() async {
    setState(() => loading = true);
    try {
      expenses = await apiService.getExpenses();
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(e.toString())));
    } finally {
      setState(() => loading = false);
    }
  }

  Future<void> deleteExpense(int id) async {
    await apiService.deleteExpense(id);
    fetchExpenses();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Expenses')),
      body: loading
          ? Center(child: CircularProgressIndicator())
          : ExpenseList(expenses: expenses, onDelete: deleteExpense),
    );
  }
}

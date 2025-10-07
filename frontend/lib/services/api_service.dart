import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/expense.dart';

class ApiService {
  static const baseUrl = 'http://localhost:8080/api';

  Future<List<Expense>> getExpenses() async {
    final response = await http.get(Uri.parse('$baseUrl/expenses'));
    if (response.statusCode == 200) {
      List data = json.decode(response.body);
      return data.map((e) => Expense.fromJson(e)).toList();
    } else {
      throw Exception('Failed to load expenses');
    }
  }

  Future<void> addExpense(Map<String, dynamic> expense) async {
    final response = await http.post(
      Uri.parse('$baseUrl/expenses'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode(expense),
    );
    if (response.statusCode != 200) {
      throw Exception('Failed to add expense');
    }
  }

  Future<void> deleteExpense(int id) async {
    final response = await http.delete(Uri.parse('$baseUrl/expenses/$id'));
    if (response.statusCode != 200) {
      throw Exception('Failed to delete expense');
    }
  }
}
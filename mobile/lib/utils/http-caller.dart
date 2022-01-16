import 'dart:convert' as convert;
import 'dart:io';

import 'package:http/http.dart' as http;

class HttpCaller {
  final String defaultErrorMessage = """
  Произошла ошибка вызова сервиса. Сервис временно недоступен. 
  Попробуйте выполнить данное действие через некоторое время.
  """;

  static T handleResponse<T>(http.Response response) {
    T result;
    if (response.statusCode == HttpStatus.ok) {
      if (response.body != null && response.body.isNotEmpty) {
        Map<String, dynamic> resultMap = convert.jsonDecode(response.body) as Map<String, dynamic>;
        if (resultMap != null) {
          result = resultMap['data'] as T;
        }
      }
      return '' as T;
    }
    return result;
  }

  static Future<T> post<T>(Uri uri, {Map<String, String> headers, dynamic params}) async {
    http.Response response = await http.post(uri, headers: headers, body: params);
    return Future.value(handleResponse(response));
  }

  static Future<T> get<T>(Uri uri, Map<String, String> headers, T params) async {
    http.Response response = await http.get(uri, headers: headers);
    return Future.value(handleResponse(response));
  }
}

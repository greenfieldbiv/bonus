import 'dart:convert' as convert;
import 'dart:developer';
import 'dart:io';

import 'package:bivbonus/utils/secure-storage.dart';
import 'package:http/http.dart' as http;

class HttpCaller {
  final SecureStorage _secureStorage = new SecureStorage();

  final String defaultErrorMessage = """
  Произошла ошибка вызова сервиса. Сервис временно недоступен. 
  Попробуйте выполнить данное действие через некоторое время.
  """;

  T handleResponse<T>(http.Response response) {
    T result;
    switch (response.statusCode) {
      case HttpStatus.ok:
        {
          _storageToken(response);
          if (response.body != null && response.body.isNotEmpty) {
            Map<String, dynamic> resultMap = convert.jsonDecode(response.body);
            if (resultMap != null) {
              return resultMap['data'] as T;
            }
          }
          return '' as T;
        }
      case HttpStatus.unauthorized:
      case HttpStatus.forbidden:
        {
          log("unauthorized || forbidden");
        }
    }
    return result;
  }

  void _storageToken(http.Response response) {
    String aToken = response.headers["a_tkn"];
    if (aToken != null) {
      _secureStorage.writeSecureData(_secureStorage.ACCESS_TOKEN, aToken);
    }
  }

  Future<T> post<T>(Uri uri, {Map<String, String> headers, dynamic params}) async {
    headers = await _interceptHttp(headers);
    http.Response response = await http.post(uri, headers: headers, body: params);
    return Future.value(handleResponse(response));
  }

  Future<T> get<T>(Uri uri) async {
    http.Response response = await http.get(uri, headers: await _interceptHttp({}));
    return Future.value(handleResponse(response));
  }

  Future<Map<String, String>> _interceptHttp(Map<String, String> headers) async {
    if (headers == null) {
      headers = {};
    }
    String authHeader = HttpHeaders.authorizationHeader[0].toUpperCase() + HttpHeaders.authorizationHeader.substring(1).toLowerCase();
    var isContainKey = await _secureStorage.isContainKey(_secureStorage.ACCESS_TOKEN);
    if (isContainKey) {
      String token = await _secureStorage.readSecureData(_secureStorage.ACCESS_TOKEN);
      headers[authHeader] = "Bearer ${token}";
    }
    return headers;
  }
}

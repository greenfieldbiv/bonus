import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class SecureStorage {
  final FlutterSecureStorage _storage = new FlutterSecureStorage();

  setDataByKey(String key, final String value) {
    if (key != null && value != null) {
      _storage.write(key: key, value: value);
    }
  }

  String getDataByKey(String key) {
    if (key != null) {
      return _storage.read(key: key) as String;
    }
    return "";
  }
}

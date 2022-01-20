import 'package:shared_preferences/shared_preferences.dart';

class SecureStorage {
  final String ACCESS_TOKEN = 'access_token';

  Future<void> writeSecureData(String key, String value) async {
    SharedPreferences storage = await SharedPreferences.getInstance();
    storage.setString(key, value ?? "");
  }

  Future<String> readSecureData(String key) async {
    SharedPreferences storage = await SharedPreferences.getInstance();
    return storage.getString(key);
  }

  Future<bool> isContainKey(String key) async {
    SharedPreferences storage = await SharedPreferences.getInstance();
    return storage.containsKey(key);
  }
}

import 'dart:convert';

import 'package:bivbonus/pages/auth/login.dart';
import 'package:bivbonus/utils/http-caller.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class Register extends StatelessWidget {
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  final httpCaller = HttpCaller();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: false,
      body: Center(
        child: Container(
          // constraints: BoxConstraints.expand(),
          child: Align(
            alignment: Alignment.topCenter,
            child: Column(
              children: [
                Padding(
                  padding: const EdgeInsets.only(top: 150),
                  // child: Image.asset("images/logo.png"),
                ),
                SizedBox(
                  height: 100,
                ),
                Padding(
                  padding: const EdgeInsets.only(left: 30.0, right: 30, bottom: 30),
                  child: Container(
                    child: TextField(
                      textAlign: TextAlign.center,
                      minLines: 1,
                      maxLines: 1,
                      autocorrect: false,
                      controller: _emailController,
                      decoration: InputDecoration(
                        icon: Icon(Icons.mail_outline),
                        hintText: 'Email',
                        filled: true,
                        fillColor: Colors.white,
                        enabledBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.all(Radius.circular(40.0)),
                          borderSide: BorderSide(color: Colors.white),
                        ),
                        focusedBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.all(Radius.circular(40.0)),
                          borderSide: BorderSide(color: Colors.grey),
                        ),
                      ),
                    ),
                  ),
                ),
                Padding(
                  padding: EdgeInsets.only(left: 30.0, right: 30, bottom: 30),
                  child: Container(
                    child: TextField(
                      obscureText: true,
                      textAlign: TextAlign.center,
                      minLines: 1,
                      maxLines: 1,
                      autocorrect: false,
                      controller: _passwordController,
                      decoration: InputDecoration(
                        icon: Icon(Icons.mail_outline),
                        hintStyle: GoogleFonts.cinzel(color: Colors.grey[500], fontSize: 16, fontWeight: FontWeight.normal),
                        hintText: 'Пароль',
                        filled: true,
                        fillColor: Colors.white,
                        enabledBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.all(Radius.circular(40.0)),
                          borderSide: BorderSide(color: Colors.white),
                        ),
                        focusedBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.all(Radius.circular(40.0)),
                          borderSide: BorderSide(color: Colors.grey),
                        ),
                      ),
                    ),
                  ),
                ),
                Container(
                  width: 350,
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.stretch,
                    children: [
                      ElevatedButton(
                        onPressed: () {
                          httpCaller.post<dynamic>(new Uri.http('10.0.2.2:8080', 'auth/registry'),
                              params: json.encode({"email": _emailController.text, "password": _passwordController.text}),
                              headers: {"Content-Type": "application/json"}).then(
                            (value) => value != null
                                ? Navigator.push(
                                    context,
                                    MaterialPageRoute(
                                      builder: (context) => LoginWidget(),
                                    ),
                                  )
                                : ScaffoldMessenger.of(context).showSnackBar(SnackBar(
                                    content: Text("Не удалось авторизоваться"),
                                  )),
                          );
                        },
                        child: Text(
                          'Зарегистрироваться',
                          style: GoogleFonts.cinzel(color: Colors.white, fontSize: 16, fontWeight: FontWeight.bold),
                          textAlign: TextAlign.center,
                        ),
                      ),
                      SizedBox(
                        height: 25,
                      ),
                      ElevatedButton(
                        onPressed: () {
                          Navigator.pop(context);
                        },
                        child: Text(
                          'Назад',
                          style: GoogleFonts.cinzel(color: Colors.white, fontSize: 16, fontWeight: FontWeight.bold),
                          textAlign: TextAlign.center,
                        ),
                      ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}

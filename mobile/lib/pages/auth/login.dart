import 'dart:convert';

import 'package:bivbonus/mainpage.dart';
import 'package:bivbonus/pages/auth/register.dart';
import 'package:bivbonus/utils/http-caller.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class LoginWidget extends StatelessWidget {
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();

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
                  child: Column(
                    children: [
                      InkWell(
                        onTap: () {
                          HttpCaller.post<dynamic>(new Uri.http('10.0.2.2:8080', 'auth/login'),
                              params: json.encode({"email": _emailController.text, "password": _passwordController.text}),
                              headers: {"Content-Type": "application/json"}).then(
                            (value) => value != null
                                ? Navigator.push(
                                    context,
                                    MaterialPageRoute(
                                      builder: (context) => MyNavigationBar(),
                                    ),
                                  )
                                : ScaffoldMessenger.of(context).showSnackBar(SnackBar(
                                    content: Text("Не удалось авторизоваться"),
                                  )),
                          );
                        },
                        child: ElevatedButton(
                          child: Text(
                            'Войти',
                            style: GoogleFonts.cinzel(color: Colors.white, fontSize: 16, fontWeight: FontWeight.bold),
                            textAlign: TextAlign.center,
                          ),
                        ),
                      ),
                      SizedBox(
                        height: 25,
                      ),
                      InkWell(
                        onTap: () {
                          Navigator.push(
                            context,
                            MaterialPageRoute(
                              builder: (context) => Register(),
                            ),
                          );
                        },
                        child: ElevatedButton(
                          child: Text(
                            'Зарегистрироваться',
                            style: GoogleFonts.cinzel(color: Colors.white, fontSize: 16, fontWeight: FontWeight.bold),
                            textAlign: TextAlign.center,
                          ),
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

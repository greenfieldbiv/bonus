import 'dart:convert';

import 'package:bivbonus/pages/auth/login.dart';
import 'package:bivbonus/utils/http-caller.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class Register extends StatelessWidget {
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: SingleChildScrollView(
          child: ConstrainedBox(
            constraints: BoxConstraints(
                //    minHeight: viewportConstraints.maxHeight,

                ),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.end,
              children: [
                Column(
                  children: [
                    Padding(
                      padding: const EdgeInsets.only(top: 150),
                      child: Image.asset("images/logo.png"),
                    ),
                    Padding(
                      padding: const EdgeInsets.only(left: 30.0, right: 30, bottom: 30),
                      child: Container(
                        child: TextField(
                          textAlign: TextAlign.center,
                          minLines: 1,
                          maxLines: 1,
                          autocorrect: false,
                          decoration: InputDecoration(
                            icon: Icon(Icons.mail_outline),
                            hintText: 'Введите Email',
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
                      padding: const EdgeInsets.only(left: 30.0, right: 30, bottom: 30),
                      child: Container(
                        decoration: BoxDecoration(
                          boxShadow: [
                            BoxShadow(
                              color: Colors.grey[200],
                              blurRadius: 6.0, // soften the shadow
                              spreadRadius: 2.0, //extend the shadow
                              offset: Offset(
                                0, // Move to right 10  horizontally
                                0, // Move to bottom 10 Vertically
                              ),
                            )
                          ],
                        ),
                        child: TextField(
                          textAlign: TextAlign.center,
                          minLines: 1,
                          maxLines: 1,
                          autocorrect: false,
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
                    Padding(
                      padding: const EdgeInsets.only(left: 30.0, right: 30, bottom: 30),
                      child: Container(
                        decoration: BoxDecoration(
                          boxShadow: [
                            BoxShadow(
                              blurRadius: 6.0, // soften the shadow
                              spreadRadius: 2.0, //extend the shadow
                              offset: Offset(
                                0, // Move to right 10  horizontally
                                0, // Move to bottom 10 Vertically
                              ),
                            )
                          ],
                        ),
                      ),
                    ),
                  ],
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Padding(
                      padding: const EdgeInsets.only(top: 28, left: 0.0),
                      child: Container(
                          height: 45,
                          width: 150,
                          decoration: BoxDecoration(
                            color: Colors.blueAccent[400],
                            borderRadius: new BorderRadius.only(
                              bottomLeft: const Radius.circular(5.0),
                              bottomRight: const Radius.circular(5.0),
                              topLeft: const Radius.circular(5.0),
                              topRight: const Radius.circular(5.0),
                            ),
                          ),
                          child: Padding(
                            padding: const EdgeInsets.only(
                              top: 0.0,
                            ),
                            child: Column(
                              children: [
                                InkWell(onTap: () {
                                  var t = HttpCaller.post<dynamic>(new Uri.http('10.0.2.2:8080', 'auth/registry'),
                                      params: json.encode({"email": _emailController.text, "password": _passwordController.text}),
                                      headers: {"Content-Type": "application/json"}).then(
                                    (value) => value != null
                                        ? Navigator.push(
                                            context,
                                            MaterialPageRoute(
                                              builder: (context) => LoginWidget(),
                                            ))
                                        : ScaffoldMessenger.of(context).showSnackBar(SnackBar(
                                            content: Text("Не удалось авторизоваться"),
                                          )),
                                  );
                                }),
                              ],
                            ),
                          )),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}

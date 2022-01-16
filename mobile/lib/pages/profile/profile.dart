import 'package:bivbonus/mainpage.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class Profile extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          // Add your onPressed code here!
        },
        child: Container(
          decoration: BoxDecoration(color: Colors.blue, borderRadius: BorderRadius.all(Radius.circular(40))),
          child: Padding(
            padding: const EdgeInsets.all(10.0),
            child: Icon(
              Icons.edit,
              color: Colors.white,
              size: 16,
            ),
          ),
        ),
      ),
      appBar: AppBar(
        elevation: 0,
        backgroundColor: Colors.white,
        leading: InkWell(
          onTap: () {
            Navigator.push(
              context,
              MaterialPageRoute(
                builder: (context) => MyNavigationBar(),
              ),
            );
          },
          child: Icon(
            Icons.arrow_back_ios_outlined,
            color: Colors.grey[700],
            size: 18,
          ),
        ),
      ),
      backgroundColor: Colors.white,
      body: SingleChildScrollView(
        child: Column(
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Padding(
                  padding: const EdgeInsets.only(left: 22.0),
                  child: Text(
                    'Профиль',
                    style: GoogleFonts.lato(color: Colors.grey[800], fontSize: 26, letterSpacing: 0, fontWeight: FontWeight.bold),
                  ),
                ),
              ],
            ),
            Center(
              child: Container(
                decoration: new BoxDecoration(color: Colors.white),
                height: 180,
                child: Stack(
                  children: <Widget>[
                    Container(
                        height: 108,
                        width: 101,
                        margin: const EdgeInsets.only(left: 15.0, right: 15, top: 25, bottom: 5),
                        padding: const EdgeInsets.all(2.0),
                        decoration: BoxDecoration(border: Border.all(color: Colors.white, width: 2), borderRadius: BorderRadius.circular(140)),
                        child: CircleAvatar(
                          backgroundImage: NetworkImage(
                            'https://cdn.now.howstuffworks.com/media-content/0b7f4e9b-f59c-4024-9f06-b3dc12850ab7-1920-1080.jpg',
                          ),
                        )),
                    Positioned(
                      bottom: 54,
                      right: 20,
                      //give the values according to your requirement
                      child: Material(
                          color: Colors.blue,
                          elevation: 10,
                          borderRadius: BorderRadius.circular(100),
                          child: Padding(
                            padding: const EdgeInsets.all(5.0),
                            child: Icon(
                              Icons.zoom_out_map,
                              size: 18,
                              color: Colors.white,
                            ),
                          )),
                    ),
                  ],
                ),
              ),
            ),
            Row(
              children: [
                Padding(
                  padding: const EdgeInsets.only(left: 24.0),
                  child: Text(
                    'Имя ',
                    style: GoogleFonts.lato(color: Colors.grey[900], fontSize: 16, letterSpacing: 0, fontWeight: FontWeight.bold),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.only(left: 37.0),
                  child: Text(
                    '   Барак Обама',
                    style: GoogleFonts.lato(color: Colors.grey[600], fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.normal),
                  ),
                ),
              ],
            ),
            SizedBox(
              height: 30,
            ),
            Row(
              children: [
                Padding(
                  padding: const EdgeInsets.only(left: 24.0),
                  child: Text(
                    'Роль ',
                    style: GoogleFonts.lato(color: Colors.grey[900], fontSize: 16, letterSpacing: 0, fontWeight: FontWeight.bold),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.only(left: 37.0),
                  child: Text(
                    '  Social Engineer Of Google',
                    style: GoogleFonts.lato(color: Colors.grey[600], fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.normal),
                  ),
                ),
              ],
            ),
            SizedBox(
              height: 30,
            ),
            Row(
              children: [
                Padding(
                  padding: const EdgeInsets.only(left: 24.0),
                  child: Text(
                    'Команда ',
                    style: GoogleFonts.lato(color: Colors.grey[900], fontSize: 16, letterSpacing: 0, fontWeight: FontWeight.bold),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.only(left: 14.0),
                  child: Text(
                    'Команда Обамы',
                    style: GoogleFonts.lato(color: Colors.grey[600], fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.normal),
                  ),
                ),
              ],
            ),
            SizedBox(
              height: 30,
            ),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 22.0),
              child: Divider(),
            ),
            SizedBox(
              height: 30,
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                Padding(
                  padding: const EdgeInsets.only(left: 54.0),
                  child: Icon(Icons.mail, color: Colors.grey[500]),
                ),
                Padding(
                  padding: const EdgeInsets.only(left: 8.0),
                  child: Text(
                    'mathewsteven92@gmail.com',
                    style: GoogleFonts.lato(color: Colors.grey[700], fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.normal),
                  ),
                ),
              ],
            ),
            SizedBox(
              height: 20,
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                Padding(
                  padding: const EdgeInsets.only(left: 54.0),
                  child: Icon(Icons.phone, color: Colors.grey[500]),
                ),
                Padding(
                  padding: const EdgeInsets.only(left: 8.0),
                  child: Text(
                    '+91 - 9560419114',
                    style: GoogleFonts.lato(color: Colors.grey[700], fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.normal),
                  ),
                ),
              ],
            ),
            SizedBox(
              height: 20,
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                Padding(
                  padding: const EdgeInsets.only(left: 54.0),
                  child: Icon(Icons.home_outlined, color: Colors.grey[500]),
                ),
                Padding(
                  padding: const EdgeInsets.only(left: 8.0),
                  child: Text(
                    'RZ- 5167, Hari Nagar, New Delhi',
                    style: GoogleFonts.lato(color: Colors.grey[700], fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.normal),
                  ),
                ),
              ],
            ),
            SizedBox(
              height: 20,
            ),
          ],
        ),
      ),
    );
  }
}

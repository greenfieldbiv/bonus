import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

import 'event.dart';

class AchievementPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: SingleChildScrollView(
        child: Column(
          children: [
            SizedBox(
              height: 40,
            ),
            Center(
              child: Text(
                'Достижения',
                style: GoogleFonts.lato(color: Colors.black87, fontSize: 16.5, letterSpacing: 1, fontWeight: FontWeight.bold),
              ),
            ),
            SizedBox(
              height: 0,
            ),
            Padding(
              padding: const EdgeInsets.only(left: 18.0, right: 18, top: 14),
              child: Material(
                borderRadius: BorderRadius.circular(18),
                elevation: 4,
                child: Column(
                  children: [
                    Container(
                      decoration: BoxDecoration(color: Colors.white),
                      height: 450,
                      child: Stack(
                        children: <Widget>[
                          Container(
                              height: 212,
                              width: MediaQuery.of(context).size.width,
                              child: ClipRRect(
                                  borderRadius: BorderRadius.only(topLeft: Radius.circular(10.0), topRight: Radius.circular(10.0)),
                                  child: Image.network(
                                    'https://cdn.now.howstuffworks.com/media-content/0b7f4e9b-f59c-4024-9f06-b3dc12850ab7-1920-1080.jpg',
                                  ))),
                          Positioned(
                            bottom: 210,
                            right: 250, //give the values according to your requirement
                            child: Container(
                              height: 80,
                              width: 80,
                              child: Material(
                                  color: Colors.white,
                                  elevation: 10,
                                  borderRadius: BorderRadius.circular(100),
                                  child: Padding(
                                    padding: const EdgeInsets.all(3.0),
                                    child: CircleAvatar(
                                      backgroundImage: NetworkImage(
                                        'https://images.pexels.com/photos/3767673/pexels-photo-3767673.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940',
                                      ),
                                    ),
                                  )),
                            ),
                          ),
                          Positioned(
                              top: 260,
                              left: 25,
                              child: Text(
                                'Wed 2, October 2019, at 9AM',
                                style: GoogleFonts.lato(color: Colors.blue[700], fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.normal),
                              )),
                          Positioned(
                              top: 283,
                              left: 25,
                              child: Text(
                                'Atlassian Open 2019',
                                style: GoogleFonts.lato(color: Colors.black87, fontSize: 16.5, letterSpacing: 1, fontWeight: FontWeight.bold),
                              )),
                          Positioned(
                              top: 303,
                              left: 26,
                              child: Text(
                                'at Sudney, Australia',
                                style: GoogleFonts.lato(color: Colors.grey[500], fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.normal),
                              )),
                          Positioned(
                            top: 393,
                            child: Padding(
                              padding: const EdgeInsets.all(8.0),
                              child: InkWell(
                                onTap: () {
                                  var route = new MaterialPageRoute(
                                    builder: (BuildContext context) => new EventPage(),
                                  );

                                  Navigator.of(context).push(route);
                                },
                                child: Material(
                                  color: Colors.blue,
                                  borderRadius: BorderRadius.all(Radius.circular(5)),
                                  elevation: 2,
                                  child: Container(
                                      height: 40,
                                      width: MediaQuery.of(context).size.width / 1.15,
                                      decoration: BoxDecoration(
                                        borderRadius: BorderRadius.all(Radius.circular(5)),
                                        border: Border.all(color: Colors.blue, width: 1),
                                      ),
                                      child: Center(
                                          child: Text(
                                        'Хочу учавствовать в событии',
                                        style:
                                            GoogleFonts.averageSans(color: Colors.white, fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.bold),
                                        textAlign: TextAlign.center,
                                      ))),
                                ),
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
            SizedBox(
              height: 40,
            ),
            Padding(
              padding: const EdgeInsets.only(left: 18.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.start,
                children: [
                  Padding(
                    padding: const EdgeInsets.only(top: 1.0),
                    child: Icon(
                      Icons.event,
                      size: 19,
                      color: Colors.blue,
                    ),
                  ),
                  SizedBox(
                    width: 5,
                  ),
                  Text(
                    'Последние полученные достижения',
                    style: GoogleFonts.averageSans(color: Colors.grey[800], fontSize: 17, letterSpacing: 1, fontWeight: FontWeight.normal),
                  ),
                ],
              ),
            ),
            Container(
              height: 250,
              child: ListView(scrollDirection: Axis.horizontal, children: [
                Padding(
                  padding: const EdgeInsets.only(left: 18.0, right: 28, top: 14, bottom: 5),
                  child: Material(
                    borderRadius: BorderRadius.only(topLeft: Radius.circular(10.0), topRight: Radius.circular(10.0)),
                    elevation: 2,
                    child: Column(
                      children: [
                        Container(
                          decoration: BoxDecoration(color: Colors.white),
                          height: 220,
                          width: 280,
                          child: Stack(
                            children: <Widget>[
                              Container(
                                  height: 100,
                                  width: MediaQuery.of(context).size.width,
                                  child: ClipRRect(
                                    borderRadius: BorderRadius.only(topLeft: Radius.circular(10.0), topRight: Radius.circular(10.0)),
                                    child: Image.network(
                                        'https://images.pexels.com/photos/1470165/pexels-photo-1470165.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940',
                                        fit: BoxFit.cover),
                                  )),
                              Positioned(
                                bottom: 100,
                                right: 210, //give the values according to your requirement
                                child: Material(
                                    color: Colors.white,
                                    elevation: 10,
                                    borderRadius: BorderRadius.circular(100),
                                    child: Padding(
                                      padding: const EdgeInsets.all(3.0),
                                      child: CircleAvatar(
                                        backgroundImage: NetworkImage(
                                          'https://images.pexels.com/photos/3767673/pexels-photo-3767673.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940',
                                        ),
                                      ),
                                    )),
                              ),
                              Positioned(
                                  top: 130,
                                  left: 25,
                                  child: Text(
                                    'Wed 2, October 2019, at 9AM',
                                    style: GoogleFonts.lato(color: Colors.blue[700], fontSize: 13, letterSpacing: 1, fontWeight: FontWeight.normal),
                                  )),
                              Positioned(
                                  top: 150,
                                  left: 25,
                                  child: Text(
                                    'Atlassian Open 2019',
                                    style: GoogleFonts.lato(color: Colors.black87, fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.bold),
                                  )),
                              Positioned(
                                top: 160,
                                child: Padding(
                                  padding: const EdgeInsets.only(top: 25, left: 26.0, bottom: 5),
                                ),
                              ),
                            ],
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.only(left: 8.0, right: 28, top: 14, bottom: 5),
                  child: Material(
                    borderRadius: BorderRadius.only(topLeft: Radius.circular(10.0), topRight: Radius.circular(10.0)),
                    elevation: 2,
                    child: Column(
                      children: [
                        Container(
                          decoration: BoxDecoration(color: Colors.white),
                          height: 220,
                          width: 280,
                          child: Stack(
                            children: <Widget>[
                              Container(
                                  height: 100,
                                  width: MediaQuery.of(context).size.width,
                                  child: ClipRRect(
                                    borderRadius: BorderRadius.only(topLeft: Radius.circular(10.0), topRight: Radius.circular(10.0)),
                                    child: Image.network(
                                        'https://images.pexels.com/photos/569996/pexels-photo-569996.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940',
                                        fit: BoxFit.cover),
                                  )),
                              Positioned(
                                bottom: 100,
                                right: 210, //give the values according to your requirement
                                child: Material(
                                    color: Colors.white,
                                    elevation: 10,
                                    borderRadius: BorderRadius.circular(100),
                                    child: Padding(
                                      padding: const EdgeInsets.all(3.0),
                                      child: CircleAvatar(
                                        backgroundImage: NetworkImage(
                                          'https://images.pexels.com/photos/6186/vintage-mockup-old-logo.jpg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940',
                                        ),
                                      ),
                                    )),
                              ),
                              Positioned(
                                  top: 130,
                                  left: 25,
                                  child: Text(
                                    'Wed 2, October 2019, at 9AM',
                                    style: GoogleFonts.lato(color: Colors.blue[700], fontSize: 13, letterSpacing: 1, fontWeight: FontWeight.normal),
                                  )),
                              Positioned(
                                  top: 150,
                                  left: 25,
                                  child: Text(
                                    'Atlassian Open 2019',
                                    style: GoogleFonts.lato(color: Colors.black87, fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.bold),
                                  )),
                            ],
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ]),
            ),
            SizedBox(
              height: 0,
            )
          ],
        ),
      ),
    );
  }
}

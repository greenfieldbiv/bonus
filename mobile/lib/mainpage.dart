import 'package:bivbonus/pages/common/achievement-page.dart';
import 'package:bivbonus/pages/common/connections.dart';
import 'package:bivbonus/pages/common/event.dart';
import 'package:bivbonus/pages/common/notification.dart';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:google_fonts/google_fonts.dart';

import 'pages/profile/profile.dart';

class MainPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: SingleChildScrollView(
        child: Column(
          children: [
            SizedBox(
              height: 15,
            ),
            Padding(
              padding: const EdgeInsets.only(top: 30.0, left: 30.0, right: 30.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(
                    'Bonus',
                    style: GoogleFonts.lato(
                      color: Colors.black,
                      fontSize: 16,
                      letterSpacing: 1,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                  InkWell(
                    onTap: () => {
                      Navigator.push(
                        context,
                        MaterialPageRoute(
                          builder: (context) => NotificationPage(),
                        ),
                      )
                    },
                    child: Padding(
                      padding: const EdgeInsets.all(0.0),
                      child: Icon(
                        Icons.notifications_active_outlined,
                        color: Colors.grey[700],
                      ),
                    ),
                  )
                ],
              ),
            ),
            Padding(
              padding: const EdgeInsets.only(top: 30.0),
              child: Text(
                'Добро пожаловать!',
                textAlign: TextAlign.start,
              ),
            ),
            SizedBox(
              height: 10,
            ),
            Card(
              color: Colors.blue,
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: <Widget>[
                  const ListTile(
                    title: Text(
                      'Участники команды онлайн:',
                      style: TextStyle(
                        color: Colors.white,
                      ),
                    ),
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: <Widget>[
                      Row(
                        children: [
                          Padding(
                            padding: const EdgeInsets.only(
                              left: 15.0,
                              top: 0.0,
                            ),
                            child: Material(
                              elevation: 4,
                              borderRadius: BorderRadius.circular(100),
                              child: Padding(
                                padding: const EdgeInsets.all(2.0),
                                child: CircleAvatar(
                                    radius: 9,
                                    backgroundImage: NetworkImage(
                                      'https://i.insider.com/5c9a115d8e436a63e42c2883?width=600&format=jpeg&auto=webp',
                                    )),
                              ),
                            ),
                          ),
                          Padding(
                            padding: const EdgeInsets.all(0.0),
                            child: Material(
                              elevation: 4,
                              borderRadius: BorderRadius.circular(100),
                              child: Padding(
                                padding: const EdgeInsets.all(2.0),
                                child: CircleAvatar(
                                    radius: 9,
                                    backgroundImage: NetworkImage(
                                      'https://play-images-prod-cms.tech.tvnz.co.nz/api/v1/web/image/content/dam/images/entertainment/shows/p/person-of-interest/personofinterest_coverimg.png.2017-03-08T11:21:33+13:00.jpg?width=960&height=540',
                                    )),
                              ),
                            ),
                          ),
                          Padding(
                            padding: const EdgeInsets.only(top: 4.0, left: 10),
                            child: Text(
                              '5 в сети, включая Вас',
                              style: GoogleFonts.lato(color: Colors.white, fontSize: 14, letterSpacing: 1, fontWeight: FontWeight.bold),
                            ),
                          ),
                        ],
                      ),
                      Padding(
                        padding: const EdgeInsets.only(top: 4.0, left: 10),
                        child: TextButton(
                          child: const Text(
                            'все',
                            style: TextStyle(
                              color: Colors.white,
                            ),
                          ),
                          onPressed: () {
                            Navigator.push(
                              context,
                              MaterialPageRoute(
                                builder: (context) => Connections(),
                              ),
                            );
                          },
                        ),
                      ),
                    ],
                  ),
                  SizedBox(
                    height: 10,
                  ),
                ],
              ),
            ),
            Padding(
              padding: const EdgeInsets.only(left: 20, top: 15),
              child: Text('Категории'),
            ),
            SizedBox(
              height: 10,
            ),
            CustomScrollView(
              scrollDirection: Axis.vertical,
              shrinkWrap: true,
              primary: false,
              slivers: <Widget>[
                SliverPadding(
                  padding: const EdgeInsets.all(20),
                  sliver: SliverGrid.count(
                    crossAxisSpacing: 15,
                    mainAxisSpacing: 15,
                    crossAxisCount: 3,
                    children: <Widget>[
                      _getCategoryButton(
                        const Text(
                          'Команды',
                        ),
                        () => {},
                      ),
                      _getCategoryButton(
                        const Text(
                          'События',
                        ),
                        () => {
                          Navigator.push(
                            context,
                            MaterialPageRoute(
                              builder: (context) => EventPage(),
                            ),
                          )
                        },
                      ),
                      _getCategoryButton(
                        const Text(
                          'Достижения',
                        ),
                        () => {
                          Navigator.push(
                            context,
                            MaterialPageRoute(
                              builder: (context) => AchievementPage(),
                            ),
                          )
                        },
                      ),
                      _getCategoryButton(
                        const Text(
                          'Потратить баллы',
                        ),
                        () => {},
                      ),
                      _getCategoryButton(
                        const Text(
                          'Как заработать баллы',
                        ),
                        () => {},
                      ),
                    ],
                  ),
                ),
              ],
            )
          ],
        ),
      ),
    );
  }

  TextButton _getCategoryButton(Text text, Function onPressed) {
    return TextButton(
      style: ButtonStyle(
        shape: MaterialStateProperty.all(
          RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(5.0),
              side: BorderSide(
                color: Colors.blue,
              )),
        ),
      ),
      child: Center(
        child: text,
      ),
      onPressed: onPressed,
    );
  }
}

class MyNavigationBar extends StatefulWidget {
  MyNavigationBar({Key key}) : super(key: key);

  @override
  _MyNavigationBarState createState() => _MyNavigationBarState();
}

class _MyNavigationBarState extends State<MyNavigationBar> {
  int _selectedIndex = 0;
  static List<Widget> _widgetOptions = <Widget>[
    MainPage(),
    Profile(),
  ];

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: _widgetOptions.elementAt(_selectedIndex),
      ),
      bottomNavigationBar: BottomNavigationBar(
        items: <BottomNavigationBarItem>[
          BottomNavigationBarItem(
              icon: Icon(
                Icons.home,
              ),
              title: SizedBox(),
              backgroundColor: Colors.green),
          BottomNavigationBarItem(
              icon: Icon(
                Icons.person,
              ),
              title: SizedBox(),
              backgroundColor: Colors.yellow),
        ],
        type: BottomNavigationBarType.fixed,
        backgroundColor: Colors.blue,
        selectedItemColor: Colors.white,
        unselectedItemColor: Colors.black,
        selectedFontSize: 14,
        unselectedFontSize: 14,
        onTap: _onItemTapped,
        currentIndex: _selectedIndex,
        iconSize: 26,
        elevation: 5,
      ),
    );
  }
}

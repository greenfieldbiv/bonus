import 'package:bivbonus/utils/http-caller.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class TeamsPage extends StatefulWidget {
  @override
  _TeamsPageState createState() => _TeamsPageState();
}

class _TeamsPageState extends State<TeamsPage> {
  List<_TeamMaster> entries = [];
  final httpCaller = HttpCaller();
  ScrollController _controller = ScrollController();

  @override
  void initState() {
    super.initState();
    loadData();
  }

  loadData() async {
    var value = await httpCaller.get<dynamic>(new Uri.http('10.0.2.2:8080', 'team/all'));
    value.map((e) => entries.add(_TeamMaster.fromJson(e))).toList();
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: SingleChildScrollView(
        scrollDirection: Axis.vertical,
        controller: _controller,
        child: Column(
          children: [
            SizedBox(
              height: 35,
            ),
            Row(children: [
              SizedBox(height: 40),
              Padding(
                padding: const EdgeInsets.only(left: 16.0),
                child: Text(
                  'Команды',
                  style: GoogleFonts.lato(color: Colors.grey[700], fontSize: 18, letterSpacing: 1, fontWeight: FontWeight.bold),
                ),
              ),
            ]),
            SizedBox(height: 40),
            Center(
              child: Text(
                'Все команды',
                style: GoogleFonts.lato(color: Colors.grey[600], fontSize: 15, letterSpacing: 1, fontWeight: FontWeight.normal),
              ),
            ),
            ListView.builder(
                padding: const EdgeInsets.only(left: 15.0, right: 15, top: 20),
                itemCount: entries.length,
                shrinkWrap: true,
                controller: _controller,
                itemBuilder: (BuildContext context, int index) {
                  return InkWell(
                    child: Container(
                      height: 90,
                      decoration: BoxDecoration(
                          color: Colors.white,
                          borderRadius: new BorderRadius.only(
                            bottomLeft: const Radius.circular(5.0),
                            bottomRight: const Radius.circular(5.0),
                            topLeft: const Radius.circular(5.0),
                            topRight: const Radius.circular(5.0),
                          ),
                          boxShadow: [
                            BoxShadow(
                              color: Colors.grey[200],
                              blurRadius: 10.0, // soften the shadow
                              spreadRadius: 2.0, //extend the shadow
                              offset: Offset(
                                0, // Move to right 10  horizontally
                                4, // Move to bottom 10 Vertically
                              ),
                            )
                          ]),
                      child: Padding(
                        padding: const EdgeInsets.only(top: 8.0),
                        child: ListTile(
                          leading: CircleAvatar(
                              radius: 40,
                              backgroundImage: NetworkImage(
                                'https://content.api.news/v3/images/bin/a6923adbc7bece73803221613f410782',
                              )),
                          title: Text(
                            entries[index].name,
                            style: GoogleFonts.lato(color: Colors.black, letterSpacing: 1, fontSize: 16, fontWeight: FontWeight.bold),
                          ),
                          subtitle: Text(
                            entries[index].sysname,
                            style: GoogleFonts.lato(color: Colors.grey[700], letterSpacing: 1, fontSize: 13, fontWeight: FontWeight.normal),
                          ),
                          isThreeLine: false,
                        ),
                      ),
                    ),
                  );
                })
          ],
        ),
      ),
    );
  }
}

class _TeamMaster {
  String name;
  String sysname;

  _TeamMaster(this.name, this.sysname);

  factory _TeamMaster.fromJson(dynamic json) {
    return _TeamMaster(
      json['name'] as String,
      json['sysname'] as String,
    );
  }
}

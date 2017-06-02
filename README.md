Devour
========

Devour aims to enhance the dining experience and solve the age old question "what do you want to eat?"

Not only do we care about you, but we care about the people you care about.
We ensure that when you are sitting around a table with the people you love, you are also eating the food you love. 

There is a magic in food and power in breaking bread together. 
We are Devour, helping you eat together, your way.

Here is our current build:
[Homepage](https://dvr.leedann.me)

Here is the everything else:
[Server](https://github.com/leedann/devour/tree/master/devoursvr)
[Client](https://github.com/leedann/devour/tree/master/dvrclient)


Current Status
=======

- Users are able to create an [account](https://dvr.leedann.me/register) and [login](https://dvr.leedann.me/)
  - be sure to [pick](https://dvr.leedann.me/survey) your diets and allergies after registration!
- Creation of an [event](https://dvr.leedann.me/create)
- Able to find your [recipes](https://dvr.leedann.me/recipes)
- NOT able to add friends and friends to event at the moment
- NOT able to add budget, check your recipebook, plan, or friends

You will be able to test group recipe searches using our test account
Username: test@test.com
Password: password

This account is already invited to an event with two other users.
The account has no diet, however is allergic to dairy.

Technology
=======

<img src="http://react-etc.net/files/2016-07/logo-578x270.png" align="left" width="200"/>
Declarative, and Component-based makes creating ui easy! It is also what I am familiar with.

<img src="http://octivi.com/wp-content/uploads/2014/05/redis.png" align="left" width="200"/>
User authentication and store, it helps keep the important information separated from the database. Simple and easy to use.

<img src="https://softwareengineeringdaily.com/wp-content/uploads/2016/10/PostgreSQL.png" align="left" width="200"/>
Open-Sourced database. Helps making many-to-many relationships easy to deal with (we have many).

<img src="https://msopentech.com/wp-content/uploads/dockericon.png" align="left"  width="200"/>
"Escapes the app dependency" matrix allows us to run our client and server without having to worry about dependencies, containers are also secure.

<img src="https://dwa5x7aod66zk.cloudfront.net/assets/pack/logo-digitalocean-3d328c1d6619d314d47aab1259c1235b1339c343e12df62a688076bf6ceac866.jpg" align="left" width="200"/>
For our cloud hosting needs, it has awesome documentation and was easy to learn (I also had a coupon for the service)

<img src="https://www.spaceotechnologies.com/wp-content/uploads/2016/01/go_lang1.png" align="left" width="200"/>
Good documentation, easy, and reliable-- I am the most familiar with Golang when it comes to server development.

Information Needs
=======

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/18/Yummly_logo.png/1200px-Yummly_logo.png" align="left" width="200"/>
For our recipe information needs. They provide a lot of details when it comes to food recipes, most of which we used in Devour.



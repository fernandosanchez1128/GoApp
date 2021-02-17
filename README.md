# GoApp


Rest endpoints

1) sort enpoint 

To use sortEndpoint hit the follow url using POST method
http://localhost:8080/sort

Request should look like 
```json
{
    "unsorted": [2,10,20,2,5,8,7,4,7,8,9,15,20,11,15,20,20]
}
```

Response will look like this
```json
{
"unsorted": [ 2, 10, 20, 2, 5, 8, 7, 4, 7, 8, 9, 15, 20, 11, 15, 20, 20 ],
"sorted": [ 2, 4, 5, 7, 8, 9, 10, 11, 15, 20, 2, 20, 20, 15, 7, 8, 20]
}
```

Notice the sorted list will put repeated elements at the of the list. THis sorting algorithm is bases on MergeSort algorithm and its complexity is O (N*log(N))

2) User info Request

To query for User info you will need to setup a Database. you can setup Db with this script
```sql
{
create table `user`(
  `id` int NOT NULL AUTO_INCREMENT,
  `document` varchar (100) NOT NULL UNIQUE,
  `name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
);


create table `phone` (
  `id` int NOT NULL AUTO_INCREMENT,
  `phone_number` varchar(15) NOT NULL,
  `owner` int NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT FK_PersonOwnwe FOREIGN KEY (owner)
  REFERENCES user(id)
 
) 

SELECT *  FROM phone p

insert into  `user` (document,name, last_name) values 
("1234", "Jess", "Thompson"),
("12345", "Will", "Ortiz");

insert into phone (phone_number, owner)values 
("3352284","1"),
("3127805409","1"),
("3261587","2"),
("3125894058","2");
}
```

* Notice you can easily change DB params on ``resources/config.properties`` file. the project suppose your ${GOPATH} variable point to you Go folder

to hit the endpoint you will need to sen request to this endpoint
http://localhost:8080/user/1234 using GET method. Note: in that request `1234` refers to person document.

it will return a response like this 
``` json
{
    "id": 1,
    "document": "1234",
    "name": "Jess",
    "lastname": "Thompson",
    "phones": [
        {
            "phoneNumber": "3352284"
        },
        {
            "phoneNumber": "3127805409"
        }
    ]
}
```

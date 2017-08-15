# API Beschreibung

## Routen für alle Zugänglich
### Teilnehmer Login Route
`POST     /participant                                 Participant.Login`
#### REQUEST
```json
{
  "name": "TestUser1",
  "password": "test12345#?!\"_-/"
}
```
#### RESPONSE
```json
{
    "status" : "success",
    "data" : {
   		"name":
   		"firstname":
   		"lastname":
        "qrhash": "2039m4c8094875043mxxncowtn"
     }
}

{
    "status" : "fail", //normalerweise fehler 4xx
    "data": {"name":"a name is required"}	//zum beispiel wenn der name fehlt
}

{
    "status" : "error",	//fehler 5xx
    "message" : "Unable to communicate with database"
}
```



### Trainer Login Route

`POST     /instructor                                  Instructor.Login`
#### REQUEST
```json
{
  "name": "TestUser1",
  "password": "test12345#?!\"_-/"
}
```
#### RESPONSE
```json
{
    "status" : "success",
    "data" : {
        "apikey": "2039m4c8094875043mxxncowtn"
     }
}

{
    "status" : "fail",
     "data": {"name":"a name is required"}
}

{
    "status" : "error",
    "message" : "Unable to communicate with database"
}

```







## Routen nur für Kursleiter zugänglich

### Kurs hinzufügen

`POST		/courses/add							Course.Add`

#### REQUEST

```json
{
	"apikey": "ahsdkfjsahdf234234234",
	"name":"kochen",
	"date": "2017-08-07 16:42:03",
}
```

#### RESPONSE

```json


{
    "status" : "success",
    "data" : {     }
}

{
    "status" : "fail",
    "data": {"name":"a name is required"}
}

{
    "status" : "error",
    "message" : "Unable to communicate with database"
}


```

 ### Kurse listen

`POST		/courses						Course.List`

#### REQUEST

```json
{
  "apikey": "09823094caqköldjadf"
}
```



#### RESPONSE

```json
{
	"status": "success",
	"data": {
		[{
				"name": "asd",
				"date": "2017-08-07 16:42:03",
				"instructor-apikey": "0ßc84m5098420",
				"participants": [{
					"vorname": "asd",
					"nachname": "asd",
					"geb": "12.12.2012",
					"qr-code": "BASE64 IMAGE",
					"qr-code-hash": "po8nm4973oc8"
				}, { /*...*/ }]
			},
			{ /*...*/ }
		]
	}
}

{
	"status": "fail",
	"data": {
		"apikey": "aapikey is wrong/required..."
	}
}

{
	"status": "error",
	"message": "Unable to communicate with database"
}
```

### Teilnehmer zu Kurs hinzufügen

`PUT		/course/:id				Course.Addparticipant`

#### REQUEST

```json
{
  "apikey": "09c8m2409mc09r",
  "qr-code-hash": "8m5c098509j8et" // Hash von "vorname_nachname_geb"
}
```

#### RESPONSE

```json
{
  "code": "SUCCESS|ERROR",
  "message": "Teilnehmer nicht gefunden.|Teilnehmer darf nicht teilnehmen."
}

{
    "status" : "success",
    "data" : { }
}

{
    "status" : "fail",
    "data": {"qr-code-hash": "Teilnehmer nicht gefunden.|Teilnehmer darf nicht teilnehmen."}
}

{
    "status" : "error",
    "message" : "Unable to communicate with database"
}

```

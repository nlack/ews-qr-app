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
  "code": "OK|ERROR",
  "message": "blah blah blah" // bei Error leer.,
  "key": "2039m4c8094875043mxxncowtn" // bei Error nicht vorhanden.
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
  "code": "OK|ERROR",
  "message": "blah blah blah" // bei Error leer.,
  "key": "2039m4c8094875043mxxncowtn" // bei Error nicht vorhanden.
}
```











## Routen nur für Teilnehmer zugänglich

### Rufe QR-Code (in Base64 codier?) ab

`POST     /participient/:id                                  Participient.Show`

#### REQUEST

```json
{
  "key": "3m85c0934095c"
}
```



#### RESPONSE

```json
{
  "vorname": "Karl",
  "nachname": "Klaus",
  "geb": "12.12.2012",
  "qr-code": "base64 codierter qr-code?", // oder bild zu image
  "qr-code-hash": "mc8w09e8rmpes98fcos98m9"
}
```



## Routen nur für Kursleiter zugänglich

### Kurs hinzufügen

`POST		/courses/add							Course.Add`

#### REQUEST

```json
{
	"key": "ahsdkfjsahdf234234234",
	"name":"kochen",
	"date": "12.12.2017",
  	"time": "12:00"
}
```

#### RESPONSE

```json
{
  "code": "OK|ERROR",
  "message": "blah blah"
}
```

 ### Kurse listen

`POST		/courses						Course.List`

#### REQUEST

```json
{
  "key": "09823094caqköldjadf"
}
```



#### RESPONSE

```json
[
  {
    "name": "asd",
    "date": "12.12.2017",
    "time": "12:00",
    "instructor-key": "0ßc84m5098420"
    "participants": [
      {
        "vorname": "asd",
        "nachname": "asd",
        "geb": "12.12.2012",
        "qr-code": "BASE64 IMAGE",
        "qr-code-hash": "po8nm4973oc8"
      },
      {...}
    ]
  },
      {
        ...
      }
]
```

### Teilnehmer zu Kurs hinzufügen

`PUT		/course/:id				Course.Addparticipant`

#### REQUEST

```json
{
  "key": "09c8m2409mc09r",
  "qr-code-hash": "8m5c098509j8et" // Hash von "vorname_nachname_geb"
}
```

#### RESPONSE

```json
{
  "code": "SUCCESS|ERROR",
  "message": "Teilnehmer nicht gefunden.|Teilnehmer darf nicht teilnehmen."
}
```




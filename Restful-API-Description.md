# API Beschreibung

## Error
```
"status": "failed | error",
"data": {
	"message": "Bad Request | Unauthorized | Server Error"
}
```

## Routen für alle Zugänglich
### Teilnehmer Login Route
`POST     /participant`
#### REQUEST
```json
{
  "name": "Ulla",
  "password": "ulla123"
}
```
#### RESPONSE
```json
{
    "data": {
        "Name": "Ute",
        "Firstname": "Ute",
        "Lastname": "Grub",
        "Qrhash": "5o9Sype4b57aSlLmAFNuz5Y4K",
        "Haspayed": true
    },
    "status": "success"
}
```



### Trainer Login Route

`POST     /instructor`
#### REQUEST
```json
{
  "name": "Klaus",
  "password": "klaus12345"
}
```
#### RESPONSE
```json
{
    "data": {
        "APIKey": "M7hfjphd3abDUwxwxt8r4gO5q"
    },
    "status": "success"
}
```
## Routen nur für Kursleiter zugänglich
### Kurs hinzufügen

`POST		/courses/add`

#### REQUEST

```json
{
	"apikey": "M7hfjphd3abDUwxwxt8r4gO5q",
	"name":"kochen",
	"date": "2017-08-07 16:42:03"
}
```

#### RESPONSE

```json
{
    "data": null,
    "status": "success"
}
```

 ### Kurse listen

`POST		/courses`

#### REQUEST

```json
{
	"apikey":"M7hfjphd3abDUwxwxt8r4gO5q"
}
```

#### RESPONSE

```json
{
    "data": [
        {
            "id": 1,
            "name": "schwimmen",
            "participants": [
                {
                    "Name": "Ulla",
                    "Firstname": "Ulla",
                    "Lastname": "Urte",
                    "Haspayed": true
                },
                {
                    "Name": "Ulrike",
                    "Firstname": "Ulrike",
                    "Lastname": "Grub",
                    "Haspayed": true
                },
                {
                    "Name": "Undine",
                    "Firstname": "Undine",
                    "Lastname": "Grub",
                    "Haspayed": true
                }
            ],
            "date": "2017-08-07T16:42:03Z",
            "instructor_id": 1
        }
    ],
    "status": "success"
}
```

### Teilnehmer zu Kurs hinzufügen

`PUT		/course/:id`

#### REQUEST

```json
{
  "apikey": "09c8m2409mc09r",
  "qrhash": "8m5c098509j8et"
}
```

#### RESPONSE
```json
{
    "data": null,
    "status": "success"
}
```

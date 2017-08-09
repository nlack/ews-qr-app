package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
	"github.com/knq/dburl"
	"github.com/nlack/ews-qr-app/restapi/models"
)

// This example show a complete (GET,PUT,POST,DELETE) conventional example of
// a REST Resource including documentation to be served by e.g. a Swagger UI
// It is recommended to create a Resource struct (ParticipantResource) that can encapsulate
// an object that provide domain access (a DAO)
// It has a Register method including the complete Route mapping to methods together
// with all the appropriate documentation
//
// POST http://localhost:8080/participants
// <Participant><Id>1</Id><Name>Melissa Raspberry</Name></Participant>
//
// GET http://localhost:8080/participants/1
//
// PUT http://localhost:8080/participants/1
// <Participant><Id>1</Id><Name>Melissa</Name></Participant>
//
// DELETE http://localhost:8080/participants/1
//

var flagVerbose = flag.Bool("v", false, "verbose")

var flagURL = flag.String("url", "mysql://root:asdf@localhost/testtt?parseTime=true&sql_mode=ansi", "url")
var db *sql.DB

type Participant struct {
	Id, Name string
}

type ParticipantResource struct {
	// normally one would use DAO (data access object)
	participant models.Participant
}

func (u ParticipantResource) Register(container *restful.Container) {
	// open database
	var err error
	db, err = dburl.Open(*flagURL)
	if err != nil {
		log.Fatal(err)
	}

	ws := new(restful.WebService)
	ws.
		Path("/participants").
		Doc("Manage Participants").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.GET("/{participant-id}").To(u.findParticipant).
		// docs
		Doc("get a participant").
		Operation("findParticipant").
		Param(ws.PathParameter("participant-id", "identifier of the participant").DataType("string")).
		Writes(Participant{})) // on the response

	container.Add(ws)
}

// GET http://localhost:8080/participants/1
//
func (u ParticipantResource) findParticipant(request *restful.Request, response *restful.Response) {
	id, err := strconv.Atoi(request.PathParameter("participant-id"))
	if err != nil {
		log.Fatal(err)
	}
	usr, err := models.ParticipantByID(db, id)
	if usr == nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "404: Participant could not be found.")
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	response.WriteEntity(usr)
}

func main() {
	// to see what happens in the package, uncomment the following
	//restful.TraceLogger(log.New(os.Stdout, "[restful] ", log.LstdFlags|log.Lshortfile))

	wsContainer := restful.NewContainer()
	var participant models.Participant
	u := ParticipantResource{participant}
	u.Register(wsContainer)

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs and enter http://localhost:8080/apidocs.json in the api input field.
	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:8080",
		ApiPath:        "/apidocs.json",

		// Optionally, specify where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "/Participants/emicklei/xProjects/swagger-ui/dist"}
	swagger.RegisterSwaggerService(config, wsContainer)

	log.Print("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gamegos/jsend"

	"github.com/go-playground/validator"
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

var db *sql.DB
var validate *validator.Validate

type ParticipantResource struct {
	// normally one would use DAO (data access object)
	participant models.Participant
}

func (u ParticipantResource) Register(container *restful.Container) {
	validate = validator.New()
	// open database
	var err error
	db, err = dburl.Open("mysql://root:asdf@localhost/testtt?parseTime=true&sql_mode=ansi")
	if err != nil {
		fmt.Println(err)
	}
	models.XOLog = func(s string, p ...interface{}) {
		fmt.Printf("-------------------------------------\nQUERY: %s\n  VAL: %v\n", s, p)
	}

	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON) // you can specify this per route as well

	ws.Route(ws.GET("/participants/{participant-id}").To(u.findParticipant).
		// docs
		Doc("get a participant").
		Operation("findParticipant").
		Param(ws.PathParameter("participant-id", "identifier of the participant").DataType("string")).
		Writes(models.Participant{})) // on the response

	ws.Route(ws.PUT("/participants/{participant-id}").To(u.updateParticipant).
		// docs
		Doc("update a participant").
		Operation("updateParticipant").
		Param(ws.PathParameter("participant-id", "identifier of the participant").DataType("string")).
		Returns(409, "duplicate participant-id", nil).
		Reads(models.Participant{})) // from the request

	ws.Route(ws.POST("/participants").To(u.createParticipant).
		// docs
		Doc("create a participant").
		Operation("createParticipant").
		Reads(models.Participant{})) // from the request

	ws.Route(ws.DELETE("/participants/{participant-id}").To(u.removeParticipant).
		// docs
		Doc("delete a participant").
		Operation("removeParticipant").
		Param(ws.PathParameter("participant-id", "identifier of the participant").DataType("string")))

	//###############################################
	//###############################################
	//###############################################

	ws.Route(ws.POST("/participant").To(u.loginParticipant).
		// docs
		Doc("participant login").
		Operation("loginParticipant").
		Reads(models.Participant{})) // from the request

	ws.Route(ws.POST("/instructor").To(u.loginInstructor).
		// docs
		Doc("instructor login").
		Operation("loginInstructor").
		Reads(models.Instructor{})) // from the request

	ws.Route(ws.POST("/courses/add").To(u.addCourse).
		// docs
		Doc("add course").
		Operation("addCourse").
		Reads(models.Instructor{})) // from the request
	//TODO READS??

	ws.Route(ws.POST("/courses").To(u.listCourses).
		// docs
		Doc("list courses").
		Operation("listCourses").
		Reads(models.Instructor{})) // from the request
	container.Add(ws)
}

func (u *ParticipantResource) addCourse(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")
	courseInfos := new(struct {
		Apikey string
		Name   string
		Date   string
	})
	err := request.ReadEntity(&courseInfos)
	if err != nil {
		jsend.Wrap(response.ResponseWriter).Status(http.StatusInternalServerError).Message(err.Error()).Send()
		log.Fatal(err.Error())
		return
	}

	err = validate.Var(courseInfos.Name, "required")
	if err != nil {
		a := new(struct{ Name string })
		a.Name = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusNotFound).Data(a).Send()
		return
	}
	err = validate.Var(courseInfos.Apikey, "required")
	if err != nil {
		a := new(struct{ Apikey string })
		a.Apikey = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusNotFound).Data(a).Send()
		return
	}
	err = validate.Var(courseInfos.Date, "required")
	if err != nil {
		a := new(struct{ Date string })
		a.Date = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusNotFound).Data(a).Send()
		return
	}

	usr, err := models.InstructorByAPIKey(db, courseInfos.Apikey)
	if err != nil {
		a := new(struct{ Apikey string })
		a.Apikey = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusNotFound).Data(a).Send()
		return
	}
	course := new(models.Course)
	course.Name = courseInfos.Name
	course.InstructorID = usr.ID
	course.Date, err = time.Parse("2006-01-02 15:04:05", courseInfos.Date)
	if err != nil {
		jsend.Wrap(response.ResponseWriter).Status(http.StatusInternalServerError).Message(err.Error()).Send()
		log.Fatal(err.Error())
		return
	}
	course.Save(db)

}

func (u *ParticipantResource) listCourses(request *restful.Request, response *restful.Response) {
	a := new(struct{ Apikey string })
	err := request.ReadEntity(&a)
	if err != nil {
		a.Apikey = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusNotFound).Data(a).Send()
		return
	}

	_, err = models.InstructorByAPIKey(db, a.Apikey)
	if err != nil {
		a := new(struct{ Apikey string })
		a.Apikey = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusNotFound).Data(a).Send()
		return
	}
	courses, err := models.ListCourses(db)
	if err != nil {
		jsend.Wrap(response.ResponseWriter).Status(http.StatusInternalServerError).Message(err.Error()).Send()
		log.Fatal(err.Error())
		return
	}
	jsend.Wrap(response.ResponseWriter).Status(http.StatusAccepted).Data(courses).Send()

}

// POST http://localhost:8080/participant
func (u *ParticipantResource) loginParticipant(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")
	p := new(models.Participant)
	err := request.ReadEntity(&p)
	if err != nil {
		jsend.Wrap(response.ResponseWriter).Status(http.StatusInternalServerError).Message(err.Error()).Send()
		log.Fatal(err.Error())
		return
	}
	err = validate.Var(p.Name, "required")
	if err != nil {
		a := new(struct{ Name string })
		a.Name = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusExpectationFailed).Data(a).Send()
		return
	}

	err = validate.Var(p.Password, "required")
	if err != nil {
		a := new(struct{ Password string })
		a.Password = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusExpectationFailed).Data(a).Send()
		return
	}

	pwd := p.Password

	p, err = models.ParticipantByName(db, p.Name)
	if err != nil {
		a := new(struct{ Name string })
		a.Name = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusNotFound).Data(a).Send()
		return
	}

	err = validate.Var(p.Password, fmt.Sprintf("eq=%s", pwd))
	if err != nil {
		a := new(struct{ Password string })
		a.Password = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusUnauthorized).Data(a).Send()
		return
	}

	jsend.Wrap(response.ResponseWriter).Status(http.StatusAccepted).Data(p).Send()
}

// POST http://localhost:8080/instructor
//
func (u *ParticipantResource) loginInstructor(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")
	usr := new(models.Instructor)
	err := request.ReadEntity(&usr)
	if err != nil {
		jsend.Wrap(response.ResponseWriter).Status(http.StatusInternalServerError).Message(err.Error()).Send()
		log.Fatal(err.Error())
		return
	}
	err = validate.Var(usr.Name, "required")
	if err != nil {
		a := new(struct{ Name string })
		a.Name = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusExpectationFailed).Data(a).Send()
		return
	}

	err = validate.Var(usr.Password, "required")
	if err != nil {
		a := new(struct{ Password string })
		a.Password = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusExpectationFailed).Data(a).Send()
		return
	}

	pwd := usr.Password

	usr, err = models.InstructorByName(db, usr.Name)
	if err != nil {
		a := new(struct{ Name string })
		a.Name = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusNotFound).Data(a).Send()
		return
	}

	err = validate.Var(usr.Password, fmt.Sprintf("eq=%s", pwd))
	if err != nil {
		a := new(struct{ Password string })
		a.Password = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusUnauthorized).Data(a).Send()
		return
	}

	a := new(struct{ APIKey string })
	a.APIKey = usr.Apikey
	jsend.Wrap(response.ResponseWriter).Status(http.StatusAccepted).Data(a).Send()
}

// GET http://localhost:8080/participants/1
//
func (u ParticipantResource) findParticipant(request *restful.Request, response *restful.Response) {
	id, err := strconv.Atoi(request.PathParameter("participant-id"))
	if err != nil {
		fmt.Println(err)
	}
	usr, err := models.ParticipantByID(db, id)
	if usr == nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "404: Participant could not be found.")
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	response.WriteEntity(usr)
}

// POST http://localhost:8080/participants
// <Participant><Name>Melissa</Name></Participant>
//
func (u *ParticipantResource) createParticipant(request *restful.Request, response *restful.Response) {
	usr := new(models.Participant)
	err := request.ReadEntity(&usr)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(err)
		return
	}
	err = validate.Struct(*usr)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(err.Error())
		return
	}
	err = usr.Save(db)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(err)
		return
	}
	//usr.Id = strconv.Itoa(len(u.participants) + 1) // simple id generation
	//u.participants[usr.Id] = *usr
	err = response.WriteHeaderAndEntity(http.StatusCreated, usr)
	if err != nil {
		fmt.Println(err)
	}
}

// PUT http://localhost:8080/participants/1
// <Participant><Id>1</Id><Name>Melissa Raspberry</Name></Participant>
//
func (u *ParticipantResource) updateParticipant(request *restful.Request, response *restful.Response) {
	id, err := strconv.Atoi(request.PathParameter("participant-id"))
	if err != nil {
		fmt.Println(err)
	}
	usr, err := models.ParticipantByID(db, id)
	if err != nil {
		fmt.Println(err)
	}
	err = request.ReadEntity(&usr)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(err)
		return
	}
	usr.ID = id
	err = usr.Update(db)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(err)
		return
	}
	response.WriteEntity(usr)
}

// DELETE http://localhost:8080/participants/1
//
func (u *ParticipantResource) removeParticipant(request *restful.Request, response *restful.Response) {
	id, err := strconv.Atoi(request.PathParameter("participant-id"))
	if err != nil {
		fmt.Println(err)
	}
	usr, err := models.ParticipantByID(db, id)
	if err != nil {
		fmt.Println(err)
	}
	err = usr.Delete(db)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(err)
		return
	}
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
	fmt.Println(server.ListenAndServe())
}

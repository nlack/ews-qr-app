package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
	"github.com/gamegos/jsend"
	"github.com/go-errors/errors"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

var db *sql.DB
var validate *validator.Validate
var noproxy = flag.Bool("noproxy", false, "no proxy")

type ParticipantResource struct { //TODO ??
	participant models.Participant
}

func (u ParticipantResource) Register(container *restful.Container) {
	validate = validator.New()
	// open database
	var err error
	db, err = dburl.Open("mysql://" + os.Getenv("DBUser") + ":" + os.Getenv("DBPassword") + "@" + os.Getenv("DBHost") + "/" + os.Getenv("DBName") + "?parseTime=true&sql_mode=ansi")
	if err != nil {
		fmt.Println(errors.Wrap(err, 1))
	}
	models.XOLog = func(s string, p ...interface{}) {
		fmt.Printf("-------------------------------------\nQUERY: %s\n  VAL: %v\n", s, p)
	}

	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST"},
		CookiesAllowed: false,
		Container:      container}

	container.Filter(cors.Filter)

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

	ws.Route(ws.POST("/participant/add").To(u.createParticipant).
		// docs
		Doc("create a participant").
		Operation("createParticipant").
		Reads(models.Participant{})) // from the request

	ws.Route(ws.POST("/instructor/add").To(u.createInstructor).
		// docs
		Doc("create a instructor").
		Operation("createInstructor").
		Reads(models.Instructor{})) // from the request

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

	ws.Route(ws.PUT("/course/{course-id}").To(u.addParticipant).
		// docs
		Doc("add participant to course").
		Operation("addParticipant").
		Reads(models.Participant{})) // from the request
	//TODO reads?
	ws.Route(ws.POST("/courses").To(u.listCourses).
		// docs
		Doc("list courses").
		Operation("listCourses").
		Reads(models.Instructor{})) // from the request
	container.Add(ws)
}

func setHeaders(response *restful.Response) {
	response.AddHeader("Access-Control-Allow-Origin", "*")
	response.AddHeader("Access-Control-Allow-Credentials", "true")
	response.AddHeader("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	response.AddHeader("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
}

func serverError(response *restful.Response, err error) {
	jsend.Wrap(response.ResponseWriter).Status(http.StatusInternalServerError).Message(http.StatusText(http.StatusInternalServerError)).Send()
	fmt.Println(errors.Wrap(err, 0).ErrorStack())
}

func badRequest(response *restful.Response, err error) {
	jsend.Wrap(response.ResponseWriter).Status(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Send()
	fmt.Println(errors.Wrap(err, 0).ErrorStack())
}

func (u *ParticipantResource) addParticipant(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")
	if *noproxy {
		setHeaders(response)
	}
	a := new(struct {
		Apikey string
		Qrhash string
	})

	var err error
	err = request.ReadEntity(&a)
	if err != nil {
		serverError(response, err)
		return
	}
	fmt.Println("apikey" + a.Apikey)
	_, err = models.InstructorByAPIKey(db, a.Apikey)
	if err != nil {
		badRequest(response, err)
		return
	}

	coursePart := new(models.Courseparticipant)

	part, err := models.ParticipantByQrhash(db, a.Qrhash)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			jsend.Wrap(response.ResponseWriter).Status(http.StatusBadRequest).Message("Participant already added").Send()
			fmt.Println(errors.Wrap(err, 0).ErrorStack())
			return
		}
		badRequest(response, err)
		return
	}
	if !part.Haspayed {
		jsend.Wrap(response.ResponseWriter).Status(http.StatusBadRequest).Message("Participant has not payed").Send()
		fmt.Println(errors.Wrap(err, 0).ErrorStack())
		return
	}
	coursePart.Courseid, err = strconv.Atoi(request.PathParameter("course-id"))
	if err != nil {
		serverError(response, err)
		return
	}
	coursePart.Participantid = part.ID
	err = coursePart.Save(db)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			jsend.Wrap(response.ResponseWriter).Status(http.StatusBadRequest).Message("Participant already added").Send()
			fmt.Println(errors.Wrap(err, 0).ErrorStack())
			return
		}
		badRequest(response, err)
		return
	}
	jsend.Wrap(response.ResponseWriter).Status(http.StatusOK).Send()

}

func (u *ParticipantResource) addCourse(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")

	if *noproxy {
		setHeaders(response)
	}

	courseInfos := new(struct {
		Apikey string
		Name   string
		Date   string
	})
	err := request.ReadEntity(&courseInfos)
	if err != nil {
		serverError(response, err)
		return
	}

	err = validate.Var(courseInfos.Name, "required")
	if err != nil {
		badRequest(response, err)
		return
	}

	usr, err := models.InstructorByAPIKey(db, courseInfos.Apikey)
	if err != nil {
		// a := new(struct{ Apikey string })
		// a.Apikey = err.Error()
		// jsend.Wrap(response.ResponseWriter).Status(http.StatusUnauthorized).Send()
		// fmt.Println(errors.Wrap(err, 1))
		badRequest(response, err)
		return
	}
	course := new(models.Course)
	course.Name = courseInfos.Name
	course.InstructorID = usr.ID
	course.Date, err = time.Parse("2006-01-02 15:04:05", courseInfos.Date)
	if err != nil {
		badRequest(response, err)
		return
	}
	err = course.Save(db)
	if err != nil {
		serverError(response, err)
		return
	}
	jsend.Wrap(response.ResponseWriter).Status(http.StatusOK).Send()
}

func (u *ParticipantResource) listCourses(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")
	if *noproxy {
		setHeaders(response)
	}

	a := new(struct{ Apikey string })
	err := request.ReadEntity(&a)
	if err != nil {
		serverError(response, err)
		return
	}

	_, err = models.InstructorByAPIKey(db, a.Apikey)
	if err != nil {
		jsend.Wrap(response.ResponseWriter).Status(http.StatusUnauthorized).Message(http.StatusText(http.StatusUnauthorized)).Send()
		fmt.Println(errors.Wrap(err, 1))
		return
	}

	courses, err := models.ListCourses(db)
	if err != nil {
		serverError(response, err)
		return
	}
	for _, c := range courses {
		cps, err := models.CourseparticipantsByCourseid(db, c.ID)
		if err != nil {
			serverError(response, err)
			return
		}
		for _, cp := range cps {
			part, err := cp.Participant(db)
			a := new(struct {
				Name      string
				Firstname string
				Lastname  string
				Haspayed  bool
			})
			a.Name = part.Name
			a.Firstname = part.Firstname
			a.Lastname = part.Lastname
			a.Haspayed = part.Haspayed
			if err != nil {
				serverError(response, err)
				return
			}
			c.Participants = append(c.Participants, *a)
		}
	}

	jsend.Wrap(response.ResponseWriter).Status(http.StatusOK).Data(courses).Send()
}

// POST http://localhost:8080/participant
func (u *ParticipantResource) loginParticipant(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")
	if *noproxy {
		setHeaders(response)
	}
	p := new(models.Participant)
	err := request.ReadEntity(&p)
	if err != nil {
		serverError(response, err)
		return
	}
	/*err = validate.Var(p.Name, "required")
	if err != nil {
		a := new(struct{ Name string })
		a.Name = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusExpectationFailed).Data(a).Send()
		fmt.Println(errors.Wrap(err, 1))
		return
	}

	err = validate.Var(p.Password, "required")
	if err != nil {
		a := new(struct{ Password string })
		a.Password = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusExpectationFailed).Data(a).Send()
		fmt.Println(errors.Wrap(err, 1))
		return
	}
	*/
	pwd := p.Password

	p, err = models.ParticipantByName(db, p.Name)
	if err != nil {
		// a := new(struct{ Name string })
		// a.Name = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Send()
		fmt.Println(errors.Wrap(err, 0).ErrorStack())
		return
	}

	err = validate.Var(p.Password, fmt.Sprintf("eq=%s", pwd))
	if err != nil {
		// a := new(struct{ Password string })
		// a.Password = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusUnauthorized).Message(http.StatusText(http.StatusUnauthorized)).Send()
		fmt.Println(errors.Wrap(err, 1))
		return
	}

	jsend.Wrap(response.ResponseWriter).Status(http.StatusOK).Data(p).Send()
}

// POST http://localhost:8080/instructor
//
func (u *ParticipantResource) loginInstructor(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")
	if *noproxy {
		setHeaders(response)
	}
	usr := new(models.Instructor)
	err := request.ReadEntity(&usr)
	if err != nil {
		serverError(response, err)
		return
	}
	// err = validate.Var(usr.Name, "required")
	// if err != nil {
	// 	a := new(struct{ Name string })
	// 	a.Name = err.Error()
	// 	jsend.Wrap(response.ResponseWriter).Status(http.StatusExpectationFailed).Data(a).Send()
	// 	fmt.Println(errors.Wrap(err, 1))
	// 	return
	// }
	//
	// err = validate.Var(usr.Password, "required")
	// if err != nil {
	// 	a := new(struct{ Password string })
	// 	a.Password = err.Error()
	// 	jsend.Wrap(response.ResponseWriter).Status(http.StatusExpectationFailed).Data(a).Send()
	// 	fmt.Println(errors.Wrap(err, 1))
	// 	return
	// }

	pwd := usr.Password

	usr, err = models.InstructorByName(db, usr.Name)
	if err != nil {
		// a := new(struct{ Name string })
		// a.Name = err.Error()
		badRequest(response, err)
		return
	}

	err = validate.Var(usr.Password, fmt.Sprintf("eq=%s", pwd))
	if err != nil {
		// a := new(struct{ Password string })
		// a.Password = err.Error()
		jsend.Wrap(response.ResponseWriter).Status(http.StatusUnauthorized).Message(http.StatusText(http.StatusUnauthorized)).Send()
		fmt.Println(errors.Wrap(err, 1))
		return
	}
	a := new(struct{ APIKey string })
	a.APIKey = usr.Apikey
	jsend.Wrap(response.ResponseWriter).Status(http.StatusOK).Data(a).Send()
}

// GET http://localhost:8080/participants/1
//
func (u ParticipantResource) findParticipant(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")
	if *noproxy {
		setHeaders(response)
	}
	id, err := strconv.Atoi(request.PathParameter("participant-id"))
	if err != nil {
		serverError(response, err)
		return
	}
	usr, err := models.ParticipantByID(db, id)
	if err != nil {
		badRequest(response, err)
		return
	}
	usr.Password = ""

	jsend.Wrap(response.ResponseWriter).Status(http.StatusOK).Data(usr).Send()

}

// POST http://localhost:8080/participants
// <Participant><Name>Melissa</Name></Participant>
//
func (u *ParticipantResource) createParticipant(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json") //TODO die headers prüfen
	if *noproxy {
		setHeaders(response)
	}
	usr := new(models.Participant)
	err := request.ReadEntity(&usr)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(errors.Wrap(err, 1))
		return
	}
	usr.Haspayed = true
	usr.Qrhash = randomString(25)
	err = validate.Struct(*usr)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(errors.Wrap(err, 1))
		return
	}
	err = usr.Save(db)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(errors.Wrap(err, 1))
		return
	}
	//usr.Id = strconv.Itoa(len(u.participants) + 1) // simple id generation
	//u.participants[usr.Id] = *usr
	err = response.WriteHeaderAndEntity(http.StatusCreated, usr)
	if err != nil {
		fmt.Println(errors.Wrap(err, 1))
	}
}

// POST http://localhost:8080/participants
// <Participant><Name>Melissa</Name></Participant>
//
func (u *ParticipantResource) createInstructor(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")
	if *noproxy {
		setHeaders(response)
	}
	usr := new(models.Instructor)
	err := request.ReadEntity(&usr)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(errors.Wrap(err, 1))
		return
	}
	usr.Apikey = randomString(25)
	err = validate.Struct(*usr)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(errors.Wrap(err, 1))
		return
	}
	err = usr.Save(db)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(errors.Wrap(err, 1))
		return
	}
	//usr.Id = strconv.Itoa(len(u.participants) + 1) // simple id generation
	//u.participants[usr.Id] = *usr
	err = response.WriteHeaderAndEntity(http.StatusCreated, usr)
	if err != nil {
		fmt.Println(errors.Wrap(err, 1))
	}
}

// PUT http://localhost:8080/participants/1
// <Participant><Id>1</Id><Name>Melissa Raspberry</Name></Participant>
//
func (u *ParticipantResource) updateParticipant(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")
	if *noproxy {
		setHeaders(response)
	}
	id, err := strconv.Atoi(request.PathParameter("participant-id"))
	if err != nil {
		fmt.Println(errors.Wrap(err, 1))
	}
	usr, err := models.ParticipantByID(db, id)
	if err != nil {
		fmt.Println(errors.Wrap(err, 1))
	}
	err = request.ReadEntity(&usr)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(errors.Wrap(err, 1))
		return
	}
	usr.ID = id
	err = usr.Update(db)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(errors.Wrap(err, 1))
		return
	}
	response.WriteEntity(usr)
}

// DELETE http://localhost:8080/participants/1
//
func (u *ParticipantResource) removeParticipant(request *restful.Request, response *restful.Response) {
	response.AddHeader("Content-Type", "application/json")
	if *noproxy {
		setHeaders(response)
	}
	id, err := strconv.Atoi(request.PathParameter("participant-id"))
	if err != nil {
		fmt.Println(errors.Wrap(err, 1))
	}
	usr, err := models.ParticipantByID(db, id)
	if err != nil {
		fmt.Println(errors.Wrap(err, 1))
	}
	err = usr.Delete(db)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		fmt.Println(errors.Wrap(err, 1))
		return
	}
}

func main() {
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		fmt.Println(errors.Wrap(err, 1))
	}
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
		WebServicesUrl: "http://localhost:" + os.Getenv("APIPORT"),
		ApiPath:        "/apidocs.json",

		// Optionally, specify where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "/Participants/emicklei/xProjects/swagger-ui/dist"}
	swagger.RegisterSwaggerService(config, wsContainer)

	log.Print("start listening on localhost:" + os.Getenv("APIPORT"))
	server := &http.Server{Addr: ":" + os.Getenv("APIPORT"), Handler: wsContainer}
	fmt.Println(server.ListenAndServe())
}

package controllers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nlack/ews-qr-app/RevelEWS/app/models"
	"github.com/revel/revel"
)

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	host := getParamString("db.host", "")
	port := getParamString("db.port", "")
	user := getParamString("db.user", "")

	pass := "asdf"

	dbname := getParamString("db.name", "")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("dbargs", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}

var InitDb func() = func() {
	connectionString := getConnectionString()
	if db, err := sql.Open("mysql", connectionString); err != nil {
		revel.ERROR.Fatal(err)
	} else {
		Dbm = &gorp.DbMap{
			Db:      db,
			Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	}
	// Defines the table for use by GORP
	// This is a function we will create soon.
	defineAndFillTables(Dbm)

	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		revel.ERROR.Fatal(err)
	}
}

func defineAndFillTables(dbm *gorp.DbMap) {

	// set "id" as primary key and autoincrement
	t := dbm.AddTable(models.Course{}).SetKeys(true, "id")
	//TODO AddTable Rel Course Participant
	t = dbm.AddTable(models.Instructor{}).SetKeys(true, "id")
	t.ColMap("username").SetUnique(true)
	t = dbm.AddTable(models.Participant{}).SetKeys(true, "id")
	t.ColMap("username").SetUnique(true)

	//TODO Insert funktioniert noch nicht
	instr := models.Participant{55, "rollcage", "asdf1234", "Michael", "Sann", "9999ddddd", "kkkkk55555"}
	dbm.Insert(instr)
}

func init() {
	revel.OnAppStart(InitDb)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}
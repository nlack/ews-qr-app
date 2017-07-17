package models

import "github.com/revel/revel"

type Course struct {
	Id   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Date string `db:"date" json:"date"`
	// TODO: array
	Participants  string `db:"participants" json:"participants"`
	InstructorKey string `db:"instructorkey" json:"instructorkey"`
}

func (b *Course) Validate(v *revel.Validation) {

	v.Check(b.Name,
		revel.ValidRequired(),
		revel.ValidMaxSize(25))
}

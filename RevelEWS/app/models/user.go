package models

import "github.com/revel/revel"

type User struct {
	Id         int64  `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Instructor bool   `db:"instructor" json:"instructor"`
}

func (b *User) Validate(v *revel.Validation) {

	v.Check(b.Name,
		revel.ValidRequired(),
		revel.ValidMaxSize(25))

	/*v.Check(b.Category,
	revel.ValidMatch(
		regexp.MustCompile(
			"^(travel|leasure|sports|entertainment)$")))
	*/
}

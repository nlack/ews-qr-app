package models

import "github.com/revel/revel"

type CourseInstructor struct {
	Id   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (b *CourseInstructor) Validate(v *revel.Validation) {

	v.Check(b.Name,
		revel.ValidRequired(),
		revel.ValidMaxSize(25))

	/*v.Check(b.Category,
		revel.ValidRequired(),
		revel.ValidMatch(
			regexp.MustCompile(
				"^(travel|leasure|sports|entertainment)$")))

	v.Check(b.EstimatedValue,
		revel.ValidRequired())

	v.Check(b.StartBid,
		revel.ValidRequired())

	v.Check(b.BidIncrement,
		revel.ValidRequired())
	*/
}

package models

import "github.com/revel/revel"

type Participant struct {
	Id        int    `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	AccessKey string `db:"accesskey" json:"accesskey"`
	QRHash    string `db:"qrhash" json:"qrhash"`
}

func (u *Participant) Validate(v *revel.Validation) {

	v.Check(u.Firstname,
		revel.ValidRequired(),
		revel.ValidMaxSize(25))

	v.Check(u.Lastname,
		revel.ValidRequired(),
		revel.ValidMaxSize(25))
	//TODO how long should accesskey be?
	v.Check(u.AccessKey,
		revel.ValidRequired(),
		revel.ValidLength(32))
	//TODO how long should the qr code be?
	v.Check(u.QRHash,
		revel.ValidRequired(),
		revel.ValidLength(32))
}

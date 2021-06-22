package model

import (
	"github.com/astaxie/beego/orm"
)

type Developer struct {
	Id          int    `orm:column(id);auto`
	TeamId      *Team  `orm:column(teamId);rel(fk)`
	Name        string `orm:column(name);size(64);null`
	PhoneNumber string `orm:column(phoneNumber);size(64);null`
}

const (
	Developers = "Developer"
)

func (t *Developer) TableName() string {
	return Developers
}

func init() {
	orm.RegisterModel(new(Developer))

}

func CreateDeveloper(t *Developer) *Developer {
	o := orm.NewOrm()
	_, err := o.Insert(t)

	if err != nil {
		panic(err)
	}
	return t

}

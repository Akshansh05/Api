package model

import (
	"../config"

	"github.com/astaxie/beego/orm"
)

type Team struct {
	Id   int    `orm:column(id);auto`
	Name string `orm:column(name);size(64);null`
}

const (
	Teams = "Team"
)

func (t *Team) TableName() string {
	return Teams
}

func init() {
	config.Connect()
	orm.RegisterModel(new(Team))

}

func CreateTeam(t *Team) *Team {
	o := orm.NewOrm()
	_, err := o.Insert(t)

	if err != nil {
		panic(err)
	}
	return t

}

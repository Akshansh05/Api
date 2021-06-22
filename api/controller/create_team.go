package controller

import (
	"net/http"

	"../model"
	"../util"
)

type CreateTeamJSON struct {
	Team       Teams       `json:"team"`
	Developers []Developer `json:""developers""`
}

type Teams struct {
	Name string `json:"name"`
}

type Developer struct {
	Name        string `json:"name"`
	PhoneNumber string `json:""phone_number""`
}

func covertJSONTOObject() {

}

func CreateTeam(w http.ResponseWriter, r *http.Request) {

	teamJSON := CreateTeamJSON{}
	//util.ParseBody(r, teamJSON)
	util.ParseBody(r, teamJSON)
	teamObject := &model.Team{}
	teamObject.Name = teamJSON.Team.Name
	t := model.CreateTeam(teamObject)

	developerObject := &model.Developer{}
	developerObject.TeamId = t

	developerObject.Name = teamJSON.Developers[0].Name
	developerObject.PhoneNumber = teamJSON.Developers[0].PhoneNumber

	//d := model.CreateDeveloper(developerObject)

}

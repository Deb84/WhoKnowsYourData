package server

import (
	"fmt"
)

type Routes struct{}

const (
	// params, used as enum
	ParamLabel     = "label"
	ParamValueUUID = "uuid"

	//API
	api       = "/api"
	value     = api + "/value"
	values    = api + "/values"
	relation  = api + "/relation"
	relations = api + "/relations"

	// Public
	register = "/register"
	login    = "/login"
)

func NewRoutes() *Routes {
	return &Routes{}
}

// API

func (routes *Routes) APIValue() string {
	return value
}

func (routes *Routes) APIValues() string {
	return values
}

func (routes *Routes) APIRelation() string {
	return relation
}

func (routes *Routes) APIRelations() string {
	return relations
}

func (routes *Routes) APIValuesLabel() string {
	return fmt.Sprintf("%s/{%s}", values, ParamLabel)
}

func (routes *Routes) APIValueUUID() string {
	return fmt.Sprintf("%s/{%s}", value, ParamValueUUID)
}

// Public

func (routes *Routes) Register() string {
	return register
}

func (routes *Routes) Login() string {
	return login
}

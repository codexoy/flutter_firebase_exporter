package models

type Registration struct {
	Competition      string `json:"competition"`
	Occupation       string `json:"occupation"`
	ReasonToAttend   string `json:"reasonToAttend"`
	IsBringingLaptop bool   `json:"laptop"`
}

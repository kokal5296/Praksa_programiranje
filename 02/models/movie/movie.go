package movie

//Movie podatki

import (
	"github.com/kokal5296/Praksa_Vaje/models/actor"
	"github.com/kokal5296/Praksa_Vaje/models/director"
)

type Movie struct {
	ID       string             `json:"id"`
	Title    string             `json:"title"`
	Director *director.Director `json:"director"`
	Actor    *actor.Actor       `json:"actor"`
}

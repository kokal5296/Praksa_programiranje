package storage

//virtualna databaza

import (
	"github.com/kokal5296/Praksa_Vaje/models/actor"
	"github.com/kokal5296/Praksa_Vaje/models/director"
	"github.com/kokal5296/Praksa_Vaje/models/movie"
)

func Movies() ([]movie.Movie, error) {
	return []movie.Movie{
		{ID: "1", Title: "LOLA",
			Director: &director.Director{ID: "4", FirstName: "tine", LastName: "kokalj"},
			Actor:    &actor.Actor{ID: "7", FirstName: "koki", LastName: "keke"}},

		{ID: "2", Title: "lala",
			Director: &director.Director{ID: "5", FirstName: "zan", LastName: "horvat"},
			Actor:    &actor.Actor{ID: "8", FirstName: "zanko", LastName: "horvi"}},

		{ID: "3", Title: "hehe",
			Director: &director.Director{ID: "6", FirstName: "luka", LastName: "potocnik"},
			Actor:    &actor.Actor{ID: "9", FirstName: "luki", LastName: "tocnik"}},
	}, nil

}

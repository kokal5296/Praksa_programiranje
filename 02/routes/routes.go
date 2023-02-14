package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kokal5296/Praksa_Vaje/models/actor"
	"github.com/kokal5296/Praksa_Vaje/models/director"
	"github.com/kokal5296/Praksa_Vaje/models/movie"
	"github.com/kokal5296/Praksa_Vaje/storage"
)

//api

func GetMovies(c *gin.Context) {
	movies, err := storage.Movies()
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, movies)
}

func GetActors(c *gin.Context) {
	movies, err := storage.Movies()
	if err != nil {
		panic(err)
	}
	for _, b := range movies {
		c.IndentedJSON(http.StatusOK, b.Actor)
	}
}

func GetDirectors(c *gin.Context) {
	movies, err := storage.Movies()
	if err != nil {
		panic(err)
	}
	for _, b := range movies {
		c.IndentedJSON(http.StatusOK, b.Director)
	}
}

func MovieByID(c *gin.Context) {
	id := c.Param("id")
	movie, err := GetMovieByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
	} else {
		c.IndentedJSON(http.StatusOK, movie)
	}
}

func ActorByID(c *gin.Context) {
	id := c.Param("id")
	actors, err := GetActorByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Actor not found"})
	} else {
		c.IndentedJSON(http.StatusOK, actors)
	}
}

func DirectorByID(c *gin.Context) {
	id := c.Param("id")
	director, err := GetDirectorByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Director not found"})
	} else {
		c.IndentedJSON(http.StatusOK, director)
	}
}

func GetMovieByID(id string) (*movie.Movie, error) {
	movies, err := storage.Movies()
	if err != nil {
		panic(err)
	}
	for i, b := range movies {
		if b.ID == id {
			return &movies[i], nil
		}
	}
	return nil, errors.New("Movie not found")
}

func GetActorByID(id string) (*actor.Actor, error) {
	movies, err := storage.Movies()
	if err != nil {
		panic(err)
	}
	for _, b := range movies {
		if b.Actor.ID == id {

			return b.Actor, nil
		}
	}
	return nil, errors.New("Actor not found")
}

func GetDirectorByID(id string) (*director.Director, error) {
	movies, err := storage.Movies()
	if err != nil {
		panic(err)
	}
	for _, b := range movies {
		if b.Director.ID == id {
			return b.Director, nil
		}
	}
	return nil, errors.New("Director not found")
}

func NewMovie(c *gin.Context) {
	newMovie := movie.Movie{}
	movies, kaj := storage.Movies()
	if kaj != nil {
		panic(kaj)
	}
	if err := c.BindJSON(&newMovie); err != nil {
		return
	}
	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)
}

func FixActor(c *gin.Context) {
	newActor := movie.Movie{}
	movies, kaj := storage.Movies()
	if kaj != nil {
		panic(kaj)
	}
	if err := c.BindJSON(&newActor.Actor); err != nil {
		return
	}
	movies = append(movies, newActor)
	c.IndentedJSON(http.StatusOK, newActor.Actor)
}

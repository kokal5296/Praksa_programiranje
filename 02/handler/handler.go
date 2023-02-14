package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/kokal5296/Praksa_Vaje/routes"
)

//REST povezave

func Handler() {

	router := gin.Default()
	router.GET("/movies", routes.GetMovies)
	router.GET("/actors", routes.GetActors)
	router.GET("/directors", routes.GetDirectors)
	router.GET("/movie/:id", routes.MovieByID)
	router.GET("/actor/:id", routes.ActorByID)
	router.GET("/director/:id", routes.DirectorByID)
	router.POST("/movie", routes.NewMovie)
	router.PUT("/movie/actor/:id", routes.FixActor)

	if err := router.Run("localhost:3000"); err != nil {
		panic(err)
	}

}

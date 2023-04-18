package Controller

import (
	Api2 "github.com/403-access-denied/main-backend/src/MainService/Controller/Api"
	"github.com/403-access-denied/main-backend/src/MainService/Token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router     *gin.Engine
	TokenMaker Token.Maker
}

func (s *Server) MainController() {
	v1 := s.Router.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/auth/otp/send", Api2.SendOTP)
			user.POST("/auth/otp/login", Api2.LoginUser)
			user.GET("/auth/google/login", Api2.OAuth2Login)
			user.GET("/auth/google/callback", Api2.GoogleCallback)
		}
		// TODO add auth middleware
		poster := v1.Group("/posters")
		{
			poster.GET("/", Api2.GetPosters)
			poster.GET("/:id", Api2.GetPoster)
			poster.POST("/", Api2.CreatePoster)
			poster.PATCH("/:id", Api2.UpdatePoster)
			poster.DELETE("/:id", Api2.DeletePoster)
		}
	}

}
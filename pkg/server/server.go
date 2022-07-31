package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type IServer interface {
	Start(port string)
	App() *gin.Engine
}

type server struct {
	app *gin.Engine
}

func (s *server) App() *gin.Engine {
	return s.app
}

func (s *server) Start(port string) {
	log.Info().Str("port", port).Msg("starting server")
	s.app.Run(fmt.Sprintf("0.0.0.0:%s", port))
}

func NewServer() IServer {
	router := gin.Default()
	router.Use(cors.Default())
	return &server{
		app: router,
	}
}

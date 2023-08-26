package main

import (
	presentation "bookManagement/Presentation"
	"bookManagement/config"
	"bookManagement/db"
	"bookManagement/logic"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	var cfg config.Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.CreateDb(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = db.CreateSchemas()
	if err != nil {
		log.Fatal(err)
	}

	jwtManager, err := logic.NewJwtManager(24)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	serverManager := presentation.ServerManager{
		Db:         db,
		JwtManager: jwtManager,
	}

	router.Handle(http.MethodPost, "/api/v1/auth/signup", serverManager.SignUp)
	router.Handle(http.MethodPost, "/api/v1/auth/login", serverManager.SignInByCredintials)
	router.Handle(http.MethodPost, "/api/v1/auth/autoLogin", serverManager.SignInByToken)
	router.Handle(http.MethodPost, "/api/v1/books", serverManager.CreateBook)
	router.Handle(http.MethodGet, "/api/v1/books", serverManager.GetAllBooks)
	router.Handle(http.MethodGet, "/api/v1/books/:id", serverManager.GetBook)
	router.Handle(http.MethodPatch, "/api/v1/books/:id", serverManager.UpdateBook)
	router.Handle(http.MethodDelete, "/api/v1/books/:id", serverManager.DeleteBook)
	log.Fatal(router.Run("0.0.0.0:3001"))

}

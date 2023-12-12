package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/zhyzhkoslava/ITEA-GO/hw10/handler"
	"github.com/zhyzhkoslava/ITEA-GO/hw10/middleware"
	"github.com/zhyzhkoslava/ITEA-GO/hw10/server"
	"github.com/zhyzhkoslava/ITEA-GO/hw10/structure"
)

const UsersFile = "hw10/users.json"

func main() {

	usersData, err := os.ReadFile(UsersFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", UsersFile, err)
		return
	}

	var users []structure.User
	if err := json.Unmarshal(usersData, &users); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	logger := log.Default()
	token := "secret"

	requestLoggerMiddleware := middleware.NewRequestLogger(logger)
	tokenMiddleware := middleware.NewTokenMiddleware(token)

	usersHandler := handler.NewUsersHandler(logger, users)
	userHandler := handler.NewUserHandler(logger, users, UsersFile)

	apiServer := server.NewAPIServer(logger)
	apiServer.AddRoute("/users", requestLoggerMiddleware.Wrap(tokenMiddleware.Wrap(usersHandler)))
	apiServer.AddRoute("/user/", requestLoggerMiddleware.Wrap(tokenMiddleware.Wrap(userHandler)))

	if err := apiServer.Start(); err != nil {
		logger.Fatal(err)
	}
}

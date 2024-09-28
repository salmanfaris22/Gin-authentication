package controllers

import (
	"sync"

	"main.go/models"
)

var (
	UserDemo = []models.User{{ID: "1", FirstName: "sdas", LastName: "fai", Email: "salman@gmail.com", Password: "salmana", Token: "DemoToken"}}
	lastId   = 1
	mu       sync.Mutex
)

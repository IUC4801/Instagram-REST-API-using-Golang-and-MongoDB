package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/IUC08/insta/controllers"
)

func main() {
	r := httprouter.New()
	// uc := controllers.NewUserController(getSession())
	uc := controllers.UserController{Session: getSession()}

	r.GET("/users", uc.GetAllUser)
	r.GET("/users/:id", uc.GetUser)
	r.GET("/posts/users/:id", uc.GetUserPost)
	r.POST("/users", uc.CreateUser)
	r.POST("/posts", uc.CreatePost)
	http.ListenAndServe("localhost:8000", r)
}

//mongodb session

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}

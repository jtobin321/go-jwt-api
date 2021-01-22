package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/jtobin321/go-jwt-api/api/config"
	"github.com/jtobin321/go-jwt-api/api/controllers/posts"
	"github.com/jtobin321/go-jwt-api/api/controllers/users"
	"github.com/jtobin321/go-jwt-api/api/controllers/views"
	"github.com/jtobin321/go-jwt-api/api/middleware"
)

// Server handles all API traffic
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Run starts http server with mux router
func (s *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}

// Initialize creates connection to db and new mux router
func (s *Server) Initialize() {
	DB, err := config.InitializeDB()
	if err != nil {
		log.Fatal("Could not connect to database because of the following error:", err)
		return
	}

	s.DB = DB
	s.Router = mux.NewRouter()
	s.initializeRoutes()
}

func (s *Server) initializeRoutes() {
	s.Router.Use(middleware.SetMiddlewareJSON)

	// Home route
	s.Router.HandleFunc("/", views.Home).Methods("GET")

	// // Login route
	s.Router.HandleFunc("/login", s.createRouterFunc(views.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", s.createRouterFunc(users.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", s.createRouterFunc(users.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", s.createRouterFunc(users.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareAuth(s.createRouterFunc(users.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareAuth(s.createRouterFunc(users.DeleteUser))).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", s.createRouterFunc(posts.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", s.createRouterFunc(posts.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", s.createRouterFunc(posts.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareAuth(s.createRouterFunc(posts.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareAuth(s.createRouterFunc(posts.DeletePost))).Methods("DELETE")

}

func (s *Server) createRouterFunc(f func(http.ResponseWriter, *http.Request, *gorm.DB)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, s.DB)
	}
}

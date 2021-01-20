package controllers

import (
	"github.com/jtobin321/go-jwt-api/api/middleware"
)

func (s *Server) initializeRoutes() {
	// Home route
	s.Router.HandleFunc("/", middleware.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login route
	s.Router.HandleFunc("/login", middleware.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middleware.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middleware.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuth(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareAuth(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middleware.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middleware.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuth(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareAuth(s.DeletePost)).Methods("DELETE")
}

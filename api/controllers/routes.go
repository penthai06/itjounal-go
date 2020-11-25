package controllers

import "itjournal/api/middlewares"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")
}

package controllers

import "itjournal/api/middlewares"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")

	server.Router.HandleFunc("/customer", middlewares.SetMiddlewareJSON(server.CustomerSave)).Methods("POST")
	server.Router.HandleFunc("/customer_status_all", middlewares.SetMiddlewareJSON(server.CustomerStatusAll)).Methods("GET")
}

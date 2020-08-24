package controllers

func (s *Server) initializeRoutes() {

	// Home route 
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
}
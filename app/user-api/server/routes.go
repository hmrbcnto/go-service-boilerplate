package server

func (srv *server) initRoutes() {
	srv.mux.HandleFunc("/users", srv.handlers.CreateUser).Methods("POST")
}

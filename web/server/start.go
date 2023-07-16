package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type GracefulShutdown struct {
	ListenAddr  string
	BaseHandler http.Handler
	httpServer  *http.Server
}

func (gs *GracefulShutdown) GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.SkipClean(true)
	router.Handle("/", gs.BaseHandler)
	return router
}

func (gs *GracefulShutdown) Start() {
	router := gs.GetRouter()
	gs.httpServer = &http.Server{
		Addr:    gs.ListenAddr,
		Handler: router,
	}

	fmt.Printf("Server is running at %s\n", gs.ListenAddr)
	if err := gs.httpServer.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

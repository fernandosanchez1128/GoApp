package main

import (
	"GoApp/Config"
	"GoApp/Controllers"
	"GoApp/Middleware"
	"GoApp/Repository"
	"GoApp/Services"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
	"net/http"
	"os"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	p := properties.MustLoadFile("${GOPATH}/src/GoApp/resources/config.properties", properties.UTF8)
	Config.CreateDbConfig(p.Map())
	svc := Services.SortService{}
	errorHandler := Middleware.ErrorHandler {Logger:logger}
	errorMiddleware := httptransport.ServerErrorHandler(errorHandler)
	sortHandler := httptransport.NewServer(
		Middleware.LoggingMiddleware(logger) (Controllers.MakeSortEndpoint(svc)),
		Controllers.DecodeSortRequest,
		Controllers.EncodeResponse,
		errorMiddleware,
	)

	userService := Services.NewUserService(&Repository.UserRepository{})

	userHandler := httptransport.NewServer(
		Middleware.LoggingMiddleware(logger) (Controllers.MakeUserEndpoint(userService)),
		Controllers.DecodeUserRequest,
		Controllers.EncodeResponse,
		errorMiddleware,
	)

	mux := mux.NewRouter().StrictSlash(true)

	mux.Handle("/sort", sortHandler).Methods("POST")
	mux.Handle("/user/{id}", userHandler).Methods("GET")

	http.Handle("/", accessControl(mux))
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

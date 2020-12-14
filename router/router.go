package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/judaro13/miaguila/handlers"
	"github.com/urfave/negroni"
)

//Route test struct
type Route struct {
	name     string
	path     string
	function http.HandlerFunc
	methods  []string
}

//NewRouter function for test
func NewRouter(ctx *models.AppContext) *mux.Router {
	router := mux.NewRouter()

	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			if req.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(rw, req)
		})
	}

	for _, route := range routes() {
		handler := negroni.New(
			negroni.Wrap(route.function),
		)
		methods := append(route.methods, "OPTIONS")
		router.Methods(methods...).Path(route.path).Name(route.name).Handler(handler)
		fmt.Printf("Setting methods %v for route %s\n", methods, route.path)
	}

	router.Use(corsMiddleware)

	return router
}

func routes() []Route {
	return []Route{
		Route{
			path: "/",
			function: func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode(models.JSONResponse{Code: "0", Message: "API Service active v1.0.1", Data: nil})
			},
			methods: []string{"GET"},
		},
		Route{
			path:     "/upload",
			function: handlers.Upload,
			methods:  []string{"POST"},
		},
	}
}
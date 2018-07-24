package main

import (
	"log"
	"net/http"
	"golang.org/x/net/websocket"
	"github.com/ant0ine/go-json-rest/rest"

	"./interfaces/api/comment"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		// rest api
		rest.Get("/monitor/health", func(writer rest.ResponseWriter, request *rest.Request) {
			writer.WriteJson("hello")
		}),
		// websocket
		rest.Get("/ws/rooms/:roomId/comments", func(writer rest.ResponseWriter, request *rest.Request) {
			handler := websocket.Handler(comment.FetchAll)

			handler.ServeHTTP(writer.(http.ResponseWriter), request.Request)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	log.Fatal(http.ListenAndServe(":9999", api.MakeHandler()))
}

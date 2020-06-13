package main

import (
	"fmt"
	"log"
	"net/http"

	"app/apperrors"
	"app/routers"
)

func main() {
	port := "8080"
	fmt.Printf("Web Server Listening on port %s\n", port)

	r := routers.CreateRouter()

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		err = apperrors.HTTPServerPortListenFailed.Wrap(err, "server cannot listen port")
		apperrors.ErrorHandler(err)
		log.Fatal(apperrors.GetType(err), "||", apperrors.GetMessage(err), "||", err)
	}
}

package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

var port = "9999"

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(`cmd`, `/c`, run, uri)
	return cmd.Start()
}

func main() {
	box := packr.NewBox("./templates")

	http.Handle("/", http.FileServer(box))

	err := Open("http://localhost:" + port)
	if err != nil {
		fmt.Printf("")
	}

	fmt.Println("Started server on - localhost:" + port)

	serveError := http.ListenAndServe(":"+port, nil)

	if serveError != nil {
		log.Println("Error starting server", serveError)
	}
}

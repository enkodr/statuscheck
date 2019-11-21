package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"kununu.com/health/config"
	"kununu.com/health/status"
)

const defaultPort = "8008"

var cnf config.Config

func main() {
	// Get command parameters
	var path string
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		dir = "."
	}
	configPath := dir + "/config.yaml"
	flag.StringVar(&path, "c", configPath, "port to listen")
	flag.Parse()
	// Load configuration
	cnf, err = config.Load(path)

	if err != nil {
		log.Printf("No config file defined. Loading from defaults.")
	}

	// Starting server
	http.HandleFunc(cnf.Endpoint, statusCheck)
	log.Printf("Listening on port %s...", cnf.Port)
	err = http.ListenAndServe(":"+cnf.Port, nil)
	if err != nil {
		panic(err)
	}
}

func statusCheck(w http.ResponseWriter, r *http.Request) {
	msg := ""
	chk := status.Make(cnf.Check.Type)
	alive, err := chk.Check(cnf.Check)
	if err != nil {
		log.Printf("ERROR: %s", err)
	}
	if alive {
		msg = cnf.OKMessage
	} else {
		msg = cnf.ErrorMessage
	}
	log.Printf("status check: %s", msg)
	w.Write([]byte(msg))
}

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var test struct {
	Name   string
	code   int
	stdout string
	stderr string
}

func execScript(path string) (code int, stdout string, stderr string, err error) {
	cmd := exec.Command(os.Getenv("SHELL"), path)

	var out bytes.Buffer
	cmd.Stdout = &out

	stderrpipe, err := cmd.StderrPipe()
	if err != nil {
		return -1, "", "", err
	}

	err = cmd.Start()
	if err != nil {
		return -1, "", "", err
	}

	stderrb, err := ioutil.ReadAll(stderrpipe)
	if err != nil {
		return -1, "", "", err
	}
	stderr = string(stderrb)

	err = cmd.Wait()
	if err != nil {
		code, _ = strconv.Atoi(err.Error()[12:])
	}

	return code, out.String(), stderr, nil
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck)
	srv := &http.Server{
		Handler: router,
		Addr:    os.Getenv("ADDRESS") + ":" + os.Getenv("PORT"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println(srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

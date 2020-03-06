package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

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

type Test struct {
	Name   string
	Code   int
	Stdout string
	Stderr string
}

type Response struct {
	Health []Test
	Error  string
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	var res Response
	files, err := filepath.Glob(os.Getenv("SCRIPT_PATH"))
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
	}
	tests := make([]Test, len(files))
	for i, filename := range files {
		code, stdout, stderr, err := execScript(filename)
		tests[i] = Test{
			Name:   filename,
			Code:   code,
			Stdout: stdout,
			Stderr: stderr,
		}
		if err != nil {
			res.Error = err.Error()
		}
	}
	res.Health = tests
	json.NewEncoder(w).Encode(res)
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
	log.Println("listening on: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

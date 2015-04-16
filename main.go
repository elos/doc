package main

import (
	"log"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/elos/hyde"
	"github.com/julienschmidt/httprouter"
)

var secret = "githubsecret"

func main() {
	doc, err := filepath.Abs("documentation")
	if err != nil {
		log.Fatal(err)
	}

	r := httprouter.New()

	r.POST("/push", GithubPush)

	h := hyde.NewHullWithRouter(":3000", r, doc)
	h.Start()
}

func GithubPush(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Print("Heard from github")
	err := exec.Command("rm", "-rf", "documentation").Run()
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("git", "clone", "https://github.com/elos/documentation.git").Run()
	if err != nil {
		log.Fatal(err)
	}
}

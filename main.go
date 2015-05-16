package main

import (
	"log"
	"os/exec"
	"path/filepath"

	"github.com/elos/ehttp/builtin"
	"github.com/elos/ehttp/serve"
	"github.com/elos/hyde"
)

var secret = "githubsecret"

func main() {
	doc, err := filepath.Abs("documentation")
	if err != nil {
		log.Fatal(err)
	}

	agents := filepath.Join(doc, "agents")
	data := filepath.Join(doc, "data")
	http := filepath.Join(doc, "http")
	server := filepath.Join(doc, "server")

	r := builtin.NewRouter()
	r.POST("/push", GithubPush)
	hull := hyde.NewHullWithRouter(r, agents, data, http, server)

	s := serve.New(&serve.Opts{
		Port:    3000,
		Handler: hull,
	})

	go s.Start()
	s.WaitStop()
}

func GithubPush(c *serve.Conn) {
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

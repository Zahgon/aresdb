package main

import (
	"net/http"

	"github.com/abiosoft/ishell"
)

const contentType = "application/json"

type shellContext struct {
	host        string
	port        int
	clusterMode bool
	client      http.Client
}

// script global context
var ctx shellContext

func show(c *ishell.Context) { _ = "STUB: not implemented"; return }

// TODO: @shz implement cluster mode with gateway.controller

// local mode

func aql(c *ishell.Context) { _ = "STUB: not implemented"; return }

func Execute() {
	_ = "STUB: not implemented"

	// ishell shell
	return
}

//TODO: add sql cmd

// cobra command

// read args

// config http client

func main() {
	Execute()
}

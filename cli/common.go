package cli

var helpTextPage = `
            ____________   _____________   _____________  _ _____________
      ___ _/           /  /            /  /   _________/  /             /
          /   ________/ _/    ____    /  /   /_________  /____     ____/
    _ ___/   /_______   /    /   /   /  /________     /      /    /
        /           /  /    /___/   /       _ __/    /     _/    /
       /    _______/  /    ____    /  _________/    /  _ __/    /
     _/    /      __ /    /   /   /  /     _    _ _/      /    /
  __ /____/         /____/   /___/  /_____________/      /____/

  The first open-source Function-as-a-Service (FaaS) platform written in Go.

  Please refer below for available FAST cli.

  - fast help
    This "help" command will show this help page.

  - fast create {your_module_name}
    This "create" command will create a .go file with default code as its content

  - fast build {your_module_name}
    This "build" command will export the .go file into .so file

  - fast rm {your_module_name}
    This "rm" command will remove both the .go and .so files

  - fast start
    This "start" command will start the FAST server
`

var defaultModuleFile = `package main

import (
  "net/http"
  "github.com/labstack/echo"
)

type {module_name} struct{}

func (m *{module_name}) Call(ctx echo.Context) error {
  return ctx.String(http.StatusOK, "Hello world.")
}

var {module_title} {module_name}
`

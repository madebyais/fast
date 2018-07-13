package common

// LogoText is a text logo
var LogoText = `
  	          ____________   _____________   _____________  _ _____________
  	    ___ _/           /  /            /  /   _________/  /             /
  	        /   ________/ _/    ____    /  /   /_________  /____     ____/
  	  _ ___/   /_______   /    /   /   /  /________     /      /    /
  	      /           /  /    /___/   /       _ __/    /     _/    /
  	     /    _______/  /    ____    /  _________/    /  _ __/    /
  	   _/    /      __ /    /   /   /  /     _    _ _/      /    /
	__ /____/         /____/   /___/  /_____________/      /____/
`

// HelpTextPage is the help text content
var HelpTextPage = `
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

// DefaultModuleFile is the default content of module file
var DefaultModuleFile = `package main

import (
  "net/http"
)

type {module_name} struct{}

func (m *{module_name}) Call() (interface{}, error) {
  response := make(map[string]interface{})
  response["code"] = http.StatusOK
  response["data"] = m
  
  return response, nil
}

var {module_title} {module_name}
`

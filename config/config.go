package config

import (
	"flag"
	"log"

	"github.com/madebyais/kilde"
)

// Schema is a schema structure for config package
// It depicts the available config that being used by FAST
type Schema struct {
	Addr string `yaml:"Addr"`

	ModulePath string `yaml:"ModulePath"`
}

var globalconf *Schema

// New is used to initiate config package
func New() {
	s := &Schema{}

	env := flag.String(`env`, `default`, `Set application environment. Default: default`)
	configpath := flag.String(`config`, `/etc/fast.yml`, `Set fast config path. Default: /etc/fast.yml`)
	flag.Parse()

	k := kilde.New()
	k.SetSchema(s)
	k.SetConfigType(`yaml`)
	k.SetFilePath(*configpath)
	k.SetEnv(*env)

	err := k.Read()
	if err != nil {
		log.Fatalf(`Failed to read config. Error=%s`, err.Error())
	}

	globalconf = s
}

// Get returns the current loaded config
func Get() *Schema {
	return globalconf
}

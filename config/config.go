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
	configpath := flag.String(`config`, `/etc/fast.yml`, `Set fast config path. Default: /etc/fast.yml`)
	flag.Parse()

	load(*configpath)
}

func load(configpath string) {
	s := &Schema{}

	k := kilde.New()
	k.SetSchema(s)
	k.SetConfigType(`yaml`)
	k.SetFilePath(configpath)

	err := k.Read()
	if err != nil {
		log.Fatalf(`Failed to read config. Please make sure that you have set config path. Error=%s`, err.Error())
	}

	globalconf = s
}

// Get returns the current loaded config
func Get() *Schema {
	return globalconf
}

// Reload is used to reload application config
func Reload(configpath string) {
	load(configpath)
}

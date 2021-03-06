package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/madebyais/fast/app"
	"github.com/madebyais/fast/common"
	"github.com/madebyais/fast/config"
)

// Interface is an interface for cli package
type Interface interface {
	Init()
}

type schema struct {
	args []string
	cmd  string
}

// New is used to initiate cli package
func New() Interface {
	return &schema{}
}

// Init is used to run the cli package
func (s *schema) Init() {
	s.args = os.Args

	err := s.execCommand()
	if err != nil {
		panic(err)
	}
}

// execCommand will run/execute command based on args
// passed through the FAST cli
func (s *schema) execCommand() error {
	if len(s.args) == 1 {
		return s.startCmd()
	}

	switch s.args[1] {
	default:
		return s.startCmd()
	case `create`:
		return s.createCmd()
	case `build`:
		return s.buildCmd()
	case `rm`:
		return s.removeCmd()
	case `help`:
		return s.helpCmd()
	}
}

// helpCmd will execute help command,
// which will show FAST description and its help page
func (s *schema) helpCmd() error {
	s.cmd = `help`

	fmt.Println(common.HelpTextPage)
	return nil
}

// createCmd will create a file with .go as file extension
// it will contain the default function file
// cmd: fast create {your_module_name}
func (s *schema) createCmd() error {
	s.cmd = `create`

	if len(s.args) < 3 {
		return errors.New(`cannot find module name`)
	}

	moduleName := s.args[2]
	moduleContent := strings.Replace(common.DefaultModuleFile, `{module_name}`, moduleName, -1)

	moduleTitle := strings.ToTitle(moduleName)
	moduleContent = strings.Replace(moduleContent, `{module_title}`, moduleTitle, -1)

	filename := fmt.Sprintf(`%s.go`, moduleName)
	filepath := fmt.Sprintf(`./%s`, filename)
	err := ioutil.WriteFile(filepath, []byte(moduleContent), 0644)
	if err != nil {
		return err
	}

	fmt.Println(`Your`, moduleTitle, `module has been created successfully. Generated filename ~>`, filename)

	return nil
}

// buildCmd will export the .go file into .so file
// which have been created using the create command
// cmd: fast build {your_module_name}
func (s *schema) buildCmd() error {
	s.cmd = `build`

	if len(s.args) < 3 {
		return errors.New(`cannot find module name`)
	}

	moduleName := s.args[2]
	gofilename := fmt.Sprintf(`%s.go`, moduleName)
	sofilename := fmt.Sprintf(`%s.so`, moduleName)

	err := exec.Command(`go`, `build`, `-buildmode=plugin`, `-o`, sofilename, gofilename).Run()
	if err != nil {
		return err
	}

	moduleTitle := strings.ToTitle(moduleName)
	fmt.Println(`Your`, moduleTitle, `module has been build successfully. Generated filename ~>`, sofilename)

	return nil
}

// removeCmd will remove both .go and .so files
// cmd: fast rm {your_module_name}
func (s *schema) removeCmd() error {
	s.cmd = `rm`

	if len(s.args) < 3 {
		return errors.New(`cannot find module name`)
	}

	moduleName := s.args[2]
	gofilename := fmt.Sprintf(`./%s.go`, moduleName)
	sofilename := fmt.Sprintf(`./%s.so`, moduleName)

	err := os.Remove(gofilename)
	if err != nil {
		return err
	}

	err = os.Remove(sofilename)
	if err != nil {
		return err
	}

	return nil
}

// startCmd will start the FAST server
func (s *schema) startCmd() error {
	s.cmd = `start`

	config.New()
	app.New().Start()

	return nil
}

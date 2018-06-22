package cli

import (
	"fmt"
	"os"
	"testing"
)

func tearDown(filename string) {
	err := os.Remove(`./` + filename + `.go`)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.Remove(`./` + filename + `.so`)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestShouldReturnNoErrorWhenExecuteHelpCmd(t *testing.T) {
	args := []string{"fast"}
	c := &schema{
		args: args,
	}

	err := c.helpCmd()
	if err != nil {
		t.Errorf(`Expected no error, but got "%s"`, err.Error())
	}
}

func TestShouldReturnNoErrorWhenExecuteCreateCmd(t *testing.T) {
	filename := `something`
	args := []string{"fast", "create", filename}
	c := &schema{
		args: args,
	}

	err := c.createCmd()
	if err != nil {
		t.Errorf(`Expected no error, but got "%s"`, err.Error())
	}

	tearDown(filename)
}

func TestShouldReturnErrorWhenExecuteCreateCmdWithoutModulename(t *testing.T) {
	args := []string{"fast", "create"}
	c := &schema{
		args: args,
	}

	err := c.createCmd()
	if err == nil {
		t.Errorf(`Expected an error, but got nothing`)
	}
}

func TestShouldReturnNoErrorWhenExecuteBuildCmd(t *testing.T) {
	filename := `something`
	args := []string{"fast", "build", filename}
	c := &schema{
		args: args,
	}

	_ = c.createCmd()
	err := c.buildCmd()
	if err != nil {
		t.Errorf(`Expected no error, but got "%s"`, err.Error())
	}

	tearDown(filename)
}

func TestShouldReturnErrorWhenExecuteBuildCmdWithoutModulename(t *testing.T) {
	args := []string{"fast", "build"}
	c := &schema{
		args: args,
	}

	err := c.buildCmd()
	if err == nil {
		t.Errorf(`Expected an error, but got nothing`)
	}
}

func TestShouldReturnNoErrorWhenExecuteRemoveCmd(t *testing.T) {
	args := []string{"fast", "rm", "something"}
	c := &schema{
		args: args,
	}

	_ = c.createCmd()
	_ = c.buildCmd()

	err := c.removeCmd()
	if err != nil {
		t.Errorf(`Expected no error, but got "%s"`, err.Error())
	}
}

func TestShouldReturnErrorWhenExecuteRemoveCmdWithoutModulename(t *testing.T) {
	args := []string{"fast", "rm"}
	c := &schema{
		args: args,
	}

	err := c.removeCmd()
	if err == nil {
		t.Errorf(`Expected an error, but got nothing`)
	}
}

func TestShouldReturnErrorWhenExecuteRemoveCmdIfGoFileNotFound(t *testing.T) {
	args := []string{"fast", "rm", "something"}
	c := &schema{
		args: args,
	}

	err := c.removeCmd()
	if err == nil {
		t.Errorf(`Expected an error, but got nothing`)
	}
}

func TestShouldReturnErrorWhenExecuteRemoveCmdIfSoFileNotFound(t *testing.T) {
	args := []string{"fast", "rm", "something"}
	c := &schema{
		args: args,
	}

	_ = c.createCmd()

	err := c.removeCmd()
	if err == nil {
		t.Errorf(`Expected an error, but got nothing`)
	}
}

func TestShouldReturnNoErrorWhenExecCommandIsCalledWithoutAnyArguments(t *testing.T) {
	args := []string{"fast"}
	c := &schema{
		args: args,
	}

	err := c.execCommand()
	if err != nil {
		t.Errorf(`Expected no error, but got "%s"`, err.Error())
	}
}

func TestShouldReturnNoErrorAndExecuteDefaultCmdWhenExecCommandIsCalledWithAnyArguments(t *testing.T) {
	args := []string{"fast", "madebyais"}
	c := &schema{
		args: args,
	}

	err := c.execCommand()
	if err != nil {
		t.Errorf(`Expected no error, but got "%s"`, err.Error())
	}
}

func TestShouldReturnNoErrorAndExecuteCreateCmdWhenExecCommandIsCalled(t *testing.T) {
	filename := `something`
	args := []string{"fast", "create", filename}
	c := &schema{
		args: args,
	}

	err := c.execCommand()
	if err != nil {
		t.Errorf(`Expected no error, but got "%s"`, err.Error())
	}

	tearDown(filename)
}

func TestShouldReturnNoErrorAndExecuteBuildCmdWhenExecCommandIsCalled(t *testing.T) {
	filename := `something`
	args := []string{"fast", "build", filename}
	c := &schema{
		args: args,
	}

	_ = c.createCmd()
	err := c.execCommand()
	if err != nil {
		t.Errorf(`Expected no error, but got "%s"`, err.Error())
	}

	tearDown(filename)
}

func TestShouldReturnNoErrorAndExecuteRemoveCmdWhenExecCommandIsCalled(t *testing.T) {
	args := []string{"fast", "rm", "yoo"}
	c := &schema{
		args: args,
	}

	_ = c.createCmd()
	_ = c.buildCmd()

	err := c.execCommand()
	if err != nil {
		t.Errorf(`Expected no error, but got "%s"`, err.Error())
	}
}
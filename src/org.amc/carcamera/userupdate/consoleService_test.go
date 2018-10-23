package userupdate

import "testing"

//TestConsoleService todo mock out log
func TestConsoleService(t *testing.T) {
	c := new(ConsoleService)

	c.Started()
	c.Stopped()
	c.Error("Error msg")
	c.Init()
	c.Close()
}

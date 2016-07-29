package golog

import (
	"testing"
)

func Test_Golog(t *testing.T) {

	log, err := NewLogger("/home/xyq/test.log", "test")
	if err != nil {
		panic(err) // Check for error
	}

	log.Critical("This is Critical!")
	log.Debug("This is Debug!")
	log.Warning("This is Warning!")
	log.Error("This is Error!")
	log.Notice("This is Notice!")
	log.Info("This is Info!")
}

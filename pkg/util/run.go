package util

import (
	"bytes"
	"log"
	"os/exec"
)

func RunAndWait(cmd *exec.Cmd) (string, string, error) {
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	log.Printf("Task %s with args %+v", cmd.Path, cmd.Args)
	err := cmd.Run()
	return outBuf.String(), errBuf.String(), err
}

func CanRun(cmd *exec.Cmd) bool {
	return cmd.Run() == nil
}

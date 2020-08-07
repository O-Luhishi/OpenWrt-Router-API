package common

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	signals = make(chan os.Signal, 100)
)

func RunBash(bashWrapper string) (bool, string, string) {
	cmd := exec.Command("/bin/sh", "-s")
	cmd.Stdin = strings.NewReader(bashWrapper)
	return finishRunning(cmd)
}

func finishRunning(cmd *exec.Cmd) (bool, string, string) {
	stdout, stderr := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	done := make(chan struct{})
	defer close(done)
	go func() {
		for {
			select {
			case <-done:
				return
			case s := <-signals:
				cmd.Process.Signal(s)
			}
		}
	}()
	if err := cmd.Run(); err != nil {
		log.Printf("Error running: %v", err)
		return false, string(stdout.Bytes()), string(stderr.Bytes())
	}
	return true, string(stdout.Bytes()), string(stderr.Bytes())
}

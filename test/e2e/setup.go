package e2e

import (
	"fmt"
	"os/exec"
	"time"
)

func setup() {
	cmd := exec.Command("up", "start")
	cmd.Dir = "../.."
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(300 * time.Millisecond)
}

func apiURL() string {
	return "http://localhost:3000"
}

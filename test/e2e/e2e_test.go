package e2e

import (
	"flag"
	"os"
	"os/exec"
	"testing"
	"time"

	baloo "gopkg.in/h2non/baloo.v3"
)

var start = flag.String("start", "no", "inidicates if the app should be started")

func TestApp(t *testing.T) {
	setup()
	if *start == "yes" {
		cmd := startApp()
		runTests(t)
		endApp(cmd)
	} else {
		runTests(t)
	}
}

func runTests(t *testing.T) {
	t.Run("getPhotos", getPhotos)
}

func startApp() *exec.Cmd {
	cmd := exec.Command("up", "start")
	cmd.Dir = "../.."
	cmd.Env = append(os.Environ(), "UP_TEST=true")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	waitApp()
	return cmd
}

func endApp(cmd *exec.Cmd) {
	if cmd != nil {
		cmd.Process.Kill()
	}
}

func waitApp() {
	test = baloo.New(apiURL())
	_, err := test.Get("/").Send()
	for err != nil {
		time.Sleep(50 * time.Millisecond)
		_, err = test.Get("/").Send()
	}
	return
}

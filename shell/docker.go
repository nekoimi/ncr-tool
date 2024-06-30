package shell

import (
	"errors"
	"fmt"
	"os/exec"
)

func Pull(image string) error {
	return run("docker", "pull", image)
}

func Tag(src, target string) error {
	return run("docker", "tag", src, target)
}

func run(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	stdout, err := cmd.Output()
	if err != nil {
		return errors.New(err.Error())
	}
	fmt.Println(string(stdout))
	return nil
}

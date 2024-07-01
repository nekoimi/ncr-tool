package shell

import (
	"bufio"
	"github.com/nekoimi/ncr-tool/logs"
	"os/exec"
	"sync"
)

var log = logs.New()

func Pull(image string) error {
	return execute("docker", "pull", image)
}

func Tag(src, target string) error {
	log.Infof("\nRenaming %s to %s\n", src, target)
	return execute("docker", "tag", src, target)
}

func RmImage(image string) error {
	return execute("docker", "rmi", image)
}

func execute(command string, args ...string) error {
	var wg sync.WaitGroup
	cmd := exec.Command(command, args...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	wg.Add(1)
	scanstd := bufio.NewScanner(stdout)
	go func() {
		for scanstd.Scan() {
			log.Infoln(scanstd.Text())
		}
		wg.Done()
	}()

	wg.Add(1)
	scanerr := bufio.NewScanner(stderr)
	go func() {
		for scanerr.Scan() {
			log.Errorln(scanerr.Text())
		}
		wg.Done()
	}()

	if err := cmd.Start(); err != nil {
		log.Fatalln(err.Error())
		return err
	}

	return cmd.Wait()
}

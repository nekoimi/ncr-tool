package shell

import (
	"github.com/nekoimi/ncr-tool/logs"
	"testing"
)

func TestExecCmd(t *testing.T) {
	logs.New()
	execute("docker", "pull", "nginx:latest")
}

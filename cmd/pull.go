package cmd

import (
	"errors"
	"github.com/nekoimi/ncr-tool/shell"
	"github.com/urfave/cli/v2"
	"log"
	"strings"
)

var registryMap = make(map[string]string)

var pull = &cli.Command{
	Name:   "pull",
	Usage:  "Pull a container image",
	Action: run,
}

func init() {
	registryMap["docker.io"] = "docker.mirror.403forbidden.run"
	registryMap["quay.io"] = "quay.mirror.403forbidden.run"
	registryMap["ghcr.io"] = "ghcr.mirror.403forbidden.run"
	registryMap["gcr.io"] = "gcr.mirror.403forbidden.run"
	registryMap["k8s.gcr.io"] = "k8sgcr.mirror.403forbidden.run"

	// registry pull command
	root.Commands = append(root.Commands, pull)
}

func run(ctx *cli.Context) error {
	image := ctx.Args().First()
	if len(image) == 0 {
		return errors.New("缺少容器镜像参数")
	}

	var targetImage string
	for k, v := range registryMap {
		if strings.HasPrefix(image, k) {
			targetImage = strings.ReplaceAll(image, k, v)
		}
	}

	if len(targetImage) == 0 {
		if strings.Count(image, "/") == 0 {
			targetImage = "docker.mirror.403forbidden.run/library/" + image
		} else {
			targetImage = "docker.mirror.403forbidden.run/" + image
		}
	}

	log.Printf("Image: %s ===> %s \n", image, targetImage)

	err := shell.Pull(targetImage)
	if err != nil {
		return err
	}

	err = shell.Tag(targetImage, image)
	if err != nil {
		return err
	}

	return nil
}

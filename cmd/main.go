package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/arashn0uri/go-server/internal/api"
)

func main() {
	app := api.LoadApplication()
	h, mountErr := app.Mount()
	if mountErr != nil {
		logrus.Error("failed to mount application", "error", mountErr)
		os.Exit(1)
	}

	err := app.Run(h)

	if err != nil {
		logrus.Error("server has faild to start", "error", err)
		os.Exit(1)
	}
}

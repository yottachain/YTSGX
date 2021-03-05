package main

import (
	"flag"
	"os/exec"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yottachain/YTSGX.git/routers"
)

func main() {
	logrus.Infof(time.Now().Format("2006-01-02 15:04:05") + "strart ......")
	flag.Parse()

	// api.StartApi()

	router := routers.InitRouter()

	err1 := router.Run(":8088")
	if err1 != nil {
		panic(err1)
	}
	logrus.Info("strart ......")

}

var commands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func OpenUrl(uri string) {
	run, _ := commands[runtime.GOOS]
	exec.Command(run, uri).Start()
}

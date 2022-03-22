package main

import (
	"flag"
	"os/exec"
	"runtime"
	"time"

	"github.com/yottachain/YTSGX/tools"

	"github.com/sirupsen/logrus"
	"github.com/yottachain/YTSGX/routers"
)

func main() {

	//gin.SetMode(gin.DebugMode)
	logrus.Infof(time.Now().Format("2006-01-02 15:04:05") + "strart ......")
	flag.Parse()

	// api.StartApi()

	tools.CreateStorageDirectory()

	router := routers.InitRouter()

	err1 := router.Run(":18080")
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

// func run() error {
// 	var values yts3Flags

// 	flagSet := flag.NewFlagSet("", 0)
// 	values.attach(flagSet)
// 	values.backendKind = "mem"
// 	values.initialBucket = "bucket"
// 	values.fsPath = "test"

// 	if err := flagSet.Parse(os.Args[1:]); err != nil {
// 		return err
// 	}

// 	stopper, err := profile(values)
// 	if err != nil {
// 		return err
// 	}
// 	defer stopper()

// 	if values.debugHost != "" {
// 		go debugServer(values.debugHost)
// 	}

// 	var backend yts3.Backend

// 	timeSource, timeSkewLimit, err := values.timeOptions()
// 	if err != nil {
// 		return err
// 	}

// 	switch values.backendKind {
// 	case "":
// 		flag.PrintDefaults()
// 		return fmt.Errorf("-backend is required")

// 	case "mem", "memory":
// 		if values.initialBucket == "" {
// 			log.Println("no buckets available; consider passing -initialbucket")
// 		}
// 		backend = s3mem.New(s3mem.WithTimeSource(timeSource))
// 		log.Println("using memory backend")

// 	default:
// 		return fmt.Errorf("unknown backend %q", values.backendKind)
// 	}

// 	if values.initialBucket != "" {
// 	}

// 	faker := yts3.New(backend,
// 		yts3.WithIntegrityCheck(!values.noIntegrity),
// 		yts3.WithTimeSkewLimit(timeSkewLimit),
// 		yts3.WithTimeSource(timeSource),
// 		yts3.WithLogger(yts3.GlobalLog()),
// 		yts3.WithHostBucket(values.hostBucket),
// 	)

// 	return listenAndServe(values.host, faker.Server())
// }

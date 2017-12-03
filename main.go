package main

import "flag"
import (
	"fmt"

	"os"

	"net/http"

	log "github.com/Sirupsen/logrus"
)

type deviceFilesFlag []string

var attachArg bool
var detachArg bool
var deviceFiles deviceFilesFlag
var hostname string
var serverMode bool
var vmName string

func (deviceFiles *deviceFilesFlag) String() string {
	return fmt.Sprint(*deviceFiles)
}

func (deviceFiles *deviceFilesFlag) Set(deviceFile string) error {
	*deviceFiles = append(*deviceFiles, deviceFile)
	return nil
}

func init() {
	log.SetLevel(log.DebugLevel)

	flag.BoolVar(&serverMode, "s", false, "Run as a server")
	flag.BoolVar(&serverMode, "server", false, "Run as a server")
	flag.Var(&deviceFiles, "f", "Path to device file")
	flag.Var(&deviceFiles, "device-file", "Path to device file")
	flag.StringVar(&vmName, "n", "", "Name of the VM to manage")
	flag.StringVar(&vmName, "name", "", "Name of the VM to manage")
	flag.StringVar(&hostname, "h", "127.0.0.1:7654", "Hostname of the service to connect to")
	flag.StringVar(&hostname, "hostname", "127.0.0.1:7654", "Hostname of the service to connect to")
	flag.BoolVar(&attachArg, "a", false, "Attach the devices")
	flag.BoolVar(&attachArg, "attach", false, "Attach the devices")
	flag.BoolVar(&detachArg, "d", false, "Detach the devices")
	flag.BoolVar(&detachArg, "detach", false, "Detach the devices")
}

func main() {
	flag.Parse()

	// Two run methods:
	// 1) As server daemon that listens for commands via REST api
	// 2) As client that sends command based on CLI args

	if serverMode {
		if vmName == "" {
			log.Error("VM name must be provided when starting in server mode")
			os.Exit(1)
		}

		log.Info("Starting in server mode...")
		startServer()
	} else {
		if !attachArg && !detachArg {
			log.Error("Must specify whether to attach or detach the devices")
			os.Exit(1)
		}

		if attachArg {
			http.Get(fmt.Sprintf("http://%s/virsh-device-daemon/attach", hostname))
		}

		if detachArg {
			http.Get(fmt.Sprintf("http://%s/virsh-device-daemon/detach", hostname))
		}
	}
}

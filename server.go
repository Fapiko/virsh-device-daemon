package main

import (
	"net/http"

	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
)

func startServer() {
	ws := new(restful.WebService)
	ws.Path("/virsh-device-daemon").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/attach").To(attach))
	ws.Route(ws.GET("/detach").To(detach))

	restful.Add(ws)
	http.ListenAndServe("0.0.0.0:7654", nil)
}

func detach(request *restful.Request, response *restful.Response) {
	for _, filename := range deviceFiles {
		cmd := exec.Command("virsh", "detach-device", vmName, filename)

		log.Debug(cmd.Args)

		err := cmd.Run()
		if err != nil {
			log.Error(err)
			response.WriteErrorString(500, err.Error())
		}
	}
}

func attach(request *restful.Request, response *restful.Response) {
	for _, filename := range deviceFiles {
		cmd := exec.Command("virsh", "attach-device", vmName, filename)

		log.Debug(cmd.Args)

		err := cmd.Run()
		if err != nil {
			log.Error(err)
			response.WriteErrorString(500, err.Error())
		}
	}
}

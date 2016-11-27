package handler

import (
	"encoding/json"
	"github.com/andyxning/host_ip_reflection/models"
	"github.com/docker/distribution/health"
	_ "github.com/docker/distribution/health/api"
	"github.com/golang/glog"
	"net"
	"net/http"
)

func getRemoteIP(w http.ResponseWriter, r *http.Request) {
	var node models.Node

	if remoteIP, exists := r.Header["X-Real-Ip"]; exists {
		node = models.Node{IP: remoteIP[0]}
	} else {
		remoteAddr := r.RemoteAddr
		remoteIP, _, err := net.SplitHostPort(remoteAddr)
		if err != nil {
			glog.Errorf("parse remote address error. %s", err)
			http.Error(w, "can not get remote ip", http.StatusInternalServerError)
			return
		}

		node = models.Node{IP: remoteIP}
	}

	content, err := json.Marshal(node)
	if err != nil {
		glog.Errorf("error in encode response body. %s", err)
		http.Error(w, "error in encode response body", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)

	glog.Infof("remote address %s", node.IP)
}

func init() {
	http.Handle("/", health.Handler(http.HandlerFunc(getRemoteIP)))
}

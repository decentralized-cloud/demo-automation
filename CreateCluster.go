// Package demoautomation contains functions required to automate edge cloud resource provisioning
package demoautomation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/decentralized-cloud/demo-automation/services/configuration"
	"github.com/decentralized-cloud/demo-automation/services/gce/gke"
)

var gkeService gke.GkeContract

func init() {
	configurationService, err := configuration.NewEnvConfigurationService()
	if err != nil {
		log.Fatalf("Failed to create configuration service. Error: %v\n", err)
	}

	gkeService, err = gke.NewGkeService(configurationService)
	if err != nil {
		log.Fatalf("Failed to create gkeService service. Error: %v\n", err)
	}
}

// CreateCluster creates a new GKE cluster
func CreateCluster(w http.ResponseWriter, r *http.Request) {
	request := gke.CreateClusterRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to decode the provided request. Error: %v\n", err)

		return
	}

	_, err := gkeService.CreateCluster(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to create cluster. Error: %v\n", err)
	}

	w.WriteHeader(http.StatusAccepted)
}

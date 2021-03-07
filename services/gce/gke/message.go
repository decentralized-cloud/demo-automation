// Package gke implements helper methods to create and list GKE cluster
package gke

// CreateClusterRequest contains the request to create a GKE cluster
type CreateClusterRequest struct {
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	ClusterLabels map[string]string `json:"clusterLabels"`
	NodeLabels    map[string]string `json:"nodeLabels"`
}

// CreateClusterResponse contains the response of creating a GKE cluster
type CreateClusterResponse struct {
}

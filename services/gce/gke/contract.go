// Package gke implements helper methods to create and list GKE cluster
package gke

import "context"

// GkeContract declares the service that provides helper methods required by different
// demo automation functions to create and list GKE cluster
type GkeContract interface {
	// CreateCluster creates a new GKE cluster using provided
	// context: Mandatory The reference to the context
	// request: Mandatory. The request to create a GKE cluster
	// Returns either the result of creating new GKE cluster or error if something goes wrong.
	CreateCluster(
		ctx context.Context,
		request *CreateClusterRequest) (*CreateClusterResponse, error)
}

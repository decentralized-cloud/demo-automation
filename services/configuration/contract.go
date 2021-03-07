// Package configuration implements configuration service required by different demo automation functions
package configuration

// ConfigurationContract declares the service that provides configuration required by different demo automation functions
type ConfigurationContract interface {
	// GetProjectId returns the Google Cloud project ID
	// Returns the Google Cloud project ID or error if something goes wrong
	GetProjectID() (string, error)

	// GetRegion returns the Google Cloud region to deploy resources to
	// Returns the Google Cloud region to deploy resources to or error if something goes wrong
	GetRegion() (string, error)

	// GetZone returns the Google Cloud zone to deploy resources to
	// Returns the Google Cloud zone to deploy resources to or error if something goes wrong
	GetZone() (string, error)

	// GetKubernetesClusterVersion returns the Kubernetes version to use when creating GKE cluster
	// Returns the Kubernetes version to use when creating GKE cluster or error if something goes wrong
	GetKubernetesClusterVersion() (string, error)

	// GetMachineType returns the name of a Google Compute Engine machine type
	// Returns the name of a Google Compute Engine machine type or error if something goes wrong
	GetMachineType() (string, error)

	// GetDiskSize returns the size of the disk attached to each node
	// Returns the size of the disk attached to each node or error if something goes wrong
	GetDiskSize() (int, error)

	// GetDiskType returns the type of the disk attached to each node (e.g. 'pd-standard', 'pd-ssd' or 'pd-balanced')
	// Returns the type of the disk attached to each node or error if something goes wrong
	GetDiskType() (string, error)

	// GetImageType returns the image type to use for this node. Note that for a given image type,
	// the latest version of it will be used.
	// Returns the image type to use for this node or error if something goes wrong
	GetImageType() (string, error)

	// GetNodeCount returns the node count for the pool
	// Returns the node count for the pool or error if something goes wrong
	GetNodeCount() (int, error)

	// GetMaxPodsPerNode returns the constraint enforced on the max num of pods per node
	// Returns the constraint enforced on the max num of pods per nodel or error if something goes wrong
	GetMaxPodsPerNode() (int, error)
}

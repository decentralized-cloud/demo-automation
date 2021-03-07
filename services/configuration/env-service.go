// Package configuration implements configuration service required by different demo automation functions
package configuration

import (
	"os"
	"strconv"
	"strings"
)

type envConfigurationService struct {
}

// NewEnvConfigurationService creates new instance of the EnvConfigurationService, setting up all dependencies and returns the instance
// Returns the new service or error if something goes wrong
func NewEnvConfigurationService() (ConfigurationContract, error) {
	return &envConfigurationService{}, nil
}

// GetProjectId returns the Google Cloud project ID
// Returns the Google Cloud project ID or error if something goes wrong
func (service *envConfigurationService) GetProjectID() (string, error) {
	value := os.Getenv("PROJECT_ID")

	if strings.Trim(value, " ") == "" {
		return "", NewUnknownError("PROJECT_ID is required")
	}

	return value, nil
}

// GetRegion returns the Google Cloud region to deploy resources to
// Returns the Google Cloud region to deploy resources to or error if something goes wrong
func (service *envConfigurationService) GetRegion() (string, error) {
	value := os.Getenv("REGION")

	if strings.Trim(value, " ") == "" {
		return "australia-southeast1", nil
	}

	return value, nil
}

// GetZone returns the Google Cloud zone to deploy resources to
// Returns the Google Cloud zone to deploy resources to or error if something goes wrong
func (service *envConfigurationService) GetZone() (string, error) {
	value := os.Getenv("ZONE")

	if strings.Trim(value, " ") == "" {
		return "australia-southeast1-a", nil
	}

	return value, nil
}

// GetKubernetesClusterVersion returns the Kubernetes version to use when creating GKE cluster
// Returns the Kubernetes version to use when creating GKE cluster or error if something goes wrong
func (service *envConfigurationService) GetKubernetesClusterVersion() (string, error) {
	value := os.Getenv("KUBERNETES_CLUSTER_VERSION")

	if strings.Trim(value, " ") == "" {
		return "latest", nil
	}

	return value, nil
}

// GetMachineType returns the name of a Google Compute Engine machine type
// Returns the name of a Google Compute Engine machine type or error if something goes wrong
func (service *envConfigurationService) GetMachineType() (string, error) {
	value := os.Getenv("MACHINE_TYPE")

	if strings.Trim(value, " ") == "" {
		return "e2-medium", nil
	}

	return value, nil
}

// GetDiskSize returns the size of the disk attached to each node
// Returns the size of the disk attached to each node or error if something goes wrong
func (service *envConfigurationService) GetDiskSize() (int, error) {
	valueStr := os.Getenv("DISK_SIZE")
	if strings.Trim(valueStr, " ") == "" {
		return 32, nil
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, NewUnknownErrorWithError("Failed to convert DISK_SIZE to integer", err)
	}

	return value, nil
}

// GetDiskType returns the type of the disk attached to each node (e.g. 'pd-standard', 'pd-ssd' or 'pd-balanced')
// Returns the type of the disk attached to each node or error if something goes wrong
func (service *envConfigurationService) GetDiskType() (string, error) {
	value := os.Getenv("DISK_TYPE")

	if strings.Trim(value, " ") == "" {
		return "pd-standard", nil
	}

	return value, nil
}

// GetImageType returns the image type to use for this node. Note that for a given image type,
// the latest version of it will be used.
// Returns the image type to use for this node or error if something goes wrong
func (service *envConfigurationService) GetImageType() (string, error) {
	value := os.Getenv("IMAGE_TYPE")

	if strings.Trim(value, " ") == "" {
		return "COS", nil
	}

	return value, nil
}

// GetNodeCount returns the node count for the pool
// Returns the node count for the pool or error if something goes wrong
func (service *envConfigurationService) GetNodeCount() (int, error) {
	valueStr := os.Getenv("NODE_COUNT")
	if strings.Trim(valueStr, " ") == "" {
		return 1, nil
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, NewUnknownErrorWithError("Failed to convert NODE_COUNT to integer", err)
	}

	return value, nil
}

// GetMaxPodsPerNode returns the constraint enforced on the max num of pods per node
// Returns the constraint enforced on the max num of pods per nodel or error if something goes wrong
func (service *envConfigurationService) GetMaxPodsPerNode() (int, error) {
	valueStr := os.Getenv("MAX_PODS_PER_NODE")
	if strings.Trim(valueStr, " ") == "" {
		return 110, nil
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, NewUnknownErrorWithError("Failed to convert MAX_PODS_PER_NODE to integer", err)
	}

	return value, nil
}

// Package gke implements helper methods to create and list GKE cluster
package gke

import (
	"context"
	"log"

	container "cloud.google.com/go/container/apiv1"
	"github.com/decentralized-cloud/demo-automation/services/configuration"
	commonErrors "github.com/micro-business/go-core/system/errors"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

type gkeService struct {
	projectID                string
	region                   string
	zone                     string
	kubernetesClusterVersion string
	machineType              string
	diskSize                 int
	diskType                 string
	imageType                string
	nodeCount                int
	maxPodsPerNode           int
}

// NewGkeService creates new instance of the GkeService, setting up all dependencies and returns the instance
// configurationService: Mandatory. Reference to the configuration service
// Returns the new service or error if something goes wrong
func NewGkeService(
	configurationService configuration.ConfigurationContract) (GkeContract, error) {
	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	projectID, err := configurationService.GetProjectID()
	if err != nil {
		return nil, NewUnknownErrorWithError("Could not read ProjectID", err)
	}

	region, err := configurationService.GetRegion()
	if err != nil {
		return nil, NewUnknownErrorWithError("Could not read Region", err)
	}

	zone, err := configurationService.GetZone()
	if err != nil {
		return nil, NewUnknownErrorWithError("Could not read Zone", err)
	}

	kubernetesClusterVersion, err := configurationService.GetKubernetesClusterVersion()
	if err != nil {
		return nil, NewUnknownErrorWithError("Could not read KubernetesClusterVersion", err)
	}

	machineType, err := configurationService.GetMachineType()
	if err != nil {
		return nil, NewUnknownErrorWithError("Could not read MachineType", err)
	}

	diskSize, err := configurationService.GetDiskSize()
	if err != nil {
		return nil, NewUnknownErrorWithError("Could not read DiskSize", err)
	}

	diskType, err := configurationService.GetDiskType()
	if err != nil {
		return nil, NewUnknownErrorWithError("Could not read DiskType", err)
	}

	imageType, err := configurationService.GetImageType()
	if err != nil {
		return nil, NewUnknownErrorWithError("Could not read ImageType", err)
	}

	nodeCount, err := configurationService.GetNodeCount()
	if err != nil {
		return nil, NewUnknownErrorWithError("Could not read NodeCount", err)
	}

	maxPodsPerNode, err := configurationService.GetMaxPodsPerNode()
	if err != nil {
		return nil, NewUnknownErrorWithError("Could not read MaxPodsPerNode", err)
	}

	return &gkeService{
		projectID:                projectID,
		region:                   region,
		zone:                     zone,
		kubernetesClusterVersion: kubernetesClusterVersion,
		machineType:              machineType,
		diskSize:                 diskSize,
		diskType:                 diskType,
		imageType:                imageType,
		nodeCount:                nodeCount,
		maxPodsPerNode:           maxPodsPerNode,
	}, nil
}

// CreateCluster creates a new GKE cluster using provided
// context: Mandatory The reference to the context
// request: Mandatory. The request to create a GKE cluster
// Returns either the result of creating new GKE cluster or error if something goes wrong.
func (service *gkeService) CreateCluster(
	ctx context.Context,
	request *CreateClusterRequest) (*CreateClusterResponse, error) {
	if ctx == nil {
		return nil, commonErrors.NewArgumentNilError("ctx", "ctx is required")
	}

	if request == nil {
		return nil, commonErrors.NewArgumentNilError("request", "request is required")
	}

	if err := request.Validate(); err != nil {
		return nil, NewUnknownErrorWithError("Request validation failed", err)
	}

	clusterManager, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return nil, NewUnknownErrorWithError("Failed to create cluster manager", err)
	}

	createClusterRequest := &containerpb.CreateClusterRequest{
		Cluster: &containerpb.Cluster{
			Name:        request.Name,
			Description: request.Description,
			MasterAuth: &containerpb.MasterAuth{
				ClientCertificateConfig: &containerpb.ClientCertificateConfig{
					IssueClientCertificate: false,
				},
			},
			Network: "projects/" + service.projectID + "/global/networks/default",
			AddonsConfig: &containerpb.AddonsConfig{
				HttpLoadBalancing: &containerpb.HttpLoadBalancing{
					Disabled: false,
				},
				HorizontalPodAutoscaling: &containerpb.HorizontalPodAutoscaling{
					Disabled: false,
				},
				DnsCacheConfig: &containerpb.DnsCacheConfig{
					Enabled: true,
				},
			},
			Subnetwork: "projects/" + service.projectID + "/regions/" + service.region + "/subnetworks/default",
			NodePools: []*containerpb.NodePool{
				{
					Name: "default-pool",
					Config: &containerpb.NodeConfig{
						MachineType: service.machineType,
						DiskSizeGb:  int32(service.diskSize),
						OauthScopes: []string{
							"https://www.googleapis.com/auth/devstorage.read_only",
							"https://www.googleapis.com/auth/logging.write",
							"https://www.googleapis.com/auth/monitoring",
							"https://www.googleapis.com/auth/servicecontrol",
							"https://www.googleapis.com/auth/service.management.readonly",
							"https://www.googleapis.com/auth/trace.append",
						},
						Metadata:    map[string]string{"disable-legacy-endpoints": "true"},
						ImageType:   service.imageType,
						Labels:      request.NodeLabels,
						Preemptible: true,
						DiskType:    service.diskType,
						ShieldedInstanceConfig: &containerpb.ShieldedInstanceConfig{
							EnableIntegrityMonitoring: false,
						},
					},
					InitialNodeCount: int32(service.nodeCount),
					Autoscaling: &containerpb.NodePoolAutoscaling{
						Enabled: false,
					},
					Management: &containerpb.NodeManagement{
						AutoUpgrade: true,
						AutoRepair:  true,
					},
					UpgradeSettings: &containerpb.NodePool_UpgradeSettings{
						MaxSurge: 1,
					},
				},
			},
			Locations:      []string{service.zone},
			ResourceLabels: request.ClusterLabels,
			NetworkPolicy: &containerpb.NetworkPolicy{
				Enabled: false,
			},
			IpAllocationPolicy: &containerpb.IPAllocationPolicy{
				UseIpAliases: true,
			},
			MasterAuthorizedNetworksConfig: &containerpb.MasterAuthorizedNetworksConfig{
				Enabled: false,
			},
			Autoscaling: &containerpb.ClusterAutoscaling{
				EnableNodeAutoprovisioning: false,
			},
			DefaultMaxPodsConstraint: &containerpb.MaxPodsConstraint{
				MaxPodsPerNode: int64(service.maxPodsPerNode),
			},
			AuthenticatorGroupsConfig: &containerpb.AuthenticatorGroupsConfig{
				Enabled: false,
			},
			PrivateClusterConfig: &containerpb.PrivateClusterConfig{
				EnablePrivateNodes:  true,
				MasterIpv4CidrBlock: "172.16.0.0/28",
			},
			DatabaseEncryption: &containerpb.DatabaseEncryption{
				State: containerpb.DatabaseEncryption_DECRYPTED,
			},
			ShieldedNodes: &containerpb.ShieldedNodes{
				Enabled: false,
			},
			ReleaseChannel: &containerpb.ReleaseChannel{
				Channel: containerpb.ReleaseChannel_REGULAR,
			},
			InitialClusterVersion: service.kubernetesClusterVersion,
			Location:              service.zone,
		},
		Parent: "projects/" + service.projectID + "/locations/" + service.zone,
	}

	operation, err := clusterManager.CreateCluster(ctx, createClusterRequest)
	if err != nil {
		return nil, NewUnknownErrorWithError("Failed to create cluster manager", err)
	}

	log.Printf("Operation result: %v\n", operation)

	return &CreateClusterResponse{}, nil
}

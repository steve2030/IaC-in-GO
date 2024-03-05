package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "cloud.google.com/go/container/apiv1"
    "google.golang.org/api/option"
    "google.golang.org/grpc"
    containerpb "google.golang.org/genproto/googleapis/container/v1"
)

const (
    projectID   = "uzapoint"
    clusterName = "gol-cluster"
    region      = "us-central1"
    zone        = "us-central1-a"
)

func main() {
    ctx := context.Background()

    // Creates a client.
    client, err := container.NewClusterManagerClient(ctx, option.WithGRPCDialOption(grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*4))))
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    // Create cluster request.
    req := &containerpb.CreateClusterRequest{
        Parent:    fmt.Sprintf("projects/%s/locations/%s", projectID, region),
        Cluster:   createCluster(),
        ClusterId: clusterName, // Corrected line
        Zone:      zone,
    }

    // Create the cluster.
    resp, err := client.CreateCluster(ctx, req)
    if err != nil {
        log.Fatalf("Failed to create cluster: %v", err)
    }

    fmt.Printf("Cluster created: %v\n", resp)

    // Additional setup for namespaces can be added here.
}

func createCluster() *containerpb.Cluster {
    return &containerpb.Cluster{
        Name:            clusterName,
        NodePools: []*containerpb.NodePool{
            {
                Name: "default-pool",
                Config: &containerpb.NodeConfig{
                    MachineType: "e2-medium",
                    DiskSizeGb:  100,
                    DiskType:    "pd-standard",
                    ImageType:   "COS",
                    OauthScopes: []string{
                        "https://www.googleapis.com/auth/compute",
                        "https://www.googleapis.com/auth/devstorage.read_only",
                        "https://www.googleapis.com/auth/logging.write",
                        "https://www.googleapis.com/auth/monitoring",
                        "https://www.googleapis.com/auth/service.management.readonly",
                        "https://www.googleapis.com/auth/servicecontrol",
                        "https://www.googleapis.com/auth/trace.append",
                    },
                    Labels: map[string]string{
                        "env": "production",
                    },
                },
                Autoscaling: &containerpb.NodePoolAutoscaling{
                    Enabled:       true,
                    MinNodeCount:  1,
                    MaxNodeCount:  5,
                },
                Management: &containerpb.NodeManagement{
                    AutoUpgrade: true,
                    AutoRepair:  true,
                },
                InitialNodeCount: 1,
            },
        },
        NetworkConfig: &containerpb.NetworkConfig{
            Network: "default",
        },
        MasterAuth: &containerpb.MasterAuth{
            Username: "admin",
        },
        LoggingService:   "logging.googleapis.com",
        MonitoringService: "monitoring.googleapis.com",
        AddonsConfig: &containerpb.AddonsConfig{
            HorizontalPodAutoscaling: &containerpb.HorizontalPodAutoscaling{
                Disabled: false,
            },
            HttpLoadBalancing: &containerpb.HttpLoadBalancing{
                Disabled: false,
            },
            KubernetesDashboard: &containerpb.KubernetesDashboard{
                Disabled: false,
            },
        },
        EnableKubernetesAlpha: false,
        ResourceLabels: map[string]string{
            "env": "production",
        },
        Locations: []string{
            zone,
        },
    }
}

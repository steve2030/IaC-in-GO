# GCP Autopilot Cluster Provisioning

This Go script provisions an Autopilot cluster in Google Cloud Platform (GCP) with two namespaces: staging and production.

## Prerequisites

Before running the script, make sure you have the following:

- Go installed on your system.
- Authentication set up for Google Cloud Platform.
- Google Cloud SDK installed and configured.

## Usage

1. Clone the repository:

    ```bash
    git clone <repository-url>
    ```

2. Navigate to the directory containing the script.

3. Run the script using the following command:

    ```bash
    go run index.go
    ```

4. Follow the prompts to provide necessary inputs.

## Configuration

Make sure to update the following variables in the script with your actual project details and configurations:

- `projectID`: Your GCP project ID.
- `clusterName`: Name for your Autopilot cluster.
- `region`: Region where the cluster will be deployed.
- `zone`: Zone where the cluster will be deployed.

## Dependencies

- cloud.google.com/go/container/apiv1
- google.golang.org/api/option
- google.golang.org/genproto/googleapis/container/v1
- google.golang.org/grpc

## License

This project is licensed under the [MIT License](LICENSE).


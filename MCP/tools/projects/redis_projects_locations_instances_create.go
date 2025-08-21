package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/google-cloud-memorystore-for-redis-api/mcp-server/config"
	"github.com/google-cloud-memorystore-for-redis-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Redis_projects_locations_instances_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		parentVal, ok := args["parent"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: parent"), nil
		}
		parent, ok := parentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: parent"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["instanceId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("instanceId=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Instance
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v1/%s/instances%s", cfg.BaseURL, parent, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Operation
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateRedis_projects_locations_instances_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_parent_instances",
		mcp.WithDescription("Creates a Redis instance based on the specified tier and memory size. By default, the instance is accessible from the project's [default network](https://cloud.google.com/vpc/docs/vpc). The creation is executed asynchronously and callers may check the returned operation to track its progress. Once the operation is completed the Redis instance will be fully functional. Completed longrunning.Operation will contain the new instance object in the response field. The returned operation is automatically deleted after a few hours, so there is no need to call DeleteOperation."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. The resource name of the instance location using the form: `projects/{project_id}/locations/{location_id}` where `location_id` refers to a GCP region.")),
		mcp.WithString("instanceId", mcp.Description("Required. The logical name of the Redis instance in the customer project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-40 characters. * Must end with a number or a letter. * Must be unique within the customer project / location")),
		mcp.WithString("customerManagedKey", mcp.Description("Input parameter: Optional. The KMS key reference that the customer provides when trying to create the instance.")),
		mcp.WithString("name", mcp.Description("Input parameter: Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.")),
		mcp.WithString("maintenanceVersion", mcp.Description("Input parameter: Optional. The self service update maintenance version. The version is date based such as \"20210712_00_00\".")),
		mcp.WithString("persistenceIamIdentity", mcp.Description("Input parameter: Output only. Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is \"serviceAccount:\". The value may change over time for a given instance so should be checked before each import/export operation.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. The time the instance was created.")),
		mcp.WithArray("serverCaCerts", mcp.Description("Input parameter: Output only. List of server CA certificates for the instance.")),
		mcp.WithNumber("port", mcp.Description("Input parameter: Output only. The port number of the exposed Redis endpoint.")),
		mcp.WithBoolean("satisfiesPzs", mcp.Description("Input parameter: Optional. Output only. Reserved for future use.")),
		mcp.WithBoolean("authEnabled", mcp.Description("Input parameter: Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to \"true\" AUTH is enabled on the instance. Default value is \"false\" meaning AUTH is disabled.")),
		mcp.WithObject("labels", mcp.Description("Input parameter: Resource labels to represent user provided metadata")),
		mcp.WithNumber("replicaCount", mcp.Description("Input parameter: Optional. The number of replica nodes. The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1. The valid value for basic tier is 0 and the default is also 0.")),
		mcp.WithObject("redisConfigs", mcp.Description("Input parameter: Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries")),
		mcp.WithArray("nodes", mcp.Description("Input parameter: Output only. Info per node.")),
		mcp.WithString("alternativeLocationId", mcp.Description("Input parameter: Optional. If specified, at least one node will be provisioned in this zone in addition to the zone specified in location_id. Only applicable to standard tier. If provided, it must be a different zone from the one provided in [location_id]. Additional nodes beyond the first 2 will be placed in zones selected by the service.")),
		mcp.WithString("redisVersion", mcp.Description("Input parameter: Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility * `REDIS_7_0` for Redis 7.0 compatibility")),
		mcp.WithString("transitEncryptionMode", mcp.Description("Input parameter: Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.")),
		mcp.WithArray("availableMaintenanceVersions", mcp.Description("Input parameter: Optional. The available maintenance versions that an instance could update to.")),
		mcp.WithNumber("readEndpointPort", mcp.Description("Input parameter: Output only. The port number of the exposed readonly redis endpoint. Standard tier only. Write requests should target 'port'.")),
		mcp.WithBoolean("satisfiesPzi", mcp.Description("Input parameter: Optional. Output only. Reserved for future use.")),
		mcp.WithObject("persistenceConfig", mcp.Description("Input parameter: Configuration of the persistence functionality.")),
		mcp.WithString("currentLocationId", mcp.Description("Input parameter: Output only. The current zone where the Redis primary node is located. In basic tier, this will always be the same as [location_id]. In standard tier, this can be the zone of any node in the instance.")),
		mcp.WithString("statusMessage", mcp.Description("Input parameter: Output only. Additional information about the current status of this instance, if available.")),
		mcp.WithString("state", mcp.Description("Input parameter: Output only. The current state of this instance.")),
		mcp.WithString("locationId", mcp.Description("Input parameter: Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone from the specified region for the instance. For standard tier, additional nodes will be added across multiple zones for protection against zonal failures. If specified, at least one node will be provisioned in this zone.")),
		mcp.WithString("tier", mcp.Description("Input parameter: Required. The service tier of the instance.")),
		mcp.WithString("reservedIpRange", mcp.Description("Input parameter: Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED the default block size is /28.")),
		mcp.WithString("host", mcp.Description("Input parameter: Output only. Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.")),
		mcp.WithObject("maintenanceSchedule", mcp.Description("Input parameter: Upcoming maintenance schedule. If no maintenance is scheduled, fields are not populated.")),
		mcp.WithNumber("memorySizeGb", mcp.Description("Input parameter: Required. Redis memory size in GiB.")),
		mcp.WithArray("suspensionReasons", mcp.Description("Input parameter: Optional. reasons that causes instance in \"SUSPENDED\" state.")),
		mcp.WithString("readReplicasMode", mcp.Description("Input parameter: Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.")),
		mcp.WithString("readEndpoint", mcp.Description("Input parameter: Output only. Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only. Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes will exhibit some lag behind the primary. Write requests must target 'host'.")),
		mcp.WithString("secondaryIpRange", mcp.Description("Input parameter: Optional. Additional IP range for node placement. Required when enabling read replicas on an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or \"auto\". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address range associated with the private service access connection, or \"auto\".")),
		mcp.WithObject("maintenancePolicy", mcp.Description("Input parameter: Maintenance policy for an instance.")),
		mcp.WithString("connectMode", mcp.Description("Input parameter: Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.")),
		mcp.WithString("authorizedNetwork", mcp.Description("Input parameter: Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.")),
		mcp.WithString("displayName", mcp.Description("Input parameter: An arbitrary and optional user-provided name for the instance.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Redis_projects_locations_instances_createHandler(cfg),
	}
}

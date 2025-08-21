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

func Redis_projects_locations_clusters_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["clusterId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("clusterId=%v", val))
		}
		if val, ok := args["requestId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("requestId=%v", val))
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
		var requestBody models.Cluster
		
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
		url := fmt.Sprintf("%s/v1/%s/clusters%s", cfg.BaseURL, parent, queryString)
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

func CreateRedis_projects_locations_clusters_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_parent_clusters",
		mcp.WithDescription("Creates a Redis cluster based on the specified properties. The creation is executed asynchronously and callers may check the returned operation to track its progress. Once the operation is completed the Redis cluster will be fully functional. The completed longrunning.Operation will contain the new cluster object in the response field. The returned operation is automatically deleted after a few hours, so there is no need to call DeleteOperation."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. The resource name of the cluster location using the form: `projects/{project_id}/locations/{location_id}` where `location_id` refers to a GCP region.")),
		mcp.WithString("clusterId", mcp.Description("Required. The logical name of the Redis cluster in the customer project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the customer project / location")),
		mcp.WithString("requestId", mcp.Description("Idempotent request UUID.")),
		mcp.WithArray("pscConnections", mcp.Description("Input parameter: Output only. PSC connections for discovery of the cluster topology and accessing the cluster.")),
		mcp.WithObject("stateInfo", mcp.Description("Input parameter: Represents additional information about the state of the cluster.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. The timestamp associated with the cluster creation request.")),
		mcp.WithString("name", mcp.Description("Input parameter: Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/clusters/{cluster_id}`")),
		mcp.WithNumber("shardCount", mcp.Description("Input parameter: Required. Number of shards for the Redis cluster.")),
		mcp.WithString("state", mcp.Description("Input parameter: Output only. The current state of this cluster. Can be CREATING, READY, UPDATING, DELETING and SUSPENDED")),
		mcp.WithString("authorizationMode", mcp.Description("Input parameter: Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster.")),
		mcp.WithArray("discoveryEndpoints", mcp.Description("Input parameter: Output only. Endpoints created on each given network, for Redis clients to connect to the cluster. Currently only one discovery endpoint is supported.")),
		mcp.WithArray("pscConfigs", mcp.Description("Input parameter: Required. Each PscConfig configures the consumer network where IPs will be designated to the cluster for client access through Private Service Connect Automation. Currently, only one PscConfig is supported.")),
		mcp.WithNumber("sizeGb", mcp.Description("Input parameter: Output only. Redis memory size in GB for the entire cluster rounded up to the next integer.")),
		mcp.WithString("uid", mcp.Description("Input parameter: Output only. System assigned, unique identifier for the cluster.")),
		mcp.WithNumber("replicaCount", mcp.Description("Input parameter: Optional. The number of replica nodes per shard.")),
		mcp.WithString("transitEncryptionMode", mcp.Description("Input parameter: Optional. The in-transit encryption for the Redis cluster. If not provided, encryption is disabled for the cluster.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Redis_projects_locations_clusters_createHandler(cfg),
	}
}

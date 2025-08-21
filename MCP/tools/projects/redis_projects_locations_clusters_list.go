package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google-cloud-memorystore-for-redis-api/mcp-server/config"
	"github.com/google-cloud-memorystore-for-redis-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Redis_projects_locations_clusters_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["pageSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageSize=%v", val))
		}
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
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
		url := fmt.Sprintf("%s/v1/%s/clusters%s", cfg.BaseURL, parent, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.ListClustersResponse
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

func CreateRedis_projects_locations_clusters_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_parent_clusters",
		mcp.WithDescription("Lists all Redis clusters owned by a project in either the specified location (region) or all locations. The location should have the following format: * `projects/{project_id}/locations/{location_id}` If `location_id` is specified as `-` (wildcard), then all regions available to the project are queried, and the results are aggregated."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. The resource name of the cluster location using the form: `projects/{project_id}/locations/{location_id}` where `location_id` refers to a GCP region.")),
		mcp.WithNumber("pageSize", mcp.Description("The maximum number of items to return. If not specified, a default value of 1000 will be used by the service. Regardless of the page_size value, the response may include a partial list and a caller should only rely on response's `next_page_token` to determine if there are more clusters left to be queried.")),
		mcp.WithString("pageToken", mcp.Description("The `next_page_token` value returned from a previous ListClusters request, if any.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Redis_projects_locations_clusters_listHandler(cfg),
	}
}

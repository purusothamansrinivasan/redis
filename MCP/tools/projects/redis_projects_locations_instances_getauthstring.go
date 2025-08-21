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

func Redis_projects_locations_instances_getauthstringHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		nameVal, ok := args["name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: name"), nil
		}
		name, ok := nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: name"), nil
		}
		queryParams := make([]string, 0)
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
		url := fmt.Sprintf("%s/v1/%s/authString%s", cfg.BaseURL, name, queryString)
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
		var result models.InstanceAuthString
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

func CreateRedis_projects_locations_instances_getauthstringTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_name_authString",
		mcp.WithDescription("Gets the AUTH string for a Redis instance. If AUTH is not enabled for the instance the response will be empty. This information is not included in the details returned to GetInstance."),
		mcp.WithString("name", mcp.Required(), mcp.Description("Required. Redis instance resource name using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` where `location_id` refers to a GCP region.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Redis_projects_locations_instances_getauthstringHandler(cfg),
	}
}

package main

import (
	"github.com/google-cloud-memorystore-for-redis-api/mcp-server/config"
	"github.com/google-cloud-memorystore-for-redis-api/mcp-server/models"
	tools_projects "github.com/google-cloud-memorystore-for-redis-api/mcp-server/tools/projects"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_projects.CreateRedis_projects_locations_instances_reschedulemaintenanceTool(cfg),
		tools_projects.CreateRedis_projects_locations_instances_upgradeTool(cfg),
		tools_projects.CreateRedis_projects_locations_clusters_listTool(cfg),
		tools_projects.CreateRedis_projects_locations_clusters_createTool(cfg),
		tools_projects.CreateRedis_projects_locations_instances_getauthstringTool(cfg),
		tools_projects.CreateRedis_projects_locations_operations_cancelTool(cfg),
		tools_projects.CreateRedis_projects_locations_instances_failoverTool(cfg),
		tools_projects.CreateRedis_projects_locations_listTool(cfg),
		tools_projects.CreateRedis_projects_locations_instances_exportTool(cfg),
		tools_projects.CreateRedis_projects_locations_operations_deleteTool(cfg),
		tools_projects.CreateRedis_projects_locations_operations_getTool(cfg),
		tools_projects.CreateRedis_projects_locations_instances_patchTool(cfg),
		tools_projects.CreateRedis_projects_locations_operations_listTool(cfg),
		tools_projects.CreateRedis_projects_locations_instances_listTool(cfg),
		tools_projects.CreateRedis_projects_locations_instances_createTool(cfg),
		tools_projects.CreateRedis_projects_locations_instances_importTool(cfg),
	}
}

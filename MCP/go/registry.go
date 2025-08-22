package main

import (
	"github.com/search-services/mcp-server/config"
	"github.com/search-services/mcp-server/models"
	tools_search "github.com/search-services/mcp-server/tools/search"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_search.CreateGet_search_v1_organicTool(cfg),
		tools_search.CreateGet_search_v1_scrapeTool(cfg),
		tools_search.CreateGet_search_v1_fieldsTool(cfg),
	}
}

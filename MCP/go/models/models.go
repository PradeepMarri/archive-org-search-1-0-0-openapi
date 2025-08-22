package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// OrganicResult represents the OrganicResult schema from the OpenAPI specification
type OrganicResult struct {
	Count int `json:"count,omitempty"` // Number of results returned
	Items []Hit `json:"items,omitempty"`
	Total int `json:"total,omitempty"` // Approximate number of results meeting query
}

// ScrapeResult represents the ScrapeResult schema from the OpenAPI specification
type ScrapeResult struct {
	Items []Hit `json:"items,omitempty"`
	Previous string `json:"previous,omitempty"` // A scroll handle
	Total int `json:"total,omitempty"` // Total number of results from this cursor point
	Count int `json:"count,omitempty"` // Number of results returned
	Cursor string `json:"cursor,omitempty"` // A scroll handle
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}

// Hit represents the Hit schema from the OpenAPI specification
type Hit struct {
}

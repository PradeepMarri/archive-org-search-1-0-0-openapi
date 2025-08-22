package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/search-services/mcp-server/config"
	"github.com/search-services/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_search_v1_scrapeHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["q"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("q=%v", val))
		}
		if val, ok := args["field"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("field=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("size=%v", val))
		}
		if val, ok := args["cursor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cursor=%v", val))
		}
		if val, ok := args["total_only"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("total_only=%v", val))
		}
		if val, ok := args["callback"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("callback=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/search/v1/scrape%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
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
		var result models.ScrapeResult
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

func CreateGet_search_v1_scrapeTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_search_v1_scrape",
		mcp.WithDescription("Scrape search results from Internet Archive, allowing a scrolling cursor
"),
		mcp.WithString("q", mcp.Description("Lucene-type search query")),
		mcp.WithString("field", mcp.Description("Metadata field")),
		mcp.WithString("sort", mcp.Description("sort collations")),
		mcp.WithNumber("size", mcp.Description("Number of query results to return")),
		mcp.WithString("cursor", mcp.Description("Cursor for scrolling (used for subsequent calls)")),
		mcp.WithBoolean("total_only", mcp.Description("Request total only; do not return hits")),
		mcp.WithString("callback", mcp.Description("Specifies a JavaScript function func, for a JSON-P response. When provided, results are wrapped as `callback(data)`, and the returned MIME type is application/javascript. This causes the caller to automatically run the func with the JSON results as its argument.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_search_v1_scrapeHandler(cfg),
	}
}

package controllers

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"github.com/rimo10/youtube-api-server/src/config"
)

var searchResults []config.Searchapi

func Get(c *fiber.Ctx) error {

	flag.Parse()
	query := c.Query("query", "")
	counts, err := strconv.ParseInt(c.Query("count", ""), 10, 64)

	if err != nil {
		return c.JSON("Incorrect query")
	}

	getresponse := make([]*config.Searchapi, 0)
	config.Database.Where("query = ?", query).Find(&getresponse)

	response := make([]struct {
		VideoId     string `json:"VideoId"`
		Title       string `json:"Title"`
		Description string `json:"Description"`
		ChannelName string `json:"ChannelName"`
		PublishedAt string `json:"PublishedAt"`
	}, len(searchResults))

	for _, result := range getresponse {
		if int64(len(response)) >= counts {
			break
		}
		response = append(response, struct {
			VideoId     string `json:"VideoId"`
			Title       string `json:"Title"`
			Description string `json:"Description"`
			ChannelName string `json:"ChannelName"`
			PublishedAt string `json:"PublishedAt"`
		}{
			VideoId:     result.VideoId,
			Title:       result.Title,
			Description: result.Description,
			ChannelName: result.ChannelName,
			PublishedAt: result.PublishedAt,
		})
	}
	return c.JSON(response)
}

package handler

import (
	"moyu/campaign"
	"moyu/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

// api/v1/campaigns (Get Campaigns by userID or all)
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.FindCampaigns(userID)
	if err != nil {
		campaignsFormatter := campaign.FormatCampaigns(campaigns)

		response := helper.APIResponse("Failed to get campaigns", "error", http.StatusBadRequest, campaignsFormatter)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	campaignsFormatter := campaign.FormatCampaigns(campaigns)

	response := helper.APIResponse("Successfuly get list of campaigns", "success", http.StatusOK, campaignsFormatter)
	c.JSON(http.StatusOK, response)
}

// api/v1/campaigns/{1} (Get Campaign by ID or Campaign Detail)
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// handler : mapping id yg di url ke struct input => kirim service, call formatter
	// service : tangkap input dari struct input, panggil repo
	// repository : get campaign by id

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign detail", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaigns, err := h.campaignService.FindCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign detail", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatDetailCampaign(campaigns)

	response := helper.APIResponse("Successfuly to get campaign detail", "success", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

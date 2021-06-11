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

// api/v1/campaigns/user_id?
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

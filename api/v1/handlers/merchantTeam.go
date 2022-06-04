package handlers

import (
	"encoding/json"
	"merchant-service/api/v1/services"
	"merchant-service/models"
	"merchant-service/utils"
	"net/http"
	"strconv"
)

type MerchantTeamHandler struct {
	Service services.MerchantTeamServiceInterface
}

func MerchantTeamHttpHandler(service services.MerchantTeamServiceInterface) *MerchantTeamHandler {
	return &MerchantTeamHandler{service}
}

func (mth MerchantTeamHandler) GetMerchantTeamList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		page := r.URL.Query().Get("page")
		pageInt, _ := strconv.Atoi(page)
		merchantID := r.URL.Query().Get("merchant_id")
		merchantIDInt, _ := strconv.Atoi(merchantID)
		response := mth.Service.GetMerchantTeams(r.Context(), pageInt, uint64(merchantIDInt))
		utils.ResponseJson(w, response)
	} else {
		var response models.ResponseJsonApi
		response.Success = false
		response.ResponseCode = http.StatusMethodNotAllowed
		response.Error = "Invalid request method."
		utils.ResponseJson(w, response)
	}
}

func (mth MerchantTeamHandler) CreateMerchantTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var request models.CreateMerchantTeamRequest
		json.NewDecoder(r.Body).Decode(&request)
		response := mth.Service.CreatetMerchantTeam(r.Context(), &request)
		utils.ResponseJson(w, response)
	} else {
		var response models.ResponseJsonApi
		response.Success = false
		response.ResponseCode = http.StatusMethodNotAllowed
		response.Error = "Invalid request method."
		utils.ResponseJson(w, response)
	}
}

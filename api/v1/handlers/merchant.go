package handlers

import (
	"encoding/json"
	"fmt"
	"merchant-service/api/v1/services"
	"merchant-service/models"
	"merchant-service/utils"
	"net/http"
	"strconv"
)

type MerchantHandler struct {
	Service services.MerchantServiceInterface
}

func MerchantHttpHandler(service services.MerchantServiceInterface) *MerchantHandler {
	return &MerchantHandler{service}
}

func (mh MerchantHandler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ping...")

}

func (mh MerchantHandler) GetMerchantList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		page := r.URL.Query().Get("page")
		pageInt, _ := strconv.Atoi(page)
		response := mh.Service.GetMerchants(r.Context(), pageInt)
		utils.ResponseJson(w, response)
	} else {
		var response models.ResponseJsonApi
		response.Success = false
		response.ResponseCode = http.StatusMethodNotAllowed
		response.Error = "Invalid request method."
		utils.ResponseJson(w, response)
	}
}

func (mh MerchantHandler) CreatetMerchant(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var request models.CreateMerchantRequest
		json.NewDecoder(r.Body).Decode(&request)
		response := mh.Service.CreatetMerchant(r.Context(), &request)
		utils.ResponseJson(w, response)
	} else {
		var response models.ResponseJsonApi
		response.Success = false
		response.ResponseCode = http.StatusMethodNotAllowed
		response.Error = "Invalid request method."
		utils.ResponseJson(w, response)
	}
}

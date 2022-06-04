package services

import (
	"context"
	"errors"
	"merchant-service/models"
	"merchant-service/repository"
	"net/http"
)

type merchantTeamService struct {
	MerchantTeamRepository repository.MerchantTeamRepositoryInterface
}

type MerchantTeamServiceInterface interface {
	GetMerchantTeams(context.Context, int, uint64) models.ResponseJsonApi
	CreatetMerchantTeam(context.Context, *models.CreateMerchantTeamRequest) models.ResponseJsonApi
}

func NewMerchantTeamService(merchantTeamRepository repository.MerchantTeamRepositoryInterface) *merchantTeamService {
	return &merchantTeamService{merchantTeamRepository}
}

func (mts merchantTeamService) GetMerchantTeams(ctx context.Context, page int, merchantID uint64) models.ResponseJsonApi {
	var response models.ResponseJsonApi
	merchants, err := mts.MerchantTeamRepository.GetMerchantTeams(ctx, page, merchantID)
	if err != nil {
		response.Success = false
		response.ResponseCode = models.StatusCustomError
		response.Error = err.Error()
		return response
	}
	response.Success = true
	response.ResponseCode = http.StatusOK
	response.Data = merchants
	return response
}

func (mts merchantTeamService) CreatetMerchantTeam(ctx context.Context, data *models.CreateMerchantTeamRequest) models.ResponseJsonApi {
	var response models.ResponseJsonApi
	err := validateCreateMerchentTeamRequest(ctx, data)
	if err != nil {
		response.Success = false
		response.ResponseCode = models.StatusCustomError
		response.Error = err.Error()
		return response
	}
	err = mts.MerchantTeamRepository.CreatetMerchantTeam(ctx, data)
	if err != nil {
		response.Success = false
		response.ResponseCode = models.StatusCustomError
		response.Error = err.Error()
		return response
	}
	response.Success = true
	response.ResponseCode = http.StatusOK
	response.Data = "Merchant team created successfully."
	return response
}

func validateCreateMerchentTeamRequest(ctx context.Context, data *models.CreateMerchantTeamRequest) error {
	if data.Name == "" || data.Email == "" || data.Status > 1 || data.MerchantID == 0 {
		return errors.New("kindly check all inputs")
	}
	return nil
}

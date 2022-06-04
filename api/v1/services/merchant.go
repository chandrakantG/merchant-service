package services

import (
	"context"
	"errors"
	"merchant-service/models"
	"merchant-service/repository"
	"net/http"
)

type merchantService struct {
	MerchantRepository repository.MerchantRepositoryInterface
}

type MerchantServiceInterface interface {
	GetMerchants(context.Context, int) models.ResponseJsonApi
	CreatetMerchant(context.Context, *models.CreateMerchantRequest) models.ResponseJsonApi
}

func NewMerchantService(merchantRepository repository.MerchantRepositoryInterface) *merchantService {
	return &merchantService{merchantRepository}
}

func (ms merchantService) GetMerchants(ctx context.Context, page int) models.ResponseJsonApi {
	var response models.ResponseJsonApi
	merchants, err := ms.MerchantRepository.GetMerchants(ctx, page)
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

func (ms merchantService) CreatetMerchant(ctx context.Context, data *models.CreateMerchantRequest) models.ResponseJsonApi {
	var response models.ResponseJsonApi
	err := validateCreateMerchentRequest(ctx, data)
	if err != nil {
		response.Success = false
		response.ResponseCode = models.StatusCustomError
		response.Error = err.Error()
		return response
	}
	err = ms.MerchantRepository.CreatetMerchant(ctx, data)
	if err != nil {
		response.Success = false
		response.ResponseCode = models.StatusCustomError
		response.Error = err.Error()
		return response
	}
	response.Success = true
	response.ResponseCode = http.StatusOK
	response.Data = "Merchant created successfully."
	return response
}

func validateCreateMerchentRequest(ctx context.Context, data *models.CreateMerchantRequest) error {
	if data.Name == "" || data.Code == "" || (data.Status > 1) {
		return errors.New("kindly check all inputs")
	}
	return nil
}

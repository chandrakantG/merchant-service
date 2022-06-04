package server

import (
	"fmt"
	"log"
	"merchant-service/api/v1/handlers"
	"merchant-service/api/v1/services"
	"merchant-service/config"
	"merchant-service/database"
	"merchant-service/models"
	"merchant-service/repository"
	"merchant-service/utils"
	"net/http"
	"os"
)

type Server interface {
	Start()
}

type server struct{}

func NewServer() Server {
	return server{}
}

func (s server) Start() {
	os.Setenv("ENVIRONMENT", "app")
	err := config.SetupConfig()
	if err != nil {
		log.Fatal("Can not load config.")
	}

	config.LoadConfig()
	serviceConfig := config.GetServiceConfig()

	dbConn := database.DBConnection().GetConnection()

	if dbConn == nil {
		log.Fatal("Expecting db connection object but received nil")
	}

	merchantRepo := repository.NewMerchantRepository(dbConn)
	merchantService := services.NewMerchantService(merchantRepo)
	merchantHandler := handlers.MerchantHttpHandler(merchantService)

	merchantTeamRepo := repository.NewMerchantTeamRepository(dbConn)
	merchantTeamService := services.NewMerchantTeamService(merchantTeamRepo)
	merchantTeamHandler := handlers.MerchantTeamHttpHandler(merchantTeamService)

	http.HandleFunc("/ping", merchantHandler.Ping)

	http.HandleFunc("/api/v1/merchant/list", auth(merchantHandler.GetMerchantList))
	http.HandleFunc("/api/v1/merchant/create", auth(merchantHandler.CreatetMerchant))

	http.HandleFunc("/api/v1/merchant-team/list", auth(merchantTeamHandler.GetMerchantTeamList))
	http.HandleFunc("/api/v1/merchant-team/create", auth(merchantTeamHandler.CreateMerchantTeam))

	fmt.Println("services started")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", serviceConfig.ServiceHost, serviceConfig.ServicePort), nil))
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var response models.ResponseJsonApi
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			response.Success = false
			response.ResponseCode = http.StatusBadRequest
			response.Error = "content type should be application/json."
			utils.ResponseJson(w, response)
			return
		}
		tokenVal := r.Header.Get("Authorization")
		if tokenVal == "" {
			response.Success = false
			response.ResponseCode = http.StatusUnauthorized
			response.Error = "authorization is required."
			utils.ResponseJson(w, response)
			return
		}
		if err := utils.ValidateToken(tokenVal); err != nil {
			response.Success = false
			response.ResponseCode = http.StatusUnauthorized
			response.Error = "authorization is invalid."
			utils.ResponseJson(w, response)
			return
		}
		next(w, r)
	})
}

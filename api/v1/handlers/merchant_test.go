package handlers

import (
	"log"
	"merchant-service/api/v1/services"
	"merchant-service/config"
	"merchant-service/database"
	"merchant-service/repository"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestMerchantHttpHandler(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	err := config.SetupConfig()
	if err != nil {
		log.Fatal("Can not load config.")
	}
	config.LoadConfig()

	dbConn := database.DBConnection().GetConnection()
	// fmt.Println("dbConn:", dbConn)
	t.Log("dbConn:", dbConn)

	if dbConn == nil {
		log.Fatal("Expecting db connection object but received nil")
	}
	merchantRepo := repository.NewMerchantRepository(dbConn)
	merchantService := services.NewMerchantService(merchantRepo)
	merchantHandler := MerchantHttpHandler(merchantService)

	type args struct {
		service services.MerchantServiceInterface
	}

	tests := []struct {
		name string
		args args
		want *MerchantHandler
	}{
		// TODO: Add test cases.
		{
			args: args{
				service: merchantService,
			},
			want: merchantHandler,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MerchantHttpHandler(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MerchantHttpHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerchantHandler_Ping(t *testing.T) {
	type fields struct {
		Service services.MerchantServiceInterface
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mh := MerchantHandler{
				Service: tt.fields.Service,
			}
			mh.Ping(tt.args.w, tt.args.r)
		})
	}
}

func TestMerchantHandler_GetMerchantList(t *testing.T) {
	type fields struct {
		Service services.MerchantServiceInterface
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mh := MerchantHandler{
				Service: tt.fields.Service,
			}
			mh.GetMerchantList(tt.args.w, tt.args.r)
		})
	}
}

func TestMerchantHandler_CreatetMerchant(t *testing.T) {
	type fields struct {
		Service services.MerchantServiceInterface
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mh := MerchantHandler{
				Service: tt.fields.Service,
			}
			mh.CreatetMerchant(tt.args.w, tt.args.r)
		})
	}
}

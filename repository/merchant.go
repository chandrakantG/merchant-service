package repository

import (
	"context"
	"database/sql"
	"errors"
	"merchant-service/models"
	"time"
)

type MerchantRepository struct {
	DB *sql.DB
}

type MerchantRepositoryInterface interface {
	GetMerchants(context.Context, int) (models.MerchantsWithPagination, error)
	CreatetMerchant(context.Context, *models.CreateMerchantRequest) error
}

func NewMerchantRepository(db *sql.DB) *MerchantRepository {
	return &MerchantRepository{DB: db}
}

func (mr MerchantRepository) GetMerchants(ctx context.Context, page int) (models.MerchantsWithPagination, error) {
	var (
		merchants               []models.Merchant
		merchant                models.Merchant
		start                   = 0
		pagination              models.Pagination
		merchantsWithPagination models.MerchantsWithPagination
	)
	count, countErr := mr.GetMerchansCount(ctx)
	if countErr != nil {
		return merchantsWithPagination, countErr
	}

	if page == 1 {
		if count >= 2 {
			pagination.NextPage = true
		} else {
			pagination.NextPage = false
		}
		pagination.PreviousPage = false
		pagination.Page = 1
	} else {
		start = (page - 1) * 2

		if count >= uint64(start+2) {
			pagination.NextPage = true
		} else {
			pagination.NextPage = false
		}
		pagination.PreviousPage = true
		pagination.Page = page
	}

	rows, err := mr.DB.Query("select id, name, code, status, created_at, updated_at from merchants order by id desc limit ?,2;", start)
	if err != nil {
		return merchantsWithPagination, err
	}

	for rows.Next() {
		if err := rows.Scan(&merchant.ID, &merchant.Name, &merchant.Code, &merchant.Status, &merchant.CreatedAt, &merchant.UpdatedAt); err != nil {
			return merchantsWithPagination, err
		}
		merchants = append(merchants, merchant)
	}
	merchantsWithPagination.Merchants = merchants
	merchantsWithPagination.Pagination = pagination
	return merchantsWithPagination, nil
}

func (mr MerchantRepository) GetMerchantByCode(ctx context.Context, code string) (models.Merchant, error) {
	var (
		merchant models.Merchant
	)

	rows, err := mr.DB.Query("select id, name, code, status, created_at, updated_at from merchants where code = ?;", code)
	if err != nil {
		return merchant, err
	}

	if rows.Next() {
		if err := rows.Scan(&merchant.ID, &merchant.Name, &merchant.Code, &merchant.Status, &merchant.CreatedAt, &merchant.UpdatedAt); err != nil {
			return merchant, err
		}
	}
	return merchant, nil
}

func (mr MerchantRepository) GetMerchansCount(ctx context.Context) (uint64, error) {
	var count uint64
	rows, err := mr.DB.Query("select count(id) from merchants;")
	if err != nil {
		return count, err
	}

	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return count, err
		}
	}
	return count, nil
}

func (mr MerchantRepository) CreatetMerchant(ctx context.Context, data *models.CreateMerchantRequest) error {
	if data.Name == "" {
		return errors.New("name should not empty")
	}
	if data.Code == "" {
		return errors.New("code should not empty")
	}
	if data.Status > 1 {
		return errors.New("status should be 0 or 1")
	}

	merchant, err := mr.GetMerchantByCode(ctx, data.Code)
	if err != nil {
		return err
	} else if merchant.ID > 0 {
		return errors.New("merchant already exist")
	}

	// TODO change table name
	stmt, err := mr.DB.Prepare("INSERT into merchants SET name=?, code=?, status=?, created_at=?, updated_at=?")
	if err != nil {
		return err

	}
	t := time.Now()
	ts := t.Format("2006-01-02 15:04:05")
	defer stmt.Close()
	_, err = stmt.Exec(data.Name, data.Code, data.Status, ts, ts)
	if err != nil {
		return err
	}

	return nil
}

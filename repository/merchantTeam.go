package repository

import (
	"context"
	"database/sql"
	"errors"
	"merchant-service/models"
	"time"
)

type MerchantTeamRepository struct {
	DB *sql.DB
}

type MerchantTeamRepositoryInterface interface {
	GetMerchantTeams(context.Context, int, uint64) (models.MerchantTeamsWithPagination, error)
	CreatetMerchantTeam(context.Context, *models.CreateMerchantTeamRequest) error
}

func NewMerchantTeamRepository(db *sql.DB) *MerchantTeamRepository {
	return &MerchantTeamRepository{DB: db}
}

func (mr MerchantTeamRepository) GetMerchantTeams(ctx context.Context, page int, merchantID uint64) (models.MerchantTeamsWithPagination, error) {
	var (
		merchantTeams               []models.MerchantTeam
		merchantTeam                models.MerchantTeam
		start                       = 0
		pagination                  models.Pagination
		merchantTeamsWithPagination models.MerchantTeamsWithPagination
	)
	count, countErr := mr.GetMerchantTeamsCount(ctx, merchantID)
	if countErr != nil {
		return merchantTeamsWithPagination, countErr
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

	rows, err := mr.DB.Query("select id, merchant_id, name, email, status, created_at, updated_at from merchant_teams where merchant_id = ? order by id desc limit ?,2;", merchantID, start)
	if err != nil {
		return merchantTeamsWithPagination, err
	}

	for rows.Next() {
		if err := rows.Scan(&merchantTeam.ID, &merchantTeam.MerchantID, &merchantTeam.Name, &merchantTeam.Email, &merchantTeam.Status, &merchantTeam.CreatedAt, &merchantTeam.UpdatedAt); err != nil {
			return merchantTeamsWithPagination, err
		}
		merchantTeams = append(merchantTeams, merchantTeam)
	}
	merchantTeamsWithPagination.MerchantTeams = merchantTeams
	merchantTeamsWithPagination.Pagination = pagination
	return merchantTeamsWithPagination, nil
}

func (mr MerchantTeamRepository) GetMerchantTeamByMerchantIDByEmail(ctx context.Context, merchantID uint, email string) (models.MerchantTeam, error) {
	var (
		merchantTeam models.MerchantTeam
	)

	rows, err := mr.DB.Query("select id, merchant_id, name, email, status, created_at, updated_at from merchant_teams where merchant_id = ? and email = ?;", merchantID, email)
	if err != nil {
		return merchantTeam, err
	}

	if rows.Next() {
		if err := rows.Scan(&merchantTeam.ID, &merchantTeam.MerchantID, &merchantTeam.Name, &merchantTeam.Email, &merchantTeam.Status, &merchantTeam.CreatedAt, &merchantTeam.UpdatedAt); err != nil {
			return merchantTeam, err
		}
	}
	return merchantTeam, nil
}

func (mr MerchantTeamRepository) GetMerchantTeamsCount(ctx context.Context, merchantID uint64) (uint64, error) {
	var count uint64
	rows, err := mr.DB.Query("select count(id) from merchant_teams where merchant_id = ?;", merchantID)
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

func (mr MerchantTeamRepository) CreatetMerchantTeam(ctx context.Context, data *models.CreateMerchantTeamRequest) error {
	if data.Name == "" {
		return errors.New("name should not empty")
	}
	if data.Email == "" {
		return errors.New("code should not empty")
	}
	if data.Status > 1 {
		return errors.New("status should be 0 or 1")
	}

	if data.MerchantID <= 0 {
		return errors.New("merchant id should be greater than 0")
	}

	merchantTeam, err := mr.GetMerchantTeamByMerchantIDByEmail(ctx, data.MerchantID, data.Email)
	if err != nil {
		return err
	} else if merchantTeam.ID > 0 {
		return errors.New("merchant team already exist")
	}

	// TODO change table name
	stmt, err := mr.DB.Prepare("INSERT into merchant_teams SET merchant_id=?, name=?, email=?, status=?, created_at=?, updated_at=?")
	if err != nil {
		return err

	}
	t := time.Now()
	ts := t.Format("2006-01-02 15:04:05")
	defer stmt.Close()
	_, err = stmt.Exec(data.MerchantID, data.Name, data.Email, data.Status, ts, ts)
	if err != nil {
		return err
	}

	return nil
}

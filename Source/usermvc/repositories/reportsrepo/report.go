package reportsrepo

import (
	"context"
	constants "usermvc"
	"usermvc/entity"
	"usermvc/repositories"
	logger2 "usermvc/utility/logger"

	"github.com/jinzhu/gorm"
)

type ReportsRepo interface {
	GetConfirmedOrders(ctx context.Context) ([]*entity.ConfirmedOrderReport, error)
}

type reportsRepo struct {
	db *gorm.DB
}

func NewReportsRepo() ReportsRepo {
	newDb, err := repositories.NewDb()
	if err != nil {
		panic(err)
	}

	newDb.AutoMigrate(&entity.ConfirmedOrderReport{})

	return &reportsRepo{
		db: newDb,
	}
}

func (r reportsRepo) GetConfirmedOrders(ctx context.Context) ([]*entity.ConfirmedOrderReport, error) {
	var result []*entity.ConfirmedOrderReport
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all confirmed order details")
	err := r.db.Table(constants.ConfirmedOrderReport).Select("*").Order("serialno asc").Scan(&result).Error
	if err != nil {
		logger.Error("error while get all confirmed order from ", constants.ConfirmedOrderReport, err.Error())
		return nil, err
	}

	return result, nil
}

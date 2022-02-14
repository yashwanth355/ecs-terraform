package controller

import (
	"usermvc/repositories/reportsrepo"
	logger2 "usermvc/utility/logger"

	"github.com/gin-gonic/gin"
)

type ReportsController interface {
	GetConfirmedOrders(ctx *gin.Context)
}

type reportsController struct {
	reportsRepo reportsrepo.ReportsRepo
}

func newReportsController() ReportsController {
	return &reportsController{
		reportsRepo: reportsrepo.NewReportsRepo(),
	}
}

func (r reportsController) GetConfirmedOrders(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	res, err := r.reportsRepo.GetConfirmedOrders(ctx)
	if err != nil {
		logger.Error("error while getting all confirmed order details", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	logger.Info("getting response from get reportsRepo.GetConfirmedOrders ", res)
	ctx.JSON(200, res)
}

package controller

import (
	"context"
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories/gcrepo"
	logger2 "usermvc/utility/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GCController interface {
	InsertGcDetails(ctx *gin.Context)
	ViewGcDetails(ctx *gin.Context)
	ListGcDetails(ctx *gin.Context)
}
type gcController struct {
	gcRepo gcrepo.GcRepo
}

func NewGcController() GCController {
	return &gcController{
		gcRepo: gcrepo.NewgcRepo(),
	}
}

func (gc gcController) InsertGcDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var gcs model.GCDetails
	if err := ctx.ShouldBindJSON(&gcs); err != nil {
		logger.Error("Error while inserting the data")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := gc.gcRepo.InsertGcDetails(context.Background(), entity.GCDetails(gcs))
	if err != nil {
		zap.S().Error("not able to parse the request", err.Error())
		ctx.JSON(200, err.Error())
		return
	}
	ctx.JSON(200, res)
} //end of function

func (gc gcController) ViewGcDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var view model.GCViewDetails
	logger.Info("start handling gc supplier")
	if err := ctx.ShouldBindJSON(&view); err != nil {
		logger.Error("Error while parsing")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := gc.gcRepo.ViewGcDetails(ctx, &view)
	if err != nil {
		ctx.JSON(503, "error while getting")
		return
	}
	ctx.JSON(200, res)
	return
}

func (gc gcController) ListGcDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var list *model.InputG
	var list1 *model.ListGCDetails
	logger.Info("start handling of list gc details")
	if err := ctx.ShouldBindJSON(&list); err != nil {
		logger.Error("Error while parsing the list gc")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := gc.gcRepo.ListGcDetails(ctx, list1)
	if err != nil {
		ctx.JSON(503, "error")
		return
	}
	ctx.JSON(200, res)
	return

}

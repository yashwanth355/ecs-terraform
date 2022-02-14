package controller

import (
	"context"
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories/supplierrepo"
	logger2 "usermvc/utility/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SupplierController interface {
	InsertSupplierDetails(ctx *gin.Context)
	ViewSupplierDetails(ctx *gin.Context)
	ListSupplierDetails(ctx *gin.Context)
}
type supplierController struct {
	supplierRepo supplierrepo.SupplierRepo
}

func NewsupplierController() SupplierController {
	return &supplierController{
		supplierRepo: supplierrepo.NewsupplierRepo(),
	}
}
func (sc supplierController) ListSupplierDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var suppliers *model.ListSupplier
	logger.Info("start handling supplier controller")
	if err := ctx.ShouldBindJSON(&suppliers); err != nil {
		logger.Error("Error while parsing the supplier info")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := sc.supplierRepo.ListSupplierDetails(ctx, suppliers)
	if err != nil {
		ctx.JSON(503, "error")
		return
	}
	ctx.JSON(200, res)
	return

}

func (sc supplierController) InsertSupplierDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var insertsupplier model.VendorDetails
	if err := ctx.ShouldBindJSON(&insertsupplier); err != nil {
		logger.Error("Error while parsing the insert supplier details")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := sc.supplierRepo.InsertSupplierDetails(context.Background(), entity.VendorDetails(insertsupplier))
	if err != nil {
		zap.S().Error("not able to parse the request", err.Error())
		ctx.JSON(200, err.Error())
		return
	}
	ctx.JSON(200, res)
}
func (sc supplierController) ViewSupplierDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var view model.VendorDetails
	logger.Info("start handling supplier module")
	if err := ctx.ShouldBindJSON(&view); err != nil {
		logger.Error("Error while parsing")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := sc.supplierRepo.ViewSupplierDetails(ctx, &view)
	if err != nil {
		ctx.JSON(503, "error while getting")
		return
	}
	ctx.JSON(200, res)
	return

}
//

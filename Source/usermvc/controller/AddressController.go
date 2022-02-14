package controller

import (
	"github.com/gin-gonic/gin"
	"usermvc/model"
	"usermvc/repositories/addressrepo"
	logger2 "usermvc/utility/logger"
)

type AddressController interface {
	GetCountries(ctx *gin.Context)
	GetStates(ctx *gin.Context)
	GetCities(ctx *gin.Context)
	
}

type addressController struct {
	addressRepo addressrepo.AddressRepo
}

func NewAddressController() AddressController {
	return addressController{
		addressRepo: addressrepo.NewAddressRepo(),
	}
}

func (ad addressController) GetCountries(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("started handling addressController.GetCountries")
	res, err := ad.addressRepo.GetCountries(ctx)
	if err != nil {
		logger.Error("error while fetching all country details from addressRepo ", err.Error())
		ctx.JSON(503, err.Error())
	}
	ctx.JSON(200, res)
}
func (ad addressController) GetStates(ctx *gin.Context){
	logger := logger2.GetLoggerWithContext(ctx)
	var getStates model.Countries
	logger.Info("started handling poController.GetPoCreationInfo")
	if err := ctx.ShouldBindJSON(&getStates); err != nil {
		logger.Error("error while parsing the getStates request body", err.Error())
		ctx.JSON(403, err.Error())
		return
	}
	res, err := ad.addressRepo.GetStates(ctx, getStates)
	if err != nil {
		ctx.JSON(503, "error while fetching states info from address repo")
		return
	}
	ctx.JSON(200, res)
	return
}
func (ad addressController) GetCities(ctx *gin.Context){
	logger := logger2.GetLoggerWithContext(ctx)
	var getCity model.States
	logger.Info("started handling addressController.GetCities")
	if err := ctx.ShouldBindJSON(&getCity); err != nil {
		logger.Error("error while parsing the getCity request body", err.Error())
		ctx.JSON(403, err.Error())
		return
	}
	res, err := ad.addressRepo.GetCities(ctx, getCity)
	if err != nil {
		ctx.JSON(503, "error while fetching cities info from repo")
		return
	}
	ctx.JSON(200, res)
	return
}
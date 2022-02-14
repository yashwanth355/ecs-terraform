package controller

import (
	"context"
	"log"
	"usermvc/model"
	"usermvc/repositories/userrepo"
	logger2 "usermvc/utility/logger"

	//"github.com/aws/aws-sdk-go/private/model/api/codegentest/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"usermvc/service/user"
)

//
//func (ctr controller) AdminConfirmSignup (c *gin.Context) {
//	var credential model.Credentials
//	log.Println("entering readFirst")
//	if err :=c.BindJSON(&credential);err != nil {
//		zap.S().Error("error while marshalling json ", err.Error())
//	}
//
//	res, err := ctr.userSvc.AdminConfirmSignUp(context.Background(), credential)
//	if err != nil {
//		c.JSON(200, err.Error())
//	}
//	c.JSON(200, res)
//}

type UserController interface {
	GetAllUsersDetail(ctx *gin.Context)
	GetAllCompanyNames(ctx *gin.Context)
	GetAllDepartmentName(ctx *gin.Context)
	GetAllDesignationName(ctx *gin.Context)
	GetDivisions(ctx *gin.Context)
	GetUserRole(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	// GetCountryDetails(ctx *gin.Context)
	// GetStateDetails(ctx *gin.Context)
	// GetCityDetails(ctx *gin.Context)
}

type userController struct {
	userRepo userrepo.UserRepo
	userSvc  user.Service
}

func newUserController() UserController {
	return userController{
		userRepo: userrepo.NewUserRepo(),
		userSvc:  user.NewuserService(),
	}
}

func (uc userController) CreateUser(c *gin.Context) {
	var user model.UserResquest
	log.Println("entering readFirst")
	if err := c.BindJSON(&user); err != nil {
		zap.S().Error("error while marshalling json ", err.Error())
	}

	res, err := uc.userSvc.CreateUser(context.Background(), user)
	if err != nil {
		c.JSON(200, err.Error())
	}
	c.JSON(200, res)
}

func (uc userController) GetUserRole(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	res, err := uc.userRepo.GetUserRole(ctx)
	if err != nil {
		logger.Error("Error while parsing the request", err.Error())
		ctx.JSON(503, err.Error())
		return
	}

	ctx.JSON(200, res)
}

func (uc userController) GetAllUsersDetail(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	res, err := uc.userRepo.GetAllUsersDetail(ctx)
	if err != nil {
		logger.Error("error while getting userDetails from userRepo ", err.Error())
		ctx.JSON(503, err.Error())
	}
	ctx.JSON(200, res)
}

func (uc userController) GetAllCompanyNames(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	res, err := uc.userRepo.GetAllCompanyNames(ctx)
	if err != nil {
		logger.Error("error while getting all companies details from userRepo ", err.Error())
		ctx.JSON(503, err.Error())
	}
	ctx.JSON(200, res)
}

func (uc userController) GetAllDepartmentName(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	res, err := uc.userRepo.GetAllDepartmentName(ctx)
	if err != nil {
		logger.Error("error while getting department names from userRepo ", err.Error())
		ctx.JSON(503, err.Error())
	}
	ctx.JSON(200, res)
}

func (uc userController) GetAllDesignationName(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	res, err := uc.userRepo.GetAllDesignationName(ctx)
	if err != nil {
		logger.Error("error while getting desination  names from userRepo ", err.Error())
		ctx.JSON(503, err.Error())
	}
	ctx.JSON(200, res)
}

func (uc userController) GetDivisions(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var req model.GetDivisionsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Info("error while parsing request into devision request", err.Error())
		ctx.JSON(503, "error while parsing request")
		return
	}
	res, err := uc.userRepo.GetDivisions(ctx, &req)
	if err != nil {
		logger.Error("error while getting desination  names from userRepo ", err.Error())
		ctx.JSON(503, err.Error())
	}
	ctx.JSON(200, res)
}

// func (uc userController) GetStateDetails(ctx *gin.Context) {
// 	logger := logger2.GetLoggerWithContext(ctx)
// 	var req model.GetCountries
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		logger.Info("error while parsing request into state request", err.Error())
// 		ctx.JSON(503, "error while parsing request")
// 		return
// 	}
// 	res, err := uc.userRepo.GetStateDetails(ctx, &req)
// 	if err != nil {
// 		logger.Error("error while getting desination  names from userRepo ", err.Error())
// 		ctx.JSON(503, err.Error())
// 	}
// 	ctx.JSON(200, res)
// }

// func (uc userController) GetCityDetails(ctx *gin.Context) {
// 	logger := logger2.GetLoggerWithContext(ctx)
// 	var reqs model.GetState
// 	if err := ctx.ShouldBindJSON(&reqs); err != nil {
// 		logger.Info("error while parsing request into state request", err.Error())
// 		ctx.JSON(503, "error while parsing request")
// 		return
// 	}
// 	res, err := uc.userRepo.GetCityDetails(ctx, &reqs)
// 	if err != nil {
// 		logger.Error("error while getting desination  names from userRepo ", err.Error())
// 		ctx.JSON(503, err.Error())
// 	}
// 	ctx.JSON(200, res)
// }

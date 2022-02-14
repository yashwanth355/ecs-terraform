package controller

import (
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories/projectrepo"
	logger2 "usermvc/utility/logger"

	"github.com/gin-gonic/gin"
)

type ProjectManagementController interface {
	GetProjectDetail(ctx *gin.Context)
	GetAllProjectManagementDetails(ctx *gin.Context)
	DeleteProjectManagementDetail(ctx *gin.Context)
	UpdateProjectDetails(ctx *gin.Context)
}

type projectManagementController struct {
	projectrepo projectrepo.ProjectRepo
}

func newProjectManagementController() ProjectManagementController {
	return &projectManagementController{
		projectrepo: projectrepo.NewProjectRepo(),
	}
}

func (pc projectManagementController) GetProjectDetail(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var projectDetailRequest model.ProjectDetailRequest
	if err := ctx.ShouldBindJSON(&projectDetailRequest); err != nil {
		logger.Error("Error while parsing the project detail")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := pc.projectrepo.GetProjectDetail(ctx, projectDetailRequest.ProjectId)
	if err != nil {
		logger.Error("error while getting project detail ", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	logger.Info("getting response from get projectrepo.GetProjectDetail ", res)
	ctx.JSON(200, res)
}

func (pc projectManagementController) GetAllProjectManagementDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	res, err := pc.projectrepo.GetAllProjectManagementDetails(ctx)
	if err != nil {
		logger.Error("error while getting all project management details", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	logger.Info("getting response from get projectrepo.GetAllProjectManagementDetails ", res)
	ctx.JSON(200, res)
}

func (pc projectManagementController) UpdateProjectDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var projectDetailRequest model.ProjectDetailRequest
	if err := ctx.ShouldBindJSON(&projectDetailRequest); err != nil {
		logger.Error("Error while parsing the task details")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := pc.projectrepo.UpdateProjectDetails(ctx, ConvertProjectDetailsToProjectManagementMaster(projectDetailRequest))
	if err != nil {
		logger.Error("error while updating project detail", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	logger.Info("getting response from get projectrepo.UpdateProjectDetails ", res)
	ctx.JSON(200, res)
}

func (pc projectManagementController) DeleteProjectManagementDetail(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var getProjectDetailRequest model.ProjectDetailRequest
	if err := ctx.ShouldBindJSON(&getProjectDetailRequest); err != nil {
		logger.Error("Error while parsing the project details")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := pc.projectrepo.DeleteProjectManagementDetail(ctx, getProjectDetailRequest.ProjectId)
	if err != nil {
		logger.Error("error while deleting project management detail", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	logger.Info("getting response from get projectrepo.GetAllProjectManagementDetails ", res)
	ctx.JSON(200, res)
}

func ConvertProjectDetailsToProjectManagementMaster(projectDetails model.ProjectDetailRequest) entity.ProjectMangementMaster {
	return entity.ProjectMangementMaster{
		ProjectId:          projectDetails.ProjectId,
		ProjectIdsNo:       projectDetails.ProjectIdsNo,
		ProjectName:        projectDetails.ProjectName,
		ProjectOwner:       projectDetails.ProjectOwner,
		ProjectDate:        projectDetails.ProjectDate,
		DueDate:            projectDetails.DueDate,
		ProjectDescription: projectDetails.ProjectDescription,
		DepId:              projectDetails.DepId,
		FinancialYearId:    projectDetails.FinancialYearId,
		ProjectType:        projectDetails.ProjectType,
	}
}

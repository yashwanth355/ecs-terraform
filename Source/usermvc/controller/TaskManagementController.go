package controller

import (
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories/taskrepo"
	logger2 "usermvc/utility/logger"

	"github.com/gin-gonic/gin"
)

type TaskManagementController interface {
	GetAllTasksDetails(ctx *gin.Context)
	UpdateTaskDetails(ctx *gin.Context)
	GetTaskDetail(ctx *gin.Context)
	DeleteTaskDetail(ctx *gin.Context)
}

type taskManagementController struct {
	taskrepo taskrepo.TaskRepo
}

func newTaskManagementController() TaskManagementController {
	return &taskManagementController{
		taskrepo: taskrepo.NewTaskRepo(),
	}
}

func (tk taskManagementController) GetAllTasksDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	res, err := tk.taskrepo.GetAllTaskDetails(ctx)
	if err != nil {
		logger.Error("error while getting all task details", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	logger.Info("getting response from get taskrepo.GetAllTasksDetails ", res)
	ctx.JSON(200, res)
}

func (tk taskManagementController) GetTaskDetail(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var taskDetailRequest model.TaskDetailRequest
	if err := ctx.ShouldBindJSON(&taskDetailRequest); err != nil {
		logger.Error("Error while parsing the task detail")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := tk.taskrepo.GetTaskDetail(ctx, taskDetailRequest.TaskId)
	if err != nil {
		logger.Error("error while getting task detail", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	logger.Info("getting response from get taskrepo.GetTaskDetail ", res)
	ctx.JSON(200, res)
}

func (tk taskManagementController) UpdateTaskDetails(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var taskDetailRequest model.TaskDetailRequest
	if err := ctx.ShouldBindJSON(&taskDetailRequest); err != nil {
		logger.Error("Error while parsing the task details")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := tk.taskrepo.UpdateTaskDetails(ctx, ConvertTaskDetailsToTaskManagementMaster(taskDetailRequest))
	if err != nil {
		logger.Error("error while updating task details", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	logger.Info("getting response from get taskrepo.UpdateTaskDetails ", res)
	ctx.JSON(200, res)
}

func (tk taskManagementController) DeleteTaskDetail(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var getTaskDetailRequest model.TaskDetailRequest
	if err := ctx.ShouldBindJSON(&getTaskDetailRequest); err != nil {
		logger.Error("Error while parsing the task details")
		ctx.JSON(403, err.Error())
		return
	}
	res, err := tk.taskrepo.DeleteTaskDetail(ctx, getTaskDetailRequest.TaskId)
	if err != nil {
		logger.Error("error while deleting task detail", err.Error())
		ctx.JSON(503, err.Error())
		return
	}
	logger.Info("getting response from get taskrepo.DeleteTaskDetail ", res)
	ctx.JSON(200, res)
}

func ConvertTaskDetailsToTaskManagementMaster(taskDetails model.TaskDetailRequest) entity.TaskMangementMaster {
	return entity.TaskMangementMaster{
		TaskId:            taskDetails.TaskId,
		TaskIdsNo:         taskDetails.TaskIdsNo,
		ProjectId:         taskDetails.ProjectId,
		TaskName:          taskDetails.TaskName,
		TaskStart:         taskDetails.TaskStart,
		CloseStatus:       taskDetails.CloseStatus,
		TaskDescription:   taskDetails.TaskDescription,
		AssignedTo:        taskDetails.AssignedTo,
		Status:            taskDetails.Status,
		CancelStatus:      taskDetails.CancelStatus,
		StartingDate:      taskDetails.StartingDate,
		EndingDate:        taskDetails.EndingDate,
		TaskStartDatetime: taskDetails.TaskStartDatetime,
		CustomerId:        taskDetails.CustomerId,
		OriginId:          taskDetails.OriginId,
	}
}

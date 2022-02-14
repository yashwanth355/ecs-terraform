package taskrepo

import (
	"context"
	constants "usermvc"
	"usermvc/entity"
	"usermvc/repositories"
	logger2 "usermvc/utility/logger"

	"github.com/jinzhu/gorm"
)

type TaskRepo interface {
	GetAllTaskDetails(ctx context.Context) ([]*entity.ListTaskDetails, error)
	UpdateTaskDetails(ctx context.Context, updateTaskDetails entity.TaskMangementMaster) (string, error)
	GetTaskDetail(ctx context.Context, taskId string) (entity.TaskMangementMaster, error)
	DeleteTaskDetail(ctx context.Context, taskId string) (string, error)
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo() TaskRepo {
	newDb, err := repositories.NewDb()
	if err != nil {
		panic(err)
	}

	newDb.AutoMigrate(&entity.ProjectMangementMaster{})

	return &taskRepo{
		db: newDb,
	}
}

func (tr taskRepo) GetAllTaskDetails(ctx context.Context) ([]*entity.ListTaskDetails, error) {
	var result []*entity.ListTaskDetails
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all task details")
	err := tr.db.Table(constants.TaskManagementTableName).Select("taskid, originid, taskname, date(sdate) as sdate, status").Order("taskidsno desc").Scan(&result).Error
	if err != nil {
		logger.Error("error while get all task details from ", constants.TaskManagementTableName, err.Error())
		return nil, err
	}

	return result, nil
}

func (tr taskRepo) GetTaskDetail(ctx context.Context, taskId string) (entity.TaskMangementMaster, error) {
	var result entity.TaskMangementMaster
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch task detail")
	err := tr.db.Table(constants.TaskManagementTableName).Where("taskid=?", taskId).Find(&result).Error
	if err != nil {
		logger.Error("error while get task detail from ", constants.TaskManagementTableName, err.Error())
		return entity.TaskMangementMaster{}, err
	}

	return result, nil
}

func (tr taskRepo) UpdateTaskDetails(ctx context.Context, updateTaskDetails entity.TaskMangementMaster) (string, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to update task details")
	_, err := tr.db.Table(constants.TaskManagementTableName).Model(&entity.TaskMangementMaster{}).Where("taskid=?", updateTaskDetails.TaskId).Update(updateTaskDetails).Rows()
	if err != nil {
		logger.Error("error while updating record into", constants.TaskManagementTableName, err.Error())
		return "Unable to update record", err
	}
	return "Successfully updated", nil
}

func (tr taskRepo) DeleteTaskDetail(ctx context.Context, taskId string) (string, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to delete task details with having task id-", taskId)
	_, err := tr.db.Table(constants.TaskManagementTableName).Where("taskid=?", taskId).Delete(&entity.TaskMangementMaster{}).Rows()
	if err != nil {
		logger.Error("error while deleting the record from ", constants.TaskManagementTableName, err.Error())
		return "Error in deleting the record", err
	}

	return "Deleted successfully", nil
}

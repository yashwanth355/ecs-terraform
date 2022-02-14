package projectrepo

import (
	"context"
	constants "usermvc"
	"usermvc/entity"
	"usermvc/repositories"
	logger2 "usermvc/utility/logger"

	"github.com/jinzhu/gorm"
)

type ProjectRepo interface {
	GetProjectDetail(ctx context.Context, projectId string) (entity.ProjectMangementMaster, error)
	GetAllProjectManagementDetails(ctx context.Context) ([]*entity.ProjectManagementDetails, error)
	DeleteProjectManagementDetail(ctx context.Context, projectId string) (string, error)
	UpdateProjectDetails(ctx context.Context, updateprojectDetails entity.ProjectMangementMaster) (string, error)
}

type projectRepo struct {
	db *gorm.DB
}

func NewProjectRepo() ProjectRepo {
	newDb, err := repositories.NewDb()
	if err != nil {
		panic(err)
	}

	newDb.AutoMigrate(&entity.ProjectMangementMaster{})

	return &projectRepo{
		db: newDb,
	}
}

func (lr projectRepo) GetProjectDetail(ctx context.Context, projectId string) (entity.ProjectMangementMaster, error) {
	var result entity.ProjectMangementMaster
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch project detail")
	err := lr.db.Table(constants.ProjectManagementTableName).Where("projectid=?", projectId).Find(&result).Error
	if err != nil {
		logger.Error("error while get project detail from ", constants.TaskManagementTableName, err.Error())
		return entity.ProjectMangementMaster{}, err
	}

	return result, nil
}

func (lr projectRepo) GetAllProjectManagementDetails(ctx context.Context) ([]*entity.ProjectManagementDetails, error) {
	var result []*entity.ProjectManagementDetails
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all project details")
	err := lr.db.Table(constants.ProjectManagementTableName).Select("projectid, initcap(projectname) as projectname").Order("projectidsno").Scan(&result).Error
	if err != nil {
		logger.Error("error while get all project details from ", constants.ProjectManagementTableName, err.Error())
		return nil, err
	}

	return result, nil
}

func (lr projectRepo) DeleteProjectManagementDetail(ctx context.Context, projectId string) (string, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to delete project details with having project id-", projectId)
	_, err := lr.db.Table(constants.ProjectManagementTableName).Where("projectid=?", projectId).Delete(&entity.ProjectMangementMaster{}).Rows()
	if err != nil {
		logger.Error("error while deleting the record from ", constants.ProjectManagementTableName, err.Error())
		return "Error in deleting the record", err
	}

	return "Deleted successfully", nil
}

func (lr projectRepo) UpdateProjectDetails(ctx context.Context, updateprojectDetails entity.ProjectMangementMaster) (string, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to update project details")
	_, err := lr.db.Table(constants.ProjectManagementTableName).Model(&entity.ProjectMangementMaster{}).Where("projectid=?", updateprojectDetails.ProjectId).Update(updateprojectDetails).Rows()
	if err != nil {
		logger.Error("error while updating record into ", constants.ProjectManagementTableName, err.Error())
		return "Unable to update record", err
	}
	return "Successfully updated", nil
}

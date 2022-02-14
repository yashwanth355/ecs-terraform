package userrepo

import (
	"context"
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories"
	logger2 "usermvc/utility/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type UserRepo interface {
	Create(context context.Context, user entity.User) error
	GetAllUsersDetail(ctx context.Context) ([]*model.UserDetails, error)
	GetAllCompanyNames(ctx context.Context) ([]*model.CompanyNames, error)
	GetAllDepartmentName(ctx context.Context) ([]*model.DepartmentNames, error)
	GetAllDesignationName(ctx context.Context) ([]*model.DesignationName, error)
	GetDivisions(ctx context.Context, request *model.GetDivisionsRequest) ([]*model.Divisions, error)
	GetUserRole(ctx context.Context) ([]*model.UserRole, error)
	// GetCountryDetails(ctx context.Context) ([]*model.Countries, error)
	// GetStateDetails(ctx context.Context, request *model.Countries) ([]*model.States, error)
	// GetCityDetails(ctx context.Context, request *model.States) ([]*model.Cities, error)
	GetNameAndUsernameByUserId(ctx context.Context, userid string) (model.UserDetails, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo() UserRepo {
	newDb, err := repositories.NewDb()
	if err != nil {
		panic(err)
	}
	newDb.AutoMigrate(&entity.User{})
	return &userRepo{
		db: newDb,
	}
}

//func newDb() (*gorm.DB, error) {
//	//need to fix that
//	const (
//		host     = "ccl-psql-dev.cclxlbtddgmn.ap-south-1.rds.amazonaws.com"
//		port     = 5432
//		user     = "postgres"
//		password = "Ccl_RDS_DB#2022"
//		dbname   = "ccldevdb"
//	)
//	conf := Config.LoadConfig()
//	fmt.Println("prting the user and password", conf.DbConfig)
//	dsn := url.URL{
//		User:     url.UserPassword(user, password),
//		Scheme:   "postgres",
//		Host:     fmt.Sprintf("%s:%d", host, port),
//		Path:     dbname,
//		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
//	}
//	fmt.Println(password)
//	db, err := gorm.Open("postgres", dsn.String())
//	if err != nil {
//		return nil, err
//	}
//	return db, nil
//}

func (u userRepo) Create(context context.Context, user entity.User) error {

	logger := logger2.GetLoggerWithContext(context)
	tt := u.db.HasTable(&entity.User{})
	logger.Debug(tt)

	// db1 := u.db.Table("dbo.users_master_newpg").Model(&entity.User{}).Create(&user)
	// logger.Info(db1)
	ii := u.db.Table("dbo.testmk").Model(&entity.User{}).Value
	logger.Debug(ii)
	if err := u.db.Table("dbo.testmk").Model(&entity.User{}).Create(&user).Error; err != nil {
		return err
	}
	return nil
}
func (u userRepo) GetUserRole(ctx context.Context) ([]*model.UserRole, error) {
	var userrolelist []*model.UserRole
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("Going to fetch get user role")
	rows, err := u.db.Raw(`select rolename from dbo.cms_user_roles_master_newpg`).Rows()
	if err != nil {
		logger.Error("error while getting the companies details", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var users model.UserRole
		//	err = rows.Scan(&users.userName)
		userrolelist = append(userrolelist, &users)
	}
	return userrolelist, nil
}

func (u userRepo) GetAllUsersDetail(ctx context.Context) ([]*model.UserDetails, error) {
	var result []*model.UserDetails
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all userdetail from userdetails_master ")
	err := u.db.Table("userdetails_master").Model(&entity.UserDetails{}).Find(&result).Error
	if err != nil {
		logger.Error("error while getting the user details", err.Error)
		return nil, err
	}

	return result, nil

}

func (u userRepo) GetAllCompanyNames(ctx context.Context) ([]*model.CompanyNames, error) {
	var result []*model.CompanyNames
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all companies name")
	rows, err := u.db.Raw(`SELECT compname FROM dbo.company_master`).Rows()
	if err != nil {
		logger.Error("error while getting the companies details", err.Error)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var companyName model.CompanyNames
		err = rows.Scan(&companyName.Compname)
		result = append(result, &companyName)
	}
	return result, nil

}

func (u userRepo) GetAllDepartmentName(ctx context.Context) ([]*model.DepartmentNames, error) {
	var result []*model.DepartmentNames
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all Department name")
	rows, err := u.db.Raw(`SELECT deptname FROM dbo.department_master_newpg`).Rows()
	if err != nil {
		logger.Error("error while getting the department details", err.Error)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var companyName model.DepartmentNames
		err = rows.Scan(&companyName.Deptname)
		result = append(result, &companyName)
	}
	return result, nil

}

func (u userRepo) GetAllDesignationName(ctx context.Context) ([]*model.DesignationName, error) {
	var result []*model.DesignationName
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all Designation name")
	rows, err := u.db.Raw(`SELECT desgname FROM dbo.designation_master_newpg`).Rows()
	if err != nil {
		logger.Error("error while getting the Designation name from designation_master ", err.Error)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var DesigName model.DesignationName
		err = rows.Scan(&DesigName.DesgName)
		result = append(result, &DesigName)
	}
	return result, nil
}

func (u userRepo) GetDivisions(ctx context.Context, request *model.GetDivisionsRequest) ([]*model.Divisions, error) {
	var result []*model.Divisions
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all getDivisions with following params", result)
	rows, err := u.db.Raw(`SELECT divmaster FROM dbo.division_master_newpg where deptname=$1`, request.DepartmentName).Rows()
	if err != nil {
		logger.Error("error while getting the Designation name from designation_master ", err.Error)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var devision model.Divisions
		err = rows.Scan(&devision.Divmaster)
		result = append(result, &devision)
	}
	return result, nil
}

/*
*
 */
func (userRepoRef userRepo) GetNameAndUsernameByUserId(ctx context.Context, userid string) (model.UserDetails, error) {
	var userInfo model.UserDetails
	query := `select username, firstname, middlename, lastname from dbo.users_master_newpg where userid=$1`
	db := userRepoRef.db
	err := db.Raw(query, userid).Row().Scan(&userInfo.Username, &userInfo.Firstname, &userInfo.Middlename, &userInfo.Lastname)
	if err != nil {
		logger := logger2.GetLoggerWithContext(ctx)
		logger.Error("Error from func: GetNameAndUsernameByUserId in querying dbo.users_master_newpg ", err.Error())
		return model.UserDetails{}, err
	}
	return userInfo, err
}

package addressrepo

import (
	"context"
	// "errors"
	// "fmt"
	"github.com/jinzhu/gorm"

	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories"
	logger2 "usermvc/utility/logger"
)

type AddressRepo interface {
	GetCountries(ctx context.Context) (interface{}, error)
	GetStates(ctx context.Context, request model.Countries) (interface{}, error)
	GetCities(ctx context.Context, request model.States) (interface{}, error)
}

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo() AddressRepo {
	newDb, err := repositories.NewDb()
	if err != nil {
		panic(err)
	}
	newDb.AutoMigrate(&entity.User{})
	return &addressRepo{
		db: newDb,
	}
}
func (ad addressRepo) GetCountries(ctx context.Context) (interface{}, error) {
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("going to fetch all Country name")
	rows, err := ad.db.Raw(`SELECT countryid, countryname FROM dbo.countries_master`).Rows()
	if err != nil {
		logger.Error("error while getting the country details", err.Error)
		return nil, err
	}

	defer rows.Close()
	var countryList []model.Countries
	
	for rows.Next() {
		var cnty model.Countries
		err = rows.Scan(&cnty.CountryId, &cnty.CountryName)
		countryList = append(countryList, cnty)
	}
	return countryList, nil	
}

func (ad addressRepo) GetStates(ctx context.Context, request model.Countries) (interface{}, error) {
	
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("Going to fetch all states")
	sqlStatement1 := `select s.statename from dbo.states_master s inner join dbo.countries_master c on s.countryid=c.countryid where c.countryname=$1`

	rows, err := ad.db.Raw(sqlStatement1,&request.CountryName).Rows()
	if err != nil {
		logger.Error("Error while getting the state details", err.Error())
		return nil, err
	}
	var stateNames []model.States
	defer rows.Close()
	for rows.Next() {
		var st model.States
		err = rows.Scan(&st.StateName)
		stateNames = append(stateNames, st)
	}
	return stateNames, nil
}
func (ad addressRepo) GetCities(ctx context.Context, request model.States) (interface{}, error) {
	
	logger := logger2.GetLoggerWithContext(ctx)
	logger.Info("Going to fetch all states")
	sqlStatement1 := `select c.cityname from dbo.cities_master c inner join dbo.states_master s on c.stateid=s.stateid where s.statename=$1`

	rows, err := ad.db.Raw(sqlStatement1,&request.StateName).Rows()
	if err != nil {
		logger.Error("Error while getting the state details", err.Error())
		return nil, err
	}
	var cityNames []model.Cities
	defer rows.Close()
	for rows.Next() {
		var c model.Cities
		err = rows.Scan(&c.CityName)
		cityNames = append(cityNames, c)
	}
	return cityNames, nil
}

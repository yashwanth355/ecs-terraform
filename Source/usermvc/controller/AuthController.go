package controller

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"

	"encoding/json"
	logger2 "usermvc/utility/logger"
)

type authController struct {
}

const (
	host         = "ccl-psql-dev.cclxlbtddgmn.ap-south-1.rds.amazonaws.com"
	port         = 5432
	password     = "Ccl_RDS_DB#2022"
	dbname       = "ccldevdb"
	flowUserType = "USER_PASSWORD_AUTH"
)

type App struct {
	CognitoClient *cognito.CognitoIdentityProvider
	UserPoolID    string
	AppClientID   string
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Result struct {
	Id       *string `json:"id"`
	Role     string  `json:"role"`
	UserName string  `json:"user_name"`
	Userid   string  `json:"user_id"`
}

type AuthController interface {
	AdminConfirmSignUp(ctx *gin.Context)
	//Login(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func NewAuthController() AuthController {
	return &authController{}
}

func (auth authController) AdminConfirmSignUp(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var credentials Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		logger.Error("error while parsing the credentials", err.Error())
		ctx.JSON(403, err.Error())
		return
	}
	mySession := session.Must(session.NewSession())
	var cognitoRegion, cognitoUserPoolId, cognitoAppClientId string

	if os.Getenv("APP_ENV") == "production" {
		// 		cognitoRegion = globals.GetConfig().Cognito.CognitoRegion
		// 		cognitoUserPoolId = globals.GetConfig().Cognito.CognitoUserPoolId
		// 		cognitoAppClientId = globals.GetConfig().Cognito.CognitoAppClientId
		cognitoRegion = os.Getenv("AWS_COGNITO_REGION")
		cognitoUserPoolId = "us-east-2_P7fjWmPXI"
		cognitoAppClientId = "7rg8e53uej42du0ndlkf6cmvr6"
	} else {
		cognitoRegion = os.Getenv("AWS_COGNITO_REGION")
		cognitoUserPoolId = "us-east-2_P7fjWmPXI"
		cognitoAppClientId = "7rg8e53uej42du0ndlkf6cmvr6"
	}

	svc := cognitoidentityprovider.New(mySession, aws.NewConfig().WithRegion(cognitoRegion))

	cognitoClient := App{
		CognitoClient: svc,
		UserPoolID:    cognitoUserPoolId,
		AppClientID:   cognitoAppClientId,
	}

	adminSignUp := &cognito.AdminConfirmSignUpInput{
		UserPoolId: aws.String(cognitoClient.UserPoolID),
		Username:   aws.String(credentials.Username),
	}
	_, err := cognitoClient.CognitoClient.AdminConfirmSignUp(adminSignUp)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(503, "error while signup to account")
	}
}

func (auth authController) Login(ctx *gin.Context) {
	logger := logger2.GetLoggerWithContext(ctx)
	var credentials Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		logger.Error("error while parsing the credentials", err.Error())
		ctx.JSON(403, err.Error())
	}

	mySession := session.Must(session.NewSession())
	cognitoRegion := "ap-south-1"
	cognitoUserPoolId := "ap-south-1_eKFsULKaQ"
	cognitoAppClientId := "3u89hkmelq33cjnt9bsb2t95gp"

	svc := cognitoidentityprovider.New(mySession, aws.NewConfig().WithRegion(cognitoRegion))

	cognitoClient := App{
		CognitoClient: svc,
		UserPoolID:    cognitoUserPoolId,
		AppClientID:   cognitoAppClientId,
	}

	flow := aws.String(flowUserType)
	params := map[string]*string{
		"USERNAME": aws.String(credentials.Username),
		"PASSWORD": aws.String(credentials.Password),
	}

	authTry := &cognito.InitiateAuthInput{
		AuthFlow:       flow,
		AuthParameters: params,
		ClientId:       aws.String(cognitoClient.AppClientID),
	}

	response, err := cognitoClient.CognitoClient.InitiateAuth(authTry)
	if err != nil {
		log.Println(err.Error())
	}

	token := response.AuthenticationResult.AccessToken
	var result Result
	result.Id = token
	if credentials.Username != "" {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, "postgres", password, dbname)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Println(err)
		}
		defer db.Close()
		var rows *sql.Rows
		log.Println(credentials.Username)
		sqlStatement1 := `select u.role,u.userid,u.username from dbo.users_master_newpg u where u.emailid=$1`
		rows, _ = db.Query(sqlStatement1, credentials.Username)

		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&result.Role, &result.Userid, &result.UserName)
		}
	}

	log.Println(result)
	res, _ := json.Marshal(result)

	log.Println(res)

	ctx.JSON(200, result)

}

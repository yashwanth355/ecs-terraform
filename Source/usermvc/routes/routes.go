package routes

import (
	"time"
	"usermvc/controller"
	"usermvc/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	controller := controller.NewController()
	r := gin.Default()
	r.Use(middleware.LoggerMiddleWare())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT, POST, GET, DELETE, OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//r.Use(middleware.ApiAuth())
	//change the path here
	grp1 := r.Group("")

	{ //need to check with this
		//----------------USER-MODULE--------------------//
		grp1.POST("gin-dev-testenvLogin", controller.AuthController.Login)

		grp1.POST("gin-dev-AdminConfirmSignUp", controller.AuthController.AdminConfirmSignUp)

		grp1.GET("gin-dev-allquoteLineitems", controller.LeadControler.GetAllQuoteLineItems)
		grp1.GET("gin-dev-allquotes", controller.LeadControler.GetAllQuotes)
		grp1.GET("gin-dev-getQuotationCreationInfo", controller.LeadControler.GetQuotationCreationInfo)
		grp1.GET("gin-dev-getQuoteInformation", controller.LeadControler.GetQuoteInformation)

		grp1.GET("gin-dev-userDetails", controller.UserController.GetAllUsersDetail)
		grp1.GET("gin-dev-getCompanyNames", controller.UserController.GetAllCompanyNames)
		grp1.GET("gin-dev-DepartmentNames", controller.UserController.GetAllDepartmentName)
		grp1.GET("gin-dev-designationName", controller.UserController.GetAllDesignationName)
		grp1.GET("gin-dev-alldivisions", controller.UserController.GetDivisions)

		//grp1.POST("gin-dev-createuser", controller.UserController.CreateUser)

		grp1.GET("gin-dev-getAllLeads", controller.LeadControler.GetAllLeadsDetails)
		grp1.GET("gin-dev-getLeadDetails", controller.LeadControler.GetLeadDetails)
		grp1.GET("gin-dev-getLeadCreationInfo", controller.LeadControler.GetLeadCreationInfo)

		grp1.POST("gin-dev-getLeadsInfo", controller.LeadControler.GetLeadsInfo)
		grp1.POST("gin-dev-insertLeadDetails", controller.LeadControler.InsertLeadDetails)

		grp1.GET("gin-dev-getPOFormInfo", controller.PoController.GetPOFormInfo)
		// grp1.GET("gin-dev-listPurchaseOrders", controller.PoController.ListPurchaseOrders)
		//need to implement
		//grp1.POST("gin-dev-SendEmail", controller.PoController.SenEmail)

		//supplier
		grp1.GET("gin-dev-allaccount", controller.GetAllAccountDetails)

		//grp1.GET("gin-dev-viewsupplier", controller.GetSupplierDetails)

		grp1.POST("gin-dev-insertSupplierDetails", controller.SupplierController.InsertSupplierDetails)
		grp1.GET("gin-dev-viewSupplier", controller.SupplierController.ViewSupplierDetails)
		grp1.GET("gin-dev-listsupplier", controller.SupplierController.ListSupplierDetails)

		grp1.POST("gin-dev-insertGcDetails", controller.GCController.InsertGcDetails)
		grp1.GET("gin-dev-viewGcDetails", controller.GCController.ViewGcDetails)
		grp1.GET("gin-dev-listGcDetails", controller.GCController.ListGcDetails)

		grp1.GET("gin-dev-getorigin", controller.PoController.GetPortandOrigin)

		grp1.GET("gin-dev-getuserrole", controller.UserController.GetUserRole)

		// grp1.GET("gin-dev-getcountries", controller.UserController.GetCountryDetails)
		// grp1.GET("gin-dev-getstates", controller.UserController.GetStateDetails)
		// grp1.GET("gin-dev-getcities", controller.UserController.GetCityDetails)

		// All Project management Api Routes
		grp1.GET("gin-dev-getprojectdetail", controller.ProjectManagementController.GetProjectDetail)
		grp1.GET("gin-dev-allprojectmanagementdetails", controller.ProjectManagementController.GetAllProjectManagementDetails)
		grp1.DELETE("gin-dev-deleteprojectdetail", controller.ProjectManagementController.DeleteProjectManagementDetail)
		grp1.PATCH("gin-dev-updateprojectdetail", controller.ProjectManagementController.UpdateProjectDetails)

		// All Task Api Routes
		grp1.GET("gin-dev-gettaskdetail", controller.TaskManagementController.GetTaskDetail)
		grp1.GET("gin-dev-alltaskdetails", controller.TaskManagementController.GetAllTasksDetails)
		grp1.PATCH("gin-dev-updatetaskdetail", controller.TaskManagementController.UpdateTaskDetails)
		grp1.DELETE("gin-dev-deletetaskdetail", controller.TaskManagementController.DeleteTaskDetail)

		// Home page report apis
		grp1.GET("gin-dev-getconfirmedorders", controller.ReportsController.GetConfirmedOrders)
		//Address APIs
		grp1.GET("gin-dev-getCountry", controller.AddressController.GetCountries)
		grp1.POST("gin-dev-getState", controller.AddressController.GetStates)
		grp1.POST("gin-dev-getCity", controller.AddressController.GetCities)
		// Green Coffee Purchase Orders Module
		grp1.POST("gin-dev-getPoCreationInfo", controller.PoController.GetPoCreationInfo)
		grp1.POST("gin-dev-getPortandOriginForPO", controller.PoController.GetPortandOrigin)
		grp1.POST("gin-dev-getBalQuoteQtyForPoOrder", controller.PoController.GetBalQuoteQtyForPoOrder)
		grp1.POST("gin-dev-viewGCPODetails", controller.PoController.ViewGCPODetails)

	}

	return r
}

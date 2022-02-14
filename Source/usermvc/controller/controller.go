package controller

import (
	"usermvc/repositories/accountrepo"
	"usermvc/service/account"
	"usermvc/service/user"
)

type Controller interface {
}
type controller struct {
	userSvc                     user.Service
	accountSvc                  account.Service
	LeadControler               LeadController
	UserController              UserController
	AccountRepo                 accountrepo.AccountRepo
	AuthController              AuthController
	PoController                PoController
	SupplierController          SupplierController
	GCController                GCController
	ProjectManagementController ProjectManagementController
	TaskManagementController    TaskManagementController
	ReportsController           ReportsController
	AddressController           AddressController
}

func NewController() *controller {
	return &controller{userSvc: user.NewuserService(),
		accountSvc:                  account.NewAccountService(),
		LeadControler:               newLeadController(),
		UserController:              newUserController(),
		AccountRepo:                 accountrepo.NewAccountRepo(),
		SupplierController:          NewsupplierController(),
		AuthController:              NewAuthController(),
		PoController:                NewPoController(),
		GCController:                NewGcController(),
		ProjectManagementController: newProjectManagementController(),
		TaskManagementController:    newTaskManagementController(),
		ReportsController:           newReportsController(),
		AddressController:			 NewAddressController(),
	}
}

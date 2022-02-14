package controller
//update
import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"usermvc/model"
)

func (ctr *controller) InsertAccountDetails(c *gin.Context) {
	var accountDetailsReq *model.AccountDetailsRequest
	fmt.Println("getting error is the best")
	if err := c.ShouldBindJSON(&accountDetailsReq); err != nil {
		zap.S().Error("not able parse request", err.Error())
		c.JSON(200, err.Error())
		return
	}
	res, err := ctr.accountSvc.InsertAccountDetails(context.Background(), accountDetailsReq)

	if err != nil {
		zap.S().Error("not able parse request", err.Error())
		c.JSON(200, err.Error())
		return
	}
	c.JSON(200, res)
}

func (ctr controller) GetAllAccountDetails(c *gin.Context) {
	res, err := ctr.AccountRepo.GetAllAccountDetails(context.Background())

	if err != nil {
		zap.S().Error("error from the getappAccountDetails ", err.Error())
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, res)
}

//
//func (ctr controller) GetAllLeadAccounts(c *gin.Context)   {
//	res, err := ctr.accountSvc.GetAllAccountDetails(context.Background())
//	if err != nil {
//		zap.S().Error("error from the getappAccountDetails ", err.Error())
//		c.JSON(500, err.Error())
//		return
//	}
//	c.JSON(200, &model.GetAllAccountDetailsResponse{
//		Status:  200,
//		Payload: res,
//	})
//}
func validateRequest(ctx gin.Context) error {
	//write validation here
	return nil
}

func (ctr controller) TestPing(c *gin.Context) {
	c.Status(200)
}

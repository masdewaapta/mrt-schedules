package station

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masdewaapta/mrt-schedules/common/response"
)

func Initiate(router *gin.RouterGroup) {

	stationService := NewService()

	station := router.Group("/stations")
	station.GET("", func(ctx *gin.Context) {
		GetAllStation(ctx, stationService)

	})

}

func GetAllStation(ctx *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Successfully get all station",
		Data:    datas,
	})
}

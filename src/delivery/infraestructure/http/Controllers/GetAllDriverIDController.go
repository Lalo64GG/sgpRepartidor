package controllers

import(
	"github.com/lalo64/sgp/src/delivery/application"
	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/shared/responses"
	"net/http"
	"strconv"
	
)

type GetAllDriverIDController struct {
	GetAllDriverIDUseCase *application.GetAllDriverIDUseCase
}

func NewGetAllDriverIDController(getAllDriverUseCase *application.GetAllDriverIDUseCase) *GetAllDriverIDController{
	return &GetAllDriverIDController{GetAllDriverIDUseCase: getAllDriverUseCase}
}

func(ctr *GetAllDriverIDController) Run(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid ID",
			Error: err.Error(),
			Data: nil,
		})
		return
	}

	deliveries, err := ctr.GetAllDriverIDUseCase.Run(id)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Internal Server Error",
			Error: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, responses.Response{
		Success: true,
		Message: "Deliveries Found",
		Error: nil,
		Data: deliveries,
	})
}
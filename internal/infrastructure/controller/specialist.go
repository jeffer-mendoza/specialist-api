package controller

import (
	"app/internal/application"
	"app/internal/domain/model"
	"app/internal/infrastructure/container"

	"github.com/jeffer-mendoza/go-common-lib/pkg/errors"
	"github.com/jeffer-mendoza/go-common-lib/pkg/web"

	"github.com/gin-gonic/gin"

	"net/http"
)

type Specialist interface {
	Create(ctx *gin.Context)
	Find(ctx *gin.Context)
	List(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type specialistImpl struct {
	container *container.Dependency
}

func NewSpecialist(container *container.Dependency) Specialist {
	return &specialistImpl{
		container: container,
	}
}


func (controller *specialistImpl) Create(ctx *gin.Context) {
	request := model.Specialist{}
	var response model.Specialist

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		err = errors.NewBadRequestApiError(err.Error())
	} else {
		response, err = controller.container.SpecialistApplication.
		(application.Specialist).Create(ctx, request)
	}

	web.ReturnInfo(ctx, response, err, http.StatusOK)
}

func (controller *specialistImpl) Find(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := controller.container.SpecialistApplication.
		(application.Specialist).Find(ctx, id)

	web.ReturnInfo(ctx, response, err, http.StatusOK)
}

func (controller *specialistImpl) List(ctx *gin.Context) {
	response, err := controller.container.SpecialistApplication.
	(application.Specialist).List(ctx)

	web.ReturnInfo(ctx, response, err, http.StatusOK)
}

func (controller *specialistImpl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := controller.container.SpecialistApplication.
	(application.Specialist).Delete(ctx, id)

	web.ReturnInfo(ctx, nil, err, http.StatusNoContent)
}


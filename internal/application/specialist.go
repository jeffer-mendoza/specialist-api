package application

import (
	"app/internal/domain/model"
	"app/internal/domain/port"
	"app/internal/infrastructure/container"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Specialist interface {
	Create(ctx *gin.Context, request model.Specialist) (model.Specialist, error)
	Find(ctx *gin.Context, id string) (model.Specialist, error)
	List(ctx *gin.Context) ([]model.Specialist, error)
	Delete(ctx *gin.Context, id string) error
}

func NewSpecialist(container *container.Dependency) Specialist {
	return &specialistImpl{
		container: container,
	}
}

type specialistImpl struct {
	container *container.Dependency
}

func (s *specialistImpl) Create(ctx *gin.Context,
	request model.Specialist) (model.Specialist, error) {
	request.ID = uuid.New().String()

	err := s.container.SpecialistDB.(port.Specialist).Create(ctx, request)

	return request, err
}

func (s *specialistImpl) Find(ctx *gin.Context, id string) (model.Specialist, error) {
	return s.container.SpecialistDB.(port.Specialist).Find(id)

}

func (s *specialistImpl) List(ctx *gin.Context) ([]model.Specialist, error) {
	return s.container.SpecialistDB.(port.Specialist).FindAll()
}

func (s *specialistImpl) Delete(ctx *gin.Context, id string) error {
	return s.container.SpecialistDB.(port.Specialist).Delete(id)
}

package port

import (
	"app/internal/domain/model"
	"github.com/gin-gonic/gin"
)

type Specialist interface {
	Create(ctx *gin.Context, specialist model.Specialist) error
	Find(ID string) (model.Specialist, error)
	FindAll() ([]model.Specialist, error)
	Delete(ID string) error
}

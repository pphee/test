package rest

import (
	"denti/pkg/logger"
	"denti/pkg/user"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type userController struct {
	log logger.LogInfoFormat
	svc user.Service
}

func NewUserController(log logger.LogInfoFormat, svc user.Service) *userController {
	return &userController{log, svc}
}

func (u *userController) GetAll(ctx *gin.Context) {
	users, err := u.svc.GetAll()
	if len(users) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *userController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	user, err := u.svc.GetByID(id)
	if user == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *userController) Store(ctx *gin.Context) {
	var user user.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.svc.Store(&user)
	ctx.Status(http.StatusCreated)
}

func (u *userController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var user user.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	u.svc.Update(&user)
	ctx.Status(http.StatusOK)
}

func (u *userController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	u.svc.Delete(id)
	ctx.Status(http.StatusNoContent)
}
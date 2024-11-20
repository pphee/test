package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pphee/test/pkg/bmi"
	"net/http"
)

type BmiService interface {
	StoreBMI(ctx context.Context) ([]*bmi.BMI, error)
	QueryBMI(ctx context.Context, queryVector []float32) ([]*bmi.BMIIV, error)
}

type BmiCtrl struct {
	svc bmi.BmiService
}

func NewBmiCtrls(svc BmiService) *BmiCtrl {
	return &BmiCtrl{svc: svc}
}

func (h *BmiCtrl) StoreBMI(ctx *gin.Context) {
	bmiRecord, err := h.svc.StoreBMI(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, bmiRecord)
}

func (h *BmiCtrl) QueryBMI(ctx *gin.Context) {
	var req struct {
		QueryVector []float32 `json:"query_vector"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	results, err := h.svc.QueryBMI(ctx.Request.Context(), req.QueryVector)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}

package server

import (
	"github.com/akbaralishaikh/denti/pkg/login"
	"github.com/akbaralishaikh/denti/pkg/user"
	"github.com/gin-gonic/gin"
	"github.com/pphee/test/pkg/bmi"
	rest "github.com/pphee/test/pkg/http/rest"
)

func (ds *dserver) MapRoutes() {

	// Group : v1
	apiV1 := ds.router.Group("api/v1")

	ds.healthRoutes(apiV1)
	ds.loginRoutes(apiV1)
	ds.userRoutes(apiV1)
}

func (ds *dserver) healthRoutes(api *gin.RouterGroup) {
	healthRoutes := api.Group("/health")
	{
		h := rest.NewHealthCtrl()
		healthRoutes.GET("/", h.Ping)
	}
}

func (ds *dserver) loginRoutes(api *gin.RouterGroup) {
	var loginSvc login.Service
	ds.cont.Invoke(func(l login.Service) {
		loginSvc = l
	})

	loginRoutes := api.Group("/login")
	{
		f := rest.NewLoginCtrl(ds.logger, loginSvc)
		loginRoutes.POST("/", f.Signin)
	}
}

func (ds *dserver) userRoutes(api *gin.RouterGroup) {
	userRoutes := api.Group("/users")
	{
		var userSvc user.Service
		ds.cont.Invoke(func(u user.Service) {
			userSvc = u
		})

		usr := rest.NewUserCtrl(ds.logger, userSvc)

		userRoutes.GET("/", usr.GetAll)
		userRoutes.POST("/", usr.Store)
		userRoutes.GET("/:id", usr.GetByID)
		userRoutes.PUT("/:id", usr.Update)
		userRoutes.DELETE("/:id", usr.Delete)
	}
}

func (ds *dserver) bmiRoutes(api *gin.RouterGroup) {
	bmiRoutes := api.Group("/bmi")
	{
		var bmiSvc bmi.BmiService
		b := rest.NewBmiCtrls(bmiSvc)
		bmiRoutes.POST("/", b.StoreBMI)
		bmiRoutes.GET("/query", b.QueryBMI)
	}
}

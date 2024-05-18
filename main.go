package main

import (
	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/component/gormc"
	"log"
	"my-shop/builder"
	"my-shop/common"
	"my-shop/component"
	"my-shop/middleware"
	productcontroller "my-shop/module/product/controller"
	productusecase "my-shop/module/product/domain/usecase"
	productmysql "my-shop/module/product/repository/mysql"
	userhttpservice "my-shop/module/user/infras/httpservice"
	userrepository "my-shop/module/user/infras/repository"
	userusecase "my-shop/module/user/usecase"
	"net/http"
)

func newService() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("my_shop"),
		sctx.WithComponent(gormc.NewGormDB(common.KeyGorm, "")),
		sctx.WithComponent(component.NewJWT(common.KeyJWT)),
		sctx.WithComponent(component.NewAWSS3Provider(common.KeyAWSS3)),
	)
}

func main() {

	service := newService()

	service.OutEnv()

	if err := service.Load(); err != nil {
		log.Fatalln(err)
	}

	db := service.MustGet(common.KeyGorm).(common.DbContext).GetDB()

	r := gin.Default()

	r.Use(middleware.Recovery())

	tokenProvider := service.MustGet(common.KeyJWT).(component.TokenProvider)
	authClient := userusecase.NewIntrospectUC(userrepository.NewUserRepo(db), userrepository.NewSessionMySQLRepo(db), tokenProvider)

	r.GET("/ping", middleware.RequireAuth(authClient), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.DELETE("/v1/revoke-token", middleware.RequireAuth(authClient), func(c *gin.Context) {
		requester := c.MustGet(common.KeyRequester).(common.Requester)

		repo := userrepository.NewSessionMySQLRepo(db)
		if err := repo.Delete(c.Request.Context(), requester.TokenId()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	})

	// Setup dependencies
	repo := productmysql.NewMysqlRepository(db)
	useCase := productusecase.NewCreateProductUseCase(repo)
	api := productcontroller.NewAPIController(useCase)

	v1 := r.Group("/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", api.CreateProductAPI(db))
		}
	}

	//userUC := userusecase.NewUseCase(userrepository.NewUserRepo(db), userrepository.NewSessionMySQLRepo(db), &common.Hasher{}, tokenProvider)

	userUseCase := userusecase.UseCaseWithBuilder(builder.NewSimpleBuilder(db, tokenProvider))

	userhttpservice.NewUserService(userUseCase, service).Routes(v1)

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

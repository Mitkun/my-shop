package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"my-shop/common"
	productcontroller "my-shop/module/product/controller"
	productusecase "my-shop/module/product/domain/usecase"
	productmysql "my-shop/module/product/repository/mysql"
	userhttpservice "my-shop/module/user/infras/httpservice"
	"my-shop/module/user/infras/repository"
	userusecase "my-shop/module/user/usecase"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
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

	userUC := userusecase.NewUseCase(repository.NewUserRepo(db), &common.Hasher{})
	userhttpservice.NewUserService(userUC).Routes(v1)

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

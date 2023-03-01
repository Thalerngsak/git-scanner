package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/thalerngsak/git-scanner/handler"
	"github.com/thalerngsak/git-scanner/repository"
	"github.com/thalerngsak/git-scanner/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

func main() {
	initTimeZone()
	initConfig()
	client := initDatabase()
	db := client.Database(viper.GetString("db.database"))
	// Initialize repository and result stores
	repositoryCollection := db.Collection("repositories")
	resultCollection := db.Collection("results")

	repositoryStore := repository.NewRepositoryDB(repositoryCollection)
	repositoryService := service.NewRepositoryService(&repositoryStore)
	repositoryHandler := handler.NewRepositoryHandler(repositoryService)

	resultStoreDB := repository.NewResultDB(resultCollection)
	_ = resultStoreDB
	//resultStoreMock := repository.NewResultRepositoryMock()

	resultService := service.NewResultService(&resultStoreDB)
	resultHandler := handler.NewResultHandler(resultService)

	scanHandler := handler.NewScanHandler(repositoryService, resultService)

	router := gin.Default()
	// Endpoint to create a repository
	router.POST("/repositories", repositoryHandler.NewRepository)

	// Endpoint to get all repositories
	router.GET("/repositories", repositoryHandler.GetRepository)
	// Endpoint to get a repository by ID
	router.GET("/repositories/:id", repositoryHandler.GetRepositoryByID)

	// Endpoint to update a repository
	router.PUT("/repositories/:id", repositoryHandler.UpdateRepository)

	// Endpoint to delete a repository
	router.DELETE("/repositories/:id", repositoryHandler.DeleteRepository)

	// Endpoint to trigger a scan
	router.POST("/scan/:id", scanHandler.Scan)

	// Endpoint to get all results
	router.GET("/results", resultHandler.GetResult)

	// Endpoint to get results by repository ID
	router.GET("/results/repository/:id", resultHandler.GetResultByRepositoryID)

	router.Run(fmt.Sprintf(":%v", viper.GetInt("app.port")))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *mongo.Client {
	// Initialize MongoDB client and database "mongodb://root:example@localhost:27017"
	uri := fmt.Sprintf("%v://%v:%v@%v:%v",
		viper.GetString("db.driver"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
	)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}

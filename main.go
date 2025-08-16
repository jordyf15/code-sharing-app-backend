package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jordyf15/code-sharing-app/controllers"
	"github.com/jordyf15/code-sharing-app/middlewares"
	sr "github.com/jordyf15/code-sharing-app/snippet/repository"
	su "github.com/jordyf15/code-sharing-app/snippet/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbClient *mongo.Database
	router   *gin.Engine
)

func initializeRoutes() {
	snippetRepo := sr.NewSnippetRepository(dbClient)

	snippetUsecase := su.NewSnippetUsecase(snippetRepo)

	snippetController := controllers.NewSnippetController(snippetUsecase)

	router.GET("_health", health)

	router.POST("snippets", snippetController.CreateSnippet)
	router.GET("snippets/:snippet_id", snippetController.GetSnippet)
	router.PATCH("snippets/:snippet_id", snippetController.UpdateSnippet)
}

func connectToDB() {
	connectionURL := os.Getenv("DB_URL")
	clientOptions := options.Client().ApplyURI(connectionURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	dbClient = client.Database(os.Getenv("DB_NAME"))
}

func health(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	connectToDB()
}

func main() {
	router = gin.Default()

	trustedProxies := strings.Split(os.Getenv("TRUSTED_PROXIES"), ",")
	if len(trustedProxies) > 0 {
		router.SetTrustedProxies(trustedProxies)
	}

	if allowedOriginsEnvValue := os.Getenv("ALLOWED_ORIGINS"); len(allowedOriginsEnvValue) > 0 {
		allowedOrigins := strings.Split(allowedOriginsEnvValue, ",")
		config := cors.DefaultConfig()
		config.AllowOrigins = allowedOrigins
		config.AllowHeaders = []string{"Origin", "Authorization"}

		router.Use(cors.New(config))
	}

	loggerMiddleware := middlewares.NewLoggerMiddleware()

	if gin.IsDebugging() {
		router.Use(loggerMiddleware.PrintHeadersAndFormParams)
	}

	router.MaxMultipartMemory = 10 << 20
	initializeRoutes()
	if os.Getenv("ROUTER_PORT") != "" {
		router.Run(fmt.Sprintf(":%s", os.Getenv("ROUTER_PORT")))
	} else {
		router.Run()
	}
}

package main

import (
    "github.com/gin-gonic/gin"
    "interview_YangYang_20241010/handlers"
    "interview_YangYang_20241010/repository"
    _ "interview_YangYang_20241010/docs"

    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerfiles "github.com/swaggo/files"
)

func main() {
    // init db
    repository.InitDB()

    router := gin.Default()

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

    // set player management route
    players := router.Group("/players")
    {
        players.GET("", handlers.GetPlayers)
        players.POST("", handlers.CreatePlayer)
        players.GET("/:id", handlers.GetPlayerByID)
        players.PUT("/:id", handlers.UpdatePlayer)
        players.DELETE("/:id", handlers.DeletePlayer)
    }

    // Set up level management routes
    levels := router.Group("/levels")
    {
        levels.GET("", handlers.GetLevels)
        levels.POST("", handlers.CreateLevel)
    }

    // Set up room management routes
    rooms := router.Group("/rooms")
    {
        rooms.GET("", handlers.GetRooms)
        rooms.POST("", handlers.CreateRoom)
        rooms.GET("/:id", handlers.GetRoomByID)
        rooms.PUT("/:id", handlers.UpdateRoom)
        rooms.DELETE("/:id", handlers.DeleteRoom)
    }

    // Set up reservation management routes
    reservations := router.Group("/reservations")
    {
        reservations.GET("", handlers.GetReservations)
        reservations.POST("", handlers.CreateReservation)
    }

    // Set up challenge management routes (new)
    challenges := router.Group("/challenges")
    {
        challenges.POST("", handlers.ParticipateChallenge)
        challenges.GET("/results", handlers.GetChallengeResults)
    }

	// Set up log management routes (new)
	logs := router.Group("/logs")
	{
		logs.GET("", handlers.GetLogs)
		logs.POST("", handlers.CreateLog)
	}

	// Set up payment management routes (new)
	payments := router.Group("/payments")
	{
		payments.POST("", handlers.ProcessPayment)
		payments.GET("/:id", handlers.GetPaymentDetails)
	}

    // start server, listen 8080 port
    router.Run(":8080")
}
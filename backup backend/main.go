package main

import (
	"github.com/Topzzson/SA-PROJECT/controller"

	"github.com/Topzzson/SA-PROJECT/entity"

	"github.com/Topzzson/SA-PROJECT/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// User Routes
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			protected.PATCH("/users", controller.UpdateUser)
			protected.DELETE("/users/:id", controller.DeleteUser)

			protected.GET("/path_statuses", controller.ListStatus)
			protected.GET("/path_status/:id", controller.GetStatus)
			protected.POST("/path_statuses", controller.CreateStatus)
			protected.PATCH("/path_statuses", controller.UpdateStatus)
			protected.DELETE("/path_statuses/:id", controller.DeleteStatus)

			protected.GET("/ambulances", controller.ListAmbulance)
			protected.GET("/ambulance/:id", controller.GetAmbulance)
			protected.POST("/ambulances", controller.CreateAmbulance)
			protected.PATCH("/ambulances", controller.UpdateAmbulance)
			protected.DELETE("/ambulances/:id", controller.DeleteAmbulance)

			protected.GET("/check_lists", controller.ListCheckList)
			protected.GET("/check_list/:id", controller.GetCheckList)
			protected.POST("/check_lists", controller.CreateCheckList)
			protected.PATCH("/check_lists", controller.UpdateCheckList)
			protected.DELETE("/check_lists/:id", controller.DeleteCheckList)

			protected.GET("/car_paths", controller.ListCarPath)
			protected.GET("/car_path/:id", controller.GetCarPath)
			protected.POST("/car_paths", controller.CreateCarPath)
			protected.PATCH("/car_paths", controller.UpdateCarPath)
			protected.DELETE("/car_paths/:id", controller.DeleteCarPath)

		}
	}

	// User Routes
	r.POST("/users", controller.CreateUser)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

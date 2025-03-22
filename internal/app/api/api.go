package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/vitalicher97/psychologist_app/internal/app/middleware"
)

// SetupRoutes initializes the routes for the API.
func SetupRoutes(r *gin.Engine) {
	gin.SetMode(gin.DebugMode)

	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.ErrorHandler())

	apiRouter := r.Group("/api/psychotherapy")

	psychologist := apiRouter.Group("psychologists")
	psychologist.GET("", GetAllPsychologists)
	psychologist.GET(":id", GetPsychologist)
	psychologist.POST("", CreatePsychologist)
	psychologist.PUT(":id", UpdatePsychologist)
	psychologist.DELETE(":id", DeletePsychologist)

	availability := apiRouter.Group("availabilities")
	availability.GET("", GetAllAvailability)
	availability.GET(":id", GetAvailability)
	availability.POST("", CreateAvailability)
	availability.PUT(":id", UpdateAvailability)
	availability.DELETE(":id", DeleteAvailability)

	consultationPricing := apiRouter.Group("consultation-pricings")
	consultationPricing.GET("", GetAllConsultationPricing)
	consultationPricing.GET(":id", GetConsultationPricing)
	consultationPricing.POST("", CreateConsultationPricing)
	consultationPricing.PUT(":id", UpdateConsultationPricing)
	consultationPricing.DELETE(":id", DeleteConsultationPricing)

	appointments := apiRouter.Group("appointments")
	appointments.GET("", GetAllAppointments)
	appointments.GET(":id", GetAppointment)
	appointments.POST("", CreateAppointment)
	appointments.PUT(":id", UpdateAppointment)
	appointments.DELETE(":id", DeleteAppointment)

	customer := apiRouter.Group("customers")
	customer.GET("", GetAllCustomers)
	customer.GET(":id", GetCustomer)
	customer.POST("", CreateCustomer)
	customer.PUT(":id", UpdateCustomer)
	customer.DELETE(":id", DeleteCustomer)

	CustomerPsychologistPrices := apiRouter.Group("customer-psychologist-prices")
	CustomerPsychologistPrices.GET("", GetCustomerPsychologistPrices)
	CustomerPsychologistPrices.POST("", CreateCustomerPsychologistPrices)
}

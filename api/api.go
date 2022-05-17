package api

import (
	"apiaive/api/controller"
	"apiaive/api/model"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Handlers() *gin.Engine {
	router := gin.Default()
	controller.InitDB()
	userRoute := router.Group("api/users")
	{
		userRoute.GET("", GetUsers)
		userRoute.POST("", PostAppointment)
	}

	vaccinationCenterRoute := router.Group("api/vaccination-center")
	{
		vaccinationCenterRoute.GET("", GetVaccinationCenters)
	}

	appointmentRoute := router.Group("api/appointment")
	{
		appointmentRoute.GET("", GetAppointments)
		//A sécuriser
		appointmentRoute.GET(":vcId", GetAppointmentsByVcId)
	}

	return router
}

func GetUsers(c *gin.Context) {
	users := controller.GetUsers()
	c.JSON(200, &users)
}

func GetVaccinationCenters(c *gin.Context) {
	vCenters := controller.GetVaccinationCenters()
	c.JSON(200, &vCenters)
}

func PostAppointment(c *gin.Context) {
	var jsonAppointment model.Appointment
	c.Bind(&jsonAppointment)
	fmt.Print(jsonAppointment.Date)
	appointment, err := controller.PostAppointment(&jsonAppointment)
	if err != nil {
		// Affichage de l'erreur
		c.JSON(422, gin.H{"error": "Fields are empty"})
	} else {
		// Affichage des données saisies
		c.JSON(201, gin.H{"success": appointment})
	}
}

func GetAppointments(c *gin.Context) {
	appointments := controller.GetAppointments()
	c.JSON(200, appointments)
}

func GetAppointmentsByVcId(c *gin.Context) {
	vcId := c.Params.ByName("vcId")
	if vcId != "" {
		apppointments := controller.GetAppointmentsByCenterId(vcId)
		c.JSON(200, apppointments)
	} else {
		c.JSON(404, gin.H{"error": "Center id not provided"})
	}
}

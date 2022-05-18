package api

import (
	"apiaive/api/controller"
	"apiaive/api/model"
	"fmt"
	"time"

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
		appointmentRoute.GET(":vcid", GetAppointmentsByVcId)
		appointmentRoute.GET(":vcid/:date", GetAppointmentsAvaibles)
	}

	tokenRoute := router.Group("api/token")
	{
		tokenRoute.GET(":token", CheckToken)
	}

	return router
}

func CheckToken(c *gin.Context) {
	generatedToken := c.Params.ByName("token")
	check, err := controller.ControlToken(generatedToken)
	if err != nil {
		c.JSON(403, gin.H{"error": err})
	} else {
		c.JSON(200, gin.H{"success": check})
	}
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
	vcId := c.Params.ByName("vcid")
	if vcId != "" {
		apppointments := controller.GetAppointmentsByCenterId(vcId)
		c.JSON(200, apppointments)
	} else {
		c.JSON(404, gin.H{"error": "vcid not provided"})
	}
}

func GetAppointmentsAvaibles(c *gin.Context) {
	date := c.Params.ByName("date")
	fmt.Println(date)
	vcId := c.Params.ByName("vcid")
	if date != "" && vcId != "" {
		t, _ := time.Parse("2006-01-02T15:04:05.000Z", date)
		fmt.Println(t)
		appointments := controller.GetAppointmentsAvaibles(vcId, t)
		c.JSON(200, gin.H{"success": appointments})
	} else {
		c.JSON(404, gin.H{"error": "date not provided"})
	}
}

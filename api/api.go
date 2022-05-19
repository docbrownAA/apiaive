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

	vaccinationCenterRoute := router.Group("api/vaccination-center")
	{
		vaccinationCenterRoute.GET("", GetVaccinationCenters)
	}

	appointmentRoute := router.Group("api/appointment")
	{
		appointmentRoute.POST("", CreateAppointment)
		appointmentRoute.GET("", GetAppointments)
		appointmentRoute.GET(":vcid/:date", GetAppointmentAvailables)
		//A sécuriser
		appointmentRoute.GET(":vcid", GetAppointmentsByVcId)
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
		fmt.Println(err)
		c.JSON(403, gin.H{"error": err.Error()})
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

func CreateAppointment(c *gin.Context) {
	var jsonAppointment model.Appointment
	c.Bind(&jsonAppointment)
	appointment, token, err := controller.CreateAppointment(&jsonAppointment)
	if err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
	} else {
		//url just to have access quickly to the validation
		c.JSON(201, gin.H{"success": appointment, "token": "http://localhost:3000/api/token/" + token.String()})
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

func GetAppointmentAvailables(c *gin.Context) {
	date := c.Params.ByName("date")
	fmt.Println(date)
	vcId := c.Params.ByName("vcid")
	if date != "" && vcId != "" {
		t, _ := time.Parse("2006-01-02T15:04:05.000Z", date)
		fmt.Println(t)
		appointments := controller.GetAppointmentAvailables(vcId, t)
		c.JSON(200, gin.H{"success": appointments})
	} else {
		c.JSON(404, gin.H{"error": "date not provided"})
	}
}

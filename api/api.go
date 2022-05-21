package api

import (
	"apiaive/api/controller"
	"apiaive/api/middlewares"
	"apiaive/api/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Handlers() *gin.Engine {
	router := gin.Default()
	controller.InitDB()

	userRoute := router.Group("api/user")
	{
		userRoute.POST("/signup", registerUser)
		userRoute.POST("/signin", GenerateToken)
	}
	tokenRoute := router.Group("api/token")
	{
		tokenRoute.GET(":token", checkToken)
	}

	vaccinationCenterRoute := router.Group("api/vaccination-center")
	{
		vaccinationCenterRoute.GET("", getVaccinationCenters)
	}

	appointmentRoute := router.Group("api/appointment")
	{
		appointmentRoute.POST("", createAppointment)
		appointmentRoute.GET(":vcid/:date", getAppointmentAvailables)
	}

	adminRoute := router.Group("api/admin/appointment").Use(middlewares.Auth())
	{
		adminRoute.GET("", getAppointmentsByVcId)
	}

	return router
}

// Check the token that was given to the user when he created an appointment
func checkToken(c *gin.Context) {
	generatedToken := c.Params.ByName("token")
	check, err := controller.ControlToken(generatedToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": check})
	}
}

// List Vaccination centers
func getVaccinationCenters(c *gin.Context) {
	vCenters := controller.GetVaccinationCenters()
	c.JSON(200, &vCenters)
}

// Create an appointment.
// Return the link with the token for the user that he has to click
func createAppointment(c *gin.Context) {
	var jsonAppointment model.Appointment
	c.Bind(&jsonAppointment)
	appointment, token, err := controller.CreateAppointment(&jsonAppointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		//url just to have access quickly to the validation
		c.JSON(http.StatusCreated, gin.H{"success": appointment, "token": "http://localhost:3000/api/token/" + token.String()})
	}
}

func getAppointmentsByVcId(c *gin.Context) {
	fmt.Println(c.Request.Header["Authorization"][0])
	username := c.Request.Header["Username"][0]
	if username != "" {
		apppointments, err := controller.GetAppointmentsByCenterId(username)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		c.JSON(200, apppointments)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not provided"})
	}
}

func getAppointmentAvailables(c *gin.Context) {
	date := c.Params.ByName("date")
	fmt.Println(date)
	vcId := c.Params.ByName("vcid")
	if date != "" && vcId != "" {
		t, _ := time.Parse("2006-01-02T15:04", date)
		fmt.Println("api:", t)
		appointments := controller.GetAppointmentAvailables(vcId, t)
		c.JSON(200, gin.H{"success": appointments})
	} else {
		c.JSON(404, gin.H{"error": "date not provided"})

	}
}

// Create a user with UserName, email, password,and vaccination center id
func registerUser(c *gin.Context) {
	var user model.User
	c.Bind(&user)
	fmt.Println(user)
	err := controller.RegisterUser(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "ok"})
}

// Generate a bearer token to authenticate user
func GenerateToken(c *gin.Context) {
	var request model.TokenRequest
	c.Bind(&request)
	tokenString, err := controller.GeneratedToken(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}

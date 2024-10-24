// PATH: auth-golang/routes/auth.go

package routes

import (
    "auth-golang/controllers"

    "github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)
    r.POST("/signup", controllers.Signup)
    r.GET("/home", controllers.Home)
    r.GET("/premium", controllers.Premium)
    r.GET("/logout", controllers.Logout)
}

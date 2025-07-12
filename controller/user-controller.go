package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/thampaponn/learn-go/initializers"
	"github.com/thampaponn/learn-go/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	//Get username/password from req body
	var body struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})

		return
	}

	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	//Create user
	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Username:  body.Username,
		Password:  string(hash),
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	//Return
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(ctx *gin.Context) {
	//Get username/password from req body
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	//Check if user exists
	var user models.User
	initializers.DB.First(&user, "username = ?", body.Username)
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid username or password"})
		return
	}

	//Check if password is correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	//Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	//Return token
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{})
}

func Validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "You're logged in",
		"user":    user,
	})
}

func DeleteUser(ctx *gin.Context) {
	//Get id from url
	id := ctx.Param("id")

	//Check if user exists
	if err := initializers.DB.First(&models.User{}, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	//Delete user
	result := initializers.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete user"})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	//Return
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

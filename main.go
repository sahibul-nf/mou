package main

import (
	"fmt"
	"log"
	"moyu/auth"
	"moyu/campaign"
	"moyu/handler"
	"moyu/helper"
	"moyu/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// membuat koneksi ke db mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "uSNF:uSNF@tcp(127.0.0.1:3306)/moyu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	campaignRepository := campaign.NewRepository(db)

	// campaigns, err := campaignRepository.FindByID(2)

	// fmt.Println(campaigns)

	campaignService := campaign.NewService(campaignRepository)

	input := campaign.GetCampaignDetailInput{
		ID: 1,
	}

	campaign, err := campaignService.FindCampaign(input)

	if campaign.ID == 0 {
		fmt.Println("No campaign found on that id")
	} else {
		fmt.Println(campaign)
	}

	campaignHanler := handler.NewCampaignHandler(campaignService)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	router.Static("/images", "./images")
	router.Static("/campaign-images", "./campaign-images")

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/campaigns", campaignHanler.GetCampaigns)

	router.Run()
}

// ambil nilai header Authorization :: Value -> "Bearer tokentoken"
// dari header Authorization, kita ambil nilai tokennya saja
// kita validasi token
// kita ambil userID
// ambil user dari db berdasarkan userID lewat service
// kita set context isinya user

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := "" // "Bearer tokentoken"
		headerTokenString := strings.Split(authHeader, " ")
		if len(headerTokenString) == 2 {
			tokenString = headerTokenString[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}

}

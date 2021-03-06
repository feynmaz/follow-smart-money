package main

import (
	"log"
	"net/http"

	scrapper "github.com/feynmaz/cypherhunterscrapper"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/feynmaz/follow-smart-money/api/v1/docs"
	"github.com/feynmaz/follow-smart-money/configs"
)

type CoinUrl struct {
	Url string `json:"url"`
}

// swagger:model Investor
type Investor struct {
	Name   string
	Link   string
	Tier   string
	TierId int
}

// @title           Follow Smart Money API
// @version         1.0
// @description     Get a list of the best investors of crypto coin. Built on top of github.com/feynmaz/cypherhunterscrapper

// @contact.name	Nickolay Mazein
// @contact.email	feynmaz@gmail.com
// @contact.url		https://github.com/feynmaz

// @BasePath  /api/v1

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		coins := v1.Group("coins")
		{
			coins.POST("/", PostCoinUrl)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + configs.PORT)
}

// @description 	PostCoinUrl
// @tags			coins
// @summary 		Get a list of the best investors of crypto coin
// @accept 			json
// @produce 		json
// @param 			url body CoinUrl true "url of coin page on cypherhunter.com"
// @success      	200 {object} Investor "success"
// @failure      	500 {string} string "fail"
// @router			/coins [post]
func PostCoinUrl(c *gin.Context) {
	var cu CoinUrl
	if err := c.BindJSON(&cu); err != nil {
		return
	}

	r, err := scrapper.NewCoinRequester(cu.Url)
	if err != nil {
		log.Fatal(err)
		return
	}

	listAll, err := scrapper.GetInvestorsAll(r)
	if err != nil {
		log.Fatal(err)
		return
	}

	listExceptional := scrapper.GetInvestorsExceptional(listAll)
	c.IndentedJSON(http.StatusOK, listExceptional)
}

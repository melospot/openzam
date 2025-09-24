package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type District struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Province struct {
	Name      string     `json:"name"`
	Code      string     `json:"code"`
	Districts []District `json:"districts"`
}

var provinces []Province

func loadData() {
	file, err := os.Open("data/provinces.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&provinces)
	if err != nil {
		panic(err)
	}
}

func GetProvinces(c *gin.Context) {
	c.JSON(http.StatusOK, provinces)
}

func GetDistricts(c *gin.Context) {
	provinceQuery := c.Query("province")
	if provinceQuery != "" {
		for _, p := range provinces {
			if p.Name == provinceQuery {
				c.JSON(http.StatusOK, p.Districts)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Province not found"})
		return
	}

	allDistricts := []District{}
	for _, p := range provinces {
		allDistricts = append(allDistricts, p.Districts...)
	}
	c.JSON(http.StatusOK, allDistricts)
}

func RegisterRoutes(router *gin.Engine) {
	loadData()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/provinces", GetProvinces)
		v1.GET("/districts", GetDistricts)
	}
}

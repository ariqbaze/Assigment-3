package controllers

import (
	"Assigment-3/database"
	"Assigment-3/models"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	d := time.NewTicker(15 * time.Second)

	go func() {
		for range d.C {
			db := database.GetDB()
			status := models.Status{}
			status.Water = rand.Int64N(100)
			status.Wind = rand.Int64N(100)

			if err := db.Debug().Create(&status).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal Server Error",
					"message": err.Error(),
				})
				return
			}
			json_struck, _ := json.Marshal(status)

			fmt.Println(string(json_struck))
			if status.Water < 5 {
				fmt.Println("status water: aman")
			} else if status.Water >= 6 && status.Water <= 8 {
				fmt.Println("status water: siaga")
			} else if status.Water > 8 {
				fmt.Println("status water: bahaya")
			}
			if status.Wind < 6 {
				fmt.Println("status wind: aman")
			} else if status.Wind >= 7 && status.Wind <= 15 {
				fmt.Println("status wind: siaga")
			} else if status.Wind > 15 {
				fmt.Println("status wind: bahaya")
			}
		}
	}()

}

package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EnvStatus struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type StatusData struct {
	Status      EnvStatus `json:"status"`
	WaterStatus string    `json:"water_state"`
	WindStatus  string    `json:"wind_state"`
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		status := randomStatus()

		waterState := determineState(status.Water, "water")
		windState := determineState(status.Wind, "wind")

		data := StatusData{
			Status:      status,
			WaterStatus: waterState,
			WindStatus:  windState,
		}

		c.HTML(http.StatusOK, "index.tmpl", data)
	})

	// update tiap 15 detik
	go updateStatus()

	router.Run(":8080")
}

// randomStatus untuk menghasilkan random nilai
func randomStatus() EnvStatus {
	return EnvStatus{
		Water: rand.Intn(17) + 1,
		Wind:  rand.Intn(17) + 1,
	}
}

// update setiap 15 detik
func updateStatus() {
	for {
		time.Sleep(15 * time.Second)
	}
}

// menentukan status
func determineState(value int, category string) string {
	switch category {
	case "water":
		switch {
		case value < 5:
			return "Aman"
		case value >= 5 && value <= 8:
			return "Siaga"
		default:
			return "Bahaya"
		}
	case "wind":
		switch {
		case value < 6:
			return "Aman"
		case value >= 6 && value <= 15:
			return "Siaga"
		default:
			return "Bahaya"
		}
	default:
		return "Tidak diketahui"
	}
}

package main

import (
	_ "embed"
	"fmt"
	"kcalendar/model"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

//go:embed festival
var festivalContent string

type Festival struct {
	Date     string
	Festival string
	Holiday  int
}

var holidays []Festival
var festivals []Festival

func init() {
	yaml.Unmarshal([]byte(festivalContent), &festivals)
	for _, festival := range festivals {
		if festival.Holiday == 0 {
			continue
		}
		holidays = append(holidays, festival)
	}
}

func main() {
	r := gin.Default()

	r.GET("isHoliday", func(ctx *gin.Context) {
		date := ctx.Query("date")
		fmt.Println(holidays)
		fmt.Println(date)
		index := sort.Search(len(holidays), func(i int) bool {
			return holidays[i].Holiday != 0 && date <= holidays[i].Date
		})
		fmt.Println(holidays[index])
		if index >= len(holidays) {
			ctx.JSON(http.StatusOK, model.Err("不支持的时间"))
			return
		}
		if date == holidays[index].Date {
			ctx.JSON(http.StatusOK, model.Suc(true))
			return
		}
		if index == 0 {
			ctx.JSON(http.StatusOK, model.Err("不支持的时间"))
			return
		}

		for index--; index >= 0 && holidays[index].Holiday == 0; index-- {
		}

		fmt.Println(holidays[index])
		festivalTime, _ := time.Parse("2006-01-02", holidays[index].Date)

		fmt.Println(festivalTime.Add(time.Duration(holidays[index].Holiday) * 24 * time.Hour).Format("2006-01-02"))
		isHoliday := festivalTime.Add(time.Duration(holidays[index].Holiday)*24*time.Hour).Format("2006-01-02") > date
		ctx.JSON(http.StatusOK, model.Suc(isHoliday))
	})

	r.GET("/getNextHoliday", func(ctx *gin.Context) {
		date := ctx.Query("date")
		index := sort.Search(len(holidays), func(i int) bool {
			return holidays[i].Holiday != 0 && date <= holidays[i].Date
		})
		if index >= len(holidays) {
			ctx.JSON(http.StatusOK, model.Err("不支持的时间"))
			return
		}
		if holidays[index].Date == date {
			index++
		}
		if index >= len(holidays) {
			ctx.JSON(http.StatusOK, model.Err("不支持的时间"))
			return
		}

		ctx.JSON(http.StatusOK, model.Suc(holidays[index]))
	})

	r.Run(":8888")
}

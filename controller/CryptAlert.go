package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kevin19930919/CryptoAlert/service"
	"net/http"
)

func AddAlert(context *gin.Context) {
	var AddAlertInfo service.SaveAlert

	// validiate adding alert info
	if err := context.ShouldBindBodyWith(&AddAlertInfo, binding.JSON); err != nil {
		fmt.Println("data valdation fail", err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	// adding alert
	if err := service.AddAlert(&AddAlertInfo); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("success add alert")
	context.JSON(http.StatusOK, AddAlertInfo)
}

func RemoveAlert(context *gin.Context) {
	var Alert service.AlertBase
	if err := context.ShouldBindBodyWith(&Alert, binding.JSON); err != nil {
		fmt.Println("data valdation fail", err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	// remove alert
	if err := Alert.RemoveAlert(); err != nil {
		fmt.Println("remove alert fail", err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("success update order")
	context.JSON(http.StatusOK, Alert)
}

func UpdateAlert(context *gin.Context) {
	// valdate update alert data
	var UpdateInfo service.UpdateAlert
	if err := context.ShouldBindBodyWith(&UpdateInfo, binding.JSON); err != nil {
		fmt.Println("data valdation fail", err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	// initial update alert object
	Alert := service.AlertBase{AlertID: UpdateInfo.AlertID}

	// if price, err := strconv.ParseFloat(UpdateInfo.Price, 64); err != nil {
	// 	fmt.Println("data valdation fail", err.Error())
	// 	context.JSON(http.StatusNotFound, gin.H{
	// 		"error": err.Error(),
	// 	})
	// }
	// UpdateInfo.Price = price
	if err := Alert.UpdateAlert(UpdateInfo); err != nil {
		fmt.Println("remove alert fail", err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, UpdateInfo)
}

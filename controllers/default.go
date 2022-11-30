package controllers

import (
	//"fmt"
	"strconv"
	"test/minicar/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Manufacturer() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "home.html"

	retrievedmanufacturers := make(map[string]orm.Params)
	retrievedmanufacturers = models.GetManufacturers()
	c.Data["RetrievedManufacturers"] = retrievedmanufacturers

	retrievedInventory := make(map[string]orm.Params)
	retrievedInventory = models.GetCarInventory()
	c.Data["RetrievedInventory"] = retrievedInventory
	//fmt.Println("***Length ", retrievedInventory)
	if retrievedInventory["inventory_0"] == nil {
		c.Data["RangeEmpty"] = "true"
	} else {
		c.Data["RangeEmpty"] = "false"
	}

	if c.Ctx.Input.Method() == "POST" {
		submitButton := c.GetString("submitButton")
		if submitButton == "manufacturer" {
			manufacturername := c.GetString("manufacturer_name")
			// if manufacturername == ""{
			// 	c.Data["NameErr"] = "Name required"
			// 	return
			// }

			insertStatus := models.AddManufacturer(manufacturername)
			if insertStatus == 1 {
				c.Redirect("/manufacturer", 303)
			} else {
				return
			}
		} else {
			modelName := c.GetString("car_model")
			manufacturerId := c.GetString("manufacturer")
			totalInventory := c.GetString("inventory")
			// if manufacturername == ""{
			// 	c.Data["NameErr"] = "Name required"
			// 	return
			// }

			insertStatus := models.AddCarModel(modelName, manufacturerId, totalInventory)
			if insertStatus == 1 {
				c.Redirect("/manufacturer", 303)
			} else {
				return
			}
		}
	}
}

func (c *MainController) UpdateInventory() {
	if c.Ctx.Input.Method() == "GET" {
		modelId := c.GetString("modelId")
		manufacturerId := c.GetString("manufacturerId")
		carCount := c.GetString("carCount")
		intCount, _ := strconv.Atoi(carCount)
		newCount := intCount - 1
		var newStatus = 1
		if newCount == 0 {
			newStatus = 0
		}

		updateStatus := models.UpdateCarInventory(modelId, manufacturerId, newCount, newStatus)

		// outputObj := &updateStruct{
		// 	UpdateStatus:      updateStatus,
		// }
		// b, _ := json.Marshal(outputObj)
		c.Ctx.Output.Body([]byte(updateStatus))
	}
}

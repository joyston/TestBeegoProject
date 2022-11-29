package models

import (
	"strconv"
	"github.com/astaxie/beego/orm"
)

func AddManufacturer(manufacturername string) int {
	o := orm.NewOrm()
	ormTransactionalErr := o.Begin()
	_, err := o.Raw("INSERT INTO manufacturer (name) VALUES (?) ", manufacturername).Exec()
	if err != nil {
		ormTransactionalErr = o.Rollback()
		return 0
	}
	if ormTransactionalErr != nil {
		return 0
	}
	ormTransactionalErr = o.Commit()
	return 1
}

func AddCarModel(modelName string, manufacturerId string, totalInventory string) int {
	o := orm.NewOrm()
	ormTransactionalErr := o.Begin()
	_, err := o.Raw("INSERT INTO car_model (model_name, manufacturer_fkid, car_count) VALUES (?,?,?) ", modelName, manufacturerId, totalInventory).Exec()
	if err != nil {
		ormTransactionalErr = o.Rollback()
		return 0
	}
	if ormTransactionalErr != nil {
		return 0
	}
	ormTransactionalErr = o.Commit()
	return 1
}

func UpdateCarInventory(modelId string, manufacturerId string, newCount int, newStatus int) string {
	o := orm.NewOrm()
	ormTransactionalErr := o.Begin()
	_, err := o.Raw("Update car_model SET car_count=?, status=? WHERE _id=? and manufacturer_fkid=? ", newCount, newStatus, modelId, manufacturerId).Exec()
	if err != nil {
		ormTransactionalErr = o.Rollback()
		return "false"
	}
	if ormTransactionalErr != nil {
		return "false"
	}
	ormTransactionalErr = o.Commit()
	return "true"
}

func GetManufacturers() map[string]orm.Params {
	result := make(map[string]orm.Params)
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM manufacturer").Values(&maps)
	if err == nil && num > 0 {
		var i int64 = 0
		for ; i < num; i++ {
			//fmt.Println(maps[i])
		}
		for k, v := range maps {

			key := "manufacturer_" + strconv.Itoa(k)
			result[key] = v
		}
	} else {
		//fmt.Println(err)
		result["status"] = nil
	}
	return result
}


func GetCarInventory() map[string]orm.Params {
	result := make(map[string]orm.Params)
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT m._id as manufacturerId, m.name, cm._id as modelId, cm.model_name, cm.car_count FROM manufacturer as m, car_model as cm WHERE m._id=cm.manufacturer_fkid and cm.status=1").Values(&maps)
	if err == nil && num > 0 {
		var i int64 = 0
		for ; i < num; i++ {
			//fmt.Println(maps[i])
		}
		for k, v := range maps {

			key := "inventory_" + strconv.Itoa(k)
			result[key] = v
		}

	} else {
		//fmt.Println(err)
		result["status"] = nil
	}
	return result
}
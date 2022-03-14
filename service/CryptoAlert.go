package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kevin19930919/CrtptoAlert/database"
	"github.com/kevin19930919/CrtptoAlert/model"
)

type (
	AlertBase struct {
		AlertID int64 `json:"alert_id"`
	}

	SaveAlert struct {
		Crypto    string  `json:"crypto"`
		Direction bool    `json:"direction"`
		Price     float64 `json:"price"`
	}

	UpdateAlert struct {
		Direction bool    `json:"direction"`
		Price     float64 `json:"price"`
	}
)

func SaveAlert(alert *SaveAlert) (err error) {
	if err := db.Create(alert).Error; err != nil {
		fmt.Println("fail to create alert")
		return err
	}
	return nil
}


func (this AlertBase) GetAlertByID(id int64) (*model.Alert, error){
	var AlertModel model.Alert
	if err:= db.Where("id = ?", this.AlertID).First(&AlertModel).Error; err!=nil{
		fmt.Println("fail to get alert record : ",err)
		return nil ,err
	}
	return &AlertModel, nil

}

func (this AlertBase) RemoveAlert() {
	AlertModel ,err := this.GetAlertByID(this.AlertID)
	if err != nil {
		return err
	}

	if err := db.Delete(AlertModel).Error,err!=nil{
		fmt.Println("fail to remove record : ", err)
		return err
	}
	return nil
}

func (this AlertBase) UpdateAlert(UpdateInfo UpdateAlert) (error) {
	AlertModel ,err := this.GetAlertByID(this.AlertID)
	if err != nil {
		return err
	}
	// Update alert in database
    if err := db.Model(&AlertModel).Update(AlertModel{Direction:UpdateInfo.Direction, Price:UpdateInfo.Price}).Error;err!=nil{
		fmt.Println("fail to update alert : ".err)
		return err
	} 
	return nil
}


// func SaveAlert(db *sql.DB, alert *Alert) (*Alert, error) {
// 	// Insert alert into database
// 	sql := `
// 		INSERT INTO alerts(crypto, direction, price)
// 		VALUES(?, ?, ?)
// 		`
// 	res, err := db.Exec(sql, alert.Crypto, alert.Direction, alert.Price)
// 	if err != nil {
// 		return nil, err
// 	}
// 	alert.ID, _ = res.LastInsertId()
// 	return alert, nil
// }

// func RemoveAlert(db *sql.DB, id int64) (*Alert, error) {
// 	// Remove alert from database
// 	record, err := GetAlertByID(db, id)
// 	if err != nil {
// 		return nil, errors.New("Alert not found")
// 	}

// 	sql := `
// 		DELETE FROM alerts
// 		WHERE id = ?
// 		`
// 	_, err = db.Exec(sql, id)
// 	if err != nil {
// 		return nil, errors.New("Alert Removal Failed")
// 	}
// 	return record, nil
}

// func GetAlertByID(db *sql.DB, id int64) (*Alert, error) {
// 	// Get alert from database
// 	sql := `
// 		SELECT id, crypto, price, direction
// 		FROM alerts
// 		WHERE id = ?
// 		`
// 	row := db.QueryRow(sql, id)
// 	alert := new(Alert)
// 	err := row.Scan(&alert.ID, &alert.Crypto, &alert.Price, &alert.Direction)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return alert, nil
// }

func GetAlerts(db *sql.DB) ([]*Alert, error) {
	// Get alerts from database
	sql := `
		SELECT id, crypto, price, direction
		FROM alerts
		`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	alerts := make([]*Alert, 0)
	for rows.Next() {
		alert := new(Alert)
		err := rows.Scan(&alert.ID, &alert.Crypto, &alert.Price, &alert.Direction)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, alert)
	}
	return alerts, nil
}

// func UpdateAlert(db *sql.DB, alert *Alert) (*Alert, error) {
// 	// Update alert in database
// 	sql := `
// 		UPDATE alerts
// 		SET price = ?, direction = ?
// 		WHERE id = ?
// 		`
// 	_, err := db.Exec(sql, alert.Price, alert.Direction, alert.ID)
// 	if err != nil {
// 		log.Debug(err)
// 		return nil, errors.New("Alert Update Failed")
// 	}
// 	return alert, nil
// }

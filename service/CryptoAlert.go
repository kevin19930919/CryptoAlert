package service

import (
	"fmt"
	"github.com/kevin19930919/CryptoAlert/database"
	"github.com/kevin19930919/CryptoAlert/model"
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
		AlertID   int64   `json:"alert_id"`
		Direction bool    `json:"direction"`
		Price     float64 `json:"price"`
	}
)

func AddAlert(alert *SaveAlert) (err error) {
	if err := database.DB.Create(alert).Error; err != nil {
		fmt.Println("fail to create alert")
		return err
	}
	return nil
}

func (this AlertBase) GetAlertByID(id int64) (*model.Alert, error) {
	var AlertModel model.Alert
	if err := database.DB.Where("id = ?", this.AlertID).First(&AlertModel).Error; err != nil {
		fmt.Println("fail to get alert record : ", err)
		return nil, err
	}
	return &AlertModel, nil

}

func (this AlertBase) RemoveAlert() error {
	AlertModel, err := this.GetAlertByID(this.AlertID)
	if err != nil {
		return err
	}

	if err := database.DB.Delete(AlertModel).Error; err != nil {
		fmt.Println("fail to remove record : ", err)
		return err
	}
	return nil
}

func (this AlertBase) UpdateAlert(UpdateInfo UpdateAlert) error {
	AlertModel, err := this.GetAlertByID(this.AlertID)
	if err != nil {
		return err
	}
	// Update alert in database
	if err := database.DB.Model(&AlertModel).Updates(UpdateInfo).Error; err != nil {
		fmt.Println("fail to update alert : ", err)
		return err
	}
	return nil
}

// func GetAlerts(database.DB *sql.database.database.DB) ([]*Alert, error) {
// 	// Get alerts from database
// 	sql := `
// 		SELECT id, crypto, price, direction
// 		FROM alerts
// 		`
// 	rows, err := database.database.DB.Query(sql)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	alerts := make([]*Alert, 0)
// 	for rows.Next() {
// 		alert := new(Alert)
// 		err := rows.Scan(&alert.ID, &alert.Crypto, &alert.Price, &alert.Direction)
// 		if err != nil {
// 			return nil, err
// 		}
// 		alerts = append(alerts, alert)
// 	}
// 	return alerts, nil
// }

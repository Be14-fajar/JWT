package models

import (
	"api/user/model"
	"errors"
	"log"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Nama_barang string `json:"name" form:"nama_barang"`
	Qty         string `json:"qty" form:"qty"`
	UserID      uint   `json:"iduser" form:"iduser"`
	User        model.User
}

type ItemsModel struct {
	DB *gorm.DB
}

func (Im *ItemsModel) InsertItems(newItem Item) (Item, error) {

	err := Im.DB.Create(&newItem).Error
	if err != nil {
		return Item{}, err
	}

	return newItem, nil
}

func (Im *ItemsModel) GetAllItems(id uint) ([]Item, error) {
	var res []Item

	tx := Im.DB.Preload("User").Where("user_id", id).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}
func (Im *ItemsModel) UpdateItem(updateditem Item) (Item, error) {
	qry := Im.DB.Model(&Item{}).Where("id = ? AND user_id = ?", updateditem.ID, updateditem.UserID).Updates(&updateditem)
	err := qry.Error

	if err != nil {
		log.Println("update query error ", err.Error())
		return Item{}, nil
	}

	return updateditem, nil
}

func (um *ItemsModel) DeleteItem(id int) error {
	qry := um.DB.Delete(&Item{}, id)

	affRow := qry.RowsAffected

	if affRow <= 0 {
		log.Println("no data processed")
		return errors.New("tidak ada data yang dihapus")
	}

	err := qry.Error

	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("tidak bisa menghapus data")
	}

	return nil
}

package controller

import (
	"api/barang/models"
	"api/middlewares"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ItemControll struct {
	MDL models.ItemsModel
}

func (uc *ItemControll) InsertItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := middlewares.ExtractToken(c)
		tmp := models.Item{}

		err := c.Bind(&tmp)
		// id := middlewares.ExtractToken(c)
		tmp.UserID = uint(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := uc.MDL.InsertItems(tmp)

		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "sukses menambahkan data"})
	}
}

func (uc *ItemControll) GetAllItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := middlewares.ExtractToken(c)
		res, err := uc.MDL.GetAllItems(uint(id))
		if err != nil {
			log.Println("query error", err.Error())
			return c.JSON(http.StatusInternalServerError, "tidak bisa diproses")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "sukses mendapatkan semua data"})
	}
}

func (uc *ItemControll) UpdateItems() echo.HandlerFunc {
	return func(c echo.Context) error {

		// paramID := c.Param("id")
		// cnvID, err := strconv.Atoi(paramID)
		// if err != nil {
		// 	log.Println("convert id error ", err.Error())
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		// 		"message": "gunakan input angka",
		// 	})
		// }
		id := middlewares.ExtractToken(c)
		tmp := models.Item{}
		err := c.Bind(&tmp)
		tmp.UserID = uint(id)
		if err != nil {
			log.Println("bind body error ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "masukkan input sesuai pola",
			})
		}

		res, err := uc.MDL.UpdateItem(tmp)

		if err != nil {
			log.Println("query error ", err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "berhasil update data",
		})

	}
}
func (uc *ItemControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := middlewares.ExtractToken(c)
		// paramID := c.Param("id")
		// cnvID, err := strconv.Atoi(paramID)
		// if err != nil {
		// 	log.Println("convert id error ", err.Error())
		// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		// 		"message": "gunakan input angka",
		// 	})
		// }

		err := uc.MDL.DeleteItem(id)

		if err != nil {
			log.Println("delete error", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "sukses menghapus data",
		})
	}
}

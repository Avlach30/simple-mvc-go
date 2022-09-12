package controller

import (
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"path"
	"simple-mvc/model"
)

func errorResponseHandler(res http.ResponseWriter, req *http.Request, status int, message string) {
	res.WriteHeader(status)

	if status == http.StatusNotFound {
		res.Write([]byte(message))
	}

	if status == http.StatusBadRequest {
		res.Write([]byte(message))
	}

	if status == http.StatusInternalServerError {
		res.Write([]byte(message))
	}
}

type repository struct {
	db *gorm.DB
}

//* Mendefinisikan function dengan mengembalikan nilai struct repository (db gorm)
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func GetAddDataFormHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		temp, err := template.ParseFiles(path.Join("views", "form-add-data.html"))
		if err != nil {
			errorResponseHandler(res, req, http.StatusInternalServerError, "Render failed! keep calm")
			return
		}

		err = temp.Execute(res, nil)
		if err != nil {
			errorResponseHandler(res, req, http.StatusInternalServerError, "Render failed! keep calm")
			return
		}
	} else {
		errorResponseHandler(res, req, http.StatusBadRequest, "Sorry! only GET method allowed by server!")
		return
	}
}

func (repo *repository) renderData(res http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles(path.Join("views", "render-data.html"))
	if err != nil {
		errorResponseHandler(res, req, http.StatusInternalServerError, "Render failed! keep calm")
		return
	}
	
	var data []model.Data

	repo.db.Find(&data)

	err = temp.Execute(res, data)

	if err != nil {
		errorResponseHandler(res, req, http.StatusInternalServerError, "Render failed! keep calm")
		return
	}
}

func (repo *repository) ProcessAddDataFormHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		err := req.ParseForm()
		if err != nil {
			errorResponseHandler(res, req, http.StatusInternalServerError, "Error occured! keep calm")
			return
		} else {
			inputData := model.Data{
				Name:        req.Form.Get("name"),
				Description: req.Form.Get("description"),
			}

			err := repo.db.Create(&inputData).Error

			if err != nil {
				errorResponseHandler(res, req, http.StatusInternalServerError, "Error occured! keep calm")
				return
			} else {
				repo.renderData(res, req)
			}
		}

	} else {
		repo.renderData(res, req)
	}
}

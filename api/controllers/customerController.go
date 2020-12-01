package controllers

import (
	"itjournal/api/models"
	"itjournal/api/responses"
	"net/http"
)

func (server *Server) CustomerSave(w http.ResponseWriter, r *http.Request) {
	customer := models.Customer{}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customer.Email = r.PostFormValue("email")
	customer.Fname = r.PostFormValue("fname")
	customer.Lname = r.PostFormValue("lname")
	customer.Phone = r.PostFormValue("phone")
	customer.Job = r.PostFormValue("job")
	customer.Sector = r.PostFormValue("sector")
	customer.Status = "0"

	customer.Prepare()
	err = customer.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerSendFile := models.CustomersSendFile{}
	customerSendFile.GovSector = r.PostFormValue("gov_sector")
	customerSendFile.Phone = r.PostFormValue("phone")
	customerSendFile.Job = r.PostFormValue("job")
	customerSendFile.SendType = r.PostFormValue("send_type")
	customerSendFile.StatusCommittee = "1"
	customerSendFile.StatusSurety = r.PostFormValue("status_surety")
	customerSendFile.Topic = r.PostFormValue("topic")

	customerSendFile.Prepare()
	err = customerSendFile.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// เช็คว่าอัพโหลดไฟล์มาหรือยัง
	customerFile := models.CustomersFile{}
	err = customerFile.CustomerUploadFile(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customerCreated, err := customer.CustomerSave(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	customerSendFile.Cid = customerCreated.ID
	customerSendFileCreated, err := customerSendFile.CustomerSaveSendFile(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	customerFile.CsfId = customerSendFileCreated.ID
	err = customerFile.CustomersFileSave(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, customer)
}

func (server *Server) CustomerStatusAll(w http.ResponseWriter, r *http.Request) {
	cs := models.CustomersStatus{}

	css, err := cs.CustomersStatusFindAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, css)
}

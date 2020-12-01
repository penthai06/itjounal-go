package models

import (
	"errors"
	"fmt"
	"io"
	"itjournal/api/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type CustomersFile struct {
	ID          int64     `json:"id"`
	CsfId       int64     `json:"csf_id"`
	FileDocName string    `json:"file_doc_name"`
	FilePdfName string    `json:"file_pdf_name"`
	FilePdfHash string    `json:"file_pdf_hash`
	FileDocHash string    `json:"file_doc_hash`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}

func (cf *CustomersFile) Prepare() {
	cf.ID = 0
	cf.CreatedAt = time.Now()
	cf.UpdateAt = time.Now()
}

func (cf *CustomersFile) CustomerUploadFile(r *http.Request) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	maxUploadSize, _ := strconv.ParseInt(os.Getenv("MAX_UPLOAD_SIZE"), 10, 64)
	maxUploadSize = maxUploadSize * 1024
	// Visit tutor https://github.com/Freshman-tech/file-upload/blob/master/main.go
	fileNameGen, _ := utils.RandomFilename(16) // 32 char

	filePdf := r.MultipartForm.File["filepdf"]
	if len(filePdf) == 0 {
		return errors.New("กรุณาใส่ไฟล์ PDF")
	}

	fileDoc := r.MultipartForm.File["filedoc"]
	if len(fileDoc) == 0 {
		return errors.New("กรุณาใส่ไฟล์ DOC")
	}

	var filePdfName string
	var filePdfOriginal string
	for _, fileHeader := range filePdf {
		if fileHeader.Size > maxUploadSize {
			return errors.New("อัพโหลดไฟล์ PDF ขนาดไม่เกิน 30MB")
		}

		filePdf, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer filePdf.Close()

		buff := make([]byte, 512)
		_, err = filePdf.Read(buff)
		if err != nil {
			return err
		}

		ext := http.DetectContentType(buff)
		if ext != "application/pdf" {
			return errors.New("อัพโหลดได้เฉพาะไฟล์ PDF")
		}

		_, err = filePdf.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}

		filePdfName = fileNameGen + filepath.Ext(fileHeader.Filename)
		filePdfOriginal = filepath.Ext(fileHeader.Filename)
		f, err := os.Create("./files/pdf/" + filePdfName)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, filePdf)
		if err != nil {
			return err
		}
	}
	cf.FilePdfHash = filePdfName
	cf.FilePdfName = filePdfOriginal

	var fileDocName string
	var fileDocOriginal string
	for _, fileHeader := range fileDoc {
		if fileHeader.Size > maxUploadSize {
			return errors.New("อัพโหลดไฟล์ DOC ขนาดไม่เกิน 30MB")
		}

		fileDoc, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer fileDoc.Close()

		buff := make([]byte, 512)
		_, err = fileDoc.Read(buff)
		if err != nil {
			return err
		}

		ext := http.DetectContentType(buff)
		if ext != "application/msword" && ext != "application/vnd.openxmlformats-officedocument.wordprocessing" {
			return errors.New("อัพโหลดได้เฉพาะไฟล์ DOC")
		}

		_, err = fileDoc.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}

		fileDocName = fileNameGen + filepath.Ext(fileHeader.Filename)
		fileDocOriginal = filepath.Ext(fileHeader.Filename)
		f, err := os.Create("./files/doc/" + fileDocName)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, fileDoc)
		if err != nil {
			return err
		}
	}
	cf.FileDocHash = fileDocName
	cf.FileDocName = fileDocOriginal

	return nil
}

func (cf *CustomersFile) CustomersFileSave(db *gorm.DB) error {
	err := db.Debug().Create(&cf).Error
	if err != nil {
		return err
	}
	return nil
}

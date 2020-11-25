package models

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"itjournal/api/utils"

	"github.com/jinzhu/gorm"
)

type ArticleFile struct {
	ID        uint      `json:"id"`
	Aid       uint      `json:"aid"`
	FileDoc   string    `json:"file_doc"`
	FilePdf   string    `json:"file_pdf"`
	CreatedAt time.Time `json:"created_at"`
}

const MAX_UPLOAD_SIZE = 30 * 1024 * 1024 // 30MB

func (a *ArticleFile) Prepare() {
	a.ID = 0
	a.CreatedAt = time.Now()
}

func (a *ArticleFile) UploadFile(r *http.Request) error {
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
	for _, fileHeader := range filePdf {
		if fileHeader.Size > MAX_UPLOAD_SIZE {
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
	a.FilePdf = filePdfName

	var fileDocName string
	for _, fileHeader := range filePdf {
		if fileHeader.Size > MAX_UPLOAD_SIZE {
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
	a.FileDoc = fileDocName

	return nil
}

func (af *ArticleFile) SaveArticleFile(db *gorm.DB) error {
	var err error
	err = db.Debug().Create(&af).Error
	if err != nil {
		return err
	}
	return nil
}

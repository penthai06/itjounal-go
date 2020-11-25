package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Fname     string    `json:"f_name"`
	Lname     string    `json:"l_name"`
	Telephone string    `json:"telephone"`
	Status    string    `json:"status"`
	Job       string    `json:"job"`
	Sector    string    `json:"sector"`
	CreatedAt time.Time `json:"created_at`
	UpdatedAt time.Time `json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Fname = html.EscapeString(strings.TrimSpace(u.Fname))
	u.Lname = html.EscapeString(strings.TrimSpace(u.Lname))
	u.Telephone = html.EscapeString(strings.TrimSpace(u.Telephone))
	u.Job = html.EscapeString(strings.TrimSpace(u.Job))
	u.Sector = html.EscapeString(strings.TrimSpace(u.Sector))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Fname == "" {
			return errors.New("กรุณากรอก ชื่อ")
		}
		if u.Lname == "" {
			return errors.New("กรุณากรอก นามสกุล")
		}
		if u.Job == "" {
			return errors.New("กรุณากรอก ตำแหน่ง")
		}
		if u.Sector == "" {
			return errors.New("กรุณากรอก หน่วยงาน")
		}
		return nil
	case "login":
		if u.Email == "" {
			return errors.New("กรุณากรอก อีเมล")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		if u.Password == "" {
			return errors.New("กรุณากรอก รหัสผ่าน")
		}
		return nil
	default:
		if u.Email == "" {
			return errors.New("กรุณากรอก อีเมล")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("รูปแบบ อีเมล ไม่ถูกต้อง")
		}
		if u.Password == "" {
			return errors.New("กรุณากรอก รหัสผ่าน")
		}
		if u.Fname == "" {
			return errors.New("กรุณากรอก ชื่อ")
		}
		if u.Lname == "" {
			return errors.New("กรุณากรอก นามสกุล")
		}
		if u.Job == "" {
			return errors.New("กรุณากรอก ตำแหน่ง")
		}
		if u.Sector == "" {
			return errors.New("กรุณากรอก หน่วยงาน")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(20).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

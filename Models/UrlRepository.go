package Models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UrlRepository interface {
	Get(url *UrlMaster, id string) error
	Save(url *UrlMaster) error
	Update(url *UrlMaster, id string) error
	Delete(url *UrlMaster, id string) error
	FindAll() ([]UrlMaster, error)
}

type database struct {
	connection *gorm.DB
}

func NewUrlRepository() UrlRepository {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("status: ", err)
	}
	//defer db.Close()
	//db.AutoMigrate(&UrlMaster{})
	return &database{
		connection: db,
	}
}

func (d database) Get(u *UrlMaster, id string) error {
	if err := d.connection.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func (d database) Save(u *UrlMaster) error {
	u.Status = "active"
	if err := d.connection.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (d database) Update(u *UrlMaster, id string) error {
	fmt.Println(u)
	d.connection.Save(u)
	return nil
}

func (d database) Delete(u *UrlMaster, id string) error {
	d.connection.Where("id = ?", id).Delete(u)
	return nil
}

func (d database) FindAll() ([]UrlMaster, error) {
	var url []UrlMaster
	if err := d.connection.Find(&url).Error; err != nil {
		return nil, err
	}
	return url, nil
}
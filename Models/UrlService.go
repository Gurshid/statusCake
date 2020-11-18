package Models

import (
	"../Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllUrl(b *[]UrlMaster) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUrl(b *UrlMaster) (err error) {
	b.Status = "active"
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetUrl(b *UrlMaster, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUrl(b *UrlMaster, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteUrl(b *UrlMaster, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}

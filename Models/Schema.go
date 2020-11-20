package Models

import (
	"github.com/jinzhu/gorm"
)

type UrlMaster struct {
	gorm.Model
	Url     				string 	`json:"url"`
	CrawlTimeout   			int8 	`json:"crawl_timeout"`
	FailureThreshold 		int8 	`json:"failure_threshold"`
	FailureCount			int8 	`json:"failure_count"`
	Status					string	`gorm:"default:active" json:"status"`
	//Frequency 			int8 	`json:"frequency"`
}

func (b *UrlMaster) TableName() string {
	return "url_master"
}

package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

//var repo = NewUrlRepository()

type urlService struct {
	urls []UrlMaster
}

func New() urlService {
	return urlService{urls: []UrlMaster{}}
}

func (*urlService) GetAllUrl() ([]UrlMaster, error) {
	return NewUrlRepository().FindAll()
}

func (*urlService) AddNewUrl(u *UrlMaster) error {
	return NewUrlRepository().Save(u)
}

func (*urlService) GetUrl(u *UrlMaster, id string) (err error) {
	if err = NewUrlRepository().Get(u, id); err == nil {
		checkUrl(u)
	}
	return err
}

func (*urlService) UpdateUrl(u *UrlMaster) error {
	return NewUrlRepository().Update(u)
}

func (*urlService) DeleteUrl(u *UrlMaster, id string) error {
	return NewUrlRepository().Delete(u, id)
}

func checkUrl(u *UrlMaster){
	res, err := http.Get(u.Url)
	if err == nil && res != nil && res.StatusCode == 200 {
		return
	} else {
		u.FailureCount += 1
		if u.FailureCount >= u.FailureThreshold {
			u.Status = "inactive"
		}
		NewUrlRepository().Update(u)
	}
}

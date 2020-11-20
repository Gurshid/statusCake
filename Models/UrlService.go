package Models

import (
	"fmt"
	"net/http"
)

type UrlService struct {
	urls []UrlMaster
	repo UrlRepository
}

func New() UrlService {
	return UrlService{urls: []UrlMaster{}, repo: NewUrlRepository()}
}

func (us *UrlService) GetAllUrl() ([]UrlMaster, error) {
	return us.repo.FindAll()
}

func (us *UrlService) AddNewUrl(u *UrlMaster) error {
	return us.repo.Save(u)
}

func (us *UrlService) GetAndCheckUrl(u *UrlMaster, id string) (err error) {
	if err = us.GetUrl(u, id); err == nil && u.Status == "active" {
		us.checkUrl(u, id)
	} else {
		fmt.Println("status: " + u.Status + " error: %v", err)
	}
	return err
}

func (us *UrlService) GetUrl(u *UrlMaster, id string) (err error) {
	return us.repo.Get(u, id)
}

func (us *UrlService) UpdateUrl(u *UrlMaster, id string) error {
	return us.repo.Update(u, id)
}

func (us *UrlService) DeleteUrl(u *UrlMaster, id string) error {
	return us.repo.Delete(u, id)
}

func (us *UrlService) checkUrl(u *UrlMaster, id string){
	res, err := http.Get(u.Url)
	if err == nil && res != nil && res.StatusCode == 200 {
		return
	} else {
		u.FailureCount += 1
		if u.FailureCount >= u.FailureThreshold {
			u.Status = "inactive"
		}
		us.repo.Update(u, id)
	}
}

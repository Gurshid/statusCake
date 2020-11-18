package ApiHelpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status 	int
	Data   	interface{}
	Error 	string
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	if payload != nil && status == 404 {
		fmt.Println("status: ", status , "\nerror: ", payload)
	}

	var res ResponseData

	res.Status = status
	if status == 200 {
		res.Data = payload
	} else {
		res.Error = fmt.Sprintf("%v", payload)
	}

	w.JSON(status, res)
}

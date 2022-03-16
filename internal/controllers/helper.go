package controllers

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// bindData binds the data from the request to the struct passed in the interface
func bindData(c *gin.Context, data interface{}) (int, error) {
	if err := c.ShouldBindJSON(&data); err != nil {
		if errors.Is(io.EOF, err) {
			return http.StatusBadRequest, errors.New("request body must not be emtpy")
		}
		log.Println(err)
		return http.StatusInternalServerError, errors.New("there was an error processing your request, please contact your server administrator")
	}
	return http.StatusOK, nil
}
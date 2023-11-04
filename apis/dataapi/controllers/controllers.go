package controllers

import (
	"antinolabsassignment/pkg/blogs"
	"antinolabsassignment/pkg/common/models/blog"
	log "antinolabsassignment/pkg/common/utilities/logger"
	"encoding/json"
	"net/http"
)

// region Struct

type Controllers struct{}

//endregion

// region public methods

// Post - to control POST requests
// Param res - it holds response data
// Param req - it holds request data
func (c Controllers) Post(res http.ResponseWriter, req *http.Request) {

	var (
		err   error
		input blog.BlogConfig
	)

	if input, err = decoder(req); err != nil {
		convertToJSON(err.Error(), nil, &res)
		return
	}

	instance, err := blogs.NewBlog(log.New())
	if err != nil || instance.Post(input) != nil {
		convertToJSON("error occurred", nil, &res)
		return
	}
	convertToJSON("", map[string]interface{}{"Message": "Successful", "Error": nil}, &res)
}

// GetById - to control GET requests
// Param res - it holds response data
// Param req - it holds request data
func (c Controllers) GetById(res http.ResponseWriter, req *http.Request) {

	var (
		err    error
		blogId string
		result *blog.BlogConfig
	)

	blogId = req.URL.Query().Get("blogId")
	if len(blogId) == 0 {
		convertToJSON("error occurred", nil, &res)
		return
	}

	instance, err := blogs.NewBlog(log.New())
	if err != nil {
		convertToJSON(err.Error(), nil, &res)
		return
	}

	result, err = instance.GetById(blogId)
	if err != nil {
		convertToJSON(err.Error(), nil, &res)
		return
	}
	convertToJSON("", result, &res)
}

// GetAll - to control GET All requests
// Param res - it holds response data
// Param req - it holds request data
func (c Controllers) GetAll(res http.ResponseWriter, req *http.Request) {

	instance, err := blogs.NewBlog(log.New())
	if err != nil {
		convertToJSON(err.Error(), nil, &res)
		return
	}

	result, err := instance.GetAll()
	if err != nil {
		convertToJSON(err.Error(), nil, &res)
		return
	}
	convertToJSON("", result, &res)
}

// UpdateById - to control UPDATE requests
// Param res - it holds response data
// Param req - it holds request data
func (c Controllers) UpdateById(res http.ResponseWriter, req *http.Request) {

	var (
		err   error
		input blog.BlogConfig
	)

	if input, err = decoder(req); err != nil {
		convertToJSON(err.Error(), nil, &res)
		return
	}

	// blogId cannot be <= 0
	if input.BlogId <= 0 {
		convertToJSON("invalid blogId found", nil, &res)
		return
	}

	instance, err := blogs.NewBlog(log.New())
	if err != nil || instance.UpdateById(input) != nil {
		convertToJSON("error occurred", nil, &res)
		return
	}
	convertToJSON("", map[string]interface{}{"Message": "Successful", "Error": nil}, &res)
}

// DeleteById - to control DELETE requests
// Param res - it holds response data
// Param req - it holds request data
func (c Controllers) DeleteById(res http.ResponseWriter, req *http.Request) {

	var (
		err    error
		blogId string
	)

	blogId = req.URL.Query().Get("blogId")
	if len(blogId) == 0 {
		convertToJSON("error occurred", nil, &res)
		return
	}

	instance, err := blogs.NewBlog(log.New())
	if err != nil || instance.DeleteById(blogId) != nil {
		convertToJSON("error occurred", nil, &res)
		return
	}
	convertToJSON("", map[string]interface{}{"Message": "Successful", "Error": nil}, &res)
}

//endregion

// region private methods

func convertToJSON(err string, data interface{}, res *http.ResponseWriter) {
	(*res).Header().Add("Content-Type", "application/json")
	if err == "" {
		(*res).WriteHeader(http.StatusOK)
		_ = json.NewEncoder(*res).Encode(data)
		return
	}
	(*res).WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(*res).Encode(map[string]interface{}{"Message": "Unsuccessful", "Error": err})
}

func decoder(req *http.Request) (blog.BlogConfig, error) {
	var (
		err  error
		data blog.BlogConfig
	)
	err = json.NewDecoder(req.Body).Decode(&data)
	return data, err
}

//endregion

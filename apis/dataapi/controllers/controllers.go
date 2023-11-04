package controllers

import (
	"antinolabsassignment/pkg/blogs"
	"antinolabsassignment/pkg/common/models/blog"
	log "antinolabsassignment/pkg/common/utilities/logger"
	"encoding/json"
	"net/http"
)

type Controllers struct{}

func (c Controllers) Post(res http.ResponseWriter, req *http.Request) {
	var (
		err   error
		input blog.BlogConfig
	)

	if input, err = decoder(req); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	instance, err := blogs.NewBlog(log.New())
	if err != nil || instance.Post(input) != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}

func (c Controllers) GetById(res http.ResponseWriter, req *http.Request) {
	var (
		err    error
		blogId string
		result *blog.BlogConfig
	)

	blogId = req.URL.Query().Get("blogId")
	if len(blogId) == 0 {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	instance, err := blogs.NewBlog(log.New())
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	result, err = instance.GetById(blogId)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(res).Encode(result)
}

func (c Controllers) UpdateById(res http.ResponseWriter, req *http.Request) {
	var (
		err   error
		input blog.BlogConfig
	)

	if input, err = decoder(req); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	instance, err := blogs.NewBlog(log.New())
	if err != nil || instance.UpdateById(input) != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}

func (c Controllers) DeleteById(res http.ResponseWriter, req *http.Request) {
	var (
		err    error
		blogId string
		result *blog.BlogConfig
	)

	blogId = req.URL.Query().Get("blogId")
	if len(blogId) == 0 {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	instance, err := blogs.NewBlog(log.New())
	if err != nil || instance.DeleteById(blogId) != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(res).Encode(result)
}

func decoder(req *http.Request) (blog.BlogConfig, error) {
	var (
		err  error
		data blog.BlogConfig
	)
	err = json.NewDecoder(req.Body).Decode(&data)
	return data, err
}

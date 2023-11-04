package blogs

import (
	"antinolabsassignment/pkg/blogs/internal"
	"antinolabsassignment/pkg/common/models/blog"
	log "antinolabsassignment/pkg/common/utilities/logger"
	"antinolabsassignment/pkg/common/utilities/viper"
	"strconv"
)

// region Struct

type Blog struct {
	log      log.Logger
	sqlDbIns *internal.SqlDbHelper
}

//endregion

// region Ctor

// NewBlog - it creates and returns new blog instance
// Param log - logger
func NewBlog(log log.Logger) (*Blog, error) {
	log.Info("get new viper instance")
	vp, err := viper.NewViper()
	if err != nil {
		return nil, err
	}

	log.Info("get sql db helper")
	sqlDbHelper, err := internal.NewSqlDbHelper(log, vp.GetString("DATABASE"))
	if err != nil {
		return nil, err
	}

	return &Blog{log: log, sqlDbIns: sqlDbHelper}, nil
}

//endregion

// region public methods

// Post - calls the post method
// Param blog - contains the blog data
func (b *Blog) Post(blog blog.BlogConfig) error {
	b.log.Info("call post method")
	return b.sqlDbIns.Post(blog)
}

// GetById - calls get method
// Param blogId - holds blog id
func (b *Blog) GetById(blogId string) (*blog.BlogConfig, error) {
	b.log.Info("parse string to int")
	key, err := strconv.Atoi(blogId)
	if err != nil {
		return nil, err
	}

	b.log.Info("call get method")
	return b.sqlDbIns.Get(key)
}

// DeleteById - calls the delete method
// Param  blogId - holds blog id
func (b *Blog) DeleteById(blogId string) error {
	b.log.Info("parse string to int")
	key, err := strconv.Atoi(blogId)
	if err != nil {
		return err
	}

	b.log.Info("call delete method")
	return b.sqlDbIns.Delete(key)
}

// UpdateById - calls the update method
// Param arg - holds blog data
func (b *Blog) UpdateById(arg blog.BlogConfig) error {
	b.log.Info("call update method")
	return b.sqlDbIns.Update(arg)
}

//endregion

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
	vp, err := viper.NewViper()
	if err != nil {
		return nil, err
	}
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
	return b.sqlDbIns.Post(blog)
}

// GetById - calls get method
// Param blogId - holds blog id
func (b *Blog) GetById(blogId string) (*blog.BlogConfig, error) {
	key, err := strconv.Atoi(blogId)
	if err != nil {
		return nil, err
	}
	return b.sqlDbIns.Get(key)
}

// DeleteById - calls the delete method
// Param  blogId - holds blog id
func (b *Blog) DeleteById(blogId string) error {
	key, err := strconv.Atoi(blogId)
	if err != nil {
		return err
	}
	return b.sqlDbIns.Delete(key)
}

// UpdateById - calls the update method
// Param arg - holds blog data
func (b *Blog) UpdateById(arg blog.BlogConfig) error {
	return b.sqlDbIns.Update(arg)
}

//endregion

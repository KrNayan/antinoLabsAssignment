package blogs

import (
	"antinolabsassignment/pkg/blogs/internal"
	"antinolabsassignment/pkg/common/models/blog"
	log "antinolabsassignment/pkg/common/utilities/logger"
	"strconv"
)

type Blog struct {
	log      log.Logger
	sqlDbIns *internal.SqlDbHelper
}

func NewBlog(log log.Logger) (*Blog, error) {
	sqlDbHelper, err := internal.NewSqlDbHelper(log, "")
	if err != nil {
		return nil, err
	}
	return &Blog{log: log, sqlDbIns: sqlDbHelper}, nil
}

func (b *Blog) Post(blog blog.BlogConfig) error {
	return b.sqlDbIns.Post(blog)
}

func (b *Blog) GetById(blogId string) (*blog.BlogConfig, error) {
	key, err := strconv.Atoi(blogId)
	if err != nil {
		return nil, err
	}
	return b.sqlDbIns.Get(key)
}

func (b *Blog) DeleteById(blogId string) error {
	key, err := strconv.Atoi(blogId)
	if err != nil {
		return err
	}
	return b.sqlDbIns.Delete(key)
}

func (b *Blog) UpdateById(param blog.BlogConfig) error {
	return b.sqlDbIns.Update(param)
}

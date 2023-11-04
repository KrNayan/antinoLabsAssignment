package internal

import (
	"antinolabsassignment/pkg/common/db/mysql"
	"antinolabsassignment/pkg/common/models/blog"
	log "antinolabsassignment/pkg/common/utilities/logger"
	"database/sql"
	"errors"
)

type SqlDbHelper struct {
	sqlDB *mysql.SqlDb
}

func NewSqlDbHelper(log log.Logger, database string) (*SqlDbHelper, error) {
	sqlDB, err := mysql.NewSQLDbInstance(log, database)
	if err != nil {
		return nil, err
	}
	return &SqlDbHelper{
		sqlDB: sqlDB,
	}, nil
}

func (sdh SqlDbHelper) Post(blog blog.BlogConfig) error {
	var query = "INSERT INTO blogs (emailId, blog) VALUES (?, ?)"
	return sdh.sqlDB.Post(query, blog.EmailId, blog.Blog)
}

func (sdh SqlDbHelper) Get(blogId int) (*blog.BlogConfig, error) {
	var (
		row    *sql.Row
		result *blog.BlogConfig
	)

	row = sdh.sqlDB.Get("SELECT * FROM blobs WHERE blogId = ?", blogId)
	if err := row.Scan(&result.BlogId, &result.EmailId, &result.Blog); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return result, err
	}
	return result, nil
}

func (sdh SqlDbHelper) Delete(blogId int) error {
	var query = "DELETE FROM blogs WHERE blogId = ?"
	return sdh.sqlDB.Delete(query, blogId)
}

func (sdh SqlDbHelper) Update(param blog.BlogConfig) error {
	var query = "UPDATE TABLE blogs SET blog = ? WHERE blogId = ?"
	return sdh.sqlDB.Update(query, param.Blog, param.BlogId)
}

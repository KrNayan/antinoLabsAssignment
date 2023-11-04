package internal

import (
	"antinolabsassignment/pkg/common/db/mysql"
	"antinolabsassignment/pkg/common/models/blog"
	log "antinolabsassignment/pkg/common/utilities/logger"
	"database/sql"
	"errors"
	"time"
)

// region Struct

type SqlDbHelper struct {
	log   log.Logger
	sqlDB *mysql.SqlDb
}

//endregion

//region Ctor

// NewSqlDbHelper - it returns sql db helper
// Param log - logger
// Param database - database
func NewSqlDbHelper(log log.Logger, database string) (*SqlDbHelper, error) {
	log.Infof("init sql db instance")
	sqlDB, err := mysql.NewSQLDbInstance(log, database)
	if err != nil {
		return nil, err
	}
	return &SqlDbHelper{
		log:   log,
		sqlDB: sqlDB,
	}, nil
}

//endregion

//region public methods

// Post - it calls post method for data posting
// Param blog - it contains the blog data
func (sdh SqlDbHelper) Post(blog blog.BlogConfig) error {
	sdh.log.Info("build post query")
	var query = "INSERT INTO `users_blog` (`emailId`, `blog`) VALUES (?, ?)"

	sdh.log.Info("call post method with query: ", query)
	return sdh.sqlDB.Post(query, blog.EmailId, blog.Blog)
}

// Get  - it calls get method for data retrieval
// Param blogId - blogId
func (sdh SqlDbHelper) Get(blogId int) (*blog.BlogConfig, error) {
	var (
		query    string
		postedOn string
		row      *sql.Row
		result   = &blog.BlogConfig{}
	)

	sdh.log.Info("build get query")
	query = "SELECT * FROM `users_blog` WHERE `blogId` = ?"

	sdh.log.Info("call get method with query: ", query)
	row = sdh.sqlDB.Get(query, blogId)
	if err := row.Scan(&result.BlogId, &result.EmailId, &result.Blog, &postedOn); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return result, err
	}

	sdh.log.Info("parse the postedOn time")
	postedOnDate, _ := time.Parse("2006-01-02 15:04:05", postedOn)
	result.PostedOn = &postedOnDate
	return result, nil
}

// Delete  - it calls delete method for blog deletion
// Param blogId - blogId
func (sdh SqlDbHelper) Delete(blogId int) error {
	sdh.log.Info("build delete query")
	var query = "DELETE FROM `users_blog` WHERE `blogId` = ?"

	sdh.log.Info("call delete method with query: ", query)
	return sdh.sqlDB.Delete(query, blogId)
}

// Update  - it calls update method for blog-update
// Param blogId - blogId
func (sdh SqlDbHelper) Update(param blog.BlogConfig) error {
	sdh.log.Info("build update query")
	var query = "UPDATE `users_blog` SET `blog` = ? WHERE `blogId` = ?"

	sdh.log.Info("call update method with query: ", query)
	return sdh.sqlDB.Update(query, param.Blog, param.BlogId)
}

//endregion

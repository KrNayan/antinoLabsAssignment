package internal

import (
	"antinolabsassignment/pkg/common/db/mysql"
	"antinolabsassignment/pkg/common/models/blog"
	log "antinolabsassignment/pkg/common/utilities/logger"
	"database/sql"
	"errors"
	"fmt"
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
func (sdh SqlDbHelper) Get(blogId int) ([]blog.BlogConfig, error) {
	var resp = make([]blog.BlogConfig, 0)

	sdh.log.Info("get rows as per requirement")
	rows, err := sdh.getRows(blogId)
	if err != nil {
		return nil, err
	}

	sdh.log.Info("iterate the rows one by one")
	for rows.Next() {
		var (
			postedOn string
			result   blog.BlogConfig
		)
		// scan the row
		err = rows.Scan(&result.BlogId, &result.EmailId, &result.Blog, &postedOn)
		if err != nil {
			// suppressing error here
			sdh.log.Error("error while scanning records: ", err.Error())
			continue
		}
		//parse the postedOn time
		postedOnDate, _ := time.Parse("2006-01-02 15:04:05", postedOn)
		result.PostedOn = &postedOnDate

		resp = append(resp, result)
	}
	return resp, nil
}

// Delete - it calls delete method for blog deletion
// Param blogId - blogId
func (sdh SqlDbHelper) Delete(blogId int) error {
	sdh.log.Info("build delete query")
	var query = "DELETE FROM `users_blog` WHERE `blogId` = ?"

	sdh.log.Info("call delete method with query: ", query)
	return sdh.sqlDB.Delete(query, blogId)
}

// Update - it calls update method for blog-update
// Param blogId - blogId
func (sdh SqlDbHelper) Update(param blog.BlogConfig) error {
	sdh.log.Info("build update query")
	var query = "UPDATE `users_blog` SET `blog` = ? WHERE `blogId` = ?"

	sdh.log.Info("call update method with query: ", query)
	return sdh.sqlDB.Update(query, param.Blog, param.BlogId)
}

//endregion

//region private methods

// getRows - wrapper method to call get
// Param blogId - holds the blogId
func (sdh SqlDbHelper) getRows(blogId int) (*sql.Rows, error) {

	sdh.log.Info("build get query")
	query := "SELECT * FROM `users_blog`"

	if blogId > 0 {
		query = fmt.Sprintf("%s %s", query, "WHERE `blogId` = ?")
		sdh.log.Info("fetch the matching records with query: ", query)
		rows, err := sdh.sqlDB.Get(query, blogId)
		if err != nil {
			return rows, errors.New("error in fetching records")
		}
		return rows, nil
	}

	sdh.log.Info("fetch all the records with query: ", query)
	rows, err := sdh.sqlDB.Get(query)
	if err != nil {
		return rows, errors.New("error in fetching records")
	}
	return rows, nil
}

//endregion

package mysql

import (
	mySql "antinolabsassignment/pkg/common/models/db"
	log "antinolabsassignment/pkg/common/utilities/logger"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// region Struct

type SqlDb struct {
	log         log.Logger
	sqlInstance *sql.DB
}

//endregion

// region Ctor

// NewSQLDbInstance - it returns new sql db instance and an error if any
// Param log - logger
// Param dataBase - database name
func NewSQLDbInstance(log log.Logger, dataBase string) (*SqlDb, error) {

	log.Info("get the sql db credentials")
	sqlDbCredentials, err := mySql.GetSqlDbCred(log)
	if err != nil {
		return nil, err
	}

	log.Info("build the connection string for sql")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", sqlDbCredentials.Username,
		sqlDbCredentials.Password,
		sqlDbCredentials.Host,
		sqlDbCredentials.Port,
		dataBase,
	)

	log.Info("open the connection to sql")
	db, err := sql.Open("mysql", connString)
	if err != nil || db.Ping() != nil {
		msg := "failed to create the connection"
		log.Errorf(msg)
		return nil, errors.New(msg)
	}

	return &SqlDb{log: log, sqlInstance: db}, nil
}

//endregion

//region public methods

// Post - it calls query executor
// Param query - it hold SQL query string
// Param args (optional) - it holds arguments
func (sd *SqlDb) Post(query string, args ...any) error {
	sd.log.Info("call the query executor")
	return sd.executeQuery(query, args...)
}

// Get - it calls query row methods
// Param query - it hold SQL query string
// Param args (optional) - it holds arguments
func (sd *SqlDb) Get(query string, args ...any) (*sql.Rows, error) {
	sd.log.Info("query to fetch the rows")
	rows, err := sd.sqlInstance.Query(query, args...)
	return rows, err
}

// Delete - it calls query executor
// Param query - it hold SQL query string
// Param args (optional) - it holds arguments
func (sd *SqlDb) Delete(query string, args ...any) error {
	sd.log.Info("call the query executor")
	return sd.executeQuery(query, args...)
}

// Update - it calls query executor for update
// Param query - it hold SQL query string
// Param args (optional) - it holds arguments
func (sd *SqlDb) Update(query string, args ...any) error {
	sd.log.Info("call the query executor")
	return sd.executeQuery(query, args...)
}

//endregion

//region private methods

// executeQuery- it calls the executor for query execution
// Param query - it contains the SQL query string
// Param args (optional) - it holds arguments
func (sd *SqlDb) executeQuery(query string, args ...any) error {
	sd.log.Info("execute the query: ", query)
	result, err := sd.sqlInstance.Exec(query, args...)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

//endregion

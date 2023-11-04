package mysql

import (
	mySql "antinolabsassignment/pkg/common/models/db"
	log "antinolabsassignment/pkg/common/utilities/logger"
	"database/sql"
	"fmt"
)

type SqlDb struct {
	log      log.Logger
	Instance *sql.DB
}

func NewSQLDbInstance(log log.Logger, dataBase string) (*SqlDb, error) {
	sqlDbCredentials := mySql.SqlDbCred{}
	// build the DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", sqlDbCredentials.Username,
		sqlDbCredentials.Password,
		sqlDbCredentials.Host,
		sqlDbCredentials.Port,
		dataBase)
	// Open the connection
	db, err := sql.Open("mysql", dsn)
	if err != nil || db.Ping() != nil {
		log.Errorf("failed to create the connection")
		return nil, err
	}
	return &SqlDb{Instance: db}, nil
}

func (sd *SqlDb) Post(query string, args ...any) error {
	return sd.executeQuery(query, args)
}

func (sd *SqlDb) Get(query string, args ...any) *sql.Row {
	row := sd.Instance.QueryRow(query, args)
	return row
}

func (sd *SqlDb) Delete(query string, args ...any) error {
	return sd.executeQuery(query, args)
}

func (sd *SqlDb) Update(query string, args ...any) error {
	return sd.executeQuery(query, args)
}

func (sd *SqlDb) executeQuery(query string, args ...any) error {
	result, err := sd.Instance.Exec(query, args)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

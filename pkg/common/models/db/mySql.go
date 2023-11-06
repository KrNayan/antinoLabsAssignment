package db

import (
	log "antinolabsassignment/pkg/common/utilities/logger"
	"antinolabsassignment/pkg/common/utilities/viper"
)

// region Struct

// SqlDbCred - struct the hold the required params
type SqlDbCred struct {
	// Username contains the username of sql db
	Username string
	// Password contains the password of sql db
	Password string
	// Database contains the sql db name
	Database string
	// Host contains the host string
	Host string
	// Port contains the port
	Port int
}

//endregion

// region public functions

// GetSqlDbCred - it extracts and returns the sql db credentials
// Param log - contains the logger
func GetSqlDbCred(log log.Logger) (*SqlDbCred, error) {
	vp, err := viper.NewViper()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &SqlDbCred{
		Username: vp.GetString("USER_NAME"),
		Password: vp.GetString("PASSWORD"),
		Database: vp.GetString("DATABASE"),
		Host:     "localhost",
		Port:     3306,
	}, nil
}

//endregion

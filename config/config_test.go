package config_test

import (
	"testing"
	"timesheet/config"
)

func Test_GetURI_Tnput_DatabaseConfig_Username_root_Password_root_Host_my_mariadb_Port_3306_Database_timesheet_Should_Be_URI(t *testing.T) {
	expectedResult := "root:root@tcp(my-mariadb:3306)/timesheet?parseTime=true"
	databaseConfig := config.DatabaseConfig{
		Username: "root",
		Password: "root",
		Host:     "my-mariadb",
		Port:     "3306",
		Database: "timesheet",
	}

	actualResult := databaseConfig.GetURI()

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v", expectedResult, actualResult)
	}
}

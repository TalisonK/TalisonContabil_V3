package util

import (
	"github.com/TalisonK/TalisonContabil/src/database"
)

/*
CheckDBStatus checks the status of the local and cloud databases
Returns the status of the local and cloud databases
*/
func CheckDBStatus() (bool, bool) {
	// Check database status
	statusDbLocal := database.CheckLocalDB()
	statusDbCloud := database.CheckCloudDB()

	return statusDbLocal, statusDbCloud
}

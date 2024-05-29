package logging

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func logHandler(message string, err error) string {
	fmt.Println(message, err, getFunctionName(), time.Now())
	return message
}

func NoDatabaseConnection() error {
	return fmt.Errorf(logHandler("No database connection available.", nil))
}

func GenericError(message string, err error) string {
	return logHandler(message, err)
}

func GenericSuccess(message string) string {
	return logHandler(message, nil)
}

func ErrorOccurred() error {
	return fmt.Errorf(logHandler("An error occurred.", nil))
}

func InvalidFields() string {
	return logHandler("Invalid fields.", nil)
}

// FAILED TO

func FailedToOpenConnection(base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to open connection to %s database.", base), err)
}

func FailedToConnectToDB(base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to connect to %s database.", base), err)
}

func FailedToFindOnDB(id string, base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to find %s on %s database.", id, base), err)
}

func FailedToCreateOnDB(id string, base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to create %s on %s database.", id, base), err)
}

func FailedToUpdateOnDB(id string, base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to update %s on %s database.", id, base), err)
}

func FailedToDeleteOnDB(id string, base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to delete %s on %s database.", id, base), err)
}

func FailedToAuthenticate(user string) string {
	return logHandler(fmt.Sprintf("Failed to authenticate user %s.", user), nil)
}

func FailedToHashPassword(err error) string {
	return logHandler("Failed to hash password.", err)
}

func FailedToGenerateSalt(err error) string {
	return logHandler("Failed to generate salt.", err)
}

func FailedToConvertPrimitive(err error) string {
	return logHandler("Failed to convert primitive.", err)
}

func FailedToPingDB(base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to ping %s database.", base), err)
}

func FailedToCloseConnection(base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to close connection to %s database.", base), err)
}

func FailedToParseBody(err error) string {
	return logHandler("Failed to parse body.", err)
}

// SUCCESS

func CreatedOnDB(id string, base string) string {
	return logHandler(fmt.Sprintf("%s created on %s database.", id, base), nil)
}

func UpdatedOnDB(id string, base string) string {
	return logHandler(fmt.Sprintf("%s updated on %s database.", id, base), nil)
}

func DeletedOnDB(id string, base string) string {
	return logHandler(fmt.Sprintf("%s deleted on %s database.", id, base), nil)
}

func FoundOnDB(id string, base string) string {
	return logHandler(fmt.Sprintf("%s found on %s database.", id, base), nil)
}

func OpenedConnection(base string) string {
	return logHandler(fmt.Sprintf("Opened connection to %s database.", base), nil)
}

// OTHERS

func EmptyPassword() string {
	return logHandler("User name or password is empty.", nil)
}

func DuplicatedEntry(id string) string {
	return logHandler(fmt.Sprintf("Duplicated entry %s on database.", id), nil)
}

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(3)
	fullName := runtime.FuncForPC(pc).Name()
	slicedPath := strings.Split(fullName, "/")
	return slicedPath[len(slicedPath)-1] // remove o ponto no início da extensão
}

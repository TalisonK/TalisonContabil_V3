package logging

import (
	"fmt"
	"time"
)

func logHandler(message string, err error, function string) string {
	fmt.Println(message, err, function, time.Now())
	return message
}

func NoDatabaseConnection(function string) error {
	return fmt.Errorf(logHandler("No database connection available.", nil, function))
}

func GenericError(message string, err error, function string) string {
	return logHandler(message, err, function)
}

func GenericSuccess(message string, function string) string {
	return logHandler(message, nil, function)
}

func ErrorOccurred(function string) error {
	return fmt.Errorf(logHandler("An error occurred.", nil, function))
}

// FAILED TO

func FailedToOpenConnection(base string, err error, function string) string {
	return logHandler(fmt.Sprintf("Failed to open connection to %s database.", base), err, function)
}

func FailedToConnectToDB(base string, err error, function string) string {
	return logHandler(fmt.Sprintf("Failed to connect to %s database.", base), err, function)
}

func FailedToFindOnDB(id string, base string, err error, function string) string {
	return logHandler(fmt.Sprintf("Failed to find %s on %s database.", id, base), err, function)
}

func FailedToCreateOnDB(id string, base string, err error, function string) string {
	return logHandler(fmt.Sprintf("Failed to create %s on %s database.", id, base), err, function)
}

func FailedToUpdateOnDB(id string, base string, err error, function string) string {
	return logHandler(fmt.Sprintf("Failed to update %s on %s database.", id, base), err, function)
}

func FailedToDeleteOnDB(id string, base string, err error, function string) string {
	return logHandler(fmt.Sprintf("Failed to delete %s on %s database.", id, base), err, function)
}

func FailedToAuthenticate(user string, function string) string {
	return logHandler(fmt.Sprintf("Failed to authenticate user %s.", user), nil, function)
}

func FailedToHashPassword(err error, function string) string {
	return logHandler("Failed to hash password.", err, function)
}

func FailedToGenerateSalt(err error, function string) string {
	return logHandler("Failed to generate salt.", err, function)
}

func FailedToConvertPrimitive(err error, function string) string {
	return logHandler("Failed to convert primitive.", err, function)
}

func FailedToPingDB(base string, err error, function string) string {
	return logHandler(fmt.Sprintf("Failed to ping %s database.", base), err, function)
}

func FailedToCloseConnection(base string, err error, function string) string {
	return logHandler(fmt.Sprintf("Failed to close connection to %s database.", base), err, function)
}

func FailedToParseBody(err error, function string) string {
	return logHandler("Failed to parse body.", err, function)
}

// SUCCESS

func CreatedOnDB(id string, base string, function string) string {
	return logHandler(fmt.Sprintf("%s created on %s database.", id, base), nil, function)
}

func UpdatedOnDB(id string, base string, function string) string {
	return logHandler(fmt.Sprintf("%s updated on %s database.", id, base), nil, function)
}

func DeletedOnDB(id string, base string, function string) string {
	return logHandler(fmt.Sprintf("%s deleted on %s database.", id, base), nil, function)
}

func FoundOnDB(id string, base string, function string) string {
	return logHandler(fmt.Sprintf("%s found on %s database.", id, base), nil, function)
}

func OpenedConnection(base string, function string) string {
	return logHandler(fmt.Sprintf("Opened connection to %s database.", base), nil, function)
}

// OTHERS

func EmptyPassword(function string) string {
	return logHandler("User name or password is empty.", nil, function)
}

func DuplicatedEntry(id string, function string) string {
	return logHandler(fmt.Sprintf("Duplicated entry %s on database.", id), nil, function)
}

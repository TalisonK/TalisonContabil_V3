package logging

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/TalisonK/TalisonContabil/pkg/timeHandler"
)

var mutex sync.Mutex
var wg sync.WaitGroup

func logHandler(message string, err error) string {
	wg.Add(1)

	funcao := getFunctionName()

	go func() {

		defer wg.Done()

		f, _ := os.OpenFile(fmt.Sprintf("../../log/talisoncontabil-%s.log", time.Now().Format(time.DateOnly)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()

		mutex.Lock()
		defer mutex.Unlock()

		if err == nil {
			f.WriteString(fmt.Sprintf("%s, %s, %s, %s\n", timeHandler.GetTimeNow(), funcao, message, "No error"))
		} else {
			f.WriteString(fmt.Sprintf("%s, %s, %s, %s\n", timeHandler.GetTimeNow(), funcao, message, err))
		}

		f.Sync()
	}()

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

func InvalidFields() error {
	return fmt.Errorf(logHandler("Invalid fields.", nil))
}

func ContextAlreadyClosed() string {
	return logHandler("Context already closed.", nil)
}

// FAILED TO

// "Failed to open connection to %s database."
func FailedToOpenConnection(base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to open connection to %s database.", base), err)
}

// "Failed to connect to %s database."
func FailedToConnectToDB(base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to connect to %s database.", base), err)
}

// "Failed to find %s on %s database."
func FailedToFindOnDB(id string, base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to find %s on %s database.", id, base), err)
}

// "Failed to create %s on %s database."
func FailedToCreateOnDB(id string, base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to create %s on %s database.", id, base), err)
}

// "Failed to update %s on %s database."
func FailedToUpdateOnDB(id string, base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to update %s on %s database.", id, base), err)
}

// "Failed to delete %s on %s database."
func FailedToDeleteOnDB(id string, base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to delete %s on %s database.", id, base), err)
}

// "Failed to authenticate user %s."
func FailedToAuthenticate(user string) string {
	return logHandler(fmt.Sprintf("Failed to authenticate user %s.", user), nil)
}

// "Failed to hash password."
func FailedToHashPassword(err error) string {
	return logHandler("Failed to hash password.", err)
}

// "Failed to generate salt."
func FailedToGenerateSalt(err error) string {
	return logHandler("Failed to generate salt.", err)
}

// "Failed to compare passwords."
func FailedToConvertPrimitive(err error) string {
	return logHandler("Failed to convert primitive.", err)
}

// "Failed to ping %s database."
func FailedToPingDB(base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to ping %s database.", base), err)
}

// "Failed to close connection to %s database."
func FailedToCloseConnection(base string, err error) string {
	return logHandler(fmt.Sprintf("Failed to close connection to %s database.", base), err)
}

// "Failed to parse body."
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

// %s found on %s database.
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

// "Duplicated entry %s on database."
func DuplicatedEntry(id string) string {
	return logHandler(fmt.Sprintf("Duplicated entry %s on database.", id), nil)
}

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(3)
	fullName := runtime.FuncForPC(pc).Name()
	slicedPath := strings.Split(fullName, "/")
	return slicedPath[len(slicedPath)-1] // remove o ponto no início da extensão
}

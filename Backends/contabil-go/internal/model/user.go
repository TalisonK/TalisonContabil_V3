package model

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/constants"
	"github.com/TalisonK/TalisonContabil/internal/database"
	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
	"github.com/TalisonK/TalisonContabil/pkg/tagError"
	"github.com/TalisonK/TalisonContabil/pkg/timeHandler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/argon2"
	_ "gorm.io/gorm"
)

func GetUsers() ([]domain.UserDTO, *tagError.TagError) {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if statusDbLocal {
		usersLocal, err := GetLocalUsers()

		if err != nil {
			logging.FailedToFindOnDB("All Users", constants.LOCAL, err)
			return nil, tagError.GetTagError(http.StatusBadRequest, err)
		}

		logging.FoundOnDB("All Users", constants.LOCAL)
		return usersLocal, nil
	}

	if statusDbCloud {
		usersCloud, err := GetCloudUsers()

		if err != nil {
			logging.FailedToFindOnDB("All Users", constants.CLOUD, err)
			return nil, tagError.GetTagError(http.StatusBadRequest, err)
		}

		return usersCloud, nil
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func CreateUser(user *domain.User) (*domain.UserDTO, *tagError.TagError) {

	if user.Name == "" || user.Password == "" {
		return nil, tagError.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.EmptyPassword()))
	}

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if _, err := FindUserByName(user.Name); err == nil {
		return nil, tagError.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.DuplicatedEntry("category")))
	}

	// Encrypt the password
	hash, salt, err := Hash(user.Password)

	if err != nil {
		logging.FailedToHashPassword(err)
		return nil, tagError.GetTagError(http.StatusInternalServerError, err)
	}

	user.Password = base64.StdEncoding.EncodeToString(hash)
	user.Salt = base64.StdEncoding.EncodeToString(salt)
	user.Role = "ROLE_USER"
	user.CreatedAt = timeHandler.GetTimeNow()
	user.UpdatedAt = timeHandler.GetTimeNow()

	if statusDbCloud {
		resultCloud, err := database.DBCloud.User.InsertOne(context.Background(), user.ToPrim())

		if err != nil {
			logging.FailedToCreateOnDB(user.Name, constants.CLOUD, err)
			return nil, tagError.GetTagError(http.StatusInternalServerError, err)
		}
		user.ID = resultCloud.InsertedID.(primitive.ObjectID).Hex()
		logging.CreatedOnDB(user.Name, constants.CLOUD)
	}

	if statusDbLocal && user.ID != "" {
		result := database.DBlocal.Create(&user)

		if result.Error != nil {
			logging.FailedToCreateOnDB(user.Name, constants.LOCAL, result.Error)
			return nil, tagError.GetTagError(http.StatusInternalServerError, result.Error)
		}
		logging.CreatedOnDB(user.Name, constants.LOCAL)

		dto := user.ToDTO()
		return &dto, nil
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func UpdateUser(user *domain.User) (*domain.UserDTO, *tagError.TagError) {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	baseUser, err := FindUserById(user.ID)

	if err != nil {
		return nil, tagError.GetTagError(http.StatusBadRequest, err.Inner)
	}

	if user.Name != "" {
		baseUser.Name = user.Name
	}
	if user.Password != "" {

		hash, salt, err := Hash(user.Password)

		if err != nil {
			return nil, tagError.GetTagError(http.StatusInternalServerError, err)
		}

		baseUser.Password = base64.StdEncoding.EncodeToString(hash)
		baseUser.Salt = base64.StdEncoding.EncodeToString(salt)
	}
	if user.Role != "" {
		baseUser.Role = user.Role
	}
	baseUser.UpdatedAt = timeHandler.GetTimeNow()

	user = baseUser

	if statusDbLocal {
		result := database.DBlocal.Save(user)
		if result.Error != nil {
			logging.FailedToUpdateOnDB(user.ID, constants.LOCAL, result.Error)
			return nil, tagError.GetTagError(http.StatusInternalServerError, result.Error)
		} else {
			logging.UpdatedOnDB(user.ID, constants.LOCAL)
		}
	}

	if statusDbCloud {
		userParse := user.ToPrim()

		id, _ := primitive.ObjectIDFromHex(user.ID)

		filter := bson.M{"_id": id}

		update := bson.M{"$set": userParse}

		_, err := database.DBCloud.User.UpdateOne(context.Background(), filter, update)
		if err != nil {
			logging.FailedToUpdateOnDB(user.ID, constants.CLOUD, err)
			return nil, tagError.GetTagError(http.StatusInternalServerError, err)
		} else {
			logging.UpdatedOnDB(user.ID, constants.CLOUD)
		}
	}

	dto := user.ToDTO()

	return &dto, nil
}

func DeleteUser(id string) *tagError.TagError {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	idParse, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logging.FailedToConvertPrimitive(err)
		return tagError.GetTagError(http.StatusInternalServerError, err)
	}

	if statusDbLocal {
		result := database.DBlocal.Delete(&domain.User{}, "id = ?", id)
		if result.Error != nil {
			logging.FailedToDeleteOnDB(id, constants.LOCAL, result.Error)
			return tagError.GetTagError(http.StatusInternalServerError, result.Error)
		} else {
			logging.DeletedOnDB(id, constants.LOCAL)
		}
	}

	if statusDbCloud {
		filter := bson.M{"_id": idParse}

		_, err = database.DBCloud.User.DeleteOne(context.Background(), filter)
		if err != nil {
			logging.FailedToDeleteOnDB(id, constants.CLOUD, err)
			return tagError.GetTagError(http.StatusInternalServerError, err)
		} else {
			logging.DeletedOnDB(id, constants.CLOUD)
		}
	}

	return nil
}

func LoginUser(user domain.User) (*domain.UserDTO, *tagError.TagError) {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if user.Name == "" || user.Password == "" {
		return nil, tagError.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.EmptyPassword()))
	}

	if statusDbLocal {
		baseUser := domain.User{}
		result := database.DBlocal.First(&baseUser, "name = ?", user.Name)

		if result.Error != nil {
			logging.FailedToAuthenticate(user.Name)
			return nil, tagError.GetTagError(http.StatusBadRequest, result.Error)
		} else {
			if Compare(user.Password, baseUser.Salt, baseUser.Password) {
				logging.FoundOnDB(user.Name, constants.LOCAL)
				dto := baseUser.ToDTO()
				return &dto, nil
			} else {
				return nil, tagError.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.FailedToAuthenticate(user.Name)))
			}
		}
	}

	if statusDbCloud {
		filter := bson.M{"name": user.Name}

		result := primitive.M{}
		database.DBCloud.User.FindOne(context.Background(), filter).Decode(&result)

		if len(result) == 0 {
			logging.FailedToFindOnDB(user.Name, constants.CLOUD, nil)
		} else {
			userCloud := domain.PrimToUser(result)

			if Compare(user.Password, userCloud.Salt, userCloud.Password) {
				logging.FoundOnDB(user.Name, constants.CLOUD)
				dto := userCloud.ToDTO()
				return &dto, nil
			} else {
				return nil, tagError.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.FailedToAuthenticate(user.Name)))
			}
		}
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

// findUserById retrieves a user by its ID
func FindUserById(id string) (*domain.User, *tagError.TagError) {
	var user domain.User = domain.User{}

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, tagError.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection())
	}

	if statusDbLocal {
		result := database.DBlocal.First(&user, "id = ?", id)

		if result.Error != nil {
			logging.FailedToFindOnDB(id, constants.LOCAL, result.Error)
			return &user, tagError.GetTagError(http.StatusBadRequest, result.Error)
			//TODO: equalize bases de users
		} else {
			logging.FoundOnDB(id, constants.LOCAL)
			return &user, nil
		}
	}

	if statusDbCloud {
		idParse, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			logging.FailedToConvertPrimitive(err)
			return &user, tagError.GetTagError(http.StatusInternalServerError, err)
		}

		filter := bson.M{"_id": idParse}
		result := primitive.M{}
		database.DBCloud.User.FindOne(context.Background(), filter).Decode(&result)

		if len(result) == 0 {
			logging.FailedToFindOnDB(id, constants.CLOUD, nil)
		} else {
			user = domain.PrimToUser(result)
			logging.FoundOnDB(id, constants.CLOUD)
			return &user, nil
		}
	}

	return nil, tagError.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred())
}

func FindUserByName(name string) (*domain.User, error) {
	var user domain.User = domain.User{}

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, logging.NoDatabaseConnection()
	}

	if statusDbLocal {
		result := database.DBlocal.First(&user, "name = ?", name)

		if result.Error != nil {
			logging.FailedToFindOnDB(name, constants.LOCAL, result.Error)
			return &user, result.Error
		} else {
			logging.FoundOnDB(name, constants.LOCAL)
			return &user, nil
		}
	}

	if statusDbCloud {
		filter := bson.M{"name": name}
		result := primitive.M{}
		database.DBCloud.User.FindOne(context.Background(), filter).Decode(&result)

		if len(result) == 0 {
			logging.FailedToFindOnDB(name, constants.CLOUD, nil)
		} else {
			user = domain.PrimToUser(result)
			logging.FoundOnDB(name, constants.CLOUD)
			return &user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

// getCloudUsers retrieves all users from the cloud database
func GetCloudUsers() ([]domain.UserDTO, error) {
	colect, err := database.DBCloud.User.Find(context.TODO(), &bson.D{})

	result := []domain.UserDTO{}

	//colect.All(context.Background(), &result)

	for colect.Next(context.Background()) {
		var user bson.M
		colect.Decode(&user)

		u := domain.PrimToUser(user)

		result = append(result, u.ToDTO())
	}

	if err != nil {
		logging.FailedToFindOnDB("All Users", constants.CLOUD, err)
		return nil, err
	}
	return result, nil
}

// getLocalUsers retrieves all users from the local database
func GetLocalUsers() ([]domain.UserDTO, error) {
	var users []domain.User
	result := database.DBlocal.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	userDto := []domain.UserDTO{}

	for i := range users {
		userDto = append(userDto, users[i].ToDTO())

	}

	return userDto, nil
}

// Hash the password
func Hash(password string) ([]byte, []byte, error) {

	salt := make([]byte, 16)

	_, err := rand.Read(salt)

	if err != nil {
		logging.FailedToGenerateSalt(err)
		return nil, nil, err
	}

	hash := hasher(password, salt)

	return hash, salt, nil
}

func Compare(password string, salt string, origin string) bool {

	salter, err := base64.StdEncoding.DecodeString(salt)

	if err != nil {
		logging.FailedToGenerateSalt(err)
		return false
	}

	hash := hasher(password, salter)

	old, _ := base64.StdEncoding.DecodeString(origin)

	if bytes.Equal(hash, old) {
		return true
	} else {
		return false
	}
}

func hasher(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, 1, 128*1024, 4, 64)
}

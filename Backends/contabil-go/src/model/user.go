package model

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/argon2"
	_ "gorm.io/gorm"
)

func GetUsers() ([]domain.UserDTO, *util.TagError) {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.GetUsers"))
	}

	var users []domain.UserDTO

	if statusDbLocal {
		usersLocal, err := GetLocalUsers()

		if err != nil {
			logging.FailedToFindOnDB("All Users", "Local", err, "model.GetUsers")
			return nil, util.GetTagError(http.StatusBadRequest, err)
		}

		users = append(users, usersLocal...)
		return users, nil
	}

	if statusDbCloud {
		usersCloud, err := GetCloudUsers()

		if err != nil {
			logging.FailedToFindOnDB("All Users", "Cloud", err, "model.GetUsers")
			return nil, util.GetTagError(http.StatusBadRequest, err)
		}

		users = append(users, usersCloud...)
		return users, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred("model.GetUsers"))
}

func CreateUser(user *domain.User) (*domain.UserDTO, *util.TagError) {

	if user.Name == "" || user.Password == "" {
		return nil, util.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.EmptyPassword("model.CreateUser")))
	}

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.GetUsers"))
	}

	if _, err := FindUserByName(user.Name); err == nil {
		return nil, util.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.DuplicatedEntry("category", "model.CreateUser")))
	}

	// Encrypt the password
	hash, salt, err := Hash(user.Password)

	if err != nil {
		logging.FailedToHashPassword(err, "model.CreateUser")
		return nil, util.GetTagError(http.StatusInternalServerError, err)
	}

	user.Password = base64.StdEncoding.EncodeToString(hash)
	user.Salt = base64.StdEncoding.EncodeToString(salt)
	user.Role = "ROLE_USER"
	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.UpdatedAt = time.Now().Format(time.RFC3339)

	if statusDbCloud {
		resultCloud, err := database.DBCloud.User.InsertOne(context.Background(), user.ToPrim())

		if err != nil {
			logging.FailedToCreateOnDB(user.Name, "Cloud", err, "model.CreateUser")
			return nil, util.GetTagError(http.StatusInternalServerError, err)
		}
		user.ID = resultCloud.InsertedID.(primitive.ObjectID).Hex()
		logging.CreatedOnDB(user.Name, "Cloud", "model.CreateUser")
	}

	if statusDbLocal && user.ID != "" {
		result := database.DBlocal.Create(&user)

		if result.Error != nil {
			logging.FailedToCreateOnDB(user.Name, "Local", result.Error, "model.CreateUser")
			return nil, util.GetTagError(http.StatusInternalServerError, result.Error)
		}
		logging.CreatedOnDB(user.Name, "Local", "model.CreateUser")

		dto := user.ToDTO()
		return &dto, nil
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred("model.CreateUser"))
}

func UpdateUser(user *domain.User) (*domain.UserDTO, *util.TagError) {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.UpdateUser"))
	}

	baseUser, err := FindUserById(user.ID)

	if err != nil {
		return nil, util.GetTagError(http.StatusBadRequest, err.Inner)
	}

	if user.Name != "" {
		baseUser.Name = user.Name
	}
	if user.Password != "" {

		hash, salt, err := Hash(user.Password)

		if err != nil {
			return nil, util.GetTagError(http.StatusInternalServerError, err)
		}

		baseUser.Password = base64.StdEncoding.EncodeToString(hash)
		baseUser.Salt = base64.StdEncoding.EncodeToString(salt)
	}
	if user.Role != "" {
		baseUser.Role = user.Role
	}
	baseUser.UpdatedAt = time.Now().Format(time.RFC3339)

	user = baseUser

	if statusDbLocal {
		result := database.DBlocal.Save(user)
		if result.Error != nil {
			logging.FailedToUpdateOnDB(user.ID, "Local", result.Error, "model.UpdateUser")
			return nil, util.GetTagError(http.StatusInternalServerError, result.Error)
		} else {
			logging.UpdatedOnDB(user.ID, "Local", "model.UpdateUser")
		}
	}

	if statusDbCloud {
		userParse := user.ToPrim()

		id, _ := primitive.ObjectIDFromHex(user.ID)

		filter := bson.M{"_id": id}

		update := bson.M{"$set": userParse}

		_, err := database.DBCloud.User.UpdateOne(context.Background(), filter, update)
		if err != nil {
			logging.FailedToUpdateOnDB(user.ID, "Cloud", err, "model.UpdateUser")
			return nil, util.GetTagError(http.StatusInternalServerError, err)
		} else {
			logging.UpdatedOnDB(user.ID, "Cloud", "model.UpdateUser")
		}
	}

	dto := user.ToDTO()

	return &dto, nil
}

func DeleteUser(id string) *util.TagError {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.DeleteUser"))
	}

	idParse, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logging.FailedToConvertPrimitive(err, "model.DeleteUser")
		return util.GetTagError(http.StatusInternalServerError, err)
	}

	if statusDbLocal {
		result := database.DBlocal.Delete(&domain.User{}, "id = ?", id)
		if result.Error != nil {
			logging.FailedToDeleteOnDB(id, "Local", result.Error, "model.DeleteUser")
			return util.GetTagError(http.StatusInternalServerError, result.Error)
		} else {
			logging.DeletedOnDB(id, "Local", "model.DeleteUser")
		}
	}

	if statusDbCloud {
		filter := bson.M{"_id": idParse}

		_, err = database.DBCloud.User.DeleteOne(context.Background(), filter)
		if err != nil {
			logging.FailedToDeleteOnDB(id, "Cloud", err, "model.DeleteUser")
			return util.GetTagError(http.StatusInternalServerError, err)
		} else {
			logging.DeletedOnDB(id, "Cloud", "model.DeleteUser")
		}
	}

	return nil
}

func LoginUser(user domain.User) (*domain.User, *util.TagError) {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.LoginUser"))
	}

	if user.Name == "" || user.Password == "" {
		return nil, util.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.EmptyPassword("model.LoginUser")))
	}

	if statusDbLocal {
		baseUser := domain.User{}
		result := database.DBlocal.First(&baseUser, "name = ?", user.Name)

		if result.Error != nil {
			logging.FailedToAuthenticate(user.Name, "model.LoginUser")
			return nil, util.GetTagError(http.StatusBadRequest, result.Error)
		} else {
			if user.Password == baseUser.Password {
				logging.FoundOnDB(user.Name, "Local", "model.LoginUser")
				return &user, nil
			}
		}
	}

	if statusDbCloud {
		filter := bson.M{"name": user.Name}

		result := primitive.M{}
		database.DBCloud.User.FindOne(context.Background(), filter).Decode(&result)

		if len(result) == 0 {
			logging.FailedToFindOnDB(user.Name, "Cloud", nil, "model.LoginUser")
		} else {
			userCloud := domain.PrimToUser(result)

			if Compare(user.Password, userCloud.Salt, userCloud.Password) {
				logging.FoundOnDB(user.Name, "Cloud", "model.LoginUser")
				return &userCloud, nil
			} else {
				return nil, util.GetTagError(http.StatusBadRequest, fmt.Errorf(logging.FailedToAuthenticate(user.Name, "model.LoginUser")))
			}
		}
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred("model.LoginUser"))
}

// findUserById retrieves a user by its ID
func FindUserById(id string) (*domain.User, *util.TagError) {
	var user domain.User = domain.User{}

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, util.GetTagError(http.StatusInternalServerError, logging.NoDatabaseConnection("model.FindUserById"))
	}

	if statusDbLocal {
		result := database.DBlocal.First(&user, "id = ?", id)

		if result.Error != nil {
			logging.FailedToFindOnDB(id, "Local", result.Error, "model.FindUserById")
			return &user, util.GetTagError(http.StatusBadRequest, result.Error)
			//TODO: equalize bases de users
		} else {
			logging.FoundOnDB(id, "Local", "model.FindUserById")
			return &user, nil
		}
	}

	if statusDbCloud {
		idParse, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			logging.FailedToConvertPrimitive(err, "model.FindUserById")
			return &user, util.GetTagError(http.StatusInternalServerError, err)
		}

		filter := bson.M{"_id": idParse}
		result := primitive.M{}
		database.DBCloud.User.FindOne(context.Background(), filter).Decode(&result)

		if len(result) == 0 {
			logging.FailedToFindOnDB(id, "Cloud", nil, "model.FindUserById")
		} else {
			user = domain.PrimToUser(result)
			logging.FoundOnDB(id, "Cloud", "model.FindUserById")
			return &user, nil
		}
	}

	return nil, util.GetTagError(http.StatusInternalServerError, logging.ErrorOccurred("model.FindUserById"))
}

func FindUserByName(name string) (*domain.User, error) {
	var user domain.User = domain.User{}

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, logging.NoDatabaseConnection("model.FindUserByName")
	}

	if statusDbLocal {
		result := database.DBlocal.First(&user, "name = ?", name)

		if result.Error != nil {
			logging.FailedToFindOnDB(name, "Local", result.Error, "model.FindUserByName")
			return &user, result.Error
		} else {
			logging.FoundOnDB(name, "Local", "model.FindUserByName")
			return &user, nil
		}
	}

	if statusDbCloud {
		filter := bson.M{"name": name}
		result := primitive.M{}
		database.DBCloud.User.FindOne(context.Background(), filter).Decode(&result)

		if len(result) == 0 {
			logging.FailedToFindOnDB(name, "Cloud", nil, "model.FindUserByName")
		} else {
			user = domain.PrimToUser(result)
			logging.FoundOnDB(name, "Cloud", "model.FindUserByName")
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
		logging.FailedToFindOnDB("All Users", "Cloud", err, "model.GetCloudUsers")
		return nil, err
	}
	return result, nil
}

// getLocalUsers retrieves all users from the local database
func GetLocalUsers() ([]domain.UserDTO, error) {
	var users []domain.UserDTO
	result := database.DBlocal.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Hash the password
func Hash(password string) ([]byte, []byte, error) {

	salt := make([]byte, 16)

	_, err := rand.Read(salt)

	if err != nil {
		logging.FailedToGenerateSalt(err, "model.Hash")
		return nil, nil, err
	}

	hash := hasher(password, salt)

	return hash, salt, nil
}

func Compare(password string, salt string, origin string) bool {

	salter, err := base64.StdEncoding.DecodeString(salt)

	if err != nil {
		logging.FailedToGenerateSalt(err, "model.Compare")
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

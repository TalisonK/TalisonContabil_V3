package model

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/domain"
	"github.com/TalisonK/TalisonContabil/src/util/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/argon2"
	_ "gorm.io/gorm"
)

func GetUsers() ([]domain.UserDTO, error) {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, fmt.Errorf(logging.NoDatabaseConnection("model.GetUsers"))
	}

	var users []domain.UserDTO

	if statusDbLocal {
		usersLocal, err := GetLocalUsers()

		if err != nil {
			logging.FailedToFindOnDB("All Users", "Local", err, "model.GetUsers")
			return nil, err
		}

		users = append(users, usersLocal...)
		return users, nil
	}

	if statusDbCloud {
		usersCloud, err := GetCloudUsers()

		if err != nil {
			logging.FailedToFindOnDB("All Users", "Cloud", err, "model.GetUsers")
			return nil, err
		}

		users = append(users, usersCloud...)
		return users, nil
	}

	return nil, fmt.Errorf(logging.ErrorOccurred("model.GetUsers"))
}

func CreateUser(user *domain.User) (*domain.UserDTO, error) {

	if user.Name == "" || user.Password == "" {
		return nil, fmt.Errorf(logging.EmptyPassword("model.CreateUser"))
	}

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, fmt.Errorf(logging.NoDatabaseConnection("model.CreateUser"))
	}

	if _, err := FindUserByName(user.Name); err == nil {
		return nil, fmt.Errorf(logging.DuplicatedEntry(user.Name, "model.CreateUser"))
	}

	// Encrypt the password
	hash, salt, err := Hash(user.Password)

	if err != nil {
		logging.FailedToHashPassword(err, "model.CreateUser")
		return nil, err
	}

	user.Password = base64.StdEncoding.EncodeToString(hash)
	user.Salt = base64.StdEncoding.EncodeToString(salt)
	user.Role = "ROLE_USER"
	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.UpdatedAt = time.Now().Format(time.RFC3339)

	if statusDbCloud {
		resultCloud, err := database.DBCloud.User.InsertOne(context.Background(), UserToPrim(*user))

		if err != nil {
			logging.FailedToCreateOnDB(user.Name, "Cloud", err, "model.CreateUser")
			return nil, err
		}
		user.ID = resultCloud.InsertedID.(primitive.ObjectID).Hex()
		logging.CreatedOnDB(user.Name, "Cloud", "model.CreateUser")
	}

	if statusDbLocal && user.ID != "" {
		result := database.DBlocal.Create(&user)

		if result.Error != nil {
			logging.FailedToCreateOnDB(user.Name, "Local", result.Error, "model.CreateUser")
			return nil, result.Error
		}
		logging.CreatedOnDB(user.Name, "Local", "model.CreateUser")
		return ToDTO(user), nil
	}

	return nil, fmt.Errorf("algo deu errado")
}

func UpdateUser(user *domain.User) (*domain.UserDTO, error) {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, fmt.Errorf(logging.NoDatabaseConnection("model.UpdateUser"))
	}

	baseUser, err := FindUserById(user.ID)

	if err != nil {
		return nil, fmt.Errorf(logging.FailedToFindOnDB(user.ID, "User", err, "model.UpdateUser"))
	}

	if user.Name != "" {
		baseUser.Name = user.Name
	}
	if user.Password != "" {

		hash, salt, err := Hash(user.Password)

		if err != nil {
			return nil, fmt.Errorf(logging.FailedToHashPassword(err, "model.UpdateUser"))
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
			return nil, result.Error
		} else {
			logging.UpdatedOnDB(user.ID, "Local", "model.UpdateUser")
		}
	}

	if statusDbCloud {
		userParse := UserToPrim(*user)

		id, _ := primitive.ObjectIDFromHex(user.ID)

		filter := bson.M{"_id": id}

		update := bson.M{"$set": userParse}

		_, err = database.DBCloud.User.UpdateOne(context.Background(), filter, update)
		if err != nil {
			logging.FailedToUpdateOnDB(user.ID, "Cloud", err, "model.UpdateUser")
			return nil, err
		} else {
			logging.UpdatedOnDB(user.ID, "Cloud", "model.UpdateUser")
		}
	}

	return ToDTO(user), nil
}

func DeleteUser(id string) error {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return fmt.Errorf(logging.NoDatabaseConnection("model.DeleteUser"))
	}

	idParse, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logging.FailedToConvertPrimitive(err, "model.DeleteUser")
		return err
	}

	if statusDbLocal {
		result := database.DBlocal.Delete(&domain.User{}, "id = ?", id)
		if result.Error != nil {
			logging.FailedToDeleteOnDB(id, "Local", result.Error, "model.DeleteUser")
			return result.Error
		} else {
			logging.DeletedOnDB(id, "Local", "model.DeleteUser")
		}
	}

	if statusDbCloud {
		filter := bson.M{"_id": idParse}

		_, err = database.DBCloud.User.DeleteOne(context.Background(), filter)
		if err != nil {
			logging.FailedToDeleteOnDB(id, "Cloud", err, "model.DeleteUser")
			return err
		} else {
			logging.DeletedOnDB(id, "Cloud", "model.DeleteUser")
		}
	}

	return nil
}

func LoginUser(user domain.User) (*domain.User, error) {

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, fmt.Errorf(logging.NoDatabaseConnection("model.LoginUser"))
	}

	if user.Name == "" || user.Password == "" {
		return nil, fmt.Errorf(logging.EmptyPassword("model.LoginUser"))
	}

	if statusDbLocal {
		baseUser := domain.User{}
		result := database.DBlocal.First(&baseUser, "name = ?", user.Name)

		if result.Error != nil {
			logging.FailedToAuthenticate(user.Name, "model.LoginUser")
			return nil, result.Error
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
			userCloud := PrimToUser(result)

			if Compare(user.Password, userCloud.Salt, userCloud.Password) {
				logging.FoundOnDB(user.Name, "Cloud", "model.LoginUser")
				return &userCloud, nil
			} else {
				return nil, fmt.Errorf("wrong password")
			}
		}
	}

	return nil, fmt.Errorf("user not found")
}

// findUserById retrieves a user by its ID
func FindUserById(id string) (*domain.User, error) {
	var user domain.User = domain.User{}

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, fmt.Errorf(logging.NoDatabaseConnection("model.FindUserById"))
	}

	if statusDbLocal {
		result := database.DBlocal.First(&user, "id = ?", id)

		if result.Error != nil {
			logging.FailedToFindOnDB(id, "Local", result.Error, "model.FindUserById")
			return &user, result.Error
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
			return &user, err
		}

		filter := bson.M{"_id": idParse}
		result := primitive.M{}
		database.DBCloud.User.FindOne(context.Background(), filter).Decode(&result)

		if len(result) == 0 {
			logging.FailedToFindOnDB(id, "Cloud", nil, "model.FindUserById")
		} else {
			user = PrimToUser(result)
			logging.FoundOnDB(id, "Cloud", "model.FindUserById")
			return &user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

func FindUserByName(name string) (*domain.User, error) {
	var user domain.User = domain.User{}

	statusDbLocal, statusDbCloud := database.CheckDBStatus()

	if !statusDbLocal && !statusDbCloud {
		return nil, fmt.Errorf(logging.NoDatabaseConnection("model.FindUserByName"))
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
			user = PrimToUser(result)
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

		u := PrimToUser(user)

		result = append(result, *ToDTO(&u))
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

// primToUser converts a primitive.M to a User
func PrimToUser(user primitive.M) domain.User {
	var usuario domain.User

	usuario.ID = user["_id"].(primitive.ObjectID).Hex()
	usuario.Name = user["name"].(string)
	usuario.Password = user["password"].(string)
	usuario.Role = user["role"].(string)
	usuario.Salt = user["salt"].(string)
	usuario.CreatedAt = user["createdAt"].(primitive.DateTime).Time().Format(time.RFC3339)
	usuario.UpdatedAt = user["updatedAt"].(primitive.DateTime).Time().Format(time.RFC3339)

	return usuario
}

// userToPrim converts a User to a primitive.M
func UserToPrim(user domain.User) primitive.M {
	usuario := primitive.M{}

	if user.ID != "" {
		usuario["_id"], _ = primitive.ObjectIDFromHex(user.ID)
	}
	if user.Name != "" {
		usuario["name"] = user.Name
	}
	if user.Password != "" {
		usuario["password"] = user.Password
	}
	if user.Role != "" {
		usuario["role"] = user.Role
	}
	createdAt, _ := time.Parse(time.RFC3339, user.CreatedAt)
	usuario["salt"] = user.Salt
	usuario["createdAt"] = primitive.NewDateTimeFromTime(createdAt)
	updatedAt, _ := time.Parse(time.RFC3339, user.UpdatedAt)
	usuario["updatedAt"] = primitive.NewDateTimeFromTime(updatedAt)

	return usuario
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

func ToDTO(user *domain.User) *domain.UserDTO {
	return &domain.UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

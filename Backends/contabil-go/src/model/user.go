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
	"github.com/TalisonK/TalisonContabil/src/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/argon2"
	_ "gorm.io/gorm"
)

func CreateUser(user *domain.User) (*domain.UserDTO, error) {

	if user.Name == "" || user.Password == "" {
		util.LogHandler("User name or password is empty.", nil, "model.CreateUser")
		return nil, fmt.Errorf("user name or password is empty")
	}

	statusDbLocal := database.CheckLocalDB()
	statusDbCloud := database.CheckCloudDB()

	if !statusDbLocal && !statusDbCloud {
		util.LogHandler("No database connection available.", nil, "model.CreateUser")
		return nil, fmt.Errorf("no database connection available")
	}

	if _, err := FindUserByName(user.Name); err == nil {
		util.LogHandler(fmt.Sprintf("User %s already exists.", user.Name), nil, "model.CreateUser")
		return nil, fmt.Errorf("user already exists")
	}

	// Encrypt the password
	hash, salt, err := Encrypt(user.Password)

	if err != nil {
		util.LogHandler("Failed to encrypt password.", err, "model.CreateUser")
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
			util.LogHandler("Failed to insert user in cloud database.", err, "model.CreateUser")
			return nil, err
		}
		user.ID = resultCloud.InsertedID.(primitive.ObjectID).Hex()
		util.LogHandler("User successfully inserted in cloud database.", nil, "model.CreateUser")
	}

	if statusDbLocal && user.ID != "" {
		result := database.DBlocal.Create(&user)

		if result.Error != nil {
			util.LogHandler("Failed to insert user in local database.", result.Error, "model.CreateUser")
			return nil, result.Error
		}
		util.LogHandler("User successfully inserted in local database.", nil, "model.CreateUser")
		return ToDTO(user), nil
	}

	return nil, fmt.Errorf("algo deu errado")
}

func UpdateUser(user *domain.User) (*domain.UserDTO, error) {

	statusDbLocal := database.CheckLocalDB()
	statusDbCloud := database.CheckCloudDB()

	if !statusDbLocal && !statusDbCloud {
		util.LogHandler("No database connection available.", nil, "model.UpdateUser")
		return nil, fmt.Errorf("no database connection available")
	}

	baseUser, err := FindUserById(user.ID)

	if err != nil {
		util.LogHandler("Fail to find user", err, "model.UpdateUser")
		return nil, fmt.Errorf("fail to find user")
	}

	if user.Name != "" {
		baseUser.Name = user.Name
	}
	if user.Password != "" {

		hash, salt, err := Encrypt(user.Password)

		if err != nil {
			util.LogHandler("Fail to encrypt new password", err, "model.UpdateUser")
			return nil, fmt.Errorf("fail to encrypt new password")
		}

		baseUser.Password = base64.StdEncoding.EncodeToString(hash)
		baseUser.Salt = base64.StdEncoding.EncodeToString(salt)
	}
	if user.Role != "" {
		baseUser.Role = user.Role
	}
	baseUser.UpdatedAt = time.Now().Format(time.RFC3339)

	user = &baseUser

	if statusDbCloud {
		userParse := UserToPrim(*user)

		id, _ := primitive.ObjectIDFromHex(user.ID)

		filter := bson.M{"_id": id}

		update := bson.M{"$set": userParse}

		_, err = database.DBCloud.User.UpdateOne(context.Background(), filter, update)
		if err != nil {
			util.LogHandler("Failed to update user in cloud database.", err, "model.UpdateUser")
			return nil, err
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully updated in cloud database.", user.ID), nil, "model.UpdateUser")
		}
	}

	if statusDbLocal {
		result := database.DBlocal.Save(user)
		if result.Error != nil {
			util.LogHandler("Failed to update user in local database.", result.Error, "model.UpdateUser")
			return nil, result.Error
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully updated in local database.", user.ID), nil, "model.UpdateUser")
		}
	}

	return ToDTO(user), nil
}

func DeleteUser(id string) error {

	statusDbLocal := database.CheckLocalDB()
	statusDbCloud := database.CheckCloudDB()

	if !statusDbLocal && !statusDbCloud {
		util.LogHandler("No database connection available.", nil, "model.DeleteUser")
		return fmt.Errorf("no database connection available")
	}

	idParse, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		util.LogHandler(fmt.Sprintf("Failed to parse user ID %s.", id), err, "model.DeleteUser")
		return err
	}

	if statusDbCloud {
		filter := bson.M{"_id": idParse}

		_, err = database.DBCloud.User.DeleteOne(context.Background(), filter)
		if err != nil {
			util.LogHandler("Failed to delete user in cloud database.", err, "model.DeleteUser")
			return err
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully deleted in cloud database.", id), nil, "model.DeleteUser")
		}
	}

	if statusDbLocal {
		result := database.DBlocal.Delete(&domain.User{}, "id = ?", id)
		if result.Error != nil {
			util.LogHandler("Failed to delete user in local database.", result.Error, "model.DeleteUser")
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully deleted in local database.", id), nil, "model.DeleteUser")
		}
	}

	return nil
}

func LoginUser(user domain.User) (domain.User, error) {

	statusDbLocal := database.CheckLocalDB()
	statusDbCloud := database.CheckCloudDB()

	if !statusDbLocal && !statusDbCloud {
		util.LogHandler("No database connection available.", nil, "model.LoginUser")
		return domain.User{}, fmt.Errorf("no database connection available")
	}

	if statusDbCloud {
		filter := bson.M{"name": user.Name}

		result := primitive.M{}
		database.DBCloud.User.FindOne(context.Background(), filter).Decode(&result)

		if len(result) == 0 {
			util.LogHandler(fmt.Sprintf("User %s not found in cloud database.", user.Name), nil, "model.LoginUser")
		} else {
			userCloud := PrimToUser(result)

			if Compare(user.Password, userCloud.Salt, userCloud.Password) {
				util.LogHandler(fmt.Sprintf("User %s successfully logged in cloud database.", user.Name), nil, "model.LoginUser")
				return userCloud, nil
			} else {
				return domain.User{}, fmt.Errorf("wrong password")
			}
		}
	}

	if database.CheckLocalDB() {
		result := database.DBlocal.First(&user, "name = ?", user.Name)

		if result.Error != nil {
			util.LogHandler("Failed to find user in local database.", result.Error, "model.LoginUser")
			return domain.User{}, result.Error
		} else {
			if user.Password == user.Password {
				util.LogHandler(fmt.Sprintf("User %s successfully logged in local database.", user.Name), nil, "model.LoginUser")
				return user, nil
			}
		}
	}

	return domain.User{}, fmt.Errorf("user not found")
}

// findUserById retrieves a user by its ID
func FindUserById(id string) (domain.User, error) {
	var user domain.User = domain.User{}

	statusDbLocal := database.CheckLocalDB()
	statusDbCloud := database.CheckCloudDB()

	if !statusDbLocal && !statusDbCloud {
		util.LogHandler("No database connection available.", nil, "findUserById")
		return user, fmt.Errorf("no database connection available")
	}

	if statusDbCloud {
		idParse, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			util.LogHandler(fmt.Sprintf("Failed to parse user ID %s.", id), err, "findUserById")
			return user, err
		}

		filter := bson.M{"_id": idParse}
		result := primitive.M{}
		database.DBCloud.User.FindOne(context.Background(), filter).Decode(&result)

		if len(result) == 0 {
			util.LogHandler(fmt.Sprintf("User %s not found in cloud database.", id), nil, "findUserById")
		} else {
			user = PrimToUser(result)
			util.LogHandler(fmt.Sprintf("User %s successfully found in cloud database.", id), nil, "findUserById")
			return user, nil
		}
	}

	if database.CheckLocalDB() {
		result := database.DBlocal.First(&user, "id = ?", id)

		if result.Error != nil {
			util.LogHandler("Failed to find user in local database.", result.Error, "findUserById")
			return user, result.Error
			//TODO: equalize bases de users
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully found in local database.", id), nil, "findUserById")
			return user, nil
		}
	}

	return user, fmt.Errorf("user not found")
}

func FindUserByName(name string) (domain.User, error) {
	var user domain.User = domain.User{}

	statusDbLocal := database.CheckLocalDB()
	statusDbCloud := database.CheckCloudDB()

	if !statusDbLocal && !statusDbCloud {
		util.LogHandler("No database connection available.", nil, "findUserByName")
		return user, fmt.Errorf("no database connection available")
	}

	if statusDbCloud {
		filter := bson.M{"name": name}
		result := primitive.M{}
		database.DBCloud.User.FindOne(context.Background(), filter).Decode(&result)

		if len(result) == 0 {
			util.LogHandler(fmt.Sprintf("User %s not found in cloud database.", name), nil, "findUserByName")
		} else {
			user = PrimToUser(result)
			util.LogHandler(fmt.Sprintf("User %s successfully found in cloud database.", name), nil, "findUserByName")
			return user, nil
		}
	}

	if database.CheckLocalDB() {
		result := database.DBlocal.First(&user, "name = ?", name)

		if result.Error != nil {
			util.LogHandler("Failed to find user in local database.", result.Error, "findUserByName")
			return user, result.Error
		} else {
			util.LogHandler(fmt.Sprintf("User %s successfully found in local database.", name), nil, "findUserByName")
			return user, nil
		}
	}

	return user, fmt.Errorf("user not found")
}

// getCloudUsers retrieves all users from the cloud database
func GetCloudUsers() ([]domain.User, error) {
	colect, err := database.DBCloud.User.Find(context.TODO(), &bson.D{})

	result := []domain.User{}

	//colect.All(context.Background(), &result)

	for colect.Next(context.Background()) {
		var user bson.M
		colect.Decode(&user)

		result = append(result, PrimToUser(user))
	}

	if err != nil {
		util.LogHandler("Failed to get users from cloud database.", err, "getUsers")
		return nil, err
	}
	return result, nil
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

// getLocalUsers retrieves all users from the local database
func GetLocalUsers() ([]domain.User, error) {
	var users []domain.User
	result := database.DBlocal.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Hash the password
func Encrypt(password string) ([]byte, []byte, error) {

	salt := make([]byte, 16)

	_, err := rand.Read(salt)

	if err != nil {
		util.LogHandler("Error creating cryptograpy steps", err, "hashPass")
		return nil, nil, err
	}

	hash := hasher(password, salt)

	return hash, salt, nil
}

func Compare(password string, salt string, origin string) bool {

	salter, err := base64.StdEncoding.DecodeString(salt)

	if err != nil {
		util.LogHandler("Failed to decode salt.", err, "model.Compare")
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

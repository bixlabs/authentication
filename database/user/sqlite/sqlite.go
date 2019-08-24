package sqlite

import (
	"fmt"
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/mappers"
	"github.com/bixlabs/authentication/database/model"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
)

type sqliteStorage struct {
	db       *gorm.DB
	Name     string `env:"AUTH_SERVER_DATABASE_NAME" envDefault:"sqlite.s3db"`
	User     string `env:"AUTH_SERVER_DATABASE_USER"`
	Password string `env:"AUTH_SERVER_DATABASE_PASSWORD"`
	Salt     string `env:"AUTH_SERVER_DATABASE_SALT"`
}

func NewSqliteStorage() (user.Repository, func()) {
	db := sqliteStorage{}
	err := env.Parse(&db)
	if err != nil {
		tools.Log().Panic("Parsing the env variables for the database failed", err)
	}
	closeDB := db.initialize()
	return db, closeDB
}

func (storage *sqliteStorage) initialize() func() {
	db := openDatabase(storage)
	db.AutoMigrate(&model.User{})
	storage.db = db

	return func() {
		_ = storage.db.Close()
	}
}

func openDatabase(storage *sqliteStorage) *gorm.DB {
	db, err := gorm.Open("sqlite3", storage.getConnectionString())
	if err != nil {
		panic(err)
	}
	storage.db = db
	return db
}

func (storage sqliteStorage) getConnectionString() string {
	// TODO: I'm not sure the authentication is working as we expect here, I'm sure in development this is not
	//  working but when creating a build it might be working as expected we need to ensure this later.
	return fmt.Sprintf("file:%s?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=ssha512&_auth_salt=%s",
		storage.Name, storage.User, storage.Password, storage.Salt)
}

func (storage sqliteStorage) Create(user structures.User) (structures.User, error) {
	modelForCreate := mappers.UserToDatabaseModel(user)
	transaction := storage.db.Begin()
	if err := transaction.Create(&modelForCreate).Error; err != nil {
		transaction.Rollback()
		return structures.User{}, err
	}

	if err := transaction.Commit().Error; err != nil {
		return structures.User{}, err
	}

	return storage.Find(user.Email)

}

func (storage sqliteStorage) Find(email string) (structures.User, error) {
	var account model.User
	if err := storage.db.First(&account, "email = ?", email).Error; err != nil {
		return structures.User{}, err
	}
	return mappers.DatabaseModelToUser(account), nil
}

func (storage sqliteStorage) IsEmailAvailable(email string) (bool, error) {
	_, err := storage.Find(email)
	if gorm.IsRecordNotFoundError(err) {
		return true, nil
	}
	return false, err
}

func (storage sqliteStorage) GetHashedPassword(email string) (string, error) {
	account, err := storage.Find(email)
	if err != nil {
		return "", err
	}

	return account.Password, nil

}

func (storage sqliteStorage) ChangePassword(email, newPassword string) error {
	transaction := storage.db.Begin()
	if err := transaction.Model(&model.User{Email: email}).Update("password", newPassword).Error; err != nil {
		transaction.Rollback()
		return err
	}

	return transaction.Commit().Error
}

func (storage sqliteStorage) UpdateResetToken(email, resetToken string) error {
	transaction := storage.db.Begin()
	if err := transaction.Model(&model.User{Email: email}).Update("reset_token", resetToken).Error; err != nil {
		transaction.Rollback()
		return err
	}

	return transaction.Commit().Error
}

func (storage sqliteStorage) Delete(user structures.User) error {
	transaction := storage.db.Begin()
	if err := transaction.Delete(&user).Error; err != nil {
		transaction.Rollback()
		return err
	}

	return transaction.Commit().Error
}

func (storage sqliteStorage) Update(email string, updateAttrs structures.User) (structures.User, error) {
	user, err := storage.Find(email)

	if err != nil {
		return structures.User{}, err
	}

	transaction := storage.db.Begin()
	modelForUpdate := mappers.UserToDatabaseModel(user)
	modelUpdateAttrs := mappers.UserToDatabaseModel(updateAttrs)

	if err := transaction.Model(&modelForUpdate).Update(modelUpdateAttrs).Error; err != nil {
		transaction.Rollback()
		return structures.User{}, err
	}

	if err := transaction.Commit().Error; err != nil {
		return structures.User{}, err
	}

	return mappers.DatabaseModelToUser(modelForUpdate), nil
}

package sqlite

import (
	"fmt"
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)
import _ "github.com/mattn/go-sqlite3"

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
	return db, db.initialize()
}

func (storage *sqliteStorage) initialize() func() {
	db := openDatabase(storage)
	db.AutoMigrate(&User{})
	storage.db = db
	storage.testDatabase()

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
	// TODO: I'm not sure the authentication is working as we expect here, I'm sure in development this is not working but in the build it might be working as expected we need to ensure this later.
	return fmt.Sprintf("file:%s?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=ssha512&_auth_salt=%s",
		storage.Name, storage.User, storage.Password, storage.Salt)
}

func (storage sqliteStorage) testDatabase() {

	// Create
	storage.db.Create(&User{Email: "jarrieta", Password: "secured_password"})

	// Read
	var account User
	storage.db.First(&account, 1)                       // find user with id 1
	storage.db.First(&account, "email = ?", "jarrieta") // find user with email jarrieta

	// Update - update user's Password to more_secured_password
	storage.db.Model(&account).Update("Password", "more_secured_password")

	// Delete - delete user
	storage.db.Delete(&account)
}

func (storage sqliteStorage) Create(user structures.User) (structures.User, error) {
	panic("implement me")
}

func (storage sqliteStorage) IsEmailAvailable(email string) (bool, error) {
	panic("implement me")
}

func (storage sqliteStorage) GetHashedPassword(email string) (string, error) {
	panic("implement me")
}

func (storage sqliteStorage) ChangePassword(email, newPassword string) error {
	panic("implement me")
}

func (storage sqliteStorage) SaveResetPasswordToken(token string) error {
	panic("implement me")
}

func (storage sqliteStorage) VerifyResetPasswordToken(token string) (bool, error) {
	panic("implement me")
}

func (storage sqliteStorage) Find(email string) (structures.User, error) {
	panic("implement me")
}

func (storage sqliteStorage) SaveResetToken(email, resetToken string) error {
	panic("implement me")
}

package service

import (
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
	"github.com/DATA-DOG/go-sqlmock"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var user_100 = &model.User{
	ID:          100,
	Google_ID:   "",
	Username:    "happyzzz",
	Full_name:   "Zoe Chen",
	Email:       "happyzzz@gmail.com",
	Password:    "jicdnwij8889",
	Description: "",
	Status:      "BASIC",
	Bean:        1000,
}

var user_101 = &model.User{
	ID:          101,
	Google_ID:   "ndjcibu156",
	Username:    "minican",
	Full_name:   "Cindy Chen",
	Email:       "minican@gmail.com",
	Password:    "",
	Description: "",
	Status:      "BASIC",
	Bean:        50,
}

type AnyPassword struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyPassword) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

func Test_FindUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	persistence.InitTestDB(gdb)

	mock.ExpectQuery(
		`SELECT * FROM "users" WHERE username = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`).
		WithArgs(user_12.Username).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "google_id", "username", "full_name", "email", "password", "description", "status", "bean"}).
				AddRow(user_12.ID, user_12.Google_ID, user_12.Username, user_12.Full_name, user_12.Email, user_12.Password, user_12.Description, user_12.Status, user_12.Bean))
	user, err := findUserByUsername("gary1030")

	require.NoError(t, err)
	require.Equal(t, user, user_12)
}

func Test_FindUserByGoogleId(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	persistence.InitTestDB(gdb)

	mock.ExpectQuery(
		`SELECT * FROM "users" WHERE google_id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`).
		WithArgs(user_101.Google_ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "google_id", "username", "full_name", "email", "password", "description", "status", "bean"}).
				AddRow(user_101.ID, user_101.Google_ID, user_101.Username, user_101.Full_name, user_101.Email, user_101.Password, user_101.Description, user_101.Status, user_101.Bean))
	user, err := findUserByGoogleId("ndjcibu156")

	require.NoError(t, err)
	require.Equal(t, user, user_101)
}

func Test_FindUserByAccountID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	persistence.InitTestDB(gdb)

	mock.ExpectQuery(
		`SELECT * FROM "users" WHERE ID = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`).
		WithArgs(user_100.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "google_id", "username", "full_name", "email", "password", "description", "status", "bean"}).
				AddRow(user_100.ID, user_100.Google_ID, user_100.Username, user_100.Full_name, user_100.Email, user_100.Password, user_100.Description, user_100.Status, user_100.Bean))
	user, err := findUserByAccountID(100)

	require.NoError(t, err)
	require.Equal(t, user, user_100)
}

func Test_AddUser_Case_1(t *testing.T) {
	// Add note without course
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = false

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: mocket.DriverName,
		DSN:        "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable",
	})) // Can be any connection string

	persistence.InitTestDB(gdb)
	commonReply := []map[string]interface{}{{"id": 100, "username": "happyzzz", "full_name": "Zoe Chen", "email": "happyzzz@gmail.com", "password": "jicdnwij8889"}}
	mocket.Catcher.NewMock().OneTime().WithQuery(`INSERT INTO "users"`).WithArgs().WithReply(commonReply)
	user, err := AddUser(user_100.Username, user_100.Full_name, user_100.Email, user_100.Password)

	require.NoError(t, err)
	require.Equal(t, user.ID, user_100.ID)
	require.Equal(t, user.Username, user_100.Username)
	require.Equal(t, user.Full_name, user_100.Full_name)
	require.Equal(t, user.Email, user_100.Email)

}

func Test_AddUser_Case_2(t *testing.T) {
	// Add note without course
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	persistence.InitTestDB(gdb)

	mock.ExpectBegin()
	mock.ExpectQuery(
		`INSERT INTO "users" ("created_at","updated_at","deleted_at","google_id","username","full_name","email","password","description","status","bean") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING "id","id"`).
		WithArgs(AnyTime{}, AnyTime{}, nil, "", user_100.Username, user_100.Full_name, user_100.Email, AnyPassword{}, "", "BASIC", 150).
		WillReturnError(errors.New("UsernameExist"))
	mock.ExpectRollback()

	_, err = AddUser(user_100.Username, user_100.Full_name, user_100.Email, user_100.Password)

	require.Error(t, err)
	require.Equal(t, err, errors.New("UsernameExist"))

}

func Test_AddGoogleUser(t *testing.T) {
	// Add note without course
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = false

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: mocket.DriverName,
		DSN:        "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable",
	})) // Can be any connection string

	persistence.InitTestDB(gdb)
	commonReply := []map[string]interface{}{{"id": 101, "username": "minican", "full_name": "Cindy Chen", "email": "minican@gmail.com", "google_id": "ndjcibu156"}}
	mocket.Catcher.NewMock().OneTime().WithQuery(`INSERT INTO "users"`).WithArgs().WithReply(commonReply)
	user, err := AddGoogleUser(user_101.Google_ID, user_101.Username, user_101.Full_name, user_101.Email)

	require.NoError(t, err)
	require.Equal(t, user.ID, user_101.ID)
	require.Equal(t, user.Username, user_101.Username)
	require.Equal(t, user.Full_name, user_101.Full_name)
	require.Equal(t, user.Email, user_101.Email)
	require.Equal(t, user.Google_ID, user_101.Google_ID)

}

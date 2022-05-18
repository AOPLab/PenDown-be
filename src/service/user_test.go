package service

import (
	"testing"
	_ "time"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var user_100 = &model.User{
	ID:          100,
	Google_ID:   "ndjcibu156",
	Username:    "happyzzz",
	Full_name:   "Zoe Chen",
	Email:       "happyzzz@gmail.com",
	Password:    "jicdnwij8889",
	Description: "",
	Status:      "BASIC",
	Bean:        1000,
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
		WithArgs(user_100.Google_ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "google_id", "username", "full_name", "email", "password", "description", "status", "bean"}).
				AddRow(user_100.ID, user_100.Google_ID, user_100.Username, user_100.Full_name, user_100.Email, user_100.Password, user_100.Description, user_100.Status, user_100.Bean))
	user, err := findUserByGoogleId("ndjcibu156")

	require.NoError(t, err)
	require.Equal(t, user, user_100)
}

// func Test_AddUser(t *testing.T) {
// 	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening database connection", err)
// 	}
// 	defer db.Close()
// 	gdb, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: db,
// 	}), &gorm.Config{})

// 	persistence.InitTestDB(gdb)

// }

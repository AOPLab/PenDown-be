package service

import (
	"errors"
	"testing"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var note_3 = &model.Note{
	ID:                  3,
	User_id:             12,
	Title:               "IM5028-Lecture 01 Overview of SPM2",
	Description:         "軟專第一堂課筆記2",
	Is_template:         false,
	Bean:                10000,
	View_cnt:            39,
	Course_id:           283,
	Pdf_filename:        "1652413412_KcB262.pdf",
	Preview_filename:    "1652413412_G8lgY2.jpg",
	Goodnotes_filename:  "",
	Notability_filename: "1652413416_9nIyu2.note",
	CreatedAt:           test_time,
}

var user_1 = &model.User{
	ID:          1,
	Google_ID:   "",
	Username:    "gary1030",
	Full_name:   "Gary Hu",
	Email:       "gary1030@gmail.com",
	Password:    "wdw54t7q",
	Description: "",
	Status:      "BASIC",
	Bean:        20000,
}

func Test_BuySale_Case_1(t *testing.T) {
	// User have not enough bean
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
	mock.ExpectQuery(
		`SELECT * FROM "notes" WHERE "notes"."id" = $1 AND "notes"."deleted_at" IS NULL AND "notes"."id" = $2 ORDER BY "notes"."id" LIMIT 1`).
		WithArgs(note_3.ID, note_3.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "course_id", "title", "description", "is_template", "bean", "view_cnt", "pdf_filename", "preview_filename", "goodnotes_filename", "notability_filename", "created_at"}).
				AddRow(note_3.ID, note_3.User_id, note_3.Course_id, note_3.Title, note_3.Description, note_3.Is_template, note_3.Bean, note_3.View_cnt, note_3.Pdf_filename, note_3.Preview_filename, note_3.Goodnotes_filename, note_3.Notability_filename, note_3.CreatedAt))

	_, err = BuyNote(100, 3)

	require.Error(t, err)
	require.Equal(t, err, errors.New("No enough beans!"))
}

func Test_BuySale_Case_2(t *testing.T) {
	// User have enough bean
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
		WithArgs(user_1.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "google_id", "username", "full_name", "email", "password", "description", "status", "bean"}).
				AddRow(user_1.ID, user_1.Google_ID, user_1.Username, user_1.Full_name, user_1.Email, user_1.Password, user_1.Description, user_1.Status, user_1.Bean))
	mock.ExpectQuery(
		`SELECT * FROM "notes" WHERE "notes"."id" = $1 AND "notes"."deleted_at" IS NULL AND "notes"."id" = $2 ORDER BY "notes"."id" LIMIT 1`).
		WithArgs(note_3.ID, note_3.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "course_id", "title", "description", "is_template", "bean", "view_cnt", "pdf_filename", "preview_filename", "goodnotes_filename", "notability_filename", "created_at"}).
				AddRow(note_3.ID, note_3.User_id, note_3.Course_id, note_3.Title, note_3.Description, note_3.Is_template, note_3.Bean, note_3.View_cnt, note_3.Pdf_filename, note_3.Preview_filename, note_3.Goodnotes_filename, note_3.Notability_filename, note_3.CreatedAt))
	mock.ExpectQuery(
		`SELECT * FROM "notes" WHERE ID = $1 AND "notes"."deleted_at" IS NULL ORDER BY "notes"."id" LIMIT 1`).
		WithArgs(note_3.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "course_id", "title", "description", "is_template", "bean", "view_cnt", "pdf_filename", "preview_filename", "goodnotes_filename", "notability_filename", "created_at"}).
				AddRow(note_3.ID, note_3.User_id, note_3.Course_id, note_3.Title, note_3.Description, note_3.Is_template, note_3.Bean, note_3.View_cnt, note_3.Pdf_filename, note_3.Preview_filename, note_3.Goodnotes_filename, note_3.Notability_filename, note_3.CreatedAt))
	mock.ExpectBegin()
	mock.ExpectQuery(
		`INSERT INTO "downloads" ("created_at","updated_at","deleted_at","user_id","note_id") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`).
		WithArgs(AnyTime{}, AnyTime{}, nil, user_1.ID, note_3.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec(
		`UPDATE "users" SET "bean"=Bean + $1,"updated_at"=$2 WHERE "users"."deleted_at" IS NULL AND "id" = $3`).
		WithArgs(float64(8000), AnyTime{}, user_12.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec(
		`UPDATE "users" SET "bean"=Bean - $1,"updated_at"=$2 WHERE "users"."deleted_at" IS NULL AND "id" = $3`).
		WithArgs(10000, AnyTime{}, user_1.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	note, err := BuyNote(1, 3)

	require.NoError(t, err)
	require.Equal(t, note_3, note)
}

package service

import (
	"database/sql/driver"
	"errors"
	"testing"
	"time"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
	"github.com/DATA-DOG/go-sqlmock"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var test_time, time_err = time.Parse(time.RFC3339, "2022-05-13 03:43:29.278922+00")

var course_283 = &model.Course{
	ID:                283,
	School_id:         1,
	Course_no:         "IM3007",
	Course_name:       "System Analysis and Design",
	View_cnt:          39,
	Note_cnt:          10,
	Last_updated_time: test_time,
}

var note_82 = &model.Note{
	ID:                  82,
	User_id:             12,
	Title:               "IM5028-Lecture 01 Overview of SPM",
	Description:         "軟專第一堂課筆記",
	Is_template:         false,
	Bean:                10,
	View_cnt:            39,
	Course_id:           283,
	Pdf_filename:        "1652413412_KcB26.pdf",
	Preview_filename:    "1652413412_G8lgY.jpg",
	Goodnotes_filename:  "",
	Notability_filename: "1652413416_9nIyu.note",
	CreatedAt:           test_time,
	Course:              *course_283,
}

var note_83 = &model.Note{
	ID:                  83,
	User_id:             12,
	Title:               "IM5028-Lecture 01 Overview of SPM2",
	Description:         "軟專第一堂課筆記2",
	Is_template:         false,
	Bean:                10,
	View_cnt:            39,
	Course_id:           283,
	Pdf_filename:        "1652413412_KcB262.pdf",
	Preview_filename:    "1652413412_G8lgY2.jpg",
	Goodnotes_filename:  "",
	Notability_filename: "1652413416_9nIyu2.note",
	CreatedAt:           test_time,
	User:                *user_12,
}

var user_12 = &model.User{
	ID:          12,
	Google_ID:   "wdwdw000",
	Username:    "gary1030",
	Full_name:   "Gary Hu",
	Email:       "gary1030@gmail.com",
	Password:    "wdw54t7q",
	Description: "",
	Status:      "BASIC",
	Bean:        2000,
}

var note_1 = &model.Note{
	ID:                  1,
	User_id:             10,
	Title:               "Wifi protocol",
	Description:         "802.11",
	Is_template:         false,
	Bean:                100,
	View_cnt:            0,
	Pdf_filename:        "",
	Preview_filename:    "",
	Goodnotes_filename:  "",
	Notability_filename: "",
	CreatedAt:           time.Now(),
}

var note_2 = &model.Note{
	ID:                  2,
	User_id:             10,
	Title:               "Wifi protocol",
	Description:         "802.11",
	Is_template:         false,
	Bean:                100,
	Course_id:           310,
	View_cnt:            0,
	Pdf_filename:        "",
	Preview_filename:    "",
	Goodnotes_filename:  "",
	Notability_filename: "",
	CreatedAt:           time.Now(),
}

var download_1 = &model.Download{
	User_id: 12,
	Note_id: 82,
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func Test_GetNoteByIdWithCourse(t *testing.T) {
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
		`SELECT * FROM "notes" WHERE "notes"."id" = $1 AND "notes"."deleted_at" IS NULL AND "notes"."id" = $2 ORDER BY "notes"."id" LIMIT 1`).
		WithArgs(note_82.ID, note_82.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "course_id", "title", "description", "is_template", "bean", "view_cnt", "pdf_filename", "preview_filename", "goodnotes_filename", "notability_filename", "created_at"}).
				AddRow(note_82.ID, note_82.User_id, note_82.Course_id, note_82.Title, note_82.Description, note_82.Is_template, note_82.Bean, note_82.View_cnt, note_82.Pdf_filename, note_82.Preview_filename, note_82.Goodnotes_filename, note_82.Notability_filename, note_82.CreatedAt))
	mock.ExpectQuery(
		`SELECT * FROM "courses" WHERE "courses"."id" = $1 AND "courses"."deleted_at" IS NULL`).
		WithArgs(note_82.Course_id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "school_id", "course_no", "course_name", "view_cnt", "note_cnt", "last_updated_time"}).
				AddRow(note_82.Course.ID, note_82.Course.School_id, note_82.Course.Course_no, note_82.Course.Course_name, note_82.Course.View_cnt, note_82.Course.Note_cnt, note_82.Course.Last_updated_time))

	// now we execute our method
	note, err := GetNoteByIdWithCourse(note_82.ID)

	require.NoError(t, err)
	require.Equal(t, note, note_82)
}

func Test_GetUserNoteById(t *testing.T) {
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
		`SELECT * FROM "notes" WHERE "notes"."id" = $1 AND "notes"."user_id" = $2 AND "notes"."deleted_at" IS NULL AND "notes"."id" = $3 ORDER BY "notes"."id" LIMIT 1`).
		WithArgs(note_82.ID, note_82.User_id, note_82.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "course_id", "title", "description", "is_template", "bean", "view_cnt", "pdf_filename", "preview_filename", "goodnotes_filename", "notability_filename", "created_at"}).
				AddRow(note_82.ID, note_82.User_id, note_82.Course_id, note_82.Title, note_82.Description, note_82.Is_template, note_82.Bean, note_82.View_cnt, note_82.Pdf_filename, note_82.Preview_filename, note_82.Goodnotes_filename, note_82.Notability_filename, note_82.CreatedAt))
	mock.ExpectQuery(
		`SELECT * FROM "courses" WHERE "courses"."id" = $1 AND "courses"."deleted_at" IS NULL`).
		WithArgs(note_82.Course_id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "school_id", "course_no", "course_name", "view_cnt", "note_cnt", "last_updated_time"}).
				AddRow(note_82.Course.ID, note_82.Course.School_id, note_82.Course.Course_no, note_82.Course.Course_name, note_82.Course.View_cnt, note_82.Course.Note_cnt, note_82.Course.Last_updated_time))

	// now we execute our method
	note, err := GetUserNoteById(note_82.User_id, note_82.ID)

	require.NoError(t, err)
	require.Equal(t, note, note_82)
}

func Test_AddNote_Case_1(t *testing.T) {
	// Add note without course
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = false

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: mocket.DriverName,
		DSN:        "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable",
	})) // Can be any connection string

	persistence.InitTestDB(gdb)

	// ("created_at","updated_at","deleted_at","user_id","title","description","view_cnt","is_template","bean")
	// VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id","course_id","pdf_filename","preview_filename","goodnotes_filename","notability_filename","id"
	commonReply := []map[string]interface{}{{"id": 1, "course_id": nil, "pdf_filename": "", "preview_filename": "", "goodnotes_filename": "", "notability_filename": ""}}
	mocket.Catcher.NewMock().OneTime().WithQuery(`INSERT INTO "notes"`).WithArgs().WithReply(commonReply)
	mocket.Catcher.NewMock().WithQuery(`UPDATE "users" SET "bean"=Bean + $1,"updated_at"=$2 WHERE "users"."deleted_at" IS NULL AND "id" = $3`)

	note, err := AddNote(note_1.User_id, note_1.Title, note_1.Description, note_1.Is_template, nil, note_1.Bean)

	require.NoError(t, err)
	require.Equal(t, note.ID, note_1.ID)
	require.Equal(t, note.User_id, note_1.User_id)
	require.Equal(t, note.Title, note_1.Title)
	require.Equal(t, note.Description, note_1.Description)
	require.Equal(t, note.Is_template, note_1.Is_template)
	require.Equal(t, note.Course_id, note_1.Course_id)
	require.Equal(t, note.Bean, note_1.Bean)
	require.Equal(t, note.Pdf_filename, note_1.Pdf_filename)
}

func Test_AddNote_Case_2(t *testing.T) {
	// Add note with course
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = false

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: mocket.DriverName,
		DSN:        "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable",
	})) // Can be any connection string

	persistence.InitTestDB(gdb)

	commonReply := []map[string]interface{}{{"id": 2, "course_id": note_2.Course_id, "pdf_filename": "", "preview_filename": "", "goodnotes_filename": "", "notability_filename": ""}}
	mocket.Catcher.NewMock().WithQuery(`UPDATE "courses" SET "note_cnt"=Note_cnt + $1,"updated_at"=$2 WHERE "courses"."deleted_at" IS NULL AND "id" = $3`)
	mocket.Catcher.NewMock().OneTime().WithQuery(`INSERT INTO "notes"`).WithArgs().WithReply(commonReply)
	mocket.Catcher.NewMock().WithQuery(`UPDATE "users" SET "bean"=Bean + $1,"updated_at"=$2 WHERE "users"."deleted_at" IS NULL AND "id" = $3`)

	note, err := AddNote(note_2.User_id, note_2.Title, note_2.Description, note_2.Is_template, &note_2.Course_id, note_2.Bean)

	require.NoError(t, err)
	require.Equal(t, note.ID, note_2.ID)
	require.Equal(t, note.User_id, note_2.User_id)
	require.Equal(t, note.Title, note_2.Title)
	require.Equal(t, note.Description, note_2.Description)
	require.Equal(t, note.Is_template, note_2.Is_template)
	require.Equal(t, note.Course_id, note_2.Course_id)
	require.Equal(t, note.Bean, note_2.Bean)
	require.Equal(t, note.Pdf_filename, note_2.Pdf_filename)
}

func Test_AddNote_Case_3(t *testing.T) {
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
		`INSERT INTO "notes" ("created_at","updated_at","deleted_at","user_id","title","description","view_cnt","is_template","bean") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id","course_id","pdf_filename","preview_filename","goodnotes_filename","notability_filename","id"`).
		WithArgs(AnyTime{}, AnyTime{}, nil, note_1.User_id, note_1.Title, note_1.Description, 0, note_1.Is_template, note_1.Bean).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "course_id", "pdf_filename", "preview_filename", "goodnotes_filename", "notability_filename", "id"}).
				AddRow(note_1.ID, nil, "", "", "", "", note_1.ID))
	mock.ExpectCommit()

	note, err := AddNote(note_1.User_id, note_1.Title, note_1.Description, note_1.Is_template, nil, note_1.Bean)

	require.NoError(t, err)
	require.Equal(t, note.ID, note_1.ID)
	require.Equal(t, note.User_id, note_1.User_id)
	require.Equal(t, note.Title, note_1.Title)
	require.Equal(t, note.Description, note_1.Description)
	require.Equal(t, note.Is_template, note_1.Is_template)
	require.Equal(t, note.Bean, note_1.Bean)
	require.Equal(t, note.Pdf_filename, note_1.Pdf_filename)
	require.Equal(t, note.Preview_filename, note_1.Preview_filename)
	require.Equal(t, note.Goodnotes_filename, note_1.Goodnotes_filename)
	require.Equal(t, note.Notability_filename, note_1.Notability_filename)
}

func Test_UpdatePdfFilename(t *testing.T) {
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = false

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: mocket.DriverName,
		DSN:        "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable",
	})) // Can be any connection string

	persistence.InitTestDB(gdb)

	mocket.Catcher.NewMock().WithQuery(`UPDATE "notes" SET "updated_at"=$1,"pdf_filename"=$2,"preview_filename"=$3 WHERE "notes"."deleted_at" IS NULL AND "id" = $4`)

	// now we execute our method
	err = UpdatePdfFilename(82, "apple", "banana")

	require.NoError(t, err)
}

func Test_SearchNoteAll(t *testing.T) {
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
		`SELECT notes.ID, notes.user_id, users.username, notes.title, notes.view_cnt, notes.preview_filename, notes.notability_filename, notes.goodnotes_filename, notes.created_at FROM "notes" JOIN users on notes.user_id = users.id WHERE lower(notes.title) LIKE $1 LIMIT 2`).
		WithArgs("%5028%").
		WillReturnRows(
			sqlmock.NewRows([]string{"ID", "user_id", "username", "title", "view_cnt", "preview_filename", "notability_filename", "goodnotes_filename", "created_at"}).
				AddRow(note_83.ID, note_83.User_id, note_83.User.Username, note_83.Title, note_83.View_cnt, note_83.Preview_filename, note_83.Notability_filename, note_83.Goodnotes_filename, note_83.CreatedAt))

	mock.ExpectQuery(
		`SELECT count(*) FROM "notes" JOIN users on notes.user_id = users.id WHERE lower(notes.title) LIKE $1`).
		WithArgs("%5028%").
		WillReturnRows(
			sqlmock.NewRows([]string{"count"}).
				AddRow(1))

	// now we execute our method
	note, note_cnt, err := SearchNote("5028", 0, 2, "all")
	var results []SearchNoteOutput
	result := SearchNoteOutput{
		ID:                  note_83.ID,
		User_id:             note_83.User_id,
		Username:            note_83.User.Username,
		Title:               note_83.Title,
		Preview_filename:    note_83.Preview_filename,
		View_cnt:            note_83.View_cnt,
		Goodnotes_filename:  note_83.Goodnotes_filename,
		Notability_filename: note_83.Notability_filename,
		CreatedAt:           note_83.CreatedAt,
	}
	results = append(results, result)

	require.NoError(t, err)
	require.Equal(t, note_cnt, int64(1))
	require.Equal(t, note, results)
}

func Test_CheckUserBuyNote_Case_1(t *testing.T) {
	// Can get buying record
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
		`SELECT * FROM "downloads" WHERE "downloads"."user_id" = $1 AND "downloads"."note_id" = $2 AND "downloads"."deleted_at" IS NULL ORDER BY "downloads"."id" LIMIT 1`).
		WithArgs(download_1.User_id, download_1.Note_id).
		WillReturnRows(
			sqlmock.NewRows([]string{"user_id", "download_id"}).
				AddRow(download_1.User_id, download_1.Note_id))

	// now we execute our method
	isBuy := CheckUserBuyNote(download_1.User_id, download_1.Note_id)

	require.Equal(t, isBuy, true)
}

func Test_CheckUserBuyNote_Case_2(t *testing.T) {
	// Cannot get buying record
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
		`SELECT * FROM "downloads" WHERE "downloads"."user_id" = $1 AND "downloads"."note_id" = $2 AND "downloads"."deleted_at" IS NULL ORDER BY "downloads"."id" LIMIT 1`).
		WithArgs(12, 80).
		WillReturnError(errors.New("NotExist"))

	// now we execute our method
	isBuy := CheckUserBuyNote(12, 80)

	require.Equal(t, isBuy, false)
}

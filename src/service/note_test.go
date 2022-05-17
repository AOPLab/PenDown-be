package service

import (
	"testing"
	"time"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var note_82_time, time_err = time.Parse(time.RFC3339, "2022-05-13 03:43:29.278922+00")

var note_82 = &model.Note{
	ID:                  82,
	User_id:             12,
	Title:               "IM5028-Lecture 01 Overview of SPM",
	Description:         "軟專第一堂課筆記",
	Is_template:         false,
	Bean:                10,
	View_cnt:            39,
	Pdf_filename:        "1652413412_KcB26.pdf",
	Preview_filename:    "1652413412_G8lgY.jpg",
	Goodnotes_filename:  "",
	Notability_filename: "1652413416_9nIyu.note",
	CreatedAt:           note_82_time,
}

func Test_GetNoteWithCourse(t *testing.T) {
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
			sqlmock.NewRows([]string{"id", "user_id", "title", "description", "is_template", "bean", "view_cnt", "pdf_filename", "preview_filename", "goodnotes_filename", "notability_filename", "created_at"}).
				AddRow(note_82.ID, note_82.User_id, note_82.Title, note_82.Description, note_82.Is_template, note_82.Bean, note_82.View_cnt, note_82.Pdf_filename, note_82.Preview_filename, note_82.Goodnotes_filename, note_82.Notability_filename, note_82.CreatedAt))

	// now we execute our method
	note, err := GetNoteByIdWithCourse(82)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	require.NoError(t, err)
	require.Equal(t, note, note_82)
}

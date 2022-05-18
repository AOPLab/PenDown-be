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
			sqlmock.NewRows([]string{"id", "user_id", "course_id", "title", "description", "is_template", "bean", "view_cnt", "pdf_filename", "preview_filename", "goodnotes_filename", "notability_filename", "created_at"}).
				AddRow(note_82.ID, note_82.User_id, note_82.Course_id, note_82.Title, note_82.Description, note_82.Is_template, note_82.Bean, note_82.View_cnt, note_82.Pdf_filename, note_82.Preview_filename, note_82.Goodnotes_filename, note_82.Notability_filename, note_82.CreatedAt))
	mock.ExpectQuery(
		`SELECT * FROM "courses" WHERE "courses"."id" = $1 AND "courses"."deleted_at" IS NULL`).
		WithArgs(note_82.Course_id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "school_id", "course_no", "course_name", "view_cnt", "note_cnt", "last_updated_time"}).
				AddRow(note_82.Course.ID, note_82.Course.School_id, note_82.Course.Course_no, note_82.Course.Course_name, note_82.Course.View_cnt, note_82.Course.Note_cnt, note_82.Course.Last_updated_time))

	// now we execute our method
	note, err := GetNoteByIdWithCourse(82)

	require.NoError(t, err)
	require.Equal(t, note, note_82)
}

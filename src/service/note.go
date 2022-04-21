package service

import (
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

func addNoteTag(note_id int64, tag_id int64) error {
	note_tag := &model.NoteTag{
		Note_id: note_id,
		Tag_id:  tag_id,
	}
	db_err := persistence.DB.Model(&model.NoteTag{}).Create(&note_tag).Error
	if db_err != nil {
		return db_err
	}
	return nil
}

func AddPost(title string, description string, is_template bool, course_id int64, tags []int64) (*model.Note, error) {

	note := &model.Note{
		Title:       title,
		Description: description,
		Is_template: is_template,
		Course_id:   course_id,
	}

	db_err := persistence.DB.Model(&model.Note{}).Create(&note).Error
	if db_err != nil {
		return nil, db_err
	}

	// TODO: add NoteTag

	return note, nil
}

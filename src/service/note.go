package service

import (
	"errors"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

func GetNoteTag(note_id int64) (*[]model.NoteTag, error) {
	var noteTags []model.NoteTag
	db_err := persistence.DB.Preload("Tag").Where(map[string]interface{}{"note_id": note_id}).Find(&noteTags).Error
	if db_err != nil {
		return nil, db_err
	}
	return &noteTags, nil
}

func AddNoteTag(user_id int64, note_id int64, tag_id int64) error {
	_, note_err := GetUserNoteById(user_id, note_id)
	if note_err != nil {
		return errors.New("NoteNotExist")
	}

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

func DeleteNoteTag(user_id int64, note_id int64, tag_id int64) error {
	_, note_err := GetUserNoteById(user_id, note_id)
	if note_err != nil {
		return errors.New("NoteNotExist")
	}

	db_err := persistence.DB.Unscoped().Where("note_id = ? AND tag_id = ?", note_id, tag_id).Delete(&model.NoteTag{}).Error
	if db_err != nil {
		return db_err
	}
	return nil
}

func AddNote(user_id int64, title string, description string, is_template bool, course_id int64, bean int) (*model.Note, error) {

	note := &model.Note{
		User_id:     user_id,
		Title:       title,
		Description: description,
		Is_template: is_template,
		Course_id:   course_id,
		Bean:        bean,
	}

	db_err := persistence.DB.Model(&model.Note{}).Create(&note).Error
	if db_err != nil {
		return nil, db_err
	}

	return note, nil
}

func GetNoteByIdWithCourse(note_id int64) (*model.Note, error) {
	note := &model.Note{
		ID: note_id,
	}
	db_err := persistence.DB.Preload("Course").Where(&note).First(&note).Error
	if db_err != nil {
		return nil, db_err
	}

	return note, nil
}

func GetUserNoteById(user_id int64, note_id int64) (*model.Note, error) {
	note := &model.Note{
		ID:      note_id,
		User_id: user_id,
	}
	db_err := persistence.DB.Preload("Course").Where(&note).First(&note).Error
	if db_err != nil {
		return nil, db_err
	}

	return note, nil
}

func GetNoteById(note_id int64) (*model.Note, error) {
	note := &model.Note{
		ID: note_id,
	}
	db_err := persistence.DB.Preload("User").Preload("Course").Preload("Course.School").Where(&note).First(&note).Error
	if db_err != nil {
		return nil, db_err
	}

	return note, nil
}

func UpdateNotabilityFilename(note_id int64, filename string) error {
	note := &model.Note{ID: note_id}
	db_err := persistence.DB.Model(&note).Update("notability_filename", filename).Error
	if db_err != nil {
		return db_err
	}

	return nil
}

func UpdateGoodnoteFilename(note_id int64, filename string) error {
	note := &model.Note{ID: note_id}
	db_err := persistence.DB.Model(&note).Update("goodnotes_filename", filename).Error
	if db_err != nil {
		return db_err
	}

	return nil
}

func UpdatePdfFilename(note_id int64, pdf_filename string) error {
	note := &model.Note{ID: note_id}
	db_err := persistence.DB.Model(&note).Updates(model.Note{Pdf_filename: pdf_filename}).Error
	if db_err != nil {
		return db_err
	}

	return nil
}

func UpdatePreviewFilename(note_id int64, preview_filename string) error {
	note := &model.Note{ID: note_id}
	db_err := persistence.DB.Model(&note).Updates(model.Note{Preview_filename: preview_filename}).Error
	if db_err != nil {
		return db_err
	}

	return nil
}

func CheckUserBuyNote(user_id int64, note_id int64) bool {
	download := &model.Download{
		User_id: user_id,
		Note_id: note_id,
	}
	db_err := persistence.DB.Where(&download).First(&download).Error
	if db_err != nil {
		return false
	}
	return true
}

func GetNoteSavedCnt(note_id int64) (int64, error) {
	var saved_cnt int64
	err := persistence.DB.Model(&model.Saved{}).Where("note_id = ?", note_id).Count(&saved_cnt).Error
	if err != nil {
		return 0, err
	}
	return saved_cnt, nil
}

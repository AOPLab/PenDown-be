package service

import (
	"errors"
	"math"

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

func GetNote(filter string, offset int64) ([]model.Note, int64, error) {
	var count int64
	var size = 6
	persistence.DB.Model(&model.Note{}).Count(&count)
	total_cnt := int64(math.Ceil(float64(count) / float64(size)))
	if offset >= total_cnt {
		return nil, 0, errors.New("offset out of range")
	}

	var notes []model.Note
	switch filter {
	case "popular":
		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset) * size).Preload("User").Preload("Course").Find(&notes)
		if results.Error != nil {
			return nil, total_cnt, results.Error
		}
		return notes, total_cnt, nil
	case "recent":
		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset) * size).Preload("User").Preload("Course").Find(&notes)
		if results.Error != nil {
			return nil, total_cnt, results.Error
		}
		return notes, total_cnt, nil
	default:
		return nil, total_cnt, nil
	}
}

type NoteOutput struct {
	Tag_id              string `json:"tag_id"`
	ID                  string `json:"note_id"`
	Title               string `json:"title"`
	Notability_filename string `json:"notability_filename"`
}

// (tag_id int64, note_type string, filter string, offset int64)
func GetNoteByTag(tag_id int64) (NoteOutput, error) {
	// Join NoteTag and Note
	// var notes []model.NoteTag
	// results := persistence.DB.Joins("Note").Where("Tag_id = ?", tag_id).Where("Note.notability_filename IS NOT NULL").Find(&notes)
	// if results.Error != nil {
	// 	return nil, 0, results.Error
	// } else {
	// 	return notes, results.RowsAffected, results.Error
	// }

	var results NoteOutput
	if err := persistence.DB.Table("notetag").Select("notetag.Tag_id, note.id, note.title, note.Notability_filename").Joins("JOIN note on note.id = notetag.Note_id").Find(&results).Error; err != nil {
		return results, err
	}
	return results, nil

	// var count int64
	// var size = 6
	// switch note_type {
	// case "notability":
	// 	persistence.DB.Model(&model.Note{}).Where("notability_filename IS NOT NULL").Count(&count)
	// case "goodnotes":
	// 	persistence.DB.Model(&model.Note{}).Where("goodnotes_filename IS NOT NULL").Count(&count)
	// default:
	// 	persistence.DB.Model(&model.Note{}).Count(&count)
	// }
	// total_cnt := int64(math.Ceil(float64(count) / float64(size)))
	// if offset >= total_cnt {
	// 	return nil, 0, errors.New("offset out of range")
	// }

	// var notes []model.Note
	// switch filter {
	// case "popular":
	// 	results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset) * size).Preload("User").Preload("Course").Find(&notes)
	// 	if results.Error != nil {
	// 		return nil, total_cnt, results.Error
	// 	}
	// 	return notes, total_cnt, nil
	// case "recent":
	// 	results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset) * size).Preload("User").Preload("Course").Find(&notes)
	// 	if results.Error != nil {
	// 		return nil, total_cnt, results.Error
	// 	}
	// 	return notes, total_cnt, nil
	// default:
	// 	return nil, total_cnt, nil
	// }
}

func GetNoteByUser(user_id int64, filter string, offset int64) ([]model.Note, int64, error) {
	var count int64
	var size = 6
	persistence.DB.Model(&model.Note{}).Where("user_id = ?", user_id).Count(&count)
	total_cnt := int64(math.Ceil(float64(count) / float64(size)))
	if offset >= total_cnt {
		return nil, 0, errors.New("offset out of range")
	}

	var notes []model.Note
	switch filter {
	case "popular":
		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)*size).Preload("User").Preload("Course").Where("user_id = ?", user_id).Find(&notes)
		if results.Error != nil {
			return nil, total_cnt, results.Error
		}
		return notes, total_cnt, nil
	case "recent":
		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)*size).Preload("User").Preload("Course").Where("user_id = ?", user_id).Find(&notes)
		if results.Error != nil {
			return nil, total_cnt, results.Error
		}
		return notes, total_cnt, nil
	default:
		return nil, total_cnt, nil
	}
}

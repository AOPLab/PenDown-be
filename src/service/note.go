package service

import (
	"errors"
	"time"

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

func UpdateNoteViewCnt(note_id int64, view_cnt int64) error {
	note := &model.Note{ID: note_id}
	db_err := persistence.DB.Model(&note).Update("view_cnt", view_cnt).Error
	if db_err != nil {
		return db_err
	}

	return nil
}

func GetNotes(filter string, offset int64) ([]model.Note, int64, error) {
	var count int64
	var size = 6

	var notes []model.Note
	switch filter {
	case "hot":
		date := time.Now()
		lastMonth := date.AddDate(0, -1, 0)
		persistence.DB.Model(&model.Note{}).Where("created_at > ?", lastMonth).Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Preload("Course").Where("created_at > ?", lastMonth).Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "popular":
		persistence.DB.Model(&model.Note{}).Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Preload("Course").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "recent":
		persistence.DB.Model(&model.Note{}).Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Preload("Course").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	default:
		return nil, 0, nil
	}
}

type NoteOutput struct {
	Note_ID             int64     `json:"note_id"`
	ID                  int64     `json:"user_id"`
	Username            string    `json:"username"`
	Title               string    `json:"title"`
	Preview_filename    string    `json:"preview_filename"`
	Goodnotes_filename  string    `json:"goodnotes_filename"`
	Notability_filename string    `json:"notability_filename"`
	View_cnt            int64     `json:"view_cnt"`
	CreatedAt           time.Time `json:"created_at"`
}

// string type: all-popular, notability-popular, goodnotes-popular, all-recent, notability-recent, goodnotes-recent
func GetNoteByTag(tag_id int64, filter string, offset int64) ([]NoteOutput, int64, error) {
	// Join NoteTag and Note
	size := 9
	var results []NoteOutput
	var count int64

	filter_col := "note_tags.note_id, users.ID, users.username, notes.title, notes.preview_filename, notes.notability_filename, notes.goodnotes_filename, notes.created_at, notes.view_cnt"

	switch filter {
	case "all-recent":
		persistence.DB.Model(&model.NoteTag{}).Where("note_tags.tag_id = ?", tag_id).Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		if err := persistence.DB.Order("notes.created_at desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	case "all-popular":
		persistence.DB.Model(&model.NoteTag{}).Where("note_tags.tag_id = ?", tag_id).Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		if err := persistence.DB.Order("notes.view_cnt desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	case "notability-recent":
		persistence.DB.Table("note_tags").Joins("JOIN notes on notes.id = note_tags.Note_id").Where("note_tags.tag_id = ?", tag_id).Where("notability_filename IS NOT NULL").Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		if err := persistence.DB.Order("notes.created_at desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Where("notability_filename IS NOT NULL").Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	case "notability-popular":
		persistence.DB.Table("note_tags").Joins("JOIN notes on notes.id = note_tags.Note_id").Where("note_tags.tag_id = ?", tag_id).Where("notability_filename IS NOT NULL").Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		if err := persistence.DB.Order("notes.view_cnt desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Where("notability_filename IS NOT NULL").Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	case "goodnotes-recent":
		persistence.DB.Table("note_tags").Joins("JOIN notes on notes.id = note_tags.Note_id").Where("note_tags.tag_id = ?", tag_id).Where("goodnotes_filename IS NOT NULL").Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		if err := persistence.DB.Order("notes.created_at desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Where("goodnotes_filename IS NOT NULL").Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	case "goodnotes-popular":
		persistence.DB.Table("note_tags").Joins("JOIN notes on notes.id = note_tags.Note_id").Where("note_tags.tag_id = ?", tag_id).Where("goodnotes_filename IS NOT NULL").Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		if err := persistence.DB.Order("notes.view_cnt desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Where("goodnotes_filename IS NOT NULL").Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	default:
		return nil, 0, nil
	}
}

// string type: all-popular, notability-popular, goodnotes-popular, all-recent, notability-recent, goodnotes-recent
func GetNoteByCourse(course_id int64, filter string, offset int64) ([]model.Note, int64, error) {
	// Join NoteTag and Note
	size := 6
	var notes []model.Note
	var count int64

	switch filter {
	case "all-recent":
		persistence.DB.Model(&model.Note{}).Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "all-popular":
		persistence.DB.Model(&model.Note{}).Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "notability-recent":
		persistence.DB.Model(&model.Note{}).Where("notability_filename IS NOT NULL").Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Where("notability_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "notability-popular":
		persistence.DB.Model(&model.Note{}).Where("notability_filename IS NOT NULL").Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Where("notability_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "goodnotes-recent":
		persistence.DB.Model(&model.Note{}).Where("goodnotes_filename IS NOT NULL").Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Where("goodnotes_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "goodnotes-popular":
		persistence.DB.Model(&model.Note{}).Where("goodnotes_filename IS NOT NULL").Count(&count)
		if offset >= count {
			return nil, 0, errors.New("offset out of range")
		}
		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Where("goodnotes_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	default:
		return nil, 0, nil
	}
}

func GetNoteByUser(user_id int64, filter string, offset int64) ([]model.Note, int64, error) {
	var count int64
	var size = 9
	persistence.DB.Model(&model.Note{}).Where("user_id = ?", user_id).Count(&count)
	total_cnt := count
	if offset >= total_cnt {
		return nil, 0, errors.New("offset out of range")
	}

	var notes []model.Note
	switch filter {
	case "popular":
		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Preload("Course").Where("user_id = ?", user_id).Find(&notes)
		if results.Error != nil {
			return nil, total_cnt, results.Error
		}
		return notes, total_cnt, nil
	case "recent":
		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Preload("Course").Where("user_id = ?", user_id).Find(&notes)
		if results.Error != nil {
			return nil, total_cnt, results.Error
		}
		return notes, total_cnt, nil
	default:
		return nil, total_cnt, nil
	}
}

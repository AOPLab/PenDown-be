package service

import (
	"errors"
	"strings"
	"time"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
	"gorm.io/gorm"
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

func AddNote(user_id int64, title string, description string, is_template bool, course_id *int64, bean int) (*model.Note, error) {
	note := &model.Note{
		User_id:     user_id,
		Title:       title,
		Description: description,
		Is_template: is_template,
		Bean:        bean,
	}

	if course_id != nil {
		note.Course_id = *course_id
		courseUpdate := &model.Course{ID: *course_id}
		err := persistence.DB.Model(&courseUpdate).Update("Note_cnt", gorm.Expr("Note_cnt + ?", 1)).Update("Last_updated_time", time.Now()).Error
		if err != nil {
			return nil, err
		}
	}

	db_err := persistence.DB.Model(&model.Note{}).Create(&note).Error
	if db_err != nil {
		return nil, db_err
	}

	addUserBean(user_id)
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

func UpdateGoodnotesFilename(note_id int64, filename string) error {
	note := &model.Note{ID: note_id}
	db_err := persistence.DB.Model(&note).Update("goodnotes_filename", filename).Error
	if db_err != nil {
		return db_err
	}

	return nil
}

func UpdatePdfFilename(note_id int64, pdf_filename string, preview_filename string) error {
	note := &model.Note{ID: note_id}
	db_err := persistence.DB.Model(&note).Updates(model.Note{Pdf_filename: pdf_filename, Preview_filename: preview_filename}).Error
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

type SearchNoteOutput struct {
	ID                  int64     `json:"note_id"`
	User_id             int64     `json:"user_id"`
	Username            string    `json:"username"`
	Title               string    `json:"title"`
	View_cnt            int64     `json:"view_cnt"`
	Preview_filename    string    `json:"preview_filename"`
	Goodnotes_filename  string    `json:"goodnotes_filename"`
	Notability_filename string    `json:"notability_filename"`
	CreatedAt           time.Time `json:"created_at"`
}

// LIKE: title
// note_type: all, notability, goodnotes
func SearchNote(q string, offset int, limit int, note_type string) ([]SearchNoteOutput, int64, error) {
	var results []SearchNoteOutput
	var count int64
	count = 0
	searchName := "%" + strings.ToLower(q) + "%"
	selectField := "notes.ID, notes.user_id, users.username, notes.title, notes.view_cnt, notes.preview_filename, notes.notability_filename, notes.goodnotes_filename, notes.created_at"
	switch note_type {
	case "all":
		if err := persistence.DB.Limit(limit).Offset(offset).Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Find(&results).Error; err != nil {
			return results, 0, err
		}
		persistence.DB.Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Count(&count)
	case "notability":
		if err := persistence.DB.Limit(limit).Offset(offset).Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Where("notability_filename IS NOT NULL").Find(&results).Error; err != nil {
			return results, 0, err
		}
		persistence.DB.Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Where("notability_filename IS NOT NULL").Count(&count)
	case "goodnotes":
		if err := persistence.DB.Limit(limit).Offset(offset).Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Where("goodnotes_filename IS NOT NULL").Find(&results).Error; err != nil {
			return results, 0, err
		}
		persistence.DB.Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Where("goodnotes_filename IS NOT NULL").Count(&count)
	default:
		break
	}

	return results, count, nil
}

// LIKE: title
// note_type: all, notability, goodnotes
func SearchTemplate(q string, offset int, limit int, note_type string) ([]SearchNoteOutput, int64, error) {
	var results []SearchNoteOutput
	var count int64
	count = 0
	searchName := "%" + strings.ToLower(q) + "%"
	selectField := "notes.ID, notes.user_id, users.username, notes.title, notes.view_cnt, notes.preview_filename, notes.notability_filename, notes.goodnotes_filename, notes.created_at"
	switch note_type {
	case "all":
		if err := persistence.DB.Limit(limit).Offset(offset).Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Where("notes.is_template = ?", true).Find(&results).Error; err != nil {
			return results, 0, err
		}
		persistence.DB.Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Where("notes.is_template = ?", true).Count(&count)
	case "notability":
		if err := persistence.DB.Limit(limit).Offset(offset).Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Where("notability_filename IS NOT NULL").Where("notes.is_template = ?", true).Find(&results).Error; err != nil {
			return results, 0, err
		}
		persistence.DB.Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Where("notability_filename IS NOT NULL").Where("notes.is_template = ?", true).Count(&count)
	case "goodnotes":
		if err := persistence.DB.Limit(limit).Offset(offset).Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Where("goodnotes_filename IS NOT NULL").Where("notes.is_template = ?", true).Find(&results).Error; err != nil {
			return results, 0, err
		}
		persistence.DB.Table("notes").Select(selectField).Joins("JOIN users on notes.user_id = users.id").Where("lower(notes.title) LIKE ?", searchName).Where("goodnotes_filename IS NOT NULL").Where("notes.is_template = ?", true).Count(&count)
	default:
		break
	}
	return results, count, nil
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

		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Preload("Course").Where("created_at > ?", lastMonth).Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "popular":
		persistence.DB.Model(&model.Note{}).Count(&count)

		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Preload("Course").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "recent":
		persistence.DB.Model(&model.Note{}).Count(&count)

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
func GetNotesByTag(tag_id int64, filter string, offset int64) ([]NoteOutput, int64, error) {
	// Join NoteTag and Note
	size := 12
	var results []NoteOutput
	var count int64

	filter_col := "note_tags.note_id, users.ID, users.username, notes.title, notes.preview_filename, notes.notability_filename, notes.goodnotes_filename, notes.created_at, notes.view_cnt"

	switch filter {
	case "all-recent":
		persistence.DB.Model(&model.NoteTag{}).Where("note_tags.tag_id = ?", tag_id).Count(&count)

		if err := persistence.DB.Order("notes.created_at desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	case "all-popular":
		persistence.DB.Model(&model.NoteTag{}).Where("note_tags.tag_id = ?", tag_id).Count(&count)

		if err := persistence.DB.Order("notes.view_cnt desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	case "notability-recent":
		persistence.DB.Table("note_tags").Joins("JOIN notes on notes.id = note_tags.Note_id").Where("note_tags.tag_id = ?", tag_id).Where("notability_filename IS NOT NULL").Count(&count)

		if err := persistence.DB.Order("notes.created_at desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Where("notability_filename IS NOT NULL").Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	case "notability-popular":
		persistence.DB.Table("note_tags").Joins("JOIN notes on notes.id = note_tags.Note_id").Where("note_tags.tag_id = ?", tag_id).Where("notability_filename IS NOT NULL").Count(&count)

		if err := persistence.DB.Order("notes.view_cnt desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Where("notability_filename IS NOT NULL").Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	case "goodnotes-recent":
		persistence.DB.Table("note_tags").Joins("JOIN notes on notes.id = note_tags.Note_id").Where("note_tags.tag_id = ?", tag_id).Where("goodnotes_filename IS NOT NULL").Count(&count)

		if err := persistence.DB.Order("notes.created_at desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Where("goodnotes_filename IS NOT NULL").Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	case "goodnotes-popular":
		persistence.DB.Table("note_tags").Joins("JOIN notes on notes.id = note_tags.Note_id").Where("note_tags.tag_id = ?", tag_id).Where("goodnotes_filename IS NOT NULL").Count(&count)

		if err := persistence.DB.Order("notes.view_cnt desc").Limit(size).Offset(int(offset)).Table("note_tags").Select(filter_col).Joins("JOIN notes on notes.id = note_tags.Note_id").Joins("JOIN users on users.id = notes.User_id").Where("note_tags.tag_id = ?", tag_id).Where("goodnotes_filename IS NOT NULL").Find(&results).Error; err != nil {
			return results, 0, err
		}
		return results, count, nil
	default:
		return nil, 0, errors.New("NoFilter")
	}
}

// string type: all-popular, notability-popular, goodnotes-popular, all-recent, notability-recent, goodnotes-recent
func GetNotesByCourse(course_id int64, filter string, offset int64) ([]model.Note, int64, error) {
	// Join NoteTag and Note
	size := 12
	var notes []model.Note
	var count int64

	switch filter {
	case "all-recent":
		persistence.DB.Model(&model.Note{}).Where("Course_id = ?", course_id).Count(&count)

		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "all-popular":
		persistence.DB.Model(&model.Note{}).Where("Course_id = ?", course_id).Count(&count)

		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "notability-recent":
		persistence.DB.Model(&model.Note{}).Where("Course_id = ?", course_id).Where("notability_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Where("notability_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "notability-popular":
		persistence.DB.Model(&model.Note{}).Where("Course_id = ?", course_id).Where("notability_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Where("notability_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "goodnotes-recent":
		persistence.DB.Model(&model.Note{}).Where("Course_id = ?", course_id).Where("goodnotes_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Where("goodnotes_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "goodnotes-popular":
		persistence.DB.Model(&model.Note{}).Where("Course_id = ?", course_id).Where("goodnotes_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Where("Course_id = ?", course_id).Where("goodnotes_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	default:
		return nil, 0, errors.New("NoFilter")
	}
}

// string type: all-popular, notability-popular, goodnotes-popular, all-recent, notability-recent, goodnotes-recent
func GetNotesByUserId(user_id int64, filter string, offset int64) ([]model.Note, int64, error) {
	size := 12
	var notes []model.Note
	var count int64

	switch filter {
	case "all-recent":
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", user_id).Count(&count)

		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("User_id = ?", user_id).Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "all-popular":
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", user_id).Count(&count)

		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Where("User_id = ?", user_id).Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "notability-recent":
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", user_id).Where("notability_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("User_id = ?", user_id).Where("notability_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "notability-popular":
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", user_id).Where("notability_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Where("User_id = ?", user_id).Where("notability_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "goodnotes-recent":
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", user_id).Where("goodnotes_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("User_id = ?", user_id).Where("goodnotes_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "goodnotes-popular":
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", user_id).Where("goodnotes_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Order("view_cnt desc").Limit(size).Offset(int(offset)).Preload("User").Where("User_id = ?", user_id).Where("goodnotes_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	default:
		return nil, 0, errors.New("NoFilter")
	}
}

// uploaded/saved/library
// string type: all, notability, goodnotes
func GetOwnUploadedNotes(user_id int64, filter string, offset int64) ([]model.Note, int64, error) {
	size := 12
	var notes []model.Note
	var count int64

	switch filter {
	case "all":
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", user_id).Count(&count)

		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("User_id = ?", user_id).Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "notability":
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", user_id).Where("notability_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("User_id = ?", user_id).Where("notability_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "goodnotes":
		persistence.DB.Model(&model.Note{}).Where("User_id = ?", user_id).Where("goodnotes_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Where("User_id = ?", user_id).Where("goodnotes_filename IS NOT NULL").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	default:
		return nil, 0, errors.New("NoFilter")
	}
}

// uploaded/saved/library
// string type: all, notability, goodnotes
func GetOwnSavedNotes(user_id int64, filter string, offset int64) ([]model.Saved, int64, error) {
	size := 12
	var notes []model.Saved
	var count int64
	switch filter {
	case "all":
		persistence.DB.Model(&model.Saved{}).Where("User_id = ?", user_id).Count(&count)

		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Preload("Note").Where("User_id = ?", user_id).Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "notability":
		persistence.DB.Table("saveds").Joins("JOIN notes on notes.id = saveds.Note_id").Where("saveds.User_id = ?", user_id).Where("notability_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Table("saveds").Order("created_at desc").Limit(size).Offset(int(offset)).Joins("JOIN notes on notes.id = saveds.Note_id").Where("saveds.User_id = ?", user_id).Where("notes.notability_filename IS NOT NULL").Preload("User").Preload("Note").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "goodnotes":
		persistence.DB.Table("saveds").Joins("JOIN notes on notes.id = saveds.Note_id").Where("saveds.User_id = ?", user_id).Where("goodnotes_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Table("saveds").Order("created_at desc").Limit(size).Offset(int(offset)).Joins("JOIN notes on notes.id = saveds.Note_id").Where("saveds.User_id = ?", user_id).Where("notes.goodnotes_filename IS NOT NULL").Preload("User").Preload("Note").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	default:
		return nil, 0, errors.New("NoFilter")
	}
}

// uploaded/saved/library
// string type: all, notability, goodnotes
func GetOwnLibraryNotes(user_id int64, filter string, offset int64) ([]model.Download, int64, error) {
	size := 12
	var notes []model.Download
	var count int64

	switch filter {
	case "all":
		persistence.DB.Model(&model.Download{}).Where("User_id = ?", user_id).Count(&count)

		results := persistence.DB.Order("created_at desc").Limit(size).Offset(int(offset)).Preload("User").Preload("Note").Where("User_id = ?", user_id).Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "notability":
		persistence.DB.Table("downloads").Joins("JOIN notes on notes.id = downloads.Note_id").Where("downloads.User_id = ?", user_id).Where("goodnotes_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Table("downloads").Order("created_at desc").Limit(size).Offset(int(offset)).Joins("JOIN notes on notes.id = downloads.Note_id").Where("downloads.User_id = ?", user_id).Where("notes.notability_filename IS NOT NULL").Preload("User").Preload("Note").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	case "goodnotes":
		persistence.DB.Model(&model.Download{}).Where("User_id = ?", user_id).Where("goodnotes_filename IS NOT NULL").Count(&count)

		results := persistence.DB.Table("downloads").Order("created_at desc").Limit(size).Offset(int(offset)).Joins("JOIN notes on notes.id = downloads.Note_id").Where("downloads.User_id = ?", user_id).Where("notes.goodnotes_filename IS NOT NULL").Preload("User").Preload("Note").Find(&notes)
		if results.Error != nil {
			return nil, count, results.Error
		}
		return notes, count, nil
	default:
		return nil, 0, errors.New("NoFilter")
	}
}

func addUserBean(account_id int64) error {
	user := &model.User{ID: account_id}
	db_err := persistence.DB.Model(&user).Update("Bean", gorm.Expr("Bean + ?", 30)).Error
	if db_err != nil {
		return db_err
	}
	return nil
}

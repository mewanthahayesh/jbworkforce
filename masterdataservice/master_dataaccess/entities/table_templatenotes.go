package entities

import "github.com/jinzhu/gorm"

type TableTemplateNotes struct {
	gorm.Model
	NoteTypeId uint   `gorm:"column:notetypeid;not_null"`
	Note       string `gorm:"column:note"`
	NoteType   TableNoteType
}

func (c TableTemplateNotes) TableName() string {
	return "table_templatenotes"
}

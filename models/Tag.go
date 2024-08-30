package models

type Tag struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"size:100;not null;unique"`

	Blogs []Blog `gorm:"many2many:blog_tags;"`
}

func (Tag) TableName() string {
	return "tags"
}

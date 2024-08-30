package models

type Category struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"size:75"`
	Slug  string `gorm:"size:100;uniqueIndex:idx_category_slug"`

	Blogs []Blog `gorm:"many2many:blog_categories;"`
}

func (Category) TableName() string {
	return "categories"
}

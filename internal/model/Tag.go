package model

type Tag struct {
	*Model
	Name  string
	State uint8
}

func (t Tag) TableName() string {
	return "blog_tag"
}

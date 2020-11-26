package model

type Article struct {
	*Model
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         uint8
}

func (a Article) TableName() string {
	return "blog_article"
}

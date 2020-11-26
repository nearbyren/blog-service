package model

type ArticleTag struct {
	*Model
	TagID     uint32
	ArticleID uint32
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

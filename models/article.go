package models

type Article struct {
	Model

	TagID int `json:"tag_id,omitempty" gorm:"index"` // 创建索引
	Tag   Tag `json:"tag"`

	Title      string `json:"title,omitempty"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
//	tx.Statement.SetColumn("CreatedOn", time.Now().UnixMilli())
//	return
//}
//
//func (a *Article) BeforeUpdate(tx *gorm.DB) (err error) {
//	tx.Statement.SetColumn("ModifiedOn", time.Now().UnixMilli())
//	return
//}

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticle(id int) (article Article) {
	var tag Tag
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Association("Tag").Find(&tag)
	article.Tag = tag
	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

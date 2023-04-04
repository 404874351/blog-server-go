package model

//
// Tag
//  @Description: 文章标签
//
type Tag struct {
	Model
	// 标签名称
	Name			string		`json:"name"         gorm:"not null;size:255"`
}

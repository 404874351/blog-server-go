package model

//
// Category
//  @Description: 文章分类
//
type Category struct {
	Model
	// 分类名称
	Name			string		`json:"name"         gorm:"not null;size:255"`
}

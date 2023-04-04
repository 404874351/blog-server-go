package model

//
// Menu
//  @Description: 菜单
//
type Menu struct {
	Model
	// 菜单代码
	Code		    string		`json:"code"        gorm:"unique;not null;size:255"`
	// 菜单名称
	Name		    string		`json:"name"        gorm:"not null;size:255"`
	// 菜单路径
	Path		    string		`json:"path"        gorm:"not null;size:255"`
	// 菜单组件
	Component		string		`json:"component"   gorm:"not null;size:255"`
	// 菜单图标
	Icon			string		`json:"icon"        gorm:"not null;size:255"`
	// 菜单类型，0具体菜单，1菜单组，默认0
	Type		    int8        `json:"type"        gorm:"not null"`
	// 菜单层级，0顶层，正数代表具体层级，默认0
	Level		    int8		`json:"level"       gorm:"not null"`
	// 父级id，null没有父级，即处于顶层
	ParentId		*int		`json:"parentId"    gorm:""`
	// 是否隐藏，0否，1是，默认0
	Hidden			*int8		`json:"hidden"      gorm:"not null"`
}

const (
	MENU_TYPE_ITEM          int8 = 0
	MENU_TYPE_GROUP         int8 = 1
	MENU_LEVEL_TOP			int8 = 0
	MENU_HIDDEN_DISABLE   	int8 = 0
	MENU_HIDDEN_ENABLE    	int8 = 1
	
	MENU_COMPONENT_LAYOUT   string = "Layout"
)



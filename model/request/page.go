package request

//
// PageVo
//  @Description: 分页 请求对象
//
type PageVo struct {
	// 当前页，默认从1开始
	Current 	int64 		`json:"current" form:"current"`
	// 页面尺寸，默认10
	Size		int64		`json:"size"    form:"size"   `

}

const (
	DEFAULT_PAGE_CURRENT = 1
	DEFAULT_PAGE_SIZE = 10
)

//
// New
//  @Description: 初始化分页
//  @receiver a
//
func (a *PageVo) New() {
	if a.Current < DEFAULT_PAGE_CURRENT {
		a.Current = DEFAULT_PAGE_CURRENT
	}
	if a.Size < DEFAULT_PAGE_SIZE {
		a.Size = DEFAULT_PAGE_SIZE
	}
}

func ValidatePageVo(pageVo PageVo) (bool, error) {
	if pageVo.Current < DEFAULT_PAGE_CURRENT || pageVo.Size < DEFAULT_PAGE_SIZE {
		return false, nil
	}
	return true, nil
}
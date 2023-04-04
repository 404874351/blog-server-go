package response
//
// Page
//  @Description: 分页查询输出
//
type Page struct {
	// 当前页
	Current 	int64			`json:"current"`
	// 页面尺寸
	Size		int64			`json:"size"`
	// 总数
	Total		int64			`json:"total"`
	// 数据列表
	Records		interface{}		`json:"records"`
}

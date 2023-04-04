package request

type ModelDeletedVo struct {
	// id
	ID        	int		`json:"id"          form:"id"`
	// 逻辑删除标志
	Deleted     int8	`json:"deleted"     form:"deleted"    binding:"ValidateDeleted"`
}

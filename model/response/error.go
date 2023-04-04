package response

//
// SystemError
//  @Description: 自定义错误
//
type SystemError struct {
	Code StateCode
}

func (e *SystemError) Error() string {
	if e == nil || e.Code == 0 {
		return "system error"
	}
	return MsgMap[e.Code]
}
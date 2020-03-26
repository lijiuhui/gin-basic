package e

var MsgText = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_EXIST:       "已存在该对象名称",
	ERROR_EXIST_FAIL:  "获取已存在对象失败",
	ERROR_NOT_EXIST:   "该不存在",
	ERROR_GET_S_FAIL:  "获取所有对象失败",
	ERROR_COUNT_FAIL:  "统计失败",
	ERROR_ADD_FAIL:    "新增失败",
	ERROR_EDIT_FAIL:   "修改失败",
	ERROR_DELETE_FAIL: "删除失败",
	ERROR_EXPORT_FAIL: "导出失败",
	ERROR_IMPORT_FAIL: "导入失败",
}

func GetMsg(code int) string {
	msg, ok := MsgText[code]
	if ok {
		return msg
	}
	return MsgText[ERROR]
}

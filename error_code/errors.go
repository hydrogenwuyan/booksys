package error_code

type ERROR_CODE int

const (
	ERROR_CODE_USER_NAME_ERROR     ERROR_CODE = 1
	ERROR_CODE_USER_PASSWORD_ERROR            = 2
	ERROR_CODE_SUCCESS             ERROR_CODE = 200
	ERROR_CODE_DB_ERROR                       = 600
	ERROR_CODE_GENERATE_TOKEN_FAIL            = 601
)

var (
	errorCodeMap = map[ERROR_CODE]string{
		ERROR_CODE_USER_NAME_ERROR:     "管理员用户名不正确",
		ERROR_CODE_USER_PASSWORD_ERROR: "管理员密码不正确",
		ERROR_CODE_DB_ERROR:            "数据库出错",
		ERROR_CODE_GENERATE_TOKEN_FAIL: "token生成失败",
	}
)

func (e ERROR_CODE) String() string {
	return errorCodeMap[e]
}

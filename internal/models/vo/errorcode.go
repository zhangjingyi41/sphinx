package vo

// 错误码常量定义，使用iota自动递增，方便维护
const (
	// 账号相关错误: 1xxxx
	AccountNotExist     Code = 10001
	AccountAlreadyExist Code = 10002
	PasswordIncorrect   Code = 10003
	AccountDisabled     Code = 10004
	InvalidPhoneFormat  Code = 10005

	// 认证相关错误: 2xxxx
	TokenExpired            Code = 20001
	TokenInvalid            Code = 20002
	UnauthorizedAccess      Code = 20003
	LoginRequired           Code = 20004
	RequestParamCheckFailed Code = 20005

	// OAuth相关错误: 3xxxx
	InvalidClient       Code = 30001
	InvalidGrant        Code = 30002
	InvalidScope        Code = 30003
	RedirectURIMismatch Code = 30004
	AccessDenied        Code = 30005

	// 系统相关错误: 9xxxx
	ServerError   Code = 90001
	DatabaseError Code = 90002
	NetworkError  Code = 90003
	UnknownError  Code = 99999
)

// 错误码与错误信息的映射
var codeMessages = map[Code]string{
	// 账号相关错误
	AccountNotExist:     "账号不存在，请注册",
	AccountAlreadyExist: "账号已存在",
	PasswordIncorrect:   "密码不正确",
	AccountDisabled:     "账号已被禁用",
	InvalidPhoneFormat:  "手机号格式不正确",

	// 认证相关错误
	TokenExpired:       "令牌已过期",
	TokenInvalid:       "无效的令牌",
	UnauthorizedAccess: "未授权的访问",
	LoginRequired:      "请先登录",

	// OAuth相关错误
	InvalidClient:       "无效的客户端",
	InvalidGrant:        "无效的授权",
	InvalidScope:        "无效的权限范围",
	RedirectURIMismatch: "重定向URI不匹配",
	AccessDenied:        "拒绝访问",

	// 系统相关错误
	ServerError:   "服务器内部错误",
	DatabaseError: "数据库错误",
	NetworkError:  "网络错误",
	UnknownError:  "未知错误",
}

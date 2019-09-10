package common

const (
	AuthError                = "AuthError"
	ServerError              = "ServerError"
	GetNavAllError           = "GetNavAllError"
	GetFishAllError          = "GetFishAllError"
	CreateNavError           = "CreateNavError"
	ParamsError              = "ParamsError"
	UsernamePasswordError    = "UsernamePasswordError"
	UserHasExistedError      = "UserHasExistedError"
	AccountNoOKError         = "AccountNoOKError"
	CreateResourceError      = "CreateResourceError"
	ListResourceError        = "ListResourceError"
	CreateCommentError       = "CreateCommentError"
	ListCommentError         = "ListCommentError"
	ListCollectionError      = "ListCollectionError"
	ListCommunicationError   = "ListCommunicationError"
	CreateCommunicationError = "CreateCommunicationError"
	CreateFeedbackError      = "CreateFeedbackError"
	CreateBookmarkError      = "CreateBookmarkError"
	TagsTooLongError         = "TagsTooLongError"
	UpdatePwdError           = "UpdatePwdError"
	UpdateInfoError          = "UpdateInfoError"
	InputPwdError            = "InputPwdError"
	ApiSecurityError         = "ApiSecurityError"
)

var errorMsg = map[string]string{
	AuthError:                "认证失败",
	ServerError:              "未知错误，请重试",
	GetNavAllError:           "获取导航列表失败",
	GetFishAllError:          "获取列表失败",
	CreateNavError:           "创建导航失败",
	ParamsError:              "参数错误",
	UsernamePasswordError:    "用户名或密码错误",
	AccountNoOKError:         "由于您的违规行为，您的账户已被冻结",
	UserHasExistedError:      "用户名已存在，请重新输入",
	CreateResourceError:      "发布资源失败，请重试",
	ListResourceError:        "获取资源列表失败，请重试",
	CreateCommentError:       "回复评论失败，请重试",
	ListCommentError:         "获取评论列表失败，请重试",
	ListCommunicationError:   "获取节点信息失败，请重试",
	CreateCommunicationError: "发布失败，请重试",
	CreateFeedbackError:      "提交失败，请重试",
	CreateBookmarkError:      "添加失败，请重试",
	TagsTooLongError:         "添加失败，标签太多",
	ListCollectionError:      "获取导航列表失败",
	UpdatePwdError:           "修改密码失败，请重试",
	InputPwdError:            "密码错误",
	UpdateInfoError:          "修改信息失败",
	ApiSecurityError:         "这位咸鱼，请不要搞事情",
}

//根据错误码获取错误信息
func GetMsg(code string) string {
	msg, ok := errorMsg[code]
	if ok {
		return msg
	}
	return errorMsg[ServerError]
}

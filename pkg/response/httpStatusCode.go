package response

const (
	StatusOk        = 20000
	CodeSuccess     = 20100
	UpdateFriend    = 20101
	ChangePassword  = 20102
	AddEmail        = 20103
	SeenNoti        = 20104
	DeleteSuccess   = 20400
	DeleteNoti      = 20401
	ErrBadRequest   = 40000
	ErrUserNotFound = 40401
)

var msg = map[int]string{
	StatusOk:        "Ok",
	CodeSuccess:     "Success",
	DeleteSuccess:   "Delete no error",
	UpdateFriend:    "Success update to friend",
	ChangePassword:  "Success to reset password",
	AddEmail:        "Add email complete",
	SeenNoti:        "Seen notification success",
	DeleteNoti:      "Delete notification success",
	ErrBadRequest:   "Bad request",
	ErrUserNotFound: "User not found",
}

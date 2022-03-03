package e

const (
	SUCCESS                        = 200
	ERROR                          = 500
	INVALID_PARAMS                 = 400 // 这种错误在函数开始可以检查出来
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 401
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 402
)

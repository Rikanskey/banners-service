package mock

var (
	adminToken = "admin_token"
	userToken  = "user_token"
)

func CheckUserToken(token string) bool {
	return userToken == token
}

func CheckAdminToken(token string) bool {
	return adminToken == token
}

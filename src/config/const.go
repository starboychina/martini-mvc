package config

const (
	//Session
	SessionToken    = "_token_"
	SessionCsrf     = "_CSRF_"
	SessionRedirect = "_redirect_"
	SessionAuth     = "_Auth_"
	//Aes-256
	SecretAes = "yVHlew1jDlZpJ/zSbJ8JPjIc2dBeoLny"
)

var Permission = map[string]interface{}{
    "user":0,
    "employee":10,
    "manage":20,
    "admin":99,
}


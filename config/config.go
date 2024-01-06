package config

func InitConfig() {
	loadEnv()
	initDB()
	initWebsocket()
	initSettings()
}

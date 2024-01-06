package config

func InitConfig() {
	loadEnv()
	initDB()
	initSettings()
}

package util

import "go-server-template/config"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(projectConfig.AppConfig.BaseConfig.JWT_SECRET)
}

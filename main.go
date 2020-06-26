package main
import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// Конфигурация
	var conf Configuration
	// Переменная окружения AWS_PROFILE
	aws_profile := os.Getenv("AWS_PROFILE")
	// По умолчанию default
	if aws_profile == "" {
		aws_profile = "default"
	}
	// Читаем параметры доступа env
	conf.Aws_access_key_id = os.Getenv("AWS_ACCESS_KEY_ID")
	conf.Aws_secret_access_key = os.Getenv("AWS_SECRET_ACCESS_KEY")
	// Если переменные окружения не установлены пробуем прочитать конфиг файл
	if conf.Aws_access_key_id == "" || conf.Aws_secret_access_key == "" {
		err := LoadConfigFromFile(&conf, aws_profile, "~/.aws/credentials")
		if err != nil {
			log.Info(err)
		}	
	} 
	
	// Если конфиг так и не появился
	if conf.Aws_access_key_id == "" || conf.Aws_secret_access_key == "" {
		log.Fatal("Cannot set credentials")
	}

	log.Info(conf)

}


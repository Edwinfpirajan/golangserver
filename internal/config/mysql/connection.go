package mysql

import (
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB = getConnection()

func getConnection() *gorm.DB {
	// dsn := "root:@/distridb"
	// pass := os.Getenv("MYSQLPASSWORD")
	// host := os.Getenv("MYSQLHOST")
	// port := os.Getenv("MYSQLPORT")
	// dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/railway?charset=utf8mb4&parseTime=True&loc=Local", pass, host, port)
	// dsn := "biometrico:_u7825Son@tcp(serverpruebas.tk:3306)/biometricov?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "server_bio:72*4xceG4@tcp(185.68.110.162:3306)/server_bio?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:9oxtg3w0NIF0Xcu9yJXG@tcp(containers-us-west-55.railway.app:6265)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:@tcp(localhost:3306)/testingbio?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err.Error())
	}

	log.Info("Connection Successfully to Mysql")

	return db
}

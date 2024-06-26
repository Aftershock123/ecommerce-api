// database/db.go

package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/Aftershock123/ecommerce/src/config"

)

var DB *gorm.DB

func ConnectDB(cfg config.DBConfig) {
    dsn := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Database + "?parseTime=true"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    
    DB = db
}

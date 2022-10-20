package models

import (
	"fmt"
	"github.com/Dorapoketto/go-gin-example/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"reflect"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)

	// 连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: tablePrefix,
			// 自动创建表时，表明不加S
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Println(err)
	}

	db.Callback().Create().Before("gorm:create").Register("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Before("gorm:update").Register("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	// 获取sql 连接
	sqlDB, _ := db.DB()

	// 设置 sql 连接属性
	// 空闲模式中与数据库最大连接数
	sqlDB.SetMaxIdleConns(50)
	// 设置与数据库的最大连接数
	sqlDB.SetMaxOpenConns(1000)
	// 最长连接时间
	sqlDB.SetConnMaxLifetime(time.Minute)

}

func CloseDB() {
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		log.Println(err)
	}
}

func updateTimeStampForCreateCallback(db *gorm.DB) {
	ctx := db.Statement.Context
	timeFieldsToInit := []string{"CreatedOn", "ModifiedOn"}
	for _, field := range timeFieldsToInit {

		if timeField := db.Statement.Schema.LookUpField(field); timeField != nil {
			switch db.Statement.ReflectValue.Kind() {
			case reflect.Slice, reflect.Array:
				for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
					if _, isZero := timeField.ValueOf(ctx, db.Statement.ReflectValue.Index(i)); isZero {
						timeField.Set(ctx, db.Statement.ReflectValue.Index(i), time.Now().UnixMilli())
					}
				}
			case reflect.Struct:
				if _, isZero := timeField.ValueOf(ctx, db.Statement.ReflectValue); isZero {
					timeField.Set(ctx, db.Statement.ReflectValue, time.Now().UnixMilli())
				}
			}
		}
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	db.Statement.SetColumn("ModifiedOn", time.Now().UnixMilli())
}

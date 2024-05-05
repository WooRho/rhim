package middleware

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"rhim/config"
	"sync"
	"time"
)

var userDbs sync.Map

// GetDb 获取指定数据库
func GetDb(mysql ...*config.Mysql) *gorm.DB {
	var (
		mysqlTmp = &config.Mysql{}
	)
	if len(mysql) == 0 {
		mysqlTmp = config.GetMysql()
	} else {
		mysqlTmp = mysql[0]
	}
	data, ok := userDbs.Load(mysqlTmp.Database)
	if !ok || data.(*gorm.DB) == nil {
		singleDb := NewDatabase(mysqlTmp)
		userDbs.Store(mysqlTmp.Database, singleDb)
		return singleDb
	}
	return data.(*gorm.DB)
}

func NewDatabase(config *config.Mysql) *gorm.DB {

	dblink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%v&collation=%v&loc=Local&parseTime=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Charset,
		config.Collation,
	)
	//加载日志
	slowLogger := logger.New(
		//将标准输出作为Writer
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			//设定慢查询时间阈值为3s
			SlowThreshold: 3 * time.Second,
			//设置日志级别，只有Warn和Info级别会输出慢查询日志
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(mysql.Open(dblink), &gorm.Config{
		SkipDefaultTransaction: false, //跳过事务执行
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "k_", //设置表前缀
			SingularTable: true, //在创建表名时采用单复数  也就是比如用结构体去创建表时 不加s
		},
		DisableForeignKeyConstraintWhenMigrating: true, //逻辑外键
		Logger:                                   slowLogger,
	})
	if err != nil {
		panic(err)
	}

	SetDB(config, db)

	return db
}

func SetDB(config *config.Mysql, db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.MaxIdleConnection)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.MaxConnection)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour * 3)

	log.Println(`正在连接"` + config.Database + `"数据库`)
	err = sqlDB.Ping()
	if err != nil {
		log.Println(config.Database + `"数据库,连接失败,err:` + err.Error())
	} else {
		log.Println(`"` + config.Database + `"数据库,连接成功`)
	}
}

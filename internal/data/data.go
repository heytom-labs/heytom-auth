package data

import (
	"heytom-auth/internal/biz"
	"heytom-auth/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,
	NewGreeterRepo,
	NewUserRepo,
	NewRoleRepo,
	NewPolicyRepo,
	NewApplicationRepo,
	NewAuthRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data) (*Data, func(), error) {
	// 连接数据库
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	// 自动迁移数据模型
	err = db.AutoMigrate(
		&User{},
		&Role{},
		&Policy{},
		&Application{},
		&biz.User{},
		&biz.Token{},
	)

	if err != nil {
		return nil, nil, err
	}

	d := &Data{
		db: db,
	}

	cleanup := func() {
		log.Info("closing the data resources")
		// 关闭数据库连接
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	return d, cleanup, nil
}

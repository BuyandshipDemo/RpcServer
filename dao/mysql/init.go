package mysql

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// TODO: how to save the authentication info.
var dsn = "root:lsh@1030@@tcp(localhost:3306)/helloworld?charset=utf8&parseTime=True&loc=Local"
// AWS mysql
// var dsn = "root:lsh66600504@tcp(personal-test.cdgyma04evtw.ap-southeast-2.rds.amazonaws.com:3306)/helloworld?charset=utf8&parseTime=True&loc=Local"


var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		klog.CtxDebugf(context.Background(), "connect mysql error: %+v", err.Error())
		panic(err)
	}
}

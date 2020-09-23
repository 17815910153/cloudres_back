package tool

import (
	"CloudRes/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)
var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

//进行数据库连接
func OrmEngine(cfg *Config) (*Orm, error) {
	//获取配置信息
	database := cfg.Database
	//连接的url
	conn := database.User + ":" + database.Password + "@(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset + "&parseTime=" + database.ParseTime
	engine, err := xorm.NewEngine(database.Driver, conn)

	if err != nil {
		return nil,err

	}
	engine.ShowSQL(database.ShowSql)
	//进行数据库表同步 TODO
	err = engine.Sync2(new(model.SmsCode), new(model.Member), new(model.FoodCategory))


	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, nil

}
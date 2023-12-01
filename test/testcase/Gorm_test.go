package testcase

import (
	"game_server/server/common/conf"
	"game_server/server/common/orm"
	"game_server/test/common"
	"testing"
	"time"
)

func TestGorm(t *testing.T) {
	config := conf.Database{
		Uri:         common.DatabaseUri,
		MaxLifetime: time.Minute * 60,
		MaxIdleTime: time.Minute * 10,
		MaxIdleConn: 5,
		MaxOpenConn: 10,
	}
	logger, _, err := common.NewLogger(true)
	if err != nil {
		t.Error(err)
		return
	}
	g, err := orm.NewGorm(config, logger)
	if err != nil {
		t.Error(err)
		return
	}
	db := g.DB()
	db.Exec("drop table if exists test_gorm_table")
	if db.Error != nil {
		t.Error(db.Error)
	}

	db.Exec("create table test_gorm_table(id integer primary key auto_increment, name varchar(200), create_at datetime)")
	if db.Error != nil {
		t.Error(db.Error)
	}

	var count int64 = -1
	db.Table("test_gorm_table").Count(&count)
	if db.Error != nil {
		t.Error(db.Error)
	}
	if count != 0 {
		t.Errorf("test_gorm_table not empty, expect: %d, but: %d", 0, count)
	}
}

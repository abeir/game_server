package orm

import (
	"game_server/server/common/conf"
	"game_server/server/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Gorm struct {
	db *gorm.DB
}

func NewGorm(config conf.Database, logger log.ILogger) (*Gorm, error) {
	g := &Gorm{}
	err := g.initialize(&config, logger)
	return g, err
}

func (g *Gorm) DB() *gorm.DB {
	return g.db
}

func (g *Gorm) initialize(config *conf.Database, logger log.ILogger) error {
	slowThreshold := 300 * time.Millisecond

	db, err := gorm.Open(mysql.Open(config.Uri), &gorm.Config{
		Logger:                                   newGormLogger(logger, slowThreshold),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	if err = g.gormPool(db, config); err != nil {
		return err
	}
	g.db = db
	return nil
}

func (g *Gorm) gormPool(db *gorm.DB, dbConfig *conf.Database) error {
	sqldb, err := db.DB()
	if err != nil {
		return err
	}

	dbConfig = g.checkDatabaseConfig(dbConfig, &conf.Database{
		MaxLifetime: time.Hour,
		MaxIdleTime: time.Minute * 15,
		MaxIdleConn: 10,
		MaxOpenConn: 50,
	})

	sqldb.SetConnMaxLifetime(dbConfig.MaxLifetime)
	sqldb.SetConnMaxIdleTime(dbConfig.MaxIdleTime)
	sqldb.SetMaxIdleConns(dbConfig.MaxIdleConn)
	sqldb.SetMaxOpenConns(dbConfig.MaxOpenConn)
	return nil
}

func (g *Gorm) checkDatabaseConfig(dbConfig, defaultConfig *conf.Database) *conf.Database {
	if dbConfig.MaxLifetime <= 0 {
		dbConfig.MaxLifetime = defaultConfig.MaxLifetime
	}
	if dbConfig.MaxIdleTime <= 0 {
		dbConfig.MaxIdleTime = defaultConfig.MaxIdleTime
	}
	if dbConfig.MaxIdleConn <= 0 {
		dbConfig.MaxIdleConn = defaultConfig.MaxIdleConn
	}
	if dbConfig.MaxOpenConn <= 0 {
		dbConfig.MaxOpenConn = defaultConfig.MaxOpenConn
	}
	return dbConfig
}

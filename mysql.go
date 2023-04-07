package gormx

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

type Mysql struct {
	init bool
	db   *gorm.DB
	cfg  *MysqlConfig
	once sync.Once
}

type MysqlConfig struct {
	User     string
	Password string
	Addr     string
	Dbname   string
}

var DefaultMysqlConfig = &MysqlConfig{
	User:     "root",
	Password: "root",
	Addr:     "127.0.0.1",
	Dbname:   "mysql",
}

func NewMysql(cfg *MysqlConfig) *Mysql {
	if cfg == nil {
		cfg = DefaultMysqlConfig
	}
	return &Mysql{
		init: true,
		cfg:  cfg,
	}
}

func (m *Mysql) DB(args ...string) *gorm.DB {
	m.once.Do(
		func() {
			m.db = m.Open(args...)
		},
	)
	return m.db
}

func (m *Mysql) Open(args ...string) *gorm.DB {
	if !m.init || m.cfg == nil {
		log.Error("mysql config not initialized")
		return nil
	}
	user := m.cfg.User
	password := m.cfg.Password
	addr := m.cfg.Addr
	dbname := m.cfg.Dbname

	switch len(args) {
	case 1: // dbname
		dbname = args[0]
	case 2:
		log.Error("num of args should be 0, 1 or 3")
		return nil
	case 3: // dbname, user, password
		dbname = args[0]
		user = args[1]
		password = args[2]
	case 4: // dbname, user, password, addr
		dbname = args[0]
		user = args[1]
		password = args[2]
		addr = args[3]
	default:
		if len(args) > 4 {
			log.Error("num of args should be 0, 1 or 3")
			return nil
		}
	}
	dsn := mysqlDsn(user, password, "tcp", addr, dbname)
	db, err := gorm.Open(
		mysql.Open(dsn), &gorm.Config{
			PrepareStmt:    true,
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
			Logger: logger.New(
				log.StandardLogger(),
				logger.Config{
					IgnoreRecordNotFoundError: true,
					SlowThreshold:             1 * time.Second,
					LogLevel:                  logger.Silent,
				},
			),
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	if err != nil {
		log.Errorf("connect to mysql failed: %v", err)
		return nil
	}
	return db
}

func mysqlDsn(user, password, protocol, addr, dbname string) string {
	params := "charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@%s(%s)/%s?%s", user, password, protocol, addr, dbname, params)
}

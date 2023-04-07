package gormx

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strconv"
	"sync"
	"time"
)

type Pg struct {
	init bool
	db   *gorm.DB
	cfg  *PgConfig
	once sync.Once
}

type PgConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
	Tz       string
}

var DefaultPgConfig = &PgConfig{
	User:     "postgres",
	Password: "postgres",
	Host:     "localhost",
	Port:     5432,
	Dbname:   "postgres",
	Tz:       "Asia/Shanghai",
}

func NewPg(cfg *PgConfig) *Pg {
	if cfg == nil {
		cfg = DefaultPgConfig
	}
	return &Pg{
		init: true,
		cfg:  cfg,
	}
}

func (p *Pg) DB(args ...string) *gorm.DB {
	p.once.Do(
		func() {
			p.db = p.Open(args...)
		},
	)
	return p.db
}

func (p *Pg) Open(args ...string) *gorm.DB {
	if !p.init || p.cfg == nil {
		log.Error("mysql config not initialized")
		return nil
	}
	user := p.cfg.User
	password := p.cfg.Password
	host := p.cfg.Host
	port := p.cfg.Port
	tz := p.cfg.Tz
	dbname := p.cfg.Dbname

	switch len(args) {
	case 1: // dbname
		dbname = args[0]
	case 2:
		log.Error("num of args should be 0, 1, 3, 4, 5 or 6")
		return nil
	case 3: // dbname, user, password
		dbname = args[0]
		user = args[1]
		password = args[2]
	case 4: // dbname, user, password, host
		dbname = args[0]
		user = args[1]
		password = args[2]
		host = args[3]
	case 5: // dbname, user, password, host, port
		dbname = args[0]
		user = args[1]
		password = args[2]
		host = args[3]
		var err error
		port, err = strconv.Atoi(args[4])
		if err != nil {
			log.Error("port should be int")
			return nil
		}
	case 6: // dbname, user, password, host, port, tz
		dbname = args[0]
		user = args[1]
		password = args[2]
		host = args[3]
		var err error
		port, err = strconv.Atoi(args[4])
		if err != nil {
			log.Error("port should be int")
			return nil
		}
		tz = args[5]
	default:
		if len(args) > 6 {
			log.Error("num of args should be 0, 1, 3, 4, 5 or 6")
			return nil
		}
	}
	dsn := pgDsn(host, user, password, dbname, tz, port)
	db, err := gorm.Open(
		postgres.Open(dsn), &gorm.Config{
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
		log.Errorf("connect to pg failed: %v", err)
		return nil
	}
	return db
}

func pgDsn(host, user, password, dbname, tz string, port int) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", host, user, password, dbname, port, tz)
}

package main

import (
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/lucky-pocket/luckyPocket-back/internal/app/notice/repository"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/batch"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/config"
	ent "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/client"
	redis "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/redis/client"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "lvl",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    "func",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	// rotator := &lumberjack.Logger{
	// 	Filename:   "./log/batch.log",
	// 	MaxSize:    5,
	// 	MaxAge:     60,
	// 	MaxBackups: 4,
	// 	LocalTime:  true,
	// }

	logger = zap.New(
		zapcore.NewTee(
			// zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(rotator), zapcore.ErrorLevel),
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), os.Stdout, zapcore.InfoLevel),
		),
	)

	if err := config.Load("./resource/app.yml"); err != nil {
		logger.Fatal(err.Error())
	}
}

func main() {
	defer logger.Sync()

	// data layer configuration.
	redisConfig := config.Data().Redis
	redis, closeRedis, err := redis.NewClient(redisConfig.Addr, redisConfig.Pass, redisConfig.DB)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer closeRedis()

	mysqlConfig := config.Data().Mysql
	ent, closeMysql, err := ent.NewClient(ent.NewMySQLDialect(ent.MysqlDialectOpts{
		User: mysqlConfig.User,
		Pass: mysqlConfig.Pass,
		Host: mysqlConfig.Host,
		Port: mysqlConfig.Port,
		DB:   mysqlConfig.DB,
	}))
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer closeMysql()

	// repository layer configuration.
	noticePool := repository.NewNoticePool(redis)
	noticeRepository := repository.NewNoticeRepository(ent)

	// batch process configuration.
	s := batch.NewScheduler(time.Local)

	if err = s.Register(30*time.Second, batch.NewNoticeSender(&batch.NoticeSenderDeps{
		NoticePool:       noticePool,
		NoticeRepository: noticeRepository,
		Logger:           logger,
	})); err != nil {
		logger.Fatal(err.Error(),
			zap.String("processor", "notice sender"),
		)
	}

	defer s.Stop()
	s.Start()

	select {}
}

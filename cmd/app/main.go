package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"gopkg.in/natefinch/lumberjack.v2"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	ent "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/client"
	redis "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/redis/client"

	blacklist_repo "github.com/lucky-pocket/luckyPocket-back/internal/app/auth/repository/blacklist"
	gamelog_repo "github.com/lucky-pocket/luckyPocket-back/internal/app/game/repository/gamelog"
	ticket_repo "github.com/lucky-pocket/luckyPocket-back/internal/app/game/repository/ticket"
	notice_repo "github.com/lucky-pocket/luckyPocket-back/internal/app/notice/repository"
	pocket_repo "github.com/lucky-pocket/luckyPocket-back/internal/app/pocket/repository"
	user_repo "github.com/lucky-pocket/luckyPocket-back/internal/app/user/repository"

	auth_uc "github.com/lucky-pocket/luckyPocket-back/internal/app/auth/usecase"
	game_uc "github.com/lucky-pocket/luckyPocket-back/internal/app/game/usecase"
	notice_uc "github.com/lucky-pocket/luckyPocket-back/internal/app/notice/usecase"
	pocket_uc "github.com/lucky-pocket/luckyPocket-back/internal/app/pocket/usecase"
	user_uc "github.com/lucky-pocket/luckyPocket-back/internal/app/user/usecase"

	auth_r "github.com/lucky-pocket/luckyPocket-back/internal/app/auth/delivery"
	game_r "github.com/lucky-pocket/luckyPocket-back/internal/app/game/delivery"
	notice_r "github.com/lucky-pocket/luckyPocket-back/internal/app/notice/delivery"
	pocket_r "github.com/lucky-pocket/luckyPocket-back/internal/app/pocket/delivery"
	user_r "github.com/lucky-pocket/luckyPocket-back/internal/app/user/delivery"

	event_type "github.com/lucky-pocket/luckyPocket-back/internal/domain/data/event"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth/jwt"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/config"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/event"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/event/dispatcher"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	v "github.com/lucky-pocket/luckyPocket-back/internal/global/validator"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/client/gauth"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/filter"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/interceptor"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/web/http/middleware"

	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	health_config "github.com/tavsec/gin-healthcheck/config"
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

	rotator := &lumberjack.Logger{
		Filename:   "/var/log/app/app.log",
		MaxSize:    5,
		MaxAge:     60,
		MaxBackups: 4,
		LocalTime:  true,
	}

	logger = zap.New(
		zapcore.NewTee(
			zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(rotator), zapcore.ErrorLevel),
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), os.Stdout, zapcore.InfoLevel),
		),
	)

	if err := config.Load("./resource/app.yml"); err != nil {
		logger.Fatal(err.Error())
	}

	err := v.Initialize(binding.Validator.Engine().(*validator.Validate))
	if err != nil {
		logger.Fatal(err.Error())
	}
}

func main() {
	defer logger.Sync()

	// data layer configuration.
	redisConf := config.Data().Redis
	redis, closeRedis, err := redis.NewClient(redisConf.Addr, redisConf.Pass, redisConf.DB)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer closeRedis()

	mysqlConf := config.Data().Mysql
	entClient, sqlDB, closeMysql, err := ent.NewClient(ent.NewMySQLDialect(ent.MysqlDialectOpts{
		User: mysqlConf.User,
		Pass: mysqlConf.Pass,
		Host: mysqlConf.Host,
		Port: mysqlConf.Port,
		DB:   mysqlConf.DB,
	}))
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer closeMysql()

	if err = ent.Migrate(context.Background(), entClient); err != nil {
		logger.Fatal(err.Error())
	}

	// repository layer configuration.
	blacklistRepo := blacklist_repo.NewBlackListRepository(redis)
	ticketRepo := ticket_repo.NewTicketRepository(redis)
	noticePool := notice_repo.NewNoticePool(redis)
	userRepo := user_repo.NewUserRepository(entClient)
	pocketRepo := pocket_repo.NewPocketRepository(entClient)
	gamelogRepo := gamelog_repo.NewGameLogRepository(entClient)
	noticeRepo := notice_repo.NewNoticeRepository(entClient)

	// helper configuration.
	txManager := tx.NewTxManager()
	jwtIssuer := jwt.NewIssuer([]byte(config.JWT().Secret))
	jwtParser := jwt.NewParser([]byte(config.JWT().Secret))

	gauthConf := config.Web().GAuth
	gauthClient := gauth.NewClient(
		gauthConf.ClientID,
		gauthConf.ClientSecret,
		gauthConf.RedirectURI,
	)

	// event manager configuration.
	noticeDumper := dispatcher.NewNoticePoolDumper(&dispatcher.NoticePoolDumperDeps{
		NoticePool: noticePool,
	})

	em := event.NewManager()
	em.Register(string(event_type.TopicPocketReceived), noticeDumper)
	em.Register(string(event_type.TopicRevealCreated), noticeDumper)

	// usecase layer configuration.
	authUcase := auth_uc.NewAuthUseCase(&auth_uc.Deps{
		UserRepository:      userRepo,
		BlackListRepository: blacklistRepo,
		TxManager:           txManager,
		GAuthClient:         gauthClient,
		JwtParser:           jwtParser,
		JwtIssuer:           jwtIssuer,
	})

	userUcase := user_uc.NewUserUseCase(&user_uc.Deps{
		UserRepository:   userRepo,
		NoticeRepository: noticeRepo,
	})

	noticeUcase := notice_uc.NewNoticeUseCase(&notice_uc.Deps{
		NoticeRepository: noticeRepo,
	})

	pocketUcase := pocket_uc.NewPocketUseCase(&pocket_uc.Deps{
		UserRepository:   userRepo,
		PocketRepository: pocketRepo,
		TxManager:        txManager,
		EventManager:     em,
	})

	gameUcase := game_uc.NewGameUseCase(&game_uc.Deps{
		UserRepository:    userRepo,
		TicketRepository:  ticketRepo,
		GameLogRepository: gamelogRepo,
		TxManager:         txManager,
	})

	// delivery layer configuration.
	authRouter := auth_r.NewAuthRouter(authUcase)
	gameRouter := game_r.NewGameRouter(gameUcase)
	noticeRouter := notice_r.NewNoticeRouter(noticeUcase)
	pocketRouter := pocket_r.NewPocketRouter(pocketUcase)
	userRouter := user_r.NewUserRouter(userUcase)

	authFilter := filter.NewAuthFilter(jwtParser)
	errorFilter := filter.NewErrorFilter()
	logHandler := interceptor.NewLogger(logger)
	ratelimiter := middleware.NewRateLimiter()
	roleFilter := filter.NewRoleFilter()

	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        false,
		SkipPaths:  []string{"/healthz", "/favicon.ico"},
		Context:    nil,
	}))
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://lucky-pocket.site", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
	}))

	healthcheck.New(e, health_config.DefaultConfig(), []checks.Check{
		checks.SqlCheck{Sql: sqlDB},
		&checks.RedisCheck{Client: redis},
	})

	e.Use(errorFilter.Register())
	e.Use(logHandler.Register())
	e.Use(ratelimiter.Register())

	game := e.Group("/games")
	{
		game.Use(authFilter.WithRequired(true))

		game.GET("/free-ticket", gameRouter.GetTicketInfo)
		game.POST("/yut", gameRouter.PlayYut)
	}

	auth := e.Group("/auth")
	{
		auth.GET("/gauth", authRouter.Login)
		auth.POST("/logout", authRouter.Logout)
		auth.POST("/refresh", authRouter.RefreshToken)
	}

	me := e.Group("/users/me")
	{
		me.Use(authFilter.WithRequired(true))

		me.GET("", userRouter.GetMyDetail)
		me.GET("/coins", userRouter.CountCoins)

		me.GET("/pockets/received", pocketRouter.GetMyPockets)

		notice := me.Group("/notices")
		{
			notice.GET("", noticeRouter.GetNotice)
			notice.PATCH("", noticeRouter.CheckAllNotices)
			notice.PATCH("/:noticeID", noticeRouter.CheckNotice)
		}
	}

	user := e.Group("/users")
	{
		user.GET("", userRouter.Search)
		user.GET("/rank", userRouter.GetRanking)
		user.GET("/:userID", userRouter.GetUserDetail)
		user.GET("/:userID/pockets", pocketRouter.GetUserPockets)
	}

	pocket := e.Group("/pockets")
	{
		pocket.GET("/:pocketID", authFilter.WithRequired(false), pocketRouter.GetPocketDetail)

		pocketAuth := pocket.Group("", authFilter.WithRequired(true))
		{
			pocketAuth.POST("", pocketRouter.SendPocket)
			pocketAuth.POST("/:pocketID/sender", pocketRouter.RevealSender)
			pocketAuth.PATCH("/:pocketID/visibility", pocketRouter.SetVisibility)
		}

	}

	admin := e.Group("/admin")
	{
		admin.Use(
			authFilter.WithRequired(true),
			roleFilter.Register(constant.RoleAdmin),
		)

		pprof.RouteRegister(admin, "debug/pprof")
	}

	httpConf := config.Web().HTTP
	if err := e.Run(fmt.Sprintf(":%d", httpConf.Port)); err != nil {
		logger.Fatal(err.Error())
	}
}

package factory

import (
	"fmt"
	"net"
	"time"
	"token-repository/internal/repository"
	"token-repository/internal/server"
	"token-repository/internal/usecase"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DBConnectRetry = 5
const DBConnectRetryInterval = 10

func JSTTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	const layout = "2006-01-02T15:04:05+09:00"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	// config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.EncoderConfig = zap.NewProductionEncoderConfig()
	config.EncoderConfig.EncodeTime = JSTTimeEncoder
	l, err := config.Build()

	l.WithOptions(zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		fmt.Printf("failed to create logger: %v\n", err)
	}
	return l, err
}

func NewOAuth2Repo(l *zap.Logger, user string, pass string, host string, port string, dbName string) (dbRepo *repository.OAuth2Repo, err error) {
	addr := net.JoinHostPort(host, port)
	dsn := user + ":" + pass + "@(" + addr + ")/" + dbName + "?parseTime=true&loc=Local"
	var gormdb *gorm.DB
	for i := 0; i < DBConnectRetry; i++ {
		gormdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			// Success DB connect
			l.Info("DB connect")
			break
		}
		l.Warn("DB connection retry")

		if i == DBConnectRetry {
			l.Error("failed to connect (DB)", zap.Error(err))
			return nil, err
		}

		time.Sleep(DBConnectRetryInterval * time.Second)
	}

	dbRepo = &repository.OAuth2Repo{Conn: gormdb}
	return dbRepo, nil
}

func NewTokenRepoService(l *zap.Logger, dbRepo *repository.OAuth2Repo) *usecase.TokenRepoService {
	return &usecase.TokenRepoService{Logger: l, Repository: dbRepo}
}

func NewServer(l *zap.Logger, port string, svc *usecase.TokenRepoService) *server.Server {
	return &server.Server{Logger: l, Port: port, Service: svc}
}

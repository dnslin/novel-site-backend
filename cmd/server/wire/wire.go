//go:build wireinject
// +build wireinject

package wire

import (
	"novel-site-backend/internal/handler"
	"novel-site-backend/internal/repository"
	"novel-site-backend/internal/server"
	"novel-site-backend/internal/service"
	"novel-site-backend/pkg/app"
	"novel-site-backend/pkg/jwt"
	"novel-site-backend/pkg/log"
	"novel-site-backend/pkg/server/http"
	"novel-site-backend/pkg/sid"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewRatingTypeRepository,
	repository.NewBookRatingRepository,
	repository.NewBookRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewRatingTypeService,
	service.NewBookRatingService,
	service.NewBookService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewRatingTypeHandler,
	handler.NewBookRatingHandler,
	handler.NewBookHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(
	httpServer *http.Server,
	job *server.Job,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("novel-site-backend"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}

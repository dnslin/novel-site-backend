package server

import (
	apiV1 "novel-site-backend/api/v1"
	"novel-site-backend/internal/handler"
	"novel-site-backend/internal/middleware"
	"novel-site-backend/pkg/jwt"
	"novel-site-backend/pkg/log"
	"novel-site-backend/pkg/server/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	bookHandler *handler.BookHandler,
	bookRatingHandler *handler.BookRatingHandler,
	ratingTypeHandler *handler.RatingTypeHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	// docs.SwaggerInfo.BasePath = "/v1"
	// s.GET("/swagger/*any", ginSwagger.WrapHandler(
	// 	swaggerfiles.Handler,
	// 	//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
	// 	ginSwagger.DefaultModelsExpandDepth(-1),
	// 	ginSwagger.PersistAuthorization(true),
	// ))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		middleware.ClientIPMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using novel-site-backend!",
		})
	})

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			// noAuthRouter.POST("/register", userHandler.Register)
			// noAuthRouter.POST("/login", userHandler.Login)
			// noAuthRouter.POST("/books", bookHandler.CreateBook)
			// noAuthRouter.PUT("/books/:id", bookHandler.UpdateBook)
			// noAuthRouter.DELETE("/books/:id", bookHandler.DeleteBook)
			noAuthRouter.GET("/books/:id", bookHandler.GetBook)
			noAuthRouter.POST("/books/list", bookHandler.ListBooks)
			noAuthRouter.POST("/books/search", bookHandler.QuickSearch)

			// 评分类型相关接口
			noAuthRouter.GET("/rating-types", ratingTypeHandler.ListRatingTypes)

			// 书籍评分相关接口
			noAuthRouter.POST("/book-ratings", bookRatingHandler.CreateBookRating)
			// noAuthRouter.PUT("/book-ratings/:id", bookRatingHandler.UpdateBookRating)
			noAuthRouter.GET("/book-ratings/:book_id/rating-stats", bookRatingHandler.GetBookRating)
			// noAuthRouter.GET("/books/:book_id/ratings", bookRatingHandler.ListBookRatings)
			noAuthRouter.GET("/books/sorts", bookHandler.GetAllSorts)
		}
		// // Non-strict permission routing group
		// noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		// {
		// 	noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		// }

		// // Strict permission routing group
		// strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		// {
		// 	strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		// }
	}

	return s
}

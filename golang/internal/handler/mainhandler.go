package handler

import (
	"bivbonus/internal/service"
	"bivbonus/pkg/middlerware"
	"bivbonus/pkg/security"
	"bivbonus/pkg/security/jwt"
	"github.com/fasthttp/router"
)

type BaseHandler struct {
	service      *service.Service
	security     *jwt.Security
	passwordHash *security.PasswordHash
}

type Handler struct {
	BaseHandler
	accountHandler     *AccountHandler
	userHandler        *UserHandler
	teamHandler        *TeamHandler
	serviceHandler     *ServiceHandler
	achievementHandler *AchievementHandler
	marketHandler      *MarketHandler
}

func NewHandler(service *service.Service, security *jwt.Security) *Handler {
	return &Handler{
		BaseHandler: BaseHandler{
			security: security,
			service:  service,
		},
		accountHandler:     NewAccountHandler(service, security),
		userHandler:        NewUserHandler(service, security),
		teamHandler:        NewTeamHandler(service, security),
		serviceHandler:     NewServiceHandler(service, security),
		achievementHandler: NewAchievementHandler(service, security),
		marketHandler:      NewMarketHandler(service, security),
	}
}

func (h *Handler) InitRouter(middlerware middlerware.IMiddleWare) *router.Router {
	r := router.New()
	// swagger
	//r.GET("/docs/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.Handler()))
	// auth
	authGroup := r.Group("/auth")
	authGroup.POST("/registry", middlerware.WareRegistry(nil, h.security))
	authGroup.POST("/login", middlerware.WareLogin(nil, h.security))
	authGroup.POST("/logout", middlerware.WareLogout(nil, h.security))
	// account
	accountGroup := r.Group("/account")
	accountGroup.GET("/full-info", middlerware.WareSecurity(h.accountHandler.GetAccountFullInfo, h.security))
	accountGroup.GET("/{id}", middlerware.WareSecurity(h.accountHandler.GetAccountInfo, h.security))
	accountGroup.GET("/all", middlerware.WareSecurity(h.accountHandler.All, h.security))
	accountGroup.POST("/add-achievement", middlerware.WareSecurity(h.accountHandler.AddAchievement, h.security))
	accountGroup.POST("/add-market", middlerware.WareSecurity(h.accountHandler.AddMarket, h.security))
	// user
	userGroup := r.Group("/user")
	userGroup.POST("/create", middlerware.WareSecurity(h.userHandler.Create, h.security))
	userGroup.GET("/{id}", middlerware.WareSecurity(h.userHandler.GetUserInfo, h.security))
	userGroup.GET("/all", middlerware.WareSecurity(h.userHandler.All, h.security))
	// team
	teamGroup := r.Group("/team")
	teamGroup.POST("/create", middlerware.WareSecurity(h.teamHandler.Create, h.security))
	teamGroup.GET("/{id}", middlerware.WareSecurity(h.teamHandler.GetTeamInfo, h.security))
	teamGroup.GET("/all", middlerware.WareSecurity(h.teamHandler.All, h.security))
	teamGroup.POST("/add-user", middlerware.WareSecurity(h.teamHandler.AddUser, h.security))
	// service
	serviceGroup := r.Group("/service")
	serviceGroup.POST("/create", middlerware.WareSecurity(h.serviceHandler.Create, h.security))
	serviceGroup.GET("/{id}", middlerware.WareSecurity(h.serviceHandler.GetTeamInfo, h.security))
	serviceGroup.GET("/all", middlerware.WareSecurity(h.serviceHandler.All, h.security))
	// market
	marketGroup := r.Group("/market")
	marketGroup.POST("/create", middlerware.WareSecurity(h.marketHandler.Create, h.security))
	marketGroup.GET("/{id}", middlerware.WareSecurity(h.marketHandler.GetMarketInfo, h.security))
	marketGroup.GET("/all", middlerware.WareSecurity(h.marketHandler.All, h.security))
	// achievement
	achievementGroup := r.Group("/achievement")
	achievementGroup.POST("/create", middlerware.WareSecurity(h.achievementHandler.Create, h.security))
	achievementGroup.GET("/{id}", middlerware.WareSecurity(h.achievementHandler.GetAchievementInfo, h.security))
	achievementGroup.GET("/all", middlerware.WareSecurity(h.achievementHandler.All, h.security))
	return r
}

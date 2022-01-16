package handler

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/service"
	"bivbonus/pkg/security/jwt"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

type UserHandler struct {
	BaseHandler
}

func NewUserHandler(service *service.Service, security *jwt.Security) *UserHandler {
	return &UserHandler{
		BaseHandler{
			service: service, security: security,
		},
	}
}

type userInfoResponse struct {
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Patronymic *string `json:"patronymic"`
}

func (u *UserHandler) GetUserInfo(ctx *fasthttp.RequestCtx) {
	pathParamId, err := strconv.ParseInt((ctx.Value("id")).(string), 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	userInfo, err := u.service.UserService.GetUserInfo(pathParamId)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	userInfoResponse := userInfoResponse{
		Name:       userInfo.Name,
		Surname:    userInfo.Surname,
		Patronymic: userInfo.Patronymic,
	}
	response, err := json.Marshal(userInfoResponse)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

type userCreateRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type userCreateResponse struct {
	Id int64 `json:"id"`
}

func (u *UserHandler) Create(ctx *fasthttp.RequestCtx) {
	var req userCreateRequest
	if err := json.Unmarshal(ctx.Request.Body(), &req); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	user := domain.User{
		Name:       &req.Name,
		Surname:    &req.Surname,
		Patronymic: &req.Patronymic,
	}
	id, err := u.service.UserService.Create(user)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(userCreateResponse{Id: id})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

func (u *UserHandler) All(ctx *fasthttp.RequestCtx) {
	all, err := u.service.UserService.All()
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(&all)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

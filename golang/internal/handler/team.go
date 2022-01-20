package handler

import (
	"bivbonus/internal/common"
	"bivbonus/internal/domain"
	"bivbonus/internal/service"
	"bivbonus/pkg/security/jwt"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

type TeamHandler struct {
	BaseHandler
}

func NewTeamHandler(service *service.Service, security *jwt.Security) *TeamHandler {
	return &TeamHandler{
		BaseHandler{
			service: service, security: security,
		},
	}
}

type teamInfoResponse struct {
	Name    *string `json:"name"`
	SysName *string `json:"sysname"`
}

func (h TeamHandler) GetTeamInfo(ctx *fasthttp.RequestCtx) {
	pathParamId, err := strconv.ParseInt((ctx.Value("id")).(string), 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	teamInfo, err := h.service.TeamService.GetTeamInfo(pathParamId)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	teamInfoResponse := teamInfoResponse{
		Name:    teamInfo.Name,
		SysName: teamInfo.Sysname,
	}
	response, err := json.Marshal(teamInfoResponse)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	_, err = ctx.Write(response)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
}

type teamCreateRequest struct {
	Name    *string `json:"name"`
	SysName *string `json:"sysname"`
}

type teamCreateResponse struct {
	Id int64 `json:"id"`
}

func (h TeamHandler) Create(ctx *fasthttp.RequestCtx) {
	var teamReq teamCreateRequest
	err := json.Unmarshal(ctx.Request.Body(), &teamReq)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	team := domain.Team{
		Name:    teamReq.Name,
		Sysname: teamReq.SysName,
	}
	id, err := h.service.TeamService.Create(team)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(&teamCreateResponse{Id: id})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

func (h TeamHandler) All(ctx *fasthttp.RequestCtx) {
	all, err := h.service.TeamService.All()
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	commonResponse := common.CommonResponse{
		Data: all,
	}
	response, err := json.Marshal(&commonResponse)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

type addUserInTeamRequest struct {
	TeamId *int64 `json:"teamId"`
	UserId *int64 `json:"userId"`
}

type addUserInTeamResponse struct {
	Id int64 `json:"id"`
}

func (h TeamHandler) AddUser(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()
	var addUserInTeamRequest addUserInTeamRequest
	if err := json.Unmarshal(body, &addUserInTeamRequest); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	id, err := h.service.TeamService.AddUser(addUserInTeamRequest.TeamId, addUserInTeamRequest.UserId)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	aUr := addUserInTeamResponse{Id: id}
	response, err := json.Marshal(aUr)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

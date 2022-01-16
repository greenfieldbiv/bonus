package handler

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/service"
	"bivbonus/pkg/security/jwt"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

type AchievementHandler struct {
	BaseHandler
}

func NewAchievementHandler(service *service.Service, security *jwt.Security) *AchievementHandler {
	return &AchievementHandler{
		BaseHandler{
			service: service, security: security,
		},
	}
}

type achievementCreateRequest struct {
	Name    *string  `json:"name"`
	SysName *string  `json:"sysname"`
	Level   *int     `json:"level"`
	Cost    *float32 `json:"cost"`
}

type achievementResponse struct {
	Id int64 `json:"id"`
}

func (h AchievementHandler) Create(ctx *fasthttp.RequestCtx) {
	var acReq achievementCreateRequest
	err := json.Unmarshal(ctx.Request.Body(), &acReq)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	achievement := domain.Achievement{
		Name:    acReq.Name,
		Sysname: acReq.SysName,
		Level:   acReq.Level,
		Cost:    acReq.Cost,
	}
	id, err := h.service.AchievementService.Create(achievement)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(&achievementResponse{Id: id})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

type achievementInfoResponse struct {
	Name    *string  `json:"name"`
	Sysname *string  `json:"sysname"`
	Level   *int     `json:"level"`
	Cost    *float32 `json:"cost"`
}

func (h AchievementHandler) GetAchievementInfo(ctx *fasthttp.RequestCtx) {
	pathParamId, err := strconv.ParseInt((ctx.Value("id")).(string), 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	achievementInfo, err := h.service.AchievementService.GetAchievementInfo(pathParamId)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	achieveInfoResp := achievementInfoResponse{
		Name:    achievementInfo.Name,
		Sysname: achievementInfo.Sysname,
		Level:   achievementInfo.Level,
		Cost:    achievementInfo.Cost,
	}
	response, err := json.Marshal(achieveInfoResp)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

func (h AchievementHandler) All(ctx *fasthttp.RequestCtx) {
	all, err := h.service.AchievementService.All()
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

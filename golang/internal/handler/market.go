package handler

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/service"
	"bivbonus/pkg/security/jwt"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

type MarketHandler struct {
	BaseHandler
}

func NewMarketHandler(service *service.Service, security *jwt.Security) *MarketHandler {
	return &MarketHandler{
		BaseHandler{
			service: service, security: security,
		},
	}
}

type marketCreateRequest struct {
	Name    *string  `json:"name"`
	SysName *string  `json:"sysname"`
	Level   *int     `json:"level"`
	Cost    *float32 `json:"cost"`
}

type marketResponse struct {
	Id int64 `json:"id"`
}

func (h MarketHandler) Create(ctx *fasthttp.RequestCtx) {
	var marketReq marketCreateRequest
	err := json.Unmarshal(ctx.Request.Body(), &marketReq)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	market := domain.Market{
		Name:    marketReq.Name,
		Sysname: marketReq.SysName,
		Level:   marketReq.Level,
		Cost:    marketReq.Cost,
	}
	id, err := h.service.MarketService.Create(market)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(&marketResponse{Id: id})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

func (h MarketHandler) All(ctx *fasthttp.RequestCtx) {
	all, err := h.service.MarketService.All()
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

type marketInfoResponse struct {
	Name    *string  `json:"name"`
	SysName *string  `json:"sysname"`
	Level   *int     `json:"level"`
	Cost    *float32 `json:"cost"`
}

func (h MarketHandler) GetMarketInfo(ctx *fasthttp.RequestCtx) {
	pathParamId, err := strconv.ParseInt((ctx.Value("id")).(string), 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	marketInfo, err := h.service.MarketService.GetMarketInfo(pathParamId)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	marketInfoResp := marketInfoResponse{
		Name:    marketInfo.Name,
		SysName: marketInfo.Sysname,
		Level:   marketInfo.Level,
		Cost:    marketInfo.Cost,
	}
	response, err := json.Marshal(marketInfoResp)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

package handler

import (
	"bivbonus/internal/domain"
	"bivbonus/internal/service"
	"bivbonus/pkg/security/jwt"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

type ServiceHandler struct {
	BaseHandler
}

func NewServiceHandler(service *service.Service, security *jwt.Security) *ServiceHandler {
	return &ServiceHandler{
		BaseHandler{
			service: service, security: security,
		},
	}
}

type serviceInfoResponse struct {
	Name    *string  `json:"name"`
	SysName *string  `json:"sysname"`
	Cost    *float32 `json:"cost"`
}

func (s ServiceHandler) GetTeamInfo(ctx *fasthttp.RequestCtx) {
	pathParamId, err := strconv.ParseInt((ctx.Value("id")).(string), 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	serviceInfo, err := s.service.ServiceService.GetServiceInfo(pathParamId)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	serviceInfoResponse := serviceInfoResponse{
		Name:    serviceInfo.Name,
		SysName: serviceInfo.SysName,
		Cost:    serviceInfo.Cost,
	}
	response, err := json.Marshal(serviceInfoResponse)
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

type serviceCreateRequest struct {
	Name    *string  `json:"name"`
	SysName *string  `json:"sysname"`
	Cost    *float32 `json:"cost"`
}

type serviceCreateResponse struct {
	Id int64 `json:"id"`
}

func (s ServiceHandler) Create(ctx *fasthttp.RequestCtx) {
	var serviceReq serviceCreateRequest
	err := json.Unmarshal(ctx.Request.Body(), &serviceReq)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	cService := domain.Service{
		Name:    serviceReq.Name,
		SysName: serviceReq.SysName,
		Cost:    serviceReq.Cost,
	}
	id, err := s.service.ServiceService.Create(cService)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(&serviceCreateResponse{Id: id})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(response)
}

func (s ServiceHandler) All(ctx *fasthttp.RequestCtx) {
	all, err := s.service.ServiceService.All()
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

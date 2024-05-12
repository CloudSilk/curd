package provider

import (
	"context"

	"github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
)

type FunctionalTemplateProvider struct {
	apipb.UnimplementedFunctionalTemplateServer
}

func (u *FunctionalTemplateProvider) Add(ctx context.Context, in *apipb.FunctionalTemplateInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	id, err := model.CreateFunctionalTemplate(model.PBToFunctionalTemplate(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *FunctionalTemplateProvider) Update(ctx context.Context, in *apipb.FunctionalTemplateInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.UpdateFunctionalTemplate(model.PBToFunctionalTemplate(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *FunctionalTemplateProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.DeleteFunctionalTemplate(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *FunctionalTemplateProvider) Query(ctx context.Context, in *apipb.QueryFunctionalTemplateRequest) (*apipb.QueryFunctionalTemplateResponse, error) {
	resp := &apipb.QueryFunctionalTemplateResponse{
		Code: apipb.Code_Success,
	}
	model.QueryFunctionalTemplate(in, resp, false)
	return resp, nil
}

func (u *FunctionalTemplateProvider) GetAll(ctx context.Context, in *apipb.GetAllRequest) (*apipb.GetAllFunctionalTemplateResponse, error) {
	resp := &apipb.GetAllFunctionalTemplateResponse{
		Code: apipb.Code_Success,
	}
	list, err := model.GetAllFunctionalTemplates()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.FunctionalTemplatesToPB(list)
	}

	return resp, nil
}

func (u *FunctionalTemplateProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetFunctionalTemplateDetailResponse, error) {
	resp := &apipb.GetFunctionalTemplateDetailResponse{
		Code: apipb.Code_Success,
	}
	f, err := model.GetFunctionalTemplateByID(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.FunctionalTemplateToPB(f)
	return resp, nil
}

func (u *FunctionalTemplateProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.EnableFunctionalTemplate(in.Id, in.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *FunctionalTemplateProvider) Copy(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	id, err := model.CopyFunctionalTemplate(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

package provider

import (
	"context"

	"github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
)

type FormProvider struct {
	apipb.UnimplementedFormServer
}

func (u *FormProvider) Add(ctx context.Context, in *apipb.FormInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.CreateForm(model.PBToForm(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *FormProvider) Update(ctx context.Context, in *apipb.FormInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.UpdateForm(model.PBToForm(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *FormProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.DeleteForm(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *FormProvider) Query(ctx context.Context, in *apipb.QueryFormRequest) (*apipb.QueryFormResponse, error) {
	resp := &apipb.QueryFormResponse{}
	model.QueryForm(in, resp, false)
	return resp, nil
}

func (u *FormProvider) GetAll(ctx context.Context, in *apipb.GetAllFormRequest) (*apipb.GetAllFormResponse, error) {
	resp := &apipb.GetAllFormResponse{}
	forms, err := model.GetAllForms(in)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.FormsToPB(forms)
	}

	return resp, nil
}

func (u *FormProvider) GetDetail(ctx context.Context, in *apipb.GetFormDetailRequest) (*apipb.GetFormDetailResponse, error) {
	resp := &apipb.GetFormDetailResponse{}
	f, err := model.GetFormById(in.Id, in.ContainerVersions)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.FormToPB(f)
	return resp, nil
}

func (u *FormProvider) GetVersion(ctx context.Context, in *apipb.DelRequest) (*apipb.GetVersionResponse, error) {
	resp := &apipb.GetVersionResponse{}
	f, err := model.GetFormVersionById(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.FormVersionToPB(&f)
	return resp, nil
}

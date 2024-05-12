package provider

import (
	"context"

	"github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
)

type PageProvider struct {
	apipb.UnimplementedPageServer
}

func (u *PageProvider) Add(ctx context.Context, in *apipb.PageInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.CreatePage(model.PBToPage(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *PageProvider) Update(ctx context.Context, in *apipb.PageInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.UpdatePage(model.PBToPage(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *PageProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.DeletePage(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *PageProvider) Query(ctx context.Context, in *apipb.QueryPageRequest) (*apipb.QueryPageResponse, error) {
	resp := &apipb.QueryPageResponse{}
	model.QueryPage(in, resp, false, false)
	return resp, nil
}

func (u *PageProvider) GetAll(ctx context.Context, in *apipb.QueryPageRequest) (*apipb.GetAllPageResponse, error) {
	resp := &apipb.GetAllPageResponse{}
	pages, err := model.GetAllPage(in)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.PagesToPB(pages)
	}

	return resp, nil
}

func (u *PageProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetPageDetailResponse, error) {
	resp := &apipb.GetPageDetailResponse{}
	f, err := model.GetPageByID(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.PageToPB(f)
	return resp, nil
}

func (u *PageProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.EnablePage(in.Id, in.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

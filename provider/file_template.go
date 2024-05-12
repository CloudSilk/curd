package provider

import (
	"context"

	"github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
)

type FileTemplateProvider struct {
	apipb.UnimplementedFileTemplateServer
}

func (u *FileTemplateProvider) Add(ctx context.Context, in *apipb.FileTemplateInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	id, err := model.CreateFileTemplate(model.PBToFileTemplate(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *FileTemplateProvider) Update(ctx context.Context, in *apipb.FileTemplateInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.UpdateFileTemplate(model.PBToFileTemplate(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *FileTemplateProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.DeleteFileTemplate(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *FileTemplateProvider) Query(ctx context.Context, in *apipb.QueryFileTemplateRequest) (*apipb.QueryFileTemplateResponse, error) {
	resp := &apipb.QueryFileTemplateResponse{
		Code: apipb.Code_Success,
	}
	model.QueryFileTemplate(in, resp, false)
	return resp, nil
}

func (u *FileTemplateProvider) GetAll(ctx context.Context, in *apipb.GetAllRequest) (*apipb.GetAllFileTemplateResponse, error) {
	resp := &apipb.GetAllFileTemplateResponse{
		Code: apipb.Code_Success,
	}
	list, err := model.GetAllFileTemplates()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.FileTemplatesToPB(list)
	}

	return resp, nil
}

func (u *FileTemplateProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetFileTemplateDetailResponse, error) {
	resp := &apipb.GetFileTemplateDetailResponse{
		Code: apipb.Code_Success,
	}
	f, err := model.GetFileTemplateByID(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.FileTemplateToPB(f)
	return resp, nil
}

func (u *FileTemplateProvider) Copy(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	id, err := model.CopyFileTemplate(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

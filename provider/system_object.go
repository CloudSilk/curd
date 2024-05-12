package provider

import (
	"context"

	"github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
)

type SystemObjectProvider struct {
	apipb.UnimplementedSystemObjectServer
}

func (u *SystemObjectProvider) Add(ctx context.Context, in *apipb.SystemObjectInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	id, err := model.CreateSystemObject(model.PBToSystemObject(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

func (u *SystemObjectProvider) Update(ctx context.Context, in *apipb.SystemObjectInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.UpdateSystemObject(model.PBToSystemObject(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *SystemObjectProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.DeleteSystemObject(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *SystemObjectProvider) Query(ctx context.Context, in *apipb.QuerySystemObjectRequest) (*apipb.QuerySystemObjectResponse, error) {
	resp := &apipb.QuerySystemObjectResponse{
		Code: apipb.Code_Success,
	}
	model.QuerySystemObject(in, resp, false)
	return resp, nil
}

func (u *SystemObjectProvider) GetAll(ctx context.Context, in *apipb.GetAllSystemObjectRequest) (*apipb.GetAllSystemObjectResponse, error) {
	resp := &apipb.GetAllSystemObjectResponse{
		Code: apipb.Code_Success,
	}
	list, err := model.GetAllSystemObjects(in)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SystemObjectsToPB(list)
	}

	return resp, nil
}

func (u *SystemObjectProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetSystemObjectDetailResponse, error) {
	resp := &apipb.GetSystemObjectDetailResponse{
		Code: apipb.Code_Success,
	}
	f, err := model.GetSystemObjectByID(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.SystemObjectToPB(f)
	return resp, nil
}

func (u *SystemObjectProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.EnableSystemObject(in.Id, in.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *SystemObjectProvider) Copy(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	id, err := model.CopySystemObject(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	return resp, nil
}

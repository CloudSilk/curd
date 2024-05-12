package provider

import (
	"context"

	"github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
)

type MetadataProvider struct {
	apipb.UnimplementedMetadataServer
}

func (u *MetadataProvider) Add(ctx context.Context, in *apipb.MetadataInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.CreateMetadata(model.PBToMetadata(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *MetadataProvider) Update(ctx context.Context, in *apipb.MetadataInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.UpdateMetadata(model.PBToMetadata(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *MetadataProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.DeleteMetadata(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *MetadataProvider) Query(ctx context.Context, in *apipb.QueryMetadataRequest) (*apipb.QueryMetadataResponse, error) {
	resp := &apipb.QueryMetadataResponse{}
	model.QueryMetadata(in, resp, false)
	return resp, nil
}

func (u *MetadataProvider) GetAll(ctx context.Context, in *apipb.QueryMetadataRequest) (*apipb.GetAllMetadataResponse, error) {
	resp := &apipb.GetAllMetadataResponse{
		Code: apipb.Code_Success,
	}
	metadatas, err := model.GetAllMetadatas(in)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MetadatasToPB(metadatas)
	}

	return resp, nil
}

func (u *MetadataProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetMetadataDetailResponse, error) {
	resp := &apipb.GetMetadataDetailResponse{}
	f, err := model.GetMetadataById(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.MetadataToPB(f)
	return resp, nil
}

package provider

import (
	"context"

	"github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
)

type CellProvider struct {
	apipb.UnimplementedCellServer
}

func (u *CellProvider) Add(ctx context.Context, in *apipb.CellInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.CreateCell(model.PBToCell(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *CellProvider) Update(ctx context.Context, in *apipb.CellInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.UpdateCell(model.PBToCell(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *CellProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.DeleteCell(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *CellProvider) Query(ctx context.Context, in *apipb.QueryCellRequest) (*apipb.QueryCellResponse, error) {
	resp := &apipb.QueryCellResponse{
		Code: apipb.Code_Success,
	}
	model.QueryCell(in, resp, false)
	return resp, nil
}

func (u *CellProvider) GetAll(ctx context.Context, in *apipb.GetAllCellRequest) (*apipb.GetAllCellResponse, error) {
	resp := &apipb.GetAllCellResponse{
		Code: apipb.Code_Success,
	}
	list, err := model.GetAllCells(in)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.CellsToPB(list)
	}

	return resp, nil
}

func (u *CellProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetCellDetailResponse, error) {
	resp := &apipb.GetCellDetailResponse{
		Code: apipb.Code_Success,
	}
	f, err := model.GetCellByID(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.CellToPB(f)
	return resp, nil
}

func (u *CellProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.EnableCell(in.Id, in.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *CellProvider) Copy(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := model.CopyCell(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

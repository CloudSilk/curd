package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/CloudSilk/pkg/utils"
	userpb "github.com/CloudSilk/usercenter/proto"
	ucprovider "github.com/CloudSilk/usercenter/provider"
)

var (
	ProjectService       = new(userpb.ProjectClientImpl)
	FormComponentService = new(userpb.FormComponentClientImpl)
)

func Init() {
	if os.Getenv("SERVICE_MODE") != "ALL" {
		config.SetConsumerService(ProjectService)
		config.SetConsumerService(FormComponentService)
	} else {
		projectProvider := new(ucprovider.ProjectProvider)
		ProjectService.Add = projectProvider.Add
		ProjectService.Delete = projectProvider.Delete
		ProjectService.Export = projectProvider.Export
		ProjectService.GetDetail = projectProvider.GetDetail
		ProjectService.Query = projectProvider.Query
		ProjectService.Update = projectProvider.Update

		formPComponentrovider := new(ucprovider.FormComponentProvider)
		FormComponentService.Add = formPComponentrovider.Add
		FormComponentService.Delete = formPComponentrovider.Delete
		FormComponentService.GetDetail = formPComponentrovider.GetDetail
		FormComponentService.Query = formPComponentrovider.Query
		FormComponentService.Update = formPComponentrovider.Update
	}
}

func GetProjectDetail(projectID string) (*userpb.GetProjectDetailResponse, error) {
	return ProjectService.GetDetail(context.Background(), &userpb.GetDetailRequest{
		Id: projectID,
	})
}

func GetProjectFormCount(projectID string) (bool, int32, error) {
	resp, err := GetProjectDetail(projectID)
	if err != nil {
		return true, 0, err
	}
	if resp.Code != userpb.Code_Success {
		return true, 0, fmt.Errorf("GetProjectFormCount Code:%d, Error:%v", resp.Code, resp.Message)
	}
	expired := utils.ParseTime(resp.Data.Expired)
	return !expired.After(time.Now()), resp.Data.FormCount, nil
}

func GetProjectPageCount(projectID string) (bool, int32, error) {
	resp, err := GetProjectDetail(projectID)
	if err != nil {
		return true, 0, err
	}
	if resp.Code != userpb.Code_Success {
		return true, 0, fmt.Errorf("GetProjectPageCount Code:%d,Error:%v", resp.Code, resp.Message)
	}
	expired := utils.ParseTime(resp.Data.Expired)
	return !expired.After(time.Now()), resp.Data.PageCount, nil
}

func GetProjectCellCount(projectID string) (bool, int32, error) {
	resp, err := ProjectService.GetDetail(context.Background(), &userpb.GetDetailRequest{
		Id: projectID,
	})
	if err != nil {
		return true, 0, err
	}
	if resp.Code != userpb.Code_Success {
		return true, 0, fmt.Errorf("GetProjectPageCount Code:%d,Error:%v", resp.Code, resp.Message)
	}
	expired := utils.ParseTime(resp.Data.Expired)
	return !expired.After(time.Now()), resp.Data.PageCount, nil
}

// true-过期
// false-未过期
func TenantExpired(projectID string) (bool, error) {
	resp, err := ProjectService.GetDetail(context.Background(), &userpb.GetDetailRequest{
		Id: projectID,
	})
	if err != nil {
		return true, err
	}
	expired := utils.ParseTime(resp.Data.Expired)
	return !expired.After(time.Now()), nil
}

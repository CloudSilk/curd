package model

import (
	apipb "github.com/CloudSilk/curd/proto"
	commonmodel "github.com/CloudSilk/pkg/model"
	"github.com/CloudSilk/pkg/utils"
)

func PBToForm(in *apipb.FormInfo) *Form {
	return &Form{
		TenantModel: commonmodel.TenantModel{
			Model: commonmodel.Model{
				ID: in.Id,
			},
			TenantID: in.TenantID,
		},
		Name:      in.Name,
		PageName:  in.PageName,
		Group:     in.Group,
		Schema:    in.Schema,
		Public:    in.Public,
		Type:      in.Type,
		Subform:   in.Subform,
		ProjectID: in.ProjectID,
		Versions:  FormVersionsToPB(in.Versions),
		IsMust:    in.IsMust,
	}
}

func FormVersionsToPB(versions []*apipb.FormVersion) []*FormVersion {
	var list []*FormVersion
	for _, version := range versions {
		list = append(list, PBToFormVersion(version))
	}
	return list
}

func PBToFormVersions(versions []*FormVersion) []*apipb.FormVersion {
	var list []*apipb.FormVersion
	for _, version := range versions {
		list = append(list, FormVersionToPB(version))
	}
	return list
}

func PBToFormVersion(in *apipb.FormVersion) *FormVersion {
	return &FormVersion{
		Model: commonmodel.Model{
			ID: in.Id,
		},
		FormID:      in.FormID,
		Version:     in.Version,
		Schema:      in.Schema,
		Description: in.Description,
	}
}

func FormToPB(in Form) *apipb.FormInfo {
	return &apipb.FormInfo{
		Id:        in.ID,
		TenantID:  in.TenantID,
		ProjectID: in.ProjectID,
		Name:      in.Name,
		PageName:  in.PageName,
		Group:     in.Group,
		Schema:    in.Schema,
		Versions:  PBToFormVersions(in.Versions),
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
		Public:    in.Public,
		Type:      in.Type,
		Subform:   in.Subform,
		IsMust:    in.IsMust,
	}
}

func FormsToPB(in []Form) []*apipb.FormInfo {
	var list []*apipb.FormInfo
	for _, f := range in {
		list = append(list, FormToPB(f))
	}
	return list
}

func FormVersionToPB(version *FormVersion) *apipb.FormVersion {
	return &apipb.FormVersion{
		Id:          version.ID,
		FormID:      version.FormID,
		Version:     version.Version,
		Schema:      version.Schema,
		Description: version.Description,
		CreatedAt:   utils.FormatTime(version.CreatedAt),
	}
}

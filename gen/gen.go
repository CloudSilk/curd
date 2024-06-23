package gen

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	stdtpl "text/template"

	curdmodel "github.com/CloudSilk/curd/model"
	apipb "github.com/CloudSilk/curd/proto"
	"github.com/CloudSilk/pkg/model"
)

const defaultGenDir = "./temp"

type GenContext struct {
	Metadata *curdmodel.Metadata
	Service  *curdmodel.Service
	CodeFile *curdmodel.CodeFile
	Template *curdmodel.Template
	Config   Config
}

func GenCode(tplID, mdID string) (string, error) {
	tpl, err := GetTemplateById(tplID)
	if err != nil {
		return "", err
	}
	md, err := GetMetadataById(mdID)
	if err != nil {
		return "", err
	}
	return genCode(tpl, &GenContext{
		Metadata: md,
		Template: tpl,
		Service: &curdmodel.Service{
			Package: md.Package,
		},
		CodeFile: &curdmodel.CodeFile{
			Package: md.Package,
		},
	})
}

func genCode(tpl *curdmodel.Template, ctx *GenContext) (string, error) {

	t := stdtpl.Must(stdtpl.New(tpl.Name).Funcs(stdtpl.FuncMap{
		"GenStructField":              GenStructField,
		"GenClassField":               GenClassField,
		"LowerSnakeCase":              LowerSnakeCase,
		"ToUpper":                     ToUpper,
		"LcFirst":                     LcFirst,
		"ToLower":                     strings.ToLower,
		"Split":                       strings.Split,
		"Join":                        strings.Join,
		"GetDefaultValue":             GetDefaultValue,
		"GenClassDefaultValue":        GenClassDefaultValue,
		"GenQueryCond":                GenQueryCond,
		"GenQueryStructField":         GenQueryStructField,
		"RemoveLastChar":              RemoveLastChar,
		"GenCreateSql":                GenCreateSql,
		"GenCreateColumnSql":          GenCreateColumnSql,
		"TableName":                   curdmodel.NamingStrategy.TableName,
		"CamelName2":                  CamelName2,
		"CamelName":                   CamelName,
		"GenPBMessageProperty":        GenPBMessageProperty,
		"GenPBQueryCond":              GenPBQueryCond,
		"GenPBToStrcut":               GenPBToStrcut,
		"GenStrcutToPB":               GenStrcutToPB,
		"GenDeleteChildren":           GenDeleteChildren,
		"GenDeleteChildrenByParentID": GenDeleteChildrenByParentID,
		"GenSwagQueryParam":           GenSwagQueryParam,
		"GetTemplateById":             GetTemplateById,
		"GetMetadataById":             GetMetadataById,
		"GetMetaDataNameByID":         GetMetaDataNameByID,
		"GetMetadataByName":           GetMetadataByName,
		"ConvertTSType":               ConvertTSType,
		"GenCreateIndexsql":           GenCreateIndexsql,
		"GenCreateUniqueIndexsql":     GenCreateUniqueIndexsql,
		"JsonMarshal":                 curdmodel.ObjectToJsonString,
		"JsonUnmarshal":               curdmodel.JsonToStringArray,
		"GenPageConfig":               GenPageConfig,
		"RecursiveGetRefMetadatas":    RecursiveGetRefMetadatas,
		"RecursiveGetRefMetadatas2":   RecursiveGetRefMetadatas2,
		"GenMenuConfig":               GenMenuConfig,
	}).Parse(tpl.Content))

	buf := &bytes.Buffer{}
	err := t.Execute(buf, ctx)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

type GenServiceRequest struct {
	ServiceID uint
}

func GenService(serviceID string, writeFile bool, resp *apipb.GenServiceResponse) {
	s, err := curdmodel.GetServiceByID(serviceID, false)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		return
	}
	err = genService(s, writeFile)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		return
	}
	resp.Data = curdmodel.ServiceToPB(s)
	resp.Dir = fmt.Sprintf("%s/%s", defaultGenDir, s.Package)
}

func genService(s *curdmodel.Service, writeFile bool) error {
	configs := make(map[string]Config)
	for _, codeFile := range s.CodeFiles {
		if !codeFile.Enable {
			continue
		}
		md, err := GetMetadataById(codeFile.MetadataID)
		if err != nil {
			return err
		}
		if codeFile.Start != "" {
			err = genCodeByTpl(configs, s, codeFile, md, codeFile.Start)
			if err != nil {
				return err
			}
		}

		tplIDs := curdmodel.JsonToStringArray(codeFile.Body)

		for _, tplID := range tplIDs {
			err = genCodeByTpl(configs, s, codeFile, md, tplID)
			if err != nil {
				return err
			}
		}
		if codeFile.End != "" {
			err = genCodeByTpl(configs, s, codeFile, md, codeFile.End)
			if err != nil {
				return err
			}
		}
		if writeFile {
			err = writeString(codeFile.Code, codeFile.Dir, codeFile.Name, s.Package)
			if err != nil {
				return err
			}
		}
	}

	for _, serviceFunctional := range s.ServiceFunctionals {
		if !serviceFunctional.Enable {
			continue
		}
		functionalTemplate, err := curdmodel.GetFunctionalTemplateByID(serviceFunctional.FunctionalTemplateID)
		if err != nil {
			return err
		}
		result, err := GenCodeByFunctionalTemplate(functionalTemplate, serviceFunctional.MetadataID, s)
		if err != nil {
			return err
		}
		if writeFile {
			for _, m := range result {
				err = writeString(m.Code, m.Dir, m.FileName, s.Package)
				if err != nil {
					return err
				}
			}

			if serviceFunctional.GenConfig {
				md, err := GetMetadataById(serviceFunctional.MetadataID)
				if err != nil {
					return err
				}
				c := GenMenuConfig(s, md)
				apiJson := JsonMarshal(c.APIs)
				err = writeString(apiJson, "deployment", fmt.Sprintf("%sAPI.json", md.Name), s.Package)
				if err != nil {
					return err
				}
				menuJson := JsonMarshal(c.Menu)
				err = writeString(menuJson, "deployment", fmt.Sprintf("%sMenu.json", md.Name), s.Package)
				if err != nil {
					return err
				}
				pageConfigJson := genPageConfig(s, md)
				err = writeString(string(pageConfigJson), "deployment", fmt.Sprintf("%sPage.json", md.Name), s.Package)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func writeString(str, dir, fileName, p string) error {
	dir = fmt.Sprintf("%s/%s/%s", defaultGenDir, p, dir)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(dir + "/" + fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(str)
	return err
}

func genCodeByTpl(configs map[string]Config, s *curdmodel.Service, codeFile *curdmodel.CodeFile, md *curdmodel.Metadata, tplID string) error {
	tpl, err := GetTemplateById(tplID)
	if err != nil {
		return err
	}
	config, ok := configs[md.ID]
	if !ok {
		config = GenMenuConfig(s, md)
		configs[md.ID] = config
	}
	code, err := genCode(tpl, &GenContext{
		Metadata: md,
		Template: tpl,
		Service:  s,
		CodeFile: codeFile,
		Config:   config,
	})
	if err != nil {
		return err
	}

	codeFile.Code += "\n\n" + code
	codeFile.TemplateContent += "\n" + tpl.Content
	return nil
}

func GenCodeByFileTemplate(codeFile *curdmodel.FileTemplate, mdID string) (string, string, error) {
	configs := make(map[string]Config)

	md, err := GetMetadataById(mdID)
	if err != nil {
		return "", "", err
	}
	service := &curdmodel.Service{
		Package: codeFile.Package,
	}
	cFile := &curdmodel.CodeFile{
		Package: codeFile.Package,
	}
	return genCodeByFileTemplate(codeFile, md, service, cFile, configs)
}

func genCodeByFileTemplate(codeFile *curdmodel.FileTemplate, md *curdmodel.Metadata, service *curdmodel.Service, cFile *curdmodel.CodeFile, configs map[string]Config) (string, string, error) {
	var err error
	if codeFile.Start != "" {
		err = genCodeByTpl(configs, service, cFile, md, codeFile.Start)
		if err != nil {
			return "", "", err
		}
	}

	tplIDs := curdmodel.JsonToStringArray(codeFile.Body)

	for _, tplID := range tplIDs {
		err = genCodeByTpl(configs, service, cFile, md, tplID)
		if err != nil {
			return "", "", err
		}
	}
	if codeFile.End != "" {
		err = genCodeByTpl(configs, service, cFile, md, codeFile.End)
		if err != nil {
			return "", "", err
		}
	}

	return cFile.TemplateContent, cFile.Code, nil
}

func GenCodeByFunctionalTemplate(functionalTemplate *curdmodel.FunctionalTemplate, mdID string, service *curdmodel.Service) ([]*apipb.GenFunctionalTemplateCodeInfo, error) {
	md, err := GetMetadataById(mdID)
	if err != nil {
		return nil, err
	}
	if service == nil {
		service = &curdmodel.Service{
			Name:    "service",
			Package: "package",
		}
	}
	return genCodeByFunctionalTemplate(functionalTemplate, md, service)
}

func genCodeByFunctionalTemplate(functionalTemplate *curdmodel.FunctionalTemplate, md *curdmodel.Metadata, service *curdmodel.Service) ([]*apipb.GenFunctionalTemplateCodeInfo, error) {
	fileTemplateIDs := curdmodel.JsonToStringArray(functionalTemplate.FileTemplateIDs)
	result := make([]*apipb.GenFunctionalTemplateCodeInfo, 0)
	configs := make(map[string]Config)
	for _, fileTemplateID := range fileTemplateIDs {
		fileTemplate, err := curdmodel.GetFileTemplateByID(fileTemplateID)
		if err != nil {
			return nil, err
		}
		cFile := &curdmodel.CodeFile{
			Name:    fileTemplate.Name,
			Dir:     fileTemplate.Dir,
			Package: fileTemplate.Package,
			Params:  fileTemplate.Params,
		}
		templateContent, code, err := genCodeByFileTemplate(fileTemplate, md, service, cFile, configs)
		if err != nil {
			return nil, err
		}
		fileName := LowerSnakeCase(md.Name)
		if fileTemplate.FileNameSuffix != "" {
			fileName += "_" + fileTemplate.FileNameSuffix
		}
		fileName += getFileExtensionByLanguage(fileTemplate.Language)
		result = append(result, &apipb.GenFunctionalTemplateCodeInfo{
			FileName: fileName,
			Dir:      fileTemplate.Dir,
			Template: templateContent,
			Code:     code,
		})
	}
	return result, nil
}

func getFileExtensionByLanguage(language string) string {
	switch language {
	case "Golang":
		return ".go"
	case "ProtocolBuffer":
		return ".proto"
	case "JSON":
		return ".json"
	case "YAML":
		return ".yaml"
	case "Typescript":
		return ".ts"
	case "Java":
		return ".java"
	case "Properties":
		return ".properties"
	case "Sql":
		return ".sql"
	case "Vue":
		return ".vue"
	case "React":
		return ".tsx"
	default:
		return ""
	}
}

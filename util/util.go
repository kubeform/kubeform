package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	. "github.com/dave/jennifer/jen"
	"github.com/hashicorp/terraform/helper/schema"
)

type TypeData struct {
	Name string
	Spec string
}

type ApisData struct {
	ProviderName string
	Version      string
	StructName   []string
}

func GenerateProviderAPIS(providerName, version string, schmeas []map[string]*schema.Schema, structNames []string) error {
	templatePath := "templates"
	goPath := os.Getenv("GOPATH")
	providerPath := filepath.Join(goPath,
		filepath.Join("src",
			filepath.Join("kubeform.dev",
				filepath.Join("kubeform",
					filepath.Join("apis", providerName)))))
	versionPath := filepath.Join(providerPath, version)

	err := os.MkdirAll(versionPath, 0777)
	if err != nil {
		return err
	}

	for i, structName := range structNames {
		var out string

		TerraformSchemaToStruct(schmeas[i], structName+"Spec", &out)
		typeData := TypeData{
			Name: structName,
			Spec: out,
		}

		templateToGoFile(filepath.Join(templatePath, "types.tmpl"), filepath.Join(versionPath, CamelCaseToSnakeCase(structName)+"_types.go"), typeData)
	}

	apiData := ApisData{
		ProviderName: providerName,
		Version:      version,
		StructName:   structNames,
	}

	templateToGoFile(filepath.Join(templatePath, "doc.tmpl"), filepath.Join(versionPath, "doc.go"), apiData)
	templateToGoFile(filepath.Join(templatePath, "register.tmpl"), filepath.Join(versionPath, "register.go"), apiData)
	templateToGoFile(filepath.Join(templatePath, "provider_register.tmpl"), filepath.Join(providerPath, "register.go"), apiData)

	return nil
}

func TerraformSchemaToStruct(s map[string]*schema.Schema, structName string, out *string) {
	var statements Statement

	for key, value := range s {
		id := SnakeCaseToCamelCase(key)
		switch value.Type {
		case schema.TypeString:
			statements = append(statements, Id(id).String().Tag(map[string]string{"json": key}))
		case schema.TypeInt:
			statements = append(statements, Id(id).Int().Tag(map[string]string{"json": key}))
		case schema.TypeBool:
			statements = append(statements, Id(id).Bool().Tag(map[string]string{"json": key}))
		case schema.TypeFloat:
			statements = append(statements, Id(id).Float64().Tag(map[string]string{"json": key}))
		case schema.TypeMap:
			switch value.Elem.(type) {
			case *schema.Schema:
				switch value.Elem.(*schema.Schema).Type {
				case schema.TypeInt:
					statements = append(statements, Id(id).Map(String()).Int().Tag(map[string]string{"json": key}))
				case schema.TypeFloat:
					statements = append(statements, Id(id).Map(String()).Float64().Tag(map[string]string{"json": key}))
				case schema.TypeBool:
					statements = append(statements, Id(id).Map(String()).Bool().Tag(map[string]string{"json": key}))
				case schema.TypeString:
					statements = append(statements, Id(id).Map(String()).String().Tag(map[string]string{"json": key}))
				}
			case *schema.Resource:
				statements = append(statements, Id(id).Map(String()).String().Tag(map[string]string{"json": key}))
				TerraformSchemaToStruct(value.Elem.(*schema.Resource).Schema, structName+id, out)
			default:
				statements = append(statements, Id(id).Map(String()).String().Tag(map[string]string{"json": key}))
			}
		default:
			switch value.Elem.(type) {
			case *schema.Schema:
				switch value.Elem.(*schema.Schema).Type {
				case schema.TypeInt:
					statements = append(statements, Id(id).Index().Int64().Tag(map[string]string{"json": key}))
				case schema.TypeFloat:
					statements = append(statements, Id(id).Index().Float64().Tag(map[string]string{"json": key}))
				case schema.TypeBool:
					statements = append(statements, Id(id).Index().Bool().Tag(map[string]string{"json": key}))
				case schema.TypeString:
					statements = append(statements, Id(id).Index().String().Tag(map[string]string{"json": key}))
				}
			case *schema.Resource:
				statements = append(statements, Id(id).Index().Id(structName).Tag(map[string]string{"json": key}))
				TerraformSchemaToStruct(value.Elem.(*schema.Resource).Schema, structName+id, out)
			default:
				statements = append(statements, Id(id).Index().String().Tag(map[string]string{"json": key}))
			}

		}
	}

	c := Type().Id(structName).Struct(statements...)
	*out = *out + fmt.Sprintf("%#v\n\n", c)
}

func templateToGoFile(templateFile, generatedFilePath string, templateData interface{}) {
	tmpl := template.Must(template.ParseFiles(templateFile))
	var buffer bytes.Buffer
	err := tmpl.Execute(&buffer, templateData)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ioutil.WriteFile(generatedFilePath, buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func SnakeCaseToCamelCase(inputUnderScoreStr string) (camelCase string) {
	isToUpper := false

	for k, v := range inputUnderScoreStr {
		if k == 0 {
			camelCase = strings.ToUpper(string(inputUnderScoreStr[0]))
		} else {
			if isToUpper {
				camelCase += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else {
					camelCase += string(v)
				}
			}
		}
	}
	return
}

func CamelCaseToSnakeCase(str string) string {
	snake := regexp.MustCompile("(.)([A-Z][a-z]+)").ReplaceAllString(str, "${1}_${2}")
	snake = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

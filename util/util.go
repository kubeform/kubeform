package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"text/template"

	. "github.com/dave/jennifer/jen"
	"github.com/gobuffalo/flect"
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

var execeptionList = map[string]string{
	"ConfigConfigurationRecorderStatus": "ConfigConfigurationRecorderStatus_",
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
		if val, ok := execeptionList[structName]; ok {
			structName = val
			structNames[i] = val
		}
		genSecret := false
		TerraformSchemaToStruct(schmeas[i], structName+"Spec", providerName, true, &genSecret, &out)
		typeData := TypeData{
			Name: structName,
			Spec: out,
		}

		templateToGoFile(filepath.Join(templatePath, "types.tmpl"), filepath.Join(versionPath, flect.Underscore(structName)+"_types.go"), typeData)
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

func TerraformSchemaToStruct(s map[string]*schema.Schema, structName, providerName string, genProviderRef bool, genSecret *bool, out *string) {
	var statements Statement
	var keys []string
	for k := range s {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		value := s[key]
		id := flect.Capitalize(flect.Camelize(key))

		if value.Removed != "" {
			continue
		}

		jk := flect.Camelize(key) // json key
		tk := key                 // terraform key
		if value.Optional || value.Computed {
			statements = append(statements, Comment("// +optional"))
			jk = jk + ",omitempty"
			tk = tk + ",omitempty"
		}

		if value.MaxItems != 0 {
			statements = append(statements, Comment("// +kubebuilder:validation:MaxItems="+strconv.Itoa(value.MaxItems)))
		}

		if value.MinItems != 0 {
			statements = append(statements, Comment("// +kubebuilder:validation:MinItems="+strconv.Itoa(value.MinItems)))
		}

		if value.Type == schema.TypeSet {
			statements = append(statements, Comment("// +kubebuilder:validation:UniqueItems=true"))
		}

		if value.Sensitive {
			if value.Type == schema.TypeString {
				statements = append(statements, Id(id).String().Tag(map[string]string{"json": "-", "tf": tk, "sensitive": "true"}))
			} else if value.Type == schema.TypeMap {
				statements = append(statements, Id(id).Map(String()).String().Tag(map[string]string{"json": "-", "tf": tk, "sensitive": "true"}))
			}
			*genSecret = true
			continue
		}

		if value.Deprecated != "" {
			statements = append(statements, Comment("// Deprecated"))
		}

		switch value.Type {
		case schema.TypeString:
			statements = append(statements, Id(id).String().Tag(map[string]string{"json": jk, "tf": tk}))
		case schema.TypeInt:
			statements = append(statements, Id(id).Int().Tag(map[string]string{"json": jk, "tf": tk}))
		case schema.TypeBool:
			statements = append(statements, Id(id).Bool().Tag(map[string]string{"json": jk, "tf": tk}))
		case schema.TypeFloat:
			statements = append(statements, Id(id).Qual("encoding/json", "Number").Tag(map[string]string{"json": jk, "tf": tk}))
		case schema.TypeMap:
			switch value.Elem.(type) {
			case *schema.Schema:
				switch value.Elem.(*schema.Schema).Type {
				case schema.TypeInt:
					statements = append(statements, Id(id).Map(String()).Int().Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeFloat:
					statements = append(statements, Id(id).Map(String()).Qual("encoding/json", "Number").Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeBool:
					statements = append(statements, Id(id).Map(String()).Bool().Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeString:
					statements = append(statements, Id(id).Map(String()).String().Tag(map[string]string{"json": jk, "tf": tk}))
				}
			case *schema.Resource:
				statements = append(statements, Id(id).Map(String()).Id(structName+id).Tag(map[string]string{"json": jk, "tf": tk}))
				TerraformSchemaToStruct(value.Elem.(*schema.Resource).Schema, structName+id, providerName, false, genSecret, out)
			default:
				statements = append(statements, Id(id).Map(String()).String().Tag(map[string]string{"json": jk, "tf": tk}))
			}
		default:
			switch value.Elem.(type) {
			case *schema.Schema:
				switch value.Elem.(*schema.Schema).Type {
				case schema.TypeInt:
					statements = append(statements, Id(id).Index().Int64().Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeFloat:
					statements = append(statements, Id(id).Index().Qual("encoding/json", "Number").Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeBool:
					statements = append(statements, Id(id).Index().Bool().Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeString:
					statements = append(statements, Id(id).Index().String().Tag(map[string]string{"json": jk, "tf": tk}))
				}
			case *schema.Resource:
				statements = append(statements, Id(id).Index().Id(structName+id).Tag(map[string]string{"json": jk, "tf": tk}))
				TerraformSchemaToStruct(value.Elem.(*schema.Resource).Schema, structName+id, providerName, false, genSecret, out)
			default:
				log.Fatalf("Provider %s has resource %s type %s.%s with unknown schema type %s", providerName, structName, structName, id, value.Elem)
			}

		}
	}

	if genProviderRef {
		if *genSecret {
			statements = append(Statement{Id("KubeFormSecret").Id("*core.LocalObjectReference").Tag(map[string]string{"json": "secret,omitempty", "tf": "-"}).Line()}, statements...)
		}
		statements = append(Statement{Id("ProviderRef").Id("core.LocalObjectReference").Tag(map[string]string{"json": "providerRef", "tf": "-"}).Line()}, statements...)
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

package config

import (
	"fmt"
	"strings"
)

type OpenAPIV2Output struct {
	Dir                        string `yaml:"dir"`
	ImportPrefix               string `yaml:"import_prefix"`
	File                       string `yaml:"file"`
	AllowDeleteBody            *bool  `yaml:"allow_delete_body"`
	GRPCApiConfiguration       string `yaml:"grpc_api_configuration"`
	AllowMerge                 *bool  `yaml:"allow_merge"`
	MergeFileName              string `yaml:"merge_file_name"`
	JSONNamesForFields         *bool  `yaml:"json_names_for_fields"`
	RepeatedPathParamSeparator string `yaml:"repeated_path_param_separator"`
	AllowRepeatedFieldsInBody  *bool  `yaml:"allow_repeated_fields_in_body"`
	IncludePackageInTags       *bool  `yaml:"include_package_in_tags"`
	FQNForOpenAPINAme          *bool  `yaml:"fqn_for_openapi_name"`
	UseGoTemplates             *bool  `yaml:"use_go_templates"`
	DisableDefaultErrors       *bool  `yaml:"disable_default_errors"`
	EnumsAsInts                *bool  `yaml:"enums_as_ints"`
	SimpleOperationIDs         *bool  `yaml:"simple_operation_ids"`
	OpenAPIConfiguration       string `yaml:"openapi_configuration"`
	GenerateUnboundMethods     *bool  `yaml:"generate_unbound_methods"`
}

func bool2Str(s *bool) string {
	if *s {
		return "true"
	}
	return "false"
}

func (output *OpenAPIV2Output) BuildArgs() []string {
	cmd := strings.Builder{}
	cmd.WriteString("--openapiv2_out=")
	cmd.WriteString(output.Dir)

	opts := make(map[string]string, 0)

	if output.ImportPrefix != "" {
		opts["import_prefix"] = output.ImportPrefix
	}
	if output.File != "" {
		opts["file"] = output.File
	}
	if output.AllowDeleteBody != nil {
		opts["allow_delete_body"] = bool2Str(output.AllowDeleteBody)
	}
	if output.GRPCApiConfiguration != "" {
		opts["grpc_api_configuration"] = output.GRPCApiConfiguration
	}
	if output.AllowMerge != nil {
		opts["allow_merge"] = bool2Str(output.AllowMerge)
	}
	if output.MergeFileName != "" {
		opts["merge_file_name"] = output.MergeFileName
	}
	if output.JSONNamesForFields != nil {
		opts["json_names_for_fields"] = bool2Str(output.JSONNamesForFields)
	}
	if output.RepeatedPathParamSeparator != "" {
		opts["repeated_path_param_separator"] = output.RepeatedPathParamSeparator
	}
	if output.AllowRepeatedFieldsInBody != nil {
		opts["allow_repeated_fields_in_body"] = bool2Str(output.AllowRepeatedFieldsInBody)
	}
	if output.IncludePackageInTags != nil {
		opts["include_package_in_tags"] = bool2Str(output.IncludePackageInTags)
	}
	if output.FQNForOpenAPINAme != nil {
		opts["fqn_for_openapi_name"] = bool2Str(output.FQNForOpenAPINAme)
	}
	if output.UseGoTemplates != nil {
		opts["use_go_templates"] = bool2Str(output.UseGoTemplates)
	}
	if output.DisableDefaultErrors != nil {
		opts["disable_default_errors"] = bool2Str(output.DisableDefaultErrors)
	}
	if output.EnumsAsInts != nil {
		opts["enums_as_ints"] = bool2Str(output.EnumsAsInts)
	}
	if output.SimpleOperationIDs != nil {
		opts["simple_operation_ids"] = bool2Str(output.SimpleOperationIDs)
	}
	if output.OpenAPIConfiguration != "" {
		opts["openapi_configuration"] = output.OpenAPIConfiguration
	}
	if output.GenerateUnboundMethods != nil {
		opts["generate_unbound_methods"] = bool2Str(output.GenerateUnboundMethods)
	}

	i := 0
	for k, v := range opts {
		if i > 0 {
			cmd.WriteString(" ")
		}
		cmd.WriteString(fmt.Sprintf("--openapiv2_opt %s=%s", k, v))
		i++
	}
	cmd.WriteString(":")
	cmd.WriteString(output.Dir)
	return []string{cmd.String()}
}

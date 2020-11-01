package config

import (
	"fmt"
	"strings"
)

type GRPCGatewayOutput struct {
	Definitions                []string `yaml:"definitions"`
	Dir                        string   `yaml:"dir"`
	ImportPrefix               string   `yaml:"import_prefix"`
	ImportPath                 string   `yaml:"import_path"`
	RegisterFuncSuffix         string   `yaml:"register_func_suffix"`
	RequestContext             *bool    `yaml:"request_context"`
	AllowDeleteBody            *bool    `yaml:"allow_delete_body"`
	GRPCApiConfiguration       string   `yaml:"grpc_api_configuration"`
	Paths                      string   `yaml:"paths"`
	Module                     string   `yaml:"module"`
	AllowRepeatedFieldsInBody  *bool    `yaml:"allow_repeated_fields_in_body"`
	RepeatedPathParamSeparator string   `yaml:"repeated_path_param_separator"`
	AllowPatchFeature          *bool    `yaml:"allow_patch_feature"`
	OmitPackageDoc             *bool    `yaml:"omit_package_doc"`
	Standalone                 *bool    `yaml:"standalone"`
	WarnOnUnboundMethods       *bool    `yaml:"warn_on_unbound_methods"`
	GenerateUnboundMethods     *bool    `yaml:"generate_unbound_methods"`
}

func (output *GRPCGatewayOutput) BuildArgs() []string {
	cmd := strings.Builder{}
	cmd.WriteString("--grpc-gateway_out")
	if len(output.Definitions) > 0 {
		cmd.WriteString("=")
	}
	for i, def := range output.Definitions {
		if i > 0 {
			cmd.WriteString(",")
		}
		cmd.WriteString(def)
	}
	cmd.WriteString(" ")
	cmd.WriteString(output.Dir)
	cmd.WriteString(" ")

	opts := make(map[string]string, 0)

	if output.ImportPrefix != "" {
		opts["import_prefix"] = output.ImportPrefix
	}
	if output.ImportPath != "" {
		opts["import_path"] = output.ImportPath
	}
	if output.RegisterFuncSuffix != "" {
		opts["register_func_suffix"] = output.RegisterFuncSuffix
	}
	if output.RequestContext != nil {
		opts["request_context"] = bool2Str(output.RequestContext)
	}
	if output.AllowDeleteBody != nil {
		opts["allow_delete_body"] = bool2Str(output.AllowDeleteBody)
	}
	if output.GRPCApiConfiguration != "" {
		opts["grpc_api_configuration"] = output.GRPCApiConfiguration
	}
	if output.Paths != "" {
		opts["paths"] = output.Paths
	}
	if output.Module != "" {
		opts["module"] = output.Module
	}
	if output.AllowRepeatedFieldsInBody != nil {
		opts["allow_repeated_fields_in_body"] = bool2Str(output.AllowRepeatedFieldsInBody)
	}
	if output.RepeatedPathParamSeparator != "" {
		opts["repeated_path_param_separator"] = output.RepeatedPathParamSeparator
	}
	if output.AllowPatchFeature != nil {
		opts["allow_patch_feature"] = bool2Str(output.AllowPatchFeature)
	}
	if output.OmitPackageDoc != nil {
		opts["omit_package_doc"] = bool2Str(output.OmitPackageDoc)
	}
	if output.Standalone != nil {
		opts["standalone"] = bool2Str(output.Standalone)
	}
	if output.WarnOnUnboundMethods != nil {
		opts["warn_on_unbound_methods"] = bool2Str(output.WarnOnUnboundMethods)
	}
	if output.GenerateUnboundMethods != nil {
		opts["generate_unbound_methods"] = bool2Str(output.GenerateUnboundMethods)
	}

	i := 0
	for k, v := range opts {
		if i > 0 {
			cmd.WriteString(" ")
		}
		cmd.WriteString(fmt.Sprintf("--grpc-gateway_opt %s=%s", k, v))
		i++
	}
	cmd.WriteString(":")
	cmd.WriteString(output.Dir)
	return []string{cmd.String()}
}

# OMG

Oh-my-GRPC is a draft of a tool to ease the pain of maintaining GRPC protofiles.

## Installation

First, you must install:

- [buf](https://buf.build/docs/installation)
- [protoc](https://grpc.io/docs/protoc-installation/)
- The protoc plugin that you want:
  - [Go](https://grpc.io/docs/languages/go/quickstart/)
  - [GRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway) (OpenAPI or Swagger generation included)
  - Other languages are not supported yet.

## Getting started

omg is a simple app that uses `buf` and `protodep` to help you to deal with gRPC code generation and dependency
management.

Instead of dealing with multiple commands with Makefiles and other building systems, you can just create an `omg.yaml`
that can configure the basics of what you need.

So, this was built on top of my needs for some projects. That is why the project has only the Go language support. At least, for now.

## Commands

### ls

    $ omg ls

Output:

    src/test1.proto
    src/test2.proto
    src/test3.proto

### lint

    $ omg lint

Output:

    src/test1.proto:64:9:Service name "accounts" should be PascalCase, such as "Accounts".
    src/test2:102:10:Field name "contentType" should be lower_snake_case, such as "content_type".
    src/test2.proto:113:10:Field name "contentType" should be lower_snake_case, such as "content_type".

### build

    $ omg build

Output:

    Processing src/test1.proto ...
      ok
    Processing src/test2.proto ...
      ok

## Configuration

All configurations are defined on a `omg.yaml` at the current directory.

- **src**: `[]string`

  List of folders, or individual files, that should be used. When building, this list will be used to generate files. These files will also be added as included files when running `protoc` (with `-I` option).

  If you need to ignore any file, you can use `!` as a prefix for the file path. Ex:

  ```
  src:
    - src
    - !src/test1.proto
  ```

- **includes**: `[]string`

  List of folders, or individual files, that should be included when generating files.

- **go**:

  Configuration for the generation for Go.

  - **dir**: `string` _Required_

    The output dir for the generated files.

  - **annotate_code**: `boolean`

    Whether to store annotations. Defines the `annotate_code` options to the `protoc-gen-go`. (Check `protoc-gen-go` documentation to know more).

  - **definitions**: `[]string`

    Map of protofiles and its implementation packages. Useful when using non standard generators like _gogo_.

  - **import_prefix**: `string`

    Prefix to prepend to import paths.

  - **paths**: `string`

    Defines the `paths` options to the `protoc-gen-go`. Possible values are `import` and `source_relative`. (Check `protoc-gen-go` documentation to know more).

  - **plugins**: `[]string`

    The list of plugins that will be used for the generation.

- **grpc_gateway**:

  Configuration for the generation of the gateway files for Go.

  - **dir**: `string` _Required_

    The output dir for the generated files.

  - **definitions**: `[]string`

    Represents the option `definitions` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **import_prefix**: `string`

    Represents the option `import_prefix` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **import_path**: `string`

    Represents the option `import_path` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **register_func_suffix**: `string`

    Represents the option `register_func_suffix` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **request_context**: `boolean`

    Represents the option `request_context` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **allow_delete_body**: `boolean`

    Represents the option `allow_delete_body` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **grpc_api_configuration**: `string`

    Represents the option `grpc_api_configuration` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **paths**: `string`

    Represents the option `paths` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **module**: `string`

    Represents the option `module` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **allow_repeated_fields_in_body**: `boolean`

    Represents the option `allow_repeated_fields_in_body` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **repeated_path_param_separator**: `string`

    Represents the option `repeated_path_param_separator` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **allow_patch_feature**: `boolean`

    Represents the option `allow_patch_feature` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **omit_package_doc**: `boolean`

    Represents the option `omit_package_doc` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **standalone**: `boolean`

    Represents the option `standalone` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **warn_on_unbound_methods**: `boolean`

    Represents the option `warn_on_unbound_methods` for `grpc-gateway` (check the grpc-gateway documentation to know more).

  - **generate_unbound_methods**: `boolean`

    Represents the option `generate_unbound_methods` for `grpc-gateway` (check the grpc-gateway documentation to know more).

- **openapiv2**:

  - **dir**: `string` _Required_

    The output dir for the generated files.

  - **import_prefix**: `string`

    Represents the option `import_prefix` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **file**: `string`

    Represents the option `file` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **allow_delete_body**: `string`

    Represents the option `allow_delete_body` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **grpc_api_configuration**: `string`

    Represents the option `grpc_api_configuration` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **allow_merge**: `string`

    Represents the option `allow_merge` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **merge_file_name**: `string`

    Represents the option `merge_file_name` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **json_names_for_fields**: `string`

    Represents the option `json_names_for_fields` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **repeated_path_param_separator**: `string`

    Represents the option `repeated_path_param_separator` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **allow_repeated_fields_in_body**: `string`

    Represents the option `allow_repeated_fields_in_body` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **include_package_in_tags**: `string`

    Represents the option `include_package_in_tags` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **fqn_for_openapi_name**: `string`

    Represents the option `fqn_for_openapi_name` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **use_go_templates**: `string`

    Represents the option `use_go_templates` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **disable_default_errors**: `string`

    Represents the option `disable_default_errors` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **enums_as_ints**: `string`

    Represents the option `enums_as_ints` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **simple_operation_ids**: `string`

    Represents the option `simple_operation_ids` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **openapi_configuration**: `string`

    Represents the option `openapi_configuration` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

  - **generate_unbound_methods**: `string`

    Represents the option `generate_unbound_methods` for `openapiv2` (check the grpc-gateway openapiv2 documentation to know more).

Example:

```yaml
src:
  - src
includes:
  - proto
go:
  definitions:
    - Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp
    - Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration
    - Mgoogle/protobuf/empty.proto=github.com/golang/protobuf/ptypes/empty
    - Mgoogle/api/annotations.proto=google/api
    - Mgoogle/api/http.proto=google/api
    - Mgoogle/protobuf/field_mask.proto=google.golang.org/genproto/protobuf/field_mask
    - Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any
    - Mgoogle/rpc/status.proto=google/rpc
  dir: output
  plugins: ["grpc"]
```

<div align="center">
    <h1>protoc-gen-validatex</h1>
</div>

<div align="center">

| [English](https://github.com/protoc-gen/protoc-gen-validatex) | [中文简体](docs/README_zh-CN.md) |

</div>

---

`protoc-gen-validatex` is a plugin for `protoc` designed to simplify and automate the generation of common parameter validation rules. By using parameter options and extensions, users can easily configure validation rules and automatically generate the corresponding validators. The plugin also supports integration with internationalization (i18n), allowing the generation of multilingual error messages based on different language environments, thereby enhancing the user experience.

## Key Features
- **Automated Rule Generation**: This is a fundamental feature of the `protoc` plugin, enabling automatic generation of common validation rules.
- **Multilingual Support**: Naturally supports multiple languages, allowing quick adaptation to various projects and business scenarios.
- **MIT License**: Users can freely fork and modify the project without concerns, under the MIT license.

## Quick Start
You can refer to the [example](./example) for a quick start. Below are some code snippets for reference:

### Definition
```proto
syntax = "proto3";

option go_package = "github.com/protoc-gen/protoc-gen-validatex/example;main";

import "validatex/validatex.proto";

message SignInRequest {
  string email = 1 [(validatex.rules).string.email = true];
  string password = 2 [(validatex.rules).string = {min_len: 5, max_len: 50}];
}
```

### Generation
```shell
go install github.com/protoc-gen/protoc-gen-validatex
protoc --proto_path=. \
       --go_out=paths=source_relative:. \
       --validatex_out=paths=source_relative:. \
       --validatex_opt=i18n_dir=./example/i18n \
       ./example/*.proto
```

### Usage:
```go
package main

import (
    "context"
    "github.com/protoc-gen/protoc-gen-validatex/pkg/validatex"
    "log"
)

func main() {
    // Create a SignInRequest instance
    req := &SignInRequest{
        Email:    "test.com",
        Password: "password123",
    }

    ctx := context.WithValue(context.Background(), validatex.KeyXLang, "zh")

    // Validate the SignInRequest instance
    if err := req.Validate(ctx); err != nil {
        log.Fatalf("validation failed: %v", err)
    }

    log.Println("validation passed")
}
```

## License
This project is licensed under the MIT License. See [LICENSE](./LICENSE) for the full license text.

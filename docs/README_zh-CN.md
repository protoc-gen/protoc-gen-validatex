<div align="center">
    <h1>protoc-gen-validatex</h1>
</div>

<div align="center">

| [English](https://github.com/protoc-gen/protoc-gen-validatex) | [中文简体](README_zh-CN.md) |

</div>

---

`protoc-gen-validatex` 是一个用于 `protoc` 的插件，旨在简化和自动化常见参数校验规则的生成。通过使用参数选项（options）和扩展（extensions），用户可以轻松配置校验规则，并自动生成相应的校验器。该插件还支持与国际化（i18n）结合，允许根据不同语言环境生成多语言的错误信息，从而提升用户体验。


## 关键特性
- **自动化校验规则生成**：这也是 `protoc` 插件的一个基础必备特性了
- **多语言支持**: 天然支持多语言，快速适配各类项目和业务场景
- **MIT**: 所有使用者可以毫无心智负担的Fork并根据情况作修改调整

## 快速开始
可以参见[example](../example)的例子快速开始，这边展示一些片段

### 定义
```proto
syntax = "proto3";

option go_package = "github.com/protoc-gen/protoc-gen-validatex/example;main";

import "validatex/validatex.proto";

message SignInRequest {
  string email = 1 [(validatex.rules).string.email = true];
  string password = 2 [(validatex.rules).string = {min_len: 5, max_len: 50}];
}
```

### 生成
```bash
go install github.com/protoc-gen/protoc-gen-validatex
protoc --proto_path=. \
       --go_out=paths=source_relative:. \
       --validatex_out=paths=source_relative:. \
       --validatex_opt=i18n_dir=./example/i18n \
       ./example/*.proto
```

### 使用
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

## 许可
该项目采用MIT进行许可，详情请参阅[LICENSE](../LICENSE)文件。

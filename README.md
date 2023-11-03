# protoc-gen-go-fiber

`protoc-gen-go-fiber` 是一个用于生成基于 [Go](https://golang.org/) 和 [Fiber](https://gofiber.io/) 的 Web 服务的 Protocol Buffers (protobuf) 文件的插件。

## 安装

使用以下命令安装 `protoc-gen-go-fiber`：

```shell
go install github.com/gopkg-dev/protoc-gen-go-fiber@latest
```

## buf

在 `buf` 的配置文件 `buf.gen.yaml` 中，可以定义插件的生成行为。以下是一个示例 `buf.gen.yaml` 文件的内容：

```yaml
version: v1
managed:
  enabled: false
plugins:
  # 生成 Go 结构体代码
  - name: go
    out: gen
    opt: paths=source_relative
  # 生成 Fiber 服务代码
  - name: go-fiber
    out: gen
    opt: paths=source_relative
  # 生成 OpenAPI v3 代码
  - plugin: buf.build/community/google-gnostic-openapi
    out: .
    opt:
      - paths=source_relative
      - fq_schema_naming=true
```

在上述示例中，我们配置了生成 Go 结构体代码、Fiber 服务代码和 OpenAPI v3 代码的插件。

## 参数校验

`protoc-gen-go-fiber` 支持使用 `protovalidate-go` 库进行参数校验。你可以在 [bufbuild/protovalidate-go](https://github.com/bufbuild/protovalidate-go) 上找到该库的详细信息和用法示例。

在你的 Protobuf 文件中定义消息结构时，可以使用 `protovalidate-go` 提供的注释来添加参数校验规则。以下是一个示例：

```protobuf
syntax = "proto3";

import "buf/validate/validate.proto";

message CreateUserRequest {
  string name = 1 [(buf.validate.field).string = { min_len: 1, max_len: 100}];
  int32 age = 2 [(buf.validate.field).int32.gt = 0];
}
```

生成的代码将自动应用参数校验规则，你无需手动处理参数校验，只需关注自己的业务逻辑。

## 文档生成

你可以使用 [google/gnostic](https://github.com/google/gnostic) 插件来生成文档。该插件可以根据你的 Protobuf 文件生成 OpenAPI 文档。

## 示例

你可以在 [protoc-gen-go-fiber 示例项目](https://github.com/gopkg-dev/protoc-gen-go-fiber/tree/main/example) 中找到一个简单的示例，演示了如何使用 `protoc-gen-go-fiber` 生成和使用基于 Fiber 的 API。你可以根据自己的需求进行更复杂的定制和扩展。

感谢你提供的具体修改建议！以下是根据你的要求更新的文档内容：

## 贡献

欢迎贡献代码、报告问题或提供改进建议！我们非常欢迎开发者的参与和贡献。

- 如果你遇到了问题或发现了 bug，请在 GitHub 存储库上提交问题，详细描述你遇到的情况和期望的行为。我们会尽力解决问题并改进插件。
- 如果你有改进建议或想要添加新的功能，请通过拉取请求的方式提交你的代码。在提交拉取请求之前，请确保你的代码符合项目的编码风格和质量要求，并附上适当的测试。
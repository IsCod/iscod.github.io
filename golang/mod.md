# go.mod


### go mod命令

```bash
go mod download # 将模块下载到本地缓存，go env中定义的GOMODCACHE路径
go mod init # 初始化一个项目
go mod tidy # 依赖对齐，添加缺少的依赖，删除未使用的依赖
go mod verify # 验证依赖是否正确
```

### go install、get、download命令

```
go install # 安装可执行插件
go mod download gopkg.in/yaml.v1  # 仅下载依赖，不会下载依赖引用的依赖
go get gopkg.in/yaml.v1 # 获取模块信息并更新go.mod文件，如果本地有缓存则引用本地缓存，如果没有则下载
go clean -modcache # 清除临时目录中的文件
```

一个go.mod包含以下内容

```go.mod
// 模块名称
module project_name

// go sds版本
go 1.19

// 当前项目依赖的包, indirect表示间接依赖的包
require (
    // dependency latest
    golang.org/x/net v0.11.0 //indirect
)

// 当前项目排除的包
exclude (
    //dependency latest
)

//使用本地替换原始包
replace (
    //source latest => target latest
)

// 项目撤回的版本
// 当其他项目引用我们的包时，撤回出问题的版本
retract (
    v1.0.0
)
```

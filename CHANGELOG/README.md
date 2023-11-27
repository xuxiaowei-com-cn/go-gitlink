# 更新日志

## [语义化版本](https://semver.org/lang/zh-CN/)

- 版本格式：主版本号.次版本号.修订号，版本号递增规则如下：

1. 主版本号：当你做了不兼容的 API 修改，
2. 次版本号：当你做了向下兼容的功能性新增，
3. 修订号：当你做了向下兼容的问题修正。

先行版本号及版本编译信息可以加到“主版本号.次版本号.修订号”的后面，作为延伸。

# [更新日志](#更新日志)

- 采用倒叙编写
- 符号含义
    - 📗 Links | 链接
    - ⭐ New Features | 新功能
    - 🐞 Bug Fixes | 漏洞修补
    - 📔 Documentation | 文档
    - 🔨 Dependency Upgrades | 依赖项升级
    - ❤ Contributors | 贡献者
    - ⚠️ Noteworthy Changes | 值得注意的变化

## v0.2.0

- ⭐ New Features | 新功能
    - 新增 `GetTags`
      [获取仓库标签列表](https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-118749619)

- ⚠️ Noteworthy Changes | 值得注意的变化
    - 作者结构体 `Author` 属于公共部分，放入 `structs.go` 中，并新增 `Id` 属性

- 🐞 Bug Fixes | 漏洞修补
    - 修改 `TestDeleteTag`
      [删除一个标签](https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-118749620)
      测试方法中的参数，防止删除重要标签

## v0.1.0

- ⭐ New Features | 新功能
    - 新增 `TestDeleteTag`
      [删除一个标签](https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-118749620)

## v0.0.3

- 🐞 Bug Fixes | 漏洞修补
    - `TestPostAttachments`
      [上传文件](https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-128323479)
      测试类 修正日志

## v0.0.2

- 🐞 Bug Fixes | 漏洞修补
    - `GetProjects`
      [项目列表](https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-102299292)
      接口返回数据：新增 status、message

## v0.0.1

- ⭐ New Features | 新功能
    - 新增 `GetProjects`
      [项目列表](https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-102299292)
    - 新增 `PostAttachments`
      [上传文件](https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-128323479)

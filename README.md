# myget

`myget` 是一个简单的 winget 包管理器包装工具，提供导入、导出和下载功能。

## 功能

- `import` - 从 JSON 文件导入包列表并安装
- `export` - 导出当前已安装的包到 JSON 文件
- `download` - 根据包 ID 批量下载包

## 安装

```bash
go build .
```

## 使用方法

### import 命令

从 JSON 文件导入包列表并安装：

```bash
myget import -i <input_file> [-d]
```

参数说明：

- `-i, --input` (必需): 输入的 JSON 文件路径，包含要导入的包列表
- `-d, --download` (可选): 如果指定此参数，在导入和安装包后还会下载这些包的安装文件

示例：

```bash
# 仅导入和安装包
myget import -i packages.json

# 导入、安装并下载包
myget import -i packages.json -d
```

### export 命令

导出当前已安装的包到 JSON 文件：

```bash
myget export -e <export_path>
```

参数说明：

- `-e, --export` (必需): 导出的 JSON 文件路径

示例：

```bash
# 导出到指定文件
myget export -e exported_packages.json

# 导出到目录（如果没有以 .json 结尾，默认文件名为 export.json）
myget export -e /path/to/directory/
```

### download 命令

根据包 ID 列表批量下载包：

```bash
myget download -i <package_ids> [-d <directory>]
```

参数说明：

- `-i, --ids` (必需): 要下载的包 ID 列表，多个 ID 之间用逗号分隔
- `-d, --directory` (可选): 下载目录，默认为当前目录

示例：

```bash
# 下载单个包
myget download -i "Microsoft.VisualStudioCode"

# 下载多个包到默认目录
myget download -i "Microsoft.VisualStudioCode,Google.Chrome"

# 下载多个包到指定目录
myget download -i "Microsoft.VisualStudioCode,Google.Chrome" -d "/path/to/downloads/"
```

## 依赖

- [winget](https://learn.microsoft.com/en-us/windows/package-manager/) (Windows 包管理器)
- [Cobra](https://github.com/spf13/cobra) (Go CLI 库)

## 架构

- `core/import.go`: 处理包的导入和安装
- `core/export.go`: 处理包的导出
- `core/download.go`: 处理包的下载
- `main.go`: CLI 界面入口

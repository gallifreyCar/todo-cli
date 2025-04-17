# 📝 todo-cli

一个使用 Go 编写的跨平台命令行任务管理工具。  
支持添加任务、标记完成、设置优先级与提醒时间，并支持自动发布构建产物。

---

## 🚀 功能特性

- ✅ 添加任务（支持设置提醒时间）
- ✅ 列出所有任务（显示状态 / 优先级 / 提醒时间）
- ✅ 标记任务完成
- ✅ 删除任务
- ✅ 修改任务内容
- ✅ 设置任务优先级（low / medium / high）
- ✅ 本地提醒调度器（`todo remind`）
- ✅ 自动发布版本并上传二进制文件（via GitHub Actions + GoReleaser）

---

## 🛠 安装方式

### 方式一：下载预编译二进制（推荐）

访问 [Releases 页面](https://github.com/gallifreyCar/todo-cli/releases)，下载适合你系统的版本：

```bash
# 例如在 Linux 下
wget https://github.com/gallifreyCar/todo-cli/releases/download/v1.1.0/todo_1.1.0_linux_amd64.zip
unzip todo_1.1.0_linux_amd64.zip
chmod +x todo
./todo
```

### 方式二：本地编译

```bash
git clone https://github.com/gallifreyCar/todo-cli.git
cd todo-cli
go build -o todo
./todo
```

---

## 📦 使用示例

### 添加任务

```bash
todo add "写日报"
todo add "开会" --remind "2025-04-17 15:00"
```

### 查看任务列表

```bash
todo list
```

### 标记任务完成

```bash
todo done 2
```

### 设置优先级

```bash
todo priority 2 high
```

### 修改任务内容

```bash
todo edit 2 "改为写周报"
```

### 删除任务

```bash
todo delete 2
```

### 启动提醒监听器（定时提醒）

```bash
todo remind
```

---

## ⚙️ 项目结构

```
.
├── cmd/                # 所有子命令
│   ├── add.go
│   ├── list.go
│   ├── done.go
│   ├── remind.go
│   └── ...
├── main.go             # CLI 入口
├── tasks.json          # 本地任务存储文件
├── .goreleaser.yaml    # 构建发布配置
├── .github/workflows/  # 自动发布 CI 配置
└── README.md
```

---

## 🧩 未来计划

- 支持任务分类 / 标签
- 支持任务搜索 / 过滤
- 添加远程同步（如 GitHub gist、Firebase）
- Web UI / TUI 图形界面

---

## 📄 License

MIT © [gallifreyCar](https://github.com/gallifreyCar)

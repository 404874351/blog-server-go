# blog-server-go

### 介绍
个人博客-服务端-go版本

blog-server-go是个人博客-服务端-Java版本的重构项目，基本还原了原本项目的业务功能，如安全控制、博客撰写、用户评论等。

本项目逻辑简单，没有太多花里胡哨的功能，适合初学者快速上手。

相比之下，由于不存在SpringBoot中的过度封装和完全面向对象等问题，go版本的业务代码显得更加轻量，更加简洁，中间件或过滤器的控制也更加灵活。

但是，在参数解析和对象自由调用等方面，go却比Java更加繁琐。同时go的开发生态也正处于起步阶段，无法像Java那样遇到啥需求很容易就能找个库开箱即用。总之，go和java各有利弊吧，结合项目需求自由选择即可。


### 主要功能

- 用户管理：系统管理员可分配用户角色
- 角色管理：基于角色为用户分配菜单和资源
- 菜单管理：基于角色动态展示后台管理页面
- 权限控制：基于Jwt和Casbin实现的权限管理
- 多方式登录：支持传统密码登录和短信验证码登录
- 博客业务：提供博客、分类、标签、评论、留言等核心功能

### 技术选型

- Golang版本：1.17 或 >= 1.16
- Web框架：Gin
- 持久层框架：Gorm
- 数据库：Mysql 8
- 缓存：Redis 6
- 安全控制：JWT身份验证，Casbin RBAC权限控制
- 文件存储：七牛云对象存储
- 短信服务：阿里云SMS
- 系统配置：Viper，配置信息存放在yaml文件
- 日志输出：Zap
- IDE推荐：JetBrains Goland

### 目录结构

```bash
server
├── api             # 接口层，提供api处理函数
├── conf            # 配置文件
├── docs            # 相关文档
├── middleware      # 中间件层
├── model           # 模型层，与数据库交互
    ├── request     # 接收参数相关结构体
    └── response    # 响应数据相关结构体
├── router          # 路由层，配置api及其对应的处理函数
├── service         # 业务层
├── static          # 静态文件（本项目为空）
├── utils           # 常用工具包
├── main.go         # 项目入口文件
```

### 安装使用

```bash
# 克隆项目
git clone https://github.com/404874351/blog-server-go.git
# 进入工程目录
cd blog-server-go
# 将conf.example.yaml复制一份到conf.yaml，作为正式的配置文件
cp ./conf/conf.example.yaml ./conf/conf.yaml
# 修改 ./conf/conf.yaml，将配置信息改为自己的
vim ./conf/conf.yaml
# 在Mysql中执行 ./docs/model.sql，初始化数据库信息
# 可使用Navicat管理Mysql，并导入sql文件，过程省略
mysql xxx
# 使用 go mod 下载并引入依赖
go generate
# 启动项目，或在Goland内启动
go run main.go
```

### 部署上线

```bash
# 以下是Linux下部署过程，windows部署方法也差不多，请自行查阅

# Windows下打包，请先修改全局变量，Linux下则忽略
set GOARCH=amd64
go env -w GOOS=linux
# 编译，完成后复制文件到Linux下任意目录
go build -o server main.go
# 进入该目录，并复制项目中conf目录到当前目录
cd <your-working-dir>
cp -r blog-server-go/conf ./conf
# 查看目录文件，此时该目录下至少应有 server文件 和 conf目录
ls
# 修改执行权限
chmod 777 ./server
# 前台启动，如没有复制conf目录，将报错：找不到配置文件
./server
# 后台启动，日志输出到文件
nohup ./server >out.log 2>&1 &
```

### 学习文档

1. Gin-Vue-Admin https://www.gin-vue-admin.com/
2. Gin https://learnku.com/docs/gin-gonic/1.7
3. Gorm https://learnku.com/docs/gorm/v1
4. JWT https://jwt.io/
5. Casbin https://casbin.org/

### 后续改进

1. 如业务条件允许，可选用原生SQL语句或SQL视图，以简化项目中的部分Gorm操作。
2. 本项目使用VO统一接收参数，使用DTO统一响应数据，本质上是Java风格的格式化操作。在较为简单的业务需求下，可以删除VO和DTO，直接使用Model参与业务即可。
3. 如依然存在问题，可联系作者。 QQ：404874351

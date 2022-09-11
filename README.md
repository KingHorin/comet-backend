# comet-backend

## 环境要求
Go 1.16, 安装并启动Mysql服务


## 部署步骤
1.终端进入comet-server, 输入go mod download, 等待依赖下载完毕

2.在pkg/database.go中对应保存有连接至Mysql的相关信息, 使其与本地数据库配置相同

3.终端输入go build，构建项目文件

4.终端输入go run main.go，后端开始工作

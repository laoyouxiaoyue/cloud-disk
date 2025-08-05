使用的命令
//新建api
goctl api new xxx
// 启动服务
go run xxx.go -f etc/xxx.yaml
// 根据api生成代码
goctl api go -api core.api -dir . -style go_zero
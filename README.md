## gtask
Go 实现任务协同系统

## 目录规范
- cmd: 可执行文件,可能有多个main文件
- internal: 内部代码,不希望被外部访问
- pkg: 公开代码
- config: 配置文件
- script: 可执行脚本
- docs:  文档
- third_party: 第三方工具
- bin: 编译打包后的二进制文件
- deploy: 部署相关
- test: 测试文件
- api: 开放的api接口

## 技术选项
- gin: Web服务组件
- gorm: 数据访问组件
- logrus: 日志组件
- redis: 缓存
- ant: 协程池
- grpc: 服务通信
- nacos: 配置中心
- viper: 配置读取组件

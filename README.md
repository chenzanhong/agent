服务器监控管理平台项目的监控代理。
运行：
```bash
  cd agent
  ./main -hostname=xxxx -token=xxxx # hostname是在服务器监控管理平台上添加该服务器时填写的主机名，token是添加该服务器时后台生成的唯一标识符，不过目前好像没实际使用价值，后面打算去除这个token
```

配置文件 agent/config/config.yaml:
```yaml
url: 113.44.170.52 # 服务器监控管理平台所在服务器
port: 8080 # 发送采集到达数据到服务器监控管理平台相关接口的请求端口号
second: 10 # 采集间隔
```yaml

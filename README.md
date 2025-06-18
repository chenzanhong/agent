# agent

## 介绍
服务器监控管理平台的客户端代理，主要功能是采集服务器的数据，并通过https请求发送到服务器监控管理平台，用于前台的数据可视化。


## 使用说明
1. 克隆本仓库
```bash
git clone https://gitee.com/chenzanhong/agent
```bash

2. 进入agent目录，对编译好的二进制文件赋予执行权限，并执行
```bash
cd agent
chmod +x main
./main
```bash

3. （可选）修改上报采集数据的时间间隔，默认30秒
```bash
cd agent/config
vim config.yaml
```bash
config.yaml中的second字段即上报采集数据的时间间隔，单位“秒”
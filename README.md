# Go SOCKS5 Proxy

一个简单的基于Go的SOCKS5代理服务器，可以部署在Render上。

## 功能特点

- 支持SOCKS5协议
- 可选的用户名/密码认证
- Docker容器化
- 易于部署到Render

## 部署到Render

1. Fork这个仓库到你的GitHub账户

2. 在Render.com创建一个新的Web Service，并连接到你的GitHub仓库

3. 设置环境变量（可选）：
   - `SOCKS5_USERNAME`: 代理认证用户名
   - `SOCKS5_PASSWORD`: 代理认证密码
   - `PORT`: 服务端口（Render会自动设置）

4. 点击 "Create Web Service"

## 本地运行

```bash
# 构建Docker镜像
docker build -t go-socks5-proxy .

# 运行容器
docker run -p 1080:1080 \
  -e SOCKS5_USERNAME=your_username \
  -e SOCKS5_PASSWORD=your_password \
  go-socks5-proxy
```

## 使用方法

设置你的SOCKS5客户端，连接到服务器：
- 主机：你的Render应用URL
- 端口：80/443（Render自动分配）
- 如果设置了认证，需要提供用户名和密码

## 注意事项

- Render免费版有一些限制，包括带宽和计算资源
- 建议在生产环境中启用认证
- 使用HTTPS以保护数据传输安全

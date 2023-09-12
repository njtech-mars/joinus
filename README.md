# Mars 工作室报名表单 ![docker](https://github.com/njtech-mars/joinus/actions/workflows/docker.yml/badge.svg)

本项目采用 monorepo 的形式，将前后端放在一个 git 仓库当中。

后端是 **GO**，前端是 **Svelte**，分别放在了 `server` 和 `client` 两个目录中，最后可以编译成一个很小的 Docker 镜像。

## 1. 环境变量

本项目用到了以下几个环境变量，其中邮箱用的是 TSL 通道，端口应该是 465：

| 变量名     | 解释                     | 备注              |
| :--------- | :----------------------- | :---------------- |
| PORT       | 项目监听端口             | 可选，默认为 5000 |
| CORS       | 允许 Cors 的地址         | 可选，默认为 空   |
| ADMIN_PASS | 用来登录后台的管理员密码 | 必需              |
| SMTP_USER  | SMTP 用户名              | 必需              |
| SMTP_PASS  | SMTP 密码                | 必需              |
| SMTP_HOST  | SMTP 主机地址            | 必需              |
| SMTP_PORT  | SMTP 主机端口            | 必需              |

## 2. 部署项目

新建一个 docker-compose.yaml 文件，根据需要修改其中的环境变量：

```yaml
version: "3"
services:
  mars-joinus:
    image: mraddict063/mars-joinus
    restart: unless-stopped
    ports:
      - 5000:5000
    environment:
      - ADMIN_PASS=admin_password
      - SMTP_PASS=smtp_password
      - SMTP_USER=smtp_username
      - SMTP_HOST=smtp_host
      - SMTP_PORT=smtp_port
    volumes:
      - ./data:/data
```

然后新建一个 data 文件夹，在其中放入一个用来发送邮件的模板文件，文件名为 `template.html`，下面是一个例子：

```html
<h1>{{.Name}}同学你好</h1>
<p>欢迎你报名Mars工作室</p>
```

最后启动项目即可：

```sh
docker-compose up -d
```

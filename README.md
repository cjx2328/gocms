<h1>GOCMS</h1>
引用自 
<div>
 基于 Gin + GORM + Casbin + vue-element-admin 实现的权限管理系统 <br/>
 基于Casbin 实现RBAC权限管理 <br/>
 前端实现： vue-element-admin <br/>
  <br/>
</div>
<br/>

## 特性

- 基于 Casbin 的 RBAC 访问控制模型
- JWT 认证
- 前后端分离

## 下载并运行

### 获取代码

```
go get -v github.com/cjx2328/gocms
```

### 运行

 
- 运行服务端：cd cmd/manageweb，go run main.go，运行成功后打开 127.0.0.1:8080，如果是在windows下操作，需要提前安装并配置好mingw（sqlite的操作库用到），安装方式请自行百度/谷歌。
- 调试/运行web：cd website/manageweb，安装：npm install，运行：npm run dev，打包：npm run build:prod
- 配置文件在(`cmd/manageweb/config.yaml`)中，用户默认为：admin/123456


#### 温馨提醒

1. 默认配置采用的是 sqlite 数据库，数据库文件(`自动生成`)在`cmd/manageweb/data/goapp.db`。如果想切换为`mysql`或`postgres`，请更改配置文件，并创建数据库（表会自动创建）。
2. 日志的配置为标准输出并写入文件。

## 前端实现

- website/manageweb：基于[vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)的实现版本

## 项目结构概览

<details>
<summary>展开查看</summary>
<pre><code>.
├── cmd  项目的主要应用
├── internal  私有应用程序和库代码
├── pkg  外部应用程序可以使用的库代码
├── vendor  项目依赖的其他第三方库
├── website  vue-element-admin
</code></pre>
</details>


 
 

## 相关文章

- [如何使用goapp写你的后台管理系统] - [https://www.cnblogs.com/hotion/p/11665837.html/](https://www.cnblogs.com/hotion/p/11665837.html/)

## 感谢以下框架的开源支持

- [Gin] - [https://gin-gonic.com/](https://gin-gonic.com/)
- [GORM] - [http://gorm.io/](http://gorm.io/)
- [Casbin] - [https://casbin.org/](https://casbin.org/)
- [vue-element-admin] - [https://github.com/PanJiaChen/vue-element-admin/](https://github.com/PanJiaChen/vue-element-admin/)
- [goapp] - [https://github.com/cjx2328/gocms/](https://github.com/cjx2328/gocms/)


## MIT License

    Copyright (c) 2020 

 





# GOblog

一个 go语言实现的博客
技术栈
golang
vue
echo 框架
前台界面目前修改于 https://firekylin.org/
开发原因
现在自己使用博客系统不方便管理、部署和二次开发
渲染速度到不到想要的效果
计划功能
 文章管理
 页面管理
 分类管理
 标签管理
 系统设置
 管理主页
 评论
 待计划
使用方式
git clone https://github.com/zxysilent/blog.git
主要若要部署，请修改关键信息[jwt、用户名、密码等] 数据库自带文章内容请在部署前删除

运行服务
配置conf.json数据库信息
还原 data目录下 z_blog.sql 数据库 数据库自带文章内容请在部署前删除
安装依赖
go run
运行后台
安装依赖 npm install
开发运行 npm run serve
浏览器打开 http://127.0.0.1:8080/
发布 npm run build 会自动发布到 res/dist目录下面
友链里面第一个为后台登陆地址默认用户名密码已经填写，可自行数据库修改
## Blog Kratos
kratos 参考文档：https://go-kratos.dev/

gorm 参考文档：https://gorm.io/zh_CN/docs/

### 相关命令
```shell
# 用户模块
# 创建user.proto
kratos proto add api/blog/user.proto 
# 创建PB
kratos proto client api/blog/user.proto
# 生成service
kratos proto server api/blog/user.proto t internal/service

# 文章模块
# 创建article.proto
kratos proto add api/blog/article.proto 
# 创建PB
kratos proto client api/blog/article.proto
# 生成service
kratos proto server api/blog/article.proto t internal/service
# 运行
kratos run
```

### 系统模块
- [x] 用户管理
  - [x] 登录
  - [x] 注册
- [ ] 文件上传
- [ ] 文章管理
- [ ] 字典管理
- [ ] 主题管理
- [ ] 友链设置
- [ ] 角色授权

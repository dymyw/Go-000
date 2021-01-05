## 学习笔记

按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架

### 目录结构

```
- project_layout
    - api                   应用接口层
        - account
            - v1
                - account.pb.go
    - cmd                   可执行文件目录
        - account_admin         运营管理
        - account-interface     BFF
        - account-job           流式任务处理
        - account-service       对内微服务
        - account-task          定时任务
    - configs               配置文件
    - internal              私有应用程序和库代码
        - biz                   业务层
        - data                  数据访问层
        - pkg                   内部可复用库代码
        - service               服务层
    - test                  测试数据
    - go.mod
    - README.md
```

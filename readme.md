## 基于gin,gorm,mysql对需求的简单实现
仓库地址 [github.com/Mutezebra/byteDance_require](https://github.com/Mutezebra/byteDance_require)  
其实课上讲师给出的需求相对简单,只需要简单的实现即可，但是我以前有过一定的项目经历， 
于是就本着复习的想法，以一个完整项目的格式来实现了需求，以下是我的项目目录

```
ALittleRequire/
├─api  
├─config  
│  └─local  
├─pkg  
│  ├─ctl  
│  ├─e  
│  └─utils  
├─repository  
│  ├─dao  
│  └─model  
├─routes  
├─service
└─types  
```
- api : 用于定义接口函数,也就是controller层
- conf : 用于存储配置文件
- pkg/e : 封装错误码
- pkg/logger : 日志打印
- repository: 仓库放置所有存储
- repository/db: 持久层MySQL仓库
- repository/db/dao: 对db进行操作的dao层
- repository/db/model: 定义所有持久层数据库表结构的model层
- routes : 路由逻辑处理
- service : 接口函数的实现
- types : 放置所有的定义的结构体

### 项目开发
1. 首先是进行了配置文件的书写,先完善config文件夹的基本信息,编写本地配置
2. 然后通过刚刚绑定的信息在repository层dao目录的init.go文件中进行mysql的配置 (不包含建库操作)
3. 接着在repository层model目录的task.go文件定义要创建到数据库中的表结构
4. 再在repository层dao目录的migrate.go文件中设置自动迁移,并在init.go的末尾进行
引用,以实现刚刚定义的结构体迁移到数据库中成表
5. 此时已经可以将上述所进行的配置函数写进main.go函数中尝试运行一遍,来判断上述操作的成功与否
6. 接着可以在pkg包里使用第三方日志管理包来配置logger,这一步非必需,可以使用log包代替
7. 然后就可以逐步的开始业务逻辑处理,先在pkg/ctl中进行基本Response的定义,然后在types中
进行Request的定义，以及一些需求的Response定义
8. 定义完简单的request后可以在routes中配置一个路由，并在main.go中使用
9. 接着就可以在routes中配置路由，在api中给到Handel,在根据需求在service中处理逻辑
根据情况在repository/dao中操作数据库
10. 每一步的具体实现都可以看源码哦,这样一步一步的就全部弄好啦

### 额外说明
1. pkg/e 主要是用来处理一些自定义的错误码，可以在程序出现问题时通过日志来更精确的定位到错误
2. 受限于时间和个人水平,项目开发描述和代码本身质量难免会有不足,欢迎大家指正
3. 仓库链接见文章顶部
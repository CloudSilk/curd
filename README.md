## 项目名称
> 请介绍一下你的项目吧  



## 运行条件
> 列出运行该项目所必须的条件和相关依赖  
* 条件一
* 条件二
* 条件三



## 运行说明
> 说明如何运行和使用你的项目，建议给出具体的步骤说明
* 添加配置文件
```yaml
dubbo:
  config-center:
    protocol: nacos
    address: 127.0.0.1:8848
    data-id: "curd"
    group: basic
    namespace: CloudSilk
```

```yaml
dubbo:
  config-center:
    protocol: nacos
    address: 127.0.0.1:8848
    data-id: "curd"
    params:
      mysql: "root:123456@(127.0.0.1:3306)/curd?charset=utf8mb4&parseTime=True&loc=Local"
      debug: "true"
  registries:
    nacos:
      protocol: nacos
      timeout: 3s
      address: 127.0.0.1:8848
  protocols:
    triple:
      name: tri
      port: 20005
  provider:
    registry-ids: nacos
    services:
      PageProvider:
        protocol-ids: triple
        interface: org.CloudSilk.Page
      MetadataProvider:
        protocol-ids: triple
        interface: org.CloudSilk.Metadata
  consumer:
    references:
      IdentityClientImpl:
        protocol: tri
        interface: org.CloudSilk.Identity
```
* 操作二
* 操作三  



## 测试说明
> 如果有测试相关内容需要说明，请填写在这里  



## 技术架构
> 使用的技术框架或系统架构图等相关说明，请填写在这里  


## 协作者
> 高效的协作会激发无尽的创造力，将他们的名字记录在这里吧

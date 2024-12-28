# AnywhereDoorControlPlane

## AnywhereDoor控制平面
* 用于与Plugin和Imsdk通讯

## IO接口

### imsdk
* 输入:
  * url: /imsdk
  * method: post
  * header:
    * username: 所属用户的用户名
    * token: imsdk的token
  * param:
    * name: imsdk名称
    * target: 目标plugin名
    * data: 原始数据
  * ret:
    * code: 详见错误码枚举
    * message: 详见错误信息枚举
* 输出
  * url: /imsdk
  * method: post
  * header:
    * token: imsdk的token
  * param:
    * name: 请求来源的plugin名称
    * target: 目标imsdk名(即当前imsdk的名字, 用来校验使用)
    * data: 原始数据

### plugin
* 输入:
    * url: /plugin
    * method: post
    * header:
        * username: 所属用户的用户名
        * token: plugin的token
    * param:
        * name: plugin名称
        * target: 目标imsdk名
        * data: 原始数据
    * ret:
      * code: 详见错误码枚举
      * message: 详见错误信息枚举
* 输出
    * url: /plugin
    * method: post
    * header:
        * token: plugin的token
    * param:
        * name: 请求来源的imsdk名称
        * target: 目标plugin名(即当前plugin的名字, 用来校验使用)
        * data: 原始数据

### config
* imsdk配置
  * url: /imsdk/config
  * method: post
  * header:
    * username: 所属用户的用户名
    * token: imsdk的token
  * param:
    * name: imsdk的名称
    * config_key: 配置键
  * ret:
      * code: 详见错误码枚举
      * message: 详见错误信息枚举
      * data: 配置数据

* plugin配置
    * url: /plugin/config
    * method: post
    * header:
        * username: 所属用户的用户名
        * token: plugin的token
    * param:
        * name: plugin的名称
        * config_key: 配置键
    * ret:
        * code: 详见错误码枚举
        * message: 详见错误信息枚举
        * data: 配置数据


## 环境变量
* DB_IP: 数据库IP地址
* DB_PORT: 数据库端口
* DB_NAME: 数据库名字, 即`anywhere_door`, 第一步中创建的数据库名字, 可以更换名字, 不建议
* DB_USER: 数据库用户
* DB_PASSWORD: 数据库密码


## 打包方式
1. 将代码clone下来
2. 安装docker及buildx
3. 打包镜像:
   * `docker buildx build --platform linux/amd64 -t 192.168.25.5:31100/maoyanluo/anywhere-door-control-plane:1.0 . --load`

## 部署方式

### Docker Command Line
1. 创建容器:
   * `docker run --name anywhere-door-control-plane -itd -p 8081:80 -e DB_IP=ip -e DB_PORT=port -e DB_NAME=anywhere_door -e DB_USER=user -e DB_PASSWORD=pwd --restart=always 192.168.25.5:31100/maoyanluo/anywhere-door-control-plane:1.0`


### Kubernetes
```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: anywhere-door-control-plane-deployment
  namespace: anywhere-door
spec:
  replicas: 1
  selector:
    matchLabels:
      app: anywhere-door-control-plane
  template:
    metadata:
      labels:
        app: anywhere-door-control-plane
    spec:
      containers:
      - name: anywhere-door-control-plane
        image: 192.168.25.5:31100/maoyanluo/anywhere-door-control-plane:1.0
        imagePullPolicy: IfNotPresent
        env:
        - name: DB_IP
          value: "anywhere-door-mysql-service.anywhere-door"
        - name: DB_PORT
          value: "3306"
        - name: DB_NAME
          value: "anywhere_door"
        - name: DB_USER
          value: "user"
        - name: DB_PASSWORD
          value: "pwd"
        ports:
        - containerPort: 80
      restartPolicy: Always
---
apiVersion: v1
kind: Service
metadata:
  name: anywhere-door-control-plane-service
  namespace: anywhere-door
  labels:
    app: anywhere-door-control-plane
spec:
  type: ClusterIP
  ports:
  - port: 80
    targetPort: 80
  selector:
    app: anywhere-door-control-plane
```

## 使用
1. 无需任何操作, 保证容器正常启动即可

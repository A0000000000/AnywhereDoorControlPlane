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
    * code: 错误码 200为无错误
    * message: 错误信息
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
        * code: 错误码 200为无错误
        * message: 错误信息
* 输出
    * url: /plugin
    * method: post
    * header:
        * token: plugin的token
    * param:
        * name: 请求来源的imsdk名称
        * target: 目标plugin名(即当前plugin的名字, 用来校验使用)
        * data: 原始数据

## 部署方式
1. 将代码clone下来
2. 安装docker及buildx
3. 打包镜像:
   * `docker buildx build --platform linux/amd64 -t 192.168.25.5:31100/maoyanluo/anywhere-door-control-plane:1.0 . --load`
4. 创建容器:
   * `docker run --name anywhere-door-control-plane -itd -p 8081:80 -e DB_IP=192.168.25.7 -e DB_PORT=3306 -e DB_NAME=anywhere_door -e DB_USER=root -e DB_PASSWORD=09251205 192.168.25.5:31100/maoyanluo/anywhere-door-control-plane:1.0`
# 望子成龙小学聊天室

## 项目介绍
望子成龙小学聊天室是一个基于vue3和go语言开发的聊天室项目，将前后端都部署在客户机，服务器只需部署数据库即可。实现了最简单的用户之间的聊天和消息推送、文件发送与接收功能。

## 项目特点
- 基于vue3和go语言开发，实现了最简单的用户之间的聊天和消息推送、文件发送与接收功能。
- 数据库使用redis。
- 服务端只提供数据中转，无需部署后端程序即可使用。
- 项目客户端程序已经编译打包完成，可以直接运行。
- 客户端程序使用webview2进行渲染，目前只能在windows10 1902及以上版本的系统下运行。

> ！ 项目还在开发过程中，诸多问题待解决... ！
> 后续会测试将本地部署的DeepSeek-R1接入聊天室，作为ai对话页面

## 项目部署
1. 下载项目代码
```bash
git clone https://gitee.com/Hu_BanXian/chatroom
```
2. 在服务器安装并配置redis数据库
> 安装教程请自行搜索...

将redis数据库配置文件redis.conf中的以下内容修改为以下内容：
```bash
bind 192.168.0.0      # 服务器ipv4地址，请自行查看并修改，并修改项目配置文件chatRoom.conf.json中的RedisAddr项
port 6379             # 服务器端口号，如需修改请一起修改项目配置文件chatRoom.conf.json中的RedisAddr项
maxclients 10000      # 最大连接数，请根据实际情况修改（一名用户会同时建立两个连接，用户上传文件时会额外建立一个连接，上传文件结束后关闭）
databases 20          # 数据库数量，默认聊天数据库为0，此外每有用户上传一个文件（图片除外）就会新建一个数据库，若数据库数量已满会导致文件上传失败，请根据实际情况修改
maxmemory 1048576000  # 最大内存请根据服务器实际情况修改
requirepass 123456    # 服务器密码，请根据实际情况修改，并修改项目配置文件chatRoom.conf.json中的RedisPassword项
```

3. 启动redis数据库
- 启动数据库命令请自行搜索...
- 请注意，redis数据库关闭后会清除所有数据，关闭服务器请提前备份好数据。

## 项目配置文件
项目配置文件位于项目根目录的gin/chatRoom.conf.json，包含以下内容：
```json
{
  "RedisPassword": "123456",    # 数据库密码
  "RedisAddr": "0.0.0.0:6379",  # 数据库地址及端口
  "RedisDB": 0,                 # 存储聊天记录的数据库
  "GinPort": ":12306",          # 客户端运行时占用的端口号
  "ListenerLastLen": 0,         # 最后接收到的消息的id，防止重复接收消息。此项在程序运行过程中会自动修改，请勿手动修改。
  "NumOfConcurrentMsg": 10      # 消息并发数量，决定了客户端在一瞬间最多可以接收多少条消息，不了解go语言的管道功能的请不要修改此项。
}
```

用户配置文件位于项目根目录的gin/user.conf.json，包含以下内容：
```json
{
  "Id": "000",                # 用户id，此项不能为空，内容和长度可以自定义，但是必须唯一（请部署者自行与用户沟通，或修改项目来防止用户修改）。
  "Name": "迷糊老师",          # 用户昵称，此项不能为空
  "Level": 10,                # 用户等级，可以为空
  "Avatar": "",               # 用户头像，可以为空  ！此功能还未完成，请勿填写此项
  "Title": "高贵的群主",        # 用户头衔，可以为空
  "TitleColor": "#f5e50dbc",  # 用户头衔背景颜色，可以为空
  "Phone": ""                 # 用户手机号，可以为空
}
```

数据文件位于项目根目录的gin/data.json，用于存储聊天记录，如需清理聊天记录请删除此文件，如需找回删除的聊天记录请将配置文件中的ListenerLastLen项修改为0。
### demo
[server](http://docmaker-server.randqun.com/)

[client](http://docmaker.randqun.com/)

#### 安装
编辑backend下面的config.json文件

>remote_server_host 使用自己搭建好的  [docMakerServer](https://github.com/zhurenkai/docMakerServer)
>
>或者使用默认的 http://docmaker-server.randqun.com 此域名只为调试使用不保证长期有效
>
>client_id和client_secret也是使用docMakerServer中生成的
>
>默认的为client_id:2, client_secret:"uBgMv7eBCmQLoOWnUxmqFNO2UNQADBSpx4SEJhNS"



```
cd backend/
./server
```

#### 开发模式

```
cd frontend/
npm run dev
```

>开发模式的域名配置在frontend/config/index.js设置proxyTable


#### 使用
1. 进入工作界面,创建项目模块，接口
2. 选择method,host,uri ,填写请求参数（支持formdata和json），填写header 选择send
3. 不出意外会返回你想要的结果，目前支持json格式化
4.确认接口正常返回之后，可以选择保存，会保存所选择的内容（可选）
5.点击查看文档
6. 选择生成，此时会生成文档的半成品，如果有些字段你曾经使用过，可能会自动填充上说明
7. 选择保存文档，此时会记录你填写的字段
8. 在弹出框底部可选菜单选择markdown，进入markdown文档编辑界面，编辑完成后点击保存
9. 登录后台即可以看见文档(自己安装的后端地址，或者默认的http://docmaker-server.randqun.com)

#### 设置

+ 进入设置界面可以为每个项目设置hosts，在此后的使用中可以在host的下拉菜单中选择

+ 可以导入本地数据库字段的注释，可以直接根据自己看注释匹配返回字段说明
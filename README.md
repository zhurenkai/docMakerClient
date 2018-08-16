####使用
编辑backend下面的config.json文件

>remote_server_host 使用自己搭建好的docMakerServer域名
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

####开发模式

```

cd frontend/
npm run dev

```

>开发模式的域名配置在frontend/config/index.js设置proxyTable

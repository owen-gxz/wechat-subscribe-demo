### 前提条件： 需要有服务器和域名（已备案），并且申请对应的订阅消息： [订阅文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/subscribe-message.html)


### app.json 配置文件使用方式

###  port： 服务对外暴露端口

### every_time： 定时任务，每天几点执行

### wechat： 微信相关配置
  * app_id: 微信appid
  * app_secret： 微信app_secret
  * temps： 模板列别
  
 
  
  |字段|作用|使用方法|
  |id|模板id|微信公众平台申请 |
  |page|跳转的小程序页面||
  |values|模板对应的字段和回复内容||

values设置：
例如模板格式为：
```
姓名: {{name01.DATA}}
金额: {{amount01.DATA}}
行程: {{thing01.DATA}}
日期: {{date01.DATA}}
```

对应微信配置为：

```
"wechat": {
    "app_id": "xxxx",
    "app_secret": "xxxxx",
    "temps": [
      {
        "id": "xxxxx",
        "page": "/pages/main/main",
        "values": [
          "name01": "张三",
          "amount01": "3块",
          "thing01": "去北京",
          "date01": "2020-10-10 xxx"
        ]
      }
    ]
  }
```


提供了两个接口，为了方便全部为get请求：
1. code2session:获取open_id接口： 用户进入系统使用 [wx.login](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/login/wx.login.html)获取code，将code发送给后台，
2. subscribe?open_id=xxx: 用户点击按钮将用户open_id发送给后台



# typora-qiniu-uploader

[English](Readme-en.md)

## 功能简介

- 基本功能：将图片等文件上传到七牛云，并返回访问链接
- 进阶功能：在[typora](https://typora.io/) 编辑器中无缝调用，自动将markdown中的本地图片转为七牛云中的图片
- [演示](https://cdn.wubw.fun/typora/210508-162521-tqu-demo.gif)

## 使用方法

1. 在release中下载编译好的二进制文件，直接运行一次
2. 首次运行会提示无配置文件，并自动将配置模板写入`$HOME/.config/typora-qiniu-uploader/config.json`
3. 编辑配置文件，示例如下
   ```json
       {
        "access_key": "登陆七牛后，从密钥管理获取",
        "secret_key": "登陆七牛后，从密钥管理获取",
        "Bucket": "存储空间名称",
        "use_https": true,
        "use_cdn_domains": true,
        "domain": "空间绑定的域名，形如https://xxxx.com.cn",
        "sub_dir": "空间里的目录：如typora"
        }
   ```
4. 配置typora：偏好设置-图像
5. 日志路径：`$HOME/.config/typora-qiniu-uploader/tqu.log`

## 注意事项

- 目前仅在mac 11.3 & typora 0.10.8 (5313)版本得到了功能验证，其他操作系统及typora版本的可用性尚未经过测试

## 参考&致谢

[qiniuyun_upload_tools](https://github.com/Han-Ya-Jun/qiniuyun_upload_tools)
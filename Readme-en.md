# typora-qiniu-uploader

[中文](Readme.md)

## Features Introduction

- Basic Feature: Upload files to Qiniu Cloud and get urls immediately
- Advanced Feature: Integrate with a super cool markdown editor and transform local images into remote ones immediately
- [Demo](https://qncdn.wubowen.com.cn/typora/210508-162521-tqu-demo.gif)

## Usage

1. Download prebuilt binary from release page.
2. Run once.Normally a first-time run will get a hint that no config file found,after which the empty config file will
   be written to `$HOME/typora-qiniu-uploader/config.json`
3. Edit the config file, example:
   ```json
       {
        "access_key": "",
        "secret_key": "",
        "Bucket": "",
        "use_https": true,
        "use_cdn_domains": true,
        "domain": "https://xxxx.com.cn",
        "sub_dir": "typora"
        }
   ```
4. Change your typora setting to use this program as its image uploader: settings-image
5. Local log located at：`$HOME/typora-qiniu-uploader/tqu.log`

## Attention

- Availability only tested on mac 11.3 and typora 0.10.8 (5313), other platforms or typora versions not fully tested, use at your own risks.

## References

[qiniuyun_upload_tools](https://github.com/Han-Ya-Jun/qiniuyun_upload_tools)
package pkg

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"path"
	"time"
)

type QiNiuClient struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	//Zone          int    `json:"zone"` //0:华东, 1:华北, 2:华南, 3:北美
	UseHTTPS      bool   `json:"use_https"`
	UseCdnDomains bool   `json:"use_cdn_domains"`
	Domain        string `json:"domain"`
}

func NewQiNiuClient(accessKey, secretKey, bucket string, useHttps, useCdnDomains bool, domain string) *QiNiuClient {
	return &QiNiuClient{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Bucket:    bucket,
		//Zone:          0,
		UseHTTPS:      useHttps,
		UseCdnDomains: useCdnDomains,
		Domain:        domain,
	}
}

func (q *QiNiuClient) UploadImages(images []string) (urls []string) {

	putPolicy := storage.PutPolicy{Scope: q.Bucket}
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	token := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.UseHTTPS = q.UseHTTPS
	cfg.UseCdnDomains = q.UseCdnDomains

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	for _, image := range images {
		_, fileName := path.Split(image)
		key := time.Now().Format("060102-150405") + "-" + fileName
		if err := formUploader.PutFile(context.Background(), ret, token, key, image, nil); err != nil {
			log.Fatalln(err)
		}
		urls = append(urls, q.Domain+key)
	}
	return
}

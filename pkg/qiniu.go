package pkg

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"path"
	"strings"
	"time"
)

type QiNiuClient struct {
	AccessKey     string `json:"access_key"`
	SecretKey     string `json:"secret_key"`
	Bucket        string `json:"bucket"`
	UseHTTPS      bool   `json:"use_https"`
	UseCdnDomains bool   `json:"use_cdn_domains"`
	Domain        string `json:"domain"`
	Subdir        string `json:"subdir"`
}

func NewQiNiuClient(accessKey, secretKey, bucket string, useHttps, useCdnDomains bool, domain, subdir string) *QiNiuClient {
	qiNiuClient := &QiNiuClient{
		AccessKey:     accessKey,
		SecretKey:     secretKey,
		Bucket:        bucket,
		UseHTTPS:      useHttps,
		UseCdnDomains: useCdnDomains,
		Domain:        domain,
		Subdir:        subdir,
	}
	if !strings.HasSuffix(qiNiuClient.Domain, "/") {
		qiNiuClient.Domain = qiNiuClient.Domain + "/"
	}
	return qiNiuClient
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
		key := q.Subdir + "/" + time.Now().Format("060102-150405") + "-" + fileName
		log.Println("Start uploading", image, "as", key)
		// TODO: go routine
		if err := formUploader.PutFile(context.Background(), &ret, token, key, image, nil); err != nil {
			log.Fatalln("Error:", err)
		}
		urls = append(urls, q.Domain+key)
		log.Println("Done uploading", image)
	}
	return
}

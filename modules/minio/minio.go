package minio

import (
	"context"
	"fmt"
	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/gs/cond"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConfig struct {
	Enable bool   `value:"${minio.enable:=true}"`    // 是否启用 HTTP
	Host   string `value:"${minio.host:=127.0.0.1}"` // HTTP host
	Port   int    `value:"${minio.port:=9000}"`      // HTTP 端口
	Access string `value:"${minio.access:=}"`        // Access
	Secret string `value:"${minio.secret:=}"`        // Secret
	Secure bool   `value:"${minio.secure:=true}"`    // Secure
	Bucket string `value:"${minio.bucket:=}"`
}

func init() {

	gs.Provide(clientMinio).
		// Bean 名称
		Name("minio-client").
		// cond.OnProperty 会检查配置项
		On(cond.OnProperty("minio.enable", cond.HavingValue("true")))
}

func clientMinio(config MinioConfig) *minio.Client {

	client, err := minio.New(fmt.Sprintf("%s:%d", config.Host, config.Port), &minio.Options{
		Creds:  credentials.NewStaticV4(config.Access, config.Secret, ""),
		Secure: config.Secure,
	})

	if err != nil {
		panic("minio client error" + err.Error())
	}

	err = client.MakeBucket(context.Background(), config.Bucket, minio.MakeBucketOptions{
		// 嘻嘻，这个桶是外太空的
		Region:        "waitaikong",
		ObjectLocking: false,
	})

	if err != nil {
		panic(fmt.Sprintf("make %s bucket error: %v", config.Bucket, err))
	}

	log.Infof("%v", client.EndpointURL())
	return client
}

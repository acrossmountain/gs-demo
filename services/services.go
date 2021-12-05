package services

import (
	"github.com/minio/minio-go/v7"
	"learn/services/filesystem/local"
	fminio "learn/services/filesystem/minio"
	"learn/types"

	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/gs/cond"
)

func init() {

	gs.Object(new(local.Service)).
		Export((*types.FileProvider)(nil)).
		// 本地存储，当 minio 不存在时才注册
		// 可以添加其它判断条件，例如 aliyun 等
		On(cond.Group(cond.And, cond.OnMissingBean((*minio.Client)(nil))))

	gs.Object(new(fminio.Service)).
		Export((*types.FileProvider)(nil)).
		On(cond.OnBean((*minio.Client)(nil)))
}

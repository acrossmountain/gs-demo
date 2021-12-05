package minio

import (
	"context"
	"errors"
	"io"
	"path"

	"github.com/go-spring/spring-base/log"

	"github.com/minio/minio-go/v7"
)

type Service struct {
	// 自动注入 minio-client
	Client *minio.Client `autowire:""`
	// 存储桶
	Bucket string `value:"${minio.bucket}"`
	// 存储路径
	Dir string `value:"${file.dir}"`
}

func (s *Service) PutObject(name string, r io.Reader, size int64) (string, error) {
	out := path.Join(s.Dir, name)
	if s.ExistsObject(out) {
		return "", errors.New("文件已存在")
	}

	_, err := s.Client.PutObject(context.Background(), s.Bucket, out, r, size, minio.PutObjectOptions{})
	if err != nil {
		log.Errorf("minio upload error: %v", err)
		return "", errors.New("文件上传失败")
	}

	return out, nil
}

func (s *Service) ExistsObject(name string) bool {
	_, err := s.Client.StatObject(context.Background(), s.Bucket, name, minio.StatObjectOptions{})
	if err != nil {
		log.Error(err)
		if err.Error() == "The specified key does not exist." {
			return false
		}
	}

	return true
}

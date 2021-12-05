package local

import (
	"errors"
	"io"
	"os"
	"path"
)

type Service struct {
	Dir string `value:"${file.dir}"`
}

func (s *Service) PutObject(name string, r io.Reader, size int64) (string, error) {

	out := path.Join(s.Dir, name)
	if s.ExistsObject(out) {
		return "", errors.New("文件已存在，请勿重复上传")
	}

	dir := path.Dir(out)
	if !s.ExistsObject(dir) {
		err := os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			return "", errors.New("文件上传失败，请稍后重试")
		}
	}

	dst, err := os.OpenFile(out, os.O_CREATE, 0644)
	if err != nil {
		return "", errors.New("文件保存失败，请稍后重试")
	}
	defer func() {
		_ = dst.Close()
	}()

	_, err = io.Copy(dst, r)
	if err != nil {
		return "", errors.New("文件保存失败")
	}

	return out, nil
}

func (s *Service) ExistsObject(name string) bool {

	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

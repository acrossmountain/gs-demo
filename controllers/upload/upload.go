package upload

import (
	"learn/types"

	"github.com/go-spring/spring-core/web"
)

type Controller struct {
	FileService types.FileProvider `autowire:""`
}

func (c *Controller) Upload(ctx web.Context) {

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"code": -1,
			"msg":  "上传文件不能为空",
		})
		return
	}

	f, err := file.Open()
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"code": -1,
			"msg":  "上传文件错误",
		})
		return
	}
	defer func() {
		// 打开的资源。一定要记住主动关闭
		_ = f.Close()
	}()

	if out, err := c.FileService.PutObject(file.Filename, f, file.Size); err == nil {
		ctx.JSON(map[string]interface{}{
			"code": 0,
			"msg":  "上传文件成功",
			"data": map[string]interface{}{
				"url": out,
			},
		})
		return
	} else {
		ctx.JSON(map[string]interface{}{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
}

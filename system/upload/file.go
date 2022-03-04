package upload

import (
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/fangjie-luoxi/growing_api/global"
)

// FilesUpload 多文件上传
func FilesUpload(c *gin.Context) {
	// 文件夹
	folder := c.PostForm("folder")
	if folder == "" {
		folder = "file"
	}
	currPath, _ := os.Getwd()
	filePath := path.Join(currPath, "static", folder)
	err := os.MkdirAll(filePath, 0755)

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		global.Resp.Error(c, 400, err, "")
		return
	}
	files := form.File["files"]

	fileMap := make(map[string]string, len(files))

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, path.Join(filePath, filename)); err != nil {
			global.Resp.Error(c, 400, err, "")
			return
		}
		fileMap[filename] = path.Join("/static", folder, filename)
	}
	global.Resp.OK(c, 200, fileMap)
}

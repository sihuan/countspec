package tarball

import (
	"archive/tar"
	"errors"
	"io"
	"mygo/core/models"
	"mygo/lib/speccpu"
	"mygo/services/sqlite"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ulikunitz/xz"
)

func CreateTarball(c *gin.Context) {
	var Tarball models.Tarball
	// 从请求中读取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	Tarball.UUID = uuid.New().String()
	Tarball.Deleted = false
	Tarball.Size = file.Size
	Tarball.Filename = file.Filename
	Tarball.ConfigPath = ""
	// 从请求中读取描述
	Tarball.Description = c.PostForm("description")

	basePath := "data/upload/tarball/" + Tarball.UUID + "/"

	returnErrorAndClenaup := func(err error) {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		os.RemoveAll(basePath)
	}

	// 上传文件到指定的路径
	// 流式解压 .tar.xz 文件
	f, err := file.Open()
	if err != nil {
		returnErrorAndClenaup(err)
		return
	}
	defer f.Close()

	r, err := xz.NewReader(f)
	if err != nil {
		returnErrorAndClenaup(err)
		return
	}
	tr := tar.NewReader(r)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			returnErrorAndClenaup(err)
			return
		}

		if header.Typeflag == tar.TypeReg {
			if strings.HasPrefix(header.Name, "config") && strings.HasSuffix(header.Name, ".cfg") {
				// 创建 confg 文件夹
				if err := os.MkdirAll(basePath+"config", os.ModePerm); err != nil {
					returnErrorAndClenaup(err)
					return
				}
				// 创建文件
				f, err := os.Create(basePath + header.Name)
				if err != nil {
					returnErrorAndClenaup(err)
					return
				}
				defer f.Close()
				_, err = io.Copy(f, tr)
				if err != nil {
					returnErrorAndClenaup(err)
					return
				}
				if Tarball.ConfigPath == "" {
					Tarball.ConfigPath = basePath + header.Name
				} else {
					returnErrorAndClenaup(errors.New("config file already exists"))
					return
				}
			} else if strings.HasPrefix(header.Name, "benchspec/CPU") {
				benchmark := strings.Split(header.Name, "/")[2]
				if speccpu.IsBenchmark(benchmark) {
					Tarball.Benchmarks = Tarball.Benchmarks.Add(benchmark)
					// 创建文件夹
					if err := os.MkdirAll(basePath+"benchmark/"+benchmark, os.ModePerm); err != nil {
						returnErrorAndClenaup(err)
						return
					}
					// 创建文件
					// 修改文件名，xxx_base.yyy -> xxx.mygo
					fileName := strings.Split(header.Name, "/")[4]
					newName := strings.Split(fileName, "_base.")[0] + ".mygo"
					f, err := os.Create(basePath + "benchmark/" + benchmark + "/" + newName)
					if err != nil {
						returnErrorAndClenaup(err)
						return
					}
					defer f.Close()

					_, err = io.Copy(f, tr)
					if err != nil {
						returnErrorAndClenaup(err)
						return
					}
					// f.Chmod(os.ModePerm)

				}
			}
		}
	}
	if err := sqlite.MyGODB.Create(&Tarball).Error; err != nil {
		returnErrorAndClenaup(err)
		return
	}
	status := http.StatusCreated
	resp := models.ResponseModel{
		Status:  status,
		Message: "success",
		Data:    Tarball,
		Errors:  nil,
	}
	c.JSON(status, resp.NewResponse())
}

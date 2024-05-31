package task

import (
	"bufio"
	"errors"
	"io"
	"log"
	"mygo/core/models"
	"mygo/services/sqlite"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TODO
// var qemuPath = "/home/sihuan/PLCT/spec/MyGO/data/qemu/"
var qemuPath = "/home/sihuan/project/mygo/data/qemu/"
var sysroot = "/opt/riscv/sysroot"

// TODO
var specCPURUN = "data/runenv/"
var qemuCreateLock sync.Mutex

// TODO
var qemuPool = make(chan bool, 100)

type createTaskRequest struct {
	TarballID  uint     `json:"tarball_id"`
	Type       string   `json:"type"`
	Benchmarks []string `json:"benchmarks"`
}

func CreateTask(c *gin.Context) {
	var req createTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if req.Type == "" {
		req.Type = "test"
	}

	if req.Type != "test" && req.Type != "ref" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "type must be test or ref",
		})
		return
	}

	var tarball models.Tarball
	if err := sqlite.MyGODB.First(&tarball, req.TarballID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, v := range req.Benchmarks {
		if !tarball.HasBenchmark(v) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "tarball does not have benchmark " + v,
			})
			return
		}
	}

	// 	task := models.Task{
	// 		Title:       req.Title,
	// 		Description: req.Description,
	// 		TarballID:   req.TarballID,
	// 		Benchmarks:  req.Benchmarks,
	// 		Completed:   false,
	// 		Type:        req.Type,
	// 		Status:      "pending",
	// 	}
	// 	if err := sqlite.MyGODB.Create(&task).Error; err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	go createTask(&task)

	// 	status := http.StatusCreated
	// 	resp := models.ResponseModel{
	// 		Status:  status,
	// 		Message: "success",
	// 		Data:    task,
	// 		Errors:  nil,
	// 	}
	// 	c.JSON(status, resp.NewResponse())
	// }

	// func createTask(t *models.Task) {
	// t.Status = "setup"
	// sqlite.MyGODB.Save(&t)
	// 查询 tarball uuid
	// var tarball models.Tarball
	// if err := sqlite.MyGODB.First(&tarball, t.TarballID).Error; err != nil {
	// 	t.Completed = true
	// 	t.Status = "failed"
	// 	sqlite.MyGODB.Save(&t)
	// 	log.Print("find tarball failed", err)
	// 	return
	// }
	// task 目录
	tarballPath := tarball.RealPath()
	taskPath := tarballPath + "/run/" + req.Type
	err := os.MkdirAll(taskPath, os.ModePerm)
	if err != nil {
		log.Print("mkdir task path failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 复制 runenv 到 task 目录下
	var newQemuTasks []models.QemuTask

	qemuCreateLock.Lock()
	defer qemuCreateLock.Unlock()
	for _, benchmark := range req.Benchmarks {

		// 如果 tarball 中有相同的 benchmark，直接跳过，相同是指 benchmark 名字相同，type 相同
		var qemuTask models.QemuTask
		result := sqlite.MyGODB.Where(&models.QemuTask{TarballID: tarball.ID, Benchmark: benchmark, Type: req.Type}).Where("status IN ?", []string{"running", "success", "pending"}).First(&qemuTask)
		if result.Error == nil && result.RowsAffected > 0 {
			log.Print("qemu task already exists", result.RowsAffected)
			continue
		} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Print("find qemu task failed", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		runenvPath := specCPURUN + benchmark + "/" + req.Type
		taskRunenvPath := taskPath + "/" + benchmark
		tarballExePath := tarballPath + "/benchmark/" + benchmark
		err := os.MkdirAll(taskRunenvPath, os.ModePerm)
		if err != nil {
			log.Print("mkdir task runenv path failed", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		cmd := exec.Command("cp", "-r", runenvPath, taskRunenvPath)
		if err := cmd.Run(); err != nil {
			log.Print("copy runenv failed", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		taskRunenvPath = taskRunenvPath + "/" + req.Type

		//遍历 tarballExePath 下的文件，复制到 taskRunenvPath 下
		dirs, err := os.ReadDir(tarballExePath)
		if err != nil {
			log.Print("read tarball exe dir failed", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		for _, dir := range dirs {
			if dir.IsDir() {
				continue
			}
			cmd := exec.Command("cp", tarballExePath+"/"+dir.Name(), taskRunenvPath)
			if err := cmd.Run(); err != nil {
				log.Print("copy exe failed", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			os.Chmod(taskRunenvPath+"/"+dir.Name(), os.ModePerm)
		}

		cmdfilePath := taskRunenvPath + "/mygocmd"
		var cmdlines []string
		cmdfile, err := os.Open(cmdfilePath)
		if err != nil {
			log.Panic("open mygocmd failed", err)
		}
		r := bufio.NewReader(cmdfile)
		for {
			// ReadLine is a low-level line-reading primitive.
			// Most callers should use ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
			bytes, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Panic("readline failed", err)
			}
			line := string(bytes)
			if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "specinvoke exit:") {
				continue
			}
			cmdlines = append(cmdlines, line)
		}
		cmdfile.Close()
		flag := false
		if len(cmdlines) > 1 {
			flag = true
		}

		for i, cmdline := range cmdlines {
			var qemuTask models.QemuTask
			qemuTask.Status = "pending"
			qemuTask.TarballID = tarball.ID
			qemuTask.Benchmark = benchmark
			qemuTask.Type = req.Type
			// qemuTask.TaskID = t.ID
			qemuTask.Path = taskRunenvPath

			if flag {
				qemuTask.Name = benchmark + "#" + strconv.Itoa(i+1)
			} else {
				qemuTask.Name = benchmark
			}
			cmdline = strings.TrimSpace(cmdline)
			//清理关闭标准输入 0<&-，2006 适配
			cmdline = strings.ReplaceAll(cmdline, "0<&-", "")
			cmdline = strings.ReplaceAll(cmdline, "  ", " ")
			temp := strings.Split(cmdline, "2>>")
			if len(temp) == 2 {
				qemuTask.Stderr = strings.TrimSpace(temp[1])
			}
			temp = strings.Split(strings.TrimSpace(temp[0]), ">")
			if len(temp) == 2 {
				qemuTask.Stdout = strings.TrimSpace(temp[1])
			}
			temp = strings.Split(strings.TrimSpace(temp[0]), "<")
			if len(temp) == 2 {
				qemuTask.Stdin = strings.TrimSpace(temp[1])
			}
			// temp = strings.Split(strings.TrimSpace(temp[0]), " ")
			// qemuTask.Cmd = temp[0]
			// qemuTask.Args = temp[1:]
			qemuTask.Cmd = strings.TrimSpace(temp[0])
			if err := sqlite.MyGODB.Create(&qemuTask).Error; err != nil {
				log.Panic("create qemu task failed", err)
			}
			newQemuTasks = append(newQemuTasks, qemuTask)
			tarball.QemuTasks = append(tarball.QemuTasks, qemuTask)
			sqlite.MyGODB.Save(&tarball)
		}
	}

	for _, q := range newQemuTasks {
		go runQemuTask(&q)
	}
	status := http.StatusCreated
	resp := models.ResponseModel{
		Status:  status,
		Message: "success",
		Data:    newQemuTasks,
		Errors:  nil,
	}
	c.JSON(status, resp.NewResponse())

}

func runQemuTask(q *models.QemuTask) {
	qemuPool <- true
	defer func() {
		<-qemuPool
	}()
	// -s 大小需要设置 TODO
	cmd := exec.Command(qemuPath+"bin/qemu-riscv64", "-L", sysroot, "-cpu", "rv64,v=true,vext_spec=v1.0", "-s", "819200000000", "-plugin", qemuPath+"/plugins/libinsn.so", "-d", "plugin")
	args := strings.Split(q.Cmd, " ")
	cmd.Args = append(cmd.Args, args...)
	cmd.Dir = q.Path
	if q.Stdin != "" {
		stdin, err := os.Open(q.Path + "/" + q.Stdin)
		if err != nil {
			log.Println("open stdin failed", err, q.Path+"/"+q.Stdin)
		}
		cmd.Stdin = stdin
	}
	if q.Stdout != "" {
		stdout, err := os.OpenFile(q.Path+"/"+q.Stdout, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Println("open stdout failed", err, q.Path+"/"+q.Stdout)
		}
		cmd.Stdout = stdout
	}
	if q.Stderr != "" {
		stderr, err := os.OpenFile(q.Path+"/"+q.Stderr, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Println("open stderr failed", err, q.Path+"/"+q.Stderr)
		}
		cmd.Stderr = stderr
	} else {
		stderr, err := os.OpenFile(q.Path+"/ins", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Println("open stderr failed", err, q.Path+"/ins")
		}
		cmd.Stderr = stderr
	}
	q.Status = "running"
	sqlite.MyGODB.Save(&q)

	log.Println("run qemu task", "tarball id", q.TarballID, "qemu id", q.ID, "benchmark", q.Benchmark, "type", q.Type)
	log.Println("cmd", cmd.String())
	log.Println("dir", cmd.Dir)

	if err := cmd.Run(); err != nil {
		log.Println("run qemu task failed", err, "tarball id", q.TarballID, "qemu id", q.ID, "benchmark", q.Benchmark, "type", q.Type)
		q.Status = "failed"
		q.Error = err.Error()
		sqlite.MyGODB.Save(&q)
		return
	}
	log.Println("run qemu task success", "tarball id", q.TarballID, "qemu id", q.ID, "benchmark", q.Benchmark, "type", q.Type)

	q.Status = "success"
	sqlite.MyGODB.Save(&q)

	// 读取 stderr 最后一行，获取 inscount
	var stderrFile string
	if q.Stderr != "" {
		stderrFile = q.Path + "/" + q.Stderr
	} else {
		stderrFile = q.Path + "/ins"
	}
	stderr, err := os.Open(stderrFile)
	if err != nil {
		log.Println("open stderr failed for ins", err, stderrFile)
		q.Error = "open stderr failed for ins"
		sqlite.MyGODB.Save(&q)
		return
	}

	r := bufio.NewReader(stderr)
	for {
		bytes, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("readline failed", err)
			q.Error = "readline failed"
			sqlite.MyGODB.Save(&q)
			return
		}
		line := string(bytes)
		if strings.HasPrefix(line, "total insns:") {
			inscount, err := strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1]))
			if err != nil {
				log.Println("inscount atoi failed", err, inscount, q.ID)
				q.Error = "inscount atoi failed"
				sqlite.MyGODB.Save(&q)
				return
			}
			q.Inscount = inscount
			sqlite.MyGODB.Save(&q)
			break
		}
	}
}

//TODO

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

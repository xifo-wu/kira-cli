package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func Auto(path string, filename string) {
	savePath := viper.GetString("save_path")
	rclonePath := viper.GetString("rclone_path")
	volumePath := viper.GetString("volume_path")

	log.Println("##############################")
	log.Println("####### Kira 开始运行 ########")
	log.Println("##############################")

	var files []string
	// 进行重命名操作
	ext := filepath.Ext(filename)
	fmt.Println(ext, "ext")
	if ext == "" {
		files = append(files, MoveToParentDir(path, filename)...)
	} else {
		files = append(files, filename)
	}

	for _, file := range files {
		newName := Rename(path, file)
		// 参数传递进来的路径，如果是 docker 可能需要替换一下路径
		src := filepath.Join(path, newName)

		localPath := src
		if volumePath != "" {
			localPath = strings.Replace(src, volumePath, savePath, 1)
		}

		dstPath := strings.Replace(src, savePath, rclonePath, 1)

		log.Printf("原始路径: %s", localPath)
		log.Printf("重命名后: %s", dstPath)

		cmd := exec.Command("rclone", "moveto", "-v", localPath, dstPath)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

func MoveToParentDir(path string, lastPath string) []string {
	dirPath := filepath.Join(path, lastPath)
	// 获取目录下的所有文件
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	filesNames := make([]string, 0, len(files))

	// 遍历所有文件
	for _, file := range files {
		// 判断是否为文件
		if !file.IsDir() {
			// 拼接旧文件路径和新文件路径
			oldFilePath := filepath.Join(dirPath, file.Name())

			newFilePath := filepath.Join(filepath.Dir(dirPath), file.Name())

			// 移动文件到上级目录
			err := os.Rename(oldFilePath, newFilePath)
			if err != nil {
				panic(err)
			}

			filesNames = append(filesNames, file.Name())
		}
	}

	// TODO 剩下空文件夹就删掉

	return filesNames
}

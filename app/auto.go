package app

import (
	"log"
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

	// 进行重命名操作
	ext := filepath.Ext(filename)
	if ext == "" {
		return
	}

	newName := Rename(path, filename)
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

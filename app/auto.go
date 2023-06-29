package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

func Auto(path string, filename string) {
	savePath := viper.GetString("save_path")
	rclonePath := viper.GetString("rclone_path")
	volumePath := viper.GetString("volume_path")
	localPath := path
	if volumePath != "" {
		localPath = strings.Replace(path, volumePath, savePath, 1)
	}

	log.Println("##############################")
	log.Println("####### Kira 开始运行 ########")
	log.Println("##############################")

	var files []string

	fmt.Println(localPath)

	file, err := os.Stat(filepath.Join(localPath, filename))
	if err != nil {
		log.Fatal(err)
	}

	if !file.Mode().IsRegular() {
		files = append(files, MoveToParentDir(localPath, filename)...)
	} else {
		files = append(files, filename)
	}

	for _, file := range files {
		newName, fileSize := Rename(localPath, file)
		// 参数传递进来的路径，如果是 docker 可能需要替换一下路径
		src := filepath.Join(localPath, newName)
		dstPath := strings.Replace(src, savePath, rclonePath, 1)

		log.Printf("原始路径: %s", src)
		log.Printf("重命名后: %s", dstPath)

		var notificationInfo string
		message := ""
		needChannelNotification := false

		videoRe := regexp.MustCompile(`\.(mp4|mov|avi|wmv|mkv|flv|webm|vob|rmvb|mpg|mpeg)$`)
		if videoRe.MatchString(strings.ToLower(newName)) {
			title := strings.Replace(src, savePath, "", 1)
			standardTitleRe := regexp.MustCompile(`S\d+E\d+`)
			info := standardTitleRe.FindString(title)
			notificationInfo = info
			needChannelNotification = true
			message = getAnimeName(path) + " " + info + " 入库成功 🎉"
			log.Println(message, "message")
		}

		cmd := exec.Command("rclone", "moveto", "-v", src, dstPath)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

		if message != "" {
			Notification(message)
		} else {
			Notification(fmt.Sprintf("上传了 %s", dstPath))
		}

		if needChannelNotification {
			// JSON 已配置的话推送消息到 TG 群里
			ChannelNotification(notificationInfo, path, fileSize)
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

func getAnimeName(path string) string {
	re := regexp.MustCompile(`(?i)(?:Season (\d+)|S(\d+))`)
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		match := re.FindStringSubmatch(parts[i])

		if len(match) >= 1 && i-1 >= 0 {
			return parts[i-1]
		}
	}

	return ""
}

package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func Rename(path string, filename string) string {
	// 判断 Filename 是否有扩展名，有时下载过来的是文件夹
	// 需要移动到上级目录
	ext := filepath.Ext(filename)
	if ext == "" {
		return filename
	}

	newFilename := GenerateSeasonAndEpisode(path, filename)

	oldPath := filepath.Join(path, filename)
	newPath := filepath.Join(path, newFilename)

	_, err := os.Stat(oldPath)
	if os.IsNotExist(err) {
		log.Println("[error] 资源不存在")
		os.Exit(1)
	}

	os.Rename(oldPath, newPath)

	return newFilename
}

func GenerateSeasonAndEpisode(path string, filename string) string {
	log.Println("======= 开始重命名 =======")
	// TODO 支持合集
	// 合集暂时跳过
	collectionRe := regexp.MustCompile(`(?i)(\d+-\d+|第\d+-\d+集|合集)`)
	matchCollection := collectionRe.FindString(filename)

	if matchCollection != "" {
		// 暂时什么都不处理
		return filename
	}

	_, lastPath := filepath.Split(path)

	re := regexp.MustCompile(`(?i)Season (\d+)`)
	matchSeason := re.FindStringSubmatch(lastPath)
	if len(matchSeason) < 1 {
		// 如果匹配不到 Season 就不需要生成季度信息和集数信息，直接返回原文件名
		return filename
	}

	seasonNum, _ := strconv.Atoi(matchSeason[1])

	// 匹配集数
	epRe := regexp.MustCompile(`(?i)( -? \d+|\[\d+]|\[\d+.?[vV]\d]|第\d+[话話集]|\[第?\d+[话話集]]|\[\d+.?END]|[e][p]?\d+|\[\d+\(\d+\)\]|\[\d+（\d+）\])`)
	matchEpisode := epRe.FindString(filename)

	episodeRe := regexp.MustCompile(`\d+`)
	episodeNum, _ := strconv.Atoi(episodeRe.FindString(matchEpisode))

	prefix := fmt.Sprintf("S%02dE%02d", seasonNum, episodeNum)

	renameFileName := strings.Replace(filename, matchEpisode, "", 1)
	renameFileName = prefix + " " + renameFileName

	return renameFileName
}

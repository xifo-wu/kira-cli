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
	// 判断 Filename 是否有扩展名，有时下载过来的是文件夹, 文件夹时不做处理
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
	// zip 文件也跳过重命名
	collectionRe := regexp.MustCompile(`(?i)(\d+-\d+|第\d+-\d+集|合集|\.zip)`)
	matchCollection := collectionRe.FindString(filename)

	if matchCollection != "" {
		// 暂时什么都不处理
		return filename
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	seasonNum, _ := strconv.Atoi(getSeasonNumber(path))
	if seasonNum == 0 {
		// 如果匹配不到 Season 就不需要生成季度信息和集数信息，直接返回原文件名
		return filename
	}

	standardTitleRe := regexp.MustCompile(`S\d+E\d+`)
	// 符合 S01E01 时直接返回文件名，不需要重命名
	if standardTitleRe.MatchString(filename) {
		return filename
	}

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

func getSeasonNumber(path string) string {
	re := regexp.MustCompile(`(?i)(?:Season (\d+)|S(\d+))`) // 编译正则表达式

	// 匹配季数
	match := re.FindStringSubmatch(path)

	seasonNumber := ""
	if len(match) < 1 {
		fmt.Println("未匹配到季数")
	} else {
		for i, _ := range re.SubexpNames() {
			if i != 0 && match[i] != "" {
				seasonNumber = match[i]
				break
			}
		}
		fmt.Println("季数为：", seasonNumber)
	}

	return seasonNumber
}

package kira_test

import (
	"testing"

	"xifo.in/kira/app"
)

// Rename 重命名文件
// 案例1：
// [桜都字幕组] 在无神世界里进行传教活动 / Kaminaki Sekai no Kamisama Katsudou [09][1080p][简繁日内封]
// 重命名
// [桜都字幕组] 在无神世界里进行传教活动 / Kaminaki Sekai no Kamisama Katsudou S01E09 [1080p][简繁日内封]
func TestGenerateSeasonAndEpisode(t *testing.T) {
	app.GenerateSeasonAndEpisode("/downloads/无神世界中的神明活动 (2023)/Season 1", "[Sakurato] Watashi no Yuri wa Oshigoto desu! [09][HEVC-10bit 1080p AAC][CHS&CHT].mkv")

	app.GenerateSeasonAndEpisode("/downloads/熊熊勇闯异世界 (2020)/Season 2", "[GJ.Y] 熊熊勇闯异世界 PUNCH！ / Kuma Kuma Kuma Bear Punch! - 11 (Baha 1920x1080 AVC AAC MP4).mkv")

	app.GenerateSeasonAndEpisode("/downloads/斗破苍穹 (2017)/Season 5", "[国漫]斗破苍穹年番第48集无水印高清迅雷下载.mp4")

	app.GenerateSeasonAndEpisode("/downloads/鬼灭之刃 (2019)/Season 4", "【豌豆字幕组&风之圣殿字幕组】★04月新番[鬼灭之刃 刀匠村篇 / Kimetsu_no_Yaiba-Katanakaji_no_Sato_Hen][10(54)][繁体][1080P][MP4].mkv")

	app.GenerateSeasonAndEpisode("/downloads/风都侦探 (2022)/Season 1", "[7³ACG] 风都侦探/风都探侦/Fuuto PI/Fuuto Tantei | 01-12 [简繁字幕] BDrip 1080p x265 FLAC [8.7GB]")

	app.GenerateSeasonAndEpisode("/downloads/名侦探柯南：犯人犯泽先生 (2022)/Season 1", "[风车字幕组][名侦探柯南番外-犯人犯泽先生][第1-12集][合集版][1080P][繁体][MP4] [2.0GB]")

	app.GenerateSeasonAndEpisode("/downloads/【我推的孩子】(2023)/Season 1", "[DMG][Oshi_no_Ko][07][1080P][GB].mp4")
}

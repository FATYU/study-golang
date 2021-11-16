package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string        // 标题
	Artist string        //艺术家
	Album  string        //专辑
	Year   int           //年份
	Length time.Duration //时长
}

var tracks = []*Track{{
	"以父之名", "jay1", "依然范特西", 2004, length("3m20s"),
}, {
	"以父之名3", "jay3", "依然范特西", 2006, length("4m20s"),
}, {
	"以父之名2", "jay2", "依然范特西", 2005, length("6m20s"),
}, {
	"以父之名4", "jay4", "依然范特西", 2007, length("5m20s"),
}}

func length(s string) time.Duration {

	result, error := time.ParseDuration(s)

	if error != nil {
		panic(s)
	}
	return result
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()

}

// 定义根据作者进行排序

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// --以上，如果想做五个字段的排序，需要进行实现五次接口-----
// 以下是通过一种结构体完成全字段排序，在真实使用中，只需要自定义比较函数即可

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// main测试代码块
func main() {

	printTracks(tracks)
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	//// 使用customSort进行自定义排序处理

}

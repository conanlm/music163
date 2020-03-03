package api

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

// Getmusic163 获取网易音乐歌曲地址
func Getmusic163() {
	os.Truncate("网易.txt", 0)
	os.Truncate("歌名.txt", 0)
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link

		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		// ioutil.WriteFile("a.txt", []byte(link), 0777)
		f, _ := os.OpenFile("网易.txt", os.O_RDWR|os.O_APPEND, 0777)  //读写模式打开，写入追加
		ff, _ := os.OpenFile("歌名.txt", os.O_RDWR|os.O_APPEND, 0777) //读写模式打开，写入追加
		defer f.Close()
		if strings.Contains(link, "song") {
			if strings.Index(link, "id=") != -1 {
				url := link[strings.Index(link, "id="):len(link)]
				num, _ := f.Write([]byte("http://music.163.com/song/media/outer/url?" + url + ".mp3"))
				num, _ = f.Write([]byte("\n"))
				ff.Write([]byte(e.Text + "\n"))
				fmt.Println(num)
			}

		}

		fmt.Println(e)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://music.163.com/album?id=71880502")

}

// Getsong 歌曲下载
func Getsong() (URLS []string, paths []string) {

	// song := []string{}
	// urls := []string{}
	var songs []string
	var urls []string

	fi, err := os.Open("歌名.txt")
	fi1, err1 := os.Open("网易.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if err1 != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	defer fi1.Close()

	br := bufio.NewReader(fi)
	br1 := bufio.NewReader(fi1)
	// fmt.Println(br)
	for {
		a, _, c := br.ReadLine()
		a1, _, c1 := br1.ReadLine()
		if c == io.EOF {
			break
		}
		if c1 == io.EOF {
			break
		}

		songs = append(songs, string(a))
		urls = append(urls, string(a1))

	}

	return urls, songs

	// DownLoadSingle(urls, song)

	// DownLoadmultithreading(urls, song, 1)

}

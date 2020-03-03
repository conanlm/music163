package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

// DownLoadSingle 单线程下载
func DownLoadSingle(URLS []string, paths []string) {

	for i, durl := range URLS {
		_, err2 := url.ParseRequestURI(durl)
		if err2 != nil {
			panic("网址错误")
		}

		filename := path.Join("歌曲/", paths[i]+".mp3")
		client := http.Client{Timeout: 900 * time.Second}
		res, err := client.Get(durl)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		io.Copy(f, res.Body)
		defer f.Close()
		defer res.Body.Close()
	}

}

// //DownLoadmultithreading 多线程下载
func DownLoadmultithreading(URLS []string, paths []string, i int) {
	fmt.Println(URLS[i])
	fmt.Println(paths[i])
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	num := 99
	loopNum := 3
	jobs := make(chan int, num) //Job为100个可以传递int类型的channel
	results := make(chan int, loopNum)

	//开启三个线程，说明线程池中只有三个线程， 在实际情况下可以动态设置开启线程数量
	for w := 1; w <= loopNum; w++ {
		go worker(w, jobs, results, paths)
	}

	// 添加9个任务
	for j := 1; j <= num; j++ {
		jobs <- j //向Jobs添加任务： 向Channel中写入数据， 传递的数据类型为int
	}
	//关闭Channel
	close(jobs)

	for a := 1; a <= num; a++ {
		res := <-results                           //从Channel中读取数据, 输出的数据类型为 int
		fmt.Println("ID = ", a, "Results = ", res) //注意 a 与 res 不对应是由于处理器调度的结果
	}

}

func worker(id int, jobs <-chan int, results chan<- int, paths []string) {
	for j := range jobs {
		fmt.Println("worker", "processing job", j, paths[j])
		time.Sleep(10 * time.Second)
		results <- j
	}

}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// SongData 歌曲数据
type SongData struct {
	url  string
	path string
}

var sum int

func myFunc(i interface{}) {
	sd, ok := i.(SongData)
	if !ok {
		return
	}
	// n := i.url()
	// m := i.path()
	// atomic.AddInt32(&sum, n)
	// fmt.Printf("run with %s\n", n)
	// fmt.Printf("run with %s\n", m)
	DownLoads(sd.url, sd.path)
	fmt.Println("下载完毕" + sd.path)
}

// Mygoroutinepool 多线程下载
func Mygoroutinepool() {

	var sd SongData
	URLS, paths := Getsong()
	runTimes := len(URLS)

	// Use the common pool.
	var wg sync.WaitGroup

	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	// d := SongData{
	// 	url:  "qwe",
	// 	path: "asd",
	// }

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		sd.path = paths[i]
		sd.url = URLS[i]
		_ = p.Invoke(sd)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}

// DownLoads 下载
func DownLoads(URL string, Path string) {
	_, err2 := url.ParseRequestURI(URL)
	if err2 != nil {
		panic("网址错误")
	}
	_dir := "歌曲"
	exist, err := PathExists(_dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if !exist {
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		}
	}

	filename := path.Join(_dir+"/", Path+".mp3")
	client := http.Client{Timeout: 900 * time.Second}
	res, err := client.Get(URL)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
	defer f.Close()
	defer res.Body.Close()

}

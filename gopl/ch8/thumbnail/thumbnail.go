package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func main() {

	fmt.Println("start:::::::::::::::>", time.Now().Format("2016-01-02 15:04:05"))
	makeThumbnail2(os.Args[1:])
	fmt.Println("  end:::::::::::::::>", time.Now().Format("2016-01-02 15:04:05"))
}

//非并发版本
func makeThumbnail(files []string) {

	for _, f := range files {
		ImageFile(f) // 忽略错误
	}
}

//启动 goroutine 异步处理，直接返回结果（执行完毕）
func makeThumbnail2(files []string) {
	for _, f := range files {
		go ImageFile(f) // 忽略错误
	}
}

//启动goroutine，并等待goroutine运行结束后，再返回结果
//通过共享channel方式实现事件通知
func makeThumbnail3(files []string) {
	channel := make(chan struct{})

	for _, file := range files {
		go func(filePath string) {
			ImageFile(filePath)
			channel <- struct{}{}
		}(file)

	}
	for range files { // 迭代文件内容，获取channel中数据，等待goroutine完成
		<-channel
	}
}

func makeThumbnail4(files []string) error {

	errors := make(chan error)
	for _, file := range files {
		go func(filePath string) {
			_, err := ImageFile(filePath)
			errors <- err
		}(file)

	}

	for range files {
		if err := <-errors; err != nil {
			//此语句会导致 goroutine 泄漏  因为在所有的goroutine未执行完毕前，可能因为某个goroutine出现error，然后函数直接return

			//解决办法： 创建一个合适大小缓存的channel（buffered channel 参见 makeThumbnail5）
			return err
		}

	}
	return nil
}

func makeThumbnail5(files []string) (thumbnailFiles []string, err error) {

	//定义结构体，包含生成的文件路径和错误信息

	type item struct {
		thumbnailFile string
		err           error
	}

	flag := make(chan item, len(files)) //创建带有缓冲区的channel

	for _, file := range files {
		go func(file string) {
			var item item
			item.thumbnailFile, item.err = ImageFile(file)
			flag <- item
		}(file)

	}

	for range files {
		f := <-flag
		if f.err != nil {
			return nil, f.err //return不会中断for循环吗？
		}
		thumbnailFiles = append(thumbnailFiles, f.thumbnailFile)
	}
	return thumbnailFiles, nil
}

// 使用sync.waitGroup方式等待所有的goroutine执行完毕，类似java中的CountdownLatch

//获取所有图片处理后的文件的总byte大小
func makeThumbnail6(files []string) int64 {

	sizes := make(chan int64) //保存每个goroutine的执行结果的字节大小

	swg := sync.WaitGroup{}

	for _, file := range files {
		swg.Add(1) //对sycn进行加1操作

		go func(file string) {
			defer swg.Done() // 进行减1操作

			filePath, err := ImageFile(file)
			if err != nil {
				log.Print(err)
				return //当前匿名函数返回，不影响for循环
			}
			stat, err := os.Stat(filePath) //统计路径文件大小
			if err != nil {
				log.Print(err)
				return
			}
			sizes <- stat.Size()
		}(file)

	}
	//等待所有的goroutine结束
	go func() {
		swg.Wait()
		close(sizes)
	}()

	var total int64

	for size := range sizes {
		total += size
	}

	return total
}

/// ----------------------------- ⬇️ 导入 thumbnail.go 文件内容函数 -----------------------------------------------

// Image returns a thumbnail-size version of src.
func Image(src image.Image) image.Image {
	// Compute thumbnail size, preserving aspect ratio.
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) // portrait
	} else {
		height = int(128 / aspect) // landscape
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// a very crude scaling algorithm
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

// ImageStream reads an image from r and
// writes a thumbnail-size version of it to w.
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

// ImageFile2 reads an image from infile and writes
// a thumbnail-size version of it to outfile.
func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}
	return out.Close()
}

// ImageFile reads an image from infile and writes
// a thumbnail-size version of it in the same directory.
// It returns the generated file name, e.g. "foo.thumb.jpeg".
func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile) // e.g., ".jpg", ".JPEG"
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}

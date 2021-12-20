package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	fmt.Println("start:::::::::::::::>", time.Now().Format("2016-01-02 15:04:05"))
	makeTumbnail2(os.Args[1:])
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
		return nil
	}

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

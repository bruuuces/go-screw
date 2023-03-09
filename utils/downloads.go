package utils

import (
	"bufio"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

func main() {

}
func DownloadFile(durl string, filepath, filename string) (*os.File, error) {
	if filename == "" {
		uri, err := url.ParseRequestURI(durl)
		if err != nil {
			return nil, errors.New("网址错误")
		}
		filename = path.Base(uri.Path)
	}
	client := http.DefaultClient
	client.Timeout = time.Second * 60 //设置超时时间
	resp, err := client.Get(durl)
	if err != nil {
		return nil, err
	}
	raw := resp.Body
	defer raw.Close()
	reader := bufio.NewReaderSize(raw, 1024*32)

	file, err := os.Create(filepath + PathSeparator + filename)
	if err != nil {
		return nil, err
	}
	writer := bufio.NewWriter(file)
	buff := make([]byte, 32*1024)
	written := 0
	go func() {
		for {
			nr, er := reader.Read(buff)
			if nr > 0 {
				nw, ew := writer.Write(buff[0:nr])
				if nw > 0 {
					written += nw
				}
				if ew != nil {
					err = ew
					break
				}
				if nr != nw {
					err = io.ErrShortBuffer
					break
				}
			}
			if er != nil {
				if er != io.EOF {
					err = er
				}
				break
			}
		}
		if err != nil {
			panic(err)
		}
	}()

	spaceTime := time.Second * 1
	ticker := time.NewTicker(spaceTime)
	lastWritten := 0
	stop := false

	for {
		select {
		case <-ticker.C:
			if written-lastWritten == 0 {
				ticker.Stop()
				stop = true
				break
			}
			lastWritten = written
		}
		if stop {
			break
		}
	}
	return file, nil
}

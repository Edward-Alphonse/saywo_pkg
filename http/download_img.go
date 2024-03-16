package http

import (
	"context"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type DownloadImgLoader struct {
	ctx    context.Context
	bucket string
	imgUrl string

	imgData []byte
	format  string
}

func NewDownloadImgLoader(ctx context.Context, imgUrl string) *DownloadImgLoader {
	return &DownloadImgLoader{
		ctx:    ctx,
		imgUrl: imgUrl,
	}
}

func (l *DownloadImgLoader) Load() error {
	response, err := http.Get(l.imgUrl)
	if err != nil {
		return errors.Wrap(err, "DownloadImgLoader load failed")
	}
	defer response.Body.Close()
	contentType := response.Header.Get("Content-Type")
	format := "webp"
	switch {
	case strings.HasPrefix(contentType, "image/jpeg"):
		format = "jpg"
	case strings.HasPrefix(contentType, "image/png"):
		format = "png"
	case strings.HasPrefix(contentType, "image/gif"):
		format = "gif"
	case strings.HasPrefix(contentType, "image/webp"):
		format = "webp"
	}

	//url, _ := url2.Parse(l.imgUrl)
	//fileName := path.Base(url.Path)
	//destination := "./" + fileName + "." + format
	//file, err := os.Create(destination)
	//if err != nil {
	//	return errors.Wrap(err, "DownloadImgLoader load failed")
	//}
	//defer file.Close()

	// 将图片数据写入文件
	//_, err = io.Copy(file, response.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "DownloadImgLoader load failed")
	}
	//fmt.Printf("图片已成功下载到 %s\n", destination)
	//l.imgPath = destination
	l.imgData = body
	l.format = format
	return nil
}

func (l *DownloadImgLoader) GetData() ([]byte, string) {
	return l.imgData, l.format
}

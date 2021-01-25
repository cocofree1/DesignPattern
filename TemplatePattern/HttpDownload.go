package TemplatePattern

import "fmt"

// 实现下载模板接口
type HttpDownload struct{}

func NewHttpDownload()Download{
	download := &HttpDownload{}
	return NewTemplateImplement(download)
}

func (h* HttpDownload)Download(url string){
	fmt.Printf("http下载%s\n", url)
}

func (h* HttpDownload)Save(){
	fmt.Println("http 保存")
}

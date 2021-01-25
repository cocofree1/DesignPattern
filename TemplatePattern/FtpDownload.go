package TemplatePattern

import "fmt"

type FtpDownload struct{}

// 返回模板实例
func NewFtpDowmload()Download{
	download := &FtpDownload{}
	return NewTemplateImplement(download)
}

// 重写对外实现接口的download and save
func (f* FtpDownload)Download(url string){
	fmt.Printf("ftp 下载 %s\n",url)
}

func (f* FtpDownload)Save(){
	fmt.Println("ftp 保存")
}

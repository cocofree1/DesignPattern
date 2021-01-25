package main

import "DesignPattern/TemplatePattern"

func main(){
	var download TemplatePattern.Download
	download = TemplatePattern.NewFtpDowmload()
	download.Download("http://www.baidu.com")
	download2 := TemplatePattern.NewHttpDownload()
	download2.Download("http://www.baidu.com")
}
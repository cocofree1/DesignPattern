package TemplatePattern

// 下载模板
type DownloadTemplate interface {
	Download(url string)
	Save()
}

// 对外的下载接口
type Download interface {
	Download(url string)
}

// 对外下载接口实现
type TemplateImplement struct {
	Template DownloadTemplate
}

func NewTemplateImplement(Template DownloadTemplate)*TemplateImplement{
	return &TemplateImplement{Template}
}

func (t* TemplateImplement)Download(url string){
	t.Template.Download(url)
	t.Template.Save()
}
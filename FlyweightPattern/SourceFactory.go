package FlyweightPattern

// 资源管理

var FactoryMap *SourceFactory  //是否重复创建的标志

type SourceFactory struct{
	sourceMap map[string]*SingleSource
}

// 单例模式，只申请一次空间
func GetFactoryMap()*SourceFactory{
	if FactoryMap == nil{
		FactoryMap = &SourceFactory{sourceMap: make(map[string]*SingleSource)}
	}
	return FactoryMap
}

func (s* SourceFactory)GetSourceByName(name string)*SingleSource{
	source := FactoryMap.sourceMap[name]
	if source == nil{
		source = NewSingleSource(name)
		FactoryMap.sourceMap[name] = source
 	}
 	return source
}
package models

// 新闻
type News struct {
	Id      int64  `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" gorm:"type:varchar(625)"` //标题
	Content string `json:"content" gorm:"type:text"`       //内容
	Image   string `json:"image" gorm:"type:varchar(625)"` //图片
}

// 谚语
type Proverb struct {
	Id         int64  `json:"id" gorm:"primaryKey"`
	Sentence   string `json:"sentence" gorm:"type:varchar(625)"`   //原本的句子
	Annotation string `json:"annotation" gorm:"type:varchar(625)"` //注解
}

// 农作物百科
type Crop struct {
	Id          int64  `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"type:varchar(625)"`  //名字
	Description string `json:"description" gorm:"type:text"`   //描述
	Image       string `json:"image" gorm:"type:varchar(625)"` //图片
}

// 短视频
type Video struct {
	Id      int64  `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" gorm:"type:varchar(625)"`   //标题
	Content string `json:"content" gorm:"type:varchar(625)"` //内容
	Image   string `json:"image" gorm:"type:varchar(625)"`   //图片
}

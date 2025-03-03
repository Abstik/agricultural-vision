package models

// 新闻
type News struct {
	Id      int64  `json:"-" gorm:"primaryKey"`
	Title   string `json:"title" gorm:"type:varchar(625)"` //标题
	Content string `json:"content" gorm:"type:text"`       //内容
	Image   string `json:"image" gorm:"type:varchar(625)"` //图片
}

// 谚语
type Proverb struct {
	Id         int64  `json:"-" gorm:"primaryKey"`
	Sentence   string `json:"sentence" gorm:"type:varchar(625)"`   //原本的句子
	Annotation string `json:"annotation" gorm:"type:varchar(625)"` //注解
}

// 农作物种类
type CropCategory struct {
	Id          int64        `json:"-" gorm:"primaryKey"`
	Category    string       `json:"category" gorm:"type:varchar(625)"`                      // 种类
	CropDetails []CropDetail `json:"crop_detail" gorm:"foreignKey:CategoryId;references:Id"` // 关联的农作物细节，外键指向 CropDetail 的 CategoryId
}

// 农作物细节
type CropDetail struct {
	Id           int64  `json:"-" gorm:"primaryKey"`
	CategoryId   int64  `json:"-"`                               // 外键
	Name         string `json:"name" gorm:"type:varchar(625)"`   // 名字
	Spell        string `json:"spell" gorm:"type:varchar(625)"`  // 拼音
	Icon         string `json:"icon" gorm:"type:varchar(625)"`   // 卡通图
	Description  string `json:"description" gorm:"type:text"`    // 描述
	Introduction string `json:"introduction" gorm:"type:text"`   // 简介
	Image1       string `json:"image1" gorm:"type:varchar(625)"` // 图片
	Image2       string `json:"image2" gorm:"type:varchar(625)"` // 图片
}

// 短视频
type Video struct {
	Id  int64  `json:"-" gorm:"primaryKey"`
	Url string `json:"url" gorm:"type:varchar(625)"` //视频链接
}

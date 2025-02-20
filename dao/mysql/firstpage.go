package mysql

import "agricultural_vision/models"

func GetNews() (news []models.News, err error) {
	if err = DB.Model(&models.News{}).Find(&news).Error; err != nil {
		return
	}
	return
}

func GetProverb() (proverbs []models.Proverb, err error) {
	if err = DB.Model(&models.Proverb{}).Find(&proverbs).Error; err != nil {
		return
	}
	return
}

func GetCrop() (crops []models.CropCategory, err error) {
	if err = DB.Preload("CropDetails").Find(&crops).Error; err != nil {
		return
	}
	return
}

func GetVideo() (videos []models.Video, err error) {
	if err = DB.Model(&models.Video{}).Find(&videos).Error; err != nil {
		return
	}
	return
}

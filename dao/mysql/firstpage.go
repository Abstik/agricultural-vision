package mysql

import "agricultural_vision/models"

func GetFirstPage() (*models.FirstPage, error) {
	firstPage := &models.FirstPage{}

	if err := DB.Model(&models.News{}).Find(&firstPage.News).Error; err != nil {
		return nil, err
	}
	if err := DB.Model(&models.Proverb{}).Find(&firstPage.Proverb).Error; err != nil {
		return nil, err
	}
	if err := DB.Model(&models.Crop{}).Find(&firstPage.Crop).Error; err != nil {
		return nil, err
	}
	if err := DB.Model(&models.Video{}).Find(&firstPage.Video).Error; err != nil {
		return nil, err
	}

	return firstPage, nil
}

/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: article
 * @Version: 1.0.0
 * @Date: 2023/3/14 22:43
 */

package model

type Article struct {
	*Model
	Title         string `json:"title"`
	Doc           string `json:"doc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

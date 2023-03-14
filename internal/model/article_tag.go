/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: article_tag
 * @Version: 1.0.0
 * @Date: 2023/3/14 22:58
 */

package model

//
// ArticleTag
//  @Description:
//
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

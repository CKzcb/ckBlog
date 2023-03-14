/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: tag
 * @Version: 1.0.0
 * @Date: 2023/3/14 22:42
 */

package model

//
// Tag
//  @Description: the class of article
//
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

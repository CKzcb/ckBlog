/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: model
 * @Version: 1.0.0
 * @Date: 2023/3/14 22:38
 */

package model

//
// Model
//  @Description: common model
//		all model reflected of db include it
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

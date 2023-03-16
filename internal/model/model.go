/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: model
 * @Version: 1.0.0
 * @Date: 2023/3/14 22:38
 */

package model

import (
	"fmt"
	"github.com/CKzcb/ckBlog/global"
	"github.com/CKzcb/ckBlog/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

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

func NewDBEngine(databaseSetting *setting.DataBaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBType,
		databaseSetting.Charset,
		databaseSetting.ParseTime)

	db, err := gorm.Open(databaseSetting.DBType, s)
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		// debug mode open db log
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConnes)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConnes)

	return db, nil
}

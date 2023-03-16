/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: pagination
 * @Version: 1.0.0
 * @Date: 2023/3/16 08:28
 */

package app

import (
	"github.com/CKzcb/ckBlog/global"
	"github.com/CKzcb/ckBlog/pkg/convert"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

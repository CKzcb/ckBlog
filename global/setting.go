/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: setting
 * @Version: 1.0.0
 * @Date: 2023/3/15 12:57
 */

package global

import (
	"github.com/CKzcb/ckBlog/pkg/logger"
	"github.com/CKzcb/ckBlog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DataBaseSettingS
	Logger          *logger.Logger
)

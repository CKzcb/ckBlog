/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: section
 * @Version: 1.0.0
 * @Date: 2023/3/15 12:51
 */

package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DataBaseSettingS struct {
	DBType        string
	UserName      string
	Password      string
	Host          string
	DBName        string
	TablePrefix   string
	Charset       string
	ParseTIme     bool
	MaxIdleConnes int
	MaxOpenConnes int
}

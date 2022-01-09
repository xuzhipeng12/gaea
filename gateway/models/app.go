/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/9/26 6:21 下午
 **/
package models

// 应用表
type App struct {
	IdField
	AppName    string `gorm:"type:varchar(128)" json:"app_name"`  // 应用名称
	AppType    string `gorm:"type:varchar(11)" json:"app_type"`   // script/charts
	FileName   string `gorm:"type:varchar(255)" json:"file_name"` // 脚本/charts名字
	Content    string `gorm:"type:text" json:"content"`           // 执行命令的内容
	Timeout    int    `gorm:"type:int" json:"timeout"`            // 超时时间
	RetryTimes int    `gorm:"type:int" json:"retry_times"`        //重试次数
	TimeField
}

func (App) TableName() string {
	return "gaea_app"
}

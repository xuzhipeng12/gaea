/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/1/9 6:22 下午
 **/
package models

// 子任务（job）表
type Job struct {
	IdField
	TaskId     int    `gorm:"type:int" json:"task_id"`            // task id
	HostName   string `gorm:"type:varchar(11)" json:"app_type"`   // script/charts
	FileName   string `gorm:"type:varchar(255)" json:"file_name"` // 脚本/charts名字
	Content    string `gorm:"type:text" json:"content"`           // 执行命令的内容
	Timeout    int    `gorm:"type:int" json:"timeout"`            // 超时时间
	RetryTimes int    `gorm:"type:int" json:"retry_times"`        // 重试次数
	StartEndTime
	TimeField
}

func (Job) TableName() string {
	return "gaea_job"
}

/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/1/9 6:22 下午
 **/
package models

// 任务（task）表
type Task struct {
	IdField
	AppId        int    `gorm:"type:int" json:"app_id"`                 // 应用id
	TaskName     string `gorm:"type:varchar(128)" json:"task_name"`     // 任务名字
	TaskType     string `gorm:"type:varchar(11)" json:"task_type"`      // script/charts
	TagsId       int    `gorm:"type:int" json:"tags_id"`                // 执行对象id
	TaskState    string `gorm:"type:text" json:"task_state"`            // 执行命令的内容
	ExecuteTimes int    `gorm:"type:varchar(128)" json:"execute_times"` // 超时时间
	StartEndTime
	TimeField
}

func (Task) TableName() string {
	return "gaea_task"
}

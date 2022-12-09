// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ScheduleJobLog is the golang structure for table schedule_job_log.
type ScheduleJobLog struct {
	LogId      int64       `json:"logId"      description:"任务日志id"`
	JobId      int64       `json:"jobId"      description:"任务id"`
	BeanName   string      `json:"beanName"   description:"spring bean名称"`
	MethodName string      `json:"methodName" description:"方法名"`
	Params     string      `json:"params"     description:"参数"`
	Status     int         `json:"status"     description:"任务执行结果"`
	Error      string      `json:"error"      description:"异常信息"`
	Times      int         `json:"times"      description:"耗时（单位：毫秒）"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
}

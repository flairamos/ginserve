package utils

import "time"

// 抓拍记录

type VideoCaptureRecord struct {
	RecId           string    `json:"recId"`           // 时间唯一标识
	IndexCode       string    `json:"indexCode"`       // 监控点唯一标识
	IntelligentType string    `json:"intelligentType"` // 报警类型
	FinishedTime    time.Time `json:"finishedTime"`
	PlateInfo       string    `json:"plateInfo"`      // 车牌号
	PlateType       string    `json:"plate_type"`     // 车牌类型
	PicUrl          string    `json:"picUrl"`         // 报警图片
	Pic             []byte    `json:"pic"`            // 报警图片
	AlarmPlace      string    `json:"alarmPlace"`     // 报警位置
	AlarmSource     string    `json:"alarmSource"`    // 报警来源，0场站  2道路
	ViolationAreas  string    `json:"violationAreas"` // 报警区域
	Confidence      float64   `json:"confidence"`     // 置信度
	Audited         float64   `json:"audited"`        // 审核状态
	CollectTime     time.Time `json:"collectTime"`    // 抓拍时间
	AuditedUser     string    `json:"auditedUser"`    // 审核人
	OrgCode         string    `json:"orgCode"`        // 区域id
	OrgPath         string    `json:"orgPath"`        // 区域path
	Comment         string    `json:"comment"`        // 备注
	Lon             float64   `json:"lon"`            // 经度
	Lat             float64   `json:"lat"`            // 纬度
	Geom            string    `json:"geom"`           // 空间位置
	RecordSource    string    `json:"recordSource"`   // 记录来源
}

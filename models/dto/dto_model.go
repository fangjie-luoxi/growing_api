package dto

import "time"

// GrIntegralR  积分记录
type GrIntegralR struct {
	Id        int        /*字段名:"id"*/
	CreatedAt *time.Time /*字段名:"created_at"，默认值:"CURRENT_TIMESTAMP"*/
	UpdatedAt *time.Time /*字段名:"updated_at"，默认值:"CURRENT_TIMESTAMP"*/
	DeletedAt *time.Time /*字段名:"deleted_at"*/
	UserId    int        /*字段名:"user_id"*/
	InType    string     /*字段名:"in_type"，长度:"10"，默认值:"i"，注释:"类型 enum:i:增加,o:扣除#"*/
	Num       float64    /*字段名:"num"，注释:"增加或扣除的积分"*/
	ReTpye    string     /*字段名:"re_tpye"，长度:"10"，注释:"关联(积分变动所关联的项) enum:tt:目标,tk:任务,re:规则#"*/
	ReId      int        /*字段名:"re_id"，注释:"关联所对应的id"*/
	Desc      string     /*字段名:"desc"，长度:"255"，注释:"描述"*/
	User      *User      /*说明:与[user] 一对多 关系，该表为从表，该表user_id为关联字段 */
}

/* [gr_integral_r] json temp
{
    "Id": 0,
    "CreatedAt": "2020-01-02T03:04:05+08:00",
    "UpdatedAt": "2020-01-02T03:04:05+08:00",
    "DeletedAt": "2020-01-02T03:04:05+08:00",
    "UserId": 0,
    "InType": "",
    "Num": 0.00,
    "ReTpye": "",
    "ReId": 0,
    "Desc": "",
    "User": {"Id": 0}
}
*/

// GrRule  规则
type GrRule struct {
	Id            int             /*字段名:"id"*/
	CreatedAt     *time.Time      /*字段名:"created_at"，默认值:"CURRENT_TIMESTAMP"*/
	UpdatedAt     *time.Time      /*字段名:"updated_at"，默认值:"CURRENT_TIMESTAMP"*/
	DeletedAt     *time.Time      /*字段名:"deleted_at"*/
	UserId        int             /*字段名:"user_id"*/
	ReName        string          /*字段名:"re_name"，长度:"45"，注释:"规则名称"*/
	Content       string          /*字段名:"content"，长度:"255"，注释:"规则内容"*/
	InType        string          /*字段名:"in_type"，长度:"10"，默认值:"i"，注释:"类型 enum:i:增加,o:扣除#"*/
	Num           float64         /*字段名:"num"，注释:"积分"*/
	Rm            string          /*字段名:"rm"，长度:"1000"，注释:"备注"*/
	User          *User           /*说明:与[user] 一对多 关系，该表为从表，该表user_id为关联字段 */
	GrRuleRecords []*GrRuleRecord /*说明:"与[gr_rule_record] 一对多 关系，该表为主表，该表id为关联字段。"，*/
}

/* [gr_rule] json temp
{
    "Id": 0,
    "CreatedAt": "2020-01-02T03:04:05+08:00",
    "UpdatedAt": "2020-01-02T03:04:05+08:00",
    "DeletedAt": "2020-01-02T03:04:05+08:00",
    "UserId": 0,
    "ReName": "",
    "Content": "",
    "InType": "",
    "Num": 0.00,
    "Rm": "",
    "User": {"Id": 0}
    "GrRuleRecords": [{"Id": 0}]
}
*/

// GrRuleRecord  规则记录表
type GrRuleRecord struct {
	Id        int        /*字段名:"id"*/
	CreatedAt *time.Time /*字段名:"created_at"，默认值:"CURRENT_TIMESTAMP"*/
	UpdatedAt *time.Time /*字段名:"updated_at"，默认值:"CURRENT_TIMESTAMP"*/
	DeletedAt *time.Time /*字段名:"deleted_at"*/
	GrRuleId  int        /*字段名:"gr_rule_id"*/
	Date      *time.Time /*字段名:"date"*/
	GrRule    *GrRule    /*说明:与[gr_rule] 一对多 关系，该表为从表，该表gr_rule_id为关联字段 */
}

/* [gr_rule_record] json temp
{
    "Id": 0,
    "CreatedAt": "2020-01-02T03:04:05+08:00",
    "UpdatedAt": "2020-01-02T03:04:05+08:00",
    "DeletedAt": "2020-01-02T03:04:05+08:00",
    "GrRuleId": 0,
    "Date": "2020-01-02T03:04:05+08:00",
    "GrRule": {"Id": 0}
}
*/

// GrTarget  目标表
type GrTarget struct {
	Id        int        /*字段名:"id"*/
	CreatedAt *time.Time /*字段名:"created_at"，默认值:"CURRENT_TIMESTAMP"*/
	UpdatedAt *time.Time /*字段名:"updated_at"，默认值:"CURRENT_TIMESTAMP"*/
	DeletedAt *time.Time /*字段名:"deleted_at"*/
	UserId    int        /*字段名:"user_id"*/
	TtTitle   string     /*字段名:"tt_title"，长度:"45"，注释:"目标标题"*/
	TtContent string     /*字段名:"tt_content"，长度:"255"，注释:"目标内容"*/
	TtType    string     /*字段名:"tt_type"，长度:"45"，注释:"目标类型 enum:y:年,m:月,d:日#"*/
	TtNum     float64    /*字段名:"tt_num"，注释:"目标数量"*/
	TtUnit    string     /*字段名:"tt_unit"，长度:"45"，注释:"目标单位"*/
	Begin     *time.Time /*字段名:"begin"，注释:"开始时间"*/
	End       *time.Time /*字段名:"end"，注释:"结束时间"*/
	Status    string     /*字段名:"status"，长度:"45"，注释:"状态 enum:s:完成,n:未完成,r:进行中#"*/
	Num       float64    /*字段名:"num"，注释:"积分"*/
	Rm        string     /*字段名:"rm"，长度:"255"，注释:"备注"*/
	GenTask   string     /*字段名:"gen_task"，长度:"10"，默认值:"y"，注释:"是否生成任务 enum:y:生成,n:不生成#"*/
	Finish    float64    /*字段名:"finish"，注释:"完成数量"*/
	User      *User      /*说明:与[user] 一对多 关系，该表为从表，该表user_id为关联字段 */
	GrTasks   []*GrTask  /*说明:"与[gr_task] 一对多 关系，该表为主表，该表id为关联字段。"，*/
}

/* [gr_target] json temp
{
    "Id": 0,
    "CreatedAt": "2020-01-02T03:04:05+08:00",
    "UpdatedAt": "2020-01-02T03:04:05+08:00",
    "DeletedAt": "2020-01-02T03:04:05+08:00",
    "UserId": 0,
    "TtTitle": "",
    "TtContent": "",
    "TtType": "",
    "TtNum": 0.00,
    "TtUnit": "",
    "Begin": "2020-01-02T03:04:05+08:00",
    "End": "2020-01-02T03:04:05+08:00",
    "Status": "",
    "Num": 0.00,
    "Rm": "",
    "GenTask": "",
    "Finish": 0.00,
    "User": {"Id": 0}
    "GrTasks": [{"Id": 0}]
}
*/

// GrTask  任务表
type GrTask struct {
	Id         int        /*字段名:"id"*/
	CreatedAt  *time.Time /*字段名:"created_at"，默认值:"CURRENT_TIMESTAMP"*/
	UpdatedAt  *time.Time /*字段名:"updated_at"，默认值:"CURRENT_TIMESTAMP"*/
	DeletedAt  *time.Time /*字段名:"deleted_at"*/
	UserId     int        /*字段名:"user_id"*/
	GrTargetId int        /*字段名:"gr_target_id"，默认值:"0"*/
	TkTitle    string     /*字段名:"tk_title"，长度:"45"，注释:"任务标题"*/
	TkContent  string     /*字段名:"tk_content"，长度:"255"，注释:"任务内容"*/
	TtNum      float64    /*字段名:"tt_num"，注释:"任务数量"*/
	TtUnit     string     /*字段名:"tt_unit"，长度:"45"，注释:"任务单位"*/
	Num        float64    /*字段名:"num"，注释:"积分"*/
	Rm         string     /*字段名:"rm"，长度:"255"，注释:"备注"*/
	Status     string     /*字段名:"status"，长度:"45"，默认值:"b"，注释:"状态 enum:s:完成,n:未完成,r:进行中,b:未开始#"*/
	Date       *time.Time /*字段名:"date"，注释:"任务日期"*/
	User       *User      /*说明:与[user] 一对多 关系，该表为从表，该表user_id为关联字段 */
	GrTarget   *GrTarget  /*说明:与[gr_target] 一对多 关系，该表为从表，该表gr_target_id为关联字段 */
}

/* [gr_task] json temp
{
    "Id": 0,
    "CreatedAt": "2020-01-02T03:04:05+08:00",
    "UpdatedAt": "2020-01-02T03:04:05+08:00",
    "DeletedAt": "2020-01-02T03:04:05+08:00",
    "UserId": 0,
    "GrTargetId": 0,
    "TkTitle": "",
    "TkContent": "",
    "TtNum": 0.00,
    "TtUnit": "",
    "Num": 0.00,
    "Rm": "",
    "Status": "",
    "Date": "2020-01-02T03:04:05+08:00",
    "User": {"Id": 0}
    "GrTarget": {"Id": 0}
}
*/

// User  用户表
type User struct {
	Id           int            /*字段名:"id"*/
	CreatedAt    *time.Time     /*字段名:"created_at"，默认值:"CURRENT_TIMESTAMP"*/
	UpdatedAt    *time.Time     /*字段名:"updated_at"，默认值:"CURRENT_TIMESTAMP"*/
	DeletedAt    *time.Time     /*字段名:"deleted_at"*/
	Integral     float64        /*字段名:"integral"*/
	GrIntegralRs []*GrIntegralR /*说明:"与[gr_integral_r] 一对多 关系，该表为主表，该表id为关联字段。"，*/
	GrRules      []*GrRule      /*说明:"与[gr_rule] 一对多 关系，该表为主表，该表id为关联字段。"，*/
	GrTargets    []*GrTarget    /*说明:"与[gr_target] 一对多 关系，该表为主表，该表id为关联字段。"，*/
	GrTasks      []*GrTask      /*说明:"与[gr_task] 一对多 关系，该表为主表，该表id为关联字段。"，*/
}

/* [user] json temp
{
    "Id": 0,
    "CreatedAt": "2020-01-02T03:04:05+08:00",
    "UpdatedAt": "2020-01-02T03:04:05+08:00",
    "DeletedAt": "2020-01-02T03:04:05+08:00",
    "Integral": 0.00,
    "GrIntegralRs": [{"Id": 0}]
    "GrRules": [{"Id": 0}]
    "GrTargets": [{"Id": 0}]
    "GrTasks": [{"Id": 0}]
}
*/

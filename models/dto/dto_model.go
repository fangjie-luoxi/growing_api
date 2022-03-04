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
	Num       int        /*字段名:"num"，注释:"增加或扣除的积分"*/
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
    "Num": 0,
    "ReTpye": "",
    "ReId": 0,
    "Desc": "",
    "User": {"Id": 0}
}
*/

// GrRule  规则
type GrRule struct {
	Id        int        /*字段名:"id"*/
	CreatedAt *time.Time /*字段名:"created_at"，默认值:"CURRENT_TIMESTAMP"*/
	UpdatedAt *time.Time /*字段名:"updated_at"，默认值:"CURRENT_TIMESTAMP"*/
	DeletedAt *time.Time /*字段名:"deleted_at"*/
	UserId    int        /*字段名:"user_id"*/
	ReName    string     /*字段名:"re_name"，长度:"45"，注释:"规则名称"*/
	Content   string     /*字段名:"content"，长度:"255"，注释:"规则内容"*/
	InType    string     /*字段名:"in_type"，长度:"10"，默认值:"i"，注释:"类型 enum:i:增加,o:扣除#"*/
	Num       int        /*字段名:"num"，注释:"积分"*/
	Rm        string     /*字段名:"rm"，长度:"1000"，注释:"备注"*/
	User      *User      /*说明:与[user] 一对多 关系，该表为从表，该表user_id为关联字段 */
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
    "Num": 0,
    "Rm": "",
    "User": {"Id": 0}
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
	TtNum     int64      /*字段名:"tt_num"，注释:"目标数量"*/
	TtUnit    string     /*字段名:"tt_unit"，长度:"45"，注释:"目标单位"*/
	Begin     *time.Time /*字段名:"begin"，注释:"开始时间"*/
	End       *time.Time /*字段名:"end"，注释:"结束时间"*/
	Status    string     /*字段名:"status"，长度:"45"，注释:"状态 enum:s:完成,n:未完成,r:进行中#"*/
	Num       int        /*字段名:"num"，注释:"积分"*/
	Rm        string     /*字段名:"rm"，长度:"255"，注释:"备注"*/
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
    "TtNum": 0,
    "TtUnit": "",
    "Begin": "2020-01-02T03:04:05+08:00",
    "End": "2020-01-02T03:04:05+08:00",
    "Status": "",
    "Num": 0,
    "Rm": "",
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
	Date       *time.Time /*字段名:"date"，注释:"任务时间"*/
	TkTitle    string     /*字段名:"tk_title"，长度:"45"，注释:"任务标题"*/
	TkContent  string     /*字段名:"tk_content"，长度:"255"，注释:"任务内容"*/
	TtNum      int        /*字段名:"tt_num"，注释:"任务数量"*/
	TtUnit     string     /*字段名:"tt_unit"，长度:"45"，注释:"任务单位"*/
	Num        int        /*字段名:"num"，注释:"积分"*/
	Rm         string     /*字段名:"rm"，长度:"255"，注释:"备注"*/
	Status     string     /*字段名:"status"，长度:"45"，注释:"状态 enum:s:完成,n:未完成,r:进行中#"*/
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
    "Date": "2020-01-02T03:04:05+08:00",
    "TkTitle": "",
    "TkContent": "",
    "TtNum": 0,
    "TtUnit": "",
    "Num": 0,
    "Rm": "",
    "Status": "",
    "User": {"Id": 0}
    "GrTarget": {"Id": 0}
}
*/

type ThirdAuth struct {
	Id        int        /*字段名:"id"*/
	CreatedAt *time.Time /*字段名:"created_at"*/
	UpdatedAt *time.Time /*字段名:"updated_at"*/
	DeletedAt *time.Time /*字段名:"deleted_at"*/
	AuthType  string     /*字段名:"auth_type"*/
	Openid    string     /*字段名:"openid"，长度:"191"*/
	Unionid   string     /*字段名:"unionid"，长度:"191"*/
	Nickname  string     /*字段名:"nickname"*/
	Avatar    string     /*字段名:"avatar"*/
	UserId    int64      /*字段名:"user_id"*/
	User      *User      /*说明:与[user] 一对多 关系，该表为从表，该表user_id为关联字段 */
}

/* [third_auth] json temp
{
    "Id": 0,
    "CreatedAt": "2020-01-02T03:04:05+08:00",
    "UpdatedAt": "2020-01-02T03:04:05+08:00",
    "DeletedAt": "2020-01-02T03:04:05+08:00",
    "AuthType": "",
    "Openid": "",
    "Unionid": "",
    "Nickname": "",
    "Avatar": "",
    "UserId": 0,
    "User": {"Id": 0}
}
*/

// User  用户表
type User struct {
	Id            int            /*字段名:"id"*/
	CreatedAt     *time.Time     /*字段名:"created_at"*/
	UpdatedAt     *time.Time     /*字段名:"updated_at"*/
	DeletedAt     *time.Time     /*字段名:"deleted_at"*/
	Integral      int64          /*字段名:"integral"*/
	SysOrgId      int            /*字段名:"sys_org_id"*/
	LoginCode     string         /*字段名:"login_code"，长度:"45"*/
	Password      string         /*字段名:"password"，长度:"255"*/
	UserCode      string         /*字段名:"user_code"，长度:"45"*/
	UserName      string         /*字段名:"user_name"，长度:"45"*/
	Phone         string         /*字段名:"phone"，长度:"45"*/
	Avatar        string         /*字段名:"avatar"，长度:"255"*/
	Email         string         /*字段名:"email"，长度:"255"*/
	UserType      string         /*字段名:"user_type"，长度:"45"*/
	Status        string         /*字段名:"status"，长度:"45"，默认值:"r"*/
	Birth         *time.Time     /*字段名:"birth"*/
	Sex           string         /*字段名:"sex"，长度:"45"*/
	CardNo        string         /*字段名:"card_no"，长度:"45"*/
	LastLoginTime *time.Time     /*字段名:"last_login_time"*/
	Address       string         /*字段名:"address"，长度:"255"*/
	Profile       string         /*字段名:"profile"*/
	SysDeptId     int64          /*字段名:"sys_dept_id"，默认值:"0"*/
	SysPostId     int64          /*字段名:"sys_post_id"，默认值:"0"*/
	GrIntegralRs  []*GrIntegralR /*说明:"与[gr_integral_r] 一对多 关系，该表为主表，该表id为关联字段。"，*/
	GrRules       []*GrRule      /*说明:"与[gr_rule] 一对多 关系，该表为主表，该表id为关联字段。"，*/
	GrTargets     []*GrTarget    /*说明:"与[gr_target] 一对多 关系，该表为主表，该表id为关联字段。"，*/
	GrTasks       []*GrTask      /*说明:"与[gr_task] 一对多 关系，该表为主表，该表id为关联字段。"，*/
	ThirdAuths    []*ThirdAuth   /*说明:"与[third_auth] 一对多 关系，该表为主表，该表id为关联字段。"，*/
}

/* [user] json temp
{
    "Id": 0,
    "CreatedAt": "2020-01-02T03:04:05+08:00",
    "UpdatedAt": "2020-01-02T03:04:05+08:00",
    "DeletedAt": "2020-01-02T03:04:05+08:00",
    "Integral": 0,
    "SysOrgId": 0,
    "LoginCode": "",
    "Password": "",
    "UserCode": "",
    "UserName": "",
    "Phone": "",
    "Avatar": "",
    "Email": "",
    "UserType": "",
    "Status": "",
    "Birth": "2020-01-02T03:04:05+08:00",
    "Sex": "",
    "CardNo": "",
    "LastLoginTime": "2020-01-02T03:04:05+08:00",
    "Address": "",
    "Profile": "",
    "SysDeptId": 0,
    "SysPostId": 0,
    "GrIntegralRs": [{"Id": 0}]
    "GrRules": [{"Id": 0}]
    "GrTargets": [{"Id": 0}]
    "GrTasks": [{"Id": 0}]
    "ThirdAuths": [{"Id": 0}]
}
*/

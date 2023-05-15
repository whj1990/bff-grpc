package stru

type ItemListParam struct {
	PageNum  int32  `form:"pageNum" json:"pageNum"  validate:"required"`
	PageSize int32  `form:"pageSize" json:"pageSize" validate:"required"`
	Id       int32  `form:"id" json:"id"`
	Title    string `form:"title" json:"title"`
	ItemId   string `form:"itemId" json:"itemId"`
	Status   int8   `form:"status" json:"status"`
}
type IdParam struct {
	Id int64 `form:"id" json:"id"`
}
type StrIdParam struct {
	Id string `form:"id" json:"id"`
}
type ChatBotsParam struct {
	Mode   string `form:"mode" json:"mode"`
	Prompt string `form:"prompt" json:"prompt"`
}
type ResourceBaseListParams struct {
	Id     int64  `form:"id" json:"id"`
	Code   string `form:"code" json:"code"`
	Status int8   `form:"status" json:"status"`
	Limit  int32  `form:"limit" json:"limit" validate:"required"`
}

type ResourceOtsData struct {
	Id        int64  `form:"id" json:"id"`
	Code      string `form:"code" json:"code" validate:"max=30"`
	Desc      string `form:"desc" json:"desc" validate:"max=100"`
	Salt      string `form:"salt" json:"salt"`
	Endpoint  string `form:"endpoint" json:"endpoint" validate:"required"`
	Instance  string `form:"instance" json:"instance" validate:"required"`
	AccessKey string `form:"accessKey" json:"accessKey" validate:"required"`
	SecretKey string `form:"secretKey" json:"secretKey" validate:"required"`
	Status    int8   `form:"status" json:"status" validate:"required"`
}
type ResourceRedisData struct {
	Id        int64  `form:"id,1" json:"id"`
	Code      string `form:"code" json:"code" validate:"max=30"`
	Desc      string `form:"desc" json:"desc" validate:"max=100"`
	Salt      string `form:"salt" json:"salt"`
	Host      string `form:"host" json:"host" validate:"required"`
	Port      int32  `form:"port" json:"port" validate:"required"`
	Password  string `form:"password" json:"password"`
	Database  int32  `form:"database" json:"database"`
	MinIdle   int32  `form:"minIdle" json:"minIdle"`
	MaxIdle   int32  `form:"maxIdle" json:"maxIdle"`
	MaxActive int32  `form:"maxActive" json:"maxActive"`
	Timeout   int32  `form:"timeout" json:"timeout"`
	Status    int8   `form:"status" json:"status"`
}
type ResourceOssData struct {
	Id        int64  `form:"id" json:"id"`
	Code      string `form:"code" json:"code" validate:"max=30"`
	Desc      string `form:"desc" json:"desc" validate:"max=100"`
	Salt      string `form:"salt" json:"salt"`
	Endpoint  string `form:"endpoint" json:"endpoint" validate:"required"`
	AccessKey string `form:"accessKey" json:"accessKey" validate:"required"`
	SecretKey string `form:"secretKey" json:"secretKey" validate:"required"`
	Status    int8   `form:"status" json:"status" validate:"required"`
}
type SwitchModeParams struct {
	Code string `form:"code" json:"code"`
}
type RoleData struct {
	Id     int64  `form:"id" json:"id"`
	Code   string `form:"code" json:"code" validate:"max=30"`
	Desc   string `form:"desc" json:"desc" validate:"max=100"`
	Status int8   `form:"status" json:"status" validate:"required"`
}

type RoleUserListParams struct {
	Id     int64  `form:"id" json:"id"`
	RrCode string `form:"rrCode" json:"rrCode"`
	Uid    string `form:"uid" json:"uid"`
	Limit  int32  `form:"limit" json:"limit"`
}
type RoleUserData struct {
	Id     int64  `form:"id" json:"id"`
	RrCode string `form:"rrCode" json:"rrCode" validate:"required,max=30"`
	Uid    string `form:"uid" json:"uid" validate:"required"`
}
type ExcelParams struct {
	RrCode string `form:"rrCode" json:"rrCode" validate:"required"`
	Data   []byte `form:"data" json:"data" validate:"required"`
}
type DictDomainListParams struct {
	Id       int64  `form:"id" json:"id"`
	Code     string `form:"code" json:"code"`
	JumpType int8   `form:"jumpType" json:"jumpType"`
	Status   int8   `form:"status" json:"status"`
	Limit    int32  `form:"limit" json:"limit" validate:"required"`
}

type DictSortListParams struct {
	Id     int64  `form:"id" json:"id"`
	Code   string `form:"code" json:"code"`
	Algo   int8   `form:"algo" json:"algo"`
	Status int8   `form:"status" json:"status"`
	Limit  int32  `form:"limit" json:"limit" validate:"required"`
}

type ReviewProjectListParams struct {
	PageNum    int32  `form:"pageNum" json:"pageNum" validate:"required"`
	PageSize   int32  `form:"pageSize" json:"pageSize" validate:"required"`
	Id         int64  `form:"id" json:"id"`
	Name       string `form:"name" json:"name"`
	ModeCode   string `form:"modeCode" json:"modeCode"`
	Status     int32  `form:"status" json:"status"`
	ShowStatus int32  `form:"showStatus" json:"showStatus"`
}

type IdsInt64Params struct {
	Ids []int64 `form:"ids[]" json:"ids[]"`
}
type ReviewProjectSaveData struct {
	Id          int64  `form:"id" json:"id"`
	Name        string `form:"name" json:"name" validate:"required"`
	ModeCode    string `form:"modeCode" json:"modeCode" validate:"required"`
	Description string `form:"description" json:"description"`
}

type ReviewProjectSaveParam struct {
	Action string                `form:"action" json:"action"  validate:"required"`
	Data   ReviewProjectSaveData `form:"data" json:"data"  validate:"required,dive,required"`
}

type StatusParam struct {
	Id     int64  `form:"id" json:"id"`
	Status int8   `form:"status" json:"status"`
	Remark string `form:"remark" json:"remark"`
}
type DropDownListParams struct {
	PageNum  int32  `form:"pageNum" json:"pageNum" validate:"required"`
	PageSize int32  `form:"pageSize" json:"pageSize" validate:"required"`
	Id       int64  `form:"id" json:"id"`
	Code     string `form:"code" json:"code"`
	Status   int8   `form:"status" json:"status"`
}
type DropDownSaveParam struct {
	Action string            `form:"action" json:"action" validate:"required"`
	Data   *DropDownSaveData `form:"data" json:"data" validate:"required,dive,required"`
}
type DropDownSaveData struct {
	Id      int64  `form:"id" json:"id"`
	Name    string `form:"name" json:"name" validate:"required"`
	Code    string `form:"code" json:"code" validate:"required"`
	Key     string `form:"key" json:"key" validate:"required"`
	KeyType int8   `form:"key_type" json:"key_type" validate:"required"`
	Value   string `form:"value" json:"value" validate:"required"`
	Desc    string `form:"desc" json:"desc"`
}

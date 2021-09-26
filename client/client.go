// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchDeleteDataLimitResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteDataLimitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDataLimitResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteDataLimitResponseBody) SetRequestId(v string) *BatchDeleteDataLimitResponseBody {
	s.RequestId = &v
	return s
}

type BatchDeleteDataLimitResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDeleteDataLimitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteDataLimitResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDataLimitResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteDataLimitResponse) SetHeaders(v map[string]*string) *BatchDeleteDataLimitResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteDataLimitResponse) SetBody(v *BatchDeleteDataLimitResponseBody) *BatchDeleteDataLimitResponse {
	s.Body = v
	return s
}

type CreateConfigRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigRequest) SetLang(v string) *CreateConfigRequest {
	s.Lang = &v
	return s
}

func (s *CreateConfigRequest) SetCode(v string) *CreateConfigRequest {
	s.Code = &v
	return s
}

func (s *CreateConfigRequest) SetDescription(v string) *CreateConfigRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigRequest) SetValue(v string) *CreateConfigRequest {
	s.Value = &v
	return s
}

type CreateConfigResponseBody struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigResponseBody) SetId(v int64) *CreateConfigResponseBody {
	s.Id = &v
	return s
}

func (s *CreateConfigResponseBody) SetRequestId(v string) *CreateConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateConfigResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigResponse) SetHeaders(v map[string]*string) *CreateConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigResponse) SetBody(v *CreateConfigResponseBody) *CreateConfigResponse {
	s.Body = v
	return s
}

type CreateDataLimitRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceType    *int32  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	ParentId        *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	AuditStatus     *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AutoScan        *int32  `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	LogStoreDay     *int32  `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	EngineType      *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	OcrStatus       *int32  `json:"OcrStatus,omitempty" xml:"OcrStatus,omitempty"`
	EventStatus     *int32  `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
}

func (s CreateDataLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataLimitRequest) GoString() string {
	return s.String()
}

func (s *CreateDataLimitRequest) SetLang(v string) *CreateDataLimitRequest {
	s.Lang = &v
	return s
}

func (s *CreateDataLimitRequest) SetResourceType(v int32) *CreateDataLimitRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateDataLimitRequest) SetServiceRegionId(v string) *CreateDataLimitRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *CreateDataLimitRequest) SetParentId(v string) *CreateDataLimitRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDataLimitRequest) SetUserName(v string) *CreateDataLimitRequest {
	s.UserName = &v
	return s
}

func (s *CreateDataLimitRequest) SetPassword(v string) *CreateDataLimitRequest {
	s.Password = &v
	return s
}

func (s *CreateDataLimitRequest) SetAuditStatus(v int32) *CreateDataLimitRequest {
	s.AuditStatus = &v
	return s
}

func (s *CreateDataLimitRequest) SetAutoScan(v int32) *CreateDataLimitRequest {
	s.AutoScan = &v
	return s
}

func (s *CreateDataLimitRequest) SetLogStoreDay(v int32) *CreateDataLimitRequest {
	s.LogStoreDay = &v
	return s
}

func (s *CreateDataLimitRequest) SetEngineType(v string) *CreateDataLimitRequest {
	s.EngineType = &v
	return s
}

func (s *CreateDataLimitRequest) SetPort(v int32) *CreateDataLimitRequest {
	s.Port = &v
	return s
}

func (s *CreateDataLimitRequest) SetOcrStatus(v int32) *CreateDataLimitRequest {
	s.OcrStatus = &v
	return s
}

func (s *CreateDataLimitRequest) SetEventStatus(v int32) *CreateDataLimitRequest {
	s.EventStatus = &v
	return s
}

type CreateDataLimitResponseBody struct {
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataLimitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataLimitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataLimitResponseBody) SetId(v int32) *CreateDataLimitResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataLimitResponseBody) SetRequestId(v string) *CreateDataLimitResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataLimitResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDataLimitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataLimitResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataLimitResponse) GoString() string {
	return s.String()
}

func (s *CreateDataLimitResponse) SetHeaders(v map[string]*string) *CreateDataLimitResponse {
	s.Headers = v
	return s
}

func (s *CreateDataLimitResponse) SetBody(v *CreateDataLimitResponseBody) *CreateDataLimitResponse {
	s.Body = v
	return s
}

type CreateRuleRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Category        *int32  `json:"Category,omitempty" xml:"Category,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	StatExpress     *string `json:"StatExpress,omitempty" xml:"StatExpress,omitempty"`
	RiskLevelId     *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RuleType        *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ProductCode     *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductId       *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	WarnLevel       *int32  `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ContentCategory *int32  `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Target          *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) SetLang(v string) *CreateRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateRuleRequest) SetName(v string) *CreateRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateRuleRequest) SetCategory(v int32) *CreateRuleRequest {
	s.Category = &v
	return s
}

func (s *CreateRuleRequest) SetContent(v string) *CreateRuleRequest {
	s.Content = &v
	return s
}

func (s *CreateRuleRequest) SetStatExpress(v string) *CreateRuleRequest {
	s.StatExpress = &v
	return s
}

func (s *CreateRuleRequest) SetRiskLevelId(v int64) *CreateRuleRequest {
	s.RiskLevelId = &v
	return s
}

func (s *CreateRuleRequest) SetRuleType(v int32) *CreateRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateRuleRequest) SetProductCode(v string) *CreateRuleRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateRuleRequest) SetProductId(v int64) *CreateRuleRequest {
	s.ProductId = &v
	return s
}

func (s *CreateRuleRequest) SetWarnLevel(v int32) *CreateRuleRequest {
	s.WarnLevel = &v
	return s
}

func (s *CreateRuleRequest) SetDescription(v string) *CreateRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateRuleRequest) SetContentCategory(v int32) *CreateRuleRequest {
	s.ContentCategory = &v
	return s
}

func (s *CreateRuleRequest) SetStatus(v int32) *CreateRuleRequest {
	s.Status = &v
	return s
}

func (s *CreateRuleRequest) SetTarget(v string) *CreateRuleRequest {
	s.Target = &v
	return s
}

type CreateRuleResponseBody struct {
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) SetId(v int32) *CreateRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRuleResponse) SetHeaders(v map[string]*string) *CreateRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRuleResponse) SetBody(v *CreateRuleResponseBody) *CreateRuleResponse {
	s.Body = v
	return s
}

type CreateScanTaskRequest struct {
	DataLimitId      *int64  `json:"DataLimitId,omitempty" xml:"DataLimitId,omitempty"`
	IntervalDay      *int32  `json:"IntervalDay,omitempty" xml:"IntervalDay,omitempty"`
	RunHour          *int32  `json:"RunHour,omitempty" xml:"RunHour,omitempty"`
	RunMinute        *int32  `json:"RunMinute,omitempty" xml:"RunMinute,omitempty"`
	TaskName         *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskUserName     *string `json:"TaskUserName,omitempty" xml:"TaskUserName,omitempty"`
	OssScanPath      *string `json:"OssScanPath,omitempty" xml:"OssScanPath,omitempty"`
	ScanRange        *int32  `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
	ScanRangeContent *string `json:"ScanRangeContent,omitempty" xml:"ScanRangeContent,omitempty"`
	ResourceType     *int64  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateScanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScanTaskRequest) SetDataLimitId(v int64) *CreateScanTaskRequest {
	s.DataLimitId = &v
	return s
}

func (s *CreateScanTaskRequest) SetIntervalDay(v int32) *CreateScanTaskRequest {
	s.IntervalDay = &v
	return s
}

func (s *CreateScanTaskRequest) SetRunHour(v int32) *CreateScanTaskRequest {
	s.RunHour = &v
	return s
}

func (s *CreateScanTaskRequest) SetRunMinute(v int32) *CreateScanTaskRequest {
	s.RunMinute = &v
	return s
}

func (s *CreateScanTaskRequest) SetTaskName(v string) *CreateScanTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateScanTaskRequest) SetTaskUserName(v string) *CreateScanTaskRequest {
	s.TaskUserName = &v
	return s
}

func (s *CreateScanTaskRequest) SetOssScanPath(v string) *CreateScanTaskRequest {
	s.OssScanPath = &v
	return s
}

func (s *CreateScanTaskRequest) SetScanRange(v int32) *CreateScanTaskRequest {
	s.ScanRange = &v
	return s
}

func (s *CreateScanTaskRequest) SetScanRangeContent(v string) *CreateScanTaskRequest {
	s.ScanRangeContent = &v
	return s
}

func (s *CreateScanTaskRequest) SetResourceType(v int64) *CreateScanTaskRequest {
	s.ResourceType = &v
	return s
}

type CreateScanTaskResponseBody struct {
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateScanTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScanTaskResponseBody) SetId(v int32) *CreateScanTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateScanTaskResponseBody) SetRequestId(v string) *CreateScanTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateScanTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScanTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateScanTaskResponse) SetHeaders(v map[string]*string) *CreateScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateScanTaskResponse) SetBody(v *CreateScanTaskResponseBody) *CreateScanTaskResponse {
	s.Body = v
	return s
}

type DeleteDataLimitRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDataLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataLimitRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLimitRequest) SetLang(v string) *DeleteDataLimitRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDataLimitRequest) SetId(v int64) *DeleteDataLimitRequest {
	s.Id = &v
	return s
}

type DeleteDataLimitResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataLimitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataLimitResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataLimitResponseBody) SetRequestId(v string) *DeleteDataLimitResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDataLimitResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDataLimitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataLimitResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataLimitResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLimitResponse) SetHeaders(v map[string]*string) *DeleteDataLimitResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLimitResponse) SetBody(v *DeleteDataLimitResponseBody) *DeleteDataLimitResponse {
	s.Body = v
	return s
}

type DeleteRuleRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) SetLang(v string) *DeleteRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteRuleRequest) SetId(v int64) *DeleteRuleRequest {
	s.Id = &v
	return s
}

type DeleteRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponse) SetHeaders(v map[string]*string) *DeleteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRuleResponse) SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse {
	s.Body = v
	return s
}

type DescribeAccountDetailRequest struct {
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserId        *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	AccountTypeId *int64  `json:"AccountTypeId,omitempty" xml:"AccountTypeId,omitempty"`
}

func (s DescribeAccountDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountDetailRequest) SetLang(v string) *DescribeAccountDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAccountDetailRequest) SetUserId(v int64) *DescribeAccountDetailRequest {
	s.UserId = &v
	return s
}

func (s *DescribeAccountDetailRequest) SetAccountTypeId(v int64) *DescribeAccountDetailRequest {
	s.AccountTypeId = &v
	return s
}

type DescribeAccountDetailResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Account   *DescribeAccountDetailResponseBodyAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
}

func (s DescribeAccountDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountDetailResponseBody) SetRequestId(v string) *DescribeAccountDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountDetailResponseBody) SetAccount(v *DescribeAccountDetailResponseBodyAccount) *DescribeAccountDetailResponseBody {
	s.Account = v
	return s
}

type DescribeAccountDetailResponseBodyAccount struct {
	DataInstance         *string                                                  `json:"DataInstance,omitempty" xml:"DataInstance,omitempty"`
	LoginPolicyName      *string                                                  `json:"LoginPolicyName,omitempty" xml:"LoginPolicyName,omitempty"`
	LoginName            *string                                                  `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	FirstLevelDepartId   *int64                                                   `json:"FirstLevelDepartId,omitempty" xml:"FirstLevelDepartId,omitempty"`
	CreateTime           *int64                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FullName             *string                                                  `json:"FullName,omitempty" xml:"FullName,omitempty"`
	DeleteStatus         *string                                                  `json:"DeleteStatus,omitempty" xml:"DeleteStatus,omitempty"`
	UserId               *int64                                                   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	AliUid               *int64                                                   `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	RoleNames            *string                                                  `json:"RoleNames,omitempty" xml:"RoleNames,omitempty"`
	Email                *string                                                  `json:"Email,omitempty" xml:"Email,omitempty"`
	CellphoneNum         *string                                                  `json:"CellphoneNum,omitempty" xml:"CellphoneNum,omitempty"`
	FirstLevelDepartName *string                                                  `json:"FirstLevelDepartName,omitempty" xml:"FirstLevelDepartName,omitempty"`
	LoginDataTime        *int64                                                   `json:"LoginDataTime,omitempty" xml:"LoginDataTime,omitempty"`
	TelephoneNum         *string                                                  `json:"TelephoneNum,omitempty" xml:"TelephoneNum,omitempty"`
	ActiveStatus         *string                                                  `json:"ActiveStatus,omitempty" xml:"ActiveStatus,omitempty"`
	Id                   *int64                                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	PrivilegeList        []*DescribeAccountDetailResponseBodyAccountPrivilegeList `json:"PrivilegeList,omitempty" xml:"PrivilegeList,omitempty" type:"Repeated"`
	EventCount           *DescribeAccountDetailResponseBodyAccountEventCount      `json:"EventCount,omitempty" xml:"EventCount,omitempty" type:"Struct"`
}

func (s DescribeAccountDetailResponseBodyAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountDetailResponseBodyAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountDetailResponseBodyAccount) SetDataInstance(v string) *DescribeAccountDetailResponseBodyAccount {
	s.DataInstance = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetLoginPolicyName(v string) *DescribeAccountDetailResponseBodyAccount {
	s.LoginPolicyName = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetLoginName(v string) *DescribeAccountDetailResponseBodyAccount {
	s.LoginName = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetFirstLevelDepartId(v int64) *DescribeAccountDetailResponseBodyAccount {
	s.FirstLevelDepartId = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetCreateTime(v int64) *DescribeAccountDetailResponseBodyAccount {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetFullName(v string) *DescribeAccountDetailResponseBodyAccount {
	s.FullName = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetDeleteStatus(v string) *DescribeAccountDetailResponseBodyAccount {
	s.DeleteStatus = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetUserId(v int64) *DescribeAccountDetailResponseBodyAccount {
	s.UserId = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetAliUid(v int64) *DescribeAccountDetailResponseBodyAccount {
	s.AliUid = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetRoleNames(v string) *DescribeAccountDetailResponseBodyAccount {
	s.RoleNames = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetEmail(v string) *DescribeAccountDetailResponseBodyAccount {
	s.Email = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetCellphoneNum(v string) *DescribeAccountDetailResponseBodyAccount {
	s.CellphoneNum = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetFirstLevelDepartName(v string) *DescribeAccountDetailResponseBodyAccount {
	s.FirstLevelDepartName = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetLoginDataTime(v int64) *DescribeAccountDetailResponseBodyAccount {
	s.LoginDataTime = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetTelephoneNum(v string) *DescribeAccountDetailResponseBodyAccount {
	s.TelephoneNum = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetActiveStatus(v string) *DescribeAccountDetailResponseBodyAccount {
	s.ActiveStatus = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetId(v int64) *DescribeAccountDetailResponseBodyAccount {
	s.Id = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetPrivilegeList(v []*DescribeAccountDetailResponseBodyAccountPrivilegeList) *DescribeAccountDetailResponseBodyAccount {
	s.PrivilegeList = v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccount) SetEventCount(v *DescribeAccountDetailResponseBodyAccountEventCount) *DescribeAccountDetailResponseBodyAccount {
	s.EventCount = v
	return s
}

type DescribeAccountDetailResponseBodyAccountPrivilegeList struct {
	UseAccountId         *string `json:"UseAccountId,omitempty" xml:"UseAccountId,omitempty"`
	ProductName          *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	UseAccountTypeName   *string `json:"UseAccountTypeName,omitempty" xml:"UseAccountTypeName,omitempty"`
	UseAccountCreateTime *int64  `json:"UseAccountCreateTime,omitempty" xml:"UseAccountCreateTime,omitempty"`
	UseAccountName       *string `json:"UseAccountName,omitempty" xml:"UseAccountName,omitempty"`
	UseAccountType       *int32  `json:"UseAccountType,omitempty" xml:"UseAccountType,omitempty"`
}

func (s DescribeAccountDetailResponseBodyAccountPrivilegeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountDetailResponseBodyAccountPrivilegeList) GoString() string {
	return s.String()
}

func (s *DescribeAccountDetailResponseBodyAccountPrivilegeList) SetUseAccountId(v string) *DescribeAccountDetailResponseBodyAccountPrivilegeList {
	s.UseAccountId = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccountPrivilegeList) SetProductName(v string) *DescribeAccountDetailResponseBodyAccountPrivilegeList {
	s.ProductName = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccountPrivilegeList) SetUseAccountTypeName(v string) *DescribeAccountDetailResponseBodyAccountPrivilegeList {
	s.UseAccountTypeName = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccountPrivilegeList) SetUseAccountCreateTime(v int64) *DescribeAccountDetailResponseBodyAccountPrivilegeList {
	s.UseAccountCreateTime = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccountPrivilegeList) SetUseAccountName(v string) *DescribeAccountDetailResponseBodyAccountPrivilegeList {
	s.UseAccountName = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccountPrivilegeList) SetUseAccountType(v int32) *DescribeAccountDetailResponseBodyAccountPrivilegeList {
	s.UseAccountType = &v
	return s
}

type DescribeAccountDetailResponseBodyAccountEventCount struct {
	Total *DescribeAccountDetailResponseBodyAccountEventCountTotal `json:"Total,omitempty" xml:"Total,omitempty" type:"Struct"`
}

func (s DescribeAccountDetailResponseBodyAccountEventCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountDetailResponseBodyAccountEventCount) GoString() string {
	return s.String()
}

func (s *DescribeAccountDetailResponseBodyAccountEventCount) SetTotal(v *DescribeAccountDetailResponseBodyAccountEventCountTotal) *DescribeAccountDetailResponseBodyAccountEventCount {
	s.Total = v
	return s
}

type DescribeAccountDetailResponseBodyAccountEventCountTotal struct {
	ExcludeCount *int64 `json:"ExcludeCount,omitempty" xml:"ExcludeCount,omitempty"`
	ConfirmCount *int64 `json:"ConfirmCount,omitempty" xml:"ConfirmCount,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UndealCount  *int64 `json:"UndealCount,omitempty" xml:"UndealCount,omitempty"`
}

func (s DescribeAccountDetailResponseBodyAccountEventCountTotal) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountDetailResponseBodyAccountEventCountTotal) GoString() string {
	return s.String()
}

func (s *DescribeAccountDetailResponseBodyAccountEventCountTotal) SetExcludeCount(v int64) *DescribeAccountDetailResponseBodyAccountEventCountTotal {
	s.ExcludeCount = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccountEventCountTotal) SetConfirmCount(v int64) *DescribeAccountDetailResponseBodyAccountEventCountTotal {
	s.ConfirmCount = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccountEventCountTotal) SetTotalCount(v int64) *DescribeAccountDetailResponseBodyAccountEventCountTotal {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccountDetailResponseBodyAccountEventCountTotal) SetUndealCount(v int64) *DescribeAccountDetailResponseBodyAccountEventCountTotal {
	s.UndealCount = &v
	return s
}

type DescribeAccountDetailResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccountDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountDetailResponse) SetHeaders(v map[string]*string) *DescribeAccountDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountDetailResponse) SetBody(v *DescribeAccountDetailResponseBody) *DescribeAccountDetailResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	LoginName   *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	FeatureType *int32  `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	DepartId    *int64  `json:"DepartId,omitempty" xml:"DepartId,omitempty"`
	QueryType   *int32  `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	OperationId *int32  `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TableId     *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	PackageId   *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	ColumnId    *string `json:"ColumnId,omitempty" xml:"ColumnId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetLang(v string) *DescribeAccountsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAccountsRequest) SetLoginName(v string) *DescribeAccountsRequest {
	s.LoginName = &v
	return s
}

func (s *DescribeAccountsRequest) SetKey(v string) *DescribeAccountsRequest {
	s.Key = &v
	return s
}

func (s *DescribeAccountsRequest) SetFeatureType(v int32) *DescribeAccountsRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeAccountsRequest) SetDepartId(v int64) *DescribeAccountsRequest {
	s.DepartId = &v
	return s
}

func (s *DescribeAccountsRequest) SetQueryType(v int32) *DescribeAccountsRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeAccountsRequest) SetProductCode(v string) *DescribeAccountsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeAccountsRequest) SetOperationId(v int32) *DescribeAccountsRequest {
	s.OperationId = &v
	return s
}

func (s *DescribeAccountsRequest) SetInstanceId(v string) *DescribeAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAccountsRequest) SetTableId(v string) *DescribeAccountsRequest {
	s.TableId = &v
	return s
}

func (s *DescribeAccountsRequest) SetPackageId(v string) *DescribeAccountsRequest {
	s.PackageId = &v
	return s
}

func (s *DescribeAccountsRequest) SetColumnId(v string) *DescribeAccountsRequest {
	s.ColumnId = &v
	return s
}

func (s *DescribeAccountsRequest) SetPageSize(v int32) *DescribeAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountsRequest) SetCurrentPage(v int32) *DescribeAccountsRequest {
	s.CurrentPage = &v
	return s
}

type DescribeAccountsResponseBody struct {
	CurrentPage *int32                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeAccountsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetCurrentPage(v int32) *DescribeAccountsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetPageSize(v int32) *DescribeAccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetTotalCount(v int32) *DescribeAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetItems(v []*DescribeAccountsResponseBodyItems) *DescribeAccountsResponseBody {
	s.Items = v
	return s
}

type DescribeAccountsResponseBodyItems struct {
	DataInstance         *string                                      `json:"DataInstance,omitempty" xml:"DataInstance,omitempty"`
	LoginPolicyName      *string                                      `json:"LoginPolicyName,omitempty" xml:"LoginPolicyName,omitempty"`
	AccountTypeId        *int64                                       `json:"AccountTypeId,omitempty" xml:"AccountTypeId,omitempty"`
	LoginName            *string                                      `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	FirstLevelDepartId   *int64                                       `json:"FirstLevelDepartId,omitempty" xml:"FirstLevelDepartId,omitempty"`
	CreateTime           *int64                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FullName             *string                                      `json:"FullName,omitempty" xml:"FullName,omitempty"`
	DeleteStatus         *string                                      `json:"DeleteStatus,omitempty" xml:"DeleteStatus,omitempty"`
	UserId               *int64                                       `json:"UserId,omitempty" xml:"UserId,omitempty"`
	AliUid               *int64                                       `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	RoleNames            *string                                      `json:"RoleNames,omitempty" xml:"RoleNames,omitempty"`
	Email                *string                                      `json:"Email,omitempty" xml:"Email,omitempty"`
	CellphoneNum         *string                                      `json:"CellphoneNum,omitempty" xml:"CellphoneNum,omitempty"`
	InstanceName         *string                                      `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	FirstLevelDepartName *string                                      `json:"FirstLevelDepartName,omitempty" xml:"FirstLevelDepartName,omitempty"`
	LoginDataTime        *int64                                       `json:"LoginDataTime,omitempty" xml:"LoginDataTime,omitempty"`
	TelephoneNum         *string                                      `json:"TelephoneNum,omitempty" xml:"TelephoneNum,omitempty"`
	ActiveStatus         *string                                      `json:"ActiveStatus,omitempty" xml:"ActiveStatus,omitempty"`
	Id                   *int64                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	EventCount           *DescribeAccountsResponseBodyItemsEventCount `json:"EventCount,omitempty" xml:"EventCount,omitempty" type:"Struct"`
}

func (s DescribeAccountsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyItems) SetDataInstance(v string) *DescribeAccountsResponseBodyItems {
	s.DataInstance = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetLoginPolicyName(v string) *DescribeAccountsResponseBodyItems {
	s.LoginPolicyName = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetAccountTypeId(v int64) *DescribeAccountsResponseBodyItems {
	s.AccountTypeId = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetLoginName(v string) *DescribeAccountsResponseBodyItems {
	s.LoginName = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetFirstLevelDepartId(v int64) *DescribeAccountsResponseBodyItems {
	s.FirstLevelDepartId = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetCreateTime(v int64) *DescribeAccountsResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetFullName(v string) *DescribeAccountsResponseBodyItems {
	s.FullName = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetDeleteStatus(v string) *DescribeAccountsResponseBodyItems {
	s.DeleteStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetUserId(v int64) *DescribeAccountsResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetAliUid(v int64) *DescribeAccountsResponseBodyItems {
	s.AliUid = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetRoleNames(v string) *DescribeAccountsResponseBodyItems {
	s.RoleNames = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetEmail(v string) *DescribeAccountsResponseBodyItems {
	s.Email = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetCellphoneNum(v string) *DescribeAccountsResponseBodyItems {
	s.CellphoneNum = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetInstanceName(v string) *DescribeAccountsResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetFirstLevelDepartName(v string) *DescribeAccountsResponseBodyItems {
	s.FirstLevelDepartName = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetLoginDataTime(v int64) *DescribeAccountsResponseBodyItems {
	s.LoginDataTime = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetTelephoneNum(v string) *DescribeAccountsResponseBodyItems {
	s.TelephoneNum = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetActiveStatus(v string) *DescribeAccountsResponseBodyItems {
	s.ActiveStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetId(v int64) *DescribeAccountsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeAccountsResponseBodyItems) SetEventCount(v *DescribeAccountsResponseBodyItemsEventCount) *DescribeAccountsResponseBodyItems {
	s.EventCount = v
	return s
}

type DescribeAccountsResponseBodyItemsEventCount struct {
	Total *DescribeAccountsResponseBodyItemsEventCountTotal `json:"Total,omitempty" xml:"Total,omitempty" type:"Struct"`
}

func (s DescribeAccountsResponseBodyItemsEventCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyItemsEventCount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyItemsEventCount) SetTotal(v *DescribeAccountsResponseBodyItemsEventCountTotal) *DescribeAccountsResponseBodyItemsEventCount {
	s.Total = v
	return s
}

type DescribeAccountsResponseBodyItemsEventCountTotal struct {
	ExcludeCount *int64 `json:"ExcludeCount,omitempty" xml:"ExcludeCount,omitempty"`
	ConfirmCount *int64 `json:"ConfirmCount,omitempty" xml:"ConfirmCount,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UndealCount  *int64 `json:"UndealCount,omitempty" xml:"UndealCount,omitempty"`
}

func (s DescribeAccountsResponseBodyItemsEventCountTotal) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyItemsEventCountTotal) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyItemsEventCountTotal) SetExcludeCount(v int64) *DescribeAccountsResponseBodyItemsEventCountTotal {
	s.ExcludeCount = &v
	return s
}

func (s *DescribeAccountsResponseBodyItemsEventCountTotal) SetConfirmCount(v int64) *DescribeAccountsResponseBodyItemsEventCountTotal {
	s.ConfirmCount = &v
	return s
}

func (s *DescribeAccountsResponseBodyItemsEventCountTotal) SetTotalCount(v int64) *DescribeAccountsResponseBodyItemsEventCountTotal {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccountsResponseBodyItemsEventCountTotal) SetUndealCount(v int64) *DescribeAccountsResponseBodyItemsEventCountTotal {
	s.UndealCount = &v
	return s
}

type DescribeAccountsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponse) SetHeaders(v map[string]*string) *DescribeAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountsResponse) SetBody(v *DescribeAccountsResponseBody) *DescribeAccountsResponse {
	s.Body = v
	return s
}

type DescribeColumnsRequest struct {
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TableId       *int64  `json:"TableId,omitempty" xml:"TableId,omitempty"`
	RuleId        *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	ProductCode   *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	InstanceId    *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	TableName     *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	RuleName      *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SensLevelName *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RiskLevelId   *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
}

func (s DescribeColumnsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeColumnsRequest) SetLang(v string) *DescribeColumnsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeColumnsRequest) SetName(v string) *DescribeColumnsRequest {
	s.Name = &v
	return s
}

func (s *DescribeColumnsRequest) SetTableId(v int64) *DescribeColumnsRequest {
	s.TableId = &v
	return s
}

func (s *DescribeColumnsRequest) SetRuleId(v int64) *DescribeColumnsRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeColumnsRequest) SetProductCode(v string) *DescribeColumnsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeColumnsRequest) SetInstanceId(v int64) *DescribeColumnsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeColumnsRequest) SetInstanceName(v string) *DescribeColumnsRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeColumnsRequest) SetTableName(v string) *DescribeColumnsRequest {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsRequest) SetRuleName(v string) *DescribeColumnsRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeColumnsRequest) SetSensLevelName(v string) *DescribeColumnsRequest {
	s.SensLevelName = &v
	return s
}

func (s *DescribeColumnsRequest) SetPageSize(v int32) *DescribeColumnsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeColumnsRequest) SetCurrentPage(v int32) *DescribeColumnsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeColumnsRequest) SetRiskLevelId(v int64) *DescribeColumnsRequest {
	s.RiskLevelId = &v
	return s
}

type DescribeColumnsResponseBody struct {
	CurrentPage *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeColumnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBody) SetCurrentPage(v int32) *DescribeColumnsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeColumnsResponseBody) SetRequestId(v string) *DescribeColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColumnsResponseBody) SetPageSize(v int32) *DescribeColumnsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeColumnsResponseBody) SetTotalCount(v int32) *DescribeColumnsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeColumnsResponseBody) SetItems(v []*DescribeColumnsResponseBodyItems) *DescribeColumnsResponseBody {
	s.Items = v
	return s
}

type DescribeColumnsResponseBodyItems struct {
	CreationTime       *int64  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	TableName          *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	DataType           *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	OdpsRiskLevelValue *int32  `json:"OdpsRiskLevelValue,omitempty" xml:"OdpsRiskLevelValue,omitempty"`
	InstanceId         *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RiskLevelId        *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleId             *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Sensitive          *bool   `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	SensLevelName      *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	InstanceName       *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RiskLevelName      *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	OdpsRiskLevelName  *string `json:"OdpsRiskLevelName,omitempty" xml:"OdpsRiskLevelName,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TableId            *int64  `json:"TableId,omitempty" xml:"TableId,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProductCode        *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RevisionStatus     *int64  `json:"RevisionStatus,omitempty" xml:"RevisionStatus,omitempty"`
	RevisionId         *int64  `json:"RevisionId,omitempty" xml:"RevisionId,omitempty"`
}

func (s DescribeColumnsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItems) SetCreationTime(v int64) *DescribeColumnsResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetTableName(v string) *DescribeColumnsResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetDataType(v string) *DescribeColumnsResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetOdpsRiskLevelValue(v int32) *DescribeColumnsResponseBodyItems {
	s.OdpsRiskLevelValue = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetInstanceId(v int64) *DescribeColumnsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRiskLevelId(v int64) *DescribeColumnsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRuleName(v string) *DescribeColumnsResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRuleId(v int64) *DescribeColumnsResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetSensitive(v bool) *DescribeColumnsResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetSensLevelName(v string) *DescribeColumnsResponseBodyItems {
	s.SensLevelName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetInstanceName(v string) *DescribeColumnsResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRiskLevelName(v string) *DescribeColumnsResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetOdpsRiskLevelName(v string) *DescribeColumnsResponseBodyItems {
	s.OdpsRiskLevelName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetName(v string) *DescribeColumnsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetTableId(v int64) *DescribeColumnsResponseBodyItems {
	s.TableId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetId(v string) *DescribeColumnsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetProductCode(v string) *DescribeColumnsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRevisionStatus(v int64) *DescribeColumnsResponseBodyItems {
	s.RevisionStatus = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRevisionId(v int64) *DescribeColumnsResponseBodyItems {
	s.RevisionId = &v
	return s
}

type DescribeColumnsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeColumnsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponse) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponse) SetHeaders(v map[string]*string) *DescribeColumnsResponse {
	s.Headers = v
	return s
}

func (s *DescribeColumnsResponse) SetBody(v *DescribeColumnsResponseBody) *DescribeColumnsResponse {
	s.Body = v
	return s
}

type DescribeConfigsRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigsRequest) SetLang(v string) *DescribeConfigsRequest {
	s.Lang = &v
	return s
}

type DescribeConfigsResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConfigList []*DescribeConfigsResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
}

func (s DescribeConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigsResponseBody) SetRequestId(v string) *DescribeConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigsResponseBody) SetConfigList(v []*DescribeConfigsResponseBodyConfigList) *DescribeConfigsResponseBody {
	s.ConfigList = v
	return s
}

type DescribeConfigsResponseBodyConfigList struct {
	Code         *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Value        *int64  `json:"Value,omitempty" xml:"Value,omitempty"`
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeConfigsResponseBodyConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigsResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeConfigsResponseBodyConfigList) SetCode(v int32) *DescribeConfigsResponseBodyConfigList {
	s.Code = &v
	return s
}

func (s *DescribeConfigsResponseBodyConfigList) SetValue(v int64) *DescribeConfigsResponseBodyConfigList {
	s.Value = &v
	return s
}

func (s *DescribeConfigsResponseBodyConfigList) SetDefaultValue(v string) *DescribeConfigsResponseBodyConfigList {
	s.DefaultValue = &v
	return s
}

func (s *DescribeConfigsResponseBodyConfigList) SetDescription(v string) *DescribeConfigsResponseBodyConfigList {
	s.Description = &v
	return s
}

func (s *DescribeConfigsResponseBodyConfigList) SetId(v int64) *DescribeConfigsResponseBodyConfigList {
	s.Id = &v
	return s
}

type DescribeConfigsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigsResponse) SetHeaders(v map[string]*string) *DescribeConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigsResponse) SetBody(v *DescribeConfigsResponseBody) *DescribeConfigsResponse {
	s.Body = v
	return s
}

type DescribeDataAssetsRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RiskLevels  *string `json:"RiskLevels,omitempty" xml:"RiskLevels,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RangeId     *int32  `json:"RangeId,omitempty" xml:"RangeId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeDataAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataAssetsRequest) SetLang(v string) *DescribeDataAssetsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetName(v string) *DescribeDataAssetsRequest {
	s.Name = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetRiskLevels(v string) *DescribeDataAssetsRequest {
	s.RiskLevels = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetRuleId(v int64) *DescribeDataAssetsRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetRangeId(v int32) *DescribeDataAssetsRequest {
	s.RangeId = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetPageSize(v int32) *DescribeDataAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetCurrentPage(v int32) *DescribeDataAssetsRequest {
	s.CurrentPage = &v
	return s
}

type DescribeDataAssetsResponseBody struct {
	CurrentPage *int32                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeDataAssetsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeDataAssetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataAssetsResponseBody) SetCurrentPage(v int32) *DescribeDataAssetsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataAssetsResponseBody) SetRequestId(v string) *DescribeDataAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataAssetsResponseBody) SetPageSize(v int32) *DescribeDataAssetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataAssetsResponseBody) SetTotalCount(v int32) *DescribeDataAssetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataAssetsResponseBody) SetItems(v []*DescribeDataAssetsResponseBodyItems) *DescribeDataAssetsResponseBody {
	s.Items = v
	return s
}

type DescribeDataAssetsResponseBodyItems struct {
	Acl               *string `json:"Acl,omitempty" xml:"Acl,omitempty"`
	CreationTime      *int64  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DataType          *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Owner             *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SensitiveRatio    *string `json:"SensitiveRatio,omitempty" xml:"SensitiveRatio,omitempty"`
	Protection        *bool   `json:"Protection,omitempty" xml:"Protection,omitempty"`
	Labelsec          *bool   `json:"Labelsec,omitempty" xml:"Labelsec,omitempty"`
	TotalCount        *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RiskLevelId       *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sensitive         *bool   `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	ObjectKey         *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	RiskLevelName     *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	OdpsRiskLevelName *string `json:"OdpsRiskLevelName,omitempty" xml:"OdpsRiskLevelName,omitempty"`
	ProductId         *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SensitiveCount    *int32  `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	Id                *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeDataAssetsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataAssetsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataAssetsResponseBodyItems) SetAcl(v string) *DescribeDataAssetsResponseBodyItems {
	s.Acl = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetCreationTime(v int64) *DescribeDataAssetsResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetDataType(v string) *DescribeDataAssetsResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetOwner(v string) *DescribeDataAssetsResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetSensitiveRatio(v string) *DescribeDataAssetsResponseBodyItems {
	s.SensitiveRatio = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetProtection(v bool) *DescribeDataAssetsResponseBodyItems {
	s.Protection = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetLabelsec(v bool) *DescribeDataAssetsResponseBodyItems {
	s.Labelsec = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetTotalCount(v int32) *DescribeDataAssetsResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetRiskLevelId(v int64) *DescribeDataAssetsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetRuleName(v string) *DescribeDataAssetsResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetSensitive(v bool) *DescribeDataAssetsResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetObjectKey(v string) *DescribeDataAssetsResponseBodyItems {
	s.ObjectKey = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetRiskLevelName(v string) *DescribeDataAssetsResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetOdpsRiskLevelName(v string) *DescribeDataAssetsResponseBodyItems {
	s.OdpsRiskLevelName = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetProductId(v string) *DescribeDataAssetsResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetName(v string) *DescribeDataAssetsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetSensitiveCount(v int32) *DescribeDataAssetsResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetId(v string) *DescribeDataAssetsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetProductCode(v string) *DescribeDataAssetsResponseBodyItems {
	s.ProductCode = &v
	return s
}

type DescribeDataAssetsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataAssetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataAssetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataAssetsResponse) SetHeaders(v map[string]*string) *DescribeDataAssetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataAssetsResponse) SetBody(v *DescribeDataAssetsResponseBody) *DescribeDataAssetsResponse {
	s.Body = v
	return s
}

type DescribeDataLimitDetailRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	NetworkType *int32  `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s DescribeDataLimitDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitDetailRequest) SetLang(v string) *DescribeDataLimitDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataLimitDetailRequest) SetId(v int64) *DescribeDataLimitDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitDetailRequest) SetNetworkType(v int32) *DescribeDataLimitDetailRequest {
	s.NetworkType = &v
	return s
}

type DescribeDataLimitDetailResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DataLimit *DescribeDataLimitDetailResponseBodyDataLimit `json:"DataLimit,omitempty" xml:"DataLimit,omitempty" type:"Struct"`
}

func (s DescribeDataLimitDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitDetailResponseBody) SetRequestId(v string) *DescribeDataLimitDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBody) SetDataLimit(v *DescribeDataLimitDetailResponseBodyDataLimit) *DescribeDataLimitDetailResponseBody {
	s.DataLimit = v
	return s
}

type DescribeDataLimitDetailResponseBodyDataLimit struct {
	LocalName        *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	CheckStatus      *int32  `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	CheckStatusName  *string `json:"CheckStatusName,omitempty" xml:"CheckStatusName,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ParentId         *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	ResourceType     *int64  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	GmtCreate        *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDataLimitDetailResponseBodyDataLimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitDetailResponseBodyDataLimit) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetLocalName(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetCheckStatus(v int32) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetPort(v int32) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.Port = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetCheckStatusName(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.CheckStatusName = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetRegionId(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetParentId(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetResourceType(v int64) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetGmtCreate(v int64) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetResourceTypeCode(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetUserName(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.UserName = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetId(v int64) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.Id = &v
	return s
}

type DescribeDataLimitDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataLimitDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataLimitDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitDetailResponse) SetHeaders(v map[string]*string) *DescribeDataLimitDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataLimitDetailResponse) SetBody(v *DescribeDataLimitDetailResponseBody) *DescribeDataLimitDetailResponse {
	s.Body = v
	return s
}

type DescribeDataLimitsRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ResourceType    *int32  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ParentId        *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	CheckStatus     *int32  `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	AuditStatus     *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	EngineType      *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	DatamaskStatus  *int32  `json:"DatamaskStatus,omitempty" xml:"DatamaskStatus,omitempty"`
	Enable          *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDataLimitsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsRequest) SetLang(v string) *DescribeDataLimitsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetPageSize(v int32) *DescribeDataLimitsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetCurrentPage(v int32) *DescribeDataLimitsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetResourceType(v int32) *DescribeDataLimitsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetParentId(v string) *DescribeDataLimitsRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetServiceRegionId(v string) *DescribeDataLimitsRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetCheckStatus(v int32) *DescribeDataLimitsRequest {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetAuditStatus(v int32) *DescribeDataLimitsRequest {
	s.AuditStatus = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetEngineType(v string) *DescribeDataLimitsRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetDatamaskStatus(v int32) *DescribeDataLimitsRequest {
	s.DatamaskStatus = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetEnable(v int32) *DescribeDataLimitsRequest {
	s.Enable = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetStartTime(v int64) *DescribeDataLimitsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetEndTime(v int64) *DescribeDataLimitsRequest {
	s.EndTime = &v
	return s
}

type DescribeDataLimitsResponseBody struct {
	CurrentPage *int32                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeDataLimitsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeDataLimitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsResponseBody) SetCurrentPage(v int32) *DescribeDataLimitsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataLimitsResponseBody) SetRequestId(v string) *DescribeDataLimitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataLimitsResponseBody) SetPageSize(v int32) *DescribeDataLimitsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataLimitsResponseBody) SetTotalCount(v int32) *DescribeDataLimitsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataLimitsResponseBody) SetItems(v []*DescribeDataLimitsResponseBodyItems) *DescribeDataLimitsResponseBody {
	s.Items = v
	return s
}

type DescribeDataLimitsResponseBodyItems struct {
	SupportEvent        *bool   `json:"SupportEvent,omitempty" xml:"SupportEvent,omitempty"`
	ErrorMessage        *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	CheckStatus         *int32  `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	LocalName           *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	TenantName          *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	NextStartTime       *int64  `json:"NextStartTime,omitempty" xml:"NextStartTime,omitempty"`
	Port                *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	CheckStatusName     *string `json:"CheckStatusName,omitempty" xml:"CheckStatusName,omitempty"`
	SamplingSize        *int32  `json:"SamplingSize,omitempty" xml:"SamplingSize,omitempty"`
	ParentId            *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	DatamaskStatus      *int32  `json:"DatamaskStatus,omitempty" xml:"DatamaskStatus,omitempty"`
	ProcessTotalCount   *int32  `json:"ProcessTotalCount,omitempty" xml:"ProcessTotalCount,omitempty"`
	ResourceType        *int64  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ErrorCode           *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	OcrStatus           *int32  `json:"OcrStatus,omitempty" xml:"OcrStatus,omitempty"`
	LogStoreDay         *int32  `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	EventStatus         *int32  `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	SupportScan         *bool   `json:"SupportScan,omitempty" xml:"SupportScan,omitempty"`
	LastFinishedTime    *int64  `json:"LastFinishedTime,omitempty" xml:"LastFinishedTime,omitempty"`
	UserName            *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	AuditStatus         *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	SupportOcr          *bool   `json:"SupportOcr,omitempty" xml:"SupportOcr,omitempty"`
	EngineType          *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TotalCount          *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	DbVersion           *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	GmtCreate           *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SupportAudit        *bool   `json:"SupportAudit,omitempty" xml:"SupportAudit,omitempty"`
	AutoScan            *int32  `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	ResourceTypeCode    *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	SupportDatamask     *bool   `json:"SupportDatamask,omitempty" xml:"SupportDatamask,omitempty"`
	ProcessStatus       *int32  `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Enable              *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeDataLimitsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportEvent(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportEvent = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetErrorMessage(v string) *DescribeDataLimitsResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetCheckStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetLocalName(v string) *DescribeDataLimitsResponseBodyItems {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetTenantName(v string) *DescribeDataLimitsResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetNextStartTime(v int64) *DescribeDataLimitsResponseBodyItems {
	s.NextStartTime = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetPort(v int32) *DescribeDataLimitsResponseBodyItems {
	s.Port = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetCheckStatusName(v string) *DescribeDataLimitsResponseBodyItems {
	s.CheckStatusName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSamplingSize(v int32) *DescribeDataLimitsResponseBodyItems {
	s.SamplingSize = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetParentId(v string) *DescribeDataLimitsResponseBodyItems {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetDatamaskStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.DatamaskStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetProcessTotalCount(v int32) *DescribeDataLimitsResponseBodyItems {
	s.ProcessTotalCount = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetResourceType(v int64) *DescribeDataLimitsResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetErrorCode(v string) *DescribeDataLimitsResponseBodyItems {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetOcrStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.OcrStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetLogStoreDay(v int32) *DescribeDataLimitsResponseBodyItems {
	s.LogStoreDay = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetEventStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.EventStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportScan(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportScan = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetLastFinishedTime(v int64) *DescribeDataLimitsResponseBodyItems {
	s.LastFinishedTime = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetUserName(v string) *DescribeDataLimitsResponseBodyItems {
	s.UserName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetAuditStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.AuditStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportOcr(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportOcr = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetEngineType(v string) *DescribeDataLimitsResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetInstanceId(v string) *DescribeDataLimitsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetTotalCount(v int32) *DescribeDataLimitsResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetInstanceDescription(v string) *DescribeDataLimitsResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetDbVersion(v string) *DescribeDataLimitsResponseBodyItems {
	s.DbVersion = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetRegionId(v string) *DescribeDataLimitsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetGmtCreate(v int64) *DescribeDataLimitsResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportAudit(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportAudit = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetAutoScan(v int32) *DescribeDataLimitsResponseBodyItems {
	s.AutoScan = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetResourceTypeCode(v string) *DescribeDataLimitsResponseBodyItems {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportDatamask(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportDatamask = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetProcessStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetId(v int64) *DescribeDataLimitsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetEnable(v int32) *DescribeDataLimitsResponseBodyItems {
	s.Enable = &v
	return s
}

type DescribeDataLimitsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataLimitsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataLimitsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsResponse) SetHeaders(v map[string]*string) *DescribeDataLimitsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataLimitsResponse) SetBody(v *DescribeDataLimitsResponseBody) *DescribeDataLimitsResponse {
	s.Body = v
	return s
}

type DescribeDataLimitSetRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceType *int32  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ParentId     *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s DescribeDataLimitSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetRequest) SetLang(v string) *DescribeDataLimitSetRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataLimitSetRequest) SetResourceType(v int32) *DescribeDataLimitSetRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitSetRequest) SetParentId(v string) *DescribeDataLimitSetRequest {
	s.ParentId = &v
	return s
}

type DescribeDataLimitSetResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DataLimitSet *DescribeDataLimitSetResponseBodyDataLimitSet `json:"DataLimitSet,omitempty" xml:"DataLimitSet,omitempty" type:"Struct"`
}

func (s DescribeDataLimitSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBody) SetRequestId(v string) *DescribeDataLimitSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBody) SetDataLimitSet(v *DescribeDataLimitSetResponseBodyDataLimitSet) *DescribeDataLimitSetResponseBody {
	s.DataLimitSet = v
	return s
}

type DescribeDataLimitSetResponseBodyDataLimitSet struct {
	ResourceTypeCode *string                                                      `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	ResourceType     *int64                                                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TotalCount       *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DataLimitList    []*DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList `json:"DataLimitList,omitempty" xml:"DataLimitList,omitempty" type:"Repeated"`
	OssBucketList    []*DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList `json:"OssBucketList,omitempty" xml:"OssBucketList,omitempty" type:"Repeated"`
	RegionList       []*DescribeDataLimitSetResponseBodyDataLimitSetRegionList    `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSet) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetResourceTypeCode(v string) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetResourceType(v int64) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetTotalCount(v int32) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetDataLimitList(v []*DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.DataLimitList = v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetOssBucketList(v []*DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.OssBucketList = v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetRegionList(v []*DescribeDataLimitSetResponseBodyDataLimitSetRegionList) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.RegionList = v
	return s
}

type DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList struct {
	SupportEvent     *bool   `json:"SupportEvent,omitempty" xml:"SupportEvent,omitempty"`
	LocalName        *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	CheckStatus      *int32  `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	Connector        *string `json:"Connector,omitempty" xml:"Connector,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	CheckStatusName  *string `json:"CheckStatusName,omitempty" xml:"CheckStatusName,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ParentId         *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	ResourceType     *int64  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	LogStoreDay      *int32  `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	EventStatus      *int32  `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	GmtCreate        *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	AuditStatus      *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetSupportEvent(v bool) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.SupportEvent = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetLocalName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetCheckStatus(v int32) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetConnector(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.Connector = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetPort(v int32) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.Port = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetCheckStatusName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.CheckStatusName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetRegionId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetParentId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetResourceType(v int64) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetLogStoreDay(v int32) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.LogStoreDay = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetEventStatus(v int32) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.EventStatus = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetGmtCreate(v int64) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetResourceTypeCode(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetUserName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.UserName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetId(v int64) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetAuditStatus(v int32) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.AuditStatus = &v
	return s
}

type DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) SetRegionId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) SetBucketName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList {
	s.BucketName = &v
	return s
}

type DescribeDataLimitSetResponseBodyDataLimitSetRegionList struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetRegionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetRegionList) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetRegionList) SetRegionId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetRegionList {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetRegionList) SetLocalName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetRegionList {
	s.LocalName = &v
	return s
}

type DescribeDataLimitSetResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataLimitSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataLimitSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponse) SetHeaders(v map[string]*string) *DescribeDataLimitSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataLimitSetResponse) SetBody(v *DescribeDataLimitSetResponseBody) *DescribeDataLimitSetResponse {
	s.Body = v
	return s
}

type DescribeDataLimitsInstanceDimRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ResourceType    *int32  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ParentId        *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	CheckStatus     *int32  `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	AuditStatus     *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	EngineType      *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	DatamaskStatus  *int32  `json:"DatamaskStatus,omitempty" xml:"DatamaskStatus,omitempty"`
	Enable          *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDataLimitsInstanceDimRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitsInstanceDimRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsInstanceDimRequest) SetLang(v string) *DescribeDataLimitsInstanceDimRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetPageSize(v int32) *DescribeDataLimitsInstanceDimRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetCurrentPage(v int32) *DescribeDataLimitsInstanceDimRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetResourceType(v int32) *DescribeDataLimitsInstanceDimRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetParentId(v string) *DescribeDataLimitsInstanceDimRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetServiceRegionId(v string) *DescribeDataLimitsInstanceDimRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetCheckStatus(v int32) *DescribeDataLimitsInstanceDimRequest {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetAuditStatus(v int32) *DescribeDataLimitsInstanceDimRequest {
	s.AuditStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetEngineType(v string) *DescribeDataLimitsInstanceDimRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetDatamaskStatus(v int32) *DescribeDataLimitsInstanceDimRequest {
	s.DatamaskStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetEnable(v int32) *DescribeDataLimitsInstanceDimRequest {
	s.Enable = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetStartTime(v int64) *DescribeDataLimitsInstanceDimRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimRequest) SetEndTime(v int64) *DescribeDataLimitsInstanceDimRequest {
	s.EndTime = &v
	return s
}

type DescribeDataLimitsInstanceDimResponseBody struct {
	CurrentPage *int32                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeDataLimitsInstanceDimResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeDataLimitsInstanceDimResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitsInstanceDimResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsInstanceDimResponseBody) SetCurrentPage(v int32) *DescribeDataLimitsInstanceDimResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBody) SetRequestId(v string) *DescribeDataLimitsInstanceDimResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBody) SetPageSize(v int32) *DescribeDataLimitsInstanceDimResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBody) SetTotalCount(v int32) *DescribeDataLimitsInstanceDimResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBody) SetItems(v []*DescribeDataLimitsInstanceDimResponseBodyItems) *DescribeDataLimitsInstanceDimResponseBody {
	s.Items = v
	return s
}

type DescribeDataLimitsInstanceDimResponseBodyItems struct {
	ErrorMessage        *string                                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	CheckStatus         *int32                                                         `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	LocalName           *string                                                        `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	NextStartTime       *int64                                                         `json:"NextStartTime,omitempty" xml:"NextStartTime,omitempty"`
	Port                *int32                                                         `json:"Port,omitempty" xml:"Port,omitempty"`
	CheckStatusName     *string                                                        `json:"CheckStatusName,omitempty" xml:"CheckStatusName,omitempty"`
	SamplingSize        *int32                                                         `json:"SamplingSize,omitempty" xml:"SamplingSize,omitempty"`
	ParentId            *string                                                        `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	DatamaskStatus      *int32                                                         `json:"DatamaskStatus,omitempty" xml:"DatamaskStatus,omitempty"`
	ProcessTotalCount   *int32                                                         `json:"ProcessTotalCount,omitempty" xml:"ProcessTotalCount,omitempty"`
	ResourceType        *int64                                                         `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ErrorCode           *string                                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	OcrStatus           *int32                                                         `json:"OcrStatus,omitempty" xml:"OcrStatus,omitempty"`
	LogStoreDay         *int32                                                         `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	SupportScan         *bool                                                          `json:"SupportScan,omitempty" xml:"SupportScan,omitempty"`
	LastFinishedTime    *int64                                                         `json:"LastFinishedTime,omitempty" xml:"LastFinishedTime,omitempty"`
	UserName            *string                                                        `json:"UserName,omitempty" xml:"UserName,omitempty"`
	AuditStatus         *int32                                                         `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	SupportOcr          *bool                                                          `json:"SupportOcr,omitempty" xml:"SupportOcr,omitempty"`
	EngineType          *string                                                        `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	TotalCount          *int32                                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	InstanceDescription *string                                                        `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	DbVersion           *string                                                        `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	RegionId            *string                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	GmtCreate           *int64                                                         `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SupportAudit        *bool                                                          `json:"SupportAudit,omitempty" xml:"SupportAudit,omitempty"`
	AutoScan            *int32                                                         `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	ResourceTypeCode    *string                                                        `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	SupportDatamask     *bool                                                          `json:"SupportDatamask,omitempty" xml:"SupportDatamask,omitempty"`
	ProcessStatus       *int32                                                         `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	Id                  *int64                                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	Enable              *int32                                                         `json:"Enable,omitempty" xml:"Enable,omitempty"`
	DataLimitList       []*DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList `json:"DataLimitList,omitempty" xml:"DataLimitList,omitempty" type:"Repeated"`
}

func (s DescribeDataLimitsInstanceDimResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitsInstanceDimResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetErrorMessage(v string) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetCheckStatus(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetLocalName(v string) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetNextStartTime(v int64) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.NextStartTime = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetPort(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.Port = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetCheckStatusName(v string) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.CheckStatusName = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetSamplingSize(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.SamplingSize = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetParentId(v string) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetDatamaskStatus(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.DatamaskStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetProcessTotalCount(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.ProcessTotalCount = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetResourceType(v int64) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetErrorCode(v string) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetOcrStatus(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.OcrStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetLogStoreDay(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.LogStoreDay = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetSupportScan(v bool) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.SupportScan = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetLastFinishedTime(v int64) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.LastFinishedTime = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetUserName(v string) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.UserName = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetAuditStatus(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.AuditStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetSupportOcr(v bool) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.SupportOcr = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetEngineType(v string) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetTotalCount(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetInstanceDescription(v string) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetDbVersion(v string) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.DbVersion = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetRegionId(v string) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetGmtCreate(v int64) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetSupportAudit(v bool) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.SupportAudit = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetAutoScan(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.AutoScan = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetResourceTypeCode(v string) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetSupportDatamask(v bool) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.SupportDatamask = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetProcessStatus(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetId(v int64) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetEnable(v int32) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.Enable = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItems) SetDataLimitList(v []*DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) *DescribeDataLimitsInstanceDimResponseBodyItems {
	s.DataLimitList = v
	return s
}

type DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList struct {
	SupportEvent      *bool   `json:"SupportEvent,omitempty" xml:"SupportEvent,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	CheckStatus       *int32  `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	LocalName         *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	NextStartTime     *int64  `json:"NextStartTime,omitempty" xml:"NextStartTime,omitempty"`
	Port              *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	CheckStatusName   *string `json:"CheckStatusName,omitempty" xml:"CheckStatusName,omitempty"`
	SamplingSize      *int32  `json:"SamplingSize,omitempty" xml:"SamplingSize,omitempty"`
	ParentId          *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	DatamaskStatus    *int32  `json:"DatamaskStatus,omitempty" xml:"DatamaskStatus,omitempty"`
	ProcessTotalCount *int32  `json:"ProcessTotalCount,omitempty" xml:"ProcessTotalCount,omitempty"`
	ResourceType      *int64  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ErrorCode         *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	OcrStatus         *int32  `json:"OcrStatus,omitempty" xml:"OcrStatus,omitempty"`
	LogStoreDay       *int32  `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	EventStatus       *int32  `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	SupportScan       *bool   `json:"SupportScan,omitempty" xml:"SupportScan,omitempty"`
	LastFinishedTime  *int64  `json:"LastFinishedTime,omitempty" xml:"LastFinishedTime,omitempty"`
	UserName          *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	AuditStatus       *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AutoCreateAccount *bool   `json:"AutoCreateAccount,omitempty" xml:"AutoCreateAccount,omitempty"`
	SupportOcr        *bool   `json:"SupportOcr,omitempty" xml:"SupportOcr,omitempty"`
	EngineType        *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	TotalCount        *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DbVersion         *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	GmtCreate         *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SupportAudit      *bool   `json:"SupportAudit,omitempty" xml:"SupportAudit,omitempty"`
	AutoScan          *int32  `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	ResourceTypeCode  *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	SupportDatamask   *bool   `json:"SupportDatamask,omitempty" xml:"SupportDatamask,omitempty"`
	ProcessStatus     *int32  `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Enable            *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetSupportEvent(v bool) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.SupportEvent = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetErrorMessage(v string) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetCheckStatus(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetLocalName(v string) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetNextStartTime(v int64) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.NextStartTime = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetPort(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.Port = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetCheckStatusName(v string) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.CheckStatusName = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetSamplingSize(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.SamplingSize = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetParentId(v string) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetDatamaskStatus(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.DatamaskStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetProcessTotalCount(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.ProcessTotalCount = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetResourceType(v int64) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetErrorCode(v string) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetOcrStatus(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.OcrStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetLogStoreDay(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.LogStoreDay = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetEventStatus(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.EventStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetSupportScan(v bool) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.SupportScan = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetLastFinishedTime(v int64) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.LastFinishedTime = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetUserName(v string) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.UserName = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetAuditStatus(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.AuditStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetAutoCreateAccount(v bool) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.AutoCreateAccount = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetSupportOcr(v bool) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.SupportOcr = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetEngineType(v string) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.EngineType = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetTotalCount(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetDbVersion(v string) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.DbVersion = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetRegionId(v string) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetGmtCreate(v int64) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetSupportAudit(v bool) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.SupportAudit = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetAutoScan(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.AutoScan = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetResourceTypeCode(v string) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetSupportDatamask(v bool) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.SupportDatamask = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetProcessStatus(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetId(v int64) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList) SetEnable(v int32) *DescribeDataLimitsInstanceDimResponseBodyItemsDataLimitList {
	s.Enable = &v
	return s
}

type DescribeDataLimitsInstanceDimResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataLimitsInstanceDimResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataLimitsInstanceDimResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitsInstanceDimResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsInstanceDimResponse) SetHeaders(v map[string]*string) *DescribeDataLimitsInstanceDimResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataLimitsInstanceDimResponse) SetBody(v *DescribeDataLimitsInstanceDimResponseBody) *DescribeDataLimitsInstanceDimResponse {
	s.Body = v
	return s
}

type DescribeDataMaskingRunHistoryRequest struct {
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SrcType       *int32  `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	DstType       *int32  `json:"DstType,omitempty" xml:"DstType,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	MainProcessId *int64  `json:"MainProcessId,omitempty" xml:"MainProcessId,omitempty"`
	SrcTableName  *string `json:"SrcTableName,omitempty" xml:"SrcTableName,omitempty"`
}

func (s DescribeDataMaskingRunHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingRunHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingRunHistoryRequest) SetLang(v string) *DescribeDataMaskingRunHistoryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetTaskId(v string) *DescribeDataMaskingRunHistoryRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetStartTime(v int64) *DescribeDataMaskingRunHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetEndTime(v int64) *DescribeDataMaskingRunHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetStatus(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.Status = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetSrcType(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.SrcType = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetDstType(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.DstType = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetPageSize(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetCurrentPage(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetMainProcessId(v int64) *DescribeDataMaskingRunHistoryRequest {
	s.MainProcessId = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetSrcTableName(v string) *DescribeDataMaskingRunHistoryRequest {
	s.SrcTableName = &v
	return s
}

type DescribeDataMaskingRunHistoryResponseBody struct {
	CurrentPage *int32                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeDataMaskingRunHistoryResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeDataMaskingRunHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingRunHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetCurrentPage(v int32) *DescribeDataMaskingRunHistoryResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetRequestId(v string) *DescribeDataMaskingRunHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetPageSize(v int32) *DescribeDataMaskingRunHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetTotalCount(v int32) *DescribeDataMaskingRunHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetItems(v []*DescribeDataMaskingRunHistoryResponseBodyItems) *DescribeDataMaskingRunHistoryResponseBody {
	s.Items = v
	return s
}

type DescribeDataMaskingRunHistoryResponseBodyItems struct {
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	SrcType         *int32  `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	SrcTableName    *string `json:"SrcTableName,omitempty" xml:"SrcTableName,omitempty"`
	MaskingCount    *int64  `json:"MaskingCount,omitempty" xml:"MaskingCount,omitempty"`
	Percentage      *int32  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	DstType         *int32  `json:"DstType,omitempty" xml:"DstType,omitempty"`
	FailMsg         *string `json:"FailMsg,omitempty" xml:"FailMsg,omitempty"`
	FailCode        *string `json:"FailCode,omitempty" xml:"FailCode,omitempty"`
	ConflictCount   *int64  `json:"ConflictCount,omitempty" xml:"ConflictCount,omitempty"`
	DstTypeCode     *string `json:"DstTypeCode,omitempty" xml:"DstTypeCode,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RunIndex        *int32  `json:"RunIndex,omitempty" xml:"RunIndex,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	HasSubProcess   *int32  `json:"HasSubProcess,omitempty" xml:"HasSubProcess,omitempty"`
	HasDownloadFile *int32  `json:"HasDownloadFile,omitempty" xml:"HasDownloadFile,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	SrcTypeCode     *string `json:"SrcTypeCode,omitempty" xml:"SrcTypeCode,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDataMaskingRunHistoryResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingRunHistoryResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetStatus(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetType(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.Type = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetSrcType(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.SrcType = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetSrcTableName(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.SrcTableName = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetMaskingCount(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.MaskingCount = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetPercentage(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.Percentage = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetDstType(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.DstType = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetFailMsg(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.FailMsg = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetFailCode(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.FailCode = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetConflictCount(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.ConflictCount = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetDstTypeCode(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.DstTypeCode = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetEndTime(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetRunIndex(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.RunIndex = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetStartTime(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetHasSubProcess(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.HasSubProcess = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetHasDownloadFile(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.HasDownloadFile = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetTaskId(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetSrcTypeCode(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.SrcTypeCode = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetId(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.Id = &v
	return s
}

type DescribeDataMaskingRunHistoryResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataMaskingRunHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataMaskingRunHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingRunHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingRunHistoryResponse) SetHeaders(v map[string]*string) *DescribeDataMaskingRunHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponse) SetBody(v *DescribeDataMaskingRunHistoryResponseBody) *DescribeDataMaskingRunHistoryResponse {
	s.Body = v
	return s
}

type DescribeDataMaskingTasksRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SearchKey   *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DstType     *int32  `json:"DstType,omitempty" xml:"DstType,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeDataMaskingTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingTasksRequest) SetLang(v string) *DescribeDataMaskingTasksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetSearchKey(v string) *DescribeDataMaskingTasksRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetStartTime(v int64) *DescribeDataMaskingTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetEndTime(v int64) *DescribeDataMaskingTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetDstType(v int32) *DescribeDataMaskingTasksRequest {
	s.DstType = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetPageSize(v int32) *DescribeDataMaskingTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetCurrentPage(v int32) *DescribeDataMaskingTasksRequest {
	s.CurrentPage = &v
	return s
}

type DescribeDataMaskingTasksResponseBody struct {
	CurrentPage *int32                                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeDataMaskingTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeDataMaskingTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingTasksResponseBody) SetCurrentPage(v int32) *DescribeDataMaskingTasksResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) SetRequestId(v string) *DescribeDataMaskingTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) SetPageSize(v int32) *DescribeDataMaskingTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) SetTotalCount(v int32) *DescribeDataMaskingTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) SetItems(v []*DescribeDataMaskingTasksResponseBodyItems) *DescribeDataMaskingTasksResponseBody {
	s.Items = v
	return s
}

type DescribeDataMaskingTasksResponseBodyItems struct {
	Status             *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Owner              *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	TaskName           *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	SrcType            *int32  `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	DstType            *int32  `json:"DstType,omitempty" xml:"DstType,omitempty"`
	HasUnfinishProcess *bool   `json:"HasUnfinishProcess,omitempty" xml:"HasUnfinishProcess,omitempty"`
	OriginalTable      *bool   `json:"OriginalTable,omitempty" xml:"OriginalTable,omitempty"`
	TriggerType        *int32  `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	DstTypeCode        *string `json:"DstTypeCode,omitempty" xml:"DstTypeCode,omitempty"`
	RunCount           *int32  `json:"RunCount,omitempty" xml:"RunCount,omitempty"`
	GmtCreate          *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	TaskId             *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	DstPath            *string `json:"DstPath,omitempty" xml:"DstPath,omitempty"`
	SrcTypeCode        *string `json:"SrcTypeCode,omitempty" xml:"SrcTypeCode,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	SrcPath            *string `json:"SrcPath,omitempty" xml:"SrcPath,omitempty"`
}

func (s DescribeDataMaskingTasksResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetStatus(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetOwner(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetTaskName(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.TaskName = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetSrcType(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.SrcType = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetDstType(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.DstType = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetHasUnfinishProcess(v bool) *DescribeDataMaskingTasksResponseBodyItems {
	s.HasUnfinishProcess = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetOriginalTable(v bool) *DescribeDataMaskingTasksResponseBodyItems {
	s.OriginalTable = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetTriggerType(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.TriggerType = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetDstTypeCode(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.DstTypeCode = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetRunCount(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.RunCount = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetGmtCreate(v int64) *DescribeDataMaskingTasksResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetTaskId(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetDstPath(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.DstPath = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetSrcTypeCode(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.SrcTypeCode = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetId(v int64) *DescribeDataMaskingTasksResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetSrcPath(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.SrcPath = &v
	return s
}

type DescribeDataMaskingTasksResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataMaskingTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataMaskingTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingTasksResponse) SetHeaders(v map[string]*string) *DescribeDataMaskingTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataMaskingTasksResponse) SetBody(v *DescribeDataMaskingTasksResponseBody) *DescribeDataMaskingTasksResponse {
	s.Body = v
	return s
}

type DescribeDataTotalCountRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	CountType   *int32  `json:"CountType,omitempty" xml:"CountType,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	InstanceId  *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	FeatureType *int32  `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
}

func (s DescribeDataTotalCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountRequest) SetLang(v string) *DescribeDataTotalCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataTotalCountRequest) SetCountType(v int32) *DescribeDataTotalCountRequest {
	s.CountType = &v
	return s
}

func (s *DescribeDataTotalCountRequest) SetProductCode(v string) *DescribeDataTotalCountRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeDataTotalCountRequest) SetInstanceId(v int64) *DescribeDataTotalCountRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataTotalCountRequest) SetFeatureType(v int32) *DescribeDataTotalCountRequest {
	s.FeatureType = &v
	return s
}

type DescribeDataTotalCountResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DataCount *DescribeDataTotalCountResponseBodyDataCount `json:"DataCount,omitempty" xml:"DataCount,omitempty" type:"Struct"`
}

func (s DescribeDataTotalCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBody) SetRequestId(v string) *DescribeDataTotalCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataTotalCountResponseBody) SetDataCount(v *DescribeDataTotalCountResponseBodyDataCount) *DescribeDataTotalCountResponseBody {
	s.DataCount = v
	return s
}

type DescribeDataTotalCountResponseBodyDataCount struct {
	Instance *DescribeDataTotalCountResponseBodyDataCountInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	Table    *DescribeDataTotalCountResponseBodyDataCountTable    `json:"Table,omitempty" xml:"Table,omitempty" type:"Struct"`
	Package  *DescribeDataTotalCountResponseBodyDataCountPackage  `json:"Package,omitempty" xml:"Package,omitempty" type:"Struct"`
	Column   *DescribeDataTotalCountResponseBodyDataCountColumn   `json:"Column,omitempty" xml:"Column,omitempty" type:"Struct"`
	Oss      *DescribeDataTotalCountResponseBodyDataCountOss      `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s DescribeDataTotalCountResponseBodyDataCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCount) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCount) SetInstance(v *DescribeDataTotalCountResponseBodyDataCountInstance) *DescribeDataTotalCountResponseBodyDataCount {
	s.Instance = v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCount) SetTable(v *DescribeDataTotalCountResponseBodyDataCountTable) *DescribeDataTotalCountResponseBodyDataCount {
	s.Table = v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCount) SetPackage(v *DescribeDataTotalCountResponseBodyDataCountPackage) *DescribeDataTotalCountResponseBodyDataCount {
	s.Package = v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCount) SetColumn(v *DescribeDataTotalCountResponseBodyDataCountColumn) *DescribeDataTotalCountResponseBodyDataCount {
	s.Column = v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCount) SetOss(v *DescribeDataTotalCountResponseBodyDataCountOss) *DescribeDataTotalCountResponseBodyDataCount {
	s.Oss = v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountInstance struct {
	LastSensitiveCount *int64                                                              `json:"LastSensitiveCount,omitempty" xml:"LastSensitiveCount,omitempty"`
	TotalCount         *int64                                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SensitiveCount     *int64                                                              `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	Count              *int64                                                              `json:"Count,omitempty" xml:"Count,omitempty"`
	LastCount          *int64                                                              `json:"LastCount,omitempty" xml:"LastCount,omitempty"`
	RiskCountList      []*DescribeDataTotalCountResponseBodyDataCountInstanceRiskCountList `json:"RiskCountList,omitempty" xml:"RiskCountList,omitempty" type:"Repeated"`
}

func (s DescribeDataTotalCountResponseBodyDataCountInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountInstance) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountInstance) SetLastSensitiveCount(v int64) *DescribeDataTotalCountResponseBodyDataCountInstance {
	s.LastSensitiveCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountInstance) SetTotalCount(v int64) *DescribeDataTotalCountResponseBodyDataCountInstance {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountInstance) SetSensitiveCount(v int64) *DescribeDataTotalCountResponseBodyDataCountInstance {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountInstance) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountInstance {
	s.Count = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountInstance) SetLastCount(v int64) *DescribeDataTotalCountResponseBodyDataCountInstance {
	s.LastCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountInstance) SetRiskCountList(v []*DescribeDataTotalCountResponseBodyDataCountInstanceRiskCountList) *DescribeDataTotalCountResponseBodyDataCountInstance {
	s.RiskCountList = v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountInstanceRiskCountList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDataTotalCountResponseBodyDataCountInstanceRiskCountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountInstanceRiskCountList) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountInstanceRiskCountList) SetName(v string) *DescribeDataTotalCountResponseBodyDataCountInstanceRiskCountList {
	s.Name = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountInstanceRiskCountList) SetId(v int64) *DescribeDataTotalCountResponseBodyDataCountInstanceRiskCountList {
	s.Id = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountInstanceRiskCountList) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountInstanceRiskCountList {
	s.Count = &v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountTable struct {
	LastSensitiveCount *int64                                                           `json:"LastSensitiveCount,omitempty" xml:"LastSensitiveCount,omitempty"`
	TotalCount         *int64                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SensitiveCount     *int64                                                           `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	Count              *int64                                                           `json:"Count,omitempty" xml:"Count,omitempty"`
	LastCount          *int64                                                           `json:"LastCount,omitempty" xml:"LastCount,omitempty"`
	RiskCountList      []*DescribeDataTotalCountResponseBodyDataCountTableRiskCountList `json:"RiskCountList,omitempty" xml:"RiskCountList,omitempty" type:"Repeated"`
}

func (s DescribeDataTotalCountResponseBodyDataCountTable) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountTable) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountTable) SetLastSensitiveCount(v int64) *DescribeDataTotalCountResponseBodyDataCountTable {
	s.LastSensitiveCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountTable) SetTotalCount(v int64) *DescribeDataTotalCountResponseBodyDataCountTable {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountTable) SetSensitiveCount(v int64) *DescribeDataTotalCountResponseBodyDataCountTable {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountTable) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountTable {
	s.Count = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountTable) SetLastCount(v int64) *DescribeDataTotalCountResponseBodyDataCountTable {
	s.LastCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountTable) SetRiskCountList(v []*DescribeDataTotalCountResponseBodyDataCountTableRiskCountList) *DescribeDataTotalCountResponseBodyDataCountTable {
	s.RiskCountList = v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountTableRiskCountList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDataTotalCountResponseBodyDataCountTableRiskCountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountTableRiskCountList) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountTableRiskCountList) SetName(v string) *DescribeDataTotalCountResponseBodyDataCountTableRiskCountList {
	s.Name = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountTableRiskCountList) SetId(v int64) *DescribeDataTotalCountResponseBodyDataCountTableRiskCountList {
	s.Id = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountTableRiskCountList) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountTableRiskCountList {
	s.Count = &v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountPackage struct {
	LastSensitiveCount *int64                                                             `json:"LastSensitiveCount,omitempty" xml:"LastSensitiveCount,omitempty"`
	TotalCount         *int64                                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SensitiveCount     *int64                                                             `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	Count              *int64                                                             `json:"Count,omitempty" xml:"Count,omitempty"`
	LastCount          *int64                                                             `json:"LastCount,omitempty" xml:"LastCount,omitempty"`
	RiskCountList      []*DescribeDataTotalCountResponseBodyDataCountPackageRiskCountList `json:"RiskCountList,omitempty" xml:"RiskCountList,omitempty" type:"Repeated"`
}

func (s DescribeDataTotalCountResponseBodyDataCountPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountPackage) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountPackage) SetLastSensitiveCount(v int64) *DescribeDataTotalCountResponseBodyDataCountPackage {
	s.LastSensitiveCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountPackage) SetTotalCount(v int64) *DescribeDataTotalCountResponseBodyDataCountPackage {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountPackage) SetSensitiveCount(v int64) *DescribeDataTotalCountResponseBodyDataCountPackage {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountPackage) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountPackage {
	s.Count = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountPackage) SetLastCount(v int64) *DescribeDataTotalCountResponseBodyDataCountPackage {
	s.LastCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountPackage) SetRiskCountList(v []*DescribeDataTotalCountResponseBodyDataCountPackageRiskCountList) *DescribeDataTotalCountResponseBodyDataCountPackage {
	s.RiskCountList = v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountPackageRiskCountList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDataTotalCountResponseBodyDataCountPackageRiskCountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountPackageRiskCountList) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountPackageRiskCountList) SetName(v string) *DescribeDataTotalCountResponseBodyDataCountPackageRiskCountList {
	s.Name = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountPackageRiskCountList) SetId(v int64) *DescribeDataTotalCountResponseBodyDataCountPackageRiskCountList {
	s.Id = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountPackageRiskCountList) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountPackageRiskCountList {
	s.Count = &v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountColumn struct {
	LastSensitiveCount *int64                                                            `json:"LastSensitiveCount,omitempty" xml:"LastSensitiveCount,omitempty"`
	TotalCount         *int64                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SensitiveCount     *int64                                                            `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	Count              *int64                                                            `json:"Count,omitempty" xml:"Count,omitempty"`
	LastCount          *int64                                                            `json:"LastCount,omitempty" xml:"LastCount,omitempty"`
	RiskCountList      []*DescribeDataTotalCountResponseBodyDataCountColumnRiskCountList `json:"RiskCountList,omitempty" xml:"RiskCountList,omitempty" type:"Repeated"`
	ColumnRuleTop      []*DescribeDataTotalCountResponseBodyDataCountColumnColumnRuleTop `json:"ColumnRuleTop,omitempty" xml:"ColumnRuleTop,omitempty" type:"Repeated"`
}

func (s DescribeDataTotalCountResponseBodyDataCountColumn) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountColumn) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumn) SetLastSensitiveCount(v int64) *DescribeDataTotalCountResponseBodyDataCountColumn {
	s.LastSensitiveCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumn) SetTotalCount(v int64) *DescribeDataTotalCountResponseBodyDataCountColumn {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumn) SetSensitiveCount(v int64) *DescribeDataTotalCountResponseBodyDataCountColumn {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumn) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountColumn {
	s.Count = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumn) SetLastCount(v int64) *DescribeDataTotalCountResponseBodyDataCountColumn {
	s.LastCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumn) SetRiskCountList(v []*DescribeDataTotalCountResponseBodyDataCountColumnRiskCountList) *DescribeDataTotalCountResponseBodyDataCountColumn {
	s.RiskCountList = v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumn) SetColumnRuleTop(v []*DescribeDataTotalCountResponseBodyDataCountColumnColumnRuleTop) *DescribeDataTotalCountResponseBodyDataCountColumn {
	s.ColumnRuleTop = v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountColumnRiskCountList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDataTotalCountResponseBodyDataCountColumnRiskCountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountColumnRiskCountList) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumnRiskCountList) SetName(v string) *DescribeDataTotalCountResponseBodyDataCountColumnRiskCountList {
	s.Name = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumnRiskCountList) SetId(v int64) *DescribeDataTotalCountResponseBodyDataCountColumnRiskCountList {
	s.Id = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumnRiskCountList) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountColumnRiskCountList {
	s.Count = &v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountColumnColumnRuleTop struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDataTotalCountResponseBodyDataCountColumnColumnRuleTop) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountColumnColumnRuleTop) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumnColumnRuleTop) SetName(v string) *DescribeDataTotalCountResponseBodyDataCountColumnColumnRuleTop {
	s.Name = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumnColumnRuleTop) SetId(v int64) *DescribeDataTotalCountResponseBodyDataCountColumnColumnRuleTop {
	s.Id = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountColumnColumnRuleTop) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountColumnColumnRuleTop {
	s.Count = &v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountOss struct {
	LastSensitiveCount *int64                                                         `json:"LastSensitiveCount,omitempty" xml:"LastSensitiveCount,omitempty"`
	TotalCount         *int64                                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SensitiveCount     *int64                                                         `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	Count              *int64                                                         `json:"Count,omitempty" xml:"Count,omitempty"`
	LastCount          *int64                                                         `json:"LastCount,omitempty" xml:"LastCount,omitempty"`
	RiskCountList      []*DescribeDataTotalCountResponseBodyDataCountOssRiskCountList `json:"RiskCountList,omitempty" xml:"RiskCountList,omitempty" type:"Repeated"`
	OssRuleTop         []*DescribeDataTotalCountResponseBodyDataCountOssOssRuleTop    `json:"OssRuleTop,omitempty" xml:"OssRuleTop,omitempty" type:"Repeated"`
}

func (s DescribeDataTotalCountResponseBodyDataCountOss) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountOss) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountOss) SetLastSensitiveCount(v int64) *DescribeDataTotalCountResponseBodyDataCountOss {
	s.LastSensitiveCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountOss) SetTotalCount(v int64) *DescribeDataTotalCountResponseBodyDataCountOss {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountOss) SetSensitiveCount(v int64) *DescribeDataTotalCountResponseBodyDataCountOss {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountOss) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountOss {
	s.Count = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountOss) SetLastCount(v int64) *DescribeDataTotalCountResponseBodyDataCountOss {
	s.LastCount = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountOss) SetRiskCountList(v []*DescribeDataTotalCountResponseBodyDataCountOssRiskCountList) *DescribeDataTotalCountResponseBodyDataCountOss {
	s.RiskCountList = v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountOss) SetOssRuleTop(v []*DescribeDataTotalCountResponseBodyDataCountOssOssRuleTop) *DescribeDataTotalCountResponseBodyDataCountOss {
	s.OssRuleTop = v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountOssRiskCountList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDataTotalCountResponseBodyDataCountOssRiskCountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountOssRiskCountList) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountOssRiskCountList) SetName(v string) *DescribeDataTotalCountResponseBodyDataCountOssRiskCountList {
	s.Name = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountOssRiskCountList) SetId(v int64) *DescribeDataTotalCountResponseBodyDataCountOssRiskCountList {
	s.Id = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountOssRiskCountList) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountOssRiskCountList {
	s.Count = &v
	return s
}

type DescribeDataTotalCountResponseBodyDataCountOssOssRuleTop struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDataTotalCountResponseBodyDataCountOssOssRuleTop) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponseBodyDataCountOssOssRuleTop) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponseBodyDataCountOssOssRuleTop) SetName(v string) *DescribeDataTotalCountResponseBodyDataCountOssOssRuleTop {
	s.Name = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountOssOssRuleTop) SetId(v int64) *DescribeDataTotalCountResponseBodyDataCountOssOssRuleTop {
	s.Id = &v
	return s
}

func (s *DescribeDataTotalCountResponseBodyDataCountOssOssRuleTop) SetCount(v int64) *DescribeDataTotalCountResponseBodyDataCountOssOssRuleTop {
	s.Count = &v
	return s
}

type DescribeDataTotalCountResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDataTotalCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataTotalCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataTotalCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataTotalCountResponse) SetHeaders(v map[string]*string) *DescribeDataTotalCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataTotalCountResponse) SetBody(v *DescribeDataTotalCountResponseBody) *DescribeDataTotalCountResponse {
	s.Body = v
	return s
}

type DescribeEventDetailRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEventDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailRequest) SetLang(v string) *DescribeEventDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventDetailRequest) SetId(v int64) *DescribeEventDetailRequest {
	s.Id = &v
	return s
}

type DescribeEventDetailResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Event     *DescribeEventDetailResponseBodyEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Struct"`
}

func (s DescribeEventDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBody) SetRequestId(v string) *DescribeEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventDetailResponseBody) SetEvent(v *DescribeEventDetailResponseBodyEvent) *DescribeEventDetailResponseBody {
	s.Event = v
	return s
}

type DescribeEventDetailResponseBodyEvent struct {
	DisplayName     *string                                               `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Status          *int32                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	DealReason      *string                                               `json:"DealReason,omitempty" xml:"DealReason,omitempty"`
	UserId          *int64                                                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	StatusName      *string                                               `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	DealTime        *int64                                                `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	DealLoginName   *string                                               `json:"DealLoginName,omitempty" xml:"DealLoginName,omitempty"`
	SubTypeName     *string                                               `json:"SubTypeName,omitempty" xml:"SubTypeName,omitempty"`
	Backed          *bool                                                 `json:"Backed,omitempty" xml:"Backed,omitempty"`
	DataInstance    *string                                               `json:"DataInstance,omitempty" xml:"DataInstance,omitempty"`
	EventTime       *int64                                                `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	LoginName       *string                                               `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	SubTypeCode     *string                                               `json:"SubTypeCode,omitempty" xml:"SubTypeCode,omitempty"`
	LogDetail       *string                                               `json:"LogDetail,omitempty" xml:"LogDetail,omitempty"`
	TypeCode        *string                                               `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	AlertTime       *int64                                                `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	DealUserId      *int64                                                `json:"DealUserId,omitempty" xml:"DealUserId,omitempty"`
	TypeName        *string                                               `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	DealDisplayName *string                                               `json:"DealDisplayName,omitempty" xml:"DealDisplayName,omitempty"`
	Id              *int64                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	ProductCode     *string                                               `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	HandleInfoList  []*DescribeEventDetailResponseBodyEventHandleInfoList `json:"HandleInfoList,omitempty" xml:"HandleInfoList,omitempty" type:"Repeated"`
	Detail          *DescribeEventDetailResponseBodyEventDetail           `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
}

func (s DescribeEventDetailResponseBodyEvent) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEvent) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEvent) SetDisplayName(v string) *DescribeEventDetailResponseBodyEvent {
	s.DisplayName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetStatus(v int32) *DescribeEventDetailResponseBodyEvent {
	s.Status = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealReason(v string) *DescribeEventDetailResponseBodyEvent {
	s.DealReason = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetUserId(v int64) *DescribeEventDetailResponseBodyEvent {
	s.UserId = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetStatusName(v string) *DescribeEventDetailResponseBodyEvent {
	s.StatusName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealTime(v int64) *DescribeEventDetailResponseBodyEvent {
	s.DealTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealLoginName(v string) *DescribeEventDetailResponseBodyEvent {
	s.DealLoginName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetSubTypeName(v string) *DescribeEventDetailResponseBodyEvent {
	s.SubTypeName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetBacked(v bool) *DescribeEventDetailResponseBodyEvent {
	s.Backed = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDataInstance(v string) *DescribeEventDetailResponseBodyEvent {
	s.DataInstance = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetEventTime(v int64) *DescribeEventDetailResponseBodyEvent {
	s.EventTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetLoginName(v string) *DescribeEventDetailResponseBodyEvent {
	s.LoginName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetSubTypeCode(v string) *DescribeEventDetailResponseBodyEvent {
	s.SubTypeCode = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetLogDetail(v string) *DescribeEventDetailResponseBodyEvent {
	s.LogDetail = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetTypeCode(v string) *DescribeEventDetailResponseBodyEvent {
	s.TypeCode = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetAlertTime(v int64) *DescribeEventDetailResponseBodyEvent {
	s.AlertTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealUserId(v int64) *DescribeEventDetailResponseBodyEvent {
	s.DealUserId = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetTypeName(v string) *DescribeEventDetailResponseBodyEvent {
	s.TypeName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealDisplayName(v string) *DescribeEventDetailResponseBodyEvent {
	s.DealDisplayName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetId(v int64) *DescribeEventDetailResponseBodyEvent {
	s.Id = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetProductCode(v string) *DescribeEventDetailResponseBodyEvent {
	s.ProductCode = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetHandleInfoList(v []*DescribeEventDetailResponseBodyEventHandleInfoList) *DescribeEventDetailResponseBodyEvent {
	s.HandleInfoList = v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDetail(v *DescribeEventDetailResponseBodyEventDetail) *DescribeEventDetailResponseBodyEvent {
	s.Detail = v
	return s
}

type DescribeEventDetailResponseBodyEventHandleInfoList struct {
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	EnableTime   *int64  `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
	HandlerValue *int32  `json:"HandlerValue,omitempty" xml:"HandlerValue,omitempty"`
	DisableTime  *int64  `json:"DisableTime,omitempty" xml:"DisableTime,omitempty"`
	HandlerName  *string `json:"HandlerName,omitempty" xml:"HandlerName,omitempty"`
	HandlerType  *string `json:"HandlerType,omitempty" xml:"HandlerType,omitempty"`
	CurrentValue *string `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEventDetailResponseBodyEventHandleInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventHandleInfoList) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetStatus(v int32) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.Status = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetEnableTime(v int64) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.EnableTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetHandlerValue(v int32) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.HandlerValue = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetDisableTime(v int64) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.DisableTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetHandlerName(v string) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.HandlerName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetHandlerType(v string) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.HandlerType = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetCurrentValue(v string) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.CurrentValue = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetId(v int64) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.Id = &v
	return s
}

type DescribeEventDetailResponseBodyEventDetail struct {
	Content      []*DescribeEventDetailResponseBodyEventDetailContent      `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	Chart        []*DescribeEventDetailResponseBodyEventDetailChart        `json:"Chart,omitempty" xml:"Chart,omitempty" type:"Repeated"`
	ResourceInfo []*DescribeEventDetailResponseBodyEventDetailResourceInfo `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty" type:"Repeated"`
}

func (s DescribeEventDetailResponseBodyEventDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetail) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetail) SetContent(v []*DescribeEventDetailResponseBodyEventDetailContent) *DescribeEventDetailResponseBodyEventDetail {
	s.Content = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetail) SetChart(v []*DescribeEventDetailResponseBodyEventDetailChart) *DescribeEventDetailResponseBodyEventDetail {
	s.Chart = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetail) SetResourceInfo(v []*DescribeEventDetailResponseBodyEventDetailResourceInfo) *DescribeEventDetailResponseBodyEventDetail {
	s.ResourceInfo = v
	return s
}

type DescribeEventDetailResponseBodyEventDetailContent struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEventDetailResponseBodyEventDetailContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetailContent) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetailContent) SetLabel(v string) *DescribeEventDetailResponseBodyEventDetailContent {
	s.Label = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailContent) SetValue(v string) *DescribeEventDetailResponseBodyEventDetailContent {
	s.Value = &v
	return s
}

type DescribeEventDetailResponseBodyEventDetailChart struct {
	Type   *string                                              `json:"Type,omitempty" xml:"Type,omitempty"`
	Label  *string                                              `json:"Label,omitempty" xml:"Label,omitempty"`
	XLabel *string                                              `json:"XLabel,omitempty" xml:"XLabel,omitempty"`
	YLabel *string                                              `json:"YLabel,omitempty" xml:"YLabel,omitempty"`
	Data   *DescribeEventDetailResponseBodyEventDetailChartData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeEventDetailResponseBodyEventDetailChart) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetailChart) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetType(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.Type = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetLabel(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.Label = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetXLabel(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.XLabel = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetYLabel(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.YLabel = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetData(v *DescribeEventDetailResponseBodyEventDetailChartData) *DescribeEventDetailResponseBodyEventDetailChart {
	s.Data = v
	return s
}

type DescribeEventDetailResponseBodyEventDetailChartData struct {
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
	X *string `json:"X,omitempty" xml:"X,omitempty"`
}

func (s DescribeEventDetailResponseBodyEventDetailChartData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetailChartData) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) SetY(v string) *DescribeEventDetailResponseBodyEventDetailChartData {
	s.Y = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) SetX(v string) *DescribeEventDetailResponseBodyEventDetailChartData {
	s.X = &v
	return s
}

type DescribeEventDetailResponseBodyEventDetailResourceInfo struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEventDetailResponseBodyEventDetailResourceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetailResourceInfo) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetailResourceInfo) SetLabel(v string) *DescribeEventDetailResponseBodyEventDetailResourceInfo {
	s.Label = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailResourceInfo) SetValue(v string) *DescribeEventDetailResponseBodyEventDetailResourceInfo {
	s.Value = &v
	return s
}

type DescribeEventDetailResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponse) SetHeaders(v map[string]*string) *DescribeEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventDetailResponse) SetBody(v *DescribeEventDetailResponseBody) *DescribeEventDetailResponse {
	s.Body = v
	return s
}

type DescribeEventsRequest struct {
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	UserId            *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DealUserId        *string `json:"DealUserId,omitempty" xml:"DealUserId,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TypeCode          *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	SubTypeCode       *string `json:"SubTypeCode,omitempty" xml:"SubTypeCode,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage       *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceName      *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	TargetProductCode *string `json:"TargetProductCode,omitempty" xml:"TargetProductCode,omitempty"`
	UserName          *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) SetLang(v string) *DescribeEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventsRequest) SetId(v int64) *DescribeEventsRequest {
	s.Id = &v
	return s
}

func (s *DescribeEventsRequest) SetUserId(v int64) *DescribeEventsRequest {
	s.UserId = &v
	return s
}

func (s *DescribeEventsRequest) SetDealUserId(v string) *DescribeEventsRequest {
	s.DealUserId = &v
	return s
}

func (s *DescribeEventsRequest) SetStatus(v string) *DescribeEventsRequest {
	s.Status = &v
	return s
}

func (s *DescribeEventsRequest) SetTypeCode(v string) *DescribeEventsRequest {
	s.TypeCode = &v
	return s
}

func (s *DescribeEventsRequest) SetSubTypeCode(v string) *DescribeEventsRequest {
	s.SubTypeCode = &v
	return s
}

func (s *DescribeEventsRequest) SetStartTime(v string) *DescribeEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventsRequest) SetEndTime(v string) *DescribeEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsRequest) SetPageSize(v int32) *DescribeEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsRequest) SetCurrentPage(v int32) *DescribeEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventsRequest) SetInstanceName(v string) *DescribeEventsRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeEventsRequest) SetProductCode(v string) *DescribeEventsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeEventsRequest) SetTargetProductCode(v string) *DescribeEventsRequest {
	s.TargetProductCode = &v
	return s
}

func (s *DescribeEventsRequest) SetUserName(v string) *DescribeEventsRequest {
	s.UserName = &v
	return s
}

type DescribeEventsResponseBody struct {
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeEventsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) SetCurrentPage(v int32) *DescribeEventsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventsResponseBody) SetRequestId(v string) *DescribeEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsResponseBody) SetPageSize(v int32) *DescribeEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsResponseBody) SetTotalCount(v int32) *DescribeEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventsResponseBody) SetItems(v []*DescribeEventsResponseBodyItems) *DescribeEventsResponseBody {
	s.Items = v
	return s
}

type DescribeEventsResponseBodyItems struct {
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Status            *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WarnLevel         *int32  `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
	UserId            *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	StatusName        *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	DealTime          *int64  `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	DealLoginName     *string `json:"DealLoginName,omitempty" xml:"DealLoginName,omitempty"`
	SubTypeName       *string `json:"SubTypeName,omitempty" xml:"SubTypeName,omitempty"`
	Backed            *bool   `json:"Backed,omitempty" xml:"Backed,omitempty"`
	EventTime         *int64  `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	LoginName         *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	SubTypeCode       *string `json:"SubTypeCode,omitempty" xml:"SubTypeCode,omitempty"`
	TargetProductCode *string `json:"TargetProductCode,omitempty" xml:"TargetProductCode,omitempty"`
	TypeCode          *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	AlertTime         *int64  `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	DealUserId        *int64  `json:"DealUserId,omitempty" xml:"DealUserId,omitempty"`
	TypeName          *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	DealDisplayName   *string `json:"DealDisplayName,omitempty" xml:"DealDisplayName,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeEventsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyItems) SetDisplayName(v string) *DescribeEventsResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetStatus(v int32) *DescribeEventsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetWarnLevel(v int32) *DescribeEventsResponseBodyItems {
	s.WarnLevel = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetUserId(v int64) *DescribeEventsResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetStatusName(v string) *DescribeEventsResponseBodyItems {
	s.StatusName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealTime(v int64) *DescribeEventsResponseBodyItems {
	s.DealTime = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealLoginName(v string) *DescribeEventsResponseBodyItems {
	s.DealLoginName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetSubTypeName(v string) *DescribeEventsResponseBodyItems {
	s.SubTypeName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetBacked(v bool) *DescribeEventsResponseBodyItems {
	s.Backed = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetEventTime(v int64) *DescribeEventsResponseBodyItems {
	s.EventTime = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetLoginName(v string) *DescribeEventsResponseBodyItems {
	s.LoginName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetSubTypeCode(v string) *DescribeEventsResponseBodyItems {
	s.SubTypeCode = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetTargetProductCode(v string) *DescribeEventsResponseBodyItems {
	s.TargetProductCode = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetTypeCode(v string) *DescribeEventsResponseBodyItems {
	s.TypeCode = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetAlertTime(v int64) *DescribeEventsResponseBodyItems {
	s.AlertTime = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealUserId(v int64) *DescribeEventsResponseBodyItems {
	s.DealUserId = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetTypeName(v string) *DescribeEventsResponseBodyItems {
	s.TypeName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealDisplayName(v string) *DescribeEventsResponseBodyItems {
	s.DealDisplayName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetId(v int64) *DescribeEventsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetProductCode(v string) *DescribeEventsResponseBodyItems {
	s.ProductCode = &v
	return s
}

type DescribeEventsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponse) SetHeaders(v map[string]*string) *DescribeEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventsResponse) SetBody(v *DescribeEventsResponseBody) *DescribeEventsResponse {
	s.Body = v
	return s
}

type DescribeEventTopRequest struct {
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserId            *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DealUserId        *string `json:"DealUserId,omitempty" xml:"DealUserId,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TypeCode          *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	SubTypeCode       *string `json:"SubTypeCode,omitempty" xml:"SubTypeCode,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage       *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceName      *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	TargetProductCode *string `json:"TargetProductCode,omitempty" xml:"TargetProductCode,omitempty"`
	UserName          *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeEventTopRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTopRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventTopRequest) SetLang(v string) *DescribeEventTopRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventTopRequest) SetUserId(v int64) *DescribeEventTopRequest {
	s.UserId = &v
	return s
}

func (s *DescribeEventTopRequest) SetDealUserId(v string) *DescribeEventTopRequest {
	s.DealUserId = &v
	return s
}

func (s *DescribeEventTopRequest) SetStatus(v string) *DescribeEventTopRequest {
	s.Status = &v
	return s
}

func (s *DescribeEventTopRequest) SetTypeCode(v string) *DescribeEventTopRequest {
	s.TypeCode = &v
	return s
}

func (s *DescribeEventTopRequest) SetSubTypeCode(v string) *DescribeEventTopRequest {
	s.SubTypeCode = &v
	return s
}

func (s *DescribeEventTopRequest) SetStartTime(v string) *DescribeEventTopRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventTopRequest) SetEndTime(v string) *DescribeEventTopRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventTopRequest) SetPageSize(v int32) *DescribeEventTopRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventTopRequest) SetCurrentPage(v int32) *DescribeEventTopRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventTopRequest) SetInstanceName(v string) *DescribeEventTopRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeEventTopRequest) SetProductCode(v string) *DescribeEventTopRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeEventTopRequest) SetTargetProductCode(v string) *DescribeEventTopRequest {
	s.TargetProductCode = &v
	return s
}

func (s *DescribeEventTopRequest) SetUserName(v string) *DescribeEventTopRequest {
	s.UserName = &v
	return s
}

type DescribeEventTopResponseBody struct {
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DataCountList []*DescribeEventTopResponseBodyDataCountList `json:"DataCountList,omitempty" xml:"DataCountList,omitempty" type:"Repeated"`
}

func (s DescribeEventTopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTopResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventTopResponseBody) SetRequestId(v string) *DescribeEventTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventTopResponseBody) SetDataCountList(v []*DescribeEventTopResponseBodyDataCountList) *DescribeEventTopResponseBody {
	s.DataCountList = v
	return s
}

type DescribeEventTopResponseBodyDataCountList struct {
	EventName    *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	RemoteIp     *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ProductCode  *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeEventTopResponseBodyDataCountList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTopResponseBodyDataCountList) GoString() string {
	return s.String()
}

func (s *DescribeEventTopResponseBodyDataCountList) SetEventName(v string) *DescribeEventTopResponseBodyDataCountList {
	s.EventName = &v
	return s
}

func (s *DescribeEventTopResponseBodyDataCountList) SetRemoteIp(v string) *DescribeEventTopResponseBodyDataCountList {
	s.RemoteIp = &v
	return s
}

func (s *DescribeEventTopResponseBodyDataCountList) SetInstanceName(v string) *DescribeEventTopResponseBodyDataCountList {
	s.InstanceName = &v
	return s
}

func (s *DescribeEventTopResponseBodyDataCountList) SetUserName(v string) *DescribeEventTopResponseBodyDataCountList {
	s.UserName = &v
	return s
}

func (s *DescribeEventTopResponseBodyDataCountList) SetId(v int64) *DescribeEventTopResponseBodyDataCountList {
	s.Id = &v
	return s
}

func (s *DescribeEventTopResponseBodyDataCountList) SetProductCode(v string) *DescribeEventTopResponseBodyDataCountList {
	s.ProductCode = &v
	return s
}

type DescribeEventTopResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEventTopResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventTopResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTopResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventTopResponse) SetHeaders(v map[string]*string) *DescribeEventTopResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventTopResponse) SetBody(v *DescribeEventTopResponseBody) *DescribeEventTopResponse {
	s.Body = v
	return s
}

type DescribeEventTypesRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ParentTypeId *int64  `json:"ParentTypeId,omitempty" xml:"ParentTypeId,omitempty"`
	ResourceId   *int32  `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesRequest) SetLang(v string) *DescribeEventTypesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventTypesRequest) SetParentTypeId(v int64) *DescribeEventTypesRequest {
	s.ParentTypeId = &v
	return s
}

func (s *DescribeEventTypesRequest) SetResourceId(v int32) *DescribeEventTypesRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeEventTypesRequest) SetStatus(v int32) *DescribeEventTypesRequest {
	s.Status = &v
	return s
}

type DescribeEventTypesResponseBody struct {
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EventTypeList []*DescribeEventTypesResponseBodyEventTypeList `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Repeated"`
}

func (s DescribeEventTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesResponseBody) SetRequestId(v string) *DescribeEventTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventTypesResponseBody) SetEventTypeList(v []*DescribeEventTypesResponseBodyEventTypeList) *DescribeEventTypesResponseBody {
	s.EventTypeList = v
	return s
}

type DescribeEventTypesResponseBodyEventTypeList struct {
	Code        *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Description *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *int64                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	SubTypeList []*DescribeEventTypesResponseBodyEventTypeListSubTypeList `json:"SubTypeList,omitempty" xml:"SubTypeList,omitempty" type:"Repeated"`
}

func (s DescribeEventTypesResponseBodyEventTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesResponseBodyEventTypeList) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetCode(v string) *DescribeEventTypesResponseBodyEventTypeList {
	s.Code = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetDescription(v string) *DescribeEventTypesResponseBodyEventTypeList {
	s.Description = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetName(v string) *DescribeEventTypesResponseBodyEventTypeList {
	s.Name = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetId(v int64) *DescribeEventTypesResponseBodyEventTypeList {
	s.Id = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetSubTypeList(v []*DescribeEventTypesResponseBodyEventTypeListSubTypeList) *DescribeEventTypesResponseBodyEventTypeList {
	s.SubTypeList = v
	return s
}

type DescribeEventTypesResponseBodyEventTypeListSubTypeList struct {
	Status            *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ConfigContentType *int32  `json:"ConfigContentType,omitempty" xml:"ConfigContentType,omitempty"`
	EventHitCount     *int32  `json:"EventHitCount,omitempty" xml:"EventHitCount,omitempty"`
	ConfigValue       *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	ConfigCode        *string `json:"ConfigCode,omitempty" xml:"ConfigCode,omitempty"`
	Code              *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ConfigDescription *string `json:"ConfigDescription,omitempty" xml:"ConfigDescription,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AdaptedProduct    *string `json:"AdaptedProduct,omitempty" xml:"AdaptedProduct,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEventTypesResponseBodyEventTypeListSubTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesResponseBodyEventTypeListSubTypeList) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetStatus(v int32) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Status = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetDescription(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Description = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigContentType(v int32) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigContentType = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetEventHitCount(v int32) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.EventHitCount = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigValue(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigValue = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigCode(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigCode = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetCode(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Code = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigDescription(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigDescription = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetName(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Name = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetAdaptedProduct(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.AdaptedProduct = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetId(v int64) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Id = &v
	return s
}

type DescribeEventTypesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEventTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesResponse) SetHeaders(v map[string]*string) *DescribeEventTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventTypesResponse) SetBody(v *DescribeEventTypesResponseBody) *DescribeEventTypesResponse {
	s.Body = v
	return s
}

type DescribeEventTypesDetailRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeEventTypesDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesDetailRequest) SetLang(v string) *DescribeEventTypesDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventTypesDetailRequest) SetCode(v string) *DescribeEventTypesDetailRequest {
	s.Code = &v
	return s
}

type DescribeEventTypesDetailResponseBody struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserStatus *DescribeEventTypesDetailResponseBodyUserStatus `json:"UserStatus,omitempty" xml:"UserStatus,omitempty" type:"Struct"`
}

func (s DescribeEventTypesDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesDetailResponseBody) SetRequestId(v string) *DescribeEventTypesDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBody) SetUserStatus(v *DescribeEventTypesDetailResponseBodyUserStatus) *DescribeEventTypesDetailResponseBody {
	s.UserStatus = v
	return s
}

type DescribeEventTypesDetailResponseBodyUserStatus struct {
	Status            *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	RiskLevel         *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	ConfigContentType *int32  `json:"ConfigContentType,omitempty" xml:"ConfigContentType,omitempty"`
	DealSuggestion    *string `json:"DealSuggestion,omitempty" xml:"DealSuggestion,omitempty"`
	EventHitCount     *int32  `json:"EventHitCount,omitempty" xml:"EventHitCount,omitempty"`
	ConfigValue       *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	AccessCount       *int32  `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
	ConfigDescription *string `json:"ConfigDescription,omitempty" xml:"ConfigDescription,omitempty"`
	RiskName          *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	DescEvent         *string `json:"DescEvent,omitempty" xml:"DescEvent,omitempty"`
	RiskEvent         *string `json:"RiskEvent,omitempty" xml:"RiskEvent,omitempty"`
	ParentTypeName    *string `json:"ParentTypeName,omitempty" xml:"ParentTypeName,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DescModel         *string `json:"DescModel,omitempty" xml:"DescModel,omitempty"`
	ConfigCode        *string `json:"ConfigCode,omitempty" xml:"ConfigCode,omitempty"`
	Code              *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ParentTypeId      *int64  `json:"ParentTypeId,omitempty" xml:"ParentTypeId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AdaptedProduct    *string `json:"AdaptedProduct,omitempty" xml:"AdaptedProduct,omitempty"`
}

func (s DescribeEventTypesDetailResponseBodyUserStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesDetailResponseBodyUserStatus) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetStatus(v int32) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.Status = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetRiskLevel(v int32) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.RiskLevel = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetConfigContentType(v int32) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.ConfigContentType = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetDealSuggestion(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.DealSuggestion = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetEventHitCount(v int32) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.EventHitCount = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetConfigValue(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.ConfigValue = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetAccessCount(v int32) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.AccessCount = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetConfigDescription(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.ConfigDescription = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetRiskName(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.RiskName = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetDescEvent(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.DescEvent = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetRiskEvent(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.RiskEvent = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetParentTypeName(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.ParentTypeName = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetDescription(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.Description = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetDescModel(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.DescModel = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetConfigCode(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.ConfigCode = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetCode(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.Code = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetParentTypeId(v int64) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.ParentTypeId = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetName(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.Name = &v
	return s
}

func (s *DescribeEventTypesDetailResponseBodyUserStatus) SetAdaptedProduct(v string) *DescribeEventTypesDetailResponseBodyUserStatus {
	s.AdaptedProduct = &v
	return s
}

type DescribeEventTypesDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEventTypesDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventTypesDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesDetailResponse) SetHeaders(v map[string]*string) *DescribeEventTypesDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventTypesDetailResponse) SetBody(v *DescribeEventTypesDetailResponseBody) *DescribeEventTypesDetailResponse {
	s.Body = v
	return s
}

type DescribeInstancePortraitRequest struct {
	ProductId  *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	InstanceId *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemKeys   *string `json:"ItemKeys,omitempty" xml:"ItemKeys,omitempty"`
}

func (s DescribeInstancePortraitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePortraitRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancePortraitRequest) SetProductId(v int64) *DescribeInstancePortraitRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeInstancePortraitRequest) SetInstanceId(v int64) *DescribeInstancePortraitRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancePortraitRequest) SetItemKeys(v string) *DescribeInstancePortraitRequest {
	s.ItemKeys = &v
	return s
}

type DescribeInstancePortraitResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Content   []*DescribeInstancePortraitResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
}

func (s DescribeInstancePortraitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePortraitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancePortraitResponseBody) SetRequestId(v string) *DescribeInstancePortraitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancePortraitResponseBody) SetContent(v []*DescribeInstancePortraitResponseBodyContent) *DescribeInstancePortraitResponseBody {
	s.Content = v
	return s
}

type DescribeInstancePortraitResponseBodyContent struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ItemKey *string `json:"ItemKey,omitempty" xml:"ItemKey,omitempty"`
}

func (s DescribeInstancePortraitResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePortraitResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeInstancePortraitResponseBodyContent) SetContent(v string) *DescribeInstancePortraitResponseBodyContent {
	s.Content = &v
	return s
}

func (s *DescribeInstancePortraitResponseBodyContent) SetItemKey(v string) *DescribeInstancePortraitResponseBodyContent {
	s.ItemKey = &v
	return s
}

type DescribeInstancePortraitResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancePortraitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancePortraitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePortraitResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancePortraitResponse) SetHeaders(v map[string]*string) *DescribeInstancePortraitResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancePortraitResponse) SetBody(v *DescribeInstancePortraitResponseBody) *DescribeInstancePortraitResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductId           *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RiskLevelId         *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RuleId              *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	ServiceRegionId     *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	FeatureType         *int32  `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage         *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	LastFinishTimeStart *int64  `json:"LastFinishTimeStart,omitempty" xml:"LastFinishTimeStart,omitempty"`
	LastFinishTimeEnd   *int64  `json:"LastFinishTimeEnd,omitempty" xml:"LastFinishTimeEnd,omitempty"`
	UserId              *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType            *int32  `json:"UserType,omitempty" xml:"UserType,omitempty"`
	SensLevelName       *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	CheckStatus         *int32  `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	NameAccurate        *string `json:"NameAccurate,omitempty" xml:"NameAccurate,omitempty"`
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime             *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetLang(v string) *DescribeInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstancesRequest) SetName(v string) *DescribeInstancesRequest {
	s.Name = &v
	return s
}

func (s *DescribeInstancesRequest) SetProductId(v int64) *DescribeInstancesRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeInstancesRequest) SetRiskLevelId(v int64) *DescribeInstancesRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeInstancesRequest) SetRuleId(v int64) *DescribeInstancesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeInstancesRequest) SetServiceRegionId(v string) *DescribeInstancesRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetFeatureType(v int32) *DescribeInstancesRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeInstancesRequest) SetProductCode(v string) *DescribeInstancesRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetCurrentPage(v int32) *DescribeInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstancesRequest) SetLastFinishTimeStart(v int64) *DescribeInstancesRequest {
	s.LastFinishTimeStart = &v
	return s
}

func (s *DescribeInstancesRequest) SetLastFinishTimeEnd(v int64) *DescribeInstancesRequest {
	s.LastFinishTimeEnd = &v
	return s
}

func (s *DescribeInstancesRequest) SetUserId(v int64) *DescribeInstancesRequest {
	s.UserId = &v
	return s
}

func (s *DescribeInstancesRequest) SetUserType(v int32) *DescribeInstancesRequest {
	s.UserType = &v
	return s
}

func (s *DescribeInstancesRequest) SetSensLevelName(v string) *DescribeInstancesRequest {
	s.SensLevelName = &v
	return s
}

func (s *DescribeInstancesRequest) SetSource(v string) *DescribeInstancesRequest {
	s.Source = &v
	return s
}

func (s *DescribeInstancesRequest) SetCheckStatus(v int32) *DescribeInstancesRequest {
	s.CheckStatus = &v
	return s
}

func (s *DescribeInstancesRequest) SetNameAccurate(v string) *DescribeInstancesRequest {
	s.NameAccurate = &v
	return s
}

func (s *DescribeInstancesRequest) SetStartTime(v int64) *DescribeInstancesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancesRequest) SetEndTime(v int64) *DescribeInstancesRequest {
	s.EndTime = &v
	return s
}

type DescribeInstancesResponseBody struct {
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetCurrentPage(v int32) *DescribeInstancesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetItems(v []*DescribeInstancesResponseBodyItems) *DescribeInstancesResponseBody {
	s.Items = v
	return s
}

type DescribeInstancesResponseBodyItems struct {
	CreationTime        *int64  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Acl                 *string `json:"Acl,omitempty" xml:"Acl,omitempty"`
	LastFinishTime      *int64  `json:"LastFinishTime,omitempty" xml:"LastFinishTime,omitempty"`
	Owner               *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	TenantName          *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	Protection          *bool   `json:"Protection,omitempty" xml:"Protection,omitempty"`
	DepartName          *string `json:"DepartName,omitempty" xml:"DepartName,omitempty"`
	Labelsec            *bool   `json:"Labelsec,omitempty" xml:"Labelsec,omitempty"`
	RiskLevelId         *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	ProductId           *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TotalCount          *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	RuleName            *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sensitive           *bool   `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	RiskLevelName       *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	OdpsRiskLevelName   *string `json:"OdpsRiskLevelName,omitempty" xml:"OdpsRiskLevelName,omitempty"`
	SensitiveCount      *int32  `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeInstancesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyItems) SetCreationTime(v int64) *DescribeInstancesResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetAcl(v string) *DescribeInstancesResponseBodyItems {
	s.Acl = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetLastFinishTime(v int64) *DescribeInstancesResponseBodyItems {
	s.LastFinishTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetOwner(v string) *DescribeInstancesResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetTenantName(v string) *DescribeInstancesResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetProtection(v bool) *DescribeInstancesResponseBodyItems {
	s.Protection = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetDepartName(v string) *DescribeInstancesResponseBodyItems {
	s.DepartName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetLabelsec(v bool) *DescribeInstancesResponseBodyItems {
	s.Labelsec = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetRiskLevelId(v int64) *DescribeInstancesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetProductId(v string) *DescribeInstancesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetName(v string) *DescribeInstancesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetTotalCount(v int32) *DescribeInstancesResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetInstanceDescription(v string) *DescribeInstancesResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetRuleName(v string) *DescribeInstancesResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetSensitive(v bool) *DescribeInstancesResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetRiskLevelName(v string) *DescribeInstancesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetOdpsRiskLevelName(v string) *DescribeInstancesResponseBodyItems {
	s.OdpsRiskLevelName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetSensitiveCount(v int32) *DescribeInstancesResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetId(v int64) *DescribeInstancesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetProductCode(v string) *DescribeInstancesResponseBodyItems {
	s.ProductCode = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeInstanceSourcesRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	EngineType      *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DbName          *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ProductId       *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Authed          *bool   `json:"Authed,omitempty" xml:"Authed,omitempty"`
}

func (s DescribeInstanceSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSourcesRequest) SetLang(v string) *DescribeInstanceSourcesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetPageSize(v int32) *DescribeInstanceSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetCurrentPage(v int32) *DescribeInstanceSourcesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetServiceRegionId(v string) *DescribeInstanceSourcesRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetEngineType(v string) *DescribeInstanceSourcesRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetInstanceId(v string) *DescribeInstanceSourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetDbName(v string) *DescribeInstanceSourcesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetProductId(v int64) *DescribeInstanceSourcesRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetAuthed(v bool) *DescribeInstanceSourcesRequest {
	s.Authed = &v
	return s
}

type DescribeInstanceSourcesResponseBody struct {
	CurrentPage *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeInstanceSourcesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSourcesResponseBody) SetCurrentPage(v int32) *DescribeInstanceSourcesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) SetRequestId(v string) *DescribeInstanceSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) SetPageSize(v int32) *DescribeInstanceSourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) SetTotalCount(v int32) *DescribeInstanceSourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) SetItems(v []*DescribeInstanceSourcesResponseBodyItems) *DescribeInstanceSourcesResponseBody {
	s.Items = v
	return s
}

type DescribeInstanceSourcesResponseBodyItems struct {
	LastModifyUserId    *string `json:"LastModifyUserId,omitempty" xml:"LastModifyUserId,omitempty"`
	PasswordStatus      *int32  `json:"PasswordStatus,omitempty" xml:"PasswordStatus,omitempty"`
	EngineType          *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	TenantName          *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	DataLimitId         *int64  `json:"DataLimitId,omitempty" xml:"DataLimitId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DbName              *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	LastModifyTime      *int64  `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	RegionName          *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	CanModifyUserName   *bool   `json:"CanModifyUserName,omitempty" xml:"CanModifyUserName,omitempty"`
	LogStoreDay         *int32  `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	GmtCreate           *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	AutoScan            *int32  `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	ProductId           *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	InstanceSize        *int64  `json:"InstanceSize,omitempty" xml:"InstanceSize,omitempty"`
	UserName            *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	AuditStatus         *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TenantId            *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Enable              *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeInstanceSourcesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSourcesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetLastModifyUserId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.LastModifyUserId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetPasswordStatus(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.PasswordStatus = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetEngineType(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetTenantName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetInstanceId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetInstanceDescription(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetDataLimitId(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.DataLimitId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetRegionId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetDbName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.DbName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetLastModifyTime(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.LastModifyTime = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetRegionName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.RegionName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetCanModifyUserName(v bool) *DescribeInstanceSourcesResponseBodyItems {
	s.CanModifyUserName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetLogStoreDay(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.LogStoreDay = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetGmtCreate(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetAutoScan(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.AutoScan = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetProductId(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetInstanceSize(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.InstanceSize = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetUserName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.UserName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetAuditStatus(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.AuditStatus = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetId(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetTenantId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.TenantId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetEnable(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.Enable = &v
	return s
}

type DescribeInstanceSourcesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSourcesResponse) SetHeaders(v map[string]*string) *DescribeInstanceSourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSourcesResponse) SetBody(v *DescribeInstanceSourcesResponseBody) *DescribeInstanceSourcesResponse {
	s.Body = v
	return s
}

type DescribeOssObjectDetailRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeOssObjectDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailRequest) SetLang(v string) *DescribeOssObjectDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssObjectDetailRequest) SetId(v int64) *DescribeOssObjectDetailRequest {
	s.Id = &v
	return s
}

type DescribeOssObjectDetailResponseBody struct {
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OssObjectDetail *DescribeOssObjectDetailResponseBodyOssObjectDetail `json:"OssObjectDetail,omitempty" xml:"OssObjectDetail,omitempty" type:"Struct"`
}

func (s DescribeOssObjectDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponseBody) SetRequestId(v string) *DescribeOssObjectDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBody) SetOssObjectDetail(v *DescribeOssObjectDetailResponseBodyOssObjectDetail) *DescribeOssObjectDetailResponseBody {
	s.OssObjectDetail = v
	return s
}

type DescribeOssObjectDetailResponseBodyOssObjectDetail struct {
	CategoryName  *string                                                       `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	Name          *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	BucketName    *string                                                       `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	RiskLevelName *string                                                       `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	RegionId      *string                                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleList      []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetail) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetCategoryName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.CategoryName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetBucketName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.BucketName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetRiskLevelName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetRegionId(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.RegionId = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetRuleList(v []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.RuleList = v
	return s
}

type DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList struct {
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	CategoryName  *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	RiskLevelId   *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	Count         *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	RuleName      *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetRiskLevelName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetCategoryName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetRiskLevelId(v int64) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetCount(v int64) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.Count = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetRuleName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.RuleName = &v
	return s
}

type DescribeOssObjectDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOssObjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOssObjectDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponse) SetHeaders(v map[string]*string) *DescribeOssObjectDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssObjectDetailResponse) SetBody(v *DescribeOssObjectDetailResponseBody) *DescribeOssObjectDetailResponse {
	s.Body = v
	return s
}

type DescribeOssObjectsRequest struct {
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RiskLevelId       *int32  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RuleId            *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage       *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ServiceRegionId   *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	LastScanTimeStart *int64  `json:"LastScanTimeStart,omitempty" xml:"LastScanTimeStart,omitempty"`
	LastScanTimeEnd   *int64  `json:"LastScanTimeEnd,omitempty" xml:"LastScanTimeEnd,omitempty"`
}

func (s DescribeOssObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsRequest) SetLang(v string) *DescribeOssObjectsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetName(v string) *DescribeOssObjectsRequest {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetRiskLevelId(v int32) *DescribeOssObjectsRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetRuleId(v int64) *DescribeOssObjectsRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetInstanceId(v string) *DescribeOssObjectsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetPageSize(v int32) *DescribeOssObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetCurrentPage(v int32) *DescribeOssObjectsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetServiceRegionId(v string) *DescribeOssObjectsRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetLastScanTimeStart(v int64) *DescribeOssObjectsRequest {
	s.LastScanTimeStart = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetLastScanTimeEnd(v int64) *DescribeOssObjectsRequest {
	s.LastScanTimeEnd = &v
	return s
}

type DescribeOssObjectsResponseBody struct {
	CurrentPage *int32                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeOssObjectsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeOssObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsResponseBody) SetCurrentPage(v int32) *DescribeOssObjectsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetRequestId(v string) *DescribeOssObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetPageSize(v int32) *DescribeOssObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetTotalCount(v int32) *DescribeOssObjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetItems(v []*DescribeOssObjectsResponseBodyItems) *DescribeOssObjectsResponseBody {
	s.Items = v
	return s
}

type DescribeOssObjectsResponseBodyItems struct {
	InstanceId     *int64                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CategoryName   *string                                        `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	RiskLevelId    *int64                                         `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RegionId       *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileId         *string                                        `json:"FileId,omitempty" xml:"FileId,omitempty"`
	Size           *int64                                         `json:"Size,omitempty" xml:"Size,omitempty"`
	BucketName     *string                                        `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	RiskLevelName  *string                                        `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	Category       *int64                                         `json:"Category,omitempty" xml:"Category,omitempty"`
	Name           *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleCount      *int32                                         `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	SensitiveCount *int32                                         `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	Id             *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	RuleList       []*DescribeOssObjectsResponseBodyItemsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeOssObjectsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsResponseBodyItems) SetInstanceId(v int64) *DescribeOssObjectsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetCategoryName(v string) *DescribeOssObjectsResponseBodyItems {
	s.CategoryName = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRiskLevelId(v int64) *DescribeOssObjectsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRegionId(v string) *DescribeOssObjectsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetFileId(v string) *DescribeOssObjectsResponseBodyItems {
	s.FileId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetSize(v int64) *DescribeOssObjectsResponseBodyItems {
	s.Size = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetBucketName(v string) *DescribeOssObjectsResponseBodyItems {
	s.BucketName = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRiskLevelName(v string) *DescribeOssObjectsResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetCategory(v int64) *DescribeOssObjectsResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetName(v string) *DescribeOssObjectsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRuleCount(v int32) *DescribeOssObjectsResponseBodyItems {
	s.RuleCount = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetSensitiveCount(v int32) *DescribeOssObjectsResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetId(v string) *DescribeOssObjectsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRuleList(v []*DescribeOssObjectsResponseBodyItemsRuleList) *DescribeOssObjectsResponseBodyItems {
	s.RuleList = v
	return s
}

type DescribeOssObjectsResponseBodyItemsRuleList struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RiskLevelId *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	Count       *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeOssObjectsResponseBodyItemsRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectsResponseBodyItemsRuleList) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) SetName(v string) *DescribeOssObjectsResponseBodyItemsRuleList {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) SetRiskLevelId(v int64) *DescribeOssObjectsResponseBodyItemsRuleList {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) SetCount(v int64) *DescribeOssObjectsResponseBodyItemsRuleList {
	s.Count = &v
	return s
}

type DescribeOssObjectsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOssObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOssObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsResponse) SetHeaders(v map[string]*string) *DescribeOssObjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssObjectsResponse) SetBody(v *DescribeOssObjectsResponseBody) *DescribeOssObjectsResponse {
	s.Body = v
	return s
}

type DescribePackagesRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductId   *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RiskLevelId *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	InstanceId  *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribePackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribePackagesRequest) SetLang(v string) *DescribePackagesRequest {
	s.Lang = &v
	return s
}

func (s *DescribePackagesRequest) SetName(v string) *DescribePackagesRequest {
	s.Name = &v
	return s
}

func (s *DescribePackagesRequest) SetProductId(v int64) *DescribePackagesRequest {
	s.ProductId = &v
	return s
}

func (s *DescribePackagesRequest) SetRiskLevelId(v int64) *DescribePackagesRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribePackagesRequest) SetRuleId(v int64) *DescribePackagesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribePackagesRequest) SetInstanceId(v int64) *DescribePackagesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackagesRequest) SetPageSize(v int32) *DescribePackagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackagesRequest) SetCurrentPage(v int32) *DescribePackagesRequest {
	s.CurrentPage = &v
	return s
}

type DescribePackagesResponseBody struct {
	CurrentPage *int32                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribePackagesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribePackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackagesResponseBody) SetCurrentPage(v int32) *DescribePackagesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribePackagesResponseBody) SetRequestId(v string) *DescribePackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackagesResponseBody) SetPageSize(v int32) *DescribePackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePackagesResponseBody) SetTotalCount(v int32) *DescribePackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePackagesResponseBody) SetItems(v []*DescribePackagesResponseBodyItems) *DescribePackagesResponseBody {
	s.Items = v
	return s
}

type DescribePackagesResponseBodyItems struct {
	CreationTime   *int64  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Sensitive      *bool   `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	Owner          *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	RiskLevelName  *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	InstanceId     *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TotalCount     *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SensitiveCount *int32  `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	RiskLevelId    *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribePackagesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePackagesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribePackagesResponseBodyItems) SetCreationTime(v int64) *DescribePackagesResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetSensitive(v bool) *DescribePackagesResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetOwner(v string) *DescribePackagesResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetRiskLevelName(v string) *DescribePackagesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetInstanceId(v int64) *DescribePackagesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetTotalCount(v int32) *DescribePackagesResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetName(v string) *DescribePackagesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetSensitiveCount(v int32) *DescribePackagesResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetRiskLevelId(v int64) *DescribePackagesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetId(v int64) *DescribePackagesResponseBodyItems {
	s.Id = &v
	return s
}

type DescribePackagesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribePackagesResponse) SetHeaders(v map[string]*string) *DescribePackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribePackagesResponse) SetBody(v *DescribePackagesResponseBody) *DescribePackagesResponse {
	s.Body = v
	return s
}

type DescribeRiskLevelsRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeRiskLevelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskLevelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskLevelsRequest) SetLang(v string) *DescribeRiskLevelsRequest {
	s.Lang = &v
	return s
}

type DescribeRiskLevelsResponseBody struct {
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RiskLevelList []*DescribeRiskLevelsResponseBodyRiskLevelList `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
}

func (s DescribeRiskLevelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskLevelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskLevelsResponseBody) SetRequestId(v string) *DescribeRiskLevelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskLevelsResponseBody) SetRiskLevelList(v []*DescribeRiskLevelsResponseBodyRiskLevelList) *DescribeRiskLevelsResponseBody {
	s.RiskLevelList = v
	return s
}

type DescribeRiskLevelsResponseBodyRiskLevelList struct {
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Alias        *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ReferenceNum *int32  `json:"ReferenceNum,omitempty" xml:"ReferenceNum,omitempty"`
}

func (s DescribeRiskLevelsResponseBodyRiskLevelList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskLevelsResponseBodyRiskLevelList) GoString() string {
	return s.String()
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetName(v string) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.Name = &v
	return s
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetDescription(v string) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.Description = &v
	return s
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetAlias(v string) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.Alias = &v
	return s
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetId(v int64) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.Id = &v
	return s
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetReferenceNum(v int32) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.ReferenceNum = &v
	return s
}

type DescribeRiskLevelsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRiskLevelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRiskLevelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskLevelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskLevelsResponse) SetHeaders(v map[string]*string) *DescribeRiskLevelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskLevelsResponse) SetBody(v *DescribeRiskLevelsResponseBody) *DescribeRiskLevelsResponse {
	s.Body = v
	return s
}

type DescribeRuleCategoryRequest struct {
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ProductId *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s DescribeRuleCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleCategoryRequest) SetLang(v string) *DescribeRuleCategoryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRuleCategoryRequest) SetProductId(v int64) *DescribeRuleCategoryRequest {
	s.ProductId = &v
	return s
}

type DescribeRuleCategoryResponseBody struct {
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleList    []*DescribeRuleCategoryResponseBodyRuleList    `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	ProductList []*DescribeRuleCategoryResponseBodyProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
}

func (s DescribeRuleCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleCategoryResponseBody) SetRequestId(v string) *DescribeRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleCategoryResponseBody) SetRuleList(v []*DescribeRuleCategoryResponseBodyRuleList) *DescribeRuleCategoryResponseBody {
	s.RuleList = v
	return s
}

func (s *DescribeRuleCategoryResponseBody) SetProductList(v []*DescribeRuleCategoryResponseBodyProductList) *DescribeRuleCategoryResponseBody {
	s.ProductList = v
	return s
}

type DescribeRuleCategoryResponseBodyRuleList struct {
	Name      *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	GroupId   *string                                              `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ChildList []*DescribeRuleCategoryResponseBodyRuleListChildList `json:"ChildList,omitempty" xml:"ChildList,omitempty" type:"Repeated"`
}

func (s DescribeRuleCategoryResponseBodyRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleCategoryResponseBodyRuleList) GoString() string {
	return s.String()
}

func (s *DescribeRuleCategoryResponseBodyRuleList) SetName(v string) *DescribeRuleCategoryResponseBodyRuleList {
	s.Name = &v
	return s
}

func (s *DescribeRuleCategoryResponseBodyRuleList) SetGroupId(v string) *DescribeRuleCategoryResponseBodyRuleList {
	s.GroupId = &v
	return s
}

func (s *DescribeRuleCategoryResponseBodyRuleList) SetChildList(v []*DescribeRuleCategoryResponseBodyRuleListChildList) *DescribeRuleCategoryResponseBodyRuleList {
	s.ChildList = v
	return s
}

type DescribeRuleCategoryResponseBodyRuleListChildList struct {
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeRuleCategoryResponseBodyRuleListChildList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleCategoryResponseBodyRuleListChildList) GoString() string {
	return s.String()
}

func (s *DescribeRuleCategoryResponseBodyRuleListChildList) SetParentId(v string) *DescribeRuleCategoryResponseBodyRuleListChildList {
	s.ParentId = &v
	return s
}

func (s *DescribeRuleCategoryResponseBodyRuleListChildList) SetGroupId(v string) *DescribeRuleCategoryResponseBodyRuleListChildList {
	s.GroupId = &v
	return s
}

func (s *DescribeRuleCategoryResponseBodyRuleListChildList) SetName(v string) *DescribeRuleCategoryResponseBodyRuleListChildList {
	s.Name = &v
	return s
}

func (s *DescribeRuleCategoryResponseBodyRuleListChildList) SetId(v int64) *DescribeRuleCategoryResponseBodyRuleListChildList {
	s.Id = &v
	return s
}

type DescribeRuleCategoryResponseBodyProductList struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeRuleCategoryResponseBodyProductList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleCategoryResponseBodyProductList) GoString() string {
	return s.String()
}

func (s *DescribeRuleCategoryResponseBodyProductList) SetLabel(v string) *DescribeRuleCategoryResponseBodyProductList {
	s.Label = &v
	return s
}

func (s *DescribeRuleCategoryResponseBodyProductList) SetValue(v string) *DescribeRuleCategoryResponseBodyProductList {
	s.Value = &v
	return s
}

type DescribeRuleCategoryResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRuleCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleCategoryResponse) SetHeaders(v map[string]*string) *DescribeRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleCategoryResponse) SetBody(v *DescribeRuleCategoryResponseBody) *DescribeRuleCategoryResponse {
	s.Body = v
	return s
}

type DescribeRulesRequest struct {
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage       *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	CustomType        *int32  `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	Category          *int32  `json:"Category,omitempty" xml:"Category,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RiskLevelId       *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RuleType          *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ProductCode       *int32  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductId         *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	WarnLevel         *int32  `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
	ContentCategory   *int32  `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	Status            *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	KeywordCompatible *bool   `json:"KeywordCompatible,omitempty" xml:"KeywordCompatible,omitempty"`
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRulesRequest) SetLang(v string) *DescribeRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRulesRequest) SetPageSize(v int32) *DescribeRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRulesRequest) SetCurrentPage(v int32) *DescribeRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRulesRequest) SetCustomType(v int32) *DescribeRulesRequest {
	s.CustomType = &v
	return s
}

func (s *DescribeRulesRequest) SetCategory(v int32) *DescribeRulesRequest {
	s.Category = &v
	return s
}

func (s *DescribeRulesRequest) SetName(v string) *DescribeRulesRequest {
	s.Name = &v
	return s
}

func (s *DescribeRulesRequest) SetRiskLevelId(v int64) *DescribeRulesRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeRulesRequest) SetRuleType(v int32) *DescribeRulesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRulesRequest) SetProductCode(v int32) *DescribeRulesRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeRulesRequest) SetProductId(v int64) *DescribeRulesRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeRulesRequest) SetWarnLevel(v int32) *DescribeRulesRequest {
	s.WarnLevel = &v
	return s
}

func (s *DescribeRulesRequest) SetContentCategory(v int32) *DescribeRulesRequest {
	s.ContentCategory = &v
	return s
}

func (s *DescribeRulesRequest) SetStatus(v int32) *DescribeRulesRequest {
	s.Status = &v
	return s
}

func (s *DescribeRulesRequest) SetKeywordCompatible(v bool) *DescribeRulesRequest {
	s.KeywordCompatible = &v
	return s
}

func (s *DescribeRulesRequest) SetGroupId(v string) *DescribeRulesRequest {
	s.GroupId = &v
	return s
}

type DescribeRulesResponseBody struct {
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeRulesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBody) SetCurrentPage(v int32) *DescribeRulesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRulesResponseBody) SetRequestId(v string) *DescribeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRulesResponseBody) SetPageSize(v int32) *DescribeRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRulesResponseBody) SetTotalCount(v int32) *DescribeRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRulesResponseBody) SetItems(v []*DescribeRulesResponseBodyItems) *DescribeRulesResponseBody {
	s.Items = v
	return s
}

type DescribeRulesResponseBodyItems struct {
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WarnLevel       *int32  `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
	UserId          *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	StatExpress     *string `json:"StatExpress,omitempty" xml:"StatExpress,omitempty"`
	GmtModified     *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RiskLevelId     *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProductId       *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Target          *string `json:"Target,omitempty" xml:"Target,omitempty"`
	LoginName       *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	CategoryName    *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ContentCategory *string `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	HitTotalCount   *int32  `json:"HitTotalCount,omitempty" xml:"HitTotalCount,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	CustomType      *int32  `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	RiskLevelName   *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	GmtCreate       *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Category        *int32  `json:"Category,omitempty" xml:"Category,omitempty"`
	MajorKey        *string `json:"MajorKey,omitempty" xml:"MajorKey,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ProductCode     *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeRulesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyItems) SetDisplayName(v string) *DescribeRulesResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetStatus(v int32) *DescribeRulesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetWarnLevel(v int32) *DescribeRulesResponseBodyItems {
	s.WarnLevel = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetUserId(v int64) *DescribeRulesResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetStatExpress(v string) *DescribeRulesResponseBodyItems {
	s.StatExpress = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetGmtModified(v int64) *DescribeRulesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetRiskLevelId(v int64) *DescribeRulesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetDescription(v string) *DescribeRulesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetProductId(v int64) *DescribeRulesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetName(v string) *DescribeRulesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetContent(v string) *DescribeRulesResponseBodyItems {
	s.Content = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetTarget(v string) *DescribeRulesResponseBodyItems {
	s.Target = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetLoginName(v string) *DescribeRulesResponseBodyItems {
	s.LoginName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetCategoryName(v string) *DescribeRulesResponseBodyItems {
	s.CategoryName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetContentCategory(v string) *DescribeRulesResponseBodyItems {
	s.ContentCategory = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetHitTotalCount(v int32) *DescribeRulesResponseBodyItems {
	s.HitTotalCount = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetGroupId(v string) *DescribeRulesResponseBodyItems {
	s.GroupId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetCustomType(v int32) *DescribeRulesResponseBodyItems {
	s.CustomType = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetRiskLevelName(v string) *DescribeRulesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetGmtCreate(v int64) *DescribeRulesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetCategory(v int32) *DescribeRulesResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetMajorKey(v string) *DescribeRulesResponseBodyItems {
	s.MajorKey = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetId(v int64) *DescribeRulesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetProductCode(v string) *DescribeRulesResponseBodyItems {
	s.ProductCode = &v
	return s
}

type DescribeRulesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponse) SetHeaders(v map[string]*string) *DescribeRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRulesResponse) SetBody(v *DescribeRulesResponseBody) *DescribeRulesResponse {
	s.Body = v
	return s
}

type DescribeTablesRequest struct {
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductId           *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RiskLevelId         *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RuleId              *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	InstanceId          *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName        *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RuleName            *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SensLevelName       *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	PackageId           *int64  `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage         *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	LastScanTimeStart   *int64  `json:"LastScanTimeStart,omitempty" xml:"LastScanTimeStart,omitempty"`
	LastScanTimeEnd     *int64  `json:"LastScanTimeEnd,omitempty" xml:"LastScanTimeEnd,omitempty"`
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime             *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) SetLang(v string) *DescribeTablesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTablesRequest) SetName(v string) *DescribeTablesRequest {
	s.Name = &v
	return s
}

func (s *DescribeTablesRequest) SetProductId(v int64) *DescribeTablesRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeTablesRequest) SetProductCode(v string) *DescribeTablesRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeTablesRequest) SetRiskLevelId(v int64) *DescribeTablesRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeTablesRequest) SetRuleId(v int64) *DescribeTablesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeTablesRequest) SetInstanceId(v int64) *DescribeTablesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTablesRequest) SetInstanceName(v string) *DescribeTablesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeTablesRequest) SetRuleName(v string) *DescribeTablesRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeTablesRequest) SetSensLevelName(v string) *DescribeTablesRequest {
	s.SensLevelName = &v
	return s
}

func (s *DescribeTablesRequest) SetPackageId(v int64) *DescribeTablesRequest {
	s.PackageId = &v
	return s
}

func (s *DescribeTablesRequest) SetPageSize(v int32) *DescribeTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesRequest) SetCurrentPage(v int32) *DescribeTablesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTablesRequest) SetLastScanTimeStart(v int64) *DescribeTablesRequest {
	s.LastScanTimeStart = &v
	return s
}

func (s *DescribeTablesRequest) SetLastScanTimeEnd(v int64) *DescribeTablesRequest {
	s.LastScanTimeEnd = &v
	return s
}

func (s *DescribeTablesRequest) SetStartTime(v int64) *DescribeTablesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTablesRequest) SetEndTime(v int64) *DescribeTablesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTablesRequest) SetInstanceDescription(v string) *DescribeTablesRequest {
	s.InstanceDescription = &v
	return s
}

type DescribeTablesResponseBody struct {
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items       []*DescribeTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBody) SetCurrentPage(v int32) *DescribeTablesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablesResponseBody) SetPageSize(v int32) *DescribeTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesResponseBody) SetTotalCount(v int32) *DescribeTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTablesResponseBody) SetItems(v []*DescribeTablesResponseBodyItems) *DescribeTablesResponseBody {
	s.Items = v
	return s
}

type DescribeTablesResponseBodyItems struct {
	CreationTime        *int64                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Owner               *string                                    `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SensitiveRatio      *string                                    `json:"SensitiveRatio,omitempty" xml:"SensitiveRatio,omitempty"`
	TenantName          *string                                    `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	RiskLevelId         *int64                                     `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	InstanceName        *string                                    `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ProductId           *string                                    `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Name                *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	InstanceId          *int64                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TotalCount          *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	InstanceDescription *string                                    `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	Sensitive           *bool                                      `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	RiskLevelName       *string                                    `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	SensitiveCount      *int32                                     `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	Id                  *int64                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	ProductCode         *string                                    `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RuleList            []*DescribeTablesResponseBodyItemsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeTablesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItems) SetCreationTime(v int64) *DescribeTablesResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetOwner(v string) *DescribeTablesResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetSensitiveRatio(v string) *DescribeTablesResponseBodyItems {
	s.SensitiveRatio = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetTenantName(v string) *DescribeTablesResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetRiskLevelId(v int64) *DescribeTablesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetInstanceName(v string) *DescribeTablesResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetProductId(v string) *DescribeTablesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetName(v string) *DescribeTablesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetInstanceId(v int64) *DescribeTablesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetTotalCount(v int32) *DescribeTablesResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetInstanceDescription(v string) *DescribeTablesResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetSensitive(v bool) *DescribeTablesResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetRiskLevelName(v string) *DescribeTablesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetSensitiveCount(v int32) *DescribeTablesResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetId(v int64) *DescribeTablesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetProductCode(v string) *DescribeTablesResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetRuleList(v []*DescribeTablesResponseBodyItemsRuleList) *DescribeTablesResponseBodyItems {
	s.RuleList = v
	return s
}

type DescribeTablesResponseBodyItemsRuleList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RiskLevelId *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	Count       *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeTablesResponseBodyItemsRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyItemsRuleList) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItemsRuleList) SetDisplayName(v string) *DescribeTablesResponseBodyItemsRuleList {
	s.DisplayName = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsRuleList) SetName(v string) *DescribeTablesResponseBodyItemsRuleList {
	s.Name = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsRuleList) SetRiskLevelId(v int64) *DescribeTablesResponseBodyItemsRuleList {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsRuleList) SetCount(v int64) *DescribeTablesResponseBodyItemsRuleList {
	s.Count = &v
	return s
}

type DescribeTablesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponse) SetHeaders(v map[string]*string) *DescribeTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTablesResponse) SetBody(v *DescribeTablesResponseBody) *DescribeTablesResponse {
	s.Body = v
	return s
}

type DescribeUserCountRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeUserCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserCountRequest) SetLang(v string) *DescribeUserCountRequest {
	s.Lang = &v
	return s
}

type DescribeUserCountResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserCount *DescribeUserCountResponseBodyUserCount `json:"UserCount,omitempty" xml:"UserCount,omitempty" type:"Struct"`
}

func (s DescribeUserCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserCountResponseBody) SetRequestId(v string) *DescribeUserCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserCountResponseBody) SetUserCount(v *DescribeUserCountResponseBodyUserCount) *DescribeUserCountResponseBody {
	s.UserCount = v
	return s
}

type DescribeUserCountResponseBodyUserCount struct {
	SaleInstanceId     *string `json:"SaleInstanceId,omitempty" xml:"SaleInstanceId,omitempty"`
	Authed             *bool   `json:"Authed,omitempty" xml:"Authed,omitempty"`
	Buyed              *bool   `json:"Buyed,omitempty" xml:"Buyed,omitempty"`
	SaleInstanceStatus *int32  `json:"SaleInstanceStatus,omitempty" xml:"SaleInstanceStatus,omitempty"`
	WarnAddCount       *int64  `json:"WarnAddCount,omitempty" xml:"WarnAddCount,omitempty"`
	WarnTotalCount     *int64  `json:"WarnTotalCount,omitempty" xml:"WarnTotalCount,omitempty"`
	UnAuthCount        *int32  `json:"UnAuthCount,omitempty" xml:"UnAuthCount,omitempty"`
	TotalAuthCount     *int32  `json:"TotalAuthCount,omitempty" xml:"TotalAuthCount,omitempty"`
	AuthCount          *int32  `json:"AuthCount,omitempty" xml:"AuthCount,omitempty"`
}

func (s DescribeUserCountResponseBodyUserCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCountResponseBodyUserCount) GoString() string {
	return s.String()
}

func (s *DescribeUserCountResponseBodyUserCount) SetSaleInstanceId(v string) *DescribeUserCountResponseBodyUserCount {
	s.SaleInstanceId = &v
	return s
}

func (s *DescribeUserCountResponseBodyUserCount) SetAuthed(v bool) *DescribeUserCountResponseBodyUserCount {
	s.Authed = &v
	return s
}

func (s *DescribeUserCountResponseBodyUserCount) SetBuyed(v bool) *DescribeUserCountResponseBodyUserCount {
	s.Buyed = &v
	return s
}

func (s *DescribeUserCountResponseBodyUserCount) SetSaleInstanceStatus(v int32) *DescribeUserCountResponseBodyUserCount {
	s.SaleInstanceStatus = &v
	return s
}

func (s *DescribeUserCountResponseBodyUserCount) SetWarnAddCount(v int64) *DescribeUserCountResponseBodyUserCount {
	s.WarnAddCount = &v
	return s
}

func (s *DescribeUserCountResponseBodyUserCount) SetWarnTotalCount(v int64) *DescribeUserCountResponseBodyUserCount {
	s.WarnTotalCount = &v
	return s
}

func (s *DescribeUserCountResponseBodyUserCount) SetUnAuthCount(v int32) *DescribeUserCountResponseBodyUserCount {
	s.UnAuthCount = &v
	return s
}

func (s *DescribeUserCountResponseBodyUserCount) SetTotalAuthCount(v int32) *DescribeUserCountResponseBodyUserCount {
	s.TotalAuthCount = &v
	return s
}

func (s *DescribeUserCountResponseBodyUserCount) SetAuthCount(v int32) *DescribeUserCountResponseBodyUserCount {
	s.AuthCount = &v
	return s
}

type DescribeUserCountResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserCountResponse) SetHeaders(v map[string]*string) *DescribeUserCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserCountResponse) SetBody(v *DescribeUserCountResponseBody) *DescribeUserCountResponse {
	s.Body = v
	return s
}

type DescribeUserStatusRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusRequest) SetLang(v string) *DescribeUserStatusRequest {
	s.Lang = &v
	return s
}

type DescribeUserStatusResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserStatus *DescribeUserStatusResponseBodyUserStatus `json:"UserStatus,omitempty" xml:"UserStatus,omitempty" type:"Struct"`
}

func (s DescribeUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponseBody) SetRequestId(v string) *DescribeUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserStatusResponseBody) SetUserStatus(v *DescribeUserStatusResponseBodyUserStatus) *DescribeUserStatusResponseBody {
	s.UserStatus = v
	return s
}

type DescribeUserStatusResponseBodyUserStatus struct {
	UseOssSize           *int64  `json:"UseOssSize,omitempty" xml:"UseOssSize,omitempty"`
	UseInstanceNum       *int32  `json:"UseInstanceNum,omitempty" xml:"UseInstanceNum,omitempty"`
	InstanceNum          *int32  `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	ChargeType           *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	IncSensitiveTables   *int64  `json:"IncSensitiveTables,omitempty" xml:"IncSensitiveTables,omitempty"`
	AuditUpgradeStatus   *bool   `json:"AuditUpgradeStatus,omitempty" xml:"AuditUpgradeStatus,omitempty"`
	DataMaskTasks        *int64  `json:"DataMaskTasks,omitempty" xml:"DataMaskTasks,omitempty"`
	Authed               *bool   `json:"Authed,omitempty" xml:"Authed,omitempty"`
	LabStatus            *int32  `json:"LabStatus,omitempty" xml:"LabStatus,omitempty"`
	OdpsSet              *bool   `json:"OdpsSet,omitempty" xml:"OdpsSet,omitempty"`
	OssBucketSet         *bool   `json:"OssBucketSet,omitempty" xml:"OssBucketSet,omitempty"`
	SensitiveTable       *int64  `json:"SensitiveTable,omitempty" xml:"SensitiveTable,omitempty"`
	RemainDays           *int32  `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	TotalDataMaskColumns *int64  `json:"TotalDataMaskColumns,omitempty" xml:"TotalDataMaskColumns,omitempty"`
	DatamaskColumns      *int64  `json:"DatamaskColumns,omitempty" xml:"DatamaskColumns,omitempty"`
	AuthProductList      *string `json:"AuthProductList,omitempty" xml:"AuthProductList,omitempty"`
	Trail                *bool   `json:"Trail,omitempty" xml:"Trail,omitempty"`
	DivulgeCount         *int64  `json:"DivulgeCount,omitempty" xml:"DivulgeCount,omitempty"`
	RdsSet               *bool   `json:"RdsSet,omitempty" xml:"RdsSet,omitempty"`
	DbAuditStatus        *int32  `json:"DbAuditStatus,omitempty" xml:"DbAuditStatus,omitempty"`
	RenewStatus          *bool   `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	SensitiveObjects     *int64  `json:"SensitiveObjects,omitempty" xml:"SensitiveObjects,omitempty"`
	DataMaskColumns      *int64  `json:"DataMaskColumns,omitempty" xml:"DataMaskColumns,omitempty"`
	DlpTotalCount        *int64  `json:"DlpTotalCount,omitempty" xml:"DlpTotalCount,omitempty"`
	SensitiveObject      *int64  `json:"SensitiveObject,omitempty" xml:"SensitiveObject,omitempty"`
	DlpCount             *int64  `json:"DlpCount,omitempty" xml:"DlpCount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccessKeyId          *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	OssSize              *int64  `json:"OssSize,omitempty" xml:"OssSize,omitempty"`
	SensitiveTotalTop    *string `json:"SensitiveTotalTop,omitempty" xml:"SensitiveTotalTop,omitempty"`
	SensitiveTables      *int64  `json:"SensitiveTables,omitempty" xml:"SensitiveTables,omitempty"`
	SensitiveTop         *string `json:"SensitiveTop,omitempty" xml:"SensitiveTop,omitempty"`
	Buyed                *bool   `json:"Buyed,omitempty" xml:"Buyed,omitempty"`
	IncSensitiveObjects  *int64  `json:"IncSensitiveObjects,omitempty" xml:"IncSensitiveObjects,omitempty"`
	InstanceStatus       *int32  `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	AssetScanned         *bool   `json:"AssetScanned,omitempty" xml:"AssetScanned,omitempty"`
	AlarmCount           *int64  `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	DisplayTime          *string `json:"DisplayTime,omitempty" xml:"DisplayTime,omitempty"`
}

func (s DescribeUserStatusResponseBodyUserStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusResponseBodyUserStatus) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetUseOssSize(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.UseOssSize = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetUseInstanceNum(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.UseInstanceNum = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetInstanceNum(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.InstanceNum = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetChargeType(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.ChargeType = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetIncSensitiveTables(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.IncSensitiveTables = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAuditUpgradeStatus(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.AuditUpgradeStatus = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetDataMaskTasks(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.DataMaskTasks = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAuthed(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.Authed = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetLabStatus(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.LabStatus = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetOdpsSet(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.OdpsSet = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetOssBucketSet(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.OssBucketSet = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetSensitiveTable(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.SensitiveTable = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetRemainDays(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.RemainDays = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetTotalDataMaskColumns(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.TotalDataMaskColumns = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetDatamaskColumns(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.DatamaskColumns = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAuthProductList(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.AuthProductList = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetTrail(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.Trail = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetDivulgeCount(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.DivulgeCount = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetRdsSet(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.RdsSet = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetDbAuditStatus(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.DbAuditStatus = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetRenewStatus(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.RenewStatus = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetSensitiveObjects(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.SensitiveObjects = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetDataMaskColumns(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.DataMaskColumns = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetDlpTotalCount(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.DlpTotalCount = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetSensitiveObject(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.SensitiveObject = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetDlpCount(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.DlpCount = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetInstanceId(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAccessKeyId(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetOssSize(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.OssSize = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetSensitiveTotalTop(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.SensitiveTotalTop = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetSensitiveTables(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.SensitiveTables = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetSensitiveTop(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.SensitiveTop = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetBuyed(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.Buyed = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetIncSensitiveObjects(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.IncSensitiveObjects = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetInstanceStatus(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAssetScanned(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.AssetScanned = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAlarmCount(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.AlarmCount = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetDisplayTime(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.DisplayTime = &v
	return s
}

type DescribeUserStatusResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponse) SetHeaders(v map[string]*string) *DescribeUserStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserStatusResponse) SetBody(v *DescribeUserStatusResponseBody) *DescribeUserStatusResponse {
	s.Body = v
	return s
}

type DisableUserConfigRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DisableUserConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableUserConfigRequest) GoString() string {
	return s.String()
}

func (s *DisableUserConfigRequest) SetLang(v string) *DisableUserConfigRequest {
	s.Lang = &v
	return s
}

func (s *DisableUserConfigRequest) SetCode(v string) *DisableUserConfigRequest {
	s.Code = &v
	return s
}

type DisableUserConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableUserConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DisableUserConfigResponseBody) SetRequestId(v string) *DisableUserConfigResponseBody {
	s.RequestId = &v
	return s
}

type DisableUserConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableUserConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableUserConfigResponse) GoString() string {
	return s.String()
}

func (s *DisableUserConfigResponse) SetHeaders(v map[string]*string) *DisableUserConfigResponse {
	s.Headers = v
	return s
}

func (s *DisableUserConfigResponse) SetBody(v *DisableUserConfigResponseBody) *DisableUserConfigResponse {
	s.Body = v
	return s
}

type ExecDatamaskRequest struct {
	TemplateId *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ExecDatamaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecDatamaskRequest) GoString() string {
	return s.String()
}

func (s *ExecDatamaskRequest) SetTemplateId(v int64) *ExecDatamaskRequest {
	s.TemplateId = &v
	return s
}

func (s *ExecDatamaskRequest) SetData(v string) *ExecDatamaskRequest {
	s.Data = &v
	return s
}

type ExecDatamaskResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecDatamaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecDatamaskResponseBody) GoString() string {
	return s.String()
}

func (s *ExecDatamaskResponseBody) SetData(v string) *ExecDatamaskResponseBody {
	s.Data = &v
	return s
}

func (s *ExecDatamaskResponseBody) SetRequestId(v string) *ExecDatamaskResponseBody {
	s.RequestId = &v
	return s
}

type ExecDatamaskResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecDatamaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecDatamaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecDatamaskResponse) GoString() string {
	return s.String()
}

func (s *ExecDatamaskResponse) SetHeaders(v map[string]*string) *ExecDatamaskResponse {
	s.Headers = v
	return s
}

func (s *ExecDatamaskResponse) SetBody(v *ExecDatamaskResponseBody) *ExecDatamaskResponse {
	s.Body = v
	return s
}

type ManualTriggerMaskingProcessRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ManualTriggerMaskingProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s ManualTriggerMaskingProcessRequest) GoString() string {
	return s.String()
}

func (s *ManualTriggerMaskingProcessRequest) SetLang(v string) *ManualTriggerMaskingProcessRequest {
	s.Lang = &v
	return s
}

func (s *ManualTriggerMaskingProcessRequest) SetId(v int64) *ManualTriggerMaskingProcessRequest {
	s.Id = &v
	return s
}

type ManualTriggerMaskingProcessResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManualTriggerMaskingProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManualTriggerMaskingProcessResponseBody) GoString() string {
	return s.String()
}

func (s *ManualTriggerMaskingProcessResponseBody) SetRequestId(v string) *ManualTriggerMaskingProcessResponseBody {
	s.RequestId = &v
	return s
}

type ManualTriggerMaskingProcessResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ManualTriggerMaskingProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ManualTriggerMaskingProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s ManualTriggerMaskingProcessResponse) GoString() string {
	return s.String()
}

func (s *ManualTriggerMaskingProcessResponse) SetHeaders(v map[string]*string) *ManualTriggerMaskingProcessResponse {
	s.Headers = v
	return s
}

func (s *ManualTriggerMaskingProcessResponse) SetBody(v *ManualTriggerMaskingProcessResponseBody) *ManualTriggerMaskingProcessResponse {
	s.Body = v
	return s
}

type ModifyDataLimitRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ResourceType    *int32  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ModifyPassword  *bool   `json:"ModifyPassword,omitempty" xml:"ModifyPassword,omitempty"`
	AuditStatus     *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	LogStoreDay     *int32  `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	EngineType      *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	AutoScan        *int32  `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
}

func (s ModifyDataLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataLimitRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataLimitRequest) SetLang(v string) *ModifyDataLimitRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDataLimitRequest) SetId(v int64) *ModifyDataLimitRequest {
	s.Id = &v
	return s
}

func (s *ModifyDataLimitRequest) SetResourceType(v int32) *ModifyDataLimitRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyDataLimitRequest) SetServiceRegionId(v string) *ModifyDataLimitRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *ModifyDataLimitRequest) SetUserName(v string) *ModifyDataLimitRequest {
	s.UserName = &v
	return s
}

func (s *ModifyDataLimitRequest) SetPassword(v string) *ModifyDataLimitRequest {
	s.Password = &v
	return s
}

func (s *ModifyDataLimitRequest) SetModifyPassword(v bool) *ModifyDataLimitRequest {
	s.ModifyPassword = &v
	return s
}

func (s *ModifyDataLimitRequest) SetAuditStatus(v int32) *ModifyDataLimitRequest {
	s.AuditStatus = &v
	return s
}

func (s *ModifyDataLimitRequest) SetLogStoreDay(v int32) *ModifyDataLimitRequest {
	s.LogStoreDay = &v
	return s
}

func (s *ModifyDataLimitRequest) SetEngineType(v string) *ModifyDataLimitRequest {
	s.EngineType = &v
	return s
}

func (s *ModifyDataLimitRequest) SetPort(v int32) *ModifyDataLimitRequest {
	s.Port = &v
	return s
}

func (s *ModifyDataLimitRequest) SetAutoScan(v int32) *ModifyDataLimitRequest {
	s.AutoScan = &v
	return s
}

type ModifyDataLimitResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataLimitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataLimitResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataLimitResponseBody) SetRequestId(v string) *ModifyDataLimitResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDataLimitResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDataLimitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDataLimitResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataLimitResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataLimitResponse) SetHeaders(v map[string]*string) *ModifyDataLimitResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataLimitResponse) SetBody(v *ModifyDataLimitResponseBody) *ModifyDataLimitResponse {
	s.Body = v
	return s
}

type ModifyDefaultLevelRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DefaultId    *int64  `json:"DefaultId,omitempty" xml:"DefaultId,omitempty"`
	SensitiveIds *string `json:"SensitiveIds,omitempty" xml:"SensitiveIds,omitempty"`
}

func (s ModifyDefaultLevelRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefaultLevelRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefaultLevelRequest) SetLang(v string) *ModifyDefaultLevelRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDefaultLevelRequest) SetDefaultId(v int64) *ModifyDefaultLevelRequest {
	s.DefaultId = &v
	return s
}

func (s *ModifyDefaultLevelRequest) SetSensitiveIds(v string) *ModifyDefaultLevelRequest {
	s.SensitiveIds = &v
	return s
}

type ModifyDefaultLevelResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefaultLevelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefaultLevelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefaultLevelResponseBody) SetRequestId(v string) *ModifyDefaultLevelResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefaultLevelResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDefaultLevelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDefaultLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefaultLevelResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefaultLevelResponse) SetHeaders(v map[string]*string) *ModifyDefaultLevelResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefaultLevelResponse) SetBody(v *ModifyDefaultLevelResponseBody) *ModifyDefaultLevelResponse {
	s.Body = v
	return s
}

type ModifyEventStatusRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Backed     *bool   `json:"Backed,omitempty" xml:"Backed,omitempty"`
	DealReason *string `json:"DealReason,omitempty" xml:"DealReason,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyEventStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEventStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyEventStatusRequest) SetLang(v string) *ModifyEventStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyEventStatusRequest) SetId(v int64) *ModifyEventStatusRequest {
	s.Id = &v
	return s
}

func (s *ModifyEventStatusRequest) SetBacked(v bool) *ModifyEventStatusRequest {
	s.Backed = &v
	return s
}

func (s *ModifyEventStatusRequest) SetDealReason(v string) *ModifyEventStatusRequest {
	s.DealReason = &v
	return s
}

func (s *ModifyEventStatusRequest) SetStatus(v int32) *ModifyEventStatusRequest {
	s.Status = &v
	return s
}

type ModifyEventStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEventStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyEventStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEventStatusResponseBody) SetRequestId(v string) *ModifyEventStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyEventStatusResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyEventStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyEventStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyEventStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyEventStatusResponse) SetHeaders(v map[string]*string) *ModifyEventStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyEventStatusResponse) SetBody(v *ModifyEventStatusResponseBody) *ModifyEventStatusResponse {
	s.Body = v
	return s
}

type ModifyEventTypeStatusRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SubTypeIds *string `json:"SubTypeIds,omitempty" xml:"SubTypeIds,omitempty"`
}

func (s ModifyEventTypeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEventTypeStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyEventTypeStatusRequest) SetLang(v string) *ModifyEventTypeStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyEventTypeStatusRequest) SetSubTypeIds(v string) *ModifyEventTypeStatusRequest {
	s.SubTypeIds = &v
	return s
}

type ModifyEventTypeStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEventTypeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyEventTypeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEventTypeStatusResponseBody) SetRequestId(v string) *ModifyEventTypeStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyEventTypeStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyEventTypeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyEventTypeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyEventTypeStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyEventTypeStatusResponse) SetHeaders(v map[string]*string) *ModifyEventTypeStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyEventTypeStatusResponse) SetBody(v *ModifyEventTypeStatusResponseBody) *ModifyEventTypeStatusResponse {
	s.Body = v
	return s
}

type ModifyRuleRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	CustomType      *int32  `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Category        *int32  `json:"Category,omitempty" xml:"Category,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	StatExpress     *string `json:"StatExpress,omitempty" xml:"StatExpress,omitempty"`
	RiskLevelId     *int64  `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	RuleType        *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ProductCode     *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductId       *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	WarnLevel       *int32  `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ContentCategory *int32  `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	Target          *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyRuleRequest) SetLang(v string) *ModifyRuleRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRuleRequest) SetId(v int64) *ModifyRuleRequest {
	s.Id = &v
	return s
}

func (s *ModifyRuleRequest) SetCustomType(v int32) *ModifyRuleRequest {
	s.CustomType = &v
	return s
}

func (s *ModifyRuleRequest) SetName(v string) *ModifyRuleRequest {
	s.Name = &v
	return s
}

func (s *ModifyRuleRequest) SetCategory(v int32) *ModifyRuleRequest {
	s.Category = &v
	return s
}

func (s *ModifyRuleRequest) SetContent(v string) *ModifyRuleRequest {
	s.Content = &v
	return s
}

func (s *ModifyRuleRequest) SetStatExpress(v string) *ModifyRuleRequest {
	s.StatExpress = &v
	return s
}

func (s *ModifyRuleRequest) SetRiskLevelId(v int64) *ModifyRuleRequest {
	s.RiskLevelId = &v
	return s
}

func (s *ModifyRuleRequest) SetRuleType(v int32) *ModifyRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyRuleRequest) SetProductCode(v string) *ModifyRuleRequest {
	s.ProductCode = &v
	return s
}

func (s *ModifyRuleRequest) SetProductId(v int64) *ModifyRuleRequest {
	s.ProductId = &v
	return s
}

func (s *ModifyRuleRequest) SetWarnLevel(v int32) *ModifyRuleRequest {
	s.WarnLevel = &v
	return s
}

func (s *ModifyRuleRequest) SetDescription(v string) *ModifyRuleRequest {
	s.Description = &v
	return s
}

func (s *ModifyRuleRequest) SetContentCategory(v int32) *ModifyRuleRequest {
	s.ContentCategory = &v
	return s
}

func (s *ModifyRuleRequest) SetTarget(v string) *ModifyRuleRequest {
	s.Target = &v
	return s
}

type ModifyRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRuleResponseBody) SetRequestId(v string) *ModifyRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyRuleResponse) SetHeaders(v map[string]*string) *ModifyRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyRuleResponse) SetBody(v *ModifyRuleResponseBody) *ModifyRuleResponse {
	s.Body = v
	return s
}

type ModifyRuleStatusRequest struct {
	Lang   *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Ids    *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyRuleStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyRuleStatusRequest) SetLang(v string) *ModifyRuleStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetId(v int64) *ModifyRuleStatusRequest {
	s.Id = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetIds(v string) *ModifyRuleStatusRequest {
	s.Ids = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetStatus(v int32) *ModifyRuleStatusRequest {
	s.Status = &v
	return s
}

type ModifyRuleStatusResponseBody struct {
	FailedIds *string `json:"FailedIds,omitempty" xml:"FailedIds,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRuleStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRuleStatusResponseBody) SetFailedIds(v string) *ModifyRuleStatusResponseBody {
	s.FailedIds = &v
	return s
}

func (s *ModifyRuleStatusResponseBody) SetRequestId(v string) *ModifyRuleStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRuleStatusResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRuleStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyRuleStatusResponse) SetHeaders(v map[string]*string) *ModifyRuleStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyRuleStatusResponse) SetBody(v *ModifyRuleStatusResponseBody) *ModifyRuleStatusResponse {
	s.Body = v
	return s
}

type StopMaskingProcessRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	FeatureType *int32  `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
}

func (s StopMaskingProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s StopMaskingProcessRequest) GoString() string {
	return s.String()
}

func (s *StopMaskingProcessRequest) SetLang(v string) *StopMaskingProcessRequest {
	s.Lang = &v
	return s
}

func (s *StopMaskingProcessRequest) SetId(v int64) *StopMaskingProcessRequest {
	s.Id = &v
	return s
}

func (s *StopMaskingProcessRequest) SetFeatureType(v int32) *StopMaskingProcessRequest {
	s.FeatureType = &v
	return s
}

type StopMaskingProcessResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopMaskingProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopMaskingProcessResponseBody) GoString() string {
	return s.String()
}

func (s *StopMaskingProcessResponseBody) SetRequestId(v string) *StopMaskingProcessResponseBody {
	s.RequestId = &v
	return s
}

type StopMaskingProcessResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopMaskingProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopMaskingProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s StopMaskingProcessResponse) GoString() string {
	return s.String()
}

func (s *StopMaskingProcessResponse) SetHeaders(v map[string]*string) *StopMaskingProcessResponse {
	s.Headers = v
	return s
}

func (s *StopMaskingProcessResponse) SetBody(v *StopMaskingProcessResponseBody) *StopMaskingProcessResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hongkong": tea.String("sddp-api.cn-hongkong.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sddp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteDataLimitWithOptions(runtime *util.RuntimeOptions) (_result *BatchDeleteDataLimitResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &BatchDeleteDataLimitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchDeleteDataLimit"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteDataLimit() (_result *BatchDeleteDataLimitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteDataLimitResponse{}
	_body, _err := client.BatchDeleteDataLimitWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConfigWithOptions(request *CreateConfigRequest, runtime *util.RuntimeOptions) (_result *CreateConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateConfig"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConfig(request *CreateConfigRequest) (_result *CreateConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConfigResponse{}
	_body, _err := client.CreateConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataLimitWithOptions(request *CreateDataLimitRequest, runtime *util.RuntimeOptions) (_result *CreateDataLimitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDataLimitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDataLimit"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataLimit(request *CreateDataLimitRequest) (_result *CreateDataLimitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataLimitResponse{}
	_body, _err := client.CreateDataLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRule"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScanTaskWithOptions(request *CreateScanTaskRequest, runtime *util.RuntimeOptions) (_result *CreateScanTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateScanTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateScanTask"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScanTask(request *CreateScanTaskRequest) (_result *CreateScanTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScanTaskResponse{}
	_body, _err := client.CreateScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataLimitWithOptions(request *DeleteDataLimitRequest, runtime *util.RuntimeOptions) (_result *DeleteDataLimitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDataLimitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDataLimit"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataLimit(request *DeleteDataLimitRequest) (_result *DeleteDataLimitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataLimitResponse{}
	_body, _err := client.DeleteDataLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRule"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountDetailWithOptions(request *DescribeAccountDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccountDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccountDetail"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccountDetail(request *DescribeAccountDetailRequest) (_result *DescribeAccountDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountDetailResponse{}
	_body, _err := client.DescribeAccountDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccounts"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (_result *DescribeAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DescribeAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeColumnsWithOptions(request *DescribeColumnsRequest, runtime *util.RuntimeOptions) (_result *DescribeColumnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeColumns"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeColumns(request *DescribeColumnsRequest) (_result *DescribeColumnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.DescribeColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigsWithOptions(request *DescribeConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConfigs"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigs(request *DescribeConfigsRequest) (_result *DescribeConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConfigsResponse{}
	_body, _err := client.DescribeConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataAssetsWithOptions(request *DescribeDataAssetsRequest, runtime *util.RuntimeOptions) (_result *DescribeDataAssetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataAssetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataAssets"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataAssets(request *DescribeDataAssetsRequest) (_result *DescribeDataAssetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataAssetsResponse{}
	_body, _err := client.DescribeDataAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataLimitDetailWithOptions(request *DescribeDataLimitDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDataLimitDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataLimitDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataLimitDetail"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataLimitDetail(request *DescribeDataLimitDetailRequest) (_result *DescribeDataLimitDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataLimitDetailResponse{}
	_body, _err := client.DescribeDataLimitDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataLimitsWithOptions(request *DescribeDataLimitsRequest, runtime *util.RuntimeOptions) (_result *DescribeDataLimitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataLimitsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataLimits"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataLimits(request *DescribeDataLimitsRequest) (_result *DescribeDataLimitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataLimitsResponse{}
	_body, _err := client.DescribeDataLimitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataLimitSetWithOptions(request *DescribeDataLimitSetRequest, runtime *util.RuntimeOptions) (_result *DescribeDataLimitSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataLimitSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataLimitSet"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataLimitSet(request *DescribeDataLimitSetRequest) (_result *DescribeDataLimitSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataLimitSetResponse{}
	_body, _err := client.DescribeDataLimitSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataLimitsInstanceDimWithOptions(request *DescribeDataLimitsInstanceDimRequest, runtime *util.RuntimeOptions) (_result *DescribeDataLimitsInstanceDimResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataLimitsInstanceDimResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataLimitsInstanceDim"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataLimitsInstanceDim(request *DescribeDataLimitsInstanceDimRequest) (_result *DescribeDataLimitsInstanceDimResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataLimitsInstanceDimResponse{}
	_body, _err := client.DescribeDataLimitsInstanceDimWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataMaskingRunHistoryWithOptions(request *DescribeDataMaskingRunHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeDataMaskingRunHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataMaskingRunHistoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataMaskingRunHistory"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataMaskingRunHistory(request *DescribeDataMaskingRunHistoryRequest) (_result *DescribeDataMaskingRunHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataMaskingRunHistoryResponse{}
	_body, _err := client.DescribeDataMaskingRunHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataMaskingTasksWithOptions(request *DescribeDataMaskingTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeDataMaskingTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataMaskingTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataMaskingTasks"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataMaskingTasks(request *DescribeDataMaskingTasksRequest) (_result *DescribeDataMaskingTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataMaskingTasksResponse{}
	_body, _err := client.DescribeDataMaskingTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataTotalCountWithOptions(request *DescribeDataTotalCountRequest, runtime *util.RuntimeOptions) (_result *DescribeDataTotalCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDataTotalCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDataTotalCount"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataTotalCount(request *DescribeDataTotalCountRequest) (_result *DescribeDataTotalCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataTotalCountResponse{}
	_body, _err := client.DescribeDataTotalCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventDetailWithOptions(request *DescribeEventDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeEventDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEventDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEventDetail"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEventDetail(request *DescribeEventDetailRequest) (_result *DescribeEventDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventDetailResponse{}
	_body, _err := client.DescribeEventDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventsWithOptions(request *DescribeEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEvents"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEvents(request *DescribeEventsRequest) (_result *DescribeEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventsResponse{}
	_body, _err := client.DescribeEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventTopWithOptions(request *DescribeEventTopRequest, runtime *util.RuntimeOptions) (_result *DescribeEventTopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEventTopResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEventTop"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEventTop(request *DescribeEventTopRequest) (_result *DescribeEventTopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventTopResponse{}
	_body, _err := client.DescribeEventTopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventTypesWithOptions(request *DescribeEventTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeEventTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEventTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEventTypes"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEventTypes(request *DescribeEventTypesRequest) (_result *DescribeEventTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventTypesResponse{}
	_body, _err := client.DescribeEventTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventTypesDetailWithOptions(request *DescribeEventTypesDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeEventTypesDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEventTypesDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEventTypesDetail"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEventTypesDetail(request *DescribeEventTypesDetailRequest) (_result *DescribeEventTypesDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventTypesDetailResponse{}
	_body, _err := client.DescribeEventTypesDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancePortraitWithOptions(request *DescribeInstancePortraitRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancePortraitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancePortraitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstancePortrait"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstancePortrait(request *DescribeInstancePortraitRequest) (_result *DescribeInstancePortraitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancePortraitResponse{}
	_body, _err := client.DescribeInstancePortraitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSourcesWithOptions(request *DescribeInstanceSourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceSourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceSources"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSources(request *DescribeInstanceSourcesRequest) (_result *DescribeInstanceSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSourcesResponse{}
	_body, _err := client.DescribeInstanceSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssObjectDetailWithOptions(request *DescribeOssObjectDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeOssObjectDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOssObjectDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOssObjectDetail"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssObjectDetail(request *DescribeOssObjectDetailRequest) (_result *DescribeOssObjectDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssObjectDetailResponse{}
	_body, _err := client.DescribeOssObjectDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssObjectsWithOptions(request *DescribeOssObjectsRequest, runtime *util.RuntimeOptions) (_result *DescribeOssObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOssObjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOssObjects"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssObjects(request *DescribeOssObjectsRequest) (_result *DescribeOssObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssObjectsResponse{}
	_body, _err := client.DescribeOssObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackagesWithOptions(request *DescribePackagesRequest, runtime *util.RuntimeOptions) (_result *DescribePackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePackagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePackages"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePackages(request *DescribePackagesRequest) (_result *DescribePackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackagesResponse{}
	_body, _err := client.DescribePackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRiskLevelsWithOptions(request *DescribeRiskLevelsRequest, runtime *util.RuntimeOptions) (_result *DescribeRiskLevelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRiskLevelsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRiskLevels"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRiskLevels(request *DescribeRiskLevelsRequest) (_result *DescribeRiskLevelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRiskLevelsResponse{}
	_body, _err := client.DescribeRiskLevelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleCategoryWithOptions(request *DescribeRuleCategoryRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRuleCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRuleCategory"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleCategory(request *DescribeRuleCategoryRequest) (_result *DescribeRuleCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleCategoryResponse{}
	_body, _err := client.DescribeRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRulesWithOptions(request *DescribeRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRules"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRules(request *DescribeRulesRequest) (_result *DescribeRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRulesResponse{}
	_body, _err := client.DescribeRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTablesWithOptions(request *DescribeTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTablesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTables"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTables(request *DescribeTablesRequest) (_result *DescribeTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTablesResponse{}
	_body, _err := client.DescribeTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserCountWithOptions(request *DescribeUserCountRequest, runtime *util.RuntimeOptions) (_result *DescribeUserCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserCount"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserCount(request *DescribeUserCountRequest) (_result *DescribeUserCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserCountResponse{}
	_body, _err := client.DescribeUserCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserStatusWithOptions(request *DescribeUserStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserStatus"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserStatus(request *DescribeUserStatusRequest) (_result *DescribeUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserStatusResponse{}
	_body, _err := client.DescribeUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableUserConfigWithOptions(request *DisableUserConfigRequest, runtime *util.RuntimeOptions) (_result *DisableUserConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableUserConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableUserConfig"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableUserConfig(request *DisableUserConfigRequest) (_result *DisableUserConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableUserConfigResponse{}
	_body, _err := client.DisableUserConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecDatamaskWithOptions(request *ExecDatamaskRequest, runtime *util.RuntimeOptions) (_result *ExecDatamaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecDatamaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecDatamask"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecDatamask(request *ExecDatamaskRequest) (_result *ExecDatamaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecDatamaskResponse{}
	_body, _err := client.ExecDatamaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ManualTriggerMaskingProcessWithOptions(request *ManualTriggerMaskingProcessRequest, runtime *util.RuntimeOptions) (_result *ManualTriggerMaskingProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ManualTriggerMaskingProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ManualTriggerMaskingProcess"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ManualTriggerMaskingProcess(request *ManualTriggerMaskingProcessRequest) (_result *ManualTriggerMaskingProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ManualTriggerMaskingProcessResponse{}
	_body, _err := client.ManualTriggerMaskingProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDataLimitWithOptions(request *ModifyDataLimitRequest, runtime *util.RuntimeOptions) (_result *ModifyDataLimitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDataLimitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDataLimit"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDataLimit(request *ModifyDataLimitRequest) (_result *ModifyDataLimitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDataLimitResponse{}
	_body, _err := client.ModifyDataLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefaultLevelWithOptions(request *ModifyDefaultLevelRequest, runtime *util.RuntimeOptions) (_result *ModifyDefaultLevelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDefaultLevelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDefaultLevel"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefaultLevel(request *ModifyDefaultLevelRequest) (_result *ModifyDefaultLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefaultLevelResponse{}
	_body, _err := client.ModifyDefaultLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyEventStatusWithOptions(request *ModifyEventStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyEventStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyEventStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyEventStatus"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyEventStatus(request *ModifyEventStatusRequest) (_result *ModifyEventStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyEventStatusResponse{}
	_body, _err := client.ModifyEventStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyEventTypeStatusWithOptions(request *ModifyEventTypeStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyEventTypeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyEventTypeStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyEventTypeStatus"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyEventTypeStatus(request *ModifyEventTypeStatusRequest) (_result *ModifyEventTypeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyEventTypeStatusResponse{}
	_body, _err := client.ModifyEventTypeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRuleWithOptions(request *ModifyRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyRule"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRule(request *ModifyRuleRequest) (_result *ModifyRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRuleResponse{}
	_body, _err := client.ModifyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRuleStatusWithOptions(request *ModifyRuleStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyRuleStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyRuleStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyRuleStatus"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRuleStatus(request *ModifyRuleStatusRequest) (_result *ModifyRuleStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRuleStatusResponse{}
	_body, _err := client.ModifyRuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopMaskingProcessWithOptions(request *StopMaskingProcessRequest, runtime *util.RuntimeOptions) (_result *StopMaskingProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopMaskingProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopMaskingProcess"), tea.String("2019-01-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopMaskingProcess(request *StopMaskingProcessRequest) (_result *StopMaskingProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopMaskingProcessResponse{}
	_body, _err := client.StopMaskingProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

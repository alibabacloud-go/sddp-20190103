// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateConfigRequest struct {
	// The code of the common configuration item. Valid values:
	//
	// 	- **access_failed_cnt**: the maximum number of access attempts allowed when Data Security Center (DSC) fails to access an unauthorized resource.
	//
	// 	- **access_permission_exprie_max_days**: the maximum idle period allowed for access permissions before an alert is triggered.
	//
	// 	- **log_datasize_avg_days**: the minimum percentage of the volume of logs of a specific type generated on the current day to the average volume of logs generated in the previous 10 days before an alert is triggered.
	//
	// example:
	//
	// access_failed_cnt
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the common configuration item.
	//
	// example:
	//
	// Maximum number of access attempts allowed when DSC fails to access an unauthorized resource: 10
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The value of the common configuration item. The meaning of this parameter varies with the value of the Code parameter.
	//
	// 	- If you set the Code parameter to **access_failed_cnt**, the Value parameter specifies the maximum number of access attempts allowed when DSC fails to access an unauthorized resource.
	//
	// 	- If you set the Code parameter to **access_permission_exprie_max_days**, the Value parameter specifies the maximum idle period allowed for access permissions before an alert is triggered.
	//
	// 	- If you set the Code parameter to **log_datasize_avg_days**, the Value parameter specifies the minimum percentage of the volume of logs of a specific type generated on the current day to the average amount of logs generated in the previous 10 days before an alert is triggered.
	//
	// example:
	//
	// 30
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigRequest) SetCode(v string) *CreateConfigRequest {
	s.Code = &v
	return s
}

func (s *CreateConfigRequest) SetDescription(v string) *CreateConfigRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigRequest) SetFeatureType(v int32) *CreateConfigRequest {
	s.FeatureType = &v
	return s
}

func (s *CreateConfigRequest) SetLang(v string) *CreateConfigRequest {
	s.Lang = &v
	return s
}

func (s *CreateConfigRequest) SetSourceIp(v string) *CreateConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateConfigRequest) SetValue(v string) *CreateConfigRequest {
	s.Value = &v
	return s
}

type CreateConfigResponseBody struct {
	// The ID of the common alert configuration.
	//
	// example:
	//
	// 12300
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1EBF271
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateConfigResponse) SetStatusCode(v int32) *CreateConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigResponse) SetBody(v *CreateConfigResponseBody) *CreateConfigResponse {
	s.Body = v
	return s
}

type CreateDataLimitRequest struct {
	// Specifies whether to enable the security audit feature. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Specifies whether to automatically trigger a re-scan after a rule is modified. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// > When a re-scan is triggered, DSC scans all data in your data asset.
	//
	// example:
	//
	// 1
	AutoScan *int32 `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	// The permissions. Valid values:
	//
	// 	- **ReadOnly**: read-only permissions
	//
	// 	- **ReadWrite**: read and write permissions
	//
	// example:
	//
	// ReadOnly
	CertificatePermission *string `json:"CertificatePermission,omitempty" xml:"CertificatePermission,omitempty"`
	// Specifies whether to enable sensitive data detection. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// > If this is your first time to authorize DSC to access the data asset, the default value is 1. If this is not your first time to authorize DSC to access the data asset, the default value is the same as that used in the last authorization operation. Both 1 and 0 are possible.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The database engine that is run by the instance. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// Specifies whether to enable anomalous event detection. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes (default)
	//
	// example:
	//
	// 1
	EventStatus *int32 `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// Specifies whether to immediately scan the authorized asset. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	InstantlyScan *bool `json:"InstantlyScan,omitempty" xml:"InstantlyScan,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The retention period of raw logs after you enable the security audit feature. Unit: days. Valid values:
	//
	// 	- **30**
	//
	// 	- **90**
	//
	// 	- **180**
	//
	// 	- **365**
	//
	// example:
	//
	// 30
	LogStoreDay *int32 `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	// Specifies whether to enable optical character recognition (OCR). Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 0
	OcrStatus *int32 `json:"OcrStatus,omitempty" xml:"OcrStatus,omitempty"`
	// The name of the asset. The value is a connection string. It consists of an instance ID and a database name, which are separated by a comma (,). This parameter is required.
	//
	// example:
	//
	// test-11**
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The password that is used to access the database.
	//
	// example:
	//
	// passwd
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port that is used to connect to the database.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of service to which the data asset belongs. Valid values:
	//
	// 	- **1*	- :MaxCompute
	//
	// 	- **2**: Object Storage Service (OSS)
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4*	- :Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of sensitive data samples that are collected after sensitive data detection is enabled. Valid values:
	//
	// 	- **0**
	//
	// 	- **5**
	//
	// 	- **10**
	//
	// example:
	//
	// 0
	SamplingSize *int32 `json:"SamplingSize,omitempty" xml:"SamplingSize,omitempty"`
	// The region in which the data asset resides. Valid values:
	//
	// 	- **cn-beijing**: China (Beijing).
	//
	// 	- **cn-zhangjiakou**: China (Zhangjiakou)
	//
	// 	- **cn-huhehaote**: China (Hohhot)
	//
	// 	- **cn-hangzhou**: China (Hangzhou)
	//
	// 	- **cn-shanghai**: China (Shanghai)
	//
	// 	- **cn-shenzhen**: China (Shenzhen)
	//
	// 	- **cn-hongkong**: China (Hong Kong)
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The username that is used to access the database.
	//
	// example:
	//
	// yhm
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateDataLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataLimitRequest) GoString() string {
	return s.String()
}

func (s *CreateDataLimitRequest) SetAuditStatus(v int32) *CreateDataLimitRequest {
	s.AuditStatus = &v
	return s
}

func (s *CreateDataLimitRequest) SetAutoScan(v int32) *CreateDataLimitRequest {
	s.AutoScan = &v
	return s
}

func (s *CreateDataLimitRequest) SetCertificatePermission(v string) *CreateDataLimitRequest {
	s.CertificatePermission = &v
	return s
}

func (s *CreateDataLimitRequest) SetEnable(v int32) *CreateDataLimitRequest {
	s.Enable = &v
	return s
}

func (s *CreateDataLimitRequest) SetEngineType(v string) *CreateDataLimitRequest {
	s.EngineType = &v
	return s
}

func (s *CreateDataLimitRequest) SetEventStatus(v int32) *CreateDataLimitRequest {
	s.EventStatus = &v
	return s
}

func (s *CreateDataLimitRequest) SetFeatureType(v int32) *CreateDataLimitRequest {
	s.FeatureType = &v
	return s
}

func (s *CreateDataLimitRequest) SetInstantlyScan(v bool) *CreateDataLimitRequest {
	s.InstantlyScan = &v
	return s
}

func (s *CreateDataLimitRequest) SetLang(v string) *CreateDataLimitRequest {
	s.Lang = &v
	return s
}

func (s *CreateDataLimitRequest) SetLogStoreDay(v int32) *CreateDataLimitRequest {
	s.LogStoreDay = &v
	return s
}

func (s *CreateDataLimitRequest) SetOcrStatus(v int32) *CreateDataLimitRequest {
	s.OcrStatus = &v
	return s
}

func (s *CreateDataLimitRequest) SetParentId(v string) *CreateDataLimitRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDataLimitRequest) SetPassword(v string) *CreateDataLimitRequest {
	s.Password = &v
	return s
}

func (s *CreateDataLimitRequest) SetPort(v int32) *CreateDataLimitRequest {
	s.Port = &v
	return s
}

func (s *CreateDataLimitRequest) SetResourceType(v int32) *CreateDataLimitRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateDataLimitRequest) SetSamplingSize(v int32) *CreateDataLimitRequest {
	s.SamplingSize = &v
	return s
}

func (s *CreateDataLimitRequest) SetServiceRegionId(v string) *CreateDataLimitRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *CreateDataLimitRequest) SetSourceIp(v string) *CreateDataLimitRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateDataLimitRequest) SetUserName(v string) *CreateDataLimitRequest {
	s.UserName = &v
	return s
}

type CreateDataLimitResponseBody struct {
	// The ID of the data asset.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateDataLimitResponse) SetStatusCode(v int32) *CreateDataLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataLimitResponse) SetBody(v *CreateDataLimitResponseBody) *CreateDataLimitResponse {
	s.Body = v
	return s
}

type CreateRuleRequest struct {
	// The content type of the sensitive data detection rule. Valid values:
	//
	// 	- **0**: keyword
	//
	// 	- **2**: regular expression
	//
	// example:
	//
	// 0
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content of the sensitive data detection rule. You can specify a regular expression or keywords that are used to match sensitive fields or text.
	//
	// This parameter is required.
	//
	// example:
	//
	// (?:\\\\D|^)((?:(?:25[0-4]|2[0-4]\\\\d|1\\\\d{2}|[1-9]\\\\d{1})\\\\.)(?:(?:25[0-5]|2[0-4]\\\\d|[01]?\\\\d?\\\\d)\\\\.){2}(?:25[0-5]|2[0-4]\\\\d|1[0-9]\\\\d|[1-9]\\\\d|[1-9]))(?:\\\\D|$)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The type of the content in the sensitive data detection rule. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates attempts to exploit SQL injections. The value 2 indicates bypass by using SQL injections. The value 3 indicates abuse of stored procedures. The value 4 indicates buffer overflow. The value 5 indicates SQL injections based on errors.
	//
	// example:
	//
	// 1
	ContentCategory *int32 `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// ID card
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The match type. Valid values:
	//
	// 	- **1**: rule-based match
	//
	// 	- **2**: dictionary-based match
	//
	// example:
	//
	// 1
	MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The IDs of the models for sensitive data audit.
	//
	// example:
	//
	// 1452
	ModelRuleIds *string `json:"ModelRuleIds,omitempty" xml:"ModelRuleIds,omitempty"`
	// The name of the sensitive data detection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-tst
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the service to which data in the column of the table belongs. Valid values include **MaxCompute**, **OSS**, **ADS**, **OTS**, and **RDS**.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data asset belongs. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 2
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The type of the sensitive data detection rule. Valid values:
	//
	// 	- **1**: sensitive data detection rule
	//
	// 	- **2**: audit rule
	//
	// 	- **3**: anomalous event detection rule
	//
	// 	- **99**: custom rule
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The statistical expression.
	//
	// example:
	//
	// 1
	StatExpress *string `json:"StatExpress,omitempty" xml:"StatExpress,omitempty"`
	// Specifies whether to enable the sensitive data detection rule. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the data asset. Valid values:
	//
	// 	- **0**: all data assets
	//
	// 	- **1**: structured data asset
	//
	// 	- **2**: unstructured data asset
	//
	// > If you set the parameter to 1 or 2, rules that support all data assets and rules that support the queried data asset type are returned.
	//
	// example:
	//
	// 1
	SupportForm *int32 `json:"SupportForm,omitempty" xml:"SupportForm,omitempty"`
	// The code of the service to which the sensitive data detection rule is applied. Valid values include **MaxCompute**, **OSS**, **ADS**, **OTS**, and **RDS**.
	//
	// example:
	//
	// MaxCompute
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The IDs of the templates that are used to audit sensitive data.
	//
	// example:
	//
	// 1
	TemplateRuleIds *string `json:"TemplateRuleIds,omitempty" xml:"TemplateRuleIds,omitempty"`
	// The risk level of the alert that is triggered. Valid values:
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
	//
	// example:
	//
	// 2
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) SetCategory(v int32) *CreateRuleRequest {
	s.Category = &v
	return s
}

func (s *CreateRuleRequest) SetContent(v string) *CreateRuleRequest {
	s.Content = &v
	return s
}

func (s *CreateRuleRequest) SetContentCategory(v int32) *CreateRuleRequest {
	s.ContentCategory = &v
	return s
}

func (s *CreateRuleRequest) SetDescription(v string) *CreateRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateRuleRequest) SetLang(v string) *CreateRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateRuleRequest) SetMatchType(v int32) *CreateRuleRequest {
	s.MatchType = &v
	return s
}

func (s *CreateRuleRequest) SetModelRuleIds(v string) *CreateRuleRequest {
	s.ModelRuleIds = &v
	return s
}

func (s *CreateRuleRequest) SetName(v string) *CreateRuleRequest {
	s.Name = &v
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

func (s *CreateRuleRequest) SetRiskLevelId(v int64) *CreateRuleRequest {
	s.RiskLevelId = &v
	return s
}

func (s *CreateRuleRequest) SetRuleType(v int32) *CreateRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateRuleRequest) SetSourceIp(v string) *CreateRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateRuleRequest) SetStatExpress(v string) *CreateRuleRequest {
	s.StatExpress = &v
	return s
}

func (s *CreateRuleRequest) SetStatus(v int32) *CreateRuleRequest {
	s.Status = &v
	return s
}

func (s *CreateRuleRequest) SetSupportForm(v int32) *CreateRuleRequest {
	s.SupportForm = &v
	return s
}

func (s *CreateRuleRequest) SetTarget(v string) *CreateRuleRequest {
	s.Target = &v
	return s
}

func (s *CreateRuleRequest) SetTemplateRuleIds(v string) *CreateRuleRequest {
	s.TemplateRuleIds = &v
	return s
}

func (s *CreateRuleRequest) SetWarnLevel(v int32) *CreateRuleRequest {
	s.WarnLevel = &v
	return s
}

type CreateRuleResponseBody struct {
	// The unique ID of the sensitive data detection rule.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1EBF271
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateRuleResponse) SetStatusCode(v int32) *CreateRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRuleResponse) SetBody(v *CreateRuleResponseBody) *CreateRuleResponse {
	s.Body = v
	return s
}

type CreateScanTaskRequest struct {
	// The unique ID of the data asset, such as an instance, a database, and a bucket. You can call the [DescribeDataLimits](~~DescribeDataLimits~~) operation to query the unique ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataLimitId *int64 `json:"DataLimitId,omitempty" xml:"DataLimitId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The interval between two consecutive custom scan tasks. Unit: days. Valid values: 1 to 2147483648.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	IntervalDay *int32 `json:"IntervalDay,omitempty" xml:"IntervalDay,omitempty"`
	// The language of the content within the request and response.
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The data to be scanned in the Object Storage Service (OSS) bucket. Prefix match, suffix match, and regular expression match are supported.
	//
	// example:
	//
	// /test/test
	OssScanPath *string `json:"OssScanPath,omitempty" xml:"OssScanPath,omitempty"`
	// The type of the service to which the data assets to be scanned belong. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates OSS. The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The time when the scan task is executed next time. Unit: hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	RunHour *int32 `json:"RunHour,omitempty" xml:"RunHour,omitempty"`
	// The time when the scan task is executed next time. Unit: minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	RunMinute *int32 `json:"RunMinute,omitempty" xml:"RunMinute,omitempty"`
	// The matching rule that specifies the scan scope of the custom scan task. This parameter takes effect only if you set the **ScanRangeContent*	- parameter. Valid values:
	//
	// 	- **0**: exact match
	//
	// 	- **1**: prefix match
	//
	// 	- **2**: suffix match
	//
	// 	- **3**: regular expression match
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ScanRange *int32 `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
	// The data to be scanned in a structured data asset. Prefix match, suffix match, and regular expression match are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// datamask/
	ScanRangeContent *string `json:"ScanRangeContent,omitempty" xml:"ScanRangeContent,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The name of the scan task.
	//
	// This parameter is required.
	//
	// example:
	//
	// scan-test-sample****
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The account that is used to create the scan task.
	//
	// example:
	//
	// demo
	TaskUserName *string `json:"TaskUserName,omitempty" xml:"TaskUserName,omitempty"`
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

func (s *CreateScanTaskRequest) SetFeatureType(v int32) *CreateScanTaskRequest {
	s.FeatureType = &v
	return s
}

func (s *CreateScanTaskRequest) SetIntervalDay(v int32) *CreateScanTaskRequest {
	s.IntervalDay = &v
	return s
}

func (s *CreateScanTaskRequest) SetLang(v string) *CreateScanTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateScanTaskRequest) SetOssScanPath(v string) *CreateScanTaskRequest {
	s.OssScanPath = &v
	return s
}

func (s *CreateScanTaskRequest) SetResourceType(v int64) *CreateScanTaskRequest {
	s.ResourceType = &v
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

func (s *CreateScanTaskRequest) SetScanRange(v int32) *CreateScanTaskRequest {
	s.ScanRange = &v
	return s
}

func (s *CreateScanTaskRequest) SetScanRangeContent(v string) *CreateScanTaskRequest {
	s.ScanRangeContent = &v
	return s
}

func (s *CreateScanTaskRequest) SetSourceIp(v string) *CreateScanTaskRequest {
	s.SourceIp = &v
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

type CreateScanTaskResponseBody struct {
	// The ID of the custom scan task.
	//
	// example:
	//
	// 100
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B1F2BB1F-04EC-5D36-B136-B4DE17FD8DE0
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateScanTaskResponse) SetStatusCode(v int32) *CreateScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScanTaskResponse) SetBody(v *CreateScanTaskResponseBody) *CreateScanTaskResponse {
	s.Body = v
	return s
}

type CreateSlrRoleRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateSlrRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSlrRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateSlrRoleRequest) SetFeatureType(v int32) *CreateSlrRoleRequest {
	s.FeatureType = &v
	return s
}

func (s *CreateSlrRoleRequest) SetLang(v string) *CreateSlrRoleRequest {
	s.Lang = &v
	return s
}

func (s *CreateSlrRoleRequest) SetSourceIp(v string) *CreateSlrRoleRequest {
	s.SourceIp = &v
	return s
}

type CreateSlrRoleResponseBody struct {
	// Indicates whether the service-linked role was created. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	HasPermission *bool `json:"HasPermission,omitempty" xml:"HasPermission,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1EBF271
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSlrRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSlrRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSlrRoleResponseBody) SetHasPermission(v bool) *CreateSlrRoleResponseBody {
	s.HasPermission = &v
	return s
}

func (s *CreateSlrRoleResponseBody) SetRequestId(v string) *CreateSlrRoleResponseBody {
	s.RequestId = &v
	return s
}

type CreateSlrRoleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSlrRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSlrRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSlrRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateSlrRoleResponse) SetHeaders(v map[string]*string) *CreateSlrRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateSlrRoleResponse) SetStatusCode(v int32) *CreateSlrRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSlrRoleResponse) SetBody(v *CreateSlrRoleResponseBody) *CreateSlrRoleResponse {
	s.Body = v
	return s
}

type DeleteDataLimitRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The ID of the data asset.
	//
	// You can call the DescribeDataLimits operation to query the IDs of data assets. The value of the Id response parameter indicates the ID of a data asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12033
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteDataLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataLimitRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLimitRequest) SetFeatureType(v int32) *DeleteDataLimitRequest {
	s.FeatureType = &v
	return s
}

func (s *DeleteDataLimitRequest) SetId(v int64) *DeleteDataLimitRequest {
	s.Id = &v
	return s
}

func (s *DeleteDataLimitRequest) SetLang(v string) *DeleteDataLimitRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDataLimitRequest) SetSourceIp(v string) *DeleteDataLimitRequest {
	s.SourceIp = &v
	return s
}

type DeleteDataLimitResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteDataLimitResponse) SetStatusCode(v int32) *DeleteDataLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLimitResponse) SetBody(v *DeleteDataLimitResponseBody) *DeleteDataLimitResponse {
	s.Body = v
	return s
}

type DeleteRuleRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The ID of the sensitive data detection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 122300
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values: **zh*	- and **en**. The value zh indicates Chinese, and the value en indicates English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) SetFeatureType(v int32) *DeleteRuleRequest {
	s.FeatureType = &v
	return s
}

func (s *DeleteRuleRequest) SetId(v int64) *DeleteRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteRuleRequest) SetLang(v string) *DeleteRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteRuleRequest) SetSourceIp(v string) *DeleteRuleRequest {
	s.SourceIp = &v
	return s
}

type DeleteRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B6*****
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteRuleResponse) SetStatusCode(v int32) *DeleteRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRuleResponse) SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse {
	s.Body = v
	return s
}

type DescribeCategoryTemplateListRequest struct {
	// Page number for paginated queries. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// Set the language type for the request and response messages, default is **zh_cn**.
	//
	// Values:
	//
	// - **zh_cn**: Chinese (Simplified)
	//
	// - **en_us**: English (United States)
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// When performing a paginated query, set the number of items per page. Default value is **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// API call scenario, default is **null**.
	//
	// Values:
	//
	// - **null**: Old version
	//
	// - **0**: Old version
	//
	// - **1**: New version
	//
	// example:
	//
	// 1
	UsageScenario *int32 `json:"UsageScenario,omitempty" xml:"UsageScenario,omitempty"`
}

func (s DescribeCategoryTemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryTemplateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateListRequest) SetCurrentPage(v int32) *DescribeCategoryTemplateListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCategoryTemplateListRequest) SetFeatureType(v int32) *DescribeCategoryTemplateListRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeCategoryTemplateListRequest) SetLang(v string) *DescribeCategoryTemplateListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCategoryTemplateListRequest) SetPageSize(v int32) *DescribeCategoryTemplateListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCategoryTemplateListRequest) SetUsageScenario(v int32) *DescribeCategoryTemplateListRequest {
	s.UsageScenario = &v
	return s
}

type DescribeCategoryTemplateListResponseBody struct {
	// Page number for paginated queries. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of industry templates.
	Items []*DescribeCategoryTemplateListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Number of items per page in a paginated query. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier generated by Alibaba Cloud for this request.
	//
	// example:
	//
	// 8491DBFD-48C0-4E11-B6FC-6F38921244A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of data items returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCategoryTemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateListResponseBody) SetCurrentPage(v int32) *DescribeCategoryTemplateListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBody) SetItems(v []*DescribeCategoryTemplateListResponseBodyItems) *DescribeCategoryTemplateListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCategoryTemplateListResponseBody) SetPageSize(v int32) *DescribeCategoryTemplateListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBody) SetRequestId(v string) *DescribeCategoryTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBody) SetTotalCount(v int32) *DescribeCategoryTemplateListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCategoryTemplateListResponseBodyItems struct {
	// Current risk level ID.
	//
	// example:
	//
	// 5
	CurrentRiskLevel *int32 `json:"CurrentRiskLevel,omitempty" xml:"CurrentRiskLevel,omitempty"`
	// Description information of the industry template.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Industry template creation time.
	//
	// example:
	//
	// 1582992000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Industry template modification time.
	//
	// example:
	//
	// 1545277010000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Unique identifier ID of the industry template.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Maximum category level.
	//
	// example:
	//
	// 2
	MaxCategoryLevel *int32 `json:"MaxCategoryLevel,omitempty" xml:"MaxCategoryLevel,omitempty"`
	// Maximum risk level ID.
	//
	// example:
	//
	// 5
	MaxRiskLevel *int32 `json:"MaxRiskLevel,omitempty" xml:"MaxRiskLevel,omitempty"`
	// Name of the industry template.
	//
	// example:
	//
	// built-in template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Status of the industry template. Values:
	//
	// - **0**: Closed status.
	//
	// - **1**: Enabled status, user\\"s current main template.
	//
	// - **2**: Active status, both enabled and active templates can be applied to recognition tasks.
	//
	// - **3**: General recognition model template status.
	//
	// > Templates in enabled and active status can be used as the industry template ID parameter for [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html).
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Whether the industry template supports editing. Values:
	//
	// - **0**: Not supported.
	//
	// - **1**: Supported.
	//
	// example:
	//
	// 0
	SupportEdit *int32 `json:"SupportEdit,omitempty" xml:"SupportEdit,omitempty"`
	// Type of the industry template.
	//
	// example:
	//
	// 6
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCategoryTemplateListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryTemplateListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetCurrentRiskLevel(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.CurrentRiskLevel = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetDescription(v string) *DescribeCategoryTemplateListResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetGmtCreate(v int64) *DescribeCategoryTemplateListResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetGmtModified(v int64) *DescribeCategoryTemplateListResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetId(v int64) *DescribeCategoryTemplateListResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetMaxCategoryLevel(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.MaxCategoryLevel = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetMaxRiskLevel(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.MaxRiskLevel = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetName(v string) *DescribeCategoryTemplateListResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetStatus(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetSupportEdit(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.SupportEdit = &v
	return s
}

func (s *DescribeCategoryTemplateListResponseBodyItems) SetType(v int32) *DescribeCategoryTemplateListResponseBodyItems {
	s.Type = &v
	return s
}

type DescribeCategoryTemplateListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCategoryTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCategoryTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryTemplateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateListResponse) SetHeaders(v map[string]*string) *DescribeCategoryTemplateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCategoryTemplateListResponse) SetStatusCode(v int32) *DescribeCategoryTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCategoryTemplateListResponse) SetBody(v *DescribeCategoryTemplateListResponseBody) *DescribeCategoryTemplateListResponse {
	s.Body = v
	return s
}

type DescribeCategoryTemplateRuleListRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sensitivity level of the data that is not compliant with the rule. Valid values: **1*	- to **11**. Default value: **null**.
	//
	// 	- **1**: No sensitive data is detected.
	//
	// 	- **2**: specifies the S1 sensitivity level.
	//
	// 	- **3**: specifies the S2 sensitivity level.
	//
	// 	- **4**: specifies the S3 sensitivity level.
	//
	// 	- **5**: specifies the S4 sensitivity level.
	//
	// 	- **6**: specifies the S5 sensitivity level.
	//
	// 	- **7**: specifies the S6 sensitivity level.
	//
	// 	- **8**: specifies the S7 sensitivity level.
	//
	// 	- **9**: specifies the S8 sensitivity level.
	//
	// 	- **10**: specifies the S9 sensitivity level.
	//
	// 	- **11**: specifies the S10 sensitivity level.
	//
	// 	- **null**: specifies all preceding sensitivity levels.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The status of the rule. Default value: **null**. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// 	- **null**: all states
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCategoryTemplateRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryTemplateRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateRuleListRequest) SetCurrentPage(v int32) *DescribeCategoryTemplateRuleListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListRequest) SetFeatureType(v int32) *DescribeCategoryTemplateRuleListRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListRequest) SetLang(v string) *DescribeCategoryTemplateRuleListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListRequest) SetPageSize(v int32) *DescribeCategoryTemplateRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListRequest) SetRiskLevelId(v int64) *DescribeCategoryTemplateRuleListRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListRequest) SetStatus(v int32) *DescribeCategoryTemplateRuleListRequest {
	s.Status = &v
	return s
}

type DescribeCategoryTemplateRuleListResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of rules.
	Items []*DescribeCategoryTemplateRuleListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 136082B3-B21F-5E9D-B68E-991FFD205D24
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of rules in the template.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCategoryTemplateRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryTemplateRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateRuleListResponseBody) SetCurrentPage(v int32) *DescribeCategoryTemplateRuleListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBody) SetItems(v []*DescribeCategoryTemplateRuleListResponseBodyItems) *DescribeCategoryTemplateRuleListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBody) SetPageSize(v int32) *DescribeCategoryTemplateRuleListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBody) SetRequestId(v string) *DescribeCategoryTemplateRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBody) SetTotalCount(v int32) *DescribeCategoryTemplateRuleListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCategoryTemplateRuleListResponseBodyItems struct {
	// The description of the rule.
	//
	// example:
	//
	// Rule for identifying ID card numbers
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the rule.
	//
	// example:
	//
	// 100
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IDs of sensitive data types. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 1001,1002
	IdentificationRuleIds *string `json:"IdentificationRuleIds,omitempty" xml:"IdentificationRuleIds,omitempty"`
	// The scan scope of the rule. The value is a JSON array of the STRING type. Each element in a JSON array indicates a scan scope that contains the following fields:
	//
	// 	- **Asset**: the data asset type. Valid values: RDS, DRDS, PolarDB, OTS, ADB, and OceanBase. The value is of the STRING type.
	//
	// 	- **Content**: the scan scope. The value is of the STRING type. Each element in a JSON array indicates a scan scope that contains the following fields:
	//
	//     	- **Range**: the matching condition. Valid values: Instance, database, table, column, project, bucket, and object. The value project is valid only for data assets in MaxCompute. The values bucket and object are valid only for data assets in Object Storage Service (OSS). The value of this parameter is of the STRING type.
	//
	//     	- **Operator**: the operator. Valid values: equals, regex, prefix, and suffix. The operator equals indicates a full match. The operator regex indicates a match by regular expression. The operator prefix indicates a match by prefix. The operator suffix indicates a match by suffix.
	//
	//     	- **Value**: the matching content. The value is of the STRING type.
	//
	// example:
	//
	// [{"Asset":"RDS","Content":[{"Range":"database","Operator":"regex","Value":"register"}]},{"Asset":"RDS","Content":[{"Range":"table","Operator":"regex","Value":"register"}]},{"Asset":"RDS","Content":[{"Range":"column","Operator":"regex","Value":"register"}]},{"Asset":"ODPS","Content":[{"Range":"project","Operator":"regex","Value":"register"}]},{"Asset":"ODPS","Content":[{"Range":"table","Operator":"regex","Value":"register"}]},{"Asset":"ODPS","Content":[{"Range":"column","Operator":"regex","Value":"register"}]}]
	IdentificationScope *string `json:"IdentificationScope,omitempty" xml:"IdentificationScope,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// ID card number
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sensitivity level of the data that is not compliant with the rule. Valid values: **1*	- to **11**.
	//
	// 	- **1**: No sensitive data is detected.
	//
	// 	- **2**: indicates the S1 sensitivity level.
	//
	// 	- **3**: indicates the S2 sensitivity level.
	//
	// 	- **4**: indicates the S3 sensitivity level.
	//
	// 	- **5**: indicates the S4 sensitivity level.
	//
	// 	- **6**: indicates the S5 sensitivity level.
	//
	// 	- **7**: indicates the S6 sensitivity level.
	//
	// 	- **8**: indicates the S7 sensitivity level.
	//
	// 	- **9**: indicates the S8 sensitivity level.
	//
	// 	- **10**: indicates the S9 sensitivity level.
	//
	// 	- **11**: indicates the S10 sensitivity level.
	//
	// 	- **null**: indicates all preceding sensitivity levels.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// 	- **null**: all states
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCategoryTemplateRuleListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryTemplateRuleListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetDescription(v string) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetId(v int64) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetIdentificationRuleIds(v string) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.IdentificationRuleIds = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetIdentificationScope(v string) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.IdentificationScope = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetName(v string) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetRiskLevelId(v int64) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponseBodyItems) SetStatus(v int32) *DescribeCategoryTemplateRuleListResponseBodyItems {
	s.Status = &v
	return s
}

type DescribeCategoryTemplateRuleListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCategoryTemplateRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCategoryTemplateRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCategoryTemplateRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateRuleListResponse) SetHeaders(v map[string]*string) *DescribeCategoryTemplateRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponse) SetStatusCode(v int32) *DescribeCategoryTemplateRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponse) SetBody(v *DescribeCategoryTemplateRuleListResponseBody) *DescribeCategoryTemplateRuleListResponse {
	s.Body = v
	return s
}

type DescribeColumnsRequest struct {
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The engine type. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **MariaDB**
	//
	// 	- **Oracle**
	//
	// 	- **PostgreSQL**
	//
	// 	- **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The ID of the instance to which data in the column of the table belongs.
	//
	// > You can call the [DescribeInstances](~~DescribeRules~~) operation to query the IDs of instances.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance to which data in the column of the table belongs.
	//
	// example:
	//
	// rm-bp17t1htja573l5i8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The data tag.
	//
	// 	- 101: personal sensitive information
	//
	// 	- 102: personal information
	//
	// example:
	//
	// 101
	ModelTagId *string `json:"ModelTagId,omitempty" xml:"ModelTagId,omitempty"`
	// The search keyword. Fuzzy match is supported.
	//
	// For example, if you enter **test**, all columns whose names contain **test*	- are retrieved.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service to which data in the column of the table belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data object belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: Object Storage Service (OSS)
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore (OTS)
	//
	// 	- **5**: ApsaraDB RDS
	//
	// 	- **6**: self-managed database
	//
	// 	- **7**: PolarDB for Xscale (PolarDB-X)
	//
	// 	- **8**: PolarDB
	//
	// 	- **9**: AnalyticDB for PostgreSQL
	//
	// 	- **10**: ApsaraDB for OceanBase
	//
	// 	- **11**: ApsaraDB for MongoDB
	//
	// 	- **25**: ApsaraDB for Redis
	//
	// example:
	//
	// 5
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **1**: N/A
	//
	// 	- **2**: S1
	//
	// 	- **3**: S2
	//
	// 	- **4**: S3
	//
	// 	- **5**: S4
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the sensitive data detection rule that data in the column of the table hits.
	//
	// > You can call the [DescribeRules](~~DescribeRules~~) operation to query the IDs of sensitive data detection rules.
	//
	// example:
	//
	// 11111
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the sensitive data detection rule that data in the column of the table hits.
	//
	// example:
	//
	// ID card number (the Chinese mainland)
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The name of the sensitivity level of the data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **N/A**: No sensitive data is detected.
	//
	// 	- **S1**: indicates the low sensitivity level.
	//
	// 	- **S2**: indicates the medium sensitivity level.
	//
	// 	- **S3**: indicates the high sensitivity level.
	//
	// 	- **S4**: indicates the highest sensitivity level.
	//
	// example:
	//
	// S2
	SensLevelName *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	// The ID of the table to which the column belongs.
	//
	// > You can call the [DescribeTables](~~DescribeTables~~) operation to query the IDs of tables.
	//
	// example:
	//
	// 11132334
	TableId *int64 `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the industry-specific classification template.
	//
	// >  You can call the [DescribeCategoryTemplateList](https://help.aliyun.com/document_detail/2399296.html) operation to obtain the IDs of industry-specific classification templates.
	//
	// example:
	//
	// 5
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the template rule that is hit.
	//
	// >  You can call the [DescribeCategoryTemplateRuleList](https://help.aliyun.com/document_detail/410143.html) operation to obtain the IDs of hit template rules.
	//
	// example:
	//
	// 1542
	TemplateRuleId *string `json:"TemplateRuleId,omitempty" xml:"TemplateRuleId,omitempty"`
}

func (s DescribeColumnsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeColumnsRequest) SetCurrentPage(v int32) *DescribeColumnsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeColumnsRequest) SetEngineType(v string) *DescribeColumnsRequest {
	s.EngineType = &v
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

func (s *DescribeColumnsRequest) SetLang(v string) *DescribeColumnsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeColumnsRequest) SetModelTagId(v string) *DescribeColumnsRequest {
	s.ModelTagId = &v
	return s
}

func (s *DescribeColumnsRequest) SetName(v string) *DescribeColumnsRequest {
	s.Name = &v
	return s
}

func (s *DescribeColumnsRequest) SetPageSize(v int32) *DescribeColumnsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeColumnsRequest) SetProductCode(v string) *DescribeColumnsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeColumnsRequest) SetProductId(v string) *DescribeColumnsRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeColumnsRequest) SetRiskLevelId(v int64) *DescribeColumnsRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeColumnsRequest) SetRuleId(v int64) *DescribeColumnsRequest {
	s.RuleId = &v
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

func (s *DescribeColumnsRequest) SetTableId(v int64) *DescribeColumnsRequest {
	s.TableId = &v
	return s
}

func (s *DescribeColumnsRequest) SetTableName(v string) *DescribeColumnsRequest {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsRequest) SetTemplateId(v string) *DescribeColumnsRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeColumnsRequest) SetTemplateRuleId(v string) *DescribeColumnsRequest {
	s.TemplateRuleId = &v
	return s
}

type DescribeColumnsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data in the columns of the table.
	Items []*DescribeColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-4******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeColumnsResponseBody) SetItems(v []*DescribeColumnsResponseBodyItems) *DescribeColumnsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeColumnsResponseBody) SetPageSize(v int32) *DescribeColumnsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeColumnsResponseBody) SetRequestId(v string) *DescribeColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColumnsResponseBody) SetTotalCount(v int32) *DescribeColumnsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeColumnsResponseBodyItems struct {
	// The time when the data in the column of the table is created. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The type of data in the column of the table.
	//
	// example:
	//
	// String
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The ID of the column of the table.
	//
	// example:
	//
	// 268
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the instance to which data in the column of the table belongs.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance to which data in the column of the table belongs.
	//
	// example:
	//
	// rm-bp17t1htja573l5i8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The column encryption status. Valid values:
	//
	// 	- **-1**: unencrypted
	//
	// 	- **1**: encrypted
	//
	// 	- **2**: encryption failed
	//
	// example:
	//
	// -1
	MaskingStatus *int32 `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	// A list of tags for data that hits the recognition model.
	ModelTags []*DescribeColumnsResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// The name of the column of the table.
	//
	// example:
	//
	// gxdata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the sensitivity level for asset. Valid values:
	//
	// 	- **N/A**: indicates that no sensitive data is detected.
	//
	// 	- **S1**: indicates the low sensitivity level.
	//
	// 	- **S2**: indicates the medium sensitivity level.
	//
	// 	- **S3**: indicates the high sensitivity level.
	//
	// 	- **S4**: indicates the highest sensitivity level.
	//
	// example:
	//
	// S3
	OdpsRiskLevelName *string `json:"OdpsRiskLevelName,omitempty" xml:"OdpsRiskLevelName,omitempty"`
	// The ID of the sensitivity level of the asset. Valid values:
	//
	// 	- **1**: N/A
	//
	// 	- **2**: S1
	//
	// 	- **3**: S2
	//
	// 	- **4**: S3
	//
	// 	- **5**: S4
	//
	// example:
	//
	// 3
	OdpsRiskLevelValue *int32 `json:"OdpsRiskLevelValue,omitempty" xml:"OdpsRiskLevelValue,omitempty"`
	// The name of the service to which data in the column of the table belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data object belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: Object Storage Service (OSS)
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore (OTS)
	//
	// 	- **5**: ApsaraDB RDS
	//
	// 	- **6**: self-managed database
	//
	// 	- **7**: PolarDB for Xscale (PolarDB-X)
	//
	// 	- **8**: PolarDB
	//
	// 	- **9**: AnalyticDB for PostgreSQL
	//
	// 	- **10**: ApsaraDB for OceanBase
	//
	// 	- **11**: ApsaraDB for MongoDB
	//
	// 	- **25**: ApsaraDB for Redis
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region in which the asset resides.
	//
	// example:
	//
	// cn-***
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the revision record.
	//
	// example:
	//
	// 12
	RevisionId *int64 `json:"RevisionId,omitempty" xml:"RevisionId,omitempty"`
	// Indicates whether the column is revised. Valid values:
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// example:
	//
	// 1
	RevisionStatus *int64 `json:"RevisionStatus,omitempty" xml:"RevisionStatus,omitempty"`
	// The ID of the sensitivity level of data in the column of the table. Valid values:
	//
	// 	- **1**: N/A
	//
	// 	- **2**: S1
	//
	// 	- **3**: S2
	//
	// 	- **4**: S3
	//
	// 	- **5**: S4
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level for data in the column of the table. Valid values:
	//
	// 	- **N/A**: indicates that no sensitive data is detected.
	//
	// 	- **S1**: indicates the low sensitivity level.
	//
	// 	- **S2**: indicates the medium sensitivity level.
	//
	// 	- **S3**: indicates the high sensitivity level.
	//
	// 	- **S4**: indicates the highest sensitivity level.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The ID of the sensitive data detection rule that data in the column of the table hits.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the sensitive data detection rule that data in the column of the table hits.
	//
	// example:
	//
	// \\*\\	- rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The name of the sensitivity level. Valid values:
	//
	// 	- **N/A**: indicates that no sensitive data is detected.
	//
	// 	- **S1**: indicates the low sensitivity level.
	//
	// 	- **S2**: indicates the medium sensitivity level.
	//
	// 	- **S3**: indicates the high sensitivity level.
	//
	// 	- **S4**: indicates the highest sensitivity level.
	//
	// example:
	//
	// S2
	SensLevelName *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	// Indicates whether the column contains sensitive data. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The ID of the table.
	//
	// example:
	//
	// 123
	TableId *int64 `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the table to which the revised column belongs.
	//
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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

func (s *DescribeColumnsResponseBodyItems) SetDataType(v string) *DescribeColumnsResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetEngineType(v string) *DescribeColumnsResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetId(v string) *DescribeColumnsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetInstanceId(v int64) *DescribeColumnsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetInstanceName(v string) *DescribeColumnsResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetMaskingStatus(v int32) *DescribeColumnsResponseBodyItems {
	s.MaskingStatus = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetModelTags(v []*DescribeColumnsResponseBodyItemsModelTags) *DescribeColumnsResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetName(v string) *DescribeColumnsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetOdpsRiskLevelName(v string) *DescribeColumnsResponseBodyItems {
	s.OdpsRiskLevelName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetOdpsRiskLevelValue(v int32) *DescribeColumnsResponseBodyItems {
	s.OdpsRiskLevelValue = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetProductCode(v string) *DescribeColumnsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetProductId(v int64) *DescribeColumnsResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRegionId(v string) *DescribeColumnsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRevisionId(v int64) *DescribeColumnsResponseBodyItems {
	s.RevisionId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRevisionStatus(v int64) *DescribeColumnsResponseBodyItems {
	s.RevisionStatus = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRiskLevelId(v int64) *DescribeColumnsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRiskLevelName(v string) *DescribeColumnsResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRuleId(v int64) *DescribeColumnsResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetRuleName(v string) *DescribeColumnsResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetSensLevelName(v string) *DescribeColumnsResponseBodyItems {
	s.SensLevelName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetSensitive(v bool) *DescribeColumnsResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetTableId(v int64) *DescribeColumnsResponseBodyItems {
	s.TableId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItems) SetTableName(v string) *DescribeColumnsResponseBodyItems {
	s.TableName = &v
	return s
}

type DescribeColumnsResponseBodyItemsModelTags struct {
	// The tag ID.
	//
	// 	- **101**: sensitive personal information
	//
	// 	- **102**: personal information
	//
	// 	- **103**: important information
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tag name.
	//
	// 	- Sensitive personal information
	//
	// 	- Personal information
	//
	// 	- Important information
	//
	// example:
	//
	// personal sensitive data
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeColumnsResponseBodyItemsModelTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItemsModelTags) SetId(v int64) *DescribeColumnsResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsModelTags) SetName(v string) *DescribeColumnsResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

type DescribeColumnsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeColumnsResponse) SetStatusCode(v int32) *DescribeColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColumnsResponse) SetBody(v *DescribeColumnsResponseBody) *DescribeColumnsResponse {
	s.Body = v
	return s
}

type DescribeColumnsV2Request struct {
	// When performing a paginated query, sets the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Engine type. Values:
	//
	// - **MySQL**.
	//
	// - **MariaDB**.
	//
	// - **Oracle**.
	//
	// - **PostgreSQL**.
	//
	// - **SQLServer**.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// ID of the asset instance to which the column data in the data asset table belongs.
	//
	// > Query the data in the columns of the data assets authorized by the Data Security Center based on the ID of the asset instance to which the column data in the data asset table belongs. The asset instance ID can be obtained by calling the [DescribeInstances](https://help.aliyun.com/document_detail/141708.html) interface.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Name of the asset instance to which the column data in the data asset table belongs.
	//
	// example:
	//
	// rm-bp17t1htja573l5i8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Sets the language type for requests and received messages, default is **zh_cn**.
	//
	// Values:
	//
	// - **zh_cn**: Simplified Chinese
	//
	// - **en_us**: English (United States)
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Search keyword, supports fuzzy matching.
	//
	// For example, entering **test*	- will search for all data information containing **test*	- in the search items.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// When performing a paginated query, sets the maximum number of data asset instances displayed per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Product name to which the column data in the data asset table belongs. Values: **MaxCompute, OSS, ADS, OTS, RDS**, etc.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Risk level ID of the sensitive data recognition rule. Values:
	//
	// - **1**: N/A.
	//
	// - **2**: S1.
	//
	// - **3**: S2.
	//
	// - **4**: S3.
	//
	// - **5**: S4.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// Unique identifier ID of the sensitive data recognition rule hit by the column data in the asset table.
	//
	// > Query the data in the columns of the data assets authorized by the Data Security Center based on the ID of the sensitive data recognition rule hit by the column data in the asset table. The sensitive data recognition rule ID can be obtained by calling the [DescribeRules](https://help.aliyun.com/document_detail/141389.html) interface.
	//
	// example:
	//
	// 11122200
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Name of the sensitive data recognition rule hit by the column data in the data asset table.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Sensitive level name. Values:
	//
	// - **N/A**: No sensitive data detected.
	//
	// - **S1**: Level 1 sensitive data.
	//
	// - **S2**: Level 2 sensitive data.
	//
	// - **S3**: Level 3 sensitive data.
	//
	// - **S4**: Level 4 sensitive data.
	//
	// example:
	//
	// S2
	SensLevelName *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	// Unique identifier ID of the asset table to which the column in MaxCompute, RDS, etc., belongs.
	//
	// > Query the data in the columns of the data assets authorized by the Data Security Center based on the ID of the asset table. The asset table ID can be obtained by calling the [DescribeTables](https://help.aliyun.com/document_detail/141709.html) interface.
	//
	// example:
	//
	// 11132334
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// Name of the data asset table.
	//
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeColumnsV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeColumnsV2Request) SetCurrentPage(v int32) *DescribeColumnsV2Request {
	s.CurrentPage = &v
	return s
}

func (s *DescribeColumnsV2Request) SetEngineType(v string) *DescribeColumnsV2Request {
	s.EngineType = &v
	return s
}

func (s *DescribeColumnsV2Request) SetInstanceId(v int64) *DescribeColumnsV2Request {
	s.InstanceId = &v
	return s
}

func (s *DescribeColumnsV2Request) SetInstanceName(v string) *DescribeColumnsV2Request {
	s.InstanceName = &v
	return s
}

func (s *DescribeColumnsV2Request) SetLang(v string) *DescribeColumnsV2Request {
	s.Lang = &v
	return s
}

func (s *DescribeColumnsV2Request) SetName(v string) *DescribeColumnsV2Request {
	s.Name = &v
	return s
}

func (s *DescribeColumnsV2Request) SetPageSize(v int32) *DescribeColumnsV2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeColumnsV2Request) SetProductCode(v string) *DescribeColumnsV2Request {
	s.ProductCode = &v
	return s
}

func (s *DescribeColumnsV2Request) SetRiskLevelId(v int64) *DescribeColumnsV2Request {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeColumnsV2Request) SetRuleId(v int64) *DescribeColumnsV2Request {
	s.RuleId = &v
	return s
}

func (s *DescribeColumnsV2Request) SetRuleName(v string) *DescribeColumnsV2Request {
	s.RuleName = &v
	return s
}

func (s *DescribeColumnsV2Request) SetSensLevelName(v string) *DescribeColumnsV2Request {
	s.SensLevelName = &v
	return s
}

func (s *DescribeColumnsV2Request) SetTableId(v string) *DescribeColumnsV2Request {
	s.TableId = &v
	return s
}

func (s *DescribeColumnsV2Request) SetTableName(v string) *DescribeColumnsV2Request {
	s.TableName = &v
	return s
}

type DescribeColumnsV2ResponseBody struct {
	// When performing a paginated query, sets the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of recognition results for the columns in the data table.
	Items []*DescribeColumnsV2ResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// When performing a paginated query, sets the maximum number of data asset instances displayed per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier generated by Alibaba Cloud for this request.
	//
	// example:
	//
	// B1F2BB1F-04EC-5D36-B136-B4DE17FD8DE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of data entries in the result.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeColumnsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColumnsV2ResponseBody) SetCurrentPage(v int32) *DescribeColumnsV2ResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeColumnsV2ResponseBody) SetItems(v []*DescribeColumnsV2ResponseBodyItems) *DescribeColumnsV2ResponseBody {
	s.Items = v
	return s
}

func (s *DescribeColumnsV2ResponseBody) SetPageSize(v int32) *DescribeColumnsV2ResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeColumnsV2ResponseBody) SetRequestId(v string) *DescribeColumnsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColumnsV2ResponseBody) SetTotalCount(v int32) *DescribeColumnsV2ResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeColumnsV2ResponseBodyItems struct {
	// The creation time of the column data in the data asset table, in milliseconds.
	//
	// example:
	//
	// 1536751124000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The data type of the column data in the data asset table.
	//
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Engine type. Values:
	//
	// - **MySQL**
	//
	// - **MariaDB**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The unique identifier ID of the column in the data asset table.
	//
	// example:
	//
	// 111111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the asset instance to which the column data in the data asset table belongs.
	//
	// example:
	//
	// 1232122
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset instance to which the column data in the data asset table belongs.
	//
	// example:
	//
	// rm-1234
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Column encryption status. Values:
	//
	// - **-1**: Not encrypted
	//
	// - **1**: Encryption successful
	//
	// - **2**: Encryption failed
	//
	// example:
	//
	// -1
	MaskingStatus *int32 `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	// Data tag list.
	ModelTags []*DescribeColumnsV2ResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// The name of the column in the data asset table.
	//
	// example:
	//
	// obj_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The risk level name of the asset. Values:
	//
	// - **N/A**：No sensitive data detected.
	//
	// - **S1**：Level 1 sensitive data.
	//
	// - **S2**：Level 2 sensitive data.
	//
	// - **S3**：Level 3 sensitive data.
	//
	// - **S4**：Level 4 sensitive data.
	//
	// example:
	//
	// S4
	OdpsRiskLevelName *string `json:"OdpsRiskLevelName,omitempty" xml:"OdpsRiskLevelName,omitempty"`
	// The risk level code of the asset. Values:
	//
	// - **1**：N/A.
	//
	// - **2**：S1.
	//
	// - **3**：S2.
	//
	// - **4**：S3.
	//
	// - **5**：S4.
	//
	// > A return value <= 1 represents N/A.
	//
	// example:
	//
	// 3
	OdpsRiskLevelValue *int32 `json:"OdpsRiskLevelValue,omitempty" xml:"OdpsRiskLevelValue,omitempty"`
	// The product name to which the column data in the data asset table belongs. Values: **MaxCompute, OSS, ADS, OTS, RDS**, etc.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID corresponding to the product name of the data asset. Values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADS
	//
	// - **4**: OTS
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region where the asset is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Correction record ID.
	//
	// example:
	//
	// 12
	RevisionId *int64 `json:"RevisionId,omitempty" xml:"RevisionId,omitempty"`
	// Correction status. Values:
	//
	// - 1: Corrected.
	//
	// - 0: Not corrected.
	//
	// example:
	//
	// 1
	RevisionStatus *int64 `json:"RevisionStatus,omitempty" xml:"RevisionStatus,omitempty"`
	// The risk level ID of the column data in the data asset table. Values:
	//
	// - **1**：N/A.
	//
	// - **2**：S1.
	//
	// - **3**：S2.
	//
	// - **4**：S3.
	//
	// - **5**：S4.
	//
	// > A return value <= 1 represents N/A.
	//
	// example:
	//
	// 4
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The risk level name of the column data in the data asset table. Values:
	//
	// - **N/A**：No sensitive data detected.
	//
	// - **S1**：Level 1 sensitive data.
	//
	// - **S2**：Level 2 sensitive data.
	//
	// - **S3**：Level 3 sensitive data.
	//
	// - **S4**：Level 4 sensitive data.
	//
	// example:
	//
	// S1
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The ID of the sensitive data recognition rule that the column data in the data asset table hits.
	//
	// example:
	//
	// 1004
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the sensitive data recognition rule that the column data in the data asset table hits.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The sensitivity level name. Values:
	//
	// - **N/A**：No sensitive data detected.
	//
	// - **S1**：Level 1 sensitive data.
	//
	// - **S2**：Level 2 sensitive data.
	//
	// - **S3**：Level 3 sensitive data.
	//
	// - **S4**：Level 4 sensitive data.
	//
	// example:
	//
	// S2
	SensLevelName *string `json:"SensLevelName,omitempty" xml:"SensLevelName,omitempty"`
	// Whether the column data in the data asset table contains sensitive data. Values:
	//
	// - true: The column data in the data asset table contains sensitive data.
	//
	// - false: The column data in the data asset table does not contain sensitive data.
	//
	// example:
	//
	// true
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The ID of the asset table to which the column data in the data asset table belongs.
	//
	// example:
	//
	// 123
	TableId *int64 `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the table to which the target column for correction belongs.
	//
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeColumnsV2ResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsV2ResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeColumnsV2ResponseBodyItems) SetCreationTime(v int64) *DescribeColumnsV2ResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetDataType(v string) *DescribeColumnsV2ResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetEngineType(v string) *DescribeColumnsV2ResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetId(v string) *DescribeColumnsV2ResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetInstanceId(v int64) *DescribeColumnsV2ResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetInstanceName(v string) *DescribeColumnsV2ResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetMaskingStatus(v int32) *DescribeColumnsV2ResponseBodyItems {
	s.MaskingStatus = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetModelTags(v []*DescribeColumnsV2ResponseBodyItemsModelTags) *DescribeColumnsV2ResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetName(v string) *DescribeColumnsV2ResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetOdpsRiskLevelName(v string) *DescribeColumnsV2ResponseBodyItems {
	s.OdpsRiskLevelName = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetOdpsRiskLevelValue(v int32) *DescribeColumnsV2ResponseBodyItems {
	s.OdpsRiskLevelValue = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetProductCode(v string) *DescribeColumnsV2ResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetProductId(v int64) *DescribeColumnsV2ResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetRegionId(v string) *DescribeColumnsV2ResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetRevisionId(v int64) *DescribeColumnsV2ResponseBodyItems {
	s.RevisionId = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetRevisionStatus(v int64) *DescribeColumnsV2ResponseBodyItems {
	s.RevisionStatus = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetRiskLevelId(v int64) *DescribeColumnsV2ResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetRiskLevelName(v string) *DescribeColumnsV2ResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetRuleId(v int64) *DescribeColumnsV2ResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetRuleName(v string) *DescribeColumnsV2ResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetSensLevelName(v string) *DescribeColumnsV2ResponseBodyItems {
	s.SensLevelName = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetSensitive(v bool) *DescribeColumnsV2ResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetTableId(v int64) *DescribeColumnsV2ResponseBodyItems {
	s.TableId = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItems) SetTableName(v string) *DescribeColumnsV2ResponseBodyItems {
	s.TableName = &v
	return s
}

type DescribeColumnsV2ResponseBodyItemsModelTags struct {
	// Data tag ID. Values:
	//
	// - **101**: Personal sensitive information
	//
	// - **102**: Personal information
	//
	// - **107**: General information
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Data tag name. Values:
	//
	// - Personal sensitive information
	//
	// - Personal information
	//
	// - General information
	//
	// example:
	//
	// personal sensitive information
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeColumnsV2ResponseBodyItemsModelTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsV2ResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *DescribeColumnsV2ResponseBodyItemsModelTags) SetId(v int64) *DescribeColumnsV2ResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *DescribeColumnsV2ResponseBodyItemsModelTags) SetName(v string) *DescribeColumnsV2ResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

type DescribeColumnsV2Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColumnsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColumnsV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeColumnsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeColumnsV2Response) SetHeaders(v map[string]*string) *DescribeColumnsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeColumnsV2Response) SetStatusCode(v int32) *DescribeColumnsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeColumnsV2Response) SetBody(v *DescribeColumnsV2ResponseBody) *DescribeColumnsV2Response {
	s.Body = v
	return s
}

type DescribeConfigsRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
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
	// An array that consists of common configuration items for alerts.
	ConfigList []*DescribeConfigsResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigsResponseBody) SetConfigList(v []*DescribeConfigsResponseBodyConfigList) *DescribeConfigsResponseBody {
	s.ConfigList = v
	return s
}

func (s *DescribeConfigsResponseBody) SetRequestId(v string) *DescribeConfigsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeConfigsResponseBodyConfigList struct {
	// The code of the common configuration item.
	//
	// example:
	//
	// 1
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the default value for the common configuration item.
	//
	// example:
	//
	// The volume of logs of a specific type that are generated on the current day is less than 30% of the average volume of logs generated in the previous 10 days.
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The description of the common configuration item.
	//
	// example:
	//
	// Anomalous log output
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the common configuration item.
	//
	// example:
	//
	// 2133
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The value of the common configuration item.
	//
	// example:
	//
	// 30
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeConfigsResponseBodyConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigsResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeConfigsResponseBodyConfigList) SetCode(v string) *DescribeConfigsResponseBodyConfigList {
	s.Code = &v
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

func (s *DescribeConfigsResponseBodyConfigList) SetValue(v string) *DescribeConfigsResponseBodyConfigList {
	s.Value = &v
	return s
}

type DescribeConfigsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeConfigsResponse) SetStatusCode(v int32) *DescribeConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigsResponse) SetBody(v *DescribeConfigsResponseBody) *DescribeConfigsResponse {
	s.Body = v
	return s
}

type DescribeDataAssetsRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The keyword that is used to search for data assets. Fuzzy search is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the data asset that you want to query. Valid values:
	//
	// 	- **1**: MaxCompute project
	//
	// 	- **2**: MaxCompute table
	//
	// 	- **3**: MaxCompute package
	//
	// 	- **11**: AnalyticDB for MySQL database
	//
	// 	- **12**: AnalyticDB for MySQL table
	//
	// 	- **21**: Object Storage Service (OSS) bucket
	//
	// 	- **22**: OSS object
	//
	// 	- **31**: Tablestore instance
	//
	// 	- **32**: Tablestore table
	//
	// 	- **51**: ApsaraDB RDS database
	//
	// 	- **52**: ApsaraDB RDS table
	//
	// 	- **61**: self-managed database hosted on an Elastic Compute Service (ECS) instance
	//
	// 	- **62**: self-managed table hosted on an ECS instance
	//
	// 	- **71**: PolarDB-X database
	//
	// 	- **72**: PolarDB-X table
	//
	// 	- **81**: PolarDB database
	//
	// 	- **82**: PolarDB table
	//
	// 	- **91**: AnalyticDB for PostgreSQL database
	//
	// 	- **92**: AnalyticDB for PostgreSQL table
	//
	// example:
	//
	// 1
	RangeId *int32 `json:"RangeId,omitempty" xml:"RangeId,omitempty"`
	// The sensitivity level of the data asset. Separate multiple sensitivity levels with commas (,). Valid values:
	//
	// 	- **2**: S1, indicating the low sensitivity level
	//
	// 	- **3**: S2, indicating the medium sensitivity level
	//
	// 	- **4**: S3, indicating the high sensitivity level
	//
	// 	- **5**: S4, indicating the highest sensitivity level
	//
	// example:
	//
	// 2
	RiskLevels *string `json:"RiskLevels,omitempty" xml:"RiskLevels,omitempty"`
	// The unique ID of the sensitive data detection rule that the data assets to be queried hit.
	//
	// > If you query sensitive data detection results based on the sensitive data detection rule that the data assets hit, you can call the [DescribeRules](~~DescribeRules~~) operation to query the ID of the sensitive data detection rule.
	//
	// example:
	//
	// 11122200
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeDataAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataAssetsRequest) SetCurrentPage(v int32) *DescribeDataAssetsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetLang(v string) *DescribeDataAssetsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetName(v string) *DescribeDataAssetsRequest {
	s.Name = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetPageSize(v int32) *DescribeDataAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataAssetsRequest) SetRangeId(v int32) *DescribeDataAssetsRequest {
	s.RangeId = &v
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

type DescribeDataAssetsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of data assets.
	Items []*DescribeDataAssetsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 71064826-726F-4ADA-B879-05D8055476FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried data assets that contain sensitive data.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeDataAssetsResponseBody) SetItems(v []*DescribeDataAssetsResponseBodyItems) *DescribeDataAssetsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataAssetsResponseBody) SetPageSize(v int32) *DescribeDataAssetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataAssetsResponseBody) SetRequestId(v string) *DescribeDataAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataAssetsResponseBody) SetTotalCount(v int32) *DescribeDataAssetsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataAssetsResponseBodyItems struct {
	// The access control list (ACL) that controls the access permissions on the OSS bucket.
	//
	// > This parameter is returned only when you set the parameter **RangeId*	- to **21**.
	//
	// example:
	//
	// private
	Acl *string `json:"Acl,omitempty" xml:"Acl,omitempty"`
	// The time when the data asset was created. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The data type of the data asset.
	//
	// example:
	//
	// OSS_BUCKET
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The ID of the data asset.
	//
	// example:
	//
	// 268
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The sensitivity tag of the data. The value is fixed as **0**. **0**, **1**, **2**, or **3*	- is returned for this parameter only when you set the parameter **RangeId*	- to **1**.
	//
	// 	- **0**: unclassified
	//
	// 	- **1**: confidential
	//
	// 	- **2**: sensitive
	//
	// 	- **3**: highly sensitive
	//
	// example:
	//
	// 0
	Labelsec *bool `json:"Labelsec,omitempty" xml:"Labelsec,omitempty"`
	// The name of the data asset.
	//
	// example:
	//
	// gxdata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The key value of the OSS object.
	//
	// > This parameter is returned only when you set the parameter **RangeId*	- to **22**.
	//
	// example:
	//
	// Internal
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The sensitivity level of the MaxCompute data asset. Valid values:
	//
	// 	- **S1**: low sensitivity level
	//
	// 	- **S2**: medium sensitivity level
	//
	// 	- **S3**: high sensitivity level
	//
	// 	- **S4**: highest sensitivity level
	//
	// > This parameter is returned only when you set the parameter **RangeId*	- to **1**.
	//
	// example:
	//
	// S4
	OdpsRiskLevelName *string `json:"OdpsRiskLevelName,omitempty" xml:"OdpsRiskLevelName,omitempty"`
	// The account that owns the data asset.
	//
	// example:
	//
	// dtdep-239-******
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the service to which the data asset belongs.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: OSS
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// example:
	//
	// 5
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// Indicates whether the data protection mechanism is enabled for the data asset. The value is fixed as **false**. **true*	- or **false*	- is returned for this parameter only when you set the parameter **RangeId*	- to **1**.
	//
	// 	- **false**: The data protection mechanism is disabled.
	//
	// 	- **true**: The data protection mechanism is enabled. Only data inbound is supported. Data outbound is not supported.
	//
	// example:
	//
	// false
	Protection *bool `json:"Protection,omitempty" xml:"Protection,omitempty"`
	// The sensitivity level of the data asset. A higher sensitivity level indicates that the identified data is more sensitive. Valid values:
	//
	// 	- **1**: No sensitive data is identified.
	//
	// 	- **2**: sensitive data at level 1.
	//
	// 	- **3**: sensitive data at level 2.
	//
	// 	- **3**: sensitive data at level 3.
	//
	// 	- **5**: sensitive data at level 4.
	//
	// 	- **6**: sensitive data at level 5.
	//
	// 	- **7**: sensitive data at level 6.
	//
	// 	- **8**: sensitive data at level 7.
	//
	// 	- **9**: sensitive data at level 8.
	//
	// 	- **10**: sensitive data at level 9.
	//
	// 	- **11**: sensitive data at level 10.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level for the data asset.
	//
	// example:
	//
	// Medium sensitivity level
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The name of the sensitive data detection rule that the data asset hits.
	//
	// example:
	//
	// \\*\\*\\	- rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Indicates whether the data asset contains sensitive data. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The total number of sensitive data assets. For example, the value can be the total number of sensitive MaxCompute projects, packages, or tables, the total number of sensitive ApsaraDB RDS databases or tables, or the total number of sensitive OSS buckets or objects.
	//
	// example:
	//
	// 24
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The percentage of sensitive data in all data assets.
	//
	// example:
	//
	// 45%
	SensitiveRatio *string `json:"SensitiveRatio,omitempty" xml:"SensitiveRatio,omitempty"`
	// The total number of data assets. For example, the value can be the total number of MaxCompute projects, packages, or tables, the total number of ApsaraDB RDS databases or tables, or the total number of OSS buckets or objects.
	//
	// example:
	//
	// 432
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeDataAssetsResponseBodyItems) SetId(v string) *DescribeDataAssetsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetLabelsec(v bool) *DescribeDataAssetsResponseBodyItems {
	s.Labelsec = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetName(v string) *DescribeDataAssetsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetObjectKey(v string) *DescribeDataAssetsResponseBodyItems {
	s.ObjectKey = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetOdpsRiskLevelName(v string) *DescribeDataAssetsResponseBodyItems {
	s.OdpsRiskLevelName = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetOwner(v string) *DescribeDataAssetsResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetProductCode(v string) *DescribeDataAssetsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetProductId(v string) *DescribeDataAssetsResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetProtection(v bool) *DescribeDataAssetsResponseBodyItems {
	s.Protection = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetRiskLevelId(v int64) *DescribeDataAssetsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetRiskLevelName(v string) *DescribeDataAssetsResponseBodyItems {
	s.RiskLevelName = &v
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

func (s *DescribeDataAssetsResponseBodyItems) SetSensitiveCount(v int32) *DescribeDataAssetsResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetSensitiveRatio(v string) *DescribeDataAssetsResponseBodyItems {
	s.SensitiveRatio = &v
	return s
}

func (s *DescribeDataAssetsResponseBodyItems) SetTotalCount(v int32) *DescribeDataAssetsResponseBodyItems {
	s.TotalCount = &v
	return s
}

type DescribeDataAssetsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDataAssetsResponse) SetStatusCode(v int32) *DescribeDataAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataAssetsResponse) SetBody(v *DescribeDataAssetsResponseBody) *DescribeDataAssetsResponse {
	s.Body = v
	return s
}

type DescribeDataLimitDetailRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The unique ID of the data asset that you want to query.
	//
	// > You can call the [DescribeDataLimits](~~DescribeDataLimits~~) operation to query the ID of the data asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12300
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Simplified Chinese.
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The network type of the data asset that you want to query. Valid values:
	//
	// 	- **1**: virtual private cloud (VPC)
	//
	// 	- **2**: classic network
	//
	// example:
	//
	// 1
	NetworkType *int32 `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s DescribeDataLimitDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitDetailRequest) SetFeatureType(v int32) *DescribeDataLimitDetailRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataLimitDetailRequest) SetId(v int64) *DescribeDataLimitDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitDetailRequest) SetLang(v string) *DescribeDataLimitDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataLimitDetailRequest) SetNetworkType(v int32) *DescribeDataLimitDetailRequest {
	s.NetworkType = &v
	return s
}

type DescribeDataLimitDetailResponseBody struct {
	// The details of the data asset.
	DataLimit *DescribeDataLimitDetailResponseBodyDataLimit `json:"DataLimit,omitempty" xml:"DataLimit,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataLimitDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitDetailResponseBody) SetDataLimit(v *DescribeDataLimitDetailResponseBodyDataLimit) *DescribeDataLimitDetailResponseBody {
	s.DataLimit = v
	return s
}

func (s *DescribeDataLimitDetailResponseBody) SetRequestId(v string) *DescribeDataLimitDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDataLimitDetailResponseBodyDataLimit struct {
	// The status of the connectivity test between the data asset and DSC. Valid values:
	//
	// 	- **2**: indicates that the data asset was being connected.
	//
	// 	- **3**: indicates that the data asset was connected to DSC.
	//
	// 	- **4**: indicates that the data asset failed to be connected.
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The result that indicates the status of the connectivity test between the data asset and DSC. Valid values:
	//
	// 	- **Passed**
	//
	// 	- **Failed**
	//
	// 	- **Testing**
	//
	// example:
	//
	// Passed
	CheckStatusName *string `json:"CheckStatusName,omitempty" xml:"CheckStatusName,omitempty"`
	// The time when the data asset was connected to DSC. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 145600000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the data asset.
	//
	// example:
	//
	// 12300
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region in which the data asset resides.
	//
	// example:
	//
	// China (Qingdao)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID and name of the data asset in the service to which the data asset belongs.
	//
	// example:
	//
	// rm-m5eup49p6o274****.RDS_example
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The port number that is used to connect to the database.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the data asset resides.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the service to which the data asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: OSS
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// example:
	//
	// 1
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The service to which the data asset belongs. Valid values:
	//
	// 	- **MaxCompute**
	//
	// 	- **OSS**
	//
	// 	- **ADS**
	//
	// 	- **OTS**
	//
	// 	- **RDS**
	//
	// example:
	//
	// RDS
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	// The account of the user who manages the data asset.
	//
	// example:
	//
	// User01
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDataLimitDetailResponseBodyDataLimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitDetailResponseBodyDataLimit) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetCheckStatus(v int32) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetCheckStatusName(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.CheckStatusName = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetGmtCreate(v int64) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetId(v int64) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetLocalName(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetParentId(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetPort(v int32) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.Port = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetRegionId(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetResourceType(v int64) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.ResourceType = &v
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

type DescribeDataLimitDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataLimitDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDataLimitDetailResponse) SetStatusCode(v int32) *DescribeDataLimitDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataLimitDetailResponse) SetBody(v *DescribeDataLimitDetailResponseBody) *DescribeDataLimitDetailResponse {
	s.Body = v
	return s
}

type DescribeDataLimitSetRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese (default)
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The parent asset ID of the data asset.
	//
	// You can call the [DescribeDataLimitDetail](~~DescribeDataLimitDetail~~) or [DescribeDataLimits](~~DescribeDataLimits~~) operation to obtain the parent asset ID of the data asset from the value of the **ParentId*	- parameter.
	//
	// example:
	//
	// db
	ParentId   *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	RegionType *string `json:"RegionType,omitempty" xml:"RegionType,omitempty"`
	// The type of service to which the data asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: OSS
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// example:
	//
	// 2
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeDataLimitSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetRequest) SetFeatureType(v int32) *DescribeDataLimitSetRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataLimitSetRequest) SetLang(v string) *DescribeDataLimitSetRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataLimitSetRequest) SetParentId(v string) *DescribeDataLimitSetRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitSetRequest) SetRegionType(v string) *DescribeDataLimitSetRequest {
	s.RegionType = &v
	return s
}

func (s *DescribeDataLimitSetRequest) SetResourceType(v int32) *DescribeDataLimitSetRequest {
	s.ResourceType = &v
	return s
}

type DescribeDataLimitSetResponseBody struct {
	// The information about the data asset.
	DataLimitSet *DescribeDataLimitSetResponseBodyDataLimitSet `json:"DataLimitSet,omitempty" xml:"DataLimitSet,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataLimitSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBody) SetDataLimitSet(v *DescribeDataLimitSetResponseBodyDataLimitSet) *DescribeDataLimitSetResponseBody {
	s.DataLimitSet = v
	return s
}

func (s *DescribeDataLimitSetResponseBody) SetRequestId(v string) *DescribeDataLimitSetResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDataLimitSetResponseBodyDataLimitSet struct {
	// An array that consists of data assets that DSC is authorized to scan.
	DataLimitList []*DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList `json:"DataLimitList,omitempty" xml:"DataLimitList,omitempty" type:"Repeated"`
	// An array consisting of the OSS objects that DSC is authorized to scan.
	OssBucketList []*DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList `json:"OssBucketList,omitempty" xml:"OssBucketList,omitempty" type:"Repeated"`
	// An array consisting of the regions in which the data assets can be scanned.
	RegionList []*DescribeDataLimitSetResponseBodyDataLimitSetRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The type of service to which the data asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: OSS
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// example:
	//
	// 2
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The service to which the data asset belongs. Valid values:
	//
	// 	- **ODPS**
	//
	// 	- **OSS**
	//
	// 	- **ADS**
	//
	// 	- **OTS**
	//
	// 	- **RDS**
	//
	// example:
	//
	// OSS
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	// The total number of data objects in the data assets.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSet) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSet) GoString() string {
	return s.String()
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

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetResourceType(v int64) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetResourceTypeCode(v string) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetTotalCount(v int32) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.TotalCount = &v
	return s
}

type DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList struct {
	// Indicates whether the test of connectivity between DSC and the data asset is passed.
	//
	// 	- **2**: The connectivity test is in progress.
	//
	// 	- **3**: The connectivity test is passed.
	//
	// 	- **4**: The connectivity test failed.
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The name of the data detection status.
	//
	// example:
	//
	// Connectivity test status
	CheckStatusName *string `json:"CheckStatusName,omitempty" xml:"CheckStatusName,omitempty"`
	// The connection string that is used to access the data asset.
	//
	// example:
	//
	// Connection string
	Connector *string `json:"Connector,omitempty" xml:"Connector,omitempty"`
	// The time when the data asset was created. Unit: milliseconds.
	//
	// example:
	//
	// 1625587423000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the data asset.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region in which the data asset resides.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The parent asset ID of the data asset.
	//
	// example:
	//
	// db
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The region in which the data asset resides.
	//
	// example:
	//
	// cn-****
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of service to which the data asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: OSS
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// example:
	//
	// 2
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The code of the service to which the data asset belongs. Valid values:
	//
	// 	- **ODPS**
	//
	// 	- **OSS**
	//
	// 	- **ADS**
	//
	// 	- **OTS**
	//
	// 	- **RDS**
	//
	// example:
	//
	// OSS
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	// The username that is used to access the data asset.
	//
	// example:
	//
	// tsts
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetCheckStatus(v int32) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetCheckStatusName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.CheckStatusName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetConnector(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.Connector = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetGmtCreate(v int64) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetId(v int64) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetLocalName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetParentId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetRegionId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetResourceType(v int64) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.ResourceType = &v
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

type DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList struct {
	// The name of the OSS bucket to which the OSS object belongs.
	//
	// example:
	//
	// oss-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The region ID of the OSS object.
	//
	// example:
	//
	// cn-****
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) SetBucketName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList {
	s.BucketName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) SetRegionId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList {
	s.RegionId = &v
	return s
}

type DescribeDataLimitSetResponseBodyDataLimitSetRegionList struct {
	// The name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-****
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetRegionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetRegionList) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetRegionList) SetLocalName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetRegionList {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetRegionList) SetRegionId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetRegionList {
	s.RegionId = &v
	return s
}

type DescribeDataLimitSetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataLimitSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDataLimitSetResponse) SetStatusCode(v int32) *DescribeDataLimitSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataLimitSetResponse) SetBody(v *DescribeDataLimitSetResponseBody) *DescribeDataLimitSetResponse {
	s.Body = v
	return s
}

type DescribeDataLimitsRequest struct {
	// Specifies whether to enable the security audit feature. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The data detection status. Valid values:
	//
	// 	- **0**: The data detection is ready.
	//
	// 	- **1**: The data detection is running.
	//
	// 	- **2**: The connectivity test is in progress.
	//
	// 	- **3**: The connectivity test passed.
	//
	// 	- **4**: The connectivity test failed.
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether DSC has the data de-identification permissions on the data asset. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	DatamaskStatus *int32 `json:"DatamaskStatus,omitempty" xml:"DatamaskStatus,omitempty"`
	// Specifies whether DSC has the data detection permissions on the data asset. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The end of the time range to query The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1616068534877
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the database engine. Valid values include **MySQL**, **SQLServer**, **Oracle**, **PostgreSQL**, and **MongoDB**.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the member.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The parent ID of the data asset to be queried. Valid values:
	//
	// 	- The name or ID of the MaxCompute project.
	//
	// 	- The name or ID of the OSS bucket.
	//
	// 	- The name or ID of the ApsaraDB RDS instance or database.
	//
	// example:
	//
	// 1112
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The type of the service to which the data asset belongs. This parameter is required. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: Object Storage Service (OSS)
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// 	- **6**: self-managed database
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The region in which the data asset resides.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The beginning of the time range to query The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1616068534877
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataLimitsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsRequest) SetAuditStatus(v int32) *DescribeDataLimitsRequest {
	s.AuditStatus = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetCheckStatus(v int32) *DescribeDataLimitsRequest {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetCurrentPage(v int32) *DescribeDataLimitsRequest {
	s.CurrentPage = &v
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

func (s *DescribeDataLimitsRequest) SetEndTime(v int64) *DescribeDataLimitsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetEngineType(v string) *DescribeDataLimitsRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetFeatureType(v int32) *DescribeDataLimitsRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetLang(v string) *DescribeDataLimitsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetMemberAccount(v int64) *DescribeDataLimitsRequest {
	s.MemberAccount = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetPageSize(v int32) *DescribeDataLimitsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetParentId(v string) *DescribeDataLimitsRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetResourceType(v int32) *DescribeDataLimitsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetServiceRegionId(v string) *DescribeDataLimitsRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeDataLimitsRequest) SetStartTime(v int64) *DescribeDataLimitsRequest {
	s.StartTime = &v
	return s
}

type DescribeDataLimitsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data assets.
	Items []*DescribeDataLimitsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeDataLimitsResponseBody) SetItems(v []*DescribeDataLimitsResponseBodyItems) *DescribeDataLimitsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataLimitsResponseBody) SetPageSize(v int32) *DescribeDataLimitsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataLimitsResponseBody) SetRequestId(v string) *DescribeDataLimitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataLimitsResponseBody) SetTotalCount(v int32) *DescribeDataLimitsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataLimitsResponseBodyItems struct {
	// Indicates whether the security audit feature is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Indicates whether the data asset can be automatically scanned. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AutoScan *int32 `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	// The data detection status. Valid values:
	//
	// 	- **0**: The data detection is ready.
	//
	// 	- **1**: The data detection is running.
	//
	// 	- **2**: The connectivity test is in progress.
	//
	// 	- **3**: The connectivity test is passed.
	//
	// 	- **4**: The connectivity test failed.
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The name of the data detection status.
	//
	// example:
	//
	// Connectivity test status
	CheckStatusName *string `json:"CheckStatusName,omitempty" xml:"CheckStatusName,omitempty"`
	// Indicates whether DSC has the data de-identification permissions on the data asset. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	DatamaskStatus *int32 `json:"DatamaskStatus,omitempty" xml:"DatamaskStatus,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 2.0
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// Indicates whether DSC has the data identification permissions on the data asset. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of the database engine. Valid values include **MySQL**, **SQLServer**, **Oracle**, **PostgreSQL**, and **MongoDB**.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The error code.
	//
	// example:
	//
	// connect_network_error
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The reason for the failure.
	//
	// example:
	//
	// The password is invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the data leak prevention feature is enabled. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes (default)
	//
	// example:
	//
	// 1
	EventStatus *int32 `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The time when the data asset was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 145600000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The unique ID of the data asset.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// 123
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The ID of the data asset to which the table belongs.
	//
	// example:
	//
	// 12332
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the last scan is performed.
	//
	// 	- The value is a UNIX timestamp.
	//
	// 	- Unit: milliseconds.
	//
	// example:
	//
	// 145600000
	LastFinishedTime *int64 `json:"LastFinishedTime,omitempty" xml:"LastFinishedTime,omitempty"`
	// The last scan start time of data assets, in milliseconds.
	//
	// example:
	//
	// 145600000
	LastStartTime *int64 `json:"LastStartTime,omitempty" xml:"LastStartTime,omitempty"`
	// The region in which the data asset resides.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The retention period of raw logs. Unit: days.
	//
	// example:
	//
	// 30
	LogStoreDay *int32 `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	// The ID of the member.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// The next time when the data asset is scanned. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1676620236000
	NextStartTime *int64 `json:"NextStartTime,omitempty" xml:"NextStartTime,omitempty"`
	// Indicates whether the optical character recognition (OCR) feature is enabled. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	OcrStatus *int32 `json:"OcrStatus,omitempty" xml:"OcrStatus,omitempty"`
	// The parent ID of the data asset that you want to query. Valid values include **bucket, db, and project**.
	//
	// example:
	//
	// project
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The port number of the self-managed database.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The status of the data asset scan. Valid values:
	//
	// 	- **-1**: invalid
	//
	// 	- **0**: waiting
	//
	// 	- **1**: being scanned
	//
	// 	- **2**: suspended
	//
	// 	- **3**: completed
	//
	// example:
	//
	// 3
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// The total number of data tables or files.
	//
	// example:
	//
	// 100
	ProcessTotalCount *int32 `json:"ProcessTotalCount,omitempty" xml:"ProcessTotalCount,omitempty"`
	// The region in which the asset resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the service to which the data asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: OSS
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// 	- **6**: self-managed database
	//
	// example:
	//
	// 5
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The code of the service to which the data asset belongs. Valid values: **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// RDS
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	// The number of sensitive data samples. Valid values: **0**, **5**, and **10**. Unit: data entries.
	//
	// example:
	//
	// 5
	SamplingSize *int32 `json:"SamplingSize,omitempty" xml:"SamplingSize,omitempty"`
	// A list of the IDs of the security groups that are used by PrivateLink when you install the DSC agent.
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" xml:"SecurityGroupIdList,omitempty" type:"Repeated"`
	// Indicates whether the security audit feature is supported. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	SupportAudit *bool `json:"SupportAudit,omitempty" xml:"SupportAudit,omitempty"`
	// Indicates whether the data de-identification feature is supported. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	SupportDatamask *bool `json:"SupportDatamask,omitempty" xml:"SupportDatamask,omitempty"`
	// Indicates whether anomalous event detection is supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SupportEvent *bool `json:"SupportEvent,omitempty" xml:"SupportEvent,omitempty"`
	// Indicates whether OCR is supported. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	SupportOcr *bool `json:"SupportOcr,omitempty" xml:"SupportOcr,omitempty"`
	// Indicates whether the data asset scan feature is supported. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	SupportScan *bool `json:"SupportScan,omitempty" xml:"SupportScan,omitempty"`
	// The alias of the tenant.
	//
	// example:
	//
	// insta_gram
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The total number of fields in the table.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The username that is used to access the data asset.
	//
	// example:
	//
	// tsts
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// A list of the IDs of the vSwitches that are used by PrivateLink when you install the DSC agent.
	VSwitchIdList []*string `json:"VSwitchIdList,omitempty" xml:"VSwitchIdList,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the data asset belongs.
	//
	// example:
	//
	// vpc-2zevcqke6hh09c41****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDataLimitsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataLimitsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsResponseBodyItems) SetAuditStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.AuditStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetAutoScan(v int32) *DescribeDataLimitsResponseBodyItems {
	s.AutoScan = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetCheckStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetCheckStatusName(v string) *DescribeDataLimitsResponseBodyItems {
	s.CheckStatusName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetDatamaskStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.DatamaskStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetDbVersion(v string) *DescribeDataLimitsResponseBodyItems {
	s.DbVersion = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetEnable(v int32) *DescribeDataLimitsResponseBodyItems {
	s.Enable = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetEngineType(v string) *DescribeDataLimitsResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetErrorCode(v string) *DescribeDataLimitsResponseBodyItems {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetErrorMessage(v string) *DescribeDataLimitsResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetEventStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.EventStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetGmtCreate(v int64) *DescribeDataLimitsResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetId(v int64) *DescribeDataLimitsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetInstanceDescription(v string) *DescribeDataLimitsResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetInstanceId(v string) *DescribeDataLimitsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetLastFinishedTime(v int64) *DescribeDataLimitsResponseBodyItems {
	s.LastFinishedTime = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetLastStartTime(v int64) *DescribeDataLimitsResponseBodyItems {
	s.LastStartTime = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetLocalName(v string) *DescribeDataLimitsResponseBodyItems {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetLogStoreDay(v int32) *DescribeDataLimitsResponseBodyItems {
	s.LogStoreDay = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetMemberAccount(v int64) *DescribeDataLimitsResponseBodyItems {
	s.MemberAccount = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetNextStartTime(v int64) *DescribeDataLimitsResponseBodyItems {
	s.NextStartTime = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetOcrStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.OcrStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetParentId(v string) *DescribeDataLimitsResponseBodyItems {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetPort(v int32) *DescribeDataLimitsResponseBodyItems {
	s.Port = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetProcessStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetProcessTotalCount(v int32) *DescribeDataLimitsResponseBodyItems {
	s.ProcessTotalCount = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetRegionId(v string) *DescribeDataLimitsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetResourceType(v int64) *DescribeDataLimitsResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetResourceTypeCode(v string) *DescribeDataLimitsResponseBodyItems {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSamplingSize(v int32) *DescribeDataLimitsResponseBodyItems {
	s.SamplingSize = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSecurityGroupIdList(v []*string) *DescribeDataLimitsResponseBodyItems {
	s.SecurityGroupIdList = v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportAudit(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportAudit = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportDatamask(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportDatamask = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportEvent(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportEvent = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportOcr(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportOcr = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportScan(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportScan = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetTenantName(v string) *DescribeDataLimitsResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetTotalCount(v int32) *DescribeDataLimitsResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetUserName(v string) *DescribeDataLimitsResponseBodyItems {
	s.UserName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetVSwitchIdList(v []*string) *DescribeDataLimitsResponseBodyItems {
	s.VSwitchIdList = v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetVpcId(v string) *DescribeDataLimitsResponseBodyItems {
	s.VpcId = &v
	return s
}

type DescribeDataLimitsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataLimitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDataLimitsResponse) SetStatusCode(v int32) *DescribeDataLimitsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataLimitsResponse) SetBody(v *DescribeDataLimitsResponseBody) *DescribeDataLimitsResponse {
	s.Body = v
	return s
}

type DescribeDataMaskingRunHistoryRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the service to which the de-identified data belongs. Valid values: **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates OSS. The value indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 2
	DstType *int32 `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1583856000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the task.
	//
	// > If a task has one or more subtasks, the value of the parameter must be the ID of the task. Otherwise, leave this parameter empty.
	//
	// example:
	//
	// 366731
	MainProcessId *int64 `json:"MainProcessId,omitempty" xml:"MainProcessId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the source table.
	//
	// example:
	//
	// add
	SrcTableName *string `json:"SrcTableName,omitempty" xml:"SrcTableName,omitempty"`
	// The type of the service to which the data to be de-identified belongs. Valid values: **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 2
	SrcType *int32 `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1582992000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the de-identification task. Valid values:
	//
	// 	- **-1**: waiting
	//
	// 	- **0**: being executed
	//
	// 	- **1**: executed
	//
	// 	- **2**: failed to be executed
	//
	// 	- **3**: terminated
	//
	// 	- **4**: partially failed
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the de-identification task.
	//
	// example:
	//
	// mt4HBgtw1B******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDataMaskingRunHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingRunHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingRunHistoryRequest) SetCurrentPage(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetDstType(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.DstType = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetEndTime(v int64) *DescribeDataMaskingRunHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetLang(v string) *DescribeDataMaskingRunHistoryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetMainProcessId(v int64) *DescribeDataMaskingRunHistoryRequest {
	s.MainProcessId = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetPageSize(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetSrcTableName(v string) *DescribeDataMaskingRunHistoryRequest {
	s.SrcTableName = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetSrcType(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.SrcType = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetStartTime(v int64) *DescribeDataMaskingRunHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetStatus(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.Status = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetTaskId(v string) *DescribeDataMaskingRunHistoryRequest {
	s.TaskId = &v
	return s
}

type DescribeDataMaskingRunHistoryResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The execution information about the de-identification task.
	Items []*DescribeDataMaskingRunHistoryResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-4******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeDataMaskingRunHistoryResponseBody) SetItems(v []*DescribeDataMaskingRunHistoryResponseBodyItems) *DescribeDataMaskingRunHistoryResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetPageSize(v int32) *DescribeDataMaskingRunHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetRequestId(v string) *DescribeDataMaskingRunHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetTotalCount(v int32) *DescribeDataMaskingRunHistoryResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataMaskingRunHistoryResponseBodyItems struct {
	// The number of rows that are in conflict with the data to be de-identified in the destination table to which the data to be de-identified is moved.
	//
	// example:
	//
	// 0
	ConflictCount *int64 `json:"ConflictCount,omitempty" xml:"ConflictCount,omitempty"`
	// The type of the service to which the de-identified data belongs. Valid values: **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates OSS. The value indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 2
	DstType *int32 `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The service that stores the de-identified data. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// OSS
	DstTypeCode *string `json:"DstTypeCode,omitempty" xml:"DstTypeCode,omitempty"`
	// The end time of the de-identification task.
	//
	// example:
	//
	// 1582251233000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error code that is returned when the de-identification task fails.
	//
	// example:
	//
	// masking_task_not_found
	FailCode *string `json:"FailCode,omitempty" xml:"FailCode,omitempty"`
	// The reason why the de-identification task fails.
	//
	// example:
	//
	// error
	FailMsg *string `json:"FailMsg,omitempty" xml:"FailMsg,omitempty"`
	// Indicates whether a file is available for download.
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	HasDownloadFile *int32 `json:"HasDownloadFile,omitempty" xml:"HasDownloadFile,omitempty"`
	// The number of created subtasks.
	//
	// example:
	//
	// 4
	HasSubProcess *int32 `json:"HasSubProcess,omitempty" xml:"HasSubProcess,omitempty"`
	// The ID of the task execution record.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of rows that are de-identified.
	//
	// example:
	//
	// 100
	MaskingCount *int64 `json:"MaskingCount,omitempty" xml:"MaskingCount,omitempty"`
	// The progress of the de-identification task.
	//
	// example:
	//
	// 100
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The number of times that the de-identification task is executed.
	//
	// example:
	//
	// 1
	RunIndex *int32 `json:"RunIndex,omitempty" xml:"RunIndex,omitempty"`
	// The name of the source table.
	//
	// example:
	//
	// add
	SrcTableName *string `json:"SrcTableName,omitempty" xml:"SrcTableName,omitempty"`
	// The type of the service to which the data to be de-identified belongs. Valid values: **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates OSS. The value indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 2
	SrcType *int32 `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	// The service to which the data to be de-identified belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// OSS
	SrcTypeCode *string `json:"SrcTypeCode,omitempty" xml:"SrcTypeCode,omitempty"`
	// The time when the de-identification task was executed. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1582251233000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the de-identification task. Valid values:
	//
	// 	- **-1**: waiting
	//
	// 	- **0**: being executed
	//
	// 	- **1**: executed
	//
	// 	- **2**: failed to be executed
	//
	// 	- **3**: terminated
	//
	// 	- **4**: partially failed
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the identification task.
	//
	// example:
	//
	// mt4HBgtw1B******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The mode in which the de-identification task is executed. Valid values:
	//
	// 	- **1**: manual
	//
	// 	- **2**: scheduled
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDataMaskingRunHistoryResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingRunHistoryResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetConflictCount(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.ConflictCount = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetDstType(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.DstType = &v
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

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetFailCode(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.FailCode = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetFailMsg(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.FailMsg = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetHasDownloadFile(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.HasDownloadFile = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetHasSubProcess(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.HasSubProcess = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetId(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.Id = &v
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

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetRunIndex(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.RunIndex = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetSrcTableName(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.SrcTableName = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetSrcType(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.SrcType = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetSrcTypeCode(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.SrcTypeCode = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetStartTime(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetStatus(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetTaskId(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetType(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.Type = &v
	return s
}

type DescribeDataMaskingRunHistoryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataMaskingRunHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDataMaskingRunHistoryResponse) SetStatusCode(v int32) *DescribeDataMaskingRunHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponse) SetBody(v *DescribeDataMaskingRunHistoryResponseBody) *DescribeDataMaskingRunHistoryResponse {
	s.Body = v
	return s
}

type DescribeDataMaskingTasksRequest struct {
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The service to which the data to be de-identified belongs. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 2
	DstType *int32 `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The end of the time range during which the de-identification tasks to be queried are created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1583856000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword used to query the de-identification tasks, which can be the task name or ID.
	//
	// example:
	//
	// test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The beginning of the time range during which the de-identification tasks to be queried are created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1582992000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataMaskingTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingTasksRequest) SetCurrentPage(v int32) *DescribeDataMaskingTasksRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetDstType(v int32) *DescribeDataMaskingTasksRequest {
	s.DstType = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetEndTime(v int64) *DescribeDataMaskingTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetLang(v string) *DescribeDataMaskingTasksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetPageSize(v int32) *DescribeDataMaskingTasksRequest {
	s.PageSize = &v
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

type DescribeDataMaskingTasksResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// A list of de-identification tasks.
	Items []*DescribeDataMaskingTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-4******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeDataMaskingTasksResponseBody) SetItems(v []*DescribeDataMaskingTasksResponseBodyItems) *DescribeDataMaskingTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) SetPageSize(v int32) *DescribeDataMaskingTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) SetRequestId(v string) *DescribeDataMaskingTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) SetTotalCount(v int32) *DescribeDataMaskingTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataMaskingTasksResponseBodyItems struct {
	// The member account to which the desensitization target belongs.
	//
	// example:
	//
	// 192479427903xxxx
	DstMemberAccount *int64 `json:"DstMemberAccount,omitempty" xml:"DstMemberAccount,omitempty"`
	// The destination path.
	DstPath *string `json:"DstPath,omitempty" xml:"DstPath,omitempty"`
	// The service to which the data to be de-identified belongs. Valid values: **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates OSS. The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 5
	DstType *int32 `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The type of the service to which the de-identified data belongs. Valid values: **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// RDS
	DstTypeCode *string `json:"DstTypeCode,omitempty" xml:"DstTypeCode,omitempty"`
	// The time when the de-identification task is created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1582992000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Indicates whether the de-identification task is running.
	//
	// example:
	//
	// false
	HasUnfinishProcess *bool `json:"HasUnfinishProcess,omitempty" xml:"HasUnfinishProcess,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the source table is de-identified.
	//
	// example:
	//
	// false
	OriginalTable *bool `json:"OriginalTable,omitempty" xml:"OriginalTable,omitempty"`
	// The user who created the de-identification task.
	//
	// example:
	//
	// owner
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The number of times that the de-identification task is run.
	//
	// example:
	//
	// 1
	RunCount *int32 `json:"RunCount,omitempty" xml:"RunCount,omitempty"`
	// The member account to which the desensitization source belongs.
	//
	// example:
	//
	// 192479427903xxxx
	SrcMemberAccount *int64 `json:"SrcMemberAccount,omitempty" xml:"SrcMemberAccount,omitempty"`
	// The source path.
	SrcPath *string `json:"SrcPath,omitempty" xml:"SrcPath,omitempty"`
	// The type of the service to which the data to be de-identified belongs. Valid values: **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates OSS. The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 5
	SrcType *int32 `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	// The type of the service to which the data to be de-identified belongs. Valid values: **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// RDS
	SrcTypeCode *string `json:"SrcTypeCode,omitempty" xml:"SrcTypeCode,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// mt4HBgtw1B******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// Task name
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The mode in which the de-identification task is run. Valid values:
	//
	// 	- **1**: manual
	//
	// 	- **2**: scheduled
	//
	// 	- **3**: manual and scheduled
	//
	// example:
	//
	// 1
	TriggerType *int32 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s DescribeDataMaskingTasksResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataMaskingTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetDstMemberAccount(v int64) *DescribeDataMaskingTasksResponseBodyItems {
	s.DstMemberAccount = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetDstPath(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.DstPath = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetDstType(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.DstType = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetDstTypeCode(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.DstTypeCode = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetGmtCreate(v int64) *DescribeDataMaskingTasksResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetHasUnfinishProcess(v bool) *DescribeDataMaskingTasksResponseBodyItems {
	s.HasUnfinishProcess = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetId(v int64) *DescribeDataMaskingTasksResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetOriginalTable(v bool) *DescribeDataMaskingTasksResponseBodyItems {
	s.OriginalTable = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetOwner(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetRunCount(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.RunCount = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetSrcMemberAccount(v int64) *DescribeDataMaskingTasksResponseBodyItems {
	s.SrcMemberAccount = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetSrcPath(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.SrcPath = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetSrcType(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.SrcType = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetSrcTypeCode(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.SrcTypeCode = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetStatus(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetTaskId(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetTaskName(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.TaskName = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetTriggerType(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.TriggerType = &v
	return s
}

type DescribeDataMaskingTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataMaskingTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDataMaskingTasksResponse) SetStatusCode(v int32) *DescribeDataMaskingTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataMaskingTasksResponse) SetBody(v *DescribeDataMaskingTasksResponseBody) *DescribeDataMaskingTasksResponse {
	s.Body = v
	return s
}

type DescribeDataObjectColumnDetailRequest struct {
	// When performing a paginated query, set the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// Set the unique identifier ID of the data object to be queried.
	//
	// > You can obtain the identifier ID by calling [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html).
	//
	// example:
	//
	// 318248
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language type for requests and responses. Default value: **zh_cn**. Values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// When performing a paginated query, set the maximum number of data asset instances displayed per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID corresponding to the product name of the data object. Values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: TableStore
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// - **7**: PolarDB-X
	//
	// - **8**: PolarDB
	//
	// - **9**: ADB-PG
	//
	// - **10**: OceanBase
	//
	// - **11**: MongoDB
	//
	// - **25**: Redis
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// Industry template ID.
	//
	// > You can obtain the industry template identifier ID by calling [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html).
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDataObjectColumnDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectColumnDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailRequest) SetCurrentPage(v int32) *DescribeDataObjectColumnDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetFeatureType(v int32) *DescribeDataObjectColumnDetailRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetId(v int64) *DescribeDataObjectColumnDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetLang(v string) *DescribeDataObjectColumnDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetPageSize(v int32) *DescribeDataObjectColumnDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetProductId(v int64) *DescribeDataObjectColumnDetailRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetTemplateId(v int64) *DescribeDataObjectColumnDetailRequest {
	s.TemplateId = &v
	return s
}

type DescribeDataObjectColumnDetailResponseBody struct {
	// When performing a paginated query, set the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of identified results for the columns of the data table.
	Items []*DescribeDataObjectColumnDetailResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// When performing a paginated query, set the maximum number of data asset instances displayed per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of this request, which is a unique identifier generated by Alibaba Cloud for the request. It can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 8C8036CC-961D-514E-88E8-3088B5A50CA9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of data items in the result.
	//
	// example:
	//
	// 61
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataObjectColumnDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectColumnDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailResponseBody) SetCurrentPage(v int32) *DescribeDataObjectColumnDetailResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBody) SetItems(v []*DescribeDataObjectColumnDetailResponseBodyItems) *DescribeDataObjectColumnDetailResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBody) SetPageSize(v int32) *DescribeDataObjectColumnDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBody) SetRequestId(v string) *DescribeDataObjectColumnDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBody) SetTotalCount(v int32) *DescribeDataObjectColumnDetailResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataObjectColumnDetailResponseBodyItems struct {
	// List of industry categories for the sensitive data.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// Comment on the column.
	//
	// example:
	//
	// column comment
	ColumnComment *string `json:"ColumnComment,omitempty" xml:"ColumnComment,omitempty"`
	// Column name.
	//
	// example:
	//
	// hide14
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// Data type of the column.
	//
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Engine type. Possible values:
	//
	// - **MySQL**
	//
	// - **MariaDB**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// Unique identifier ID of the column object.
	//
	// example:
	//
	// 1509415150052786176
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Instance name of the data asset table.
	//
	// example:
	//
	// rm-1234
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Column encryption status. Possible values:
	//
	// - **-1**：Not encrypted
	//
	// - **1**：Encryption successful
	//
	// - **2**：Encryption failed
	//
	// example:
	//
	// -1
	MaskingStatus *int32 `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	// List of data tags.
	ModelTags []*DescribeDataObjectColumnDetailResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// Whether the column is a primary key. Values:
	//
	// - **true**: Primary key.
	//
	// - **false**: Not a primary key.
	//
	// example:
	//
	// true
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// ID corresponding to the product name of the data object. Possible values:
	//
	// - **1**：MaxCompute
	//
	// - **2**：OSS
	//
	// - **3**：ADB-MYSQL
	//
	// - **4**：TableStore
	//
	// - **5**：RDS
	//
	// - **6**：SELF_DB
	//
	// - **7**：PolarDB-X
	//
	// - **8**：PolarDB
	//
	// - **9**：ADB-PG
	//
	// - **10**：OceanBase
	//
	// - **11**：MongoDB
	//
	// - **25**：Redis
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// Region where the asset is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Risk level ID. Values:
	//
	// - **1**: N/A: No sensitive data detected.
	//
	// - **2**: S1: Level 1 sensitive data.
	//
	// - **3**: S2: Level 2 sensitive data.
	//
	// - **4**: S3: Level 3 sensitive data.
	//
	// - **5**: S4: Level 4 sensitive data.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// Risk level name. Possible values:
	//
	// - **N/A**：No sensitive data detected.
	//
	// - **S1**：Level 1 sensitive data.
	//
	// - **S2**：Level 2 sensitive data.
	//
	// - **S3**：Level 3 sensitive data.
	//
	// - **S4**：Level 4 sensitive data.
	//
	// example:
	//
	// S1
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// ID of the matched identification model.
	//
	// example:
	//
	// 1004
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Name of the matched identification model.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Table name.
	//
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeDataObjectColumnDetailResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectColumnDetailResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetCategories(v []*string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.Categories = v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetColumnComment(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.ColumnComment = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetColumnName(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.ColumnName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetDataType(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetEngineType(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetId(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetInstanceName(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetMaskingStatus(v int32) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.MaskingStatus = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetModelTags(v []*DescribeDataObjectColumnDetailResponseBodyItemsModelTags) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetPrimaryKey(v bool) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetProductId(v int64) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetRegionId(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetRiskLevelId(v int64) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetRiskLevelName(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetRuleId(v int64) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetRuleName(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItems) SetTableName(v string) *DescribeDataObjectColumnDetailResponseBodyItems {
	s.TableName = &v
	return s
}

type DescribeDataObjectColumnDetailResponseBodyItemsModelTags struct {
	// ID corresponding to the data tag name. Possible values:
	//
	// - **101**：Personal sensitive information
	//
	// - **102**：Personal information
	//
	// - **107**：General information
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Data tag name. Possible values:
	//
	// - **101**：Personal sensitive information
	//
	// - **102**：Personal information
	//
	// - **107**：General information
	//
	// example:
	//
	// personal sensitive information
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDataObjectColumnDetailResponseBodyItemsModelTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectColumnDetailResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailResponseBodyItemsModelTags) SetId(v int64) *DescribeDataObjectColumnDetailResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponseBodyItemsModelTags) SetName(v string) *DescribeDataObjectColumnDetailResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

type DescribeDataObjectColumnDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataObjectColumnDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataObjectColumnDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectColumnDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailResponse) SetHeaders(v map[string]*string) *DescribeDataObjectColumnDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataObjectColumnDetailResponse) SetStatusCode(v int32) *DescribeDataObjectColumnDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponse) SetBody(v *DescribeDataObjectColumnDetailResponseBody) *DescribeDataObjectColumnDetailResponse {
	s.Body = v
	return s
}

type DescribeDataObjectColumnDetailV2Request struct {
	// When performing a paginated query, set the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// Set the unique identifier ID of the data object to be queried.
	//
	// > You can obtain the identifier ID by calling [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 13456723343
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language type for the request and response messages, default is **zh_cn**. Values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// When performing a paginated query, set the maximum number of data asset instances to display per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID corresponding to the product name of the data object. Values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: TableStore
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// - **7**: PolarDB-X
	//
	// - **8**: PolarDB
	//
	// - **9**: ADB-PG
	//
	// - **10**: OceanBase
	//
	// - **11**: MongoDB
	//
	// - **25**: Redis
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// Industry template ID.
	//
	// > You can obtain the industry template identifier ID by calling [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDataObjectColumnDetailV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectColumnDetailV2Request) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailV2Request) SetCurrentPage(v int32) *DescribeDataObjectColumnDetailV2Request {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetFeatureType(v int32) *DescribeDataObjectColumnDetailV2Request {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetId(v string) *DescribeDataObjectColumnDetailV2Request {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetLang(v string) *DescribeDataObjectColumnDetailV2Request {
	s.Lang = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetPageSize(v int32) *DescribeDataObjectColumnDetailV2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetProductId(v int64) *DescribeDataObjectColumnDetailV2Request {
	s.ProductId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetTemplateId(v int64) *DescribeDataObjectColumnDetailV2Request {
	s.TemplateId = &v
	return s
}

type DescribeDataObjectColumnDetailV2ResponseBody struct {
	// Page number for paginated queries. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of recognition results for the columns in the data table.
	Items []*DescribeDataObjectColumnDetailV2ResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// When performing a paginated query, set the maximum number of data asset instances to display per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of this call request, which is a unique identifier generated by Alibaba Cloud for the request, and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of data entries.
	//
	// example:
	//
	// 231
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataObjectColumnDetailV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectColumnDetailV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) SetCurrentPage(v int32) *DescribeDataObjectColumnDetailV2ResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) SetItems(v []*DescribeDataObjectColumnDetailV2ResponseBodyItems) *DescribeDataObjectColumnDetailV2ResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) SetPageSize(v int32) *DescribeDataObjectColumnDetailV2ResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) SetRequestId(v string) *DescribeDataObjectColumnDetailV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBody) SetTotalCount(v int32) *DescribeDataObjectColumnDetailV2ResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataObjectColumnDetailV2ResponseBodyItems struct {
	// List of industry categories for sensitive data.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// Comment for the column.
	//
	// example:
	//
	// column comment
	ColumnComment *string `json:"ColumnComment,omitempty" xml:"ColumnComment,omitempty"`
	// Column name.
	//
	// example:
	//
	// hide14
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// Data type of the column.
	//
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Engine type. Values:
	//
	// - **MySQL**
	//
	// - **MariaDB**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// Unique identifier ID of the data object.
	//
	// example:
	//
	// 1392973973691383808
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Instance name of the data asset table.
	//
	// example:
	//
	// rm-bp17t1htja573l5i8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Column encryption status. Values:
	//
	// - **-1**: Not encrypted
	//
	// - **1**: Encryption successful
	//
	// - **2**: Encryption failed
	//
	// example:
	//
	// -1
	MaskingStatus *int32 `json:"MaskingStatus,omitempty" xml:"MaskingStatus,omitempty"`
	// List of data tags.
	ModelTags []*DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// Whether the column is a primary key. Value explanation:
	//
	// - **true**: Primary key.
	//
	// - **false**: Not a primary key.
	//
	// example:
	//
	// true
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// ID corresponding to the product name of the data object. Values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: TableStore
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// - **7**: PolarDB-X
	//
	// - **8**: PolarDB
	//
	// - **9**: ADB-PG
	//
	// - **10**: OceanBase
	//
	// - **11**: MongoDB
	//
	// - **25**: Redis
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// Region where the asset is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Risk level ID. Values:
	//
	// - **1**: N/A: No sensitive data detected.
	//
	// - **2**: S1: Level 1 sensitive data.
	//
	// - **3**: S2: Level 2 sensitive data.
	//
	// - **4**: S3: Level 3 sensitive data.
	//
	// - **5**: S4: Level 4 sensitive data.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// 风险等级名称。取值：
	//
	// - **N/A**：未识别到敏感数据。
	//
	// - **S1**：1级敏感数据。
	//
	// - **S2**：2级敏感数据。
	//
	// - **S3**：3级敏感数据。
	//
	// - **S4**：4级敏感数据。
	//
	// example:
	//
	// S1
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The ID of the matched recognition model.
	//
	// example:
	//
	// 51
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the matched recognition model.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Table name.
	//
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeDataObjectColumnDetailV2ResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectColumnDetailV2ResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetCategories(v []*string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.Categories = v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetColumnComment(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.ColumnComment = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetColumnName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.ColumnName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetDataType(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetEngineType(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetId(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetInstanceName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetMaskingStatus(v int32) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.MaskingStatus = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetModelTags(v []*DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetPrimaryKey(v bool) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetProductId(v int64) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetRegionId(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetRiskLevelId(v int64) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetRiskLevelName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetRuleId(v int64) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetRuleName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItems) SetTableName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItems {
	s.TableName = &v
	return s
}

type DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags struct {
	// ID corresponding to the data tag name. Values:
	//
	// - **101**: Personal Sensitive Information
	//
	// - **102**: Personal Information
	//
	// - **107**: General Information
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Data tag name. Values:
	//
	// - Personal Sensitive Information
	//
	// - Personal Information
	//
	// - General Information
	//
	// example:
	//
	// personal sensitive information
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) SetId(v int64) *DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags) SetName(v string) *DescribeDataObjectColumnDetailV2ResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

type DescribeDataObjectColumnDetailV2Response struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataObjectColumnDetailV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataObjectColumnDetailV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectColumnDetailV2Response) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailV2Response) SetHeaders(v map[string]*string) *DescribeDataObjectColumnDetailV2Response {
	s.Headers = v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Response) SetStatusCode(v int32) *DescribeDataObjectColumnDetailV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Response) SetBody(v *DescribeDataObjectColumnDetailV2ResponseBody) *DescribeDataObjectColumnDetailV2Response {
	s.Body = v
	return s
}

type DescribeDataObjectsRequest struct {
	// Page number for the paginated query. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// ID of the data domain to which the data asset belongs.
	//
	// example:
	//
	// 2
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// File category code.
	//
	// example:
	//
	// 1
	FileCategoryCode *int64 `json:"FileCategoryCode,omitempty" xml:"FileCategoryCode,omitempty"`
	// OSS file types that are supported for recognition.
	//
	// > You can obtain the supported OSS file types by calling [DescribeDocTypes](https://help.aliyun.com/document_detail/2536492.html), using the Code field value from the response. This parameter is only valid for querying OSS-type assets.
	//
	// example:
	//
	// 100001
	FileType *int64 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// Keyword for the asset instance ID.
	//
	// example:
	//
	// 8vb54hn2g9j191ddz
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language type for request and response messages, default is **zh_cn**. Values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Member account ID.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// Model IDs of the industry template, separated by commas.
	//
	// > You can obtain the industry template model identifier ID by calling [DescribeTemplateAllRules](https://help.aliyun.com/document_detail/2536491.html).
	//
	// example:
	//
	// 101
	ModelIds *string `json:"ModelIds,omitempty" xml:"ModelIds,omitempty"`
	// Data labels to be queried, separated by commas. Values:
	//
	// - **101**: Personal Sensitive Information.
	//
	// - **102**: Personal Information.
	//
	// - **107**: General Information.
	//
	// example:
	//
	// 101,102
	ModelTagIds *string `json:"ModelTagIds,omitempty" xml:"ModelTagIds,omitempty"`
	// When performing a paginated query, set the maximum number of data asset instances to display per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of parent category IDs for the template to be queried, separated by commas.
	//
	// example:
	//
	// 234,236,238
	ParentCategoryIds *string `json:"ParentCategoryIds,omitempty" xml:"ParentCategoryIds,omitempty"`
	// It is recommended to fill in the list of product IDs to be queried, separated by commas. Values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: TableStore
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// - **7**: PolarDB-X
	//
	// - **8**: PolarDB
	//
	// - **9**: ADB-PG
	//
	// - **10**: OceanBase
	//
	// - **11**: MongoDB
	//
	// - **25**: Redis
	//
	// > OSS is mutually exclusive with other products, meaning if OSS is included in the query, no other products can be listed; by default, non-OSS products are queried.
	//
	// example:
	//
	// 1,5
	ProductIds *string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty"`
	// Keyword for the data object to be queried.
	//
	// example:
	//
	// t_sddp_selfmysql_pers0
	QueryName *string `json:"QueryName,omitempty" xml:"QueryName,omitempty"`
	// Specify the risk levels of the data assets to be queried, separated by commas if multiple.
	//
	// - **2**: S1, low risk level.
	//
	// - **3**: S2, medium risk level.
	//
	// - **4**: S3, high risk level.
	//
	// - **5**: S4, highest risk level.
	//
	// example:
	//
	// 2
	RiskLevels *string `json:"RiskLevels,omitempty" xml:"RiskLevels,omitempty"`
	// Region where the asset is located. Values:
	//
	// - **cn-beijing**: North China 2 (Beijing).
	//
	// - **cn-zhangjiakou**: North China 3 (Zhangjiakou).
	//
	// - **cn-huhehaote**: North China 5 (Hohhot).
	//
	// - **cn-hangzhou**: East China 1 (Hangzhou).
	//
	// - **cn-shanghai**: East China 2 (Shanghai).
	//
	// - **cn-shenzhen**: South China 1 (Shenzhen).
	//
	// - **cn-hongkong**: Hong Kong, China.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// Industry template ID.
	//
	// > You can obtain the industry template identifier ID by calling [DescribeCategoryTemplateList](https://help.aliyun.com/document_detail/2399296.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDataObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsRequest) SetCurrentPage(v int32) *DescribeDataObjectsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetDomainId(v int64) *DescribeDataObjectsRequest {
	s.DomainId = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetFeatureType(v int32) *DescribeDataObjectsRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetFileCategoryCode(v int64) *DescribeDataObjectsRequest {
	s.FileCategoryCode = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetFileType(v int64) *DescribeDataObjectsRequest {
	s.FileType = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetInstanceId(v string) *DescribeDataObjectsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetLang(v string) *DescribeDataObjectsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetMemberAccount(v int64) *DescribeDataObjectsRequest {
	s.MemberAccount = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetModelIds(v string) *DescribeDataObjectsRequest {
	s.ModelIds = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetModelTagIds(v string) *DescribeDataObjectsRequest {
	s.ModelTagIds = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetPageSize(v int32) *DescribeDataObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetParentCategoryIds(v string) *DescribeDataObjectsRequest {
	s.ParentCategoryIds = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetProductIds(v string) *DescribeDataObjectsRequest {
	s.ProductIds = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetQueryName(v string) *DescribeDataObjectsRequest {
	s.QueryName = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetRiskLevels(v string) *DescribeDataObjectsRequest {
	s.RiskLevels = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetServiceRegionId(v string) *DescribeDataObjectsRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeDataObjectsRequest) SetTemplateId(v int64) *DescribeDataObjectsRequest {
	s.TemplateId = &v
	return s
}

type DescribeDataObjectsResponseBody struct {
	// When performing a paginated query, set the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of data objects.
	Items []*DescribeDataObjectsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// When performing a paginated query, this sets the maximum number of data asset instances to display per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 本次调用请求的ID，是由阿里云为该请求生成的唯一标识符，可用于排查和定位问题。
	//
	// example:
	//
	// E6F6460E-4330-549A-BD89-C183FB17571E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of query results.
	//
	// example:
	//
	// 21
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsResponseBody) SetCurrentPage(v int32) *DescribeDataObjectsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectsResponseBody) SetItems(v []*DescribeDataObjectsResponseBodyItems) *DescribeDataObjectsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataObjectsResponseBody) SetPageSize(v int32) *DescribeDataObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectsResponseBody) SetRequestId(v string) *DescribeDataObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataObjectsResponseBody) SetTotalCount(v int32) *DescribeDataObjectsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDataObjectsResponseBodyItems struct {
	// List of industry categories for the sensitive data.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The unique identifier ID of the data object.
	//
	// example:
	//
	// 20000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance description of the data object.
	//
	// example:
	//
	// instance description
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// Data asset instance ID.
	//
	// example:
	//
	// rm-12*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Latest file modification time, in milliseconds.
	//
	// example:
	//
	// 1687676649830
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The timestamp of the last scan, in milliseconds.
	//
	// example:
	//
	// 1687676649830
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// Member account ID.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// List of data tags.
	ModelTags []*DescribeDataObjectsResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// The name of the data object.
	//
	// example:
	//
	// t_sddp_selfmysql_pers0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// File category name.
	//
	// example:
	//
	// text file
	ObjectFileCategory *string `json:"ObjectFileCategory,omitempty" xml:"ObjectFileCategory,omitempty"`
	// The type of the data object.
	//
	// example:
	//
	// text type
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The path of the data object.
	//
	// example:
	//
	// rm-12**.db_***
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The product name that the data object belongs to. Values:
	//
	// - **MaxCompute**
	//
	// - **OSS**
	//
	// - **ADB-MYSQL**
	//
	// - **TableStore**
	//
	// - **RDS**
	//
	// - **SELF_DB**
	//
	// - **PolarDB-X**
	//
	// - **PolarDB**
	//
	// - **ADB-PG**
	//
	// - **OceanBase**
	//
	// - **MongoDB**
	//
	// - **Redis**
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID corresponding to the product name that the data object belongs to. Values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: TableStore
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// - **7**: PolarDB-X
	//
	// - **8**: PolarDB
	//
	// - **9**: ADB-PG
	//
	// - **10**: OceanBase
	//
	// - **11**: MongoDB
	//
	// - **25**: Redis
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// OSS存储对象所属区域ID。
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Region name.
	//
	// example:
	//
	// cn-hangzhou
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// List of matched identification models.
	RuleList []*DescribeDataObjectsResponseBodyItemsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// Number of sensitive data items.
	//
	// example:
	//
	// 1
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// Industry template ID
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDataObjectsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsResponseBodyItems) SetCategories(v []*string) *DescribeDataObjectsResponseBodyItems {
	s.Categories = v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetId(v string) *DescribeDataObjectsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetInstanceDescription(v string) *DescribeDataObjectsResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetInstanceId(v string) *DescribeDataObjectsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetLastModifiedTime(v int64) *DescribeDataObjectsResponseBodyItems {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetLastScanTime(v int64) *DescribeDataObjectsResponseBodyItems {
	s.LastScanTime = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetMemberAccount(v int64) *DescribeDataObjectsResponseBodyItems {
	s.MemberAccount = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetModelTags(v []*DescribeDataObjectsResponseBodyItemsModelTags) *DescribeDataObjectsResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetName(v string) *DescribeDataObjectsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetObjectFileCategory(v string) *DescribeDataObjectsResponseBodyItems {
	s.ObjectFileCategory = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetObjectType(v string) *DescribeDataObjectsResponseBodyItems {
	s.ObjectType = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetPath(v string) *DescribeDataObjectsResponseBodyItems {
	s.Path = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetProductCode(v string) *DescribeDataObjectsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetProductId(v int64) *DescribeDataObjectsResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetRegionId(v string) *DescribeDataObjectsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetRegionName(v string) *DescribeDataObjectsResponseBodyItems {
	s.RegionName = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetRuleList(v []*DescribeDataObjectsResponseBodyItemsRuleList) *DescribeDataObjectsResponseBodyItems {
	s.RuleList = v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetSensitiveCount(v int32) *DescribeDataObjectsResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItems) SetTemplateId(v int64) *DescribeDataObjectsResponseBodyItems {
	s.TemplateId = &v
	return s
}

type DescribeDataObjectsResponseBodyItemsModelTags struct {
	// Data tag ID. Values:
	//
	// - **101**: Personal sensitive information.
	//
	// - **102**: Personal information.
	//
	// - **107**: General information.
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Data tag name. Values:
	//
	// - **Personal sensitive information.**
	//
	// - **Personal information.**
	//
	// - **General information.**
	//
	// example:
	//
	// Personal sensitive information
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDataObjectsResponseBodyItemsModelTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectsResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsResponseBodyItemsModelTags) SetId(v int64) *DescribeDataObjectsResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsModelTags) SetName(v string) *DescribeDataObjectsResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

type DescribeDataObjectsResponseBodyItemsRuleList struct {
	// Risk level ID for sensitive data identification rules. Values:
	//
	// - **1**: N/A: No sensitive data identified.
	//
	// - **2**: S1: Level 1 sensitive data.
	//
	// - **3**: S2: Level 2 sensitive data.
	//
	// - **4**: S3: Level 3 sensitive data.
	//
	// - **5**: S4: Level 4 sensitive data.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// Risk level name for the data asset table. Values:
	//
	// - **N/A**: No sensitive data identified.
	//
	// - **S1**: Level 1 sensitive data.
	//
	// - **S2**: Level 2 sensitive data.
	//
	// - **S3**: Level 3 sensitive data.
	//
	// - **S4**: Level 4 sensitive data.
	//
	// example:
	//
	// S1
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// Number of matched identification models.
	//
	// example:
	//
	// 590
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// Identification model ID.
	//
	// example:
	//
	// 1080
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Identification model name.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeDataObjectsResponseBodyItemsRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectsResponseBodyItemsRuleList) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetRiskLevelId(v int64) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetRiskLevelName(v string) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetRuleCount(v int32) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.RuleCount = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetRuleId(v int64) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.RuleId = &v
	return s
}

func (s *DescribeDataObjectsResponseBodyItemsRuleList) SetRuleName(v string) *DescribeDataObjectsResponseBodyItemsRuleList {
	s.RuleName = &v
	return s
}

type DescribeDataObjectsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataObjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectsResponse) SetHeaders(v map[string]*string) *DescribeDataObjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataObjectsResponse) SetStatusCode(v int32) *DescribeDataObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataObjectsResponse) SetBody(v *DescribeDataObjectsResponseBody) *DescribeDataObjectsResponse {
	s.Body = v
	return s
}

type DescribeDocTypesRequest struct {
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDocTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocTypesRequest) SetLang(v string) *DescribeDocTypesRequest {
	s.Lang = &v
	return s
}

type DescribeDocTypesResponseBody struct {
	// A list of OSS object types.
	DocTypeList []*DescribeDocTypesResponseBodyDocTypeList `json:"DocTypeList,omitempty" xml:"DocTypeList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-4******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDocTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocTypesResponseBody) SetDocTypeList(v []*DescribeDocTypesResponseBodyDocTypeList) *DescribeDocTypesResponseBody {
	s.DocTypeList = v
	return s
}

func (s *DescribeDocTypesResponseBody) SetRequestId(v string) *DescribeDocTypesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDocTypesResponseBodyDocTypeList struct {
	// The code of the object type.
	//
	// example:
	//
	// 100001
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the object type.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the object type.
	//
	// example:
	//
	// C/C++ Source Code
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDocTypesResponseBodyDocTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocTypesResponseBodyDocTypeList) GoString() string {
	return s.String()
}

func (s *DescribeDocTypesResponseBodyDocTypeList) SetCode(v int64) *DescribeDocTypesResponseBodyDocTypeList {
	s.Code = &v
	return s
}

func (s *DescribeDocTypesResponseBodyDocTypeList) SetId(v int64) *DescribeDocTypesResponseBodyDocTypeList {
	s.Id = &v
	return s
}

func (s *DescribeDocTypesResponseBodyDocTypeList) SetName(v string) *DescribeDocTypesResponseBodyDocTypeList {
	s.Name = &v
	return s
}

type DescribeDocTypesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDocTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDocTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDocTypesResponse) SetHeaders(v map[string]*string) *DescribeDocTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDocTypesResponse) SetStatusCode(v int32) *DescribeDocTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDocTypesResponse) SetBody(v *DescribeDocTypesResponseBody) *DescribeDocTypesResponse {
	s.Body = v
	return s
}

type DescribeEventDetailRequest struct {
	// The ID of the anomalous event.
	//
	// > You can call the **DescribeEvents*	- operation to query the ID of the anomalous event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13456723343
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeEventDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailRequest) SetId(v int64) *DescribeEventDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeEventDetailRequest) SetLang(v string) *DescribeEventDetailRequest {
	s.Lang = &v
	return s
}

type DescribeEventDetailResponseBody struct {
	// The details of the anomalous event.
	Event *DescribeEventDetailResponseBodyEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 69FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEventDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBody) SetEvent(v *DescribeEventDetailResponseBodyEvent) *DescribeEventDetailResponseBody {
	s.Event = v
	return s
}

func (s *DescribeEventDetailResponseBody) SetRequestId(v string) *DescribeEventDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEventDetailResponseBodyEvent struct {
	// The time when the alert for the anomalous event was generated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545829129000
	AlertTime *int64 `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	// Indicates whether the handling result of the anomalous event is used to enhance the detection of anomalous events. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// > If you enhance the detection of anomalous events, the detection accuracy and the rate of triggering alerts for anomalous events are improved.
	//
	// example:
	//
	// false
	Backed *bool `json:"Backed,omitempty" xml:"Backed,omitempty"`
	// The instance name of the service in which the anomalous event was detected.
	//
	// example:
	//
	// in-222***
	DataInstance *string `json:"DataInstance,omitempty" xml:"DataInstance,omitempty"`
	// The display name of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// yundunsr
	DealDisplayName *string `json:"DealDisplayName,omitempty" xml:"DealDisplayName,omitempty"`
	// The username of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// det1111
	DealLoginName *string `json:"DealLoginName,omitempty" xml:"DealLoginName,omitempty"`
	// The reason why the anomalous event is handled.
	//
	// example:
	//
	// Anomaly confirmed
	DealReason *string `json:"DealReason,omitempty" xml:"DealReason,omitempty"`
	// The time when the anomalous event was handled. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1230000
	DealTime *int64 `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// The ID of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// 229157443385014***
	DealUserId *int64 `json:"DealUserId,omitempty" xml:"DealUserId,omitempty"`
	// The content in the details of the anomalous event.
	Detail *DescribeEventDetailResponseBodyEventDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// The display name of the account that triggered the anomalous event.
	//
	// example:
	//
	// yundunsr
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the anomalous event occurred. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545829129000
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// An array that consists of the handling records of the anomalous event.
	HandleInfoList []*DescribeEventDetailResponseBodyEventHandleInfoList `json:"HandleInfoList,omitempty" xml:"HandleInfoList,omitempty" type:"Repeated"`
	// The unique ID of the anomalous event.
	//
	// example:
	//
	// 52234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The details of the alert logs.
	//
	// example:
	//
	// {"client_ip": ["106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX", "106.11.XX.XX"], "start_time": "2020-05-10 00:00:01", "instance": ["omniscience-data", "punish-beaver-data"], "end_time": "2020-05-10 00:21:22", "client_ua": ["Java/1.8.0_152", "Java/1.8.0_92", "aliyun-sdk-java/2.0.0", "aliyun-sdk-java/2.8.0(Linux/4.9.151-015.ali3000.alios7.x86_64/amd64;1.8.0_152)"], "user_name": 1512222261295262}
	LogDetail *string `json:"LogDetail,omitempty" xml:"LogDetail,omitempty"`
	// The username of the account that triggered the anomalous event.
	//
	// example:
	//
	// det1111
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// Whether it is a new version of the alarm. Value:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	NewAlarm *bool `json:"NewAlarm,omitempty" xml:"NewAlarm,omitempty"`
	// The name of the service in which the anomalous event was detected. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The handling status for the anomalous event. Valid values:
	//
	// 	- **0**: unhandled
	//
	// 	- **1**: confirmed
	//
	// 	- **2**: marked as false positive
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the handling status for the anomalous event.
	//
	// example:
	//
	// Pending
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// The code of the anomalous event subtype.
	//
	// example:
	//
	// 020008
	SubTypeCode *string `json:"SubTypeCode,omitempty" xml:"SubTypeCode,omitempty"`
	// The name of the anomalous event subtype.
	//
	// example:
	//
	// Anomalous volume of downloaded data
	SubTypeName *string `json:"SubTypeName,omitempty" xml:"SubTypeName,omitempty"`
	// The code of the anomalous event type.
	//
	// example:
	//
	// 02
	TypeCode *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	// The name of the anomalous event type. Valid values:
	//
	// 	- **01**: anomalous permission usage
	//
	// 	- **02**: anomalous data flow
	//
	// 	- **03**: anomalous data operation
	//
	// example:
	//
	// Anomalous data flow
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// The ID of the account that triggered the anomalous event.
	//
	// example:
	//
	// 229157443385014***
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEventDetailResponseBodyEvent) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEvent) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEvent) SetAlertTime(v int64) *DescribeEventDetailResponseBodyEvent {
	s.AlertTime = &v
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

func (s *DescribeEventDetailResponseBodyEvent) SetDealDisplayName(v string) *DescribeEventDetailResponseBodyEvent {
	s.DealDisplayName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealLoginName(v string) *DescribeEventDetailResponseBodyEvent {
	s.DealLoginName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealReason(v string) *DescribeEventDetailResponseBodyEvent {
	s.DealReason = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealTime(v int64) *DescribeEventDetailResponseBodyEvent {
	s.DealTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDealUserId(v int64) *DescribeEventDetailResponseBodyEvent {
	s.DealUserId = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDetail(v *DescribeEventDetailResponseBodyEventDetail) *DescribeEventDetailResponseBodyEvent {
	s.Detail = v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetDisplayName(v string) *DescribeEventDetailResponseBodyEvent {
	s.DisplayName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetEventTime(v int64) *DescribeEventDetailResponseBodyEvent {
	s.EventTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetHandleInfoList(v []*DescribeEventDetailResponseBodyEventHandleInfoList) *DescribeEventDetailResponseBodyEvent {
	s.HandleInfoList = v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetId(v int64) *DescribeEventDetailResponseBodyEvent {
	s.Id = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetLogDetail(v string) *DescribeEventDetailResponseBodyEvent {
	s.LogDetail = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetLoginName(v string) *DescribeEventDetailResponseBodyEvent {
	s.LoginName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetNewAlarm(v bool) *DescribeEventDetailResponseBodyEvent {
	s.NewAlarm = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetProductCode(v string) *DescribeEventDetailResponseBodyEvent {
	s.ProductCode = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetStatus(v int32) *DescribeEventDetailResponseBodyEvent {
	s.Status = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetStatusName(v string) *DescribeEventDetailResponseBodyEvent {
	s.StatusName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetSubTypeCode(v string) *DescribeEventDetailResponseBodyEvent {
	s.SubTypeCode = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetSubTypeName(v string) *DescribeEventDetailResponseBodyEvent {
	s.SubTypeName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetTypeCode(v string) *DescribeEventDetailResponseBodyEvent {
	s.TypeCode = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetTypeName(v string) *DescribeEventDetailResponseBodyEvent {
	s.TypeName = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEvent) SetUserId(v int64) *DescribeEventDetailResponseBodyEvent {
	s.UserId = &v
	return s
}

type DescribeEventDetailResponseBodyEventDetail struct {
	// The baseline behavior chart of the anomalous event.
	Chart []*DescribeEventDetailResponseBodyEventDetailChart `json:"Chart,omitempty" xml:"Chart,omitempty" type:"Repeated"`
	// The content in the anomalous event.
	Content []*DescribeEventDetailResponseBodyEventDetailContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// An array that consists of the source from which the information of the anomalous event is recorded.
	ResourceInfo []*DescribeEventDetailResponseBodyEventDetailResourceInfo `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty" type:"Repeated"`
}

func (s DescribeEventDetailResponseBodyEventDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetail) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetail) SetChart(v []*DescribeEventDetailResponseBodyEventDetailChart) *DescribeEventDetailResponseBodyEventDetail {
	s.Chart = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetail) SetContent(v []*DescribeEventDetailResponseBodyEventDetailContent) *DescribeEventDetailResponseBodyEventDetail {
	s.Content = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetail) SetResourceInfo(v []*DescribeEventDetailResponseBodyEventDetailResourceInfo) *DescribeEventDetailResponseBodyEventDetail {
	s.ResourceInfo = v
	return s
}

type DescribeEventDetailResponseBodyEventDetailChart struct {
	// The type of the chart. Valid values:
	//
	// 	- **1**: column chart
	//
	// 	- **2**: line chart
	//
	// >This field will be returned only when NewAlarm is true.
	//
	// example:
	//
	// 1
	ChatType *int32 `json:"ChatType,omitempty" xml:"ChatType,omitempty"`
	// The data in the baseline behavior profile of the anomalous event.
	Data *DescribeEventDetailResponseBodyEventDetailChartData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The name of the baseline behavior chart of the anomalous event.
	//
	// example:
	//
	// Baseline behavior chart
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Icon title.
	//
	// >This field will be returned only when NewAlarm is true.
	//
	// example:
	//
	// misskingm
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the chart. Valid values:
	//
	// 	- **1**: column chart
	//
	// 	- **2**: line chart
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The descriptive label of data items on the X axis.
	//
	// example:
	//
	// Number of days
	XLabel *string `json:"XLabel,omitempty" xml:"XLabel,omitempty"`
	// The descriptive label of data items on the Y axis.
	//
	// example:
	//
	// Value
	YLabel *string `json:"YLabel,omitempty" xml:"YLabel,omitempty"`
	// The descriptive label of data items on the Z axis.
	//
	// >This field will be returned only when NewAlarm is true.
	//
	// example:
	//
	// chart description
	ZLabel *string `json:"ZLabel,omitempty" xml:"ZLabel,omitempty"`
}

func (s DescribeEventDetailResponseBodyEventDetailChart) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetailChart) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetChatType(v int32) *DescribeEventDetailResponseBodyEventDetailChart {
	s.ChatType = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetData(v *DescribeEventDetailResponseBodyEventDetailChartData) *DescribeEventDetailResponseBodyEventDetailChart {
	s.Data = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetLabel(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.Label = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetName(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.Name = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetType(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.Type = &v
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

func (s *DescribeEventDetailResponseBodyEventDetailChart) SetZLabel(v string) *DescribeEventDetailResponseBodyEventDetailChart {
	s.ZLabel = &v
	return s
}

type DescribeEventDetailResponseBodyEventDetailChartData struct {
	// The value of the data item on the X axis.
	//
	// example:
	//
	// [test1,test2,...]
	X []*string `json:"X,omitempty" xml:"X,omitempty" type:"Repeated"`
	// The value of the data item on the Y axis.
	//
	// example:
	//
	// [1,2,3,...]
	Y []*string `json:"Y,omitempty" xml:"Y,omitempty" type:"Repeated"`
	// The value of the data item for the Z axis.
	Z []*string `json:"Z,omitempty" xml:"Z,omitempty" type:"Repeated"`
}

func (s DescribeEventDetailResponseBodyEventDetailChartData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventDetailChartData) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) SetX(v []*string) *DescribeEventDetailResponseBodyEventDetailChartData {
	s.X = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) SetY(v []*string) *DescribeEventDetailResponseBodyEventDetailChartData {
	s.Y = v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailChartData) SetZ(v []*string) *DescribeEventDetailResponseBodyEventDetailChartData {
	s.Z = v
	return s
}

type DescribeEventDetailResponseBodyEventDetailContent struct {
	// The title of the content in the anomalous event.
	//
	// example:
	//
	// Anomaly description
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Exception event name.
	//
	// example:
	//
	// daliaoyuncom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the content in the anomalous event.
	//
	// example:
	//
	// The account was used to access OSS from an unusual terminal whose IP address is 1.2.3.4 from 00:06:45 on September 9, 2019 to 00:57:37 on September 9, 2019.
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

func (s *DescribeEventDetailResponseBodyEventDetailContent) SetName(v string) *DescribeEventDetailResponseBodyEventDetailContent {
	s.Name = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventDetailContent) SetValue(v string) *DescribeEventDetailResponseBodyEventDetailContent {
	s.Value = &v
	return s
}

type DescribeEventDetailResponseBodyEventDetailResourceInfo struct {
	// The source title.
	//
	// example:
	//
	// Risk
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The source description.
	//
	// example:
	//
	// Based on the record of authentication by using an unusual terminal, an attacker may have obtained the access permission of the account, or an employee accessed data from a personal terminal.
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

type DescribeEventDetailResponseBodyEventHandleInfoList struct {
	// The account that is used to handle the anomalous event.
	//
	// example:
	//
	// sddp-test2
	CurrentValue *string `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	// The time when the account is disabled. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1611139155000
	DisableTime *int64 `json:"DisableTime,omitempty" xml:"DisableTime,omitempty"`
	// The time when the disabled account is enabled. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1611139155000
	EnableTime *int64 `json:"EnableTime,omitempty" xml:"EnableTime,omitempty"`
	// The handling method.
	//
	// example:
	//
	// Remove from the whitelist
	HandlerName *string `json:"HandlerName,omitempty" xml:"HandlerName,omitempty"`
	// The type of the handling method.
	//
	// example:
	//
	// rds_security_ip
	HandlerType *string `json:"HandlerType,omitempty" xml:"HandlerType,omitempty"`
	// The duration for which the handling operation takes effect. If you leave this parameter empty, the handling operation is permanently valid. Unit: minutes.
	//
	// example:
	//
	// 10
	HandlerValue *int32 `json:"HandlerValue,omitempty" xml:"HandlerValue,omitempty"`
	// The ID of the handling rule.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the account that triggered the anomalous event. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// 	- **-1**: failed to disable the account
	//
	// 	- **-2**: failed to enable the account
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventDetailResponseBodyEventHandleInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDetailResponseBodyEventHandleInfoList) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetCurrentValue(v string) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.CurrentValue = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetDisableTime(v int64) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.DisableTime = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetEnableTime(v int64) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.EnableTime = &v
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

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetHandlerValue(v int32) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.HandlerValue = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetId(v int64) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.Id = &v
	return s
}

func (s *DescribeEventDetailResponseBodyEventHandleInfoList) SetStatus(v int32) *DescribeEventDetailResponseBodyEventHandleInfoList {
	s.Status = &v
	return s
}

type DescribeEventDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeEventDetailResponse) SetStatusCode(v int32) *DescribeEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventDetailResponse) SetBody(v *DescribeEventDetailResponseBody) *DescribeEventDetailResponse {
	s.Body = v
	return s
}

type DescribeEventTypesRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of anomalous event for which you want to query the anomalous events of subtypes. Valid values:
	//
	// 	- **01**: anomalous permission usage
	//
	// 	- **02**: anomalous data flow
	//
	// 	- **03**: anomalous data operation
	//
	// example:
	//
	// 01
	ParentTypeId *int64 `json:"ParentTypeId,omitempty" xml:"ParentTypeId,omitempty"`
	// The type of the resource. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 5
	ResourceId *int32 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The status of the anomalous event. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **2**: disabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesRequest) SetFeatureType(v int32) *DescribeEventTypesRequest {
	s.FeatureType = &v
	return s
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
	// An array that consists of the types of anomalous events.
	//
	// > If you leave the ParentTypeId parameter empty, anomalous event types are returned. If you set the ParentTypeId parameter, anomalous event subtypes under the specified anomalous event type are returned.
	EventTypeList []*DescribeEventTypesResponseBodyEventTypeList `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEventTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesResponseBody) SetEventTypeList(v []*DescribeEventTypesResponseBodyEventTypeList) *DescribeEventTypesResponseBody {
	s.EventTypeList = v
	return s
}

func (s *DescribeEventTypesResponseBody) SetRequestId(v string) *DescribeEventTypesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEventTypesResponseBodyEventTypeList struct {
	// The code of the anomalous event type.
	//
	// example:
	//
	// 01
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the anomalous event type.
	//
	// example:
	//
	// Anomalous permission usage,\\*\\*\\*\\*
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the anomalous event type.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the anomalous event type.
	//
	// example:
	//
	// Anomalous permission usage
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// An array that consists of anomalous event subtypes.
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

func (s *DescribeEventTypesResponseBodyEventTypeList) SetId(v int64) *DescribeEventTypesResponseBodyEventTypeList {
	s.Id = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetName(v string) *DescribeEventTypesResponseBodyEventTypeList {
	s.Name = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetSubTypeList(v []*DescribeEventTypesResponseBodyEventTypeListSubTypeList) *DescribeEventTypesResponseBodyEventTypeList {
	s.SubTypeList = v
	return s
}

type DescribeEventTypesResponseBodyEventTypeListSubTypeList struct {
	// The service to which the anomalous event detection rule applies. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// RDS
	AdaptedProduct *string `json:"AdaptedProduct,omitempty" xml:"AdaptedProduct,omitempty"`
	// The code of the anomalous event subtype.
	//
	// example:
	//
	// 020008
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The code of the configuration.
	//
	// example:
	//
	// 0100**
	ConfigCode *string `json:"ConfigCode,omitempty" xml:"ConfigCode,omitempty"`
	// The content format of anomalous event detection rule. Valid values:
	//
	// 	- **0**: numeric values such as thresholds
	//
	// 	- **1**: text such as IP addresses
	//
	// example:
	//
	// 1
	ConfigContentType *int32 `json:"ConfigContentType,omitempty" xml:"ConfigContentType,omitempty"`
	// The description of the configuration.
	//
	// example:
	//
	// The period of time for which the permission is not used exceeds the threshold. The specified threshold is ${value} calendar days.
	ConfigDescription *string `json:"ConfigDescription,omitempty" xml:"ConfigDescription,omitempty"`
	// The value of the configuration.
	//
	// example:
	//
	// 90
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The description of the anomalous event subtype.
	//
	// example:
	//
	// Inappropriate configuration-No protection for the MaxCompute sensitive project, \\*\\*\\*\\*
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of times that the anomalous event hits the anomalous event detection rule.
	//
	// example:
	//
	// 2
	EventHitCount *int32 `json:"EventHitCount,omitempty" xml:"EventHitCount,omitempty"`
	// The ID of the anomalous event subtype.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the anomalous event subtype.
	//
	// example:
	//
	// Inappropriate configuration-No protection for the MaxCompute sensitive project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether detection is enabled for the anomalous event subtype. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventTypesResponseBodyEventTypeListSubTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventTypesResponseBodyEventTypeListSubTypeList) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetAdaptedProduct(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.AdaptedProduct = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetCode(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Code = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigCode(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigCode = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigContentType(v int32) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigContentType = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigDescription(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigDescription = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigValue(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigValue = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetDescription(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Description = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetEventHitCount(v int32) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.EventHitCount = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetId(v int64) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Id = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetName(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Name = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetStatus(v int32) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Status = &v
	return s
}

type DescribeEventTypesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeEventTypesResponse) SetStatusCode(v int32) *DescribeEventTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventTypesResponse) SetBody(v *DescribeEventTypesResponseBody) *DescribeEventTypesResponse {
	s.Body = v
	return s
}

type DescribeEventsRequest struct {
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the account that handles the anomalous event.
	//
	// example:
	//
	// yundun-***
	DealUserId *string `json:"DealUserId,omitempty" xml:"DealUserId,omitempty"`
	// The end of the time range during which the anomalous events are detected. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1698700000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The unique ID of the anomalous event.
	//
	// example:
	//
	// 789026
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data asset.
	//
	// example:
	//
	// rm-uf6yzvbc2tg90iuxk.l****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service to which the table belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// OSS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The beginning of the time range during which the anomalous events are detected. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1657900000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The handling status of the anomalous event. Valid values:
	//
	// 	- 0: unhandled
	//
	// 	- 1: confirmed
	//
	// 	- 2: marked as false positive
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the anomalous event subtype.
	//
	// > You can call the **DescribeEventTypes*	- operation to query the name of the anomalous event subtype.
	//
	// example:
	//
	// Anomalous volume of downloaded data
	SubTypeCode *string `json:"SubTypeCode,omitempty" xml:"SubTypeCode,omitempty"`
	// The name of the destination service in an anomalous data flow. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**
	//
	// example:
	//
	// RDS
	TargetProductCode *string `json:"TargetProductCode,omitempty" xml:"TargetProductCode,omitempty"`
	// The name of the anomalous event type. Valid values:
	//
	// 	- 01: anomalous permission usage
	//
	// 	- 02: anomalous data flow
	//
	// 	- 03: anomalous data operation
	//
	// example:
	//
	// 02
	TypeCode *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	// The ID of the account that triggered the anomalous event.
	//
	// example:
	//
	// 1978132506596***
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the RAM user.
	//
	// example:
	//
	// name
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The risk level of the alert that is triggered. Valid values:
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
	//
	// example:
	//
	// 1
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) SetCurrentPage(v int32) *DescribeEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventsRequest) SetDealUserId(v string) *DescribeEventsRequest {
	s.DealUserId = &v
	return s
}

func (s *DescribeEventsRequest) SetEndTime(v string) *DescribeEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsRequest) SetId(v int64) *DescribeEventsRequest {
	s.Id = &v
	return s
}

func (s *DescribeEventsRequest) SetInstanceName(v string) *DescribeEventsRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeEventsRequest) SetLang(v string) *DescribeEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventsRequest) SetPageSize(v int32) *DescribeEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsRequest) SetProductCode(v string) *DescribeEventsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeEventsRequest) SetStartTime(v string) *DescribeEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventsRequest) SetStatus(v string) *DescribeEventsRequest {
	s.Status = &v
	return s
}

func (s *DescribeEventsRequest) SetSubTypeCode(v string) *DescribeEventsRequest {
	s.SubTypeCode = &v
	return s
}

func (s *DescribeEventsRequest) SetTargetProductCode(v string) *DescribeEventsRequest {
	s.TargetProductCode = &v
	return s
}

func (s *DescribeEventsRequest) SetTypeCode(v string) *DescribeEventsRequest {
	s.TypeCode = &v
	return s
}

func (s *DescribeEventsRequest) SetUserId(v int64) *DescribeEventsRequest {
	s.UserId = &v
	return s
}

func (s *DescribeEventsRequest) SetUserName(v string) *DescribeEventsRequest {
	s.UserName = &v
	return s
}

func (s *DescribeEventsRequest) SetWarnLevel(v int32) *DescribeEventsRequest {
	s.WarnLevel = &v
	return s
}

type DescribeEventsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of the anomalous events.
	Items []*DescribeEventsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeEventsResponseBody) SetItems(v []*DescribeEventsResponseBodyItems) *DescribeEventsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeEventsResponseBody) SetPageSize(v int32) *DescribeEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsResponseBody) SetRequestId(v string) *DescribeEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsResponseBody) SetTotalCount(v int32) *DescribeEventsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEventsResponseBodyItems struct {
	// The time when an alert was triggered for the anomalous event. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 154529000
	AlertTime *int64 `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	// Indicates whether the detection of anomalous events is enhanced. If the detection of anomalous events is enhanced, the detection accuracy and the rate of triggering alerts for anomalous events are improved. Valid values:
	//
	// 	- true: yes
	//
	// 	- false: no
	//
	// example:
	//
	// false
	Backed *bool `json:"Backed,omitempty" xml:"Backed,omitempty"`
	// The display name of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// yundunsr
	DealDisplayName *string `json:"DealDisplayName,omitempty" xml:"DealDisplayName,omitempty"`
	// The username of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// det1111
	DealLoginName *string `json:"DealLoginName,omitempty" xml:"DealLoginName,omitempty"`
	// The time when the anomalous event was handled. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 12223300
	DealTime *int64 `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// The ID of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// 229157443385014***
	DealUserId *int64 `json:"DealUserId,omitempty" xml:"DealUserId,omitempty"`
	// The display name of the account that triggered the anomalous event.
	//
	// example:
	//
	// yundunsr
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the anomalous event occurred. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545829129000
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// The ID of the anomalous event.
	//
	// example:
	//
	// 42233335555
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username of the account that triggered the anomalous event.
	//
	// example:
	//
	// det1111
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// The name of the service in which the anomalous event was detected.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The handling status for the anomalous event. Valid values:
	//
	// 	- 0: unhandled
	//
	// 	- 1: confirmed
	//
	// 	- 2: marked as false positive
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the handling status for the anomalous event.
	//
	// example:
	//
	// Pending
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// The code of the anomalous event subtype.
	//
	// example:
	//
	// 020008
	SubTypeCode *string `json:"SubTypeCode,omitempty" xml:"SubTypeCode,omitempty"`
	// The name of the anomalous event subtype.
	//
	// example:
	//
	// Anomalous volume of downloaded data
	SubTypeName *string `json:"SubTypeName,omitempty" xml:"SubTypeName,omitempty"`
	// The name of the destination service in an anomalous data flow.
	//
	// example:
	//
	// RDS
	TargetProductCode *string `json:"TargetProductCode,omitempty" xml:"TargetProductCode,omitempty"`
	// The code of the anomalous event type.
	//
	// example:
	//
	// 02
	TypeCode *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	// The name of the anomalous event type.
	//
	// example:
	//
	// Anomalous data flow
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// The ID of the account that triggered the anomalous event.
	//
	// example:
	//
	// 1978132506596***
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The severity of the anomalous event.
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
	//
	// example:
	//
	// 2
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s DescribeEventsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyItems) SetAlertTime(v int64) *DescribeEventsResponseBodyItems {
	s.AlertTime = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetBacked(v bool) *DescribeEventsResponseBodyItems {
	s.Backed = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealDisplayName(v string) *DescribeEventsResponseBodyItems {
	s.DealDisplayName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealLoginName(v string) *DescribeEventsResponseBodyItems {
	s.DealLoginName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealTime(v int64) *DescribeEventsResponseBodyItems {
	s.DealTime = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealUserId(v int64) *DescribeEventsResponseBodyItems {
	s.DealUserId = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDisplayName(v string) *DescribeEventsResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetEventTime(v int64) *DescribeEventsResponseBodyItems {
	s.EventTime = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetId(v int64) *DescribeEventsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetLoginName(v string) *DescribeEventsResponseBodyItems {
	s.LoginName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetProductCode(v string) *DescribeEventsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetStatus(v int32) *DescribeEventsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetStatusName(v string) *DescribeEventsResponseBodyItems {
	s.StatusName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetSubTypeCode(v string) *DescribeEventsResponseBodyItems {
	s.SubTypeCode = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetSubTypeName(v string) *DescribeEventsResponseBodyItems {
	s.SubTypeName = &v
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

func (s *DescribeEventsResponseBodyItems) SetTypeName(v string) *DescribeEventsResponseBodyItems {
	s.TypeName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetUserId(v int64) *DescribeEventsResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetWarnLevel(v int32) *DescribeEventsResponseBodyItems {
	s.WarnLevel = &v
	return s
}

type DescribeEventsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeEventsResponse) SetStatusCode(v int32) *DescribeEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventsResponse) SetBody(v *DescribeEventsResponseBody) *DescribeEventsResponse {
	s.Body = v
	return s
}

type DescribeIdentifyTaskStatusRequest struct {
	// Task ID, obtained from the ID field in the response after calling CreateScanTask or ScanOssObjectV1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 268
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Language type for request and response messages, default is **zh_cn**. Values:
	//
	// - **zh_cn**: Chinese (Simplified)
	//
	// - **en_us**: English (United States)
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeIdentifyTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIdentifyTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeIdentifyTaskStatusRequest) SetId(v int64) *DescribeIdentifyTaskStatusRequest {
	s.Id = &v
	return s
}

func (s *DescribeIdentifyTaskStatusRequest) SetLang(v string) *DescribeIdentifyTaskStatusRequest {
	s.Lang = &v
	return s
}

type DescribeIdentifyTaskStatusResponseBody struct {
	// -
	Content *DescribeIdentifyTaskStatusResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of this call request, a unique identifier generated by Alibaba Cloud for this request, which can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// 71064826-726F-4ADA-B879-05D8055476FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIdentifyTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIdentifyTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIdentifyTaskStatusResponseBody) SetContent(v *DescribeIdentifyTaskStatusResponseBodyContent) *DescribeIdentifyTaskStatusResponseBody {
	s.Content = v
	return s
}

func (s *DescribeIdentifyTaskStatusResponseBody) SetRequestId(v string) *DescribeIdentifyTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIdentifyTaskStatusResponseBodyContent struct {
	// 任务运行状态。
	//
	// - 0：未开始
	//
	// - 1：运行中
	//
	// - 2：已暂停
	//
	// - 3：已终止
	//
	// - 200：已完成
	//
	// example:
	//
	// 200
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIdentifyTaskStatusResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeIdentifyTaskStatusResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeIdentifyTaskStatusResponseBodyContent) SetStatus(v int32) *DescribeIdentifyTaskStatusResponseBodyContent {
	s.Status = &v
	return s
}

type DescribeIdentifyTaskStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIdentifyTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIdentifyTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIdentifyTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeIdentifyTaskStatusResponse) SetHeaders(v map[string]*string) *DescribeIdentifyTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeIdentifyTaskStatusResponse) SetStatusCode(v int32) *DescribeIdentifyTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIdentifyTaskStatusResponse) SetBody(v *DescribeIdentifyTaskStatusResponseBody) *DescribeIdentifyTaskStatusResponse {
	s.Body = v
	return s
}

type DescribeInstanceSourcesRequest struct {
	// Specifies whether to enable the security audit feature. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Specifies whether DSC is authorized to access the data asset.
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	AuthStatus *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the database engine. Valid values: **MySQL, MariaDB, Oracle, PostgreSQL, and SQLServer**.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// instance-demo-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese (default)
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service to which the data asset to query belongs. Valid values: **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data asset to query belongs. Valid values: **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The content based on which a fuzzy search is performed.
	//
	// example:
	//
	// 1
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The data asset type based on which a fuzzy search is performed.
	//
	// 	- **InstanceId**: the ID of the instance.
	//
	// 	- **InstanceName**: the name of the instance.
	//
	// 	- **DatabaseName**: the name of the database.
	//
	// example:
	//
	// InstanceId
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	// The region in which the data asset resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/214257.html).
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
}

func (s DescribeInstanceSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSourcesRequest) SetAuditStatus(v int32) *DescribeInstanceSourcesRequest {
	s.AuditStatus = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetAuthStatus(v int32) *DescribeInstanceSourcesRequest {
	s.AuthStatus = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetCurrentPage(v int32) *DescribeInstanceSourcesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetEngineType(v string) *DescribeInstanceSourcesRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetFeatureType(v int32) *DescribeInstanceSourcesRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetInstanceId(v string) *DescribeInstanceSourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetLang(v string) *DescribeInstanceSourcesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetPageSize(v int32) *DescribeInstanceSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetProductCode(v string) *DescribeInstanceSourcesRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetProductId(v int64) *DescribeInstanceSourcesRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetSearchKey(v string) *DescribeInstanceSourcesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetSearchType(v string) *DescribeInstanceSourcesRequest {
	s.SearchType = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetServiceRegionId(v string) *DescribeInstanceSourcesRequest {
	s.ServiceRegionId = &v
	return s
}

type DescribeInstanceSourcesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of the queried data assets.
	Items []*DescribeInstanceSourcesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5A7E8FB9-5011-5A90-97D9-AE587047****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeInstanceSourcesResponseBody) SetItems(v []*DescribeInstanceSourcesResponseBodyItems) *DescribeInstanceSourcesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) SetPageSize(v int32) *DescribeInstanceSourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) SetRequestId(v string) *DescribeInstanceSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) SetTotalCount(v int32) *DescribeInstanceSourcesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInstanceSourcesResponseBodyItems struct {
	// Indicates whether the security audit feature is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Indicates whether the automatic scan feature is enabled to detect sensitive data. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	AutoScan *int32 `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	// Indicates whether the username and password can be changed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	CanModifyUserName *bool `json:"CanModifyUserName,omitempty" xml:"CanModifyUserName,omitempty"`
	// The data detection status. Valid values:
	//
	// 	- **0**: The data detection is ready.
	//
	// 	- **1**: The data detection is running.
	//
	// 	- **2**: The connectivity test is in progress.
	//
	// 	- **3**: The connectivity test passed.
	//
	// 	- **4**: The connectivity test failed.
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// Indicates whether DSC has the data de-identification permissions on the data asset. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	DatamaskStatus *int32 `json:"DatamaskStatus,omitempty" xml:"DatamaskStatus,omitempty"`
	// The name of the database to which the data asset belongs.
	//
	// example:
	//
	// demo
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates whether sensitive data detection is enabled for the data asset. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of the database engine. Valid values: **MySQL, MariaDB, Oracle, PostgreSQL, and SQLServer**.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The reason for the failure.
	//
	// example:
	//
	// The password is invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the data asset was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1625587423000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The unique ID of the data asset.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// Test
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// rm-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The storage space size of the instance. This parameter is valid only if the value of the ProductId parameter is 2. Unit: bytes.
	//
	// example:
	//
	// 409600
	InstanceSize *int64 `json:"InstanceSize,omitempty" xml:"InstanceSize,omitempty"`
	// The time when the data asset was last modified. Unit: milliseconds.
	//
	// example:
	//
	// 1625587423000
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The ID of the account that is last used to modify the data asset.
	//
	// example:
	//
	// demo
	LastModifyUserId *string `json:"LastModifyUserId,omitempty" xml:"LastModifyUserId,omitempty"`
	// The retention period of raw logs. Unit: days.
	//
	// example:
	//
	// 30
	LogStoreDay *int32 `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	// Indicates whether the password is used. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	PasswordStatus *int32 `json:"PasswordStatus,omitempty" xml:"PasswordStatus,omitempty"`
	// The ID of the service to which the data asset belongs. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates OSS. The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 2
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The number of sensitive data samples. Valid values: **0**, **5**, and **10**. Unit: data entries.
	//
	// example:
	//
	// 10
	SamplingSize *int32 `json:"SamplingSize,omitempty" xml:"SamplingSize,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 11
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The name of the tenant.
	//
	// example:
	//
	// user1
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// User01
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeInstanceSourcesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSourcesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetAuditStatus(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.AuditStatus = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetAutoScan(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.AutoScan = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetCanModifyUserName(v bool) *DescribeInstanceSourcesResponseBodyItems {
	s.CanModifyUserName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetCheckStatus(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.CheckStatus = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetDatamaskStatus(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.DatamaskStatus = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetDbName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.DbName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetEnable(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.Enable = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetEngineType(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetErrorMessage(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetGmtCreate(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetId(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetInstanceDescription(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetInstanceId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetInstanceSize(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.InstanceSize = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetLastModifyTime(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.LastModifyTime = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetLastModifyUserId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.LastModifyUserId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetLogStoreDay(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.LogStoreDay = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetPasswordStatus(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.PasswordStatus = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetProductId(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetRegionId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetRegionName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.RegionName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetSamplingSize(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.SamplingSize = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetTenantId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.TenantId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetTenantName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetUserName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.UserName = &v
	return s
}

type DescribeInstanceSourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstanceSourcesResponse) SetStatusCode(v int32) *DescribeInstanceSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSourcesResponse) SetBody(v *DescribeInstanceSourcesResponseBody) *DescribeInstanceSourcesResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The keyword that is used to search for data assets. DSC searches for data assets based on the keyword that you specify in fuzzy match mode. For example, if you specify data, all data assets whose names contain data are queried.
	//
	// example:
	//
	// data
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service to which the data asset belongs, such as MaxCompute, OSS, and ApsaraDB RDS. For more information about the types of data assets from which DSC can scan for sensitive data, see [Supported data assets](https://help.aliyun.com/document_detail/212906.html).
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data asset belongs. You can call the [DescribeDataAssets](~~DescribeDataAssets~~) operation to query the ID of the service.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level ID of the data asset. A higher sensitivity level indicates that the identified data is more sensitive. Valid values:
	//
	// 	- **1**: No sensitive data is identified.
	//
	// 	- **2**: sensitive data at level 1.
	//
	// 	- **3**: sensitive data at level 2.
	//
	// 	- **4**: sensitive data at level 3
	//
	// 	- **5**: sensitive data at level 4.
	//
	// 	- **6**: sensitive data at level 5.
	//
	// 	- **7**: sensitive data at level 6.
	//
	// 	- **8**: sensitive data at level 7.
	//
	// 	- **9**: sensitive data at level 8.
	//
	// 	- **10**: sensitive data at level 9.
	//
	// 	- **11**: sensitive data at level 10.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the sensitive data detection rule that the data asset hits. You can call the [DescribeRules](~~DescribeRules~~) operation and obtain the ID of the sensitive data detection rule from the **Id*	- response parameter.
	//
	// example:
	//
	// 1111111
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The region where the data asset resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/214257.html).
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetCurrentPage(v int32) *DescribeInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstancesRequest) SetFeatureType(v int32) *DescribeInstancesRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeInstancesRequest) SetLang(v string) *DescribeInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstancesRequest) SetName(v string) *DescribeInstancesRequest {
	s.Name = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetProductCode(v string) *DescribeInstancesRequest {
	s.ProductCode = &v
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

type DescribeInstancesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data assets.
	Items []*DescribeInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 71064826-726F-4ADA-B879-05D8055476FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of data assets.
	//
	// example:
	//
	// 231
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeInstancesResponseBody) SetItems(v []*DescribeInstancesResponseBodyItems) *DescribeInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInstancesResponseBodyItems struct {
	// The time when the data asset was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1637226782000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The name of the department to which the data asset belongs.
	//
	// example:
	//
	// ***DemoCenter
	DepartName *string `json:"DepartName,omitempty" xml:"DepartName,omitempty"`
	// The unique ID of the data asset in DSC.
	//
	// example:
	//
	// 11111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the data asset.
	//
	// example:
	//
	// Data asset Information 1
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The security status of the data asset. Valid values:
	//
	// 	- **true**: The data asset is secure.
	//
	// 	- **false**: The data asset is insecure.
	//
	// example:
	//
	// true
	Labelsec *bool `json:"Labelsec,omitempty" xml:"Labelsec,omitempty"`
	// The time when the data asset was last scanned. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1637622793000
	LastFinishTime *int64 `json:"LastFinishTime,omitempty" xml:"LastFinishTime,omitempty"`
	// If the management account has opened multiple accounts and the asset belongs to other member accounts, this field displays the UID of the member accounts.
	//
	// example:
	//
	// 12567890126
	MemberAliUid *string `json:"MemberAliUid,omitempty" xml:"MemberAliUid,omitempty"`
	// A list of tags.
	ModelTags []*DescribeInstancesResponseBodyItemsModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// The name of the data asset.
	//
	// example:
	//
	// gxdata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	OdpsRiskLevelName *string `json:"OdpsRiskLevelName,omitempty" xml:"OdpsRiskLevelName,omitempty"`
	// The Alibaba Cloud account to which the data asset belongs.
	//
	// example:
	//
	// dtdep-239-******
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the service to which the data asset belongs, such as MaxCompute, OSS, and ApsaraDB RDS. For more information about the types of data assets that DSC can scan to detect sensitive data, see [Supported data assets](https://help.aliyun.com/document_detail/212906.html).
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data asset belongs.
	//
	// example:
	//
	// 5
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The protection status of the data asset. Valid values:
	//
	// 	- **true**: The data asset is being protected.
	//
	// 	- **false**: The data asset is not protected.
	//
	// example:
	//
	// false
	Protection *bool `json:"Protection,omitempty" xml:"Protection,omitempty"`
	// The ID of the sensitivity level for the data asset. A higher sensitivity level ID indicates that the identified data is more sensitive.
	//
	// 	- **1**: No sensitive data is detected.
	//
	// 	- **2**: sensitive data at level 1.
	//
	// 	- **3**: sensitive data at level 2.
	//
	// 	- **4**: sensitive data at level 3.
	//
	// 	- **5**: sensitive data at level 4.
	//
	// 	- **6**: sensitive data at level 5.
	//
	// 	- **7**: sensitive data at level 6.
	//
	// 	- **8**: sensitive data at level 7.
	//
	// 	- **9**: sensitive data at level 8.
	//
	// 	- **10**: sensitive data at level 9.
	//
	// 	- **11**: sensitive data at level 10.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level for the data asset.
	//
	// example:
	//
	// Sensitive data at level 1
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The name of the sensitive data detection rule that the data asset hits.
	//
	// example:
	//
	// \\*\\*\\	- rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Indicates whether the data asset contains sensitive data. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The number of sensitive data objects in the data asset. For example, if the data asset is an ApsaraDB RDS instance, the value indicates the number of sensitive tables in all databases of the instance.
	//
	// example:
	//
	// 123
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The name of the tenant.
	//
	// example:
	//
	// Tenant 1
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The total number of data objects in the data asset. For example, if the data asset is an ApsaraDB RDS instance, the value indicates the total number of tables in all databases of the instance.
	//
	// example:
	//
	// 231
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeInstancesResponseBodyItems) SetDepartName(v string) *DescribeInstancesResponseBodyItems {
	s.DepartName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetId(v int64) *DescribeInstancesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetInstanceDescription(v string) *DescribeInstancesResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetLabelsec(v bool) *DescribeInstancesResponseBodyItems {
	s.Labelsec = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetLastFinishTime(v int64) *DescribeInstancesResponseBodyItems {
	s.LastFinishTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetMemberAliUid(v string) *DescribeInstancesResponseBodyItems {
	s.MemberAliUid = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetModelTags(v []*DescribeInstancesResponseBodyItemsModelTags) *DescribeInstancesResponseBodyItems {
	s.ModelTags = v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetName(v string) *DescribeInstancesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetOdpsRiskLevelName(v string) *DescribeInstancesResponseBodyItems {
	s.OdpsRiskLevelName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetOwner(v string) *DescribeInstancesResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetProductCode(v string) *DescribeInstancesResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetProductId(v string) *DescribeInstancesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetProtection(v bool) *DescribeInstancesResponseBodyItems {
	s.Protection = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetRiskLevelId(v int64) *DescribeInstancesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetRiskLevelName(v string) *DescribeInstancesResponseBodyItems {
	s.RiskLevelName = &v
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

func (s *DescribeInstancesResponseBodyItems) SetSensitiveCount(v int32) *DescribeInstancesResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetTenantName(v string) *DescribeInstancesResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeInstancesResponseBodyItems) SetTotalCount(v int32) *DescribeInstancesResponseBodyItems {
	s.TotalCount = &v
	return s
}

type DescribeInstancesResponseBodyItemsModelTags struct {
	// The ID of the tag. Valid values:
	//
	// 	- **101**: personal sensitive information
	//
	// 	- **102**: personal information
	//
	// 	- **107**: general information
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the tag. Valid values:
	//
	// 	- Personal sensitive information
	//
	// 	- Personal information
	//
	// 	- General information
	//
	// example:
	//
	// personal sensitive data
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeInstancesResponseBodyItemsModelTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyItemsModelTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyItemsModelTags) SetId(v int64) *DescribeInstancesResponseBodyItemsModelTags {
	s.Id = &v
	return s
}

func (s *DescribeInstancesResponseBodyItemsModelTags) SetName(v string) *DescribeInstancesResponseBodyItemsModelTags {
	s.Name = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstancesResponse) SetStatusCode(v int32) *DescribeInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeOssObjectDetailRequest struct {
	// The ID of the OSS object.
	//
	// >  You can call the [DescribeOssObjects](https://help.aliyun.com/document_detail/410152.html) operation to obtain the ID of the OSS object.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345213
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeOssObjectDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailRequest) SetId(v int64) *DescribeOssObjectDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeOssObjectDetailRequest) SetLang(v string) *DescribeOssObjectDetailRequest {
	s.Lang = &v
	return s
}

type DescribeOssObjectDetailResponseBody struct {
	// The details of the OSS object.
	OssObjectDetail *DescribeOssObjectDetailResponseBodyOssObjectDetail `json:"OssObjectDetail,omitempty" xml:"OssObjectDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOssObjectDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponseBody) SetOssObjectDetail(v *DescribeOssObjectDetailResponseBodyOssObjectDetail) *DescribeOssObjectDetailResponseBody {
	s.OssObjectDetail = v
	return s
}

func (s *DescribeOssObjectDetailResponseBody) SetRequestId(v string) *DescribeOssObjectDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOssObjectDetailResponseBodyOssObjectDetail struct {
	// The name of the OSS bucket to which the OSS object belongs.
	//
	// example:
	//
	// bucke***
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The type of the OSS object.
	//
	// example:
	//
	// Excel file
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The name of the OSS object.
	//
	// example:
	//
	// obj_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the OSS object.
	//
	// example:
	//
	// cn-***
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the sensitivity level for the OSS object.
	//
	// example:
	//
	// Medium sensitivity level
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// A list of the sensitive data detection rules that the OSS object hits.
	RuleList []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetail) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetBucketName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.BucketName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetCategoryName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.CategoryName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetRegionId(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.RegionId = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetRiskLevelName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetRuleList(v []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.RuleList = v
	return s
}

type DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList struct {
	// The type of the OSS object.
	//
	// example:
	//
	// Excel file
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The number of times that the OSS object hits the sensitive data detection rule.
	//
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// A list of tags for data that hits the recognition model.
	ModelTags []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// The ID of the sensitivity level of the OSS object.
	//
	// 	- **1**: No sensitive data is detected.
	//
	// 	- **2**: indicates the low sensitivity level.
	//
	// 	- **3**: indicates the medium sensitivity level.
	//
	// 	- **4**: indicates the high sensitivity level.
	//
	// 	- **5**: indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level for the OSS object.
	//
	// example:
	//
	// Medium sensitivity level
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The name of the sensitive data detection rule.
	//
	// example:
	//
	// \\*\\*\\	- rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetCategoryName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetCount(v int64) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.Count = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetModelTags(v []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.ModelTags = v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetRiskLevelId(v int64) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetRiskLevelName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetRuleName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.RuleName = &v
	return s
}

type DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags struct {
	// The tag ID.
	//
	// 	- **101**: sensitive personal information
	//
	// 	- **102**: personal information
	//
	// 	- **103**: important information
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tag name.
	//
	// 	- Sensitive personal information
	//
	// 	- Personal information
	//
	// 	- Important information
	//
	// example:
	//
	// personal sensitive data
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) SetId(v int64) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags {
	s.Id = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) SetName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags {
	s.Name = &v
	return s
}

type DescribeOssObjectDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssObjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeOssObjectDetailResponse) SetStatusCode(v int32) *DescribeOssObjectDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssObjectDetailResponse) SetBody(v *DescribeOssObjectDetailResponseBody) *DescribeOssObjectDetailResponse {
	s.Body = v
	return s
}

type DescribeOssObjectDetailV2Request struct {
	// Bucket name.
	//
	// example:
	//
	// sddp-api-scan-demo
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The unique identifier ID of the OSS storage object.
	//
	// > Call the [DescribeOssObjects](https://help.aliyun.com/document_detail/410152.html) interface to get the ID.
	//
	// example:
	//
	// 12300
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Sets the language type for request and response messages. The default value is **zh_cn**. Values:
	//
	// - **zh_cn**: Simplified Chinese
	//
	// - **en_us**: English (US)
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The full file name of the file stored on OSS.
	//
	// example:
	//
	// dir1/test.png
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// Service region ID, i.e., the region ID where the Bucket is located.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// Industry template ID.
	//
	// > You can obtain the industry template ID by calling the [DescribeCategoryTemplateList](https://help.aliyun.com/document_detail/2399296.html) interface.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeOssObjectDetailV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailV2Request) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailV2Request) SetBucketName(v string) *DescribeOssObjectDetailV2Request {
	s.BucketName = &v
	return s
}

func (s *DescribeOssObjectDetailV2Request) SetId(v string) *DescribeOssObjectDetailV2Request {
	s.Id = &v
	return s
}

func (s *DescribeOssObjectDetailV2Request) SetLang(v string) *DescribeOssObjectDetailV2Request {
	s.Lang = &v
	return s
}

func (s *DescribeOssObjectDetailV2Request) SetObjectKey(v string) *DescribeOssObjectDetailV2Request {
	s.ObjectKey = &v
	return s
}

func (s *DescribeOssObjectDetailV2Request) SetServiceRegionId(v string) *DescribeOssObjectDetailV2Request {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeOssObjectDetailV2Request) SetTemplateId(v int64) *DescribeOssObjectDetailV2Request {
	s.TemplateId = &v
	return s
}

type DescribeOssObjectDetailV2ResponseBody struct {
	// Detailed information about the OSS storage object.
	OssObjectDetail *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail `json:"OssObjectDetail,omitempty" xml:"OssObjectDetail,omitempty" type:"Struct"`
	// The ID of this call request, which is a unique identifier generated by Alibaba Cloud for the request and can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOssObjectDetailV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailV2ResponseBody) SetOssObjectDetail(v *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) *DescribeOssObjectDetailV2ResponseBody {
	s.OssObjectDetail = v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBody) SetRequestId(v string) *DescribeOssObjectDetailV2ResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOssObjectDetailV2ResponseBodyOssObjectDetail struct {
	// The name of the Bucket to which the OSS storage object belongs.
	//
	// example:
	//
	// sddp-api-scan-demo
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The major category of the model.
	//
	// example:
	//
	// Excel file
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// File type.
	//
	// example:
	//
	// Text file
	FileCategoryName *string `json:"FileCategoryName,omitempty" xml:"FileCategoryName,omitempty"`
	// The unique ID of the OSS object.
	//
	// example:
	//
	// 1757262735738932224
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Data labels, separated by commas. Values:
	//
	// - **101**: Personal Sensitive Information.
	//
	// - **102**: Personal Information.
	//
	// - **107**: General Information.
	//
	// example:
	//
	// 101,102
	ModelTagIds *string `json:"ModelTagIds,omitempty" xml:"ModelTagIds,omitempty"`
	// OSS storage object name.
	//
	// example:
	//
	// dir1/test.png
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// OSS Object ACL
	//
	// example:
	//
	// private
	ObjectAcl *string `json:"ObjectAcl,omitempty" xml:"ObjectAcl,omitempty"`
	// The region ID to which the OSS storage object belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The sensitivity level of the OSS object. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 4
	RiskLevelId *int32 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The risk level name of the OSS storage object.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// A list of sensitive data recognition rules hit by the OSS storage object.
	RuleList []*DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// File size. Unit: Byte.
	//
	// example:
	//
	// 1024
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetBucketName(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.BucketName = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetCategoryName(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.CategoryName = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetFileCategoryName(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.FileCategoryName = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetId(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.Id = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetModelTagIds(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.ModelTagIds = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetName(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetObjectAcl(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.ObjectAcl = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetRegionId(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.RegionId = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetRiskLevelId(v int32) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetRiskLevelName(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetRuleList(v []*DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.RuleList = v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail) SetSize(v int64) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetail {
	s.Size = &v
	return s
}

type DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList struct {
	// The major category of the model.
	//
	// example:
	//
	// Excel
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The number of times the sensitive data recognition rule was hit.
	//
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// List of data tags.
	ModelTags []*DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleListModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// The risk level ID of the OSS storage object.
	//
	// - **1**: No sensitive data detected.
	//
	// - **2**: Level 1 sensitive data.
	//
	// - **3**: Level 2 sensitive data.
	//
	// - **4**: Level 3 sensitive data.
	//
	// - **5**: Level 4 sensitive data.
	//
	// example:
	//
	// 3
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The risk level name of the OSS storage object.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The name of the sensitive data recognition rule that was hit.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList) SetCategoryName(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList) SetCount(v int64) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList {
	s.Count = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList) SetModelTags(v []*DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleListModelTags) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList {
	s.ModelTags = v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList) SetRiskLevelId(v int64) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList) SetRiskLevelName(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList) SetRuleName(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleList {
	s.RuleName = &v
	return s
}

type DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleListModelTags struct {
	// ID of the data label for the recognition model.
	//
	// - **101**: Personal sensitive information.
	//
	// - **102**: Personal information.
	//
	// - **103**: Important data.
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Name of the data label for the recognition model.
	//
	// - Personal sensitive information.
	//
	// - Personal information.
	//
	// - Important data.
	//
	// example:
	//
	// personal sensitive data
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleListModelTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleListModelTags) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleListModelTags) SetId(v int64) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleListModelTags {
	s.Id = &v
	return s
}

func (s *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleListModelTags) SetName(v string) *DescribeOssObjectDetailV2ResponseBodyOssObjectDetailRuleListModelTags {
	s.Name = &v
	return s
}

type DescribeOssObjectDetailV2Response struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssObjectDetailV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssObjectDetailV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectDetailV2Response) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailV2Response) SetHeaders(v map[string]*string) *DescribeOssObjectDetailV2Response {
	s.Headers = v
	return s
}

func (s *DescribeOssObjectDetailV2Response) SetStatusCode(v int32) *DescribeOssObjectDetailV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssObjectDetailV2Response) SetBody(v *DescribeOssObjectDetailV2ResponseBody) *DescribeOssObjectDetailV2Response {
	s.Body = v
	return s
}

type DescribeOssObjectsRequest struct {
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The code of the file type.
	//
	// example:
	//
	// 1
	FileCategoryCode *int64 `json:"FileCategoryCode,omitempty" xml:"FileCategoryCode,omitempty"`
	// The ID of the instance to which the OSS object belongs.
	//
	// > You can call the **DescribeInstances*	- operation to query the instance ID.
	//
	// example:
	//
	// ins-2222
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The end time of the last scan. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	LastScanTimeEnd *int64 `json:"LastScanTimeEnd,omitempty" xml:"LastScanTimeEnd,omitempty"`
	// The start time of the last scan. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	LastScanTimeStart *int64 `json:"LastScanTimeStart,omitempty" xml:"LastScanTimeStart,omitempty"`
	// When you query data by page, use the `Marker` parameter to query the data that follows the `Marker` value.
	//
	// example:
	//
	// 1754786235714378752
	Marker *int64 `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The search keyword. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sensitivity level of the OSS object. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int32 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the sensitive data detection rule that the OSS object hits.
	//
	// > You can call the **DescribeRules*	- operation to query the ID of the sensitive data detection rule.
	//
	// example:
	//
	// 1222
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The region in which the data asset resides.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the industry-specific rule template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeOssObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsRequest) SetCurrentPage(v int32) *DescribeOssObjectsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetFileCategoryCode(v int64) *DescribeOssObjectsRequest {
	s.FileCategoryCode = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetInstanceId(v string) *DescribeOssObjectsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetLang(v string) *DescribeOssObjectsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetLastScanTimeEnd(v int64) *DescribeOssObjectsRequest {
	s.LastScanTimeEnd = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetLastScanTimeStart(v int64) *DescribeOssObjectsRequest {
	s.LastScanTimeStart = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetMarker(v int64) *DescribeOssObjectsRequest {
	s.Marker = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetName(v string) *DescribeOssObjectsRequest {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetPageSize(v int32) *DescribeOssObjectsRequest {
	s.PageSize = &v
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

func (s *DescribeOssObjectsRequest) SetServiceRegionId(v string) *DescribeOssObjectsRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeOssObjectsRequest) SetTemplateId(v int64) *DescribeOssObjectsRequest {
	s.TemplateId = &v
	return s
}

type DescribeOssObjectsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The OSS objects.
	Items []*DescribeOssObjectsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// This parameter is deprecated.
	//
	// example:
	//
	// -1
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The ID value from which the next page of results starts.
	//
	// >  This parameter is returned only when the `Truncated` parameter is set to `true`.
	//
	// example:
	//
	// 1754786235714378752
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Truncated *bool `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
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

func (s *DescribeOssObjectsResponseBody) SetItems(v []*DescribeOssObjectsResponseBodyItems) *DescribeOssObjectsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetMarker(v string) *DescribeOssObjectsResponseBody {
	s.Marker = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetNextMarker(v string) *DescribeOssObjectsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetPageSize(v int32) *DescribeOssObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetRequestId(v string) *DescribeOssObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetTotalCount(v int32) *DescribeOssObjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetTruncated(v bool) *DescribeOssObjectsResponseBody {
	s.Truncated = &v
	return s
}

type DescribeOssObjectsResponseBodyItems struct {
	// The name of the bucket.
	//
	// example:
	//
	// oss-duplicate-***
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The type of the OSS object. Valid values include **900001**, **800015**, or **800005**, which indicates the MP4 file, PDF file, or OSS configuration file, respectively.
	//
	// example:
	//
	// 900001
	Category *int64 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The name of the file type.
	//
	// example:
	//
	// MP4 file
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The code of the file type.
	//
	// example:
	//
	// 1
	FileCategoryCode *int64 `json:"FileCategoryCode,omitempty" xml:"FileCategoryCode,omitempty"`
	// The name of the file type.
	//
	// example:
	//
	// text file
	FileCategoryName *string `json:"FileCategoryName,omitempty" xml:"FileCategoryName,omitempty"`
	// The file ID of the OSS object.
	//
	// example:
	//
	// file-22***
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the OSS object.
	//
	// example:
	//
	// 17383
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the instance to which the OSS object belongs.
	//
	// example:
	//
	// 1232122
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the file was last modified.
	//
	// example:
	//
	// 1536751124000
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The name of the OSS object.
	//
	// example:
	//
	// obj_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the OSS object.
	//
	// example:
	//
	// cn-***
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the sensitivity level of the OSS object. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level for the OSS object.
	//
	// example:
	//
	// Medium sensitivity level
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The number of rules that are hit.
	//
	// example:
	//
	// 100
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// The rules.
	RuleList []*DescribeOssObjectsResponseBodyItemsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// The number of fields that are hit.
	//
	// example:
	//
	// 50
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeOssObjectsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsResponseBodyItems) SetBucketName(v string) *DescribeOssObjectsResponseBodyItems {
	s.BucketName = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetCategory(v int64) *DescribeOssObjectsResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetCategoryName(v string) *DescribeOssObjectsResponseBodyItems {
	s.CategoryName = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetFileCategoryCode(v int64) *DescribeOssObjectsResponseBodyItems {
	s.FileCategoryCode = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetFileCategoryName(v string) *DescribeOssObjectsResponseBodyItems {
	s.FileCategoryName = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetFileId(v string) *DescribeOssObjectsResponseBodyItems {
	s.FileId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetId(v string) *DescribeOssObjectsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetInstanceId(v int64) *DescribeOssObjectsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetLastModifiedTime(v int64) *DescribeOssObjectsResponseBodyItems {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetName(v string) *DescribeOssObjectsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRegionId(v string) *DescribeOssObjectsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRiskLevelId(v int64) *DescribeOssObjectsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRiskLevelName(v string) *DescribeOssObjectsResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRuleCount(v int32) *DescribeOssObjectsResponseBodyItems {
	s.RuleCount = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRuleList(v []*DescribeOssObjectsResponseBodyItemsRuleList) *DescribeOssObjectsResponseBodyItems {
	s.RuleList = v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetSensitiveCount(v int32) *DescribeOssObjectsResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetSize(v int64) *DescribeOssObjectsResponseBodyItems {
	s.Size = &v
	return s
}

type DescribeOssObjectsResponseBodyItemsRuleList struct {
	// The number of times that the rule is hit.
	//
	// example:
	//
	// 100
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The search keyword. Fuzzy match is supported.
	//
	// example:
	//
	// ID card
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the sensitivity level of the OSS object. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
}

func (s DescribeOssObjectsResponseBodyItemsRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssObjectsResponseBodyItemsRuleList) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) SetCount(v int64) *DescribeOssObjectsResponseBodyItemsRuleList {
	s.Count = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) SetName(v string) *DescribeOssObjectsResponseBodyItemsRuleList {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) SetRiskLevelId(v int64) *DescribeOssObjectsResponseBodyItemsRuleList {
	s.RiskLevelId = &v
	return s
}

type DescribeOssObjectsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeOssObjectsResponse) SetStatusCode(v int32) *DescribeOssObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssObjectsResponse) SetBody(v *DescribeOssObjectsResponseBody) *DescribeOssObjectsResponse {
	s.Body = v
	return s
}

type DescribePackagesRequest struct {
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the instance to which the package belongs.
	//
	// > You can call the **DescribeInstances*	- operation to query the ID of the instance.
	//
	// example:
	//
	// 12321
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The search keyword. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the service to which the package belongs.
	//
	// > You can call the **DescribeDataAssets*	- operation to query the ID of the service.
	//
	// example:
	//
	// 2566600
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the package. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the sensitive data detection rule that the package hits.
	//
	// > You can call the **DescribeRules*	- operation to query the ID of the sensitive data detection rule.
	//
	// example:
	//
	// 266666
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribePackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribePackagesRequest) SetCurrentPage(v int32) *DescribePackagesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePackagesRequest) SetInstanceId(v int64) *DescribePackagesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackagesRequest) SetLang(v string) *DescribePackagesRequest {
	s.Lang = &v
	return s
}

func (s *DescribePackagesRequest) SetName(v string) *DescribePackagesRequest {
	s.Name = &v
	return s
}

func (s *DescribePackagesRequest) SetPageSize(v int32) *DescribePackagesRequest {
	s.PageSize = &v
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

type DescribePackagesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of the information about the packages.
	Items []*DescribePackagesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribePackagesResponseBody) SetItems(v []*DescribePackagesResponseBodyItems) *DescribePackagesResponseBody {
	s.Items = v
	return s
}

func (s *DescribePackagesResponseBody) SetPageSize(v int32) *DescribePackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePackagesResponseBody) SetRequestId(v string) *DescribePackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackagesResponseBody) SetTotalCount(v int32) *DescribePackagesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePackagesResponseBodyItems struct {
	// The point in time when the MaxCompute package was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the package.
	//
	// example:
	//
	// 111111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the instance to which the package belongs.
	//
	// example:
	//
	// 223453332
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the package.
	//
	// example:
	//
	// gxdata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account of the user that owns the package.
	//
	// example:
	//
	// cou-2221
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The sensitivity level of the package. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 4
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level for the package.
	//
	// example:
	//
	// Highest sensitivity level
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// Indicates whether the package contains sensitive data. Valid values:
	//
	// 	- true: yes
	//
	// 	- false: no
	//
	// example:
	//
	// true
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The total volume of sensitive data in the package. For example, the value can be the total number of sensitive tables in the MaxCompute package.
	//
	// example:
	//
	// 123
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The total volume of data in the package. For example, the value can be the total number of tables in the MaxCompute package.
	//
	// example:
	//
	// 321
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribePackagesResponseBodyItems) SetId(v int64) *DescribePackagesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetInstanceId(v int64) *DescribePackagesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetName(v string) *DescribePackagesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetOwner(v string) *DescribePackagesResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetRiskLevelId(v int64) *DescribePackagesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetRiskLevelName(v string) *DescribePackagesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetSensitive(v bool) *DescribePackagesResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetSensitiveCount(v int32) *DescribePackagesResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribePackagesResponseBodyItems) SetTotalCount(v int32) *DescribePackagesResponseBodyItems {
	s.TotalCount = &v
	return s
}

type DescribePackagesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribePackagesResponse) SetStatusCode(v int32) *DescribePackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackagesResponse) SetBody(v *DescribePackagesResponseBody) *DescribePackagesResponse {
	s.Body = v
	return s
}

type DescribeParentInstanceRequest struct {
	// Authorization status of the data asset instance.
	//
	// - **0**: Unauthorized
	//
	// - **1**: Authorized
	//
	// example:
	//
	// 0
	AuthStatus *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// Connection status of the instance or the database under the instance. Values:
	//
	// - **-3**: Database not created
	//
	// - **-2**: Released
	//
	// - **-1**: Not connected
	//
	// - **2**: Connectivity test in progress
	//
	// - **3**: Connected
	//
	// - **4**: Connection failed
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// Instance status.
	//
	// - **Running**: Running
	//
	// - **Released**: Released
	//
	// - **DatabaseNotCreated**: Database not created
	//
	// example:
	//
	// Running
	ClusterStatus *string `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// When performing a paginated query, set the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Database name.
	//
	// example:
	//
	// db_**t
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Engine type. Values:
	//
	// - **MySQL**
	//
	// - **MariaDB**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The instance ID to which the data in the data asset table belongs.
	//
	// example:
	//
	// rm-*******xx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Language type for request and response messages. Values:
	//
	// - **zh_cn**: Default, Simplified Chinese
	//
	// - **en_us**: English (US)
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Member account ID.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// When performing a paginated query, set the number of rows per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. Valid values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: TableStore
	//
	// - **5**: RDS
	//
	// - **6**: SelfDB
	//
	// - **7**: PolarDB-X
	//
	// - **8**: PolarDB
	//
	// - **9**: ADB-PG
	//
	// - **10**: OceanBase
	//
	// - **11**: MongoDB
	//
	// - **25**: Redis
	//
	// example:
	//
	// 5
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The region where the asset is located. Values:
	//
	// - **cn-beijing**: China (Beijing)
	//
	// - **cn-zhangjiakou**: China (Zhangjiakou)
	//
	// - **cn-huhehaote**: China (Hohhot)
	//
	// - **cn-hangzhou**: China (Hangzhou)
	//
	// - **cn-shanghai**: China (Shanghai)
	//
	// - **cn-shenzhen**: China (Shenzhen)
	//
	// - **cn-hongkong**:  China (Hong Kong)
	//
	// example:
	//
	// cn-shanghai
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
}

func (s DescribeParentInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeParentInstanceRequest) SetAuthStatus(v int32) *DescribeParentInstanceRequest {
	s.AuthStatus = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetCheckStatus(v int32) *DescribeParentInstanceRequest {
	s.CheckStatus = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetClusterStatus(v string) *DescribeParentInstanceRequest {
	s.ClusterStatus = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetCurrentPage(v int32) *DescribeParentInstanceRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetDbName(v string) *DescribeParentInstanceRequest {
	s.DbName = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetEngineType(v string) *DescribeParentInstanceRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetInstanceId(v string) *DescribeParentInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetLang(v string) *DescribeParentInstanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetMemberAccount(v int64) *DescribeParentInstanceRequest {
	s.MemberAccount = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetPageSize(v int32) *DescribeParentInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetResourceType(v int64) *DescribeParentInstanceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetServiceRegionId(v string) *DescribeParentInstanceRequest {
	s.ServiceRegionId = &v
	return s
}

type DescribeParentInstanceResponseBody struct {
	// When performing a paginated query, set the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The assets.
	Items []*DescribeParentInstanceResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// When performing a paginated query, set the maximum number of data asset instances displayed per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID of the result.
	//
	// example:
	//
	// ACEF9334-BB50-525D-8CF3-6CF504E4D1B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of data items in the result.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeParentInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParentInstanceResponseBody) SetCurrentPage(v int32) *DescribeParentInstanceResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeParentInstanceResponseBody) SetItems(v []*DescribeParentInstanceResponseBodyItems) *DescribeParentInstanceResponseBody {
	s.Items = v
	return s
}

func (s *DescribeParentInstanceResponseBody) SetPageSize(v int32) *DescribeParentInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeParentInstanceResponseBody) SetRequestId(v string) *DescribeParentInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParentInstanceResponseBody) SetTotalCount(v int32) *DescribeParentInstanceResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeParentInstanceResponseBodyItems struct {
	// Audit authorization status. The values are as follows:
	//
	// - **1**: Authorized
	//
	// - **0**: Unauthorized
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Authorization status of the data asset instance.
	//
	// - **0**: Unauthorized
	//
	// - **1**: Authorized
	//
	// example:
	//
	// 1
	AuthStatus *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// Instance authorization time, in milliseconds.
	//
	// example:
	//
	// 1719882941000
	AuthTime *int64 `json:"AuthTime,omitempty" xml:"AuthTime,omitempty"`
	// Instance status.
	//
	// example:
	//
	// Running
	ClusterStatus *string `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// Connection node type, valid only for MongoDB assets.
	//
	// example:
	//
	// Primary
	ConnectNode *string `json:"ConnectNode,omitempty" xml:"ConnectNode,omitempty"`
	// Number of databases under the instance.
	//
	// example:
	//
	// 3
	DbNum *string `json:"DbNum,omitempty" xml:"DbNum,omitempty"`
	// The engine type. Valid values:
	//
	// - **MySQL**
	//
	// - **MariaDB**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// Description of the instance.
	//
	// example:
	//
	// instance description
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// rm-*******t2vz
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance space size, valid only for OSS assets. Unit: bytes.
	//
	// example:
	//
	// 409600
	InstanceSize *int64 `json:"InstanceSize,omitempty" xml:"InstanceSize,omitempty"`
	// Region name. The values are as follows:
	//
	// - **China (Hangzhou)**
	//
	// - **China (Shanghai)**
	//
	// - **China (Beijing)**
	//
	// - **China (Zhangjiakou)**
	//
	// - **China (Shenzhen)**
	//
	// - **China (Guangzhou)**
	//
	// - **China (Hong Kong)**
	//
	// - **Singapore**
	//
	// - **US (Silicon Valley)**
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// Member account ID.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// Identifier for the authorized asset. For structured data, it is identified by `instanceID.databaseName`.
	//
	// example:
	//
	// rm-******xxx.**st
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The region in which the asset resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Asset type name. The values are as follows:
	//
	// - **MaxCompute**
	//
	// - **OSS**
	//
	// - **ADB-MYSQL**
	//
	// - **TableStore**
	//
	// - **RDS**
	//
	// - **SelfDB**
	//
	// - **PolarDB-X**
	//
	// - **PolarDB**
	//
	// - **ADB-PG**
	//
	// - **OceanBase**
	//
	// - **MongoDB**
	//
	// - **Redis**
	//
	// example:
	//
	// RDS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Supported connection nodes, separated by commas.
	//
	// example:
	//
	// Primary,Secondary
	SupportConnectNodes *string `json:"SupportConnectNodes,omitempty" xml:"SupportConnectNodes,omitempty"`
	// Tenant ID, valid only for OceanBase assets.
	//
	// example:
	//
	// HB***-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Tenant name, valid only for OceanBase assets.
	//
	// example:
	//
	// user1
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// Number of unconnected databases under the instance.
	//
	// example:
	//
	// 1
	UnConnectDbCount *string `json:"UnConnectDbCount,omitempty" xml:"UnConnectDbCount,omitempty"`
	// Reason for not supporting one-click authorization.
	//
	// example:
	//
	// engine type not support
	UnSupportOneClickAuthReason *string `json:"UnSupportOneClickAuthReason,omitempty" xml:"UnSupportOneClickAuthReason,omitempty"`
}

func (s DescribeParentInstanceResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentInstanceResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeParentInstanceResponseBodyItems) SetAuditStatus(v int32) *DescribeParentInstanceResponseBodyItems {
	s.AuditStatus = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetAuthStatus(v int32) *DescribeParentInstanceResponseBodyItems {
	s.AuthStatus = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetAuthTime(v int64) *DescribeParentInstanceResponseBodyItems {
	s.AuthTime = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetClusterStatus(v string) *DescribeParentInstanceResponseBodyItems {
	s.ClusterStatus = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetConnectNode(v string) *DescribeParentInstanceResponseBodyItems {
	s.ConnectNode = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetDbNum(v string) *DescribeParentInstanceResponseBodyItems {
	s.DbNum = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetEngineType(v string) *DescribeParentInstanceResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetInstanceDescription(v string) *DescribeParentInstanceResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetInstanceId(v string) *DescribeParentInstanceResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetInstanceSize(v int64) *DescribeParentInstanceResponseBodyItems {
	s.InstanceSize = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetLocalName(v string) *DescribeParentInstanceResponseBodyItems {
	s.LocalName = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetMemberAccount(v int64) *DescribeParentInstanceResponseBodyItems {
	s.MemberAccount = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetParentId(v string) *DescribeParentInstanceResponseBodyItems {
	s.ParentId = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetRegionId(v string) *DescribeParentInstanceResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetResourceType(v string) *DescribeParentInstanceResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetSupportConnectNodes(v string) *DescribeParentInstanceResponseBodyItems {
	s.SupportConnectNodes = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetTenantId(v string) *DescribeParentInstanceResponseBodyItems {
	s.TenantId = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetTenantName(v string) *DescribeParentInstanceResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetUnConnectDbCount(v string) *DescribeParentInstanceResponseBodyItems {
	s.UnConnectDbCount = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetUnSupportOneClickAuthReason(v string) *DescribeParentInstanceResponseBodyItems {
	s.UnSupportOneClickAuthReason = &v
	return s
}

type DescribeParentInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParentInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParentInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParentInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeParentInstanceResponse) SetHeaders(v map[string]*string) *DescribeParentInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeParentInstanceResponse) SetStatusCode(v int32) *DescribeParentInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParentInstanceResponse) SetBody(v *DescribeParentInstanceResponseBody) *DescribeParentInstanceResponse {
	s.Body = v
	return s
}

type DescribeRiskLevelsRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- zh_cn: Chinese (default)
	//
	// 	- en_us: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the industry-specific rule template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeRiskLevelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskLevelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskLevelsRequest) SetFeatureType(v int32) *DescribeRiskLevelsRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeRiskLevelsRequest) SetLang(v string) *DescribeRiskLevelsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskLevelsRequest) SetTemplateId(v int64) *DescribeRiskLevelsRequest {
	s.TemplateId = &v
	return s
}

type DescribeRiskLevelsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 136082B3-B21F-5E9D-B68E-991FFD205D24
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of sensitivity levels.
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
	// The description of the sensitivity level. You can enter a custom description.
	//
	// The following list describes the sensitivity level names and the corresponding default description:
	//
	// 	- **NA**: which indicates that no sensitive data is detected.
	//
	// 	- **S1**: which indicates the sensitive data at sensitivity level 1.
	//
	// 	- **S2**: which indicates the sensitive data at sensitivity level 2.
	//
	// 	- **S3**: which indicates the sensitive data at sensitivity level 3.
	//
	// 	- **S4**: which indicates the sensitive data at sensitivity level 4.
	//
	// 	- **S5**: which indicates the sensitive data at sensitivity level 5.
	//
	// 	- **S6**: which indicates the sensitive data at sensitivity level 6.
	//
	// 	- **S7**: which indicates the sensitive data at sensitivity level 7.
	//
	// 	- **S8**: which indicates the sensitive data at sensitivity level 8.
	//
	// 	- **S9**: which indicates the sensitive data at sensitivity level 9.
	//
	// 	- **S10**: which indicates the sensitive data at sensitivity level 10.
	//
	// example:
	//
	// Sensitive data at sensitivity level 1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the sensitivity level. Valid values: 1 to 11. Each sensitivity level ID maps a sensitivity level. For example, the sensitivity level ID of 2 corresponds to the sensitivity level S1.
	//
	// For more information, see the description of the Name parameter.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the sensitivity level. The highest sensitivity level varies based on rule templates. The highest sensitivity level is S10. The highest sensitivity level of the **Built-in data security classification templates for Alibaba and Ant Group*	- is S4, and that of the **built-in classification templates for financial data*	- and **built-in classification templates for assets*	- is S5. For more information about the built-in classification templates for financial data, see Guidelines for Security Classification of Financial Data and Security Data - JRT 0197-2020. For a copied rule template, the highest sensitivity level is S10. The following list describes the sensitivity level names and the corresponding IDs:
	//
	// 	- **NA**: 1
	//
	// 	- **S1**: 2
	//
	// 	- **S2**: 3
	//
	// 	- **S3**: 4
	//
	// 	- **S4**: 5
	//
	// 	- **S5**: 6
	//
	// 	- **S6**: 7
	//
	// 	- **S7**: 8
	//
	// 	- **S8**: 9
	//
	// 	- **S9**: 10
	//
	// 	- **S10**: 11
	//
	// example:
	//
	// S1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of times that each sensitivity level is referenced in the rule template. Default value: 0.
	//
	// example:
	//
	// 20
	ReferenceNum *int32 `json:"ReferenceNum,omitempty" xml:"ReferenceNum,omitempty"`
}

func (s DescribeRiskLevelsResponseBodyRiskLevelList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskLevelsResponseBodyRiskLevelList) GoString() string {
	return s.String()
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetDescription(v string) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.Description = &v
	return s
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetId(v int64) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.Id = &v
	return s
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetName(v string) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.Name = &v
	return s
}

func (s *DescribeRiskLevelsResponseBodyRiskLevelList) SetReferenceNum(v int32) *DescribeRiskLevelsResponseBodyRiskLevelList {
	s.ReferenceNum = &v
	return s
}

type DescribeRiskLevelsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskLevelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRiskLevelsResponse) SetStatusCode(v int32) *DescribeRiskLevelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskLevelsResponse) SetBody(v *DescribeRiskLevelsResponseBody) *DescribeRiskLevelsResponse {
	s.Body = v
	return s
}

type DescribeRulesRequest struct {
	// The content type of the sensitive data detection rule. Valid values:
	//
	// 	- **0**: keyword
	//
	// 	- **2**: regular expression
	//
	// example:
	//
	// 2
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The type of the content in the sensitive data detection rule. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates attempts to exploit SQL injections. The value 2 indicates bypass by using SQL injections. The value 3 indicates abuse of stored procedures. The value 4 indicates buffer overflow. The value 5 indicates SQL injections based on errors.
	//
	// example:
	//
	// 1
	ContentCategory *int32 `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	// The external cooperation channel. Valid values:
	//
	// 	- DAS
	//
	// 	- YAOCHI
	//
	// example:
	//
	// DAS
	CooperationChannel *string `json:"CooperationChannel,omitempty" xml:"CooperationChannel,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the sensitive data detection rule. Valid values:
	//
	// 	- **0**: built-in rule
	//
	// 	- **1**: custom rule
	//
	// example:
	//
	// 1
	CustomType *int32 `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The parent group type of the rule.
	//
	// example:
	//
	// 4_1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether to allow earlier versions of request parameters to support keywords that are supported in later versions of request parameters. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// > To specify keywords as the content type of the sensitive data detection rule, you can set the Category parameter to 0 for earlier versions of request parameters and set the Category parameter to 5 for later versions of request parameters. You can specify the KeywordCompatible parameter based on your business requirements.
	//
	// example:
	//
	// true
	KeywordCompatible *bool `json:"KeywordCompatible,omitempty" xml:"KeywordCompatible,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The match type. Valid values:
	//
	// 	- 1: rule-based match
	//
	// 	- 2: dictionary-based match
	//
	// example:
	//
	// 1
	MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The name of the sensitive data detection rule. Fuzzy match is supported.
	//
	// example:
	//
	// \\*\\*\\	- rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service to which the data asset belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	ProductCode *int32 `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the sensitive data detection rule is applied. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The type of the sensitive data detection rule. Valid values:
	//
	// 	- **1**: sensitive data detection rule
	//
	// 	- **2**: audit rule
	//
	// 	- **3**: anomalous event detection rule
	//
	// 	- **99**: custom rule
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// Specifies whether to query a simplified rule. The simplified rule contains only the rule name. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Simplify *bool `json:"Simplify,omitempty" xml:"Simplify,omitempty"`
	// The status of the sensitive data detection rule. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the data asset. Valid values:
	//
	// 	- **0**: all data assets
	//
	// 	- **1**: structured data asset
	//
	// 	- **2**: unstructured data asset
	//
	// > If you set the parameter to 1 or 2, rules that support all data assets and rules that support the queried data asset type are returned.
	//
	// example:
	//
	// 1
	SupportForm *int32 `json:"SupportForm,omitempty" xml:"SupportForm,omitempty"`
	// The severity level of the alert. Valid values:
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
	//
	// example:
	//
	// 2
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s DescribeRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRulesRequest) SetCategory(v int32) *DescribeRulesRequest {
	s.Category = &v
	return s
}

func (s *DescribeRulesRequest) SetContentCategory(v int32) *DescribeRulesRequest {
	s.ContentCategory = &v
	return s
}

func (s *DescribeRulesRequest) SetCooperationChannel(v string) *DescribeRulesRequest {
	s.CooperationChannel = &v
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

func (s *DescribeRulesRequest) SetFeatureType(v int32) *DescribeRulesRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeRulesRequest) SetGroupId(v string) *DescribeRulesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeRulesRequest) SetKeywordCompatible(v bool) *DescribeRulesRequest {
	s.KeywordCompatible = &v
	return s
}

func (s *DescribeRulesRequest) SetLang(v string) *DescribeRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRulesRequest) SetMatchType(v int32) *DescribeRulesRequest {
	s.MatchType = &v
	return s
}

func (s *DescribeRulesRequest) SetName(v string) *DescribeRulesRequest {
	s.Name = &v
	return s
}

func (s *DescribeRulesRequest) SetPageSize(v int32) *DescribeRulesRequest {
	s.PageSize = &v
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

func (s *DescribeRulesRequest) SetRiskLevelId(v int64) *DescribeRulesRequest {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeRulesRequest) SetRuleType(v int32) *DescribeRulesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRulesRequest) SetSimplify(v bool) *DescribeRulesRequest {
	s.Simplify = &v
	return s
}

func (s *DescribeRulesRequest) SetStatus(v int32) *DescribeRulesRequest {
	s.Status = &v
	return s
}

func (s *DescribeRulesRequest) SetSupportForm(v int32) *DescribeRulesRequest {
	s.SupportForm = &v
	return s
}

func (s *DescribeRulesRequest) SetWarnLevel(v int32) *DescribeRulesRequest {
	s.WarnLevel = &v
	return s
}

type DescribeRulesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The sensitive data detection rules.
	Items []*DescribeRulesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeRulesResponseBody) SetItems(v []*DescribeRulesResponseBodyItems) *DescribeRulesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeRulesResponseBody) SetPageSize(v int32) *DescribeRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRulesResponseBody) SetRequestId(v string) *DescribeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRulesResponseBody) SetTotalCount(v int32) *DescribeRulesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeRulesResponseBodyItems struct {
	// The content type of the sensitive data detection rule. Valid values:
	//
	// 	- **0**: keyword
	//
	// 	- **2**: regular expression
	//
	// example:
	//
	// 2
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The name of the content type of the sensitive data detection rule.
	//
	// example:
	//
	// Regular expression
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The content in the sensitive data detection rule.
	//
	// >  A built-in detection rule whose CustomType is 0 does not return the content of the rule.
	//
	// example:
	//
	// (?:\\\\D|^)((?:(?:25[0-4]|2[0-4]\\\\d|1\\\\d{2}|[1-9]\\\\d{1})\\\\.)(?:(?:25[0-5]|2[0-4]\\\\d|[01]?\\\\d?\\\\d)\\\\.){2}(?:25[0-5]|2[0-4]\\\\d|1[0-9]\\\\d|[1-9]\\\\d|[1-9]))(?:\\\\D|$)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The type of the content in the sensitive data detection rule. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates attempts to exploit SQL injections. The value 2 indicates bypass by using SQL injections. The value 3 indicates abuse of stored procedures. The value 4 indicates buffer overflow. The value 5 indicates SQL injections based on errors.
	//
	// example:
	//
	// 1
	ContentCategory *string `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	// The type of the sensitive data detection rule.
	//
	// 	- 0: built-in rule
	//
	// 	- 1: custom rule
	//
	// example:
	//
	// 1
	CustomType *int32 `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	// The description of the sensitive data detection rule.
	//
	// example:
	//
	// The sensitive data detection rule is used to detect IP addresses.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the account that is used to create the sensitive data detection rule.
	//
	// example:
	//
	// ****test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the sensitive data detection rule is created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545277010000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the sensitive data detection rule is modified. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545277010000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The parent group type of the rule.
	//
	// example:
	//
	// 4_1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of times that the sensitive data detection rule is hit.
	//
	// example:
	//
	// 3
	HitTotalCount *int32 `json:"HitTotalCount,omitempty" xml:"HitTotalCount,omitempty"`
	// The ID of the sensitive data detection rule.
	//
	// example:
	//
	// 20000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username of the account that is used to create the sensitive data detection rule.
	//
	// example:
	//
	// det1111
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// The key of the primary dimension.
	//
	// example:
	//
	// key
	MajorKey *string `json:"MajorKey,omitempty" xml:"MajorKey,omitempty"`
	// The match type. Valid values:
	//
	// 	- **1**: rule-based match
	//
	// 	- **2**: dictionary-based match
	//
	// example:
	//
	// 1
	MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The IDs of the models for sensitive data audit.
	//
	// example:
	//
	// 1452
	ModelRuleIds *string `json:"ModelRuleIds,omitempty" xml:"ModelRuleIds,omitempty"`
	// The name of the sensitive data detection rule.
	//
	// example:
	//
	// IP address
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the service to which the data asset belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the sensitive data detection rule is applied. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates OSS. The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 2
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The sensitivity level of data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **N/A**: indicates that no sensitive data is detected.
	//
	// 	- **S1**: indicates the low sensitivity level.
	//
	// 	- **S2**: indicates the medium sensitivity level.
	//
	// 	- **S3**: indicates the high sensitivity level.
	//
	// 	- **S4**: indicates the highest sensitivity level.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The statistical expression.
	//
	// example:
	//
	// 1
	StatExpress *string `json:"StatExpress,omitempty" xml:"StatExpress,omitempty"`
	// The status of the sensitive data detection rule. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The data asset type that is supported by the sensitive data detection rule. Valid values:
	//
	// 	- **0**: all data assets
	//
	// 	- **1**: structured data assets
	//
	// 	- **2**: unstructured data assets
	//
	// example:
	//
	// 2
	SupportForm *int32 `json:"SupportForm,omitempty" xml:"SupportForm,omitempty"`
	// The name of the service to which the data asset belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The IDs of the templates that are used to audit sensitive data.
	//
	// example:
	//
	// 1
	TemplateRuleIds *string `json:"TemplateRuleIds,omitempty" xml:"TemplateRuleIds,omitempty"`
	// The ID of the account that is used to create the sensitive data detection rule.
	//
	// example:
	//
	// 0
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The severity level. Valid values:
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
	//
	// example:
	//
	// 2
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s DescribeRulesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyItems) SetCategory(v int32) *DescribeRulesResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetCategoryName(v string) *DescribeRulesResponseBodyItems {
	s.CategoryName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetContent(v string) *DescribeRulesResponseBodyItems {
	s.Content = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetContentCategory(v string) *DescribeRulesResponseBodyItems {
	s.ContentCategory = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetCustomType(v int32) *DescribeRulesResponseBodyItems {
	s.CustomType = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetDescription(v string) *DescribeRulesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetDisplayName(v string) *DescribeRulesResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetGmtCreate(v int64) *DescribeRulesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetGmtModified(v int64) *DescribeRulesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetGroupId(v string) *DescribeRulesResponseBodyItems {
	s.GroupId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetHitTotalCount(v int32) *DescribeRulesResponseBodyItems {
	s.HitTotalCount = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetId(v int64) *DescribeRulesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetLoginName(v string) *DescribeRulesResponseBodyItems {
	s.LoginName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetMajorKey(v string) *DescribeRulesResponseBodyItems {
	s.MajorKey = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetMatchType(v int32) *DescribeRulesResponseBodyItems {
	s.MatchType = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetModelRuleIds(v string) *DescribeRulesResponseBodyItems {
	s.ModelRuleIds = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetName(v string) *DescribeRulesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetProductCode(v string) *DescribeRulesResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetProductId(v int64) *DescribeRulesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetRiskLevelId(v int64) *DescribeRulesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetRiskLevelName(v string) *DescribeRulesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetStatExpress(v string) *DescribeRulesResponseBodyItems {
	s.StatExpress = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetStatus(v int32) *DescribeRulesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetSupportForm(v int32) *DescribeRulesResponseBodyItems {
	s.SupportForm = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetTarget(v string) *DescribeRulesResponseBodyItems {
	s.Target = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetTemplateRuleIds(v string) *DescribeRulesResponseBodyItems {
	s.TemplateRuleIds = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetUserId(v int64) *DescribeRulesResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetWarnLevel(v int32) *DescribeRulesResponseBodyItems {
	s.WarnLevel = &v
	return s
}

type DescribeRulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRulesResponse) SetStatusCode(v int32) *DescribeRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRulesResponse) SetBody(v *DescribeRulesResponseBody) *DescribeRulesResponse {
	s.Body = v
	return s
}

type DescribeTablesRequest struct {
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the data asset to which the table belongs. You can call the [DescribeInstances](~~DescribeInstances~~) operation to obtain the ID of the data asset.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The search keyword. Fuzzy match is supported. For example, if you specify test, all tables whose names contain test are retrieved.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the package to which the table belongs. You can call the [DescribePackages](~~DescribePackages~~) operation to obtain the ID of the package.
	//
	// example:
	//
	// 555555
	PackageId *int64 `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service to which the table belongs, such as MaxCompute, OSS, and ApsaraDB RDS. For more information about the types of data assets from which Data Security Center (DSC) can scan for sensitive data, see [Supported data assets](https://help.aliyun.com/document_detail/212906.html).
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the table belongs. You can call the [DescribeDataAssets](~~DescribeDataAssets~~) operation to obtain the ID of the service.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the table. Each sensitivity level ID corresponds to a sensitivity level name. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The ID of the sensitive data detection rule that the table hits. You can call the [DescribeRules](~~DescribeRules~~) operation to obtain the ID of the sensitive data detection rule.
	//
	// example:
	//
	// 333322
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The region in which DSC is activated. For more information, see [Supported regions](https://help.aliyun.com/document_detail/214257.html).
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the industry-specific rule template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) SetCurrentPage(v int32) *DescribeTablesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTablesRequest) SetInstanceId(v int64) *DescribeTablesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTablesRequest) SetLang(v string) *DescribeTablesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTablesRequest) SetName(v string) *DescribeTablesRequest {
	s.Name = &v
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

func (s *DescribeTablesRequest) SetProductCode(v string) *DescribeTablesRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeTablesRequest) SetProductId(v int64) *DescribeTablesRequest {
	s.ProductId = &v
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

func (s *DescribeTablesRequest) SetServiceRegionId(v string) *DescribeTablesRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeTablesRequest) SetTemplateId(v int64) *DescribeTablesRequest {
	s.TemplateId = &v
	return s
}

type DescribeTablesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of tables.
	Items []*DescribeTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeTablesResponseBody) SetItems(v []*DescribeTablesResponseBodyItems) *DescribeTablesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTablesResponseBody) SetPageSize(v int32) *DescribeTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablesResponseBody) SetTotalCount(v int32) *DescribeTablesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTablesResponseBodyItems struct {
	// The point in time when the table was created. Unit: milliseconds.
	//
	// example:
	//
	// 1536751124000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the table.
	//
	// example:
	//
	// 222
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the data asset.
	//
	// example:
	//
	// Description 1
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The ID of the data asset to which the table belongs.
	//
	// example:
	//
	// 1
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the data asset to which the table belongs.
	//
	// example:
	//
	// Data Asset 1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// gxdata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account to which the table belongs.
	//
	// example:
	//
	// dtdep-239-******
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the service to which the table belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**. For more information about the types of data assets from which DSC can scan for sensitive data, see [Supported data assets](https://help.aliyun.com/document_detail/212906.html).
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the table belongs.
	//
	// example:
	//
	// 1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the table. Each sensitivity level ID corresponds to a sensitivity level name. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level for the table. Valid values:
	//
	// 	- **N/A**: indicates that no sensitive data is detected.
	//
	// 	- **S1**: indicates the low sensitivity level.
	//
	// 	- **S2**: indicates the medium sensitivity level.
	//
	// 	- **S3**: indicates the high sensitivity level.
	//
	// 	- **S4**: indicates the highest sensitivity level.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The information about the sensitive data detection rules that are hit.
	RuleList []*DescribeTablesResponseBodyItemsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// Indicates whether the table contains sensitive fields. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
	// The total number of sensitive fields in the table.
	//
	// example:
	//
	// 32
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The percentage of sensitive fields in the table.
	//
	// example:
	//
	// 21%
	SensitiveRatio *string `json:"SensitiveRatio,omitempty" xml:"SensitiveRatio,omitempty"`
	// The name of the tenant.
	//
	// example:
	//
	// Tenant 1
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The total number of fields in the table.
	//
	// example:
	//
	// 1234
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeTablesResponseBodyItems) SetId(v int64) *DescribeTablesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetInstanceDescription(v string) *DescribeTablesResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetInstanceId(v int64) *DescribeTablesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetInstanceName(v string) *DescribeTablesResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetName(v string) *DescribeTablesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetOwner(v string) *DescribeTablesResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetProductCode(v string) *DescribeTablesResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetProductId(v string) *DescribeTablesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetRiskLevelId(v int64) *DescribeTablesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetRiskLevelName(v string) *DescribeTablesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetRuleList(v []*DescribeTablesResponseBodyItemsRuleList) *DescribeTablesResponseBodyItems {
	s.RuleList = v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetSensitive(v bool) *DescribeTablesResponseBodyItems {
	s.Sensitive = &v
	return s
}

func (s *DescribeTablesResponseBodyItems) SetSensitiveCount(v int32) *DescribeTablesResponseBodyItems {
	s.SensitiveCount = &v
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

func (s *DescribeTablesResponseBodyItems) SetTotalCount(v int32) *DescribeTablesResponseBodyItems {
	s.TotalCount = &v
	return s
}

type DescribeTablesResponseBodyItemsRuleList struct {
	// The total number of rules.
	//
	// example:
	//
	// 12
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// Rule name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 1
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
}

func (s DescribeTablesResponseBodyItemsRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTablesResponseBodyItemsRuleList) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItemsRuleList) SetCount(v int64) *DescribeTablesResponseBodyItemsRuleList {
	s.Count = &v
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

type DescribeTablesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeTablesResponse) SetStatusCode(v int32) *DescribeTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTablesResponse) SetBody(v *DescribeTablesResponseBody) *DescribeTablesResponse {
	s.Body = v
	return s
}

type DescribeTemplateAllRulesRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language type for the request and response, default is **zh_cn**. Values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Industry template ID.
	//
	// > You can obtain the industry template ID by calling [DescribeCategoryTemplateList](https://help.aliyun.com/document_detail/2399296.html). If this parameter is not provided, the model list of the primary template will be returned by default.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeTemplateAllRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateAllRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAllRulesRequest) SetFeatureType(v int32) *DescribeTemplateAllRulesRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeTemplateAllRulesRequest) SetLang(v string) *DescribeTemplateAllRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTemplateAllRulesRequest) SetTemplateId(v int64) *DescribeTemplateAllRulesRequest {
	s.TemplateId = &v
	return s
}

type DescribeTemplateAllRulesResponseBody struct {
	// The unique identifier generated by Alibaba Cloud for this request.
	//
	// example:
	//
	// 769FB3C1-F4C9-4******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of model objects.
	RuleList []*DescribeTemplateAllRulesResponseBodyRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeTemplateAllRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateAllRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAllRulesResponseBody) SetRequestId(v string) *DescribeTemplateAllRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateAllRulesResponseBody) SetRuleList(v []*DescribeTemplateAllRulesResponseBodyRuleList) *DescribeTemplateAllRulesResponseBody {
	s.RuleList = v
	return s
}

type DescribeTemplateAllRulesResponseBodyRuleList struct {
	// Unique ID of the model.
	//
	// example:
	//
	// 376
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Model name.
	//
	// example:
	//
	// Model Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeTemplateAllRulesResponseBodyRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateAllRulesResponseBodyRuleList) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAllRulesResponseBodyRuleList) SetId(v int64) *DescribeTemplateAllRulesResponseBodyRuleList {
	s.Id = &v
	return s
}

func (s *DescribeTemplateAllRulesResponseBodyRuleList) SetName(v string) *DescribeTemplateAllRulesResponseBodyRuleList {
	s.Name = &v
	return s
}

type DescribeTemplateAllRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplateAllRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplateAllRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateAllRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAllRulesResponse) SetHeaders(v map[string]*string) *DescribeTemplateAllRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateAllRulesResponse) SetStatusCode(v int32) *DescribeTemplateAllRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateAllRulesResponse) SetBody(v *DescribeTemplateAllRulesResponseBody) *DescribeTemplateAllRulesResponse {
	s.Body = v
	return s
}

type DescribeUserStatusRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese (default)
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusRequest) SetFeatureType(v int32) *DescribeUserStatusRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeUserStatusRequest) SetLang(v string) *DescribeUserStatusRequest {
	s.Lang = &v
	return s
}

type DescribeUserStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the current account.
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
	// The AccessKey ID of the current account.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// Indicates whether the SQL Explorer feature can be disabled. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	AuditClosable *bool `json:"AuditClosable,omitempty" xml:"AuditClosable,omitempty"`
	// Indicates whether the audit resources can be released.
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	AuditReleasable *bool `json:"AuditReleasable,omitempty" xml:"AuditReleasable,omitempty"`
	// Indicates whether DSC has permission to access user resources within the current account. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Authed *bool `json:"Authed,omitempty" xml:"Authed,omitempty"`
	// The billing method of DCS that is purchased by using the current account. Valid values:
	//
	// 	- **PREPAY**: subscription
	//
	// 	- **POSTPAY**: pay-as-you-go
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The permissions that the current account has. Valid values:
	//
	// 	- **0**: The current account has the administrative permissions or read-only permissions on Data Security Center.
	//
	// 	- **1**: The current account has the permissions to manage data domains.
	//
	// example:
	//
	// 1
	DataManagerRole *int32 `json:"DataManagerRole,omitempty" xml:"DataManagerRole,omitempty"`
	// The ID of the data security center instance purchased by the main account.
	//
	// example:
	//
	// sddp-cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of instances within the current account.
	//
	// example:
	//
	// 32
	InstanceNum *int32 `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 10
	InstanceTotalCount *int64 `json:"InstanceTotalCount,omitempty" xml:"InstanceTotalCount,omitempty"`
	// Indicates whether the data security lab feature is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	LabStatus *int32 `json:"LabStatus,omitempty" xml:"LabStatus,omitempty"`
	// OSS total storage capacity. Unit: Bytes.
	//
	// example:
	//
	// 2048
	OssTotalSize *int64 `json:"OssTotalSize,omitempty" xml:"OssTotalSize,omitempty"`
	// Accumulate the number of days to protect user assets.
	//
	// example:
	//
	// 2
	ProtectionDays *int32 `json:"ProtectionDays,omitempty" xml:"ProtectionDays,omitempty"`
	// Indicates whether DSC is purchased. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Purchased *bool `json:"Purchased,omitempty" xml:"Purchased,omitempty"`
	// The grace period between when DSC is expired and when DSC is released. Unit: days.
	//
	// example:
	//
	// 15
	ReleaseDays *int32 `json:"ReleaseDays,omitempty" xml:"ReleaseDays,omitempty"`
	// The time when the audit resources are released. Unit: milliseconds.
	//
	// example:
	//
	// 15000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The remaining period for which the data assets within the current account can be protected by DSC.
	//
	// example:
	//
	// 131
	RemainDays *int32 `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	// Indicates whether the current account uses a free trial of DSC. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Trail *bool `json:"Trail,omitempty" xml:"Trail,omitempty"`
	// Indicates whether the agent audit feature is used. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	UseAgentAudit *bool `json:"UseAgentAudit,omitempty" xml:"UseAgentAudit,omitempty"`
	// The number of instances that are used.
	//
	// example:
	//
	// 125
	UseInstanceNum *int32 `json:"UseInstanceNum,omitempty" xml:"UseInstanceNum,omitempty"`
	// The occupied space of the Object Storage Service (OSS) bucket. Unit: bytes.
	//
	// example:
	//
	// 234
	UseOssSize *int64 `json:"UseOssSize,omitempty" xml:"UseOssSize,omitempty"`
}

func (s DescribeUserStatusResponseBodyUserStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusResponseBodyUserStatus) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAccessKeyId(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAuditClosable(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.AuditClosable = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAuditReleasable(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.AuditReleasable = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAuthed(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.Authed = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetChargeType(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.ChargeType = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetDataManagerRole(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.DataManagerRole = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetInstanceId(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetInstanceNum(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.InstanceNum = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetInstanceTotalCount(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.InstanceTotalCount = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetLabStatus(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.LabStatus = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetOssTotalSize(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.OssTotalSize = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetProtectionDays(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.ProtectionDays = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetPurchased(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.Purchased = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetReleaseDays(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.ReleaseDays = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetReleaseTime(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetRemainDays(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.RemainDays = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetTrail(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.Trail = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetUseAgentAudit(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.UseAgentAudit = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetUseInstanceNum(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.UseInstanceNum = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetUseOssSize(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.UseOssSize = &v
	return s
}

type DescribeUserStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeUserStatusResponse) SetStatusCode(v int32) *DescribeUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserStatusResponse) SetBody(v *DescribeUserStatusResponseBody) *DescribeUserStatusResponse {
	s.Body = v
	return s
}

type DisableUserConfigRequest struct {
	// The code of the configuration item. You can call the [DescribeConfigs](~~DescribeConfigs~~) operation to obtain the code of the configuration item.
	//
	// example:
	//
	// access_failed_cnt
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh_cn**: Chinese (default)
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DisableUserConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableUserConfigRequest) GoString() string {
	return s.String()
}

func (s *DisableUserConfigRequest) SetCode(v string) *DisableUserConfigRequest {
	s.Code = &v
	return s
}

func (s *DisableUserConfigRequest) SetFeatureType(v int32) *DisableUserConfigRequest {
	s.FeatureType = &v
	return s
}

func (s *DisableUserConfigRequest) SetLang(v string) *DisableUserConfigRequest {
	s.Lang = &v
	return s
}

type DisableUserConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AC314611-D907-5EBF-B6D8-70425E5A8643
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DisableUserConfigResponse) SetStatusCode(v int32) *DisableUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableUserConfigResponse) SetBody(v *DisableUserConfigResponseBody) *DisableUserConfigResponse {
	s.Body = v
	return s
}

type ExecDatamaskRequest struct {
	// The sensitive data to be de-identified. The value is a JSON string that contains the following parameters:
	//
	// 	- **dataHeaderList**: the names of the columns in which data needs to be de-identified. Specify the column names in accordance with the order of data that needs to be de-identified.
	//
	// 	- **dataList**: the data that needs to be de-identified.
	//
	// 	- **ruleList**: the IDs of sensitive data detection rules used to detect data that needs to be de-identified. Specify the rule IDs in accordance with the order of data that needs to be de-identified. Each ID identifies a sensitive data detection rule that is used to detect a type of sensitive data. You can call the [DescribeRules](~~DescribeRules~~) operation to query the IDs of sensitive data detection rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"dataHeaderList":["name","age"],"dataList":[["lily",18],["lucy",17]],"ruleList":[1002,null]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the de-identification template. The ID is generated after you create the de-identification template in the [Data Security Center (DSC) console](https://yundun.console.aliyun.com/?\\&p=sddpnext#/sddp/dm/template). You can choose **Data desensitization*	- > **Desensitization Template*	- in the left-side navigation pane and obtain the ID of the de-identification template from the **Desensitization Template*	- page.
	//
	// 	- If you select **Field name*	- as the matching mode of the template, DSC matches data based on the columns specified by the **dataHeaderList*	- parameter in the **Data*	- parameter.
	//
	// 	- If you select **Sensitive type*	- as the matching mode of the template, DSC matches data based on the sensitive data detection rules specified by the **ruleList*	- parameter in the **Data*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ExecDatamaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecDatamaskRequest) GoString() string {
	return s.String()
}

func (s *ExecDatamaskRequest) SetData(v string) *ExecDatamaskRequest {
	s.Data = &v
	return s
}

func (s *ExecDatamaskRequest) SetFeatureType(v int32) *ExecDatamaskRequest {
	s.FeatureType = &v
	return s
}

func (s *ExecDatamaskRequest) SetLang(v string) *ExecDatamaskRequest {
	s.Lang = &v
	return s
}

func (s *ExecDatamaskRequest) SetTemplateId(v int64) *ExecDatamaskRequest {
	s.TemplateId = &v
	return s
}

type ExecDatamaskResponseBody struct {
	// The de-identified data, which is described in a JSON string. The JSON string contains the following parameters:
	//
	// 	- **dataHeaderList**: the names of columns that contain the de-identified data.
	//
	// 	- **dataList**: the de-identified data. The column order of the de-identified data is the same as that indicated by the dataHeaderList parameter.
	//
	// 	- **ruleList**: the IDs of sensitive data detection rules.
	//
	// example:
	//
	// {"dataHeaderList":["name","age"],"dataList":[["l***",18],["l***",17]],"ruleList":[1002,null]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 813BA9FA-D062-42C4-8CD5-11A7640B96E6
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecDatamaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ExecDatamaskResponse) SetStatusCode(v int32) *ExecDatamaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecDatamaskResponse) SetBody(v *ExecDatamaskResponseBody) *ExecDatamaskResponse {
	s.Body = v
	return s
}

type ManualTriggerMaskingProcessRequest struct {
	// The ID of the de-identification task.
	//
	// The ID of the de-identification task is a string. You can call the DescribeDataMaskingTasks operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response, default value zh_cn. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ManualTriggerMaskingProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s ManualTriggerMaskingProcessRequest) GoString() string {
	return s.String()
}

func (s *ManualTriggerMaskingProcessRequest) SetId(v int64) *ManualTriggerMaskingProcessRequest {
	s.Id = &v
	return s
}

func (s *ManualTriggerMaskingProcessRequest) SetLang(v string) *ManualTriggerMaskingProcessRequest {
	s.Lang = &v
	return s
}

type ManualTriggerMaskingProcessResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-4******
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManualTriggerMaskingProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ManualTriggerMaskingProcessResponse) SetStatusCode(v int32) *ManualTriggerMaskingProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *ManualTriggerMaskingProcessResponse) SetBody(v *ManualTriggerMaskingProcessResponseBody) *ManualTriggerMaskingProcessResponse {
	s.Body = v
	return s
}

type ModifyDataLimitRequest struct {
	// Specifies whether to enable the security audit feature. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Specifies whether to automatically trigger a re-scan after a rule is modified. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// > When a re-scan is triggered, DSC scans all data in your data asset.
	//
	// example:
	//
	// 1
	AutoScan *int32 `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	// The database engine that is run by the instance. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The unique ID of the data asset for which you want to modify configuration items.
	//
	// > You can call the [DescribeDataLimits](~~DescribeDataLimits~~) operation to query the ID of the data asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The retention period of raw logs after you enable the security audit feature. Unit: days. Valid values:
	//
	// 	- **30**
	//
	// 	- **90**
	//
	// 	- **180**
	//
	// 	- **365**
	//
	// example:
	//
	// 30
	LogStoreDay *int32 `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	// Specifies whether to change the username and password that are used to log on to the ApsaraDB RDS database. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	ModifyPassword *bool `json:"ModifyPassword,omitempty" xml:"ModifyPassword,omitempty"`
	// The password used to log on to the ApsaraDB RDS database that you authorize DSC to access.
	//
	// example:
	//
	// ********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port that is used to connect to the database.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the service to which the data asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: Object Storage Service (OSS)
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of sensitive data samples tht are collected after sensitive data detection is enabled. Valid values:
	//
	// 	- **0**
	//
	// 	- **5**
	//
	// 	- **10**
	//
	// example:
	//
	// 0
	SamplingSize *int32 `json:"SamplingSize,omitempty" xml:"SamplingSize,omitempty"`
	// The security group that is used by PrivateLink when you install the DSC agent.
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" xml:"SecurityGroupIdList,omitempty" type:"Repeated"`
	// The region in which the data asset resides. Valid values:
	//
	// 	- **cn-beijing**: China (Beijing)
	//
	// 	- **cn-zhangjiakou**: China (Zhangjiakou)
	//
	// 	- **cn-huhehaote**: China (Hohhot)
	//
	// 	- **cn-hangzhou**: China (Hangzhou)
	//
	// 	- **cn-shanghai**: China (Shanghai)
	//
	// 	- **cn-shenzhen**: China (Shenzhen)
	//
	// 	- **cn-hongkong**: China (Hong Kong)
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The username used to log on to the ApsaraDB RDS database that you authorize DSC to access.
	//
	// example:
	//
	// User01
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The vSwitch that is used by PrivateLink when you install the DSC agent.
	VSwitchIdList []*string `json:"VSwitchIdList,omitempty" xml:"VSwitchIdList,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the data asset belongs.
	//
	// example:
	//
	// vpc-2zevcqke6hh09c41****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyDataLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDataLimitRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataLimitRequest) SetAuditStatus(v int32) *ModifyDataLimitRequest {
	s.AuditStatus = &v
	return s
}

func (s *ModifyDataLimitRequest) SetAutoScan(v int32) *ModifyDataLimitRequest {
	s.AutoScan = &v
	return s
}

func (s *ModifyDataLimitRequest) SetEngineType(v string) *ModifyDataLimitRequest {
	s.EngineType = &v
	return s
}

func (s *ModifyDataLimitRequest) SetFeatureType(v int32) *ModifyDataLimitRequest {
	s.FeatureType = &v
	return s
}

func (s *ModifyDataLimitRequest) SetId(v int64) *ModifyDataLimitRequest {
	s.Id = &v
	return s
}

func (s *ModifyDataLimitRequest) SetLang(v string) *ModifyDataLimitRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDataLimitRequest) SetLogStoreDay(v int32) *ModifyDataLimitRequest {
	s.LogStoreDay = &v
	return s
}

func (s *ModifyDataLimitRequest) SetModifyPassword(v bool) *ModifyDataLimitRequest {
	s.ModifyPassword = &v
	return s
}

func (s *ModifyDataLimitRequest) SetPassword(v string) *ModifyDataLimitRequest {
	s.Password = &v
	return s
}

func (s *ModifyDataLimitRequest) SetPort(v int32) *ModifyDataLimitRequest {
	s.Port = &v
	return s
}

func (s *ModifyDataLimitRequest) SetResourceType(v int32) *ModifyDataLimitRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyDataLimitRequest) SetSamplingSize(v int32) *ModifyDataLimitRequest {
	s.SamplingSize = &v
	return s
}

func (s *ModifyDataLimitRequest) SetSecurityGroupIdList(v []*string) *ModifyDataLimitRequest {
	s.SecurityGroupIdList = v
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

func (s *ModifyDataLimitRequest) SetVSwitchIdList(v []*string) *ModifyDataLimitRequest {
	s.VSwitchIdList = v
	return s
}

func (s *ModifyDataLimitRequest) SetVpcId(v string) *ModifyDataLimitRequest {
	s.VpcId = &v
	return s
}

type ModifyDataLimitResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyDataLimitResponse) SetStatusCode(v int32) *ModifyDataLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataLimitResponse) SetBody(v *ModifyDataLimitResponseBody) *ModifyDataLimitResponse {
	s.Body = v
	return s
}

type ModifyDefaultLevelRequest struct {
	// The default sensitivity level of data that Data Security Center (DSC) cannot classify as sensitive or insensitive. Valid values:
	//
	// 	- **1**: N/A
	//
	// 	- **2**: S1
	//
	// 	- **3**: S2
	//
	// 	- **4**: S3
	//
	// 	- **5**: S4
	//
	// example:
	//
	// 4
	DefaultId *int64 `json:"DefaultId,omitempty" xml:"DefaultId,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The sensitivity level ID of data that DSC classifies as sensitive. Separate multiple IDs with commas (,). Valid values:
	//
	// 	- **1**: N/A
	//
	// 	- **2**: S1
	//
	// 	- **3**: S2
	//
	// 	- **4**: S3
	//
	// 	- **5**: S4
	//
	// example:
	//
	// 1,2,3,4
	SensitiveIds *string `json:"SensitiveIds,omitempty" xml:"SensitiveIds,omitempty"`
}

func (s ModifyDefaultLevelRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefaultLevelRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefaultLevelRequest) SetDefaultId(v int64) *ModifyDefaultLevelRequest {
	s.DefaultId = &v
	return s
}

func (s *ModifyDefaultLevelRequest) SetLang(v string) *ModifyDefaultLevelRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDefaultLevelRequest) SetSensitiveIds(v string) *ModifyDefaultLevelRequest {
	s.SensitiveIds = &v
	return s
}

type ModifyDefaultLevelResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1EBF271
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefaultLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyDefaultLevelResponse) SetStatusCode(v int32) *ModifyDefaultLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefaultLevelResponse) SetBody(v *ModifyDefaultLevelResponseBody) *ModifyDefaultLevelResponse {
	s.Body = v
	return s
}

type ModifyEventStatusRequest struct {
	// Specifies whether to enhance the detection of anomalous events. If you enhance the detection of anomalous events, the detection accuracy and the rate of triggering alerts for anomalous events are improved. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Backed *bool `json:"Backed,omitempty" xml:"Backed,omitempty"`
	// The reason why the anomalous event is handled.
	//
	// example:
	//
	// Anomaly confirmed
	DealReason *string `json:"DealReason,omitempty" xml:"DealReason,omitempty"`
	// The ID of the anomalous event.
	//
	// > You can call the **DescribeEvents*	- operation to query the ID of the anomalous event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The method to handle the anomalous event. Valid values:
	//
	// 	- **1**: marks the anomalous event as a false positive.
	//
	// 	- **2**: confirms and handles the anomalous event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyEventStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEventStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyEventStatusRequest) SetBacked(v bool) *ModifyEventStatusRequest {
	s.Backed = &v
	return s
}

func (s *ModifyEventStatusRequest) SetDealReason(v string) *ModifyEventStatusRequest {
	s.DealReason = &v
	return s
}

func (s *ModifyEventStatusRequest) SetId(v int64) *ModifyEventStatusRequest {
	s.Id = &v
	return s
}

func (s *ModifyEventStatusRequest) SetLang(v string) *ModifyEventStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyEventStatusRequest) SetStatus(v int32) *ModifyEventStatusRequest {
	s.Status = &v
	return s
}

type ModifyEventStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8491DBFD-48C0-4E11-B6FC-6F38921244A9
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEventStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyEventStatusResponse) SetStatusCode(v int32) *ModifyEventStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEventStatusResponse) SetBody(v *ModifyEventStatusResponseBody) *ModifyEventStatusResponse {
	s.Body = v
	return s
}

type ModifyEventTypeStatusRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Valid values: **zh*	- and **en**. The value zh indicates Chinese, and the value en indicates English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the anomalous event subtype. Separate multiple IDs with commas (,).
	//
	// > You can call the **DescribeEventTypes*	- operation to query the ID of anomalous event subtype.
	//
	// example:
	//
	// 020008
	SubTypeIds *string `json:"SubTypeIds,omitempty" xml:"SubTypeIds,omitempty"`
}

func (s ModifyEventTypeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEventTypeStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyEventTypeStatusRequest) SetFeatureType(v int32) *ModifyEventTypeStatusRequest {
	s.FeatureType = &v
	return s
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
	// The ID of the request.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1E*****
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEventTypeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyEventTypeStatusResponse) SetStatusCode(v int32) *ModifyEventTypeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEventTypeStatusResponse) SetBody(v *ModifyEventTypeStatusResponseBody) *ModifyEventTypeStatusResponse {
	s.Body = v
	return s
}

type ModifyReportTaskStatusRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies the status of the report task. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// > This parameter is required.
	//
	// example:
	//
	// 0
	ReportTaskStatus *int32 `json:"ReportTaskStatus,omitempty" xml:"ReportTaskStatus,omitempty"`
}

func (s ModifyReportTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyReportTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyReportTaskStatusRequest) SetFeatureType(v int32) *ModifyReportTaskStatusRequest {
	s.FeatureType = &v
	return s
}

func (s *ModifyReportTaskStatusRequest) SetLang(v string) *ModifyReportTaskStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyReportTaskStatusRequest) SetReportTaskStatus(v int32) *ModifyReportTaskStatusRequest {
	s.ReportTaskStatus = &v
	return s
}

type ModifyReportTaskStatusResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1EBF271
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyReportTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyReportTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReportTaskStatusResponseBody) SetRequestId(v string) *ModifyReportTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyReportTaskStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyReportTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyReportTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyReportTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyReportTaskStatusResponse) SetHeaders(v map[string]*string) *ModifyReportTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyReportTaskStatusResponse) SetStatusCode(v int32) *ModifyReportTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyReportTaskStatusResponse) SetBody(v *ModifyReportTaskStatusResponseBody) *ModifyReportTaskStatusResponse {
	s.Body = v
	return s
}

type ModifyRuleRequest struct {
	// The content type of the sensitive data detection rule. Valid values:
	//
	// 	- **2**: regular expression
	//
	// 	- **3**: algorithm
	//
	// 	- **5**: keyword
	//
	// example:
	//
	// 2
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The content of the sensitive data detection rule. You can specify a regular expression, an algorithm, or keywords that are used to match sensitive fields or text.
	//
	// This parameter is required.
	//
	// example:
	//
	// (?:\\\\D|^)((?:(?:25[0-4]|2[0-4]\\\\d|1\\\\d{2}|[1-9]\\\\d{1})\\\\.)(?:(?:25[0-5]|2[0-4]\\\\d|[01]?\\\\d?\\\\d)\\\\.){2}(?:25[0-5]|2[0-4]\\\\d|1[0-9]\\\\d|[1-9]\\\\d|[1-9]))(?:\\\\D|$)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the sensitive data detection rule.
	//
	// You can call the [DescribeRules](~~DescribeRules~~) operation to obtain the rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The match type. Valid values:
	//
	// 	- **1**: rule-based match
	//
	// 	- **2**: dictionary-based match
	//
	// example:
	//
	// 1
	MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The IDs of the models for sensitive data audit.
	//
	// example:
	//
	// 1452
	ModelRuleIds *string `json:"ModelRuleIds,omitempty" xml:"ModelRuleIds,omitempty"`
	// The name of the sensitive data detection rule.
	//
	// You can call the [DescribeRules](~~DescribeRules~~) operation to obtain the rule name.
	//
	// This parameter is required.
	//
	// example:
	//
	// esw
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The service to which the sensitive data detection rule is applied. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the sensitive data detection rule is applied. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The type of the sensitive data detection rule. Valid values:
	//
	// 	- **1**: data detection rule
	//
	// 	- **2**: audit rule
	//
	// 	- **3**: anomalous event detection rule
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The data assets supported by the sensitive data detection rule. Valid values:
	//
	// 	- **0**: all data assets
	//
	// 	- **1**: structured data assets
	//
	// 	- **2**: unstructured data assets
	//
	// example:
	//
	// 1
	SupportForm *int32 `json:"SupportForm,omitempty" xml:"SupportForm,omitempty"`
	// The IDs of the templates that are used to audit sensitive data.
	//
	// example:
	//
	// 1
	TemplateRuleIds *string `json:"TemplateRuleIds,omitempty" xml:"TemplateRuleIds,omitempty"`
	// The risk level of the alert that is triggered by the sensitive data detection rule. Valid values:
	//
	// 	- **1**: low level
	//
	// 	- **2**: medium level
	//
	// 	- **3**: high level
	//
	// example:
	//
	// 1
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s ModifyRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyRuleRequest) SetCategory(v int32) *ModifyRuleRequest {
	s.Category = &v
	return s
}

func (s *ModifyRuleRequest) SetContent(v string) *ModifyRuleRequest {
	s.Content = &v
	return s
}

func (s *ModifyRuleRequest) SetId(v int64) *ModifyRuleRequest {
	s.Id = &v
	return s
}

func (s *ModifyRuleRequest) SetLang(v string) *ModifyRuleRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRuleRequest) SetMatchType(v int32) *ModifyRuleRequest {
	s.MatchType = &v
	return s
}

func (s *ModifyRuleRequest) SetModelRuleIds(v string) *ModifyRuleRequest {
	s.ModelRuleIds = &v
	return s
}

func (s *ModifyRuleRequest) SetName(v string) *ModifyRuleRequest {
	s.Name = &v
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

func (s *ModifyRuleRequest) SetRiskLevelId(v int64) *ModifyRuleRequest {
	s.RiskLevelId = &v
	return s
}

func (s *ModifyRuleRequest) SetRuleType(v int32) *ModifyRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyRuleRequest) SetSupportForm(v int32) *ModifyRuleRequest {
	s.SupportForm = &v
	return s
}

func (s *ModifyRuleRequest) SetTemplateRuleIds(v string) *ModifyRuleRequest {
	s.TemplateRuleIds = &v
	return s
}

func (s *ModifyRuleRequest) SetWarnLevel(v int32) *ModifyRuleRequest {
	s.WarnLevel = &v
	return s
}

type ModifyRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyRuleResponse) SetStatusCode(v int32) *ModifyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRuleResponse) SetBody(v *ModifyRuleResponseBody) *ModifyRuleResponse {
	s.Body = v
	return s
}

type ModifyRuleStatusRequest struct {
	// The ID of the sensitive data detection rule.
	//
	// > You can query the ID of the sensitive data detection rule by calling the **DescribeRules*	- operation.
	//
	// example:
	//
	// 12341
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the sensitive data detection rule. Separate multiple IDs with commas (,).
	//
	// > You can query the ID of the sensitive data detection rule by calling the **DescribeRules*	- operation.
	//
	// example:
	//
	// 1,2,3,4
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to enable or disable the sensitive data detection rule. Valid values:
	//
	// 	- **0**: disables the sensitive data detection rule.
	//
	// 	- **1**: enables the sensitive data detection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyRuleStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyRuleStatusRequest) SetId(v int64) *ModifyRuleStatusRequest {
	s.Id = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetIds(v string) *ModifyRuleStatusRequest {
	s.Ids = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetLang(v string) *ModifyRuleStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRuleStatusRequest) SetStatus(v int32) *ModifyRuleStatusRequest {
	s.Status = &v
	return s
}

type ModifyRuleStatusResponseBody struct {
	// The IDs of sensitive data detection rules whose status failed to be changed. Multiple IDs are separated with commas (,).
	//
	// example:
	//
	// 1,2,3,4
	FailedIds *string `json:"FailedIds,omitempty" xml:"FailedIds,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyRuleStatusResponse) SetStatusCode(v int32) *ModifyRuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRuleStatusResponse) SetBody(v *ModifyRuleStatusResponseBody) *ModifyRuleStatusResponse {
	s.Body = v
	return s
}

type ScanOssObjectV1Request struct {
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// sddp-api-demo-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The objects in the OSS bucket that you want to scan. You can specify up to 50 objects at a time.
	//
	// This parameter is required.
	ObjectKeyList []*string `json:"ObjectKeyList,omitempty" xml:"ObjectKeyList,omitempty" type:"Repeated"`
	// The ID of the region in which the OSS bucket is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the industry-specific classification template.
	//
	// >  You can call the **DescribeCategoryTemplateList*	- operation to query industry-specific classification templates. If you do not specify this parameter, the system automatically uses the main template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ScanOssObjectV1Request) String() string {
	return tea.Prettify(s)
}

func (s ScanOssObjectV1Request) GoString() string {
	return s.String()
}

func (s *ScanOssObjectV1Request) SetBucketName(v string) *ScanOssObjectV1Request {
	s.BucketName = &v
	return s
}

func (s *ScanOssObjectV1Request) SetLang(v string) *ScanOssObjectV1Request {
	s.Lang = &v
	return s
}

func (s *ScanOssObjectV1Request) SetObjectKeyList(v []*string) *ScanOssObjectV1Request {
	s.ObjectKeyList = v
	return s
}

func (s *ScanOssObjectV1Request) SetServiceRegionId(v string) *ScanOssObjectV1Request {
	s.ServiceRegionId = &v
	return s
}

func (s *ScanOssObjectV1Request) SetTemplateId(v int64) *ScanOssObjectV1Request {
	s.TemplateId = &v
	return s
}

type ScanOssObjectV1ShrinkRequest struct {
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// sddp-api-demo-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The objects in the OSS bucket that you want to scan. You can specify up to 50 objects at a time.
	//
	// This parameter is required.
	ObjectKeyListShrink *string `json:"ObjectKeyList,omitempty" xml:"ObjectKeyList,omitempty"`
	// The ID of the region in which the OSS bucket is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the industry-specific classification template.
	//
	// >  You can call the **DescribeCategoryTemplateList*	- operation to query industry-specific classification templates. If you do not specify this parameter, the system automatically uses the main template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ScanOssObjectV1ShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanOssObjectV1ShrinkRequest) GoString() string {
	return s.String()
}

func (s *ScanOssObjectV1ShrinkRequest) SetBucketName(v string) *ScanOssObjectV1ShrinkRequest {
	s.BucketName = &v
	return s
}

func (s *ScanOssObjectV1ShrinkRequest) SetLang(v string) *ScanOssObjectV1ShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ScanOssObjectV1ShrinkRequest) SetObjectKeyListShrink(v string) *ScanOssObjectV1ShrinkRequest {
	s.ObjectKeyListShrink = &v
	return s
}

func (s *ScanOssObjectV1ShrinkRequest) SetServiceRegionId(v string) *ScanOssObjectV1ShrinkRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *ScanOssObjectV1ShrinkRequest) SetTemplateId(v int64) *ScanOssObjectV1ShrinkRequest {
	s.TemplateId = &v
	return s
}

type ScanOssObjectV1ResponseBody struct {
	// The ID of the identification task that is returned after the identification task is created.
	//
	// example:
	//
	// 268
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScanOssObjectV1ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScanOssObjectV1ResponseBody) GoString() string {
	return s.String()
}

func (s *ScanOssObjectV1ResponseBody) SetId(v int64) *ScanOssObjectV1ResponseBody {
	s.Id = &v
	return s
}

func (s *ScanOssObjectV1ResponseBody) SetRequestId(v string) *ScanOssObjectV1ResponseBody {
	s.RequestId = &v
	return s
}

type ScanOssObjectV1Response struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScanOssObjectV1ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScanOssObjectV1Response) String() string {
	return tea.Prettify(s)
}

func (s ScanOssObjectV1Response) GoString() string {
	return s.String()
}

func (s *ScanOssObjectV1Response) SetHeaders(v map[string]*string) *ScanOssObjectV1Response {
	s.Headers = v
	return s
}

func (s *ScanOssObjectV1Response) SetStatusCode(v int32) *ScanOssObjectV1Response {
	s.StatusCode = &v
	return s
}

func (s *ScanOssObjectV1Response) SetBody(v *ScanOssObjectV1ResponseBody) *ScanOssObjectV1Response {
	s.Body = v
	return s
}

type StopMaskingProcessRequest struct {
	// The unique ID of the de-identification task. You can query the task ID by calling the [DescribeDataMaskingTasks](~~DescribeDataMaskingTasks~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese (default)
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s StopMaskingProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s StopMaskingProcessRequest) GoString() string {
	return s.String()
}

func (s *StopMaskingProcessRequest) SetId(v int64) *StopMaskingProcessRequest {
	s.Id = &v
	return s
}

func (s *StopMaskingProcessRequest) SetLang(v string) *StopMaskingProcessRequest {
	s.Lang = &v
	return s
}

type StopMaskingProcessResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopMaskingProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StopMaskingProcessResponse) SetStatusCode(v int32) *StopMaskingProcessResponse {
	s.StatusCode = &v
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

// Summary:
//
// Modifies the configurations of a common configuration item for alerts.
//
// Description:
//
// You can call this operation to create or restore configurations based on the codes of common configuration items. This allows you to manage the configurations of common configuration items.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigResponse
func (client *Client) CreateConfigWithOptions(request *CreateConfigRequest, runtime *util.RuntimeOptions) (_result *CreateConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConfig"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a common configuration item for alerts.
//
// Description:
//
// You can call this operation to create or restore configurations based on the codes of common configuration items. This allows you to manage the configurations of common configuration items.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateConfigRequest
//
// @return CreateConfigResponse
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

// Summary:
//
// Authorizes Data Security Center (DSC) to scan data assets. The data assets can be a database, a project, or a bucket.
//
// Description:
//
// You can call this operation to authorize DSC to scan data assets to ensure the security of the data assets.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateDataLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLimitResponse
func (client *Client) CreateDataLimitWithOptions(request *CreateDataLimitRequest, runtime *util.RuntimeOptions) (_result *CreateDataLimitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditStatus)) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AutoScan)) {
		query["AutoScan"] = request.AutoScan
	}

	if !tea.BoolValue(util.IsUnset(request.CertificatePermission)) {
		query["CertificatePermission"] = request.CertificatePermission
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.EngineType)) {
		query["EngineType"] = request.EngineType
	}

	if !tea.BoolValue(util.IsUnset(request.EventStatus)) {
		query["EventStatus"] = request.EventStatus
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.InstantlyScan)) {
		query["InstantlyScan"] = request.InstantlyScan
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LogStoreDay)) {
		query["LogStoreDay"] = request.LogStoreDay
	}

	if !tea.BoolValue(util.IsUnset(request.OcrStatus)) {
		query["OcrStatus"] = request.OcrStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SamplingSize)) {
		query["SamplingSize"] = request.SamplingSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRegionId)) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataLimit"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes Data Security Center (DSC) to scan data assets. The data assets can be a database, a project, or a bucket.
//
// Description:
//
// You can call this operation to authorize DSC to scan data assets to ensure the security of the data assets.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateDataLimitRequest
//
// @return CreateDataLimitResponse
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

// Summary:
//
// Creates a custom sensitive data detection rule.
//
// @param request - CreateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRuleResponse
func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentCategory)) {
		query["ContentCategory"] = request.ContentCategory
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.ModelRuleIds)) {
		query["ModelRuleIds"] = request.ModelRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevelId)) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StatExpress)) {
		query["StatExpress"] = request.StatExpress
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SupportForm)) {
		query["SupportForm"] = request.SupportForm
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateRuleIds)) {
		query["TemplateRuleIds"] = request.TemplateRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.WarnLevel)) {
		query["WarnLevel"] = request.WarnLevel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRule"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom sensitive data detection rule.
//
// @param request - CreateRuleRequest
//
// @return CreateRuleResponse
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

// Summary:
//
// Creates a custom scan task. The custom scan task is used to scan data assets on which Data Security Center (DSC) is granted the scan permissions for sensitive data.
//
// Description:
//
// You can call this operation to create a custom scan task for authorized data assets. You can customize the interval between two consecutive scan tasks and the time when the scan task is executed next time.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScanTaskResponse
func (client *Client) CreateScanTaskWithOptions(request *CreateScanTaskRequest, runtime *util.RuntimeOptions) (_result *CreateScanTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataLimitId)) {
		query["DataLimitId"] = request.DataLimitId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalDay)) {
		query["IntervalDay"] = request.IntervalDay
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OssScanPath)) {
		query["OssScanPath"] = request.OssScanPath
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.RunHour)) {
		query["RunHour"] = request.RunHour
	}

	if !tea.BoolValue(util.IsUnset(request.RunMinute)) {
		query["RunMinute"] = request.RunMinute
	}

	if !tea.BoolValue(util.IsUnset(request.ScanRange)) {
		query["ScanRange"] = request.ScanRange
	}

	if !tea.BoolValue(util.IsUnset(request.ScanRangeContent)) {
		query["ScanRangeContent"] = request.ScanRangeContent
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUserName)) {
		query["TaskUserName"] = request.TaskUserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScanTask"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom scan task. The custom scan task is used to scan data assets on which Data Security Center (DSC) is granted the scan permissions for sensitive data.
//
// Description:
//
// You can call this operation to create a custom scan task for authorized data assets. You can customize the interval between two consecutive scan tasks and the time when the scan task is executed next time.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateScanTaskRequest
//
// @return CreateScanTaskResponse
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

// Summary:
//
// Creates a service-linked role for Data Security Center (DSC) to grant DSC the permissions to access data assets in other services.
//
// Description:
//
// You can call this operation to allow DSC to access the data assets in services such as Object Storage Service (OSS), ApsaraDB RDS, and MaxCompute. After you call this operation, the system automatically creates a service-linked role named AliyunServiceRoleForSDDP and attaches the AliyunServiceRolePolicyForSDDP policy to the role.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateSlrRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSlrRoleResponse
func (client *Client) CreateSlrRoleWithOptions(request *CreateSlrRoleRequest, runtime *util.RuntimeOptions) (_result *CreateSlrRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSlrRole"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSlrRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role for Data Security Center (DSC) to grant DSC the permissions to access data assets in other services.
//
// Description:
//
// You can call this operation to allow DSC to access the data assets in services such as Object Storage Service (OSS), ApsaraDB RDS, and MaxCompute. After you call this operation, the system automatically creates a service-linked role named AliyunServiceRoleForSDDP and attaches the AliyunServiceRolePolicyForSDDP policy to the role.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateSlrRoleRequest
//
// @return CreateSlrRoleResponse
func (client *Client) CreateSlrRole(request *CreateSlrRoleRequest) (_result *CreateSlrRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSlrRoleResponse{}
	_body, _err := client.CreateSlrRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the scan permissions on a data asset. The data asset can be a database, an instance, or a bucket.
//
// Description:
//
// You can call this operation to revoke the permissions on a data asset from Data Security Center (DSC).
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteDataLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLimitResponse
func (client *Client) DeleteDataLimitWithOptions(request *DeleteDataLimitRequest, runtime *util.RuntimeOptions) (_result *DeleteDataLimitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataLimit"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the scan permissions on a data asset. The data asset can be a database, an instance, or a bucket.
//
// Description:
//
// You can call this operation to revoke the permissions on a data asset from Data Security Center (DSC).
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteDataLimitRequest
//
// @return DeleteDataLimitResponse
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

// Summary:
//
// Deletes a custom sensitive data detection rule from Data Security Center (DSC).
//
// @param request - DeleteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleResponse
func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRule"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom sensitive data detection rule from Data Security Center (DSC).
//
// @param request - DeleteRuleRequest
//
// @return DeleteRuleResponse
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

// Summary:
//
// Call this interface to query the list of industry templates.
//
// @param request - DescribeCategoryTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCategoryTemplateListResponse
func (client *Client) DescribeCategoryTemplateListWithOptions(request *DescribeCategoryTemplateListRequest, runtime *util.RuntimeOptions) (_result *DescribeCategoryTemplateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UsageScenario)) {
		query["UsageScenario"] = request.UsageScenario
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCategoryTemplateList"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCategoryTemplateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call this interface to query the list of industry templates.
//
// @param request - DescribeCategoryTemplateListRequest
//
// @return DescribeCategoryTemplateListResponse
func (client *Client) DescribeCategoryTemplateList(request *DescribeCategoryTemplateListRequest) (_result *DescribeCategoryTemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCategoryTemplateListResponse{}
	_body, _err := client.DescribeCategoryTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries rules in a classification template by page.
//
// Description:
//
// You can call this operation to query rules in a classification template.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCategoryTemplateRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCategoryTemplateRuleListResponse
func (client *Client) DescribeCategoryTemplateRuleListWithOptions(request *DescribeCategoryTemplateRuleListRequest, runtime *util.RuntimeOptions) (_result *DescribeCategoryTemplateRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevelId)) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCategoryTemplateRuleList"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCategoryTemplateRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries rules in a classification template by page.
//
// Description:
//
// You can call this operation to query rules in a classification template.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeCategoryTemplateRuleListRequest
//
// @return DescribeCategoryTemplateRuleListResponse
func (client *Client) DescribeCategoryTemplateRuleList(request *DescribeCategoryTemplateRuleListRequest) (_result *DescribeCategoryTemplateRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCategoryTemplateRuleListResponse{}
	_body, _err := client.DescribeCategoryTemplateRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries data in the columns of the tables that Data Security Center (DSC) is authorized to access. The tables include the tables of MaxCompute and ApsaraDB RDS.
//
// Description:
//
// You can call this operation to query the data in columns of a table that may contain sensitive data. This helps you analyze sensitive data.
//
// ## [](#)Precautions
//
// The DescribeColumns operation is changed to DescribeColumnsV2. We recommend that you call the DescribeColumnsV2 operation when you develop your applications.
//
// ## [](#qps)Limits
//
// Each Alibaba Cloud account can call this operation up to 10 times per second. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnsResponse
func (client *Client) DescribeColumnsWithOptions(request *DescribeColumnsRequest, runtime *util.RuntimeOptions) (_result *DescribeColumnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EngineType)) {
		query["EngineType"] = request.EngineType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ModelTagId)) {
		query["ModelTagId"] = request.ModelTagId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevelId)) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.SensLevelName)) {
		query["SensLevelName"] = request.SensLevelName
	}

	if !tea.BoolValue(util.IsUnset(request.TableId)) {
		query["TableId"] = request.TableId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateRuleId)) {
		query["TemplateRuleId"] = request.TemplateRuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeColumns"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data in the columns of the tables that Data Security Center (DSC) is authorized to access. The tables include the tables of MaxCompute and ApsaraDB RDS.
//
// Description:
//
// You can call this operation to query the data in columns of a table that may contain sensitive data. This helps you analyze sensitive data.
//
// ## [](#)Precautions
//
// The DescribeColumns operation is changed to DescribeColumnsV2. We recommend that you call the DescribeColumnsV2 operation when you develop your applications.
//
// ## [](#qps)Limits
//
// Each Alibaba Cloud account can call this operation up to 10 times per second. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeColumnsRequest
//
// @return DescribeColumnsResponse
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

// Summary:
//
// Query data in columns of data assets such as MaxCompute, RDS, etc., that are authorized by the Data Security Center.
//
// @param request - DescribeColumnsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnsV2Response
func (client *Client) DescribeColumnsV2WithOptions(request *DescribeColumnsV2Request, runtime *util.RuntimeOptions) (_result *DescribeColumnsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EngineType)) {
		query["EngineType"] = request.EngineType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevelId)) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.SensLevelName)) {
		query["SensLevelName"] = request.SensLevelName
	}

	if !tea.BoolValue(util.IsUnset(request.TableId)) {
		query["TableId"] = request.TableId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeColumnsV2"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeColumnsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query data in columns of data assets such as MaxCompute, RDS, etc., that are authorized by the Data Security Center.
//
// @param request - DescribeColumnsV2Request
//
// @return DescribeColumnsV2Response
func (client *Client) DescribeColumnsV2(request *DescribeColumnsV2Request) (_result *DescribeColumnsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeColumnsV2Response{}
	_body, _err := client.DescribeColumnsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries common configuration items for alerts.
//
// @param request - DescribeConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConfigsResponse
func (client *Client) DescribeConfigsWithOptions(request *DescribeConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConfigs"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries common configuration items for alerts.
//
// @param request - DescribeConfigsRequest
//
// @return DescribeConfigsResponse
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

// Summary:
//
// Queries the sensitive data detection results of data assets that Data Security Center (DSC) is authorized to access.
//
// @param request - DescribeDataAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataAssetsResponse
func (client *Client) DescribeDataAssetsWithOptions(request *DescribeDataAssetsRequest, runtime *util.RuntimeOptions) (_result *DescribeDataAssetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RangeId)) {
		query["RangeId"] = request.RangeId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevels)) {
		query["RiskLevels"] = request.RiskLevels
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataAssets"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataAssetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the sensitive data detection results of data assets that Data Security Center (DSC) is authorized to access.
//
// @param request - DescribeDataAssetsRequest
//
// @return DescribeDataAssetsResponse
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

// Summary:
//
// Queries the details of a data asset, such as a MaxCompute project, an ApsaraDB RDS database, or an Object Storage Service (OSS) bucket, that you authorize Data Security Center (DSC) to access.
//
// @param request - DescribeDataLimitDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataLimitDetailResponse
func (client *Client) DescribeDataLimitDetailWithOptions(request *DescribeDataLimitDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDataLimitDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataLimitDetail"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataLimitDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data asset, such as a MaxCompute project, an ApsaraDB RDS database, or an Object Storage Service (OSS) bucket, that you authorize Data Security Center (DSC) to access.
//
// @param request - DescribeDataLimitDetailRequest
//
// @return DescribeDataLimitDetailResponse
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

// Summary:
//
// Queries data assets, such as instances, databases, and Object Storage Service (OSS) buckets, that you authorize Data Security Center (DSC) to scan in a service.
//
// Description:
//
// You can call this operation to query the data assets that are authorized to be scanned. This facilitates resource search and aggregation.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDataLimitSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataLimitSetResponse
func (client *Client) DescribeDataLimitSetWithOptions(request *DescribeDataLimitSetRequest, runtime *util.RuntimeOptions) (_result *DescribeDataLimitSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionType)) {
		query["RegionType"] = request.RegionType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataLimitSet"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataLimitSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data assets, such as instances, databases, and Object Storage Service (OSS) buckets, that you authorize Data Security Center (DSC) to scan in a service.
//
// Description:
//
// You can call this operation to query the data assets that are authorized to be scanned. This facilitates resource search and aggregation.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDataLimitSetRequest
//
// @return DescribeDataLimitSetResponse
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

// Summary:
//
// Queries the data assets such as instances, databases, or buckets that Data Security Center (DSC) is authorized to access.
//
// @param request - DescribeDataLimitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataLimitsResponse
func (client *Client) DescribeDataLimitsWithOptions(request *DescribeDataLimitsRequest, runtime *util.RuntimeOptions) (_result *DescribeDataLimitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditStatus)) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CheckStatus)) {
		query["CheckStatus"] = request.CheckStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DatamaskStatus)) {
		query["DatamaskStatus"] = request.DatamaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EngineType)) {
		query["EngineType"] = request.EngineType
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberAccount)) {
		query["MemberAccount"] = request.MemberAccount
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRegionId)) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataLimits"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataLimitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data assets such as instances, databases, or buckets that Data Security Center (DSC) is authorized to access.
//
// @param request - DescribeDataLimitsRequest
//
// @return DescribeDataLimitsResponse
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

// Summary:
//
// Queries the execution information about a de-identification task.
//
// Description:
//
// You can call this operation to query the execution information of a static de-identification task, including the status and progress.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDataMaskingRunHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataMaskingRunHistoryResponse
func (client *Client) DescribeDataMaskingRunHistoryWithOptions(request *DescribeDataMaskingRunHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeDataMaskingRunHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DstType)) {
		query["DstType"] = request.DstType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MainProcessId)) {
		query["MainProcessId"] = request.MainProcessId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SrcTableName)) {
		query["SrcTableName"] = request.SrcTableName
	}

	if !tea.BoolValue(util.IsUnset(request.SrcType)) {
		query["SrcType"] = request.SrcType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataMaskingRunHistory"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataMaskingRunHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution information about a de-identification task.
//
// Description:
//
// You can call this operation to query the execution information of a static de-identification task, including the status and progress.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDataMaskingRunHistoryRequest
//
// @return DescribeDataMaskingRunHistoryResponse
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

// Summary:
//
// Queries de-identification tasks.
//
// Description:
//
// You can call this operation to query static de-identification tasks. This facilitates task queries and management.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDataMaskingTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataMaskingTasksResponse
func (client *Client) DescribeDataMaskingTasksWithOptions(request *DescribeDataMaskingTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeDataMaskingTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DstType)) {
		query["DstType"] = request.DstType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataMaskingTasks"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataMaskingTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries de-identification tasks.
//
// Description:
//
// You can call this operation to query static de-identification tasks. This facilitates task queries and management.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDataMaskingTasksRequest
//
// @return DescribeDataMaskingTasksResponse
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

// Summary:
//
// # View data object column details
//
// Description:
//
// ## Notes
//
// The DescribeDataObjectColumnDetail interface has been revised to DescribeDataObjectColumnDetailV2. It is recommended that you use the newer version, DescribeDataObjectColumnDetailV2, when developing your application.
//
// @param request - DescribeDataObjectColumnDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataObjectColumnDetailResponse
func (client *Client) DescribeDataObjectColumnDetailWithOptions(request *DescribeDataObjectColumnDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDataObjectColumnDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataObjectColumnDetail"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataObjectColumnDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View data object column details
//
// Description:
//
// ## Notes
//
// The DescribeDataObjectColumnDetail interface has been revised to DescribeDataObjectColumnDetailV2. It is recommended that you use the newer version, DescribeDataObjectColumnDetailV2, when developing your application.
//
// @param request - DescribeDataObjectColumnDetailRequest
//
// @return DescribeDataObjectColumnDetailResponse
func (client *Client) DescribeDataObjectColumnDetail(request *DescribeDataObjectColumnDetailRequest) (_result *DescribeDataObjectColumnDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataObjectColumnDetailResponse{}
	_body, _err := client.DescribeDataObjectColumnDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # View Data Object Column Details V2
//
// @param request - DescribeDataObjectColumnDetailV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataObjectColumnDetailV2Response
func (client *Client) DescribeDataObjectColumnDetailV2WithOptions(request *DescribeDataObjectColumnDetailV2Request, runtime *util.RuntimeOptions) (_result *DescribeDataObjectColumnDetailV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataObjectColumnDetailV2"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataObjectColumnDetailV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View Data Object Column Details V2
//
// @param request - DescribeDataObjectColumnDetailV2Request
//
// @return DescribeDataObjectColumnDetailV2Response
func (client *Client) DescribeDataObjectColumnDetailV2(request *DescribeDataObjectColumnDetailV2Request) (_result *DescribeDataObjectColumnDetailV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataObjectColumnDetailV2Response{}
	_body, _err := client.DescribeDataObjectColumnDetailV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Paginated Query of Data Catalog Objects
//
// @param request - DescribeDataObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataObjectsResponse
func (client *Client) DescribeDataObjectsWithOptions(request *DescribeDataObjectsRequest, runtime *util.RuntimeOptions) (_result *DescribeDataObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.FileCategoryCode)) {
		query["FileCategoryCode"] = request.FileCategoryCode
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberAccount)) {
		query["MemberAccount"] = request.MemberAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ModelIds)) {
		query["ModelIds"] = request.ModelIds
	}

	if !tea.BoolValue(util.IsUnset(request.ModelTagIds)) {
		query["ModelTagIds"] = request.ModelTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCategoryIds)) {
		query["ParentCategoryIds"] = request.ParentCategoryIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProductIds)) {
		query["ProductIds"] = request.ProductIds
	}

	if !tea.BoolValue(util.IsUnset(request.QueryName)) {
		query["QueryName"] = request.QueryName
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevels)) {
		query["RiskLevels"] = request.RiskLevels
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRegionId)) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataObjects"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Paginated Query of Data Catalog Objects
//
// @param request - DescribeDataObjectsRequest
//
// @return DescribeDataObjectsResponse
func (client *Client) DescribeDataObjects(request *DescribeDataObjectsRequest) (_result *DescribeDataObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDataObjectsResponse{}
	_body, _err := client.DescribeDataObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of OSS object types that can be identified.
//
// @param request - DescribeDocTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocTypesResponse
func (client *Client) DescribeDocTypesWithOptions(request *DescribeDocTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeDocTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDocTypes"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDocTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of OSS object types that can be identified.
//
// @param request - DescribeDocTypesRequest
//
// @return DescribeDocTypesResponse
func (client *Client) DescribeDocTypes(request *DescribeDocTypesRequest) (_result *DescribeDocTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDocTypesResponse{}
	_body, _err := client.DescribeDocTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an anomalous event. The details include the time when the anomalous event occurred, and the description and handling status of the anomalous event.
//
// @param request - DescribeEventDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventDetailResponse
func (client *Client) DescribeEventDetailWithOptions(request *DescribeEventDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeEventDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEventDetail"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEventDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an anomalous event. The details include the time when the anomalous event occurred, and the description and handling status of the anomalous event.
//
// @param request - DescribeEventDetailRequest
//
// @return DescribeEventDetailResponse
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

// Summary:
//
// Queries the types of anomalous events.
//
// @param request - DescribeEventTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventTypesResponse
func (client *Client) DescribeEventTypesWithOptions(request *DescribeEventTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeEventTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ParentTypeId)) {
		query["ParentTypeId"] = request.ParentTypeId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEventTypes"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEventTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the types of anomalous events.
//
// @param request - DescribeEventTypesRequest
//
// @return DescribeEventTypesResponse
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

// Summary:
//
// Queries anomalous events.
//
// Description:
//
// You can call this operation to query anomalous events that may involve data leaks. This helps you search for and handle anomalous events.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsResponse
func (client *Client) DescribeEventsWithOptions(request *DescribeEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DealUserId)) {
		query["DealUserId"] = request.DealUserId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SubTypeCode)) {
		query["SubTypeCode"] = request.SubTypeCode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetProductCode)) {
		query["TargetProductCode"] = request.TargetProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.TypeCode)) {
		query["TypeCode"] = request.TypeCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.WarnLevel)) {
		query["WarnLevel"] = request.WarnLevel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEvents"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries anomalous events.
//
// Description:
//
// You can call this operation to query anomalous events that may involve data leaks. This helps you search for and handle anomalous events.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeEventsRequest
//
// @return DescribeEventsResponse
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

// Summary:
//
// # Query the status of an identification task
//
// Description:
//
// ## QPS Limit
//
// The QPS limit for this interface per user is 10 times/second. Exceeding the limit will result in API calls being rate-limited, which may affect your business. Please call it reasonably.
//
// @param request - DescribeIdentifyTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIdentifyTaskStatusResponse
func (client *Client) DescribeIdentifyTaskStatusWithOptions(request *DescribeIdentifyTaskStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeIdentifyTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIdentifyTaskStatus"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIdentifyTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the status of an identification task
//
// Description:
//
// ## QPS Limit
//
// The QPS limit for this interface per user is 10 times/second. Exceeding the limit will result in API calls being rate-limited, which may affect your business. Please call it reasonably.
//
// @param request - DescribeIdentifyTaskStatusRequest
//
// @return DescribeIdentifyTaskStatusResponse
func (client *Client) DescribeIdentifyTaskStatus(request *DescribeIdentifyTaskStatusRequest) (_result *DescribeIdentifyTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIdentifyTaskStatusResponse{}
	_body, _err := client.DescribeIdentifyTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data assets.
//
// Description:
//
// You can query a list of unauthorized or authorized data assets based on the value of AuthStatus.
//
// This operation is no longer used for the KMS console of the new version.
//
// # [](#qps-)QPS limits
//
// This operation can be called up to 10 times per second for each Alibaba Cloud account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSourcesResponse
func (client *Client) DescribeInstanceSourcesWithOptions(request *DescribeInstanceSourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditStatus)) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AuthStatus)) {
		query["AuthStatus"] = request.AuthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EngineType)) {
		query["EngineType"] = request.EngineType
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.SearchType)) {
		query["SearchType"] = request.SearchType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRegionId)) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceSources"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data assets.
//
// Description:
//
// You can query a list of unauthorized or authorized data assets based on the value of AuthStatus.
//
// This operation is no longer used for the KMS console of the new version.
//
// # [](#qps-)QPS limits
//
// This operation can be called up to 10 times per second for each Alibaba Cloud account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceSourcesRequest
//
// @return DescribeInstanceSourcesResponse
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

// Summary:
//
// Queries data assets such as MaxCompute, ApsaraDB RDS, and Object Storage Service (OSS) that you authorize Data Security Center (DSC) to access.
//
// Description:
//
// When you call the DescribeInstances operation, you can specify parameters such as Name and RiskLevelId to query data assets that meet filter conditions.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevelId)) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRegionId)) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstances"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data assets such as MaxCompute, ApsaraDB RDS, and Object Storage Service (OSS) that you authorize Data Security Center (DSC) to access.
//
// Description:
//
// When you call the DescribeInstances operation, you can specify parameters such as Name and RiskLevelId to query data assets that meet filter conditions.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
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

// Summary:
//
// Queries the details of an Object Storage Service (OSS) object that Data Security Center (DSC) is authorized to access.
//
// Description:
//
// You can call this operation to query the details of an Object Storage Service (OSS) object. This helps you locate sensitive data detected in OSS.
//
// ## [](#)Precautions
//
// The DescribeOssObjectDetail operation is chagned to DescribeOssObjectDetailV2. We recommend that you call the DescribeOssObjectDetailV2 operation when you develop your applications.
//
// ## [](#qps)Limits
//
// Each Alibaba Cloud account can call this operation up to 10 times per second. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeOssObjectDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssObjectDetailResponse
func (client *Client) DescribeOssObjectDetailWithOptions(request *DescribeOssObjectDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeOssObjectDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssObjectDetail"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssObjectDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an Object Storage Service (OSS) object that Data Security Center (DSC) is authorized to access.
//
// Description:
//
// You can call this operation to query the details of an Object Storage Service (OSS) object. This helps you locate sensitive data detected in OSS.
//
// ## [](#)Precautions
//
// The DescribeOssObjectDetail operation is chagned to DescribeOssObjectDetailV2. We recommend that you call the DescribeOssObjectDetailV2 operation when you develop your applications.
//
// ## [](#qps)Limits
//
// Each Alibaba Cloud account can call this operation up to 10 times per second. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeOssObjectDetailRequest
//
// @return DescribeOssObjectDetailResponse
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

// Summary:
//
// Call this interface to query the details of a single storage object in OSS that is authorized by the Data Security Center.
//
// Description:
//
// This interface is generally used to query the detailed information of OSS storage objects, which facilitates the accurate positioning of sensitive OSS assets.
//
// @param request - DescribeOssObjectDetailV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssObjectDetailV2Response
func (client *Client) DescribeOssObjectDetailV2WithOptions(request *DescribeOssObjectDetailV2Request, runtime *util.RuntimeOptions) (_result *DescribeOssObjectDetailV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectKey)) {
		query["ObjectKey"] = request.ObjectKey
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRegionId)) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssObjectDetailV2"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssObjectDetailV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call this interface to query the details of a single storage object in OSS that is authorized by the Data Security Center.
//
// Description:
//
// This interface is generally used to query the detailed information of OSS storage objects, which facilitates the accurate positioning of sensitive OSS assets.
//
// @param request - DescribeOssObjectDetailV2Request
//
// @return DescribeOssObjectDetailV2Response
func (client *Client) DescribeOssObjectDetailV2(request *DescribeOssObjectDetailV2Request) (_result *DescribeOssObjectDetailV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssObjectDetailV2Response{}
	_body, _err := client.DescribeOssObjectDetailV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Object Storage Service (OSS) objects that you authorize Data Security Center (DSC) to access.
//
// @param request - DescribeOssObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssObjectsResponse
func (client *Client) DescribeOssObjectsWithOptions(request *DescribeOssObjectsRequest, runtime *util.RuntimeOptions) (_result *DescribeOssObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FileCategoryCode)) {
		query["FileCategoryCode"] = request.FileCategoryCode
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LastScanTimeEnd)) {
		query["LastScanTimeEnd"] = request.LastScanTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.LastScanTimeStart)) {
		query["LastScanTimeStart"] = request.LastScanTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevelId)) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRegionId)) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssObjects"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Object Storage Service (OSS) objects that you authorize Data Security Center (DSC) to access.
//
// @param request - DescribeOssObjectsRequest
//
// @return DescribeOssObjectsResponse
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

// Summary:
//
// Queries information about the MaxCompute packages that Data Security Center (DSC) is authorized to access. The information includes the names of MaxCompute packages, the accounts of MaxCompute package owners, and the sensitivity levels of MaxCompute packages.
//
// Description:
//
// You can call this operation to query MaxCompute packages that are scanned by DSC. This helps you search for MaxCompute packages and view the summary of MaxCompute packages.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackagesResponse
func (client *Client) DescribePackagesWithOptions(request *DescribePackagesRequest, runtime *util.RuntimeOptions) (_result *DescribePackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevelId)) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePackages"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the MaxCompute packages that Data Security Center (DSC) is authorized to access. The information includes the names of MaxCompute packages, the accounts of MaxCompute package owners, and the sensitivity levels of MaxCompute packages.
//
// Description:
//
// You can call this operation to query MaxCompute packages that are scanned by DSC. This helps you search for MaxCompute packages and view the summary of MaxCompute packages.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePackagesRequest
//
// @return DescribePackagesResponse
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

// Summary:
//
// Gets the list of first-level authorizations.
//
// @param request - DescribeParentInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParentInstanceResponse
func (client *Client) DescribeParentInstanceWithOptions(request *DescribeParentInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeParentInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthStatus)) {
		query["AuthStatus"] = request.AuthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CheckStatus)) {
		query["CheckStatus"] = request.CheckStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterStatus)) {
		query["ClusterStatus"] = request.ClusterStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		query["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EngineType)) {
		query["EngineType"] = request.EngineType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberAccount)) {
		query["MemberAccount"] = request.MemberAccount
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRegionId)) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParentInstance"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeParentInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the list of first-level authorizations.
//
// @param request - DescribeParentInstanceRequest
//
// @return DescribeParentInstanceResponse
func (client *Client) DescribeParentInstance(request *DescribeParentInstanceRequest) (_result *DescribeParentInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParentInstanceResponse{}
	_body, _err := client.DescribeParentInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the sensitivity levels that are defined in a rule template provided by Data Security Center (DSC).
//
// Description:
//
// You can call this operation to query the sensitivity levels that are defined in the current rule template provided by DSC. This helps you learn about the number of times that each sensitivity level is referenced in the rule template and the highest sensitivity level.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeRiskLevelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskLevelsResponse
func (client *Client) DescribeRiskLevelsWithOptions(request *DescribeRiskLevelsRequest, runtime *util.RuntimeOptions) (_result *DescribeRiskLevelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRiskLevels"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRiskLevelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the sensitivity levels that are defined in a rule template provided by Data Security Center (DSC).
//
// Description:
//
// You can call this operation to query the sensitivity levels that are defined in the current rule template provided by DSC. This helps you learn about the number of times that each sensitivity level is referenced in the rule template and the highest sensitivity level.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeRiskLevelsRequest
//
// @return DescribeRiskLevelsResponse
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

// Summary:
//
// Queries sensitive data detection rules.
//
// @param request - DescribeRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRulesResponse
func (client *Client) DescribeRulesWithOptions(request *DescribeRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ContentCategory)) {
		query["ContentCategory"] = request.ContentCategory
	}

	if !tea.BoolValue(util.IsUnset(request.CooperationChannel)) {
		query["CooperationChannel"] = request.CooperationChannel
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.CustomType)) {
		query["CustomType"] = request.CustomType
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordCompatible)) {
		query["KeywordCompatible"] = request.KeywordCompatible
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevelId)) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.Simplify)) {
		query["Simplify"] = request.Simplify
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SupportForm)) {
		query["SupportForm"] = request.SupportForm
	}

	if !tea.BoolValue(util.IsUnset(request.WarnLevel)) {
		query["WarnLevel"] = request.WarnLevel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRules"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries sensitive data detection rules.
//
// @param request - DescribeRulesRequest
//
// @return DescribeRulesResponse
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

// Summary:
//
// Queries tables in data assets, such as MaxCompute projects and ApsaraDB RDS instances, that you authorize Data Security Center (DSC) to access.
//
// Description:
//
// When you call the DescribeTables operation to query tables, you can specify parameters such as Name and RiskLevelId to filter tables.
//
// # Limits
//
// You can send up to 10 requests per second to call this operation by using your Alibaba Cloud account. If you send excessive requests, throttling is implemented, and your business may be affected.
//
// @param request - DescribeTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTablesResponse
func (client *Client) DescribeTablesWithOptions(request *DescribeTablesRequest, runtime *util.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PackageId)) {
		query["PackageId"] = request.PackageId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevelId)) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRegionId)) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTables"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tables in data assets, such as MaxCompute projects and ApsaraDB RDS instances, that you authorize Data Security Center (DSC) to access.
//
// Description:
//
// When you call the DescribeTables operation to query tables, you can specify parameters such as Name and RiskLevelId to filter tables.
//
// # Limits
//
// You can send up to 10 requests per second to call this operation by using your Alibaba Cloud account. If you send excessive requests, throttling is implemented, and your business may be affected.
//
// @param request - DescribeTablesRequest
//
// @return DescribeTablesResponse
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

// Summary:
//
// Call this interface to query all models list of industry templates.
//
// @param request - DescribeTemplateAllRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplateAllRulesResponse
func (client *Client) DescribeTemplateAllRulesWithOptions(request *DescribeTemplateAllRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeTemplateAllRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTemplateAllRules"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTemplateAllRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call this interface to query all models list of industry templates.
//
// @param request - DescribeTemplateAllRulesRequest
//
// @return DescribeTemplateAllRulesResponse
func (client *Client) DescribeTemplateAllRules(request *DescribeTemplateAllRulesRequest) (_result *DescribeTemplateAllRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTemplateAllRulesResponse{}
	_body, _err := client.DescribeTemplateAllRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an account.
//
// Description:
//
// You can call this operation to query the information about the current account. This helps you get familiar with your account that accesses Data Security Center (DSC).
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserStatusResponse
func (client *Client) DescribeUserStatusWithOptions(request *DescribeUserStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserStatus"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an account.
//
// Description:
//
// You can call this operation to query the information about the current account. This helps you get familiar with your account that accesses Data Security Center (DSC).
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeUserStatusRequest
//
// @return DescribeUserStatusResponse
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

// Summary:
//
// Disables a configuration item. After you disable a configuration item, you can call the CreateConfig operation to enable the configuration item by specifying the code of the configuration item for the Code parameter in the request.
//
// Description:
//
// You can call this operation to disable a configuration item based on the code of the configuration item. This helps you modify configurations at the earliest opportunity.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DisableUserConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableUserConfigResponse
func (client *Client) DisableUserConfigWithOptions(request *DisableUserConfigRequest, runtime *util.RuntimeOptions) (_result *DisableUserConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableUserConfig"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableUserConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a configuration item. After you disable a configuration item, you can call the CreateConfig operation to enable the configuration item by specifying the code of the configuration item for the Code parameter in the request.
//
// Description:
//
// You can call this operation to disable a configuration item based on the code of the configuration item. This helps you modify configurations at the earliest opportunity.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DisableUserConfigRequest
//
// @return DisableUserConfigResponse
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

// Summary:
//
// Dynamically de-identifies sensitive data.
//
// @param request - ExecDatamaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecDatamaskResponse
func (client *Client) ExecDatamaskWithOptions(request *ExecDatamaskRequest, runtime *util.RuntimeOptions) (_result *ExecDatamaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecDatamask"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecDatamaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dynamically de-identifies sensitive data.
//
// @param request - ExecDatamaskRequest
//
// @return ExecDatamaskResponse
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

// Summary:
//
// Triggers a de-identification task.
//
// @param request - ManualTriggerMaskingProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManualTriggerMaskingProcessResponse
func (client *Client) ManualTriggerMaskingProcessWithOptions(request *ManualTriggerMaskingProcessRequest, runtime *util.RuntimeOptions) (_result *ManualTriggerMaskingProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ManualTriggerMaskingProcess"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ManualTriggerMaskingProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers a de-identification task.
//
// @param request - ManualTriggerMaskingProcessRequest
//
// @return ManualTriggerMaskingProcessResponse
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

// Summary:
//
// Modifies configuration items for a data asset that you authorize Data Security Center (DSC) to access.
//
// @param request - ModifyDataLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataLimitResponse
func (client *Client) ModifyDataLimitWithOptions(request *ModifyDataLimitRequest, runtime *util.RuntimeOptions) (_result *ModifyDataLimitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditStatus)) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AutoScan)) {
		query["AutoScan"] = request.AutoScan
	}

	if !tea.BoolValue(util.IsUnset(request.EngineType)) {
		query["EngineType"] = request.EngineType
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LogStoreDay)) {
		query["LogStoreDay"] = request.LogStoreDay
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyPassword)) {
		query["ModifyPassword"] = request.ModifyPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SamplingSize)) {
		query["SamplingSize"] = request.SamplingSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupIdList)) {
		query["SecurityGroupIdList"] = request.SecurityGroupIdList
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRegionId)) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchIdList)) {
		query["VSwitchIdList"] = request.VSwitchIdList
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDataLimit"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDataLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies configuration items for a data asset that you authorize Data Security Center (DSC) to access.
//
// @param request - ModifyDataLimitRequest
//
// @return ModifyDataLimitResponse
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

// Summary:
//
// Changes the sensitivity levels of sensitive data. You can change the default sensitivity levels of data that cannot be classified as sensitive or insensitive, and the sensitivity levels of data that can be classified as sensitive.
//
// Description:
//
// You can call this operation to modify the sensitivity levels of data. This helps you manage the sensitivity levels.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyDefaultLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefaultLevelResponse
func (client *Client) ModifyDefaultLevelWithOptions(request *ModifyDefaultLevelRequest, runtime *util.RuntimeOptions) (_result *ModifyDefaultLevelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefaultId)) {
		query["DefaultId"] = request.DefaultId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SensitiveIds)) {
		query["SensitiveIds"] = request.SensitiveIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefaultLevel"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefaultLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the sensitivity levels of sensitive data. You can change the default sensitivity levels of data that cannot be classified as sensitive or insensitive, and the sensitivity levels of data that can be classified as sensitive.
//
// Description:
//
// You can call this operation to modify the sensitivity levels of data. This helps you manage the sensitivity levels.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyDefaultLevelRequest
//
// @return ModifyDefaultLevelResponse
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

// Summary:
//
// Handles an anomalous event.
//
// Description:
//
// You can call this operation to handle anomalous events that involve data leaks. This helps protect your data assets at the earliest opportunity.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyEventStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEventStatusResponse
func (client *Client) ModifyEventStatusWithOptions(request *ModifyEventStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyEventStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Backed)) {
		query["Backed"] = request.Backed
	}

	if !tea.BoolValue(util.IsUnset(request.DealReason)) {
		query["DealReason"] = request.DealReason
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyEventStatus"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyEventStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Handles an anomalous event.
//
// Description:
//
// You can call this operation to handle anomalous events that involve data leaks. This helps protect your data assets at the earliest opportunity.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyEventStatusRequest
//
// @return ModifyEventStatusResponse
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

// Summary:
//
// Enables the detection of anomalous events of subtypes.
//
// @param request - ModifyEventTypeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEventTypeStatusResponse
func (client *Client) ModifyEventTypeStatusWithOptions(request *ModifyEventTypeStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyEventTypeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SubTypeIds)) {
		query["SubTypeIds"] = request.SubTypeIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyEventTypeStatus"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyEventTypeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the detection of anomalous events of subtypes.
//
// @param request - ModifyEventTypeStatusRequest
//
// @return ModifyEventTypeStatusResponse
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

// Summary:
//
// Enables or disables the report task.
//
// Description:
//
// You can call this operation to enable or disable the report task. After you activate Data Security Center (DSC), the report task is enabled by default. After you disable the report task, you cannot view statistics that are newly generated in the Report Center module, on the Overview page of the Cloud Native Data Audit module, and in the Data security lab module. Existing statistics are not affected.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyReportTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyReportTaskStatusResponse
func (client *Client) ModifyReportTaskStatusWithOptions(request *ModifyReportTaskStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyReportTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureType)) {
		query["FeatureType"] = request.FeatureType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ReportTaskStatus)) {
		query["ReportTaskStatus"] = request.ReportTaskStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyReportTaskStatus"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyReportTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the report task.
//
// Description:
//
// You can call this operation to enable or disable the report task. After you activate Data Security Center (DSC), the report task is enabled by default. After you disable the report task, you cannot view statistics that are newly generated in the Report Center module, on the Overview page of the Cloud Native Data Audit module, and in the Data security lab module. Existing statistics are not affected.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyReportTaskStatusRequest
//
// @return ModifyReportTaskStatusResponse
func (client *Client) ModifyReportTaskStatus(request *ModifyReportTaskStatusRequest) (_result *ModifyReportTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyReportTaskStatusResponse{}
	_body, _err := client.ModifyReportTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a custom sensitive data detection rule in Data Security Center (DSC).
//
// Description:
//
// When you call this operation, you must configure request parameters to specify the rule name, rule ID, and rule content.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRuleResponse
func (client *Client) ModifyRuleWithOptions(request *ModifyRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.ModelRuleIds)) {
		query["ModelRuleIds"] = request.ModelRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevelId)) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.SupportForm)) {
		query["SupportForm"] = request.SupportForm
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateRuleIds)) {
		query["TemplateRuleIds"] = request.TemplateRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.WarnLevel)) {
		query["WarnLevel"] = request.WarnLevel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRule"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a custom sensitive data detection rule in Data Security Center (DSC).
//
// Description:
//
// When you call this operation, you must configure request parameters to specify the rule name, rule ID, and rule content.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyRuleRequest
//
// @return ModifyRuleResponse
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

// Summary:
//
// Enables or disables a sensitive data detection rule.
//
// @param request - ModifyRuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRuleStatusResponse
func (client *Client) ModifyRuleStatusWithOptions(request *ModifyRuleStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyRuleStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRuleStatus"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyRuleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables a sensitive data detection rule.
//
// @param request - ModifyRuleStatusRequest
//
// @return ModifyRuleStatusResponse
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

// Summary:
//
// Creates an identification task to scan sensitive data in Object Storage Service (OSS) objects.
//
// Description:
//
// ### [](#)Prerequisites
//
// To call this operation, make sure that asset authorization for your OSS bucket is complete and the bucket is connected. If the authorization is not complete, the bucket_not_authorized error code is returned when you call the operation.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// ### [](#)Additional information
//
// After you call this operation, you can obtain the task ID. You can specify the task ID in the DescribeIdentifyTaskDetail operation to query the state of the task.
//
// After the task is complete, you can call the DescribeOssObjectDetailV2 operation to query the identification results of sensitive data in the related OSS objects. When you call the DescribeOssObjectDetailV2 operation, you must specify BucketName, ServiceRegionId, and ObjectKey.
//
// @param tmpReq - ScanOssObjectV1Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScanOssObjectV1Response
func (client *Client) ScanOssObjectV1WithOptions(tmpReq *ScanOssObjectV1Request, runtime *util.RuntimeOptions) (_result *ScanOssObjectV1Response, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ScanOssObjectV1ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ObjectKeyList)) {
		request.ObjectKeyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectKeyList, tea.String("ObjectKeyList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectKeyListShrink)) {
		query["ObjectKeyList"] = request.ObjectKeyListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceRegionId)) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ScanOssObjectV1"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ScanOssObjectV1Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an identification task to scan sensitive data in Object Storage Service (OSS) objects.
//
// Description:
//
// ### [](#)Prerequisites
//
// To call this operation, make sure that asset authorization for your OSS bucket is complete and the bucket is connected. If the authorization is not complete, the bucket_not_authorized error code is returned when you call the operation.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// ### [](#)Additional information
//
// After you call this operation, you can obtain the task ID. You can specify the task ID in the DescribeIdentifyTaskDetail operation to query the state of the task.
//
// After the task is complete, you can call the DescribeOssObjectDetailV2 operation to query the identification results of sensitive data in the related OSS objects. When you call the DescribeOssObjectDetailV2 operation, you must specify BucketName, ServiceRegionId, and ObjectKey.
//
// @param request - ScanOssObjectV1Request
//
// @return ScanOssObjectV1Response
func (client *Client) ScanOssObjectV1(request *ScanOssObjectV1Request) (_result *ScanOssObjectV1Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScanOssObjectV1Response{}
	_body, _err := client.ScanOssObjectV1WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a de-identification task. After you stop a de-identification task, you can resume the task by calling the ManualTriggerMaskingProcess operation.
//
// Description:
//
// You can call this operation to stop a de-identification task that is running. For example, you can stop a de-identification task that is used to de-identify specific data.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - StopMaskingProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMaskingProcessResponse
func (client *Client) StopMaskingProcessWithOptions(request *StopMaskingProcessRequest, runtime *util.RuntimeOptions) (_result *StopMaskingProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopMaskingProcess"),
		Version:     tea.String("2019-01-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopMaskingProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a de-identification task. After you stop a de-identification task, you can resume the task by calling the ManualTriggerMaskingProcess operation.
//
// Description:
//
// You can call this operation to stop a de-identification task that is running. For example, you can stop a de-identification task that is used to de-identify specific data.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - StopMaskingProcessRequest
//
// @return StopMaskingProcessResponse
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

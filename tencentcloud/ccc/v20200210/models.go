// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20200210

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BindStaffSkillGroupListRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 坐席邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 绑定技能组列表
	SkillGroupList []*int64 `json:"SkillGroupList,omitempty" name:"SkillGroupList"`
}

func (r *BindStaffSkillGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindStaffSkillGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffEmail")
	delete(f, "SkillGroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindStaffSkillGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindStaffSkillGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindStaffSkillGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindStaffSkillGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallInMetrics struct {

	// IVR驻留数量
	IvrCount *int64 `json:"IvrCount,omitempty" name:"IvrCount"`

	// 排队中数量
	QueueCount *int64 `json:"QueueCount,omitempty" name:"QueueCount"`

	// 振铃中数量
	RingCount *int64 `json:"RingCount,omitempty" name:"RingCount"`

	// 接通中数量
	AcceptCount *int64 `json:"AcceptCount,omitempty" name:"AcceptCount"`

	// 客服转接外线中数量
	TransferOuterCount *int64 `json:"TransferOuterCount,omitempty" name:"TransferOuterCount"`

	// 最大排队时长
	MaxQueueDuration *int64 `json:"MaxQueueDuration,omitempty" name:"MaxQueueDuration"`

	// 平均排队时长
	AvgQueueDuration *int64 `json:"AvgQueueDuration,omitempty" name:"AvgQueueDuration"`

	// 最大振铃时长
	MaxRingDuration *int64 `json:"MaxRingDuration,omitempty" name:"MaxRingDuration"`

	// 平均振铃时长
	AvgRingDuration *int64 `json:"AvgRingDuration,omitempty" name:"AvgRingDuration"`

	// 最大接通时长
	MaxAcceptDuration *int64 `json:"MaxAcceptDuration,omitempty" name:"MaxAcceptDuration"`

	// 平均接通时长
	AvgAcceptDuration *int64 `json:"AvgAcceptDuration,omitempty" name:"AvgAcceptDuration"`
}

type CallInNumberMetrics struct {

	// 线路号码
	Number *string `json:"Number,omitempty" name:"Number"`

	// 线路相关指标
	Metrics *CallInMetrics `json:"Metrics,omitempty" name:"Metrics"`

	// 所属技能组相关指标
	SkillGroupMetrics []*CallInSkillGroupMetrics `json:"SkillGroupMetrics,omitempty" name:"SkillGroupMetrics"`
}

type CallInSkillGroupMetrics struct {

	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 数据指标
	Metrics *CallInMetrics `json:"Metrics,omitempty" name:"Metrics"`

	// 技能组名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CreateSDKLoginTokenRequest struct {
	*tchttp.BaseRequest

	// 应用ID。
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 坐席账号。
	SeatUserId *string `json:"SeatUserId,omitempty" name:"SeatUserId"`
}

func (r *CreateSDKLoginTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSDKLoginTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SeatUserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSDKLoginTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSDKLoginTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SDK 登录 Token。
		Token *string `json:"Token,omitempty" name:"Token"`

		// 过期时间戳，Unix 时间戳。
		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

		// SDK 加载路径会随着 SDK 的发布而变动。
		SdkURL *string `json:"SdkURL,omitempty" name:"SdkURL"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSDKLoginTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSDKLoginTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStaffRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 客服信息，个数不超过 10
	Staffs []*SeatUserInfo `json:"Staffs,omitempty" name:"Staffs"`
}

func (r *CreateStaffRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaffRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Staffs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStaffRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateStaffResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误坐席列表及错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrorStaffList []*ErrStaffItem `json:"ErrorStaffList,omitempty" name:"ErrorStaffList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStaffResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserSigRequest struct {
	*tchttp.BaseRequest

	// 应用 ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 用户 ID
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 有效期，单位秒，不超过 1 小时
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 用户签名数据
	ClientData *string `json:"ClientData,omitempty" name:"ClientData"`
}

func (r *CreateUserSigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Uid")
	delete(f, "ExpiredTime")
	delete(f, "ClientData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserSigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserSigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 签名结果
		UserSig *string `json:"UserSig,omitempty" name:"UserSig"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserSigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStaffRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 待删除客服邮箱列表
	StaffList []*string `json:"StaffList,omitempty" name:"StaffList"`
}

func (r *DeleteStaffRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStaffRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStaffRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStaffResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 无法删除的状态为在线的客服列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		OnlineStaffList []*string `json:"OnlineStaffList,omitempty" name:"OnlineStaffList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteStaffResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStaffResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCCBuyInfoListRequest struct {
	*tchttp.BaseRequest

	// 应用ID列表，不传时查询所有应用
	SdkAppIds []*int64 `json:"SdkAppIds,omitempty" name:"SdkAppIds"`
}

func (r *DescribeCCCBuyInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCCBuyInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCCBuyInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCCBuyInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 应用购买信息列表
		SdkAppIdBuyList []*SdkAppIdBuyInfo `json:"SdkAppIdBuyList,omitempty" name:"SdkAppIdBuyList"`

		// 套餐包购买信息列表
		PackageBuyList []*PackageBuyInfo `json:"PackageBuyList,omitempty" name:"PackageBuyList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCCBuyInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCCBuyInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCallInMetricsRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 是否返回技能组维度信息，默认“是”
	EnabledSkillGroup *bool `json:"EnabledSkillGroup,omitempty" name:"EnabledSkillGroup"`

	// 是否返回线路维度信息，默认“否”
	EnabledNumber *bool `json:"EnabledNumber,omitempty" name:"EnabledNumber"`
}

func (r *DescribeCallInMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallInMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "EnabledSkillGroup")
	delete(f, "EnabledNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallInMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCallInMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 时间戳
		Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

		// 总体指标
		TotalMetrics *CallInMetrics `json:"TotalMetrics,omitempty" name:"TotalMetrics"`

		// 线路维度指标
	// 注意：此字段可能返回 null，表示取不到有效值。
		NumberMetrics []*CallInNumberMetrics `json:"NumberMetrics,omitempty" name:"NumberMetrics"`

		// 技能组维度指标
	// 注意：此字段可能返回 null，表示取不到有效值。
		SkillGroupMetrics []*CallInSkillGroupMetrics `json:"SkillGroupMetrics,omitempty" name:"SkillGroupMetrics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCallInMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallInMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChatMessagesRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 服务记录ID
	CdrId *string `json:"CdrId,omitempty" name:"CdrId"`

	// 返回记录条数 最大为100默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回记录偏移 默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1为从早到晚，2为从晚到早，默认为2
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 服务记录SessionID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

func (r *DescribeChatMessagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChatMessagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SdkAppId")
	delete(f, "CdrId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChatMessagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChatMessagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 消息列表
		Messages []*MessageBody `json:"Messages,omitempty" name:"Messages"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeChatMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChatMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIMCdrsRequest struct {
	*tchttp.BaseRequest

	// 起始时间
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 结束时间
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" name:"EndTimestamp"`

	// 实例ID
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 返回记录条数 最大为100默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 返回记录偏移 默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 1为全媒体，2为文本客服，不填则查询全部
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeIMCdrsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIMCdrsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimestamp")
	delete(f, "EndTimestamp")
	delete(f, "InstanceId")
	delete(f, "SdkAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIMCdrsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIMCdrsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 服务记录列表
		IMCdrs []*IMCdrInfo `json:"IMCdrs,omitempty" name:"IMCdrs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIMCdrsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIMCdrsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePSTNActiveSessionListRequest struct {
	*tchttp.BaseRequest

	// 应用 ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 数据偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回的数据条数，最大 25
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePSTNActiveSessionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePSTNActiveSessionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePSTNActiveSessionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePSTNActiveSessionListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表总条数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 列表内容
		Sessions []*PSTNSessionInfo `json:"Sessions,omitempty" name:"Sessions"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePSTNActiveSessionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePSTNActiveSessionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSeatUserListRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSeatUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSeatUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSeatUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSeatUserListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 此实例的坐席用户总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 坐席用户信息列表
		SeatUsers []*SeatUserInfo `json:"SeatUsers,omitempty" name:"SeatUsers"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSeatUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSeatUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSkillGroupInfoListRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 技能组ID，查询单个技能组时使用
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 查询修改时间大于等于ModifiedTime的技能组时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

func (r *DescribeSkillGroupInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillGroupInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "SkillGroupId")
	delete(f, "ModifiedTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillGroupInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSkillGroupInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 技能组总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 技能组信息列表
		SkillGroupList []*SkillGroupInfoItem `json:"SkillGroupList,omitempty" name:"SkillGroupList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSkillGroupInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillGroupInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStaffInfoListRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 坐席账号，查询单个坐席时使用
	StaffMail *string `json:"StaffMail,omitempty" name:"StaffMail"`

	// 查询修改时间大于等于ModifiedTime的坐席时使用
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

func (r *DescribeStaffInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "StaffMail")
	delete(f, "ModifiedTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStaffInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStaffInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 坐席用户总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 坐席用户信息列表
		StaffList []*StaffInfo `json:"StaffList,omitempty" name:"StaffList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStaffInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStaffStatusMetricsRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 筛选坐席列表，默认不传返回全部坐席信息
	StaffList []*string `json:"StaffList,omitempty" name:"StaffList"`
}

func (r *DescribeStaffStatusMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffStatusMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStaffStatusMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStaffStatusMetricsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 坐席状态实时信息
		Metrics []*StaffStatusMetrics `json:"Metrics,omitempty" name:"Metrics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStaffStatusMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaffStatusMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTelCallInfoRequest struct {
	*tchttp.BaseRequest

	// 起始时间戳，Unix 时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 时间戳，查询时间范围最大为90天
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 应用ID列表，多个ID时，返回值为多个ID使用总和
	SdkAppIdList []*int64 `json:"SdkAppIdList,omitempty" name:"SdkAppIdList"`
}

func (r *DescribeTelCallInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCallInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	delete(f, "SdkAppIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelCallInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTelCallInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 电话呼出统计分钟数
		TelCallOutCount *int64 `json:"TelCallOutCount,omitempty" name:"TelCallOutCount"`

		// 电话呼入统计分钟数
		TelCallInCount *int64 `json:"TelCallInCount,omitempty" name:"TelCallInCount"`

		// 坐席使用统计个数
		SeatUsedCount *int64 `json:"SeatUsedCount,omitempty" name:"SeatUsedCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTelCallInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCallInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTelCdrRequest struct {
	*tchttp.BaseRequest

	// 起始时间戳，Unix 时间戳
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" name:"StartTimeStamp"`

	// 结束时间戳，Unix 时间戳
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" name:"EndTimeStamp"`

	// 返回数据条数，上限（deprecated）
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移（deprecated）
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 实例 ID（deprecated）
	InstanceId *int64 `json:"InstanceId,omitempty" name:"InstanceId"`

	// 应用 ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 分页尺寸，上限 100
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 分页页码，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 按手机号筛选
	Phones []*string `json:"Phones,omitempty" name:"Phones"`

	// 按SessionId筛选
	SessionIds []*string `json:"SessionIds,omitempty" name:"SessionIds"`
}

func (r *DescribeTelCdrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCdrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStamp")
	delete(f, "EndTimeStamp")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceId")
	delete(f, "SdkAppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Phones")
	delete(f, "SessionIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelCdrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTelCdrResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 话单记录总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 话单记录
		TelCdrs []*TelCdrInfo `json:"TelCdrs,omitempty" name:"TelCdrs"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTelCdrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelCdrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTelSessionRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 会话ID
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

func (r *DescribeTelSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTelSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTelSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 会话信息
		Session *PSTNSession `json:"Session,omitempty" name:"Session"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTelSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTelSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrStaffItem struct {

	// 坐席邮箱地址
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 错误描述
	Message *string `json:"Message,omitempty" name:"Message"`
}

type IMCdrInfo struct {

	// 服务记录ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 服务时长秒数
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 结束状态
	EndStatus *int64 `json:"EndStatus,omitempty" name:"EndStatus"`

	// 用户昵称
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 服务类型 1为全媒体，2为文本客服
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 客服ID
	StaffId *string `json:"StaffId,omitempty" name:"StaffId"`

	// 服务时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}

type IVRKeyPressedElement struct {

	// 按键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 按键关联的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`
}

type Message struct {

	// 消息类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 消息内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

type MessageBody struct {

	// 消息时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 发消息的用户ID
	From *string `json:"From,omitempty" name:"From"`

	// 消息列表
	Messages []*Message `json:"Messages,omitempty" name:"Messages"`
}

type PSTNSession struct {

	// 会话 ID
	SessionID *string `json:"SessionID,omitempty" name:"SessionID"`

	// 会话临时房间 ID
	RoomID *string `json:"RoomID,omitempty" name:"RoomID"`

	// 主叫
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 被叫
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 开始时间，Unix 时间戳
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 振铃时间，Unix 时间戳
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`

	// 接听时间，Unix 时间戳
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 坐席邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 坐席工号
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`

	// 会话状态
	// ringing 振铃中
	// seatJoining  等待坐席接听
	// inProgress 进行中
	// finished 已完成
	SessionStatus *string `json:"SessionStatus,omitempty" name:"SessionStatus"`

	// 会话呼叫方向， 0 呼入 | 1 - 呼出
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// 转外线使用的号码（转外线主叫）
	OutBoundCaller *string `json:"OutBoundCaller,omitempty" name:"OutBoundCaller"`

	// 转外线被叫
	OutBoundCallee *string `json:"OutBoundCallee,omitempty" name:"OutBoundCallee"`

	// 主叫号码保护ID，开启号码保护映射功能时有效，且Caller字段置空
	ProtectedCaller *string `json:"ProtectedCaller,omitempty" name:"ProtectedCaller"`

	// 被叫号码保护ID，开启号码保护映射功能时有效，且Callee字段置空
	ProtectedCallee *string `json:"ProtectedCallee,omitempty" name:"ProtectedCallee"`
}

type PSTNSessionInfo struct {

	// 会话 ID
	SessionID *string `json:"SessionID,omitempty" name:"SessionID"`

	// 会话临时房间 ID
	RoomID *string `json:"RoomID,omitempty" name:"RoomID"`

	// 主叫
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 被叫
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 开始时间，Unix 时间戳
	StartTimestamp *string `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 接听时间，Unix 时间戳
	AcceptTimestamp *string `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 坐席邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 坐席工号
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`

	// 坐席状态 inProgress 进行中
	SessionStatus *string `json:"SessionStatus,omitempty" name:"SessionStatus"`

	// 会话呼叫方向， 0 呼入 | 1 - 呼出
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// 振铃时间，Unix 时间戳
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`

	// 主叫号码保护ID，开启号码保护映射功能时有效，且Caller字段置空
	ProtectedCaller *string `json:"ProtectedCaller,omitempty" name:"ProtectedCaller"`

	// 被叫号码保护ID，开启号码保护映射功能时有效，且Callee字段置空
	ProtectedCallee *string `json:"ProtectedCallee,omitempty" name:"ProtectedCallee"`
}

type PackageBuyInfo struct {

	// 套餐包Id
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 套餐包类型，0-外呼套餐包|1-400呼入套餐包
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 套餐包总量
	CapacitySize *int64 `json:"CapacitySize,omitempty" name:"CapacitySize"`

	// 套餐包剩余量
	CapacityRemain *int64 `json:"CapacityRemain,omitempty" name:"CapacityRemain"`

	// 购买时间戳
	BuyTime *int64 `json:"BuyTime,omitempty" name:"BuyTime"`

	// 截至时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type PhoneNumBuyInfo struct {

	// 电话号码
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 号码类型，0-固话|1-虚商号码|2-运营商号码|3-400号码
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 号码呼叫类型，1-呼入|2-呼出|3-呼入呼出
	CallType *int64 `json:"CallType,omitempty" name:"CallType"`

	// 购买时间戳
	BuyTime *int64 `json:"BuyTime,omitempty" name:"BuyTime"`

	// 截至时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 号码状态，1正常|2停用
	State *int64 `json:"State,omitempty" name:"State"`
}

type SdkAppIdBuyInfo struct {

	// 应用ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 应用名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐席购买数（还在有效期内）
	StaffBuyNum *int64 `json:"StaffBuyNum,omitempty" name:"StaffBuyNum"`

	// 坐席购买列表 （还在有效期内）
	StaffBuyList []*StaffBuyInfo `json:"StaffBuyList,omitempty" name:"StaffBuyList"`

	// 号码购买列表
	PhoneNumBuyList []*PhoneNumBuyInfo `json:"PhoneNumBuyList,omitempty" name:"PhoneNumBuyList"`
}

type SeatUserInfo struct {

	// 坐席名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐席邮箱
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 坐席电话号码（带0086前缀）
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 坐席昵称
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 坐席关联的技能组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupNameList []*string `json:"SkillGroupNameList,omitempty" name:"SkillGroupNameList"`

	// 工号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`
}

type ServeParticipant struct {

	// 坐席邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 坐席电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 振铃时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`

	// 接听时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 结束时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndedTimestamp *int64 `json:"EndedTimestamp,omitempty" name:"EndedTimestamp"`

	// 录音 ID，能够索引到坐席侧的录音
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// 参与者类型，"staffSeat", "outboundSeat", "staffPhoneSeat"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 转接来源坐席信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferFrom *string `json:"TransferFrom,omitempty" name:"TransferFrom"`

	// 转接去向坐席信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferTo *string `json:"TransferTo,omitempty" name:"TransferTo"`

	// 转接去向参与者类型，取值与 Type 一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferToType *string `json:"TransferToType,omitempty" name:"TransferToType"`

	// 技能组 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 结束状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndStatusString *string `json:"EndStatusString,omitempty" name:"EndStatusString"`

	// 录音 URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordURL *string `json:"RecordURL,omitempty" name:"RecordURL"`

	// 参与者序号，从 0 开始
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sequence *int64 `json:"Sequence,omitempty" name:"Sequence"`

	// 开始时间戳，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 技能组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`
}

type SkillGroupInfoItem struct {

	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`

	// 类型：IM、TEL、ALL（全媒体）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 会话分配策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoutePolicy *string `json:"RoutePolicy,omitempty" name:"RoutePolicy"`

	// 会话分配是否优先上次服务坐席
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsingLastSeat *int64 `json:"UsingLastSeat,omitempty" name:"UsingLastSeat"`

	// 单客服最大并发数（电话类型默认1）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" name:"MaxConcurrency"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTimestamp *int64 `json:"LastModifyTimestamp,omitempty" name:"LastModifyTimestamp"`
}

type SkillGroupItem struct {

	// 技能组ID
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// 技能组名称
	SkillGroupName *string `json:"SkillGroupName,omitempty" name:"SkillGroupName"`

	// 优先级
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 类型：IM、TEL、ALL（全媒体）
	Type *string `json:"Type,omitempty" name:"Type"`
}

type StaffBuyInfo struct {

	// 购买坐席数量
	Num *int64 `json:"Num,omitempty" name:"Num"`

	// 购买时间戳
	BuyTime *int64 `json:"BuyTime,omitempty" name:"BuyTime"`

	// 截至时间戳
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type StaffInfo struct {

	// 坐席名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 坐席邮箱
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 坐席电话号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 坐席昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nick *string `json:"Nick,omitempty" name:"Nick"`

	// 坐席工号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StaffNumber *string `json:"StaffNumber,omitempty" name:"StaffNumber"`

	// 所属技能组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupList []*SkillGroupItem `json:"SkillGroupList,omitempty" name:"SkillGroupList"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTimestamp *int64 `json:"LastModifyTimestamp,omitempty" name:"LastModifyTimestamp"`
}

type StaffStatusExtra struct {

	// im - 文本 | tel - 电话 | all - 全媒体
	Type *string `json:"Type,omitempty" name:"Type"`

	// in - 呼入 | out - 呼出
	Direct *string `json:"Direct,omitempty" name:"Direct"`
}

type StaffStatusMetrics struct {

	// 坐席邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 坐席状态 free 示闲 | busy 忙碌 | rest 小休 | notReady 示忙 | afterCallWork 话后调整 | offline 离线
	Status *string `json:"Status,omitempty" name:"Status"`

	// 坐席状态补充信息
	StatusExtra *StaffStatusExtra `json:"StatusExtra,omitempty" name:"StatusExtra"`

	// 当天在线总时长
	OnlineDuration *int64 `json:"OnlineDuration,omitempty" name:"OnlineDuration"`

	// 当天示闲总时长
	FreeDuration *int64 `json:"FreeDuration,omitempty" name:"FreeDuration"`

	// 当天忙碌总时长
	BusyDuration *int64 `json:"BusyDuration,omitempty" name:"BusyDuration"`

	// 当天示忙总时长
	NotReadyDuration *int64 `json:"NotReadyDuration,omitempty" name:"NotReadyDuration"`

	// 当天小休总时长
	RestDuration *int64 `json:"RestDuration,omitempty" name:"RestDuration"`

	// 当天话后调整总时长
	AfterCallWorkDuration *int64 `json:"AfterCallWorkDuration,omitempty" name:"AfterCallWorkDuration"`

	// 小休原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 是否预约小休
	ReserveRest *bool `json:"ReserveRest,omitempty" name:"ReserveRest"`

	// 是否预约示忙
	ReserveNotReady *bool `json:"ReserveNotReady,omitempty" name:"ReserveNotReady"`
}

type TelCdrInfo struct {

	// 主叫号码
	Caller *string `json:"Caller,omitempty" name:"Caller"`

	// 被叫号码
	Callee *string `json:"Callee,omitempty" name:"Callee"`

	// 呼叫发起时间戳，Unix 时间戳
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 呼入呼出方向 0 呼入 1 呼出
	Direction *int64 `json:"Direction,omitempty" name:"Direction"`

	// 通话时长
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 录音信息
	RecordURL *string `json:"RecordURL,omitempty" name:"RecordURL"`

	// 坐席信息
	SeatUser *SeatUserInfo `json:"SeatUser,omitempty" name:"SeatUser"`

	// 结束状态
	// 0	错误
	// 1	正常结束
	// 2	未接通
	// 17	坐席未接
	// 100	黑名单
	// 101	坐席转接
	// 102	IVR 期间用户放弃
	// 103	会话排队期间用户放弃
	// 104	会话振铃期间用户放弃
	// 105	无坐席在线
	// 106	非工作时间
	// 107	IVR后直接结束
	// 201	未知状态
	// 202	未接听
	// 203	拒接挂断
	// 204	关机
	// 205	空号
	// 206	通话中
	// 207	欠费
	// 208	运营商线路异常
	// 209	主叫取消
	// 210	不在服务区
	EndStatus *int64 `json:"EndStatus,omitempty" name:"EndStatus"`

	// 技能组名称
	SkillGroup *string `json:"SkillGroup,omitempty" name:"SkillGroup"`

	// 主叫归属地
	CallerLocation *string `json:"CallerLocation,omitempty" name:"CallerLocation"`

	// IVR 阶段耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	IVRDuration *int64 `json:"IVRDuration,omitempty" name:"IVRDuration"`

	// 振铃时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	RingTimestamp *int64 `json:"RingTimestamp,omitempty" name:"RingTimestamp"`

	// 接听时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	AcceptTimestamp *int64 `json:"AcceptTimestamp,omitempty" name:"AcceptTimestamp"`

	// 结束时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndedTimestamp *int64 `json:"EndedTimestamp,omitempty" name:"EndedTimestamp"`

	// IVR 按键信息 ，e.g. ["1","2","3"]
	// 注意：此字段可能返回 null，表示取不到有效值。
	IVRKeyPressed []*string `json:"IVRKeyPressed,omitempty" name:"IVRKeyPressed"`

	// 挂机方 seat 坐席 user 用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	HungUpSide *string `json:"HungUpSide,omitempty" name:"HungUpSide"`

	// 服务参与者列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServeParticipants []*ServeParticipant `json:"ServeParticipants,omitempty" name:"ServeParticipants"`

	// 技能组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" name:"SkillGroupId"`

	// error                   错误
	// ok                       正常结束
	// unconnected      未接通
	// seatGiveUp         坐席未接
	// blackList             黑名单
	// seatForward       坐席转接
	// ivrGiveUp           IVR 期间用户放弃
	// waitingGiveUp   会话排队期间用户放弃
	// ringingGiveUp   会话振铃期间用户放弃
	// noSeatOnline     无坐席在线
	// notWorkTime     非工作时间
	// ivrEnd                 IVR后直接结束
	// unknown            未知状态
	// notAnswer          未接听
	// userReject          拒接挂断
	// powerOff            关机
	// numberNotExist  空号
	// busy                    通话中
	// outOfCredit        欠费
	// operatorError     运营商线路异常
	// callerCancel        主叫取消
	// notInService       不在服务区
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndStatusString *string `json:"EndStatusString,omitempty" name:"EndStatusString"`

	// 会话开始时间戳，UNIX 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 进入排队时间，Unix 秒级时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueuedTimestamp *int64 `json:"QueuedTimestamp,omitempty" name:"QueuedTimestamp"`

	// 后置IVR按键信息（e.g. [{"Key":"1","Label":"非常满意"}]）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostIVRKeyPressed []*IVRKeyPressedElement `json:"PostIVRKeyPressed,omitempty" name:"PostIVRKeyPressed"`

	// 排队技能组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueuedSkillGroupId *int64 `json:"QueuedSkillGroupId,omitempty" name:"QueuedSkillGroupId"`

	// 会话 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 主叫号码保护ID，开启号码保护映射功能时有效，且Caller字段置空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectedCaller *string `json:"ProtectedCaller,omitempty" name:"ProtectedCaller"`

	// 被叫号码保护ID，开启号码保护映射功能时有效，且Callee字段置空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectedCallee *string `json:"ProtectedCallee,omitempty" name:"ProtectedCallee"`

	// 客户自定义数据（User-to-User Interface）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uui *string `json:"Uui,omitempty" name:"Uui"`

	// IVR按键信息（e.g. [{"Key":"1","Label":"非常满意"}]）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IVRKeyPressedEx []*IVRKeyPressedElement `json:"IVRKeyPressedEx,omitempty" name:"IVRKeyPressedEx"`
}

type UnbindStaffSkillGroupListRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	SdkAppId *int64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 客服邮箱
	StaffEmail *string `json:"StaffEmail,omitempty" name:"StaffEmail"`

	// 解绑技能组列表
	SkillGroupList []*int64 `json:"SkillGroupList,omitempty" name:"SkillGroupList"`
}

func (r *UnbindStaffSkillGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindStaffSkillGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StaffEmail")
	delete(f, "SkillGroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindStaffSkillGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnbindStaffSkillGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindStaffSkillGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindStaffSkillGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

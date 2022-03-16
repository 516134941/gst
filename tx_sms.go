package gst

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	tSms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type SmsSender struct {
	// SecretID
	SecretID string
	// SecretKey
	SecretKey string
	// SdkAppID APPID
	SdkAppID string
	// SignName 短信签名信息
	SignName string
}

// SendVerificationCode 验证码发送
// region为空默认广州节点, 北京 "ap-beijing" ,成都 "ap-chengdu",广州 "ap-guangzhou"
// tmpID 模板ID 不能为空
func (s *SmsSender) SendVerificationCode(mobile string, region string, tmpID string) (*tSms.SendSmsResponse, error) {
	credential := common.NewCredential(s.SecretID, s.SecretKey)
	if region == "" {
		region = regions.Guangzhou
	}
	if tmpID == "" {
		return nil, errors.New("模板ID不能为空")
	}
	client, err := tSms.NewClient(credential, region, profile.NewClientProfile())
	if err != nil {
		log.Error("SendMessage err", err)
		return nil, err
	}
	req := tSms.NewSendSmsRequest()
	sdkAppID := s.SdkAppID
	var mobiles []*string

	mobile = "+86" + mobile // 默认使用国内手机
	mobiles = append(mobiles, &mobile)
	signName := s.SignName

	req.PhoneNumberSet = mobiles
	req.SmsSdkAppId = &sdkAppID
	req.TemplateId = &tmpID
	req.SignName = &signName
	res, err := client.SendSms(req)
	if err != nil {
		log.Error("SendMessage err", err)
		return nil, err
	}
	return res, nil
}

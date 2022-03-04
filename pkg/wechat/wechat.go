package wechat

import (
	"github.com/fangjie-luoxi/tools/config"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/openplatform"
	openConfig "github.com/silenceper/wechat/v2/openplatform/config"
	"github.com/silenceper/wechat/v2/pay"
	payConfig "github.com/silenceper/wechat/v2/pay/config"
	"github.com/silenceper/wechat/v2/work"
	workConfig "github.com/silenceper/wechat/v2/work/config"
)

type Wechat struct {
	wc     *wechat.Wechat // 微信核心
	config *config.Config // 配置

	MiniProgram     *miniprogram.MiniProgram         // 微信小程序
	OfficialAccount *officialaccount.OfficialAccount // 微信公众号
	OpenPlat        *openplatform.OpenPlatform       // 微信开放平台
	Work            *work.Work                       // 企业微信
	Pay             *pay.Pay                         // 微信支付
}

func NewWechat(conf *config.Config) *Wechat {
	var wc Wechat

	wc.wc = wechat.NewWechat()
	if conf.String("wechat.mini.app_id") != "" {
		wc.MiniProgram = wc.wc.GetMiniProgram(&miniConfig.Config{
			AppID:     conf.String("wechat.mini.app_id"),
			AppSecret: conf.String("wechat.mini.app_secret"),
			Cache:     cache.NewMemory(),
		})
	}
	if conf.String("wechat.off.app_id") != "" {
		wc.OfficialAccount = wc.wc.GetOfficialAccount(&offConfig.Config{
			AppID:     conf.String("wechat.off.app_id"),
			AppSecret: conf.String("wechat.off.app_secret"),
			Token:     conf.String("wechat.off.token"),
			Cache:     cache.NewMemory(),
		})
	}
	if conf.String("wechat.open_plat.app_id") != "" {
		wc.OpenPlat = wc.wc.GetOpenPlatform(&openConfig.Config{
			AppID:     conf.String("wechat.open_plat.app_id"),
			AppSecret: conf.String("wechat.open_plat.app_secret"),
			Token:     conf.String("wechat.open_plat.token"),
			// EncodingAESKey: "",
			Cache: cache.NewMemory(),
		})
	}
	if conf.String("wechat.work.corp_id") != "" {
		wc.Work = wc.wc.GetWork(&workConfig.Config{
			CorpID:     conf.String("wechat.work.corp_id"),
			CorpSecret: conf.String("wechat.work.corp_secret"),
			AgentID:    conf.String("wechat.work.agent_id"),
			Cache:      cache.NewMemory(),
			//EncodingAESKey: conf.String("corp_id"),
		})
	}
	if conf.String("wechat.pay.app_id") != "" {
		wc.Pay = wc.wc.GetPay(&payConfig.Config{
			AppID:     conf.String("wechat.pay.app_id"),
			MchID:     conf.String("wechat.pay.mch_id"),
			Key:       conf.String("wechat.pay.key"),
			NotifyURL: conf.String("wechat.pay.notify_url"),
		})
	}
	return &wc
}

package member

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
	"zero.com/gtask-common/abstructs"
	"zero.com/gtask-common/cache"
	"zero.com/gtask-common/response/errors"
	"zero.com/gtask-common/response/mono"
	. "zero.com/gtask-common/session"
	"zero.com/gtask-common/utils"
)

// HandlerMember Member的路由处理函数
type HandlerMember struct {
	// 缓存客户端
	cacheClient abstructs.Cache
}

// New 创建Member的路由处理函数
func New() *HandlerMember {
	return &HandlerMember{
		cacheClient: cache.New(),
	}
}

// MemberLogin 登录
func (*HandlerMember) MemberLogin(ctx *gin.Context) {
	mono.OkWithData(map[string]any{"name": "张三"}, ctx)
}

// GetSms 根据手机号码,发送验证码短信
func (this *HandlerMember) GetSms(ctx *gin.Context) {
	mobile := ctx.PostForm("mobile")
	//1. 校验参数(手机号码)
	if err := utils.VerifyMobile(mobile); err != nil {
		mono.FailWithErr(err, ctx)
		return
	}
	//2. 生成验证码(随机4-6位数字)
	code := utils.RandomMobileCode()

	//TODO 后续引入协程池
	// 从协程池中开启一个任务
	err := TaskPool.Submit(func() {
		//3. 调用第三方短信平台
		//TODO
		//4. 将验证码存储到redis中
		// 设置缓存超时时间为3秒
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		// 验证码的key
		key := "sms_key_mobile_" + mobile
		// 有效期15分钟
		err := this.cacheClient.Put(timeout, key, code, time.Minute*15)
		if err != nil {
			Logger.Errorf("验证码存入Redis错误, err: %s", err)
			return
		}
		Logger.Info("验证码存入Redis code:%s", code)
	})
	if err != nil {
		// 获取协程错误
		panic(errors.NewWithCodeMessage(err.Error(), errors.PoolTaskCreateErrCode))
	}
	mono.Ok(ctx)
}

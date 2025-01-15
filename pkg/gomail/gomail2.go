package gomail

//
//import (
//	"agricultural_vision/responsecode"
//	"fmt"
//	"github.com/jordan-wright/email"
//	"github.com/patrickmn/go-cache"
//	"math/rand"
//	"net/smtp"
//	"time"
//)
//
//var (
//	// 缓存中的验证代码将在创建后5分钟内有效，且每隔10分钟进行一次清理。
//	verificationCodeCache = cache.New(5*time.Minute, 10*time.Minute)
//)
//
//// 发送验证码并存储到缓存中
//func SendVerificationCode(to string) error {
//	// 随机生成6位数的验证码
//	code := generateVerificationCode()
//
//	// 发送验证码
//	err := sendVerificationCode(to, code)
//	if err != nil {
//		return err
//	}
//
//	// 将验证码存储到缓存中，并设置过期时间
//	verificationCodeCache.Set(to, code, cache.DefaultExpiration)
//
//	return nil
//}
//
//// 发送验证代码到指定的邮箱。
//// 参数 to: 邮件接收人的邮箱地址。
//// 参数 responsecode: 需要发送的验证代码。
//func sendVerificationCode(to string, code string) error {
//	// 创建一个新的邮件实例
//	em := email.NewEmail()
//
//	// 打印调试信息，确保From和To设置正确
//	fmt.Println("Sending email from:", "农视界 <2455494167@qq.com>")
//	fmt.Println("Sending email to:", to)
//
//	em.From = "农视界 <2455494167@qq.com>"
//	em.To = []string{to}
//	em.Subject = "Verification VerificationCode"
//	// 设置邮件的HTML内容
//	em.HTML = []byte(`
//		<h1>Verification VerificationCode</h1>
//		<p>Your verification responsecode is: <strong>` + code + `</strong></p>
//	`)
//
//	// 发送邮件(这里使用QQ进行发送邮件验证码)
//	err := em.Send("smtp.qq.com:465", smtp.PlainAuth("", "2455494167@qq.com", "rpqcsjeyqoesecbd", "smtp.qq.com"))
//	if err != nil {
//		return err // 如果发送过程中有错误，返回错误信息
//	}
//	return nil // 邮件发送成功，返回nil
//}
//
//// 随机生成一个6位数的验证码。
//func generateVerificationCode() string {
//	rand.Seed(time.Now().UnixNano())
//	code := fmt.Sprintf("%06d", rand.Intn(1000000))
//	return code
//}
//
//// 验证用户输入的验证码是否正确。
//func VerifyVerificationCode(email string, code string) error {
//	// 从缓存中获取验证码
//	cachedCode, found := verificationCodeCache.Get(email)
//	// 如果没有找到验证码或者验证码过期
//	if !found {
//		return responsecode.ErrorInvalidEmailCode
//	}
//
//	// 如果找到验证码，但验证码不匹配
//	if cachedCode != code {
//		return responsecode.ErrorInvalidEmailCode
//	}
//
//	return nil
//}

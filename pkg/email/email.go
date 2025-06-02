package email

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
)

func generateCode(n int) string {
	digits := make([]byte, n)
	for i := range digits {
		digits[i] = byte(rand.Intn(10)) + '0'
	}
	return string(digits)
}

func SendCodeEmail(address string) (code string, err error) {
	// 准备邮件内容
	code = generateCode(6)
	m := gomail.NewMessage()
	m.SetHeader("From", "254057739@qq.com")
	m.SetHeader("To", address)
	m.SetHeader("Subject", "登录验证码")
	m.SetBody("text/plain", fmt.Sprintf("您的验证码为：%s，有效期 5 分钟。", code))

	// 发送邮件信息
	d := gomail.NewDialer("smtp.qq.com", 465, "254057739@qq.com", "hhwgvkdjzpxabggg")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 如遇证书问题可加
	if err := d.DialAndSend(m); err != nil {
		return "", fmt.Errorf("邮件发送失败: %w", err)
	}
	return
}

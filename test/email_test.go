package test

//func (svc *CodeService) SendEmail(ctx context.Context, sendTo string, subject string, body string) error {
//	from := svc.EmailInfo.from
//	password := svc.EmailInfo.password // 邮箱授权码
//	smtpServer := svc.EmailInfo.smtpServer
//
//	// 设置 PlainAuth
//	auth := smtp.PlainAuth("", from, password, "smtp.qq.com")
//
//	// 创建 tls 配置
//	tlsconfig := &tls.Config{
//		InsecureSkipVerify: true,
//		ServerName:         "smtp.qq.com",
//	}
//
//	// 连接到 SMTP 服务器
//	conn, err := tls.Dial("tcp", smtpServer, tlsconfig)
//	if err != nil {
//		return fmt.Errorf("TLS 连接失败: %v", err)
//	}
//	defer conn.Close()
//
//	client, err := smtp.NewClient(conn, "smtp.qq.com")
//	if err != nil {
//		return fmt.Errorf("SMTP 客户端创建失败: %v", err)
//	}
//	defer client.Quit()
//
//	// 使用 auth 进行认证
//	if err = client.Auth(auth); err != nil {
//		return fmt.Errorf("认证失败: %v", err)
//	}
//
//	// 设置发件人和收件人
//	if err = client.Mail(from); err != nil {
//		return fmt.Errorf("发件人设置失败: %v", err)
//	}
//	if err = client.Rcpt(sendTo); err != nil {
//		return fmt.Errorf("收件人设置失败: %v", err)
//	}
//
//	// 写入邮件内容
//	wc, err := client.Data()
//	if err != nil {
//		return fmt.Errorf("数据写入失败: %v", err)
//	}
//	defer wc.Close()
//
//	msg := []byte("From: gstecp <" + from + ">\r\n" +
//		"To: " + sendTo + "\r\n" +
//		"Subject: " + subject + "\r\n" +
//		"\r\n" +
//		body + "\r\n")
//	_, err = wc.Write(msg)
//	if err != nil {
//		return fmt.Errorf("消息发送失败: %v", err)
//	}
//
//	return nil
//}

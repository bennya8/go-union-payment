package str

import (
	"math/rand"
	"net/url"
)

// 产生随机字符串，不长于32位
func GetNonce(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}



// 生成测试使用的二维码
func QrImg(text string) string {
	newText := url.QueryEscape(text)
	return "https://cli.im/api/qrcode/code?text=" + newText + "&mhid=tBHED17pz8shMHcmKdxROK4"
}

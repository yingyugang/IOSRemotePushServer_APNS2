package main

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
)
/*
 * apnsでプッシューを送る、テストずみ。
 */
func main() {
	//certPath := flag.String("cert", "../cert.p12", "")
	//token := flag.String("token", "4648c0188ea303345012df2dee723b509dc2d1dab470cb0adb3515914156ef95", "")
	//topic := flag.String("topic", "com.moba1", "")
	//flag.Parse()
	//if *certPath == "" || *token == "" || *topic == "" {
	//	flag.PrintDefaults()
	//	os.Exit(1)
	//}
	//用工程下的APNsDev.p12
	cert, err := certificate.FromP12File("C:/Users/yingy/Documents/IOSRemotePushServer_APNS2/APNsDev.p12", "1fd94d19bbe3c")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}
	notification := &apns2.Notification{}
	//在XCodeLog里面找到DeviceToken在这里替换。每次重新安装应用的时候会产生新的Token。
	notification.DeviceToken = "a467c01f57345848a5167e8d0f55b871064dcb5f73b81c82734c71ec18340b1b"
	notification.Topic = "com.moba"
	notification.Payload = []byte(`{
			"aps" : {
				"alert":"Hello!",
        		"badge":1,
        		"mutable-content":1
			},
			"image-url": "https://yingyugang.s3-ap-northeast-1.amazonaws.com/goldIco.png"
		}`)

	client := apns2.NewClient(cert).Development()//   Production()
	res, err := client.Push(notification)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}

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
	//用工程下的APNsDev.p12
	cert, err := certificate.FromP12File("C:/Users/yingy/Documents/IOSRemotePushServer_APNS2/APNsDev.p12", "1fd94d19bbe3c")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}
	notification := &apns2.Notification{}
	//在XCodeLog里面找到DeviceToken在这里替换。每次重新安装应用的时候会产生新的Token。
	notification.DeviceToken = "79ecd5681860feadc8579353b1fa4d98f79de7defa9bed371d6e37571ba4ba80"
	notification.Topic = "com.moba"
	notification.Payload = []byte(`{
			"aps" : {
				"alert":"Hello123ddd444!",
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

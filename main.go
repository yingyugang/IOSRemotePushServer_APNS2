package main

import (
	"flag"
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
	"os"
)
/*
	apnsでプッシューを送る、テストずみ。
 */
func main() {
	certPath := flag.String("cert", "../cert.p12", "")
	token := flag.String("token", "11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7", "")
	topic := flag.String("topic", "com.moba1", "")
	flag.Parse()
	if *certPath == "" || *token == "" || *topic == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	//用工程下的APNsDev.p12
	cert, err := certificate.FromP12File("C:/Users/yingy/Documents/IOSRemotePushServer_APNS2/APNsDev.p12", "1fd94d19bbe3c")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}
	notification := &apns2.Notification{}
	//在XCodeLog里面找到DeviceToken在这里替换。每次重新安装应用的时候会产生新的Token。
	notification.DeviceToken = "0dd11ed4d1219b2e1f7a65f56a0e22bbb40a2f292d72eba4a2dd317c5f639950"
	notification.Topic = "com.moba"
	notification.Payload = []byte(`{
			"aps" : {
				"alert" : "Hello!"
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

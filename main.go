package main

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
	"net/http"
)

/*
 * apnsでプッシューを送る、テストずみ。
 */
func main() {
	http.HandleFunc("/getDeviceToken", getDeviceToken)
	http.HandleFunc("/sendSampleAPNs", sendSampleAPNs)
	http.HandleFunc("/getNotificationService", getNotificationService)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var deviceToken string

func getDeviceToken(w http.ResponseWriter, r *http.Request){
	deviceToken = r.URL.Query()["deviceToken"][0]
	fmt.Printf("%v\n", deviceToken)
}

func getNotificationService(w http.ResponseWriter, r *http.Request){
	var msg = r.URL.Query()["msg"][0]
	fmt.Printf("%v\n", msg)
}

func sendSampleAPNs(w http.ResponseWriter, r *http.Request){
	//用工程下的APNsDev.p12
	//cert, err := certificate.FromP12File("C:/Users/yingy/Documents/IOSRemotePushServer_APNS2/APNsDev.p12", "1fd94d19bbe3c")
	cert, err := certificate.FromP12File("C:/Users/yingy/Documents/IOSRemotePushServer_APNS2/AdsCLDev2.p12", "Yingyugang2017")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}
	notification := &apns2.Notification{}
	//在XCodeLog里面找到DeviceToken在这里替换。每次重新安装应用的时候会产生新的Token。
	notification.DeviceToken = "41c41626e18e66cc69ae2f5c4c6f18189580a4d2491b013ec41d65de4bf4d218"
	if len(deviceToken) > 0 {
		notification.DeviceToken = deviceToken
	}
	//notification.Topic = "com.moba"
	notification.Topic = "com.platinum-egg.crosslink-dev2"
	notification.Payload = []byte(`{
			"aps" : {
				"alert":"Hello123ddd444!Hello123ddd444Hello123ddd444Hello123ddd444Hello123ddd444Hello123ddd444Hello123ddd444Hello123ddd444Hello123ddd444",
        		"badge":1,
        		"mutable-content":1
			},
			"image-url": "https://yingyugang.s3-ap-northeast-1.amazonaws.com/Special.wav"
		}`)
	//https://yingyugang.s3-ap-northeast-1.amazonaws.com/TID_ARTILLERY.png
	//"https://yingyugang.s3-ap-northeast-1.amazonaws.com/Special.wav"
	//"https://yingyugang.s3-ap-northeast-1.amazonaws.com/goldIco.png"

	//测试手机是开发环境（XCode上Signing&Capabilities - Signing - SigningCertificate ）
	//如果是Develop，这里要用Development。如果是Distribution，这里要用Production
	//client := apns2.NewClient(cert).Development()//   Production()
	client := apns2.NewClient(cert).Production()//   Production()

	res, err := client.Push(notification)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
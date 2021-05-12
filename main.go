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
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var deviceToken string

func getDeviceToken(w http.ResponseWriter, r *http.Request){
	deviceToken = r.URL.Query()["deviceToken"][0]
}

func sendSampleAPNs(w http.ResponseWriter, r *http.Request){
	//用工程下的APNsDev.p12
	cert, err := certificate.FromP12File("C:/Users/yingy/Documents/IOSRemotePushServer_APNS2/APNsDev.p12", "1fd94d19bbe3c")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}
	notification := &apns2.Notification{}
	//在XCodeLog里面找到DeviceToken在这里替换。每次重新安装应用的时候会产生新的Token。
	notification.DeviceToken = "5b4e466af6171c9991e04f8934ed7247702c02e70c3dd59dc741b36ec006189d"
	if len(deviceToken) > 0 {
		notification.DeviceToken = deviceToken
	}
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
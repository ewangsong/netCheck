//检测http访问
package model

import (
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	//Timeout: time.Second * 5,
}

func GetHttpInfo() {
	for _, url := range CheckHttp {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("user-agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 Edg/87.0.664.75`)
		//访问开始时间

		start := time.Now()
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println(err)
			continue

		}
		resp.Body.Close()
		//访问结束时间
		end := time.Now()
		status := resp.Status
		subtime := end.Sub(start).String()
		fmt.Printf("%s 状态码%s 花费时间%s\n", url, status, subtime)

	}
	//	defer resp.Body.Close()
	// b, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("get resp failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Println(string(b))

}

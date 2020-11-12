package sourcecode

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// getData http get请求
func getData() {
	client := &http.Client{}
	resp, err := client.Get("")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
}

// http.Post
func httpPost() {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

// http.postForm
func postData() {
	//client := &http.Client{}
	resp, err := http.PostForm("https://www.pgyer.com/apiv2/app/view", url.Values{"appKey": {"62c99290f0cb2c567cb153c1fba75d867e"},
		"_api_key": {"584f29517115df2034348b0c06b3dc57"}, "buildKey": {"22d4944d06354c8dcfb16c4285d04e41"}})
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(body))

}

// 对于比较复杂的http请求，我们可以用到http.do的方式进行请求
func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

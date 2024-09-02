package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"

	"github.com/CocaineCong/gin-mall/pkg/utils/ctl"
	"github.com/CocaineCong/gin-mall/pkg/utils/log"
	"github.com/CocaineCong/gin-mall/types"
)

var ProtectSrvIns *ProtectSrv
var ProtectSrvOnce sync.Once

type ProtectSrv struct {
}

func GetProtectSrv() *ProtectSrv {
	ProtectSrvOnce.Do(func() {
		ProtectSrvIns = &ProtectSrv{}
	})
	return ProtectSrvIns
}

// http://10.236.163.233:8080
// 创建商品
func (s *ProtectSrv) ProtectCreate(ctx context.Context, req *types.ProtectCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	uId := u.Id
	fmt.Println(uId)

	data := make(map[string]interface{})
	data["uid"] = req.Uid
	fmt.Printf("%#v \n", data)

	response, err := PostRequest("http://10.236.163.233:8080/v1/config/"+req.ID+"/copy", data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 定义一个 map 变量来存储转换后的结果
	var result map[string]interface{}
	err = json.Unmarshal([]byte(response), &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	return result["data"], nil

	// res1, err := GetRequest("http://10.236.163.233:8080/v1/config/lists?page=1&size=10&uid=" + req.Uid)
	// if err != nil {
	// 	return nil, err
	// }
	// // 定义一个 map 变量来存储转换后的结果
	// var result map[string]interface{}
	// err = json.Unmarshal([]byte(res1), &result)
	// if err != nil {
	// 	fmt.Println("Error unmarshalling JSON:", err)
	// 	return
	// }
	// return result["data"], nil
}

func (s *ProtectSrv) ConfigAdd(uid, remark, startTime, endTime string) (string, error) {
	data := map[string]interface{}{
		"uid":       uid,
		"remark":    remark,
		"startTime": startTime,
		"endTime":   endTime,
	}
	return PostRequest("/v1/config/", data)
}

func (s *ProtectSrv) ConfigEdit(id int, uid, remark, startTime, endTime string) (string, error) {
	data := map[string]interface{}{
		"uid":       uid,
		"remark":    remark,
		"startTime": startTime,
		"endTime":   endTime,
	}
	return PostRequest(fmt.Sprintf("/v1/config/%d/edit", id), data)
}

func (s *ProtectSrv) ConfigCopy(id int, uid string) (string, error) {
	data := map[string]interface{}{
		"uid": uid,
	}
	return PostRequest(fmt.Sprintf("/v1/config/%d/copy", id), data)
}

func (s *ProtectSrv) ConfigDelete(id int, uid string) (string, error) {
	data := map[string]interface{}{
		"uid": uid,
	}
	return PostRequest(fmt.Sprintf("/v1/config/%d/delete", id), data)
}

func (s *ProtectSrv) ConfigLists(uid string, page, size int) (string, error) {
	data := map[string]interface{}{
		"uid":  uid,
		"page": page,
		"size": size,
	}
	// 构建查询参数字符串
	queryParams := url.Values{}
	for key, value := range data {
		queryParams.Add(key, fmt.Sprintf("%v", value))
	}
	fullURL := "/v1/config/lists?" + queryParams.Encode()

	return GetRequest(fullURL)
}

func GetRequest(url string) (string, error) {
	// 创建新的 GET 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// 设置请求头
	req.Header.Set("key", "TIANYUWEB")

	// 使用 http.Client 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// 返回响应体内容
	return string(body), nil
}

func PostRequest(url string, data map[string]interface{}) (string, error) {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// 遍历 map 并写入 multipart/form-data 部分
	for key, value := range data {
		part, err := writer.CreateFormField(key)
		if err != nil {
			return "", fmt.Errorf("failed to create form field: %w", err)
		}
		_, err = part.Write([]byte(fmt.Sprintf("%v", value)))
		if err != nil {
			return "", fmt.Errorf("failed to write form field value: %w", err)
		}
	}

	// 关闭 writer，以便写入结尾 boundary
	writer.Close()

	// 创建新的 POST 请求
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// 设置请求头
	req.Header.Add("key", "TIANYUWEB")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("User-Agent", "PostmanRuntime-ApipostRuntime/1.1.0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// 发送请求
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer res.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

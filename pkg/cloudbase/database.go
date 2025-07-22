package cloudbase

import "fmt"

// 参考腾讯云开发 OPENAPI 文档 https://docs.cloudbase.net/http-api/model/%E6%95%B0%E6%8D%AE%E6%A8%A1%E5%9E%8B-openapi

type Record map[string]interface{}

// GetModelRecordListRes
type GetModelListRes struct {
	Data struct {
		Records []Record `json:"records"`
		Total   int      `json:"total"`
	} `json:"data"`
}

// GetModelRecordList
func (c *Client) GetModelRecordList(modelName string) (*GetModelListRes, error) {
	url := fmt.Sprintf("https://%s.%s/v1/model/prod/%s/list", c.EnvId, domain, modelName)

	token, err := c.SignIn()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Authorization"] = token
	res := &GetModelListRes{}
	err = httpGet(url, headers, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

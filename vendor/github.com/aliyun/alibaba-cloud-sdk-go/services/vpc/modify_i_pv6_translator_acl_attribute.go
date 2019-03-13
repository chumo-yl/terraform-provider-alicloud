package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyIPv6TranslatorAclAttribute invokes the vpc.ModifyIPv6TranslatorAclAttribute API synchronously
// api document: https://help.aliyun.com/api/vpc/modifyipv6translatoraclattribute.html
func (client *Client) ModifyIPv6TranslatorAclAttribute(request *ModifyIPv6TranslatorAclAttributeRequest) (response *ModifyIPv6TranslatorAclAttributeResponse, err error) {
	response = CreateModifyIPv6TranslatorAclAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyIPv6TranslatorAclAttributeWithChan invokes the vpc.ModifyIPv6TranslatorAclAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyipv6translatoraclattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIPv6TranslatorAclAttributeWithChan(request *ModifyIPv6TranslatorAclAttributeRequest) (<-chan *ModifyIPv6TranslatorAclAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyIPv6TranslatorAclAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyIPv6TranslatorAclAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyIPv6TranslatorAclAttributeWithCallback invokes the vpc.ModifyIPv6TranslatorAclAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyipv6translatoraclattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIPv6TranslatorAclAttributeWithCallback(request *ModifyIPv6TranslatorAclAttributeRequest, callback func(response *ModifyIPv6TranslatorAclAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyIPv6TranslatorAclAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyIPv6TranslatorAclAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyIPv6TranslatorAclAttributeRequest is the request struct for api ModifyIPv6TranslatorAclAttribute
type ModifyIPv6TranslatorAclAttributeRequest struct {
	*requests.RpcRequest
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AclName              string           `position:"Query" name:"AclName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyIPv6TranslatorAclAttributeResponse is the response struct for api ModifyIPv6TranslatorAclAttribute
type ModifyIPv6TranslatorAclAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyIPv6TranslatorAclAttributeRequest creates a request to invoke ModifyIPv6TranslatorAclAttribute API
func CreateModifyIPv6TranslatorAclAttributeRequest() (request *ModifyIPv6TranslatorAclAttributeRequest) {
	request = &ModifyIPv6TranslatorAclAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyIPv6TranslatorAclAttribute", "vpc", "openAPI")
	return
}

// CreateModifyIPv6TranslatorAclAttributeResponse creates a response to parse from ModifyIPv6TranslatorAclAttribute response
func CreateModifyIPv6TranslatorAclAttributeResponse() (response *ModifyIPv6TranslatorAclAttributeResponse) {
	response = &ModifyIPv6TranslatorAclAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

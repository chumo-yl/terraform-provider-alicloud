package cr

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

// DeleteRepoWebhook invokes the cr.DeleteRepoWebhook API synchronously
// api document: https://help.aliyun.com/api/cr/deleterepowebhook.html
func (client *Client) DeleteRepoWebhook(request *DeleteRepoWebhookRequest) (response *DeleteRepoWebhookResponse, err error) {
	response = CreateDeleteRepoWebhookResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRepoWebhookWithChan invokes the cr.DeleteRepoWebhook API asynchronously
// api document: https://help.aliyun.com/api/cr/deleterepowebhook.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRepoWebhookWithChan(request *DeleteRepoWebhookRequest) (<-chan *DeleteRepoWebhookResponse, <-chan error) {
	responseChan := make(chan *DeleteRepoWebhookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRepoWebhook(request)
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

// DeleteRepoWebhookWithCallback invokes the cr.DeleteRepoWebhook API asynchronously
// api document: https://help.aliyun.com/api/cr/deleterepowebhook.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRepoWebhookWithCallback(request *DeleteRepoWebhookRequest, callback func(response *DeleteRepoWebhookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRepoWebhookResponse
		var err error
		defer close(result)
		response, err = client.DeleteRepoWebhook(request)
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

// DeleteRepoWebhookRequest is the request struct for api DeleteRepoWebhook
type DeleteRepoWebhookRequest struct {
	*requests.RoaRequest
	RepoNamespace string           `position:"Path" name:"RepoNamespace"`
	WebhookId     requests.Integer `position:"Path" name:"WebhookId"`
	RepoName      string           `position:"Path" name:"RepoName"`
}

// DeleteRepoWebhookResponse is the response struct for api DeleteRepoWebhook
type DeleteRepoWebhookResponse struct {
	*responses.BaseResponse
}

// CreateDeleteRepoWebhookRequest creates a request to invoke DeleteRepoWebhook API
func CreateDeleteRepoWebhookRequest() (request *DeleteRepoWebhookRequest) {
	request = &DeleteRepoWebhookRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "DeleteRepoWebhook", "/repos/[RepoNamespace]/[RepoName]/webhooks/[WebhookId]", "acr", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteRepoWebhookResponse creates a response to parse from DeleteRepoWebhook response
func CreateDeleteRepoWebhookResponse() (response *DeleteRepoWebhookResponse) {
	response = &DeleteRepoWebhookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
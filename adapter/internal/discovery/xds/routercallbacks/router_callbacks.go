/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package routercallbacks

import (
	"context"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/wso2/apk/adapter/internal/discovery/xds/common"
	logger "github.com/wso2/apk/adapter/internal/loggers"
	logging "github.com/wso2/apk/adapter/internal/logging"
)

var nodeQueueInstance *common.NodeQueue

func init() {
	nodeQueueInstance = common.GenerateNodeQueue()
}

// Callbacks is used to debug the xds server related communication.
type Callbacks struct {
}

// Report logs the fetches and requests.
func (cb *Callbacks) Report() {}

// OnStreamOpen prints debug logs
func (cb *Callbacks) OnStreamOpen(_ context.Context, id int64, typ string) error {
	logger.LoggerRouterXdsCallbacks.Debugf("stream %d open for %s\n", id, typ)
	return nil
}

// OnStreamClosed prints debug logs
func (cb *Callbacks) OnStreamClosed(id int64, node *core.Node) {
	logger.LoggerRouterXdsCallbacks.Debugf("stream %d closed\n", id)
}

// OnStreamRequest prints debug logs
func (cb *Callbacks) OnStreamRequest(id int64, request *discovery.DiscoveryRequest) error {
	nodeIdentifier := common.GetNodeIdentifier(request)
	if nodeQueueInstance.IsNewNode(nodeIdentifier) {
		logger.LoggerRouterXdsCallbacks.Infof("stream request on stream id: %d, from node: %s, version: %s",
			id, nodeIdentifier, request.VersionInfo)
	}
	logger.LoggerRouterXdsCallbacks.Debugf("stream request on stream id: %d, from node: %s, version: %s, for type: %s",
		id, nodeIdentifier, request.VersionInfo, request.TypeUrl)
	if request.ErrorDetail != nil {
		logger.LoggerRouterXdsCallbacks.ErrorC(logging.GetErrorByCode(logging.Error1401, request.GetTypeUrl(),
			id, nodeIdentifier, request.ErrorDetail.Message))
	}
	return nil
}

// OnStreamResponse prints debug logs
func (cb *Callbacks) OnStreamResponse(context context.Context, id int64, request *discovery.DiscoveryRequest,
	response *discovery.DiscoveryResponse) {
	nodeIdentifier := common.GetNodeIdentifier(request)
	logger.LoggerRouterXdsCallbacks.Debugf("stream response on stream id: %d, to node: %s, version: %s, for type: %v", id,
		nodeIdentifier, response.VersionInfo, response.TypeUrl)
}

// OnFetchRequest prints debug logs
func (cb *Callbacks) OnFetchRequest(_ context.Context, req *discovery.DiscoveryRequest) error {
	logger.LoggerRouterXdsCallbacks.Debugf("fetch request from node %s, version: %s, for type %s", common.GetNodeIdentifier(req),
		req.VersionInfo, req.TypeUrl)
	return nil
}

// OnFetchResponse prints debug logs
func (cb *Callbacks) OnFetchResponse(req *discovery.DiscoveryRequest, res *discovery.DiscoveryResponse) {
	logger.LoggerRouterXdsCallbacks.Debugf("fetch response to node: %s, version: %s, for type %s", common.GetNodeIdentifier(req),
		req.VersionInfo, res.TypeUrl)
}

// OnDeltaStreamOpen is unused.
func (cb *Callbacks) OnDeltaStreamOpen(_ context.Context, id int64, typ string) error {
	return nil
}

// OnDeltaStreamClosed is unused.
func (cb *Callbacks) OnDeltaStreamClosed(id int64, node *core.Node) {
}

// OnStreamDeltaResponse is unused.
func (cb *Callbacks) OnStreamDeltaResponse(id int64, req *discovery.DeltaDiscoveryRequest, res *discovery.DeltaDiscoveryResponse) {
}

// OnStreamDeltaRequest is unused.
func (cb *Callbacks) OnStreamDeltaRequest(id int64, req *discovery.DeltaDiscoveryRequest) error {
	return nil
}

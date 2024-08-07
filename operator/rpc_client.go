package operator

import (
	"fmt"
	"net/rpc"
	"time"

	"anzen-avs/aggregator"
	"anzen-avs/metrics"

	"github.com/Layr-Labs/eigensdk-go/logging"
)

type AggregatorRpcClienter interface {
	SendSignedOraclePullTaskReponseToAggregator(signedOraclePullTaskResponse *aggregator.SignedOraclePullTaskResponse)
}

type AggregatorRpcClient struct {
	rpcClient            *rpc.Client
	metrics              metrics.Metrics
	logger               logging.Logger
	aggregatorIpPortAddr string
}

func NewAggregatorRpcClient(aggregatorIpPortAddr string, logger logging.Logger, metrics metrics.Metrics) (*AggregatorRpcClient, error) {
	return &AggregatorRpcClient{
		// set to nil so that we can create an rpc client even if the aggregator is not running
		rpcClient:            nil,
		metrics:              metrics,
		logger:               logger,
		aggregatorIpPortAddr: aggregatorIpPortAddr,
	}, nil
}

func (c *AggregatorRpcClient) dialAggregatorRpcClient() error {
	client, err := rpc.DialHTTP("tcp", c.aggregatorIpPortAddr)
	if err != nil {
		return err
	}
	c.rpcClient = client
	return nil
}

// SendSignedOraclePullTaskReponseToAggregator sends a signed oracle pull task response to the aggregator.
// it is meant to be ran inside a go thread, so doesn't return anything.
// this is because sending the signed oracle pull task response to the aggregator is time sensitive,
// so there is no point in retrying if it fails for a few times.
// Currently hardcoded to retry sending the signed oracle pull task response 5 times, waiting 2 seconds in between each attempt.
func (c *AggregatorRpcClient) SendSignedOraclePullTaskReponseToAggregator(signedOraclePullTaskResponse *aggregator.SignedOraclePullTaskResponse) {
	if c.rpcClient == nil {
		c.logger.Info("rpc client is nil. Dialing aggregator rpc client")
		err := c.dialAggregatorRpcClient()
		if err != nil {
			c.logger.Error("Could not dial aggregator rpc client. Not sending signed oracle pull task response header to aggregator. Is aggregator running?", "err", err)
			return
		}
	}
	// we don't check this bool. It's just needed because rpc.Call requires rpc methods to have a return value
	var reply bool
	// We try to send the response 5 times to the aggregator, waiting 2 times in between each attempt.
	// This is mostly only necessary for local testing, since the aggregator sometimes is not ready to process task responses
	// before the operator gets the new task created log from anvil (because blocks are mined instantly)
	// the aggregator needs to read some onchain data related to quorums before it can accept operator signed task responses.
	c.logger.Info("Sending signed oracle pull task response header to aggregator", "signedOraclePullTaskResponse", fmt.Sprintf("%#v", signedOraclePullTaskResponse))
	for i := 0; i < 5; i++ {
		err := c.rpcClient.Call("Aggregator.ProcessSignedOraclePullTaskResponse", signedOraclePullTaskResponse, &reply)
		if err != nil {
			c.logger.Info("Received error from aggregator", "err", err)
		} else {
			c.logger.Info("Signed oracle pull task response header accepted by aggregator.", "reply", reply)
			c.metrics.IncNumTasksAcceptedByAggregator()
			return
		}
		c.logger.Infof("Retrying in 2 seconds")
		time.Sleep(2 * time.Second)
	}
	c.logger.Errorf("Could not send signed oracle pull task response to aggregator. Tried 5 times.")
}

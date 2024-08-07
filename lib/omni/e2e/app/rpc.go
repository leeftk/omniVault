package app

import (
	"context"
	"fmt"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	rpctypes "github.com/cometbft/cometbft/rpc/core/types"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/types"
)

// waitForHeight waits for the network to reach a certain height (or above),
// returning the highest height seen. Errors if the network is not making
// progress at all.
func waitForHeight(ctx context.Context, testnet *e2e.Testnet, height int64) (*types.Block, *types.BlockID, error) {
	var (
		err          error
		maxResult    *rpctypes.ResultBlock
		clients      = map[string]*rpchttp.HTTP{}
		lastIncrease = time.Now()
	)

	currentBlock := func(ctx context.Context, client *rpchttp.HTTP) (*rpctypes.ResultBlock, error) {
		ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()
		resp, err := client.Block(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, "get block")
		}

		return resp, nil
	}

	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	logTicker := time.NewTicker(5 * time.Second)
	defer logTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, nil, errors.Wrap(ctx.Err(), "context canceled")
		case <-logTicker.C:
			current := "unknown"
			if maxResult != nil {
				current = fmt.Sprint(maxResult.Block.Height)
			}
			log.Debug(ctx, "Still waiting for height", "current_height", current, "wait_for_height", height)
		case <-queryTicker.C:
			for _, node := range testnet.Nodes {
				if node.Stateless() {
					continue
				}
				client, ok := clients[node.Name]
				if !ok {
					client, err = node.Client()
					if err != nil {
						continue
					}
					clients[node.Name] = client
				}

				result, err := currentBlock(ctx, client)
				if errors.Is(err, context.DeadlineExceeded) {
					return nil, nil, errors.Wrap(err, "timeout")
				} else if err != nil {
					continue
				}

				if result.Block != nil && (maxResult == nil || result.Block.Height > maxResult.Block.Height) {
					maxResult = result
					lastIncrease = time.Now()
				}
				if maxResult != nil && maxResult.Block.Height >= height {
					return maxResult.Block, &maxResult.BlockID, nil
				}
			}

			if len(clients) == 0 {
				return nil, nil, errors.New("unable to connect to any network nodes")
			}
			if time.Since(lastIncrease) >= 20*time.Second {
				if maxResult == nil {
					return nil, nil, errors.New("chain stalled at unknown height")
				}

				return nil, nil, errors.New("chain stalled", "height", maxResult.Block.Height)
			}
		}
	}
}

// waitForNode waits for a node to become available and catch up to the given block height.
func waitForNode(ctx context.Context, node *e2e.Node, height int64, timeout time.Duration,
) (*rpctypes.ResultStatus, error) {
	client, err := node.Client()
	if err != nil {
		return nil, errors.Wrap(err, "getting client")
	}

	timer := time.NewTimer(0)
	defer timer.Stop()

	var lastHeight int64
	var lastErr error
	lastChanged := time.Now()

	for {
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "context canceled")
		case <-timer.C:
			if time.Since(lastChanged) > timeout {
				return nil, errors.New("timed out waiting for height", "name", node.Name, "height", height, "last_err", lastErr, "last_height", lastHeight)
			}

			status, err := client.Status(ctx)
			switch {
			case err != nil:
				lastErr = err
			case status.SyncInfo.LatestBlockHeight >= height && (height == 0 || !status.SyncInfo.CatchingUp):
				return status, nil
			case lastHeight < status.SyncInfo.LatestBlockHeight:
				lastHeight = status.SyncInfo.LatestBlockHeight
				lastChanged = time.Now()
			}

			timer.Reset(300 * time.Millisecond)
		}
	}
}

// waitForAllNodes waits for all nodes to become available and catch up to the given block height.
func waitForAllNodes(ctx context.Context, testnet *e2e.Testnet, height int64, timeout time.Duration) (int64, error) {
	var lastHeight int64

	deadline := time.Now().Add(timeout)

	for _, node := range testnet.Nodes {
		if node.Mode == e2e.ModeSeed {
			continue
		}

		status, err := waitForNode(ctx, node, height, time.Until(deadline))
		if err != nil {
			return 0, err
		}

		if status.SyncInfo.LatestBlockHeight > lastHeight {
			lastHeight = status.SyncInfo.LatestBlockHeight
		}
	}

	return lastHeight, nil
}

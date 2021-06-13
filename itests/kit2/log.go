package kit2

import (
	"github.com/filecoin-project/lotus/lib/lotuslog"
	logging "github.com/ipfs/go-log/v2"
)

func QuietMiningLogs() {
	lotuslog.SetupLogLevels()

	_ = logging.SetLogLevel("miner", "ERROR")
	_ = logging.SetLogLevel("chainstore", "ERROR")
	_ = logging.SetLogLevel("chain", "ERROR")
	_ = logging.SetLogLevel("sub", "ERROR")
	_ = logging.SetLogLevel("storageminer", "ERROR")
	_ = logging.SetLogLevel("pubsub", "ERROR")
	_ = logging.SetLogLevel("gen", "ERROR")
	_ = logging.SetLogLevel("dht/RtRefreshManager", "ERROR")
}
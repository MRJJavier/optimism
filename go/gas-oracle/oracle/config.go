package oracle

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum-optimism/optimism/go/gas-oracle/flags"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
)

// Config represents the configuration options for the gas oracle
type Config struct {
	chainID                      *big.Int
	ethereumHttpUrl              string
	gasPriceOracleAddress        common.Address
	privateKey                   *ecdsa.PrivateKey
	gasPrice                     *big.Int
	waitForReceipt               bool
	floorPrice                   float64
	targetGasPerSecond           float64
	maxPercentChangePerEpoch     float64
	averageBlockGasLimitPerEpoch float64
	epochLengthSeconds           uint64
	significantFactor            float64
}

// NewConfig creates a new Config
func NewConfig(ctx *cli.Context) *Config {
	cfg := Config{}
	cfg.ethereumHttpUrl = ctx.GlobalString(flags.EthereumHttpUrlFlag.Name)
	addr := ctx.GlobalString(flags.GasPriceOracleAddressFlag.Name)
	cfg.gasPriceOracleAddress = common.HexToAddress(addr)
	cfg.targetGasPerSecond = ctx.GlobalFloat64(flags.TargetGasPerSecondFlag.Name)
	cfg.maxPercentChangePerEpoch = ctx.GlobalFloat64(flags.MaxPercentChangePerEpochFlag.Name)
	cfg.averageBlockGasLimitPerEpoch = ctx.GlobalFloat64(flags.AverageBlockGasLimitPerEpochFlag.Name)
	cfg.epochLengthSeconds = ctx.GlobalUint64(flags.EpochLengthSecondsFlag.Name)
	cfg.significantFactor = ctx.GlobalFloat64(flags.SignificantFactorFlag.Name)
	cfg.floorPrice = ctx.GlobalFloat64(flags.FloorPriceFlag.Name)

	if ctx.GlobalIsSet(flags.PrivateKeyFlag.Name) {
		hex := ctx.GlobalString(flags.PrivateKeyFlag.Name)
		hex = strings.TrimPrefix(hex, "0x")
		key, err := crypto.HexToECDSA(hex)
		if err != nil {
			log.Error(fmt.Sprintf("Option %q: %v", flags.PrivateKeyFlag.Name, err))
		}
		cfg.privateKey = key
	} else {
		log.Crit("No private key configured")
	}

	if ctx.GlobalIsSet(flags.ChainIDFlag.Name) {
		chainID := ctx.GlobalUint64(flags.ChainIDFlag.Name)
		cfg.chainID = new(big.Int).SetUint64(chainID)
	}

	if ctx.GlobalIsSet(flags.TransactionGasPriceFlag.Name) {
		gasPrice := ctx.GlobalUint64(flags.TransactionGasPriceFlag.Name)
		cfg.gasPrice = new(big.Int).SetUint64(gasPrice)
	}

	if ctx.GlobalIsSet(flags.WaitForReceiptFlag.Name) {
		cfg.waitForReceipt = true
	}

	return &cfg
}

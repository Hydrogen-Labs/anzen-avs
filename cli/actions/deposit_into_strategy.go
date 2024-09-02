package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"anzen-avs/core/config"
	"anzen-avs/operator"
	"anzen-avs/types"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

func DepositIntoStrategy(ctx *cli.Context) error {

	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}
	// need to make sure we don't register the operator on startup
	// when using the cli commands to register the operator.
	nodeConfig.RegisterOperatorOnStartup = false
	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		return err
	}

	strategyAddrStr := ctx.String("strategy-addr")
	strategyAddr := common.HexToAddress(strategyAddrStr)
	amountStr := ctx.String("amount")
	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		fmt.Println("Error converting amount to big.Int")
		return err
	}

	if amount.Cmp(big.NewInt(0)) == -1 {
		fmt.Println("Deposit amount must be greater than 0")
		return err
	}

	err = operator.DepositIntoStrategy(strategyAddr, amount)
	if err != nil {
		return err
	}

	return nil
}

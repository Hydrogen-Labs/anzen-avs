package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/urfave/cli"

	"anzen-avs/core/config"
	"anzen-avs/operator"
	"anzen-avs/types"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{config.ConfigFileFlag, config.AnzenDeploymentFileFlag}
	app.Name = "anzen-operator"
	app.Usage = "Anzen Operator"
	app.Description = "Service that reads AVS economic secuirty metrics posted onchain, verifies, signs, and sends them to the aggregator."

	app.Action = operatorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed. Message:", err)
	}
}

func operatorMain(ctx *cli.Context) error {

	log.Println("Initializing Operator")
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}

	var anzenDeploymentRaw config.AnzenDeploymentRaw
	anzenDeploymentFilePath := ctx.GlobalString(config.AnzenDeploymentFileFlag.Name)
	if _, err := os.Stat(anzenDeploymentFilePath); errors.Is(err, os.ErrNotExist) {
		panic("Path " + anzenDeploymentFilePath + " does not exist")
	}
	sdkutils.ReadJsonConfig(anzenDeploymentFilePath, &anzenDeploymentRaw)

	nodeConfig.AVSRegistryCoordinatorAddress = anzenDeploymentRaw.Addresses.RegistryCoordinatorAddr
	nodeConfig.OperatorStateRetrieverAddress = anzenDeploymentRaw.Addresses.OperatorStateRetrieverAddr
	nodeConfig.SafetyFactorOracleAddr = anzenDeploymentRaw.Addresses.SafetyFactorOracleAddr

	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	log.Println("initializing operator")
	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		return err
	}
	log.Println("initialized operator")

	log.Println("starting operator")
	err = operator.Start(context.Background())
	if err != nil {
		return err
	}
	log.Println("started operator")

	return nil

}

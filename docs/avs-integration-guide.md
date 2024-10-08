## AVS Developer Guide

### Introduction

TODO: Add a brief introduction to the Anzen AVS.

### Intregrating with Anzen Guide

#### 1. Economic Security Analysis

Anzen will perform an economic security analysis of your AVS, and create a safety factor module for your AVS. This is the module which the Anzen AVS will use to determine the economic safety needs of your AVS and reward operators accordingly.

#### 2. Safety Factor Module Creation

Based on the analysis, we will create a [module](../safety-factor/modules/) for your AVS. This module will be used to determine the safety factor of your AVS, and will be used by the Anzen AVS to determine the reward payments to operators. This what ultimately goes on-chain.

#### 3. Creating the AVSReservesManager

At the smart contract level, Anzen will create your AVS an `AVSReservesManager.sol` contract which will custody the reward payments. This contract will be upgradeable, and an addressed provided by your AVS will be set as the admin to modify etc. This is created through the [AVSReservesManagerFactory](../contracts/src/AVSReservesManagerFactory.sol) contract.

Anzen should be provided with the following information to create the AVSReservesManager:

- `proxyAdmin`: The address that will be able to upgrade the AVSReservesManager contract.
- `safetyFactorConfig`: The safety factor configuration for your AVS. For more information refer to [safetyFactorConfig](../contracts//src/static/Structs.sol).
- `avsGov`: The address of the AVS governance contract.
- `avsServiceManager`: The address of the AVS Service Manager contract.
- `initial_rewardTokens`: The initial reward token addresses for the AVS.
- `initial_tokenFlowsPerSecond`: The initial token flows per second for the AVS.

The `SafetyFactorConfig` struct is defined as follows:

```
struct SafetyFactorConfig {
    int256 TARGET_SF_LOWER_BOUND;
    int256 TARGET_SF_UPPER_BOUND;
    uint256 MAX_REDUCTION_FACTOR;
    uint256 MAX_INCREASE_FACTOR;
    uint256 minEpochDuration;
}
```

- `TARGET_SF_LOWER_BOUND`: The lower bound of the target safety factor, it must be greater than `0`.
- `TARGET_SF_UPPER_BOUND`: The upper bound of the target safety factor, it must be between `TARGET_SF_LOWER_BOUND` and `10e9`.
- `MAX_REDUCTION_FACTOR`: The maximum reward token tps reduction factor for the safety factor, it must be between `0` and `10e9`. I.e `2*10e8` means a maximum 20% reduction.
- `MAX_INCREASE_FACTOR`: The maximum reward token tps increase factor for the safety factor, it must be between `0` and `10e9`. I.e `2*10e8` means a maximum 20% increase.
- `minEpochDuration`: The minimum duration in seconds for which the safety factor can be updated, it must be greater than `0`.

#### 4. Registering the AVSReservesManager with your AVS Service Manager

The AVSReservesManager will be registered with your [AVS Service Manager](https://github.com/Layr-Labs/eigenlayer-middleware/blob/dev/src/ServiceManagerBase.sol), which will be used to make programmatic payments to operators.

````

function setRewardsInitiator( address \_rewardsInitiator ) {}

```

#### 5. Anzen Node Upgrade

After the proper integration, the Anzen operator node will be upgraded to include your AVS in the list of AVSs it can operate on and begin to post safety factor responses.
```
````

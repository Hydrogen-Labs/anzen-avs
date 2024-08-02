// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import "forge-std/console.sol";
import {
    Accumulator, SafetyFactorConfig, BPS_DENOMINATOR, PRECISION, MAX_PERFORMANCE_FEE_BPS
} from "../static/Structs.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

library AccumulatorLib {
    function init(Accumulator storage accumulator, uint256 tokensPerSecond, int256 initSafetyFactor) external {
        accumulator.tokensPerSecond = tokensPerSecond;
        accumulator.prevTokensPerSecond = tokensPerSecond;
        accumulator.lastSafetyFactor = initSafetyFactor;
    }

    function adjustEpochFlow(
        Accumulator storage accumulator,
        SafetyFactorConfig memory config,
        int256 currentSafetyFactor,
        uint256 performanceFeeBPS,
        uint256 lastEpochUpdateTimestamp
    ) external {
        _adjustClaimableTokens(accumulator, performanceFeeBPS, lastEpochUpdateTimestamp);

        if (accumulator.lastSafetyFactor == currentSafetyFactor) {
            return;
        }

        uint256 prevTokensPerSecond = accumulator.tokensPerSecond;
        uint256 newTokensPerSecond = prevTokensPerSecond;

        if (currentSafetyFactor > config.TARGET_SF_UPPER_BOUND) {
            uint256 factorDifference = uint256(currentSafetyFactor - config.TARGET_SF_UPPER_BOUND);
            uint256 reduction = (prevTokensPerSecond * factorDifference) / uint256(config.TARGET_SF_UPPER_BOUND);
            uint256 maxReduction = _maxReduction(prevTokensPerSecond, config.MAX_REDUCTION_FACTOR);
            reduction = Math.min(reduction, maxReduction);
            newTokensPerSecond = prevTokensPerSecond - reduction;
        } else if (currentSafetyFactor < config.TARGET_SF_LOWER_BOUND) {
            uint256 factorDifference = uint256(config.TARGET_SF_LOWER_BOUND - currentSafetyFactor);
            uint256 increase = (prevTokensPerSecond * factorDifference) / uint256(config.TARGET_SF_LOWER_BOUND);
            uint256 maxIncrease = _maxIncrease(prevTokensPerSecond, config.MAX_INCREASE_FACTOR);

            increase = Math.min(increase, maxIncrease);

            newTokensPerSecond = prevTokensPerSecond + increase;
        }

        accumulator.tokensPerSecond = newTokensPerSecond;
        accumulator.prevTokensPerSecond = prevTokensPerSecond;
        accumulator.lastSafetyFactor = currentSafetyFactor;
    }

    function claim(Accumulator storage accumulator, uint256 performanceFeeBPS, uint256 lastEpochUpdateTimestamp)
        external
        returns (uint256, uint256)
    {
        _adjustClaimableTokens(accumulator, performanceFeeBPS, lastEpochUpdateTimestamp);

        uint256 claimableTokens = accumulator.claimableTokens;
        uint256 claimableFees = accumulator.claimableFees;

        accumulator.claimableTokens = 0;
        accumulator.claimableFees = 0;

        return (claimableTokens, claimableFees);
    }

    function overrideTokensPerSecond(
        Accumulator storage accumulator,
        uint256 tokensPerSecond,
        uint256 lastEpochUpdateTimestamp
    ) external {
        _adjustClaimableTokens(accumulator, 0, lastEpochUpdateTimestamp);

        accumulator.tokensPerSecond = tokensPerSecond;
        accumulator.prevTokensPerSecond = tokensPerSecond;
    }

    function overrideClaimableTokens(Accumulator storage accumulator, uint256 claimableTokens) external {
        accumulator.claimableTokens = claimableTokens;
    }

    function _calculateClaimableTokensAndFee(
        Accumulator memory accumulator,
        uint256 performanceFeeBPS,
        uint256 currentTimestamp,
        uint256 lastEpochUpdateTimestamp
    ) internal pure returns (uint256 tokensGained, uint256 fee) {
        uint256 elapsedTime = currentTimestamp - lastEpochUpdateTimestamp;

        if (accumulator.prevTokensPerSecond > accumulator.tokensPerSecond) {
            uint256 tokensSaved = elapsedTime * (accumulator.prevTokensPerSecond - accumulator.tokensPerSecond);
            fee = (tokensSaved * performanceFeeBPS) / BPS_DENOMINATOR;
        }

        tokensGained = (elapsedTime * accumulator.tokensPerSecond) - fee;
    }

    function _adjustClaimableTokens(
        Accumulator storage accumulator,
        uint256 performanceFeeBPS,
        uint256 lastEpochUpdateTimestamp
    ) internal {
        (uint256 tokensGained, uint256 fee) =
            _calculateClaimableTokensAndFee(accumulator, performanceFeeBPS, block.timestamp, lastEpochUpdateTimestamp);

        accumulator.claimableTokens += tokensGained;
        accumulator.claimableFees += fee;
    }

    function _maxReduction(uint256 tokensPerSecond, uint256 reductionFactor) internal pure returns (uint256) {
        return (tokensPerSecond * reductionFactor) / PRECISION;
    }

    function _maxIncrease(uint256 tokensPerSecond, uint256 increaseFactor) internal pure returns (uint256) {
        return (tokensPerSecond * increaseFactor) / PRECISION;
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import "../../static/Structs.sol";
import {ISafetyFactorOracle} from "../../interfaces/ISafetyFactorOracle.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract MockSafetyFactorOracle is ISafetyFactorOracle, Ownable {
    mapping(uint32 => SafetyFactorSnapshot) public safetyFactorSnapshots;
    mapping(uint32 => address) public protocolToReservesManager;
    mapping(uint32 => bool) public activeProtocols;

    // Safety Factor snapshots for each protocol

    function addProtocol(uint32 protocolIdToAdd, address reservesManager) external {
        // add the protocol
        activeProtocols[protocolIdToAdd] = true;
        protocolToReservesManager[protocolIdToAdd] = reservesManager;
    }

    function removeProtocol(uint32 protocolIdToRemove) external {
        // remove the protocol
        activeProtocols[protocolIdToRemove] = false;
        protocolToReservesManager[protocolIdToRemove] = address(0);
    }

    function setSafetyFactor(uint32 protocolId, int256 safetyFactor) external {
        // set the safety factor
        safetyFactorSnapshots[protocolId] = SafetyFactorSnapshot(safetyFactor, block.timestamp);
    }

    function getSafetyFactor(uint32 protocolId) external view returns (SafetyFactorSnapshot memory safetyFactor) {
        return safetyFactorSnapshots[protocolId];
    }
}

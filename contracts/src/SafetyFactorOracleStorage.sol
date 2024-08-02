// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import "./static/Structs.sol";
import {ISafetyFactorOracle} from "./interfaces/ISafetyFactorOracle.sol";

abstract contract SafetyFactorOracleStorage is ISafetyFactorOracle {
    address public anzenTaskManager; // Anzen Task Manager address
    address public anzenGov; // Anzen Governance address
    address public fallBackSafetyFactorPoster; // Fallback safety factor poster in the case of a pending $EIGEN dispute
    address public avsReservesManagerFactory; // AVSReservesManagerFactory address

    bool public pendingDispute; // Pending dispute status

    mapping(uint32 => SafetyFactorSnapshot) public safetyFactorSnapshots; // Safety Factor snapshots for each protocol
    mapping(uint32 => address) public protocolToReservesManager; // Mapping of protocol to AVSReservesManager
    mapping(uint32 => bool) public activeProtocols; // Mapping of protocol to active status
}

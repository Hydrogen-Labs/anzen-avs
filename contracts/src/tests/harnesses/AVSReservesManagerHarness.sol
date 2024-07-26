// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../../AVSReservesManager.sol";

contract AVSReservesManagerHarness is AVSReservesManager {
    constructor(address _avsServiceManager) AVSReservesManager(_avsServiceManager) {}

    // Expose internal functions for testing
    function createAllRangePayments() public returns (IPaymentCoordinator.RangePayment[] memory) {
        return _createAllRangePayments();
    }
}

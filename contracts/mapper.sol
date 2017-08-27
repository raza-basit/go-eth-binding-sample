pragma solidity ^0.4.12;

contract Mapper {

    event AddressMapped(address primary, address secondary);
    event Error(uint code, address sender);

    mapping (address => address) public primaryToSecondary;
    mapping (address => bool) public secondaryInUse;

    modifier secondaryAddressMustBeUnique(address secondary) {
        if(secondaryInUse[secondary]) {
            Error(1, msg.sender);
            throw;
        }
        _;
    }

    function mapAddress(address secondary)
        secondaryAddressMustBeUnique(secondary) {
        // If there is no mapping, this does nothing
        secondaryInUse[primaryToSecondary[msg.sender]] = false;

        primaryToSecondary[msg.sender] = secondary;
        secondaryInUse[secondary] = true;

        AddressMapped(msg.sender, secondary);
    }
}

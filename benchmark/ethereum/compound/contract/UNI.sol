pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract UNI is ERC20 {

    constructor(string memory tokenName,string memory tokenSymbol) ERC20(tokenName, tokenSymbol) {
    }

    function mint(address _address, uint _amount) public {
        _mint(_address, _amount);
    }

}
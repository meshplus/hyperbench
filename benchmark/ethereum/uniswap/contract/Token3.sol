pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Token3 is ERC20 {
    constructor() ERC20("Token3", "TK3") {
    }

    // 铸造代币
    function mint(address recipient, uint256 amount) public {
        _mint(recipient, amount);
    }

    // 销毁代币
    function burn(address from, uint256 amount) public {
        _burn(from, amount);
    }
}
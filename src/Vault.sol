// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC4626.sol";
import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

contract Vault is ERC4626, ReentrancyGuard {
    IUniswapV2Router02 public uniswapRouter;
      address[] path;


    constructor(
        IERC20 _token,
        IUniswapV2Router02 _uniswapRouter
    ) ERC20("lootWIF", "lWIF") ERC4626(_token) {
        uniswapRouter = _uniswapRouter;
        path = new address[](2);
        path[0] = uniswapRouter.WETH();
        path[1] = address(_token);
   
    }

    function depositETH(uint amountOutMin, uint deadline) external payable nonReentrant {
        require(msg.value > 0, "Must send ETH");
        require(path[path.length - 1] == address(asset()), "Invalid path");

        // Swap ETH for tokens
        uint[] memory amounts = uniswapRouter.swapExactETHForTokens{ value: msg.value }(
            amountOutMin,
            path,
            address(this),
            deadline
        );
        uint256 amountReceived = amounts[amounts.length - 1];

        // Mint vault tokens to the LayerZero communicator
        uint256 shares = convertToShares(amountReceived);
        _mint(msg.sender, shares);
    }

    function convertToShares(uint256 assets) public view override returns (uint256) {
        // Simple 1:1 conversion for demonstration
        return assets;
    }
}

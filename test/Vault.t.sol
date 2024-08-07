// SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.23;

import {Test} from "forge-std/Test.sol";

import {MockPortal} from "omni/core/test/utils/MockPortal.sol";
import {ConfLevel} from "omni/core/src/libraries/ConfLevel.sol";

import {Vault} from "src/Vault.sol";
import {GreetingBook} from "src/GreetingBook.sol";

contract GreeterTest is Test {
    address user;
    MockPortal portal;
     Vault vault;
    IERC20 token;
    IUniswapV2Router02 uniswapRouter;
    address layerZeroCommunicator = address(0x123); // Replace with actual LayerZero Communicator address
    address owner = address(0x456); // Replace with actual owner address
    address[] path;


    function setUp() public {
        user = address(this);
        portal = new MockPortal();
        greeter = new Greeter(address(portal), address(0xdeadbeef));
           // Fork mainnet
        string memory rpcUrl = "https://eth.llamarpc.com";
        vm.createSelectFork(rpcUrl);

        // Set up contracts using mainnet addresses
        token = IERC20(0xEE2a03Aa6Dacf51C18679C516ad5283d8E7C2637); // Replace with the actual token address
        uniswapRouter = IUniswapV2Router02(0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D); // Replace with the Uniswap V2 Router address

        // Deploy the vault
        vault = new Vault(token, uniswapRouter);

        // Set the swap path (ETH -> Token)
        path = new address[](2);
        path[0] = uniswapRouter.WETH();
        path[1] = address(token);
    }

    function testGreet() public {
        string memory greeting = "Hello, world!";

        bytes memory xcalldata = abi.encodeCall(GreetingBook.greet, (user, greeting));

        vm.expectCall(
            address(portal),
            abi.encodeCall(
                MockPortal.xcall,
                (
                    portal.omniChainId(),
                    ConfLevel.Latest,
                    greeter.greetingBook(),
                    xcalldata,
                    greeter.XGREET_GAS_LIMIT()
                )
            )
        );
        
        uint256 fee = portal.feeFor(portal.omniChainId(), xcalldata, greeter.XGREET_GAS_LIMIT());
        greeter.greet{value: fee}(greeting);
    }

}

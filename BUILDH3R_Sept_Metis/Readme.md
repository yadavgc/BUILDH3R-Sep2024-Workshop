# BUILDH3R Sept Metis

- Clone Template Repo:
    ```sh
    git clone https://github.com/scaffold-eth/scaffold-eth-2.git
    cd scaffold-eth-2 && rm -rf .git
    ```

## Setup hardhat network for Local development:
### Install packages:
- Command:
    ```sh
    yarn install
    ```

### Start local Network:
- Do not close this tab. Command:
    ```sh
    yarn chain
    ```


### Deploy Contract:
- Command:
    ```sh
    yarn deploy
    ```

- Output:
    <details><summary> Detailed Output </summary><blockquote>

    ~~~
    > yarn deploy           
    Generating typings for: 2 artifacts in dir: typechain-types for target: ethers-v6
    Successfully generated 6 typings!
    Compiled 2 Solidity files successfully (evm target: london).
    deploying "YourContract" (tx: 0xb69d03305a7561c34568d7dcd4aa2efe1799fa7d6b4db637eed5c4278bee0359)...: deployed at 0x5FbDB2315678afecb367f032d93F642f64180aa3 with 532743 gas
    👋 Initial greeting: Building Unstoppable Apps!!!
    📝 Updated TypeScript contract definition file on ../nextjs/contracts/deployedContracts.ts
    ~~~

    </blockquote></details>

    <img src="./Assets/contract-deploy-to-hardhat.png">

### Start Frontend:
- Command:
    ```sh
    yarn start
    ```
- We can verify our app is running in `http://localhost:3000`


### Modify Frontend:

- Edit `packages/nextjs/app/page.tsx` file:
    ```javascript
    </div>
        <p className="text-center text-lg mt-4">
        <strong style={{ fontSize: '36px', color: '#04D1fF'}}>DEPLOY ON METIS BY GC</strong>
        </p>
        {/* <p className="text-center text-lg">
        Get started by editing{" "}
        <code className="italic bg-base-300 text-base font-bold max-w-full break-words break-all inline-block">
            packages/nextjs/app/page.tsx
        </code>
        </p> */}
        <p className="text-center text-lg">
        Edited smart contract{" "}
        <code className="italic bg-base-300 text-base font-bold max-w-full break-words break-all inline-block">
            YourContract.sol
        </code>{" "}
        in{" "}
        <code className="italic bg-base-300 text-base font-bold max-w-full break-words break-all inline-block">
            packages/hardhat/contracts
        </code>{ " "}
        By adding <code className="italic bg-base-300 text-base font-bold max-w-full break-words break-all inline-block">storeNumber</code> function
        </p>
    </div>
    ```

### Modify Backend:
- Edit `packages/hardhat/contracts`:
    ```javascript
    /**
    * Function that allows the owner to withdraw all the Ether in the contract
    * The function can only be called by the owner of the contract as defined by the isOwner modifier
    */
    function withdraw() public isOwner {
        (bool success, ) = owner.call{ value: address(this).balance }("");
        require(success, "Failed to send Ether");
    }
    // Mapping from address to number
    mapping(address => uint) public userNumbers;
    event NumberUpdated(address indexed user, uint number);
    // Function to store a number
    function storeNumber(uint _number) public {
        userNumbers[msg.sender] = _number;
        emit NumberUpdated(msg.sender, _number);
    }
    ```


### Re-Deploy Contract:
- Make sure `yarn chain` is running in background tab:
    ```sh
    yarn deploy
    ```

    Output:
    <details><summary> Detailed Output </summary><blockquote>

    ~~~
    > yarn deploy
    Generating typings for: 1 artifacts in dir: typechain-types for target: ethers-v6
    Successfully generated 6 typings!
    Compiled 1 Solidity file successfully (evm target: london).
    deploying "YourContract" (tx: 0xce0f865828791286d6d8177e470c079efd1ae83b310233714b241af52a22dc9f)...: deployed at 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512 with 574168 gas
    👋 Initial greeting: Building Unstoppable Apps!!!
    📝 Updated TypeScript contract definition file on ../nextjs/contracts/deployedContracts.ts
    ~~~
 
    </blockquote></details>


### Verify:
- Visit [http://localhost:3000](http://localhost:3000):
    <img src="./Assets/verify-ui.png">

- Visit: [http://localhost:3000/debug](http://localhost:3000/debug):
    <img src="./Assets/verify-be-function.png">

- Create new Greet:

    <img src="./Assets/newGretting.png">

- Blockexplorer [http://localhost:3000/blockexplorer](http://localhost:3000/blockexplorer):\
    <img src="./Assets/blockexplorer.png">



## Deploy To Metis Sepolia Network:
### Generate Account:
- Command:
    ```sh
    yarn generate
    ```

    <details><summary> Detailed Output </summary><blockquote>

    ~~~
    > yarn generate
    👛 Generating new Wallet
    📄 Private Key saved to packages/hardhat/.env file
    🪄 Generated wallet address: 0x763bb2A2935B7fb4f5E3027e1368F82169eD8c26
    ~~~

    </blockquote></details>



### Faucet Link:
- https://faucet-427702.uc.r.appspot.com/

### Add Metis Sepolia Config to `Hardhat.config.ts` file:
- Code Snippet:
    ```typescript
    andromeda : {
      url: "https://andromeda.metis.io/?owner=1888",
      accounts: [deployerPrivateKey],
      verify: {
        etherscan: {
          apiKey: "apikey is not required, just set a placeholder",
          apiUrl: "https://api.routescan.io/v2/network/mainnet/evm/1088/etherscan",
        },
      },
    },
    metisSepolia: {
      url: "https://sepolia.metisdevops.link/",
      accounts: [deployerPrivateKey],
        verify: {
          etherscan: {
            apiKey: "apikey is not required, just set a placeholder", 
            apiUrl: "https://sepolia.explorer.metisdevops.link",
          },
        },
    },
    ```

### Deploy contract to Metis Sepolia Network:
- Command:
    ```sh
    yarn deploy --network metisSepolia
    ```

    <details><summary> Detailed Output </summary><blockquote>

    ~~~
    > yarn deploy --network metisSepolia
    Nothing to compile
    No need to generate any newer typings.
    deploying "YourContract"
    (tx: 0xd3b729d5437b46a9d8c15314acba16652f417fe5e73e791e1825626af7559e79)...: deployed at 0x8B2144A6Bcf8695721467C38a2b440becA851202 with 573988 gas
    👋 Initial greeting: Building Unstoppable Apps!!!
    📝 Updated TypeScript contract definition file on ../nextjs/contracts/deployedContracts.ts
    ~~~

    </blockquote></details>
    
    

- Here, 
    - Contract Address: `0x8B2144A6Bcf8695721467C38a2b440becA851202`
    - Txn Hash: `0xd3b729d5437b46a9d8c15314acba16652f417fe5e73e791e1825626af7559e79`
    - Txn Link: [https://sepolia-explorer.metisdevops.link/tx/0xd3b729d5437b46a9d8c15314acba16652f417fe5e73e791e1825626af7559e79](https://sepolia-explorer.metisdevops.link/tx/0xd3b729d5437b46a9d8c15314acba16652f417fe5e73e791e1825626af7559e79)


### Deploy contract to Metis Andromeda  Network:
- Command:
    ```sh
    yarn deploy --network andromeda
    ```

    <details><summary> Detailed Output </summary><blockquote>

    ~~~
    > yarn deploy --network andromeda
    Nothing to compile
    No need to generate any newer typings.
    deploying "YourContract"
    (tx: 0x22f6ec948cd1e718cc441dc7b7e34a5709eacde82c3bf168917f8f23bb7665bf)...: deployed at 0x8B2144A6Bcf8695721467C38a2b440becA851202 with 573988 gas
    👋 Initial greeting: Building Unstoppable Apps!!!
    📝 Updated TypeScript contract definition file on ../nextjs/contracts/deployedContracts.ts
    ~~~

    </blockquote></details>
    
    

- Here, 
    - Contract Address: `0x8B2144A6Bcf8695721467C38a2b440becA851202`
    - Txn Hash: `0x22f6ec948cd1e718cc441dc7b7e34a5709eacde82c3bf168917f8f23bb7665bf`
    - Txn Link: [https://andromeda-explorer.metis.io/tx/0x22f6ec948cd1e718cc441dc7b7e34a5709eacde82c3bf168917f8f23bb7665bf](https://andromeda-explorer.metis.io/tx/0x22f6ec948cd1e718cc441dc7b7e34a5709eacde82c3bf168917f8f23bb7665bf)


### Update Frontend for Metis:
- Edit `/packages/nextjs/scaffold.config.ts`:
    ```javascript
    targetNetworks: [chains.metis],
    ```

- Visit [http://localhost:3000](http://localhost:3000). Confirm our UI is connected to Metis Network instead of Hardhat network:
    
    <img src="./Assets/confirm-ui-metis.png">

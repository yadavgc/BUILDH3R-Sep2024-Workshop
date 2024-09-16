# BUILDH3R Sept Leo

## System Setup
-  Install `rust` locally. Below command will install cli. [More deatils](https://www.rust-lang.org/tools/install)
    ```sh
    curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
    ```

- Install `leo`. [More Details](https://github.com/ProvableHQ/leo/):
    ```sh
    # Download the source code
    git clone https://github.com/ProvableHQ/leo/
    cd leo
    git checkout testnet-beta
    cargo install --path .
    ```

- Install `SnarkOS`. [More Details](https://github.com/ProvableHQ/snarkOS):
    ```sh
    git clone https://github.com/AleoHQ/snarkOS.git --depth 1
    cd snarkOS
    git checkout testnet-beta
    ./build_ubuntu.sh
    cargo install --path .
    ```

- Install SnarkVM. [More Details](https://github.com/ProvableHQ/snarkVM):
    ```sh
    cargo install snarkvm
    ```

- Make sure all the binaries are added in PATH.

## Create Aleo Project
### Create new project:
- Command:
    ```sh
    leo new gcsandab_leo_auction
    ```

- Since, we are trying to create an auction app, there must be bidder and owner. So let's create 2 bidder account and 1 owner.
- Command:
    ```sh
    snarkos account new
    ```
- So, run above command 3 times in terminal:
    <details><summary> Detailed Output </summary><blockquote>

    ~~~
    $ snarkos account new

    Private Key  <redacted>
        View Key  <redacted>
        Address  aleo1yhvl35uwjcy72hmacufn4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk

    $ snarkos account new

    Private Key  <redacted>
        View Key  <redacted>
        Address  aleo1cj9vcgknthw6dzs042ys9gxcr3ar095h80j7k37w7zxq69ycpsgqjencwa

    $ snarkos account new

    Private Key  <redacted>
        View Key  <redacted>
        Address  aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq
    ~~~
    
    </blockquote></details>

- In above image, suppose 1 a/c with Address `aleo1yhvl35uwjcy72hmacufn4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk` is owner and remaining 2nd and 3rd a/cs are bidders. 

- Modify `auction/src/main.leo`. Do not forget to replace owner address in `main.leo` file.

- `.env` is created during the creation of the project. By default, it contains:
    ```sh
    NETWORK=testnet
    PRIVATE_KEY=APrivateKey1zkp8CZNn3yeCseEtxuVPbDCwSyhGW6yZKUYKfgXmcpoGPWH
    ENDPOINT=https://api.explorer.aleo.org/v1
    ```

- Value of Endpoint must be changed with: `https://api.explorer.provable.com/v1`
- New `.env`:
    ```sh
    NETWORK=testnet
    PRIVATE_KEY=APrivateKey1zkp8CZNn3yeCseEtxuVPbDCwSyhGW6yZKUYKfgXmcpoGPWH
    ENDPOINT=https://api.explorer.provable.com/v1
    ```

### Run Test
#### 1st Bid:
- Now, the value of `PRIVATE_KEY` in `.env` must be replaced with the `PRIVATE_KEY` of the 1st bidder.
- Run `place_bid` transaction:
    - Code Snippet:
        ```sh
        leo run place_bid <Address> <Amount>
        ```
    
    - Command:
        ```sh
        leo run place_bid aleo1cj9vcgknthw6dzs042ys9gxcr3ar095h80j7k37w7zxq69ycpsgqjencwa 25u64
        ```
        <details><summary> Detailed Output </summary><blockquote>

        ~~~
        $ leo run place_bid aleo1cj9vcgknthw6dzs042ys9gxcr3ar095h80j7k37w7zxq69ycpsgqjencwa 25u64
            Leo ✅ Compiled 'gcsandab_leo_auction.aleo' into Aleo instructions

        ⛓  Constraints

        •  'gcsandab_leo_auction.aleo/place_bid' - 2,026 constraints (called 1 time)

        ➡️  Output

        • {
        owner: aleo1yhvl35uwjcy72hmacufn4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk.private,
        bidder: aleo1cj9vcgknthw6dzs042ys9gxcr3ar095h80j7k37w7zxq69ycpsgqjencwa.private,
        amount: 25u64.private,
        is_winner: false.private,
        _nonce: 3196613987589781425323511887799532353386645590231891087187113149337195731667group.public
        }

            Leo ✅ Finished 'gcsandab_leo_auction.aleo/place_bid'
        ~~~

        </blockquote></details>
       
        <img src="./Assets/1st-bidder.png">
        
#### 2nd Bid:
- Now, the value of `PRIVATE_KEY` in `.env` must be replaced with the `PRIVATE_KEY` of the 2nd bidder.
- Run `place_bid` transaction:
    - Code Snippet:
        ```sh
        leo run place_bid <Address> <Amount>
        ```
    
    - Command:
        ```sh
        leo run place_bid aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq 50u64
        ```
        <details><summary> Detailed Output </summary><blockquote>

        ~~~
        $ leo run place_bid aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq 50u64
            Leo ✅ Compiled 'gcsandab_leo_auction.aleo' into Aleo instructions

        ⛓  Constraints

        •  'gcsandab_leo_auction.aleo/place_bid' - 2,026 constraints (called 1 time)

        ➡️  Output

        • {
        owner: aleo1yhvl35uwjcy72hmacufn4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk.private,
        bidder: aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        amount: 50u64.private,
        is_winner: false.private,
        _nonce: 1470786995366990318525485214894290178443403556897021385974898265083187105668group.public
        }

            Leo ✅ Finished 'gcsandab_leo_auction.aleo/place_bid'
        ~~~

        </blockquote></details>
       
        <img src="./Assets/2nd-bidder.png">

#### Resolve Time:
- Since, both the bidder has succesfully bidded, it is the time of owner to resolve who wins the bid.
- Now, the value of `PRIVATE_KEY` in `.env` must be replaced with the `PRIVATE_KEY` of the owner.
- Run `resolve` transaction:
    - Code Snippet:
        ```sh
        leo run resolve <Output_Of_1st_Bid> <Output_Of_2nd_Bid>
        ```
    
    - Command:
        ```sh
        leo run resolve '{
        owner: aleo1yhvl35uwjcy72hmacufn4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk.private,
        bidder: aleo1cj9vcgknthw6dzs042ys9gxcr3ar095h80j7k37w7zxq69ycpsgqjencwa.private,
        amount: 25u64.private,
        is_winner: false.private,
        _nonce: 3196613987589781425323511887799532353386645590231891087187113149337195731667group.public
        }' '{
        owner: aleo1yhvl35uwjcy72hmacufn4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk.private,
        bidder: aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        amount: 50u64.private,
        is_winner: false.private,
        _nonce: 1470786995366990318525485214894290178443403556897021385974898265083187105668group.public
        }'
        ```
        <details><summary> Detailed Output </summary><blockquote>

        ~~~
        $ leo run resolve '{
        owner: aleo1yhvl35uwjcy72hmacuf> n4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk.private,
        b> idder: aleo1cj9vcgknthw6dzs042ys9gxcr3ar095h80j7k37w7zxq69ycpsgqjencwa.private,
        amount: 25u64.privat> e,
        is_winner: false.private,
        _nonce: 3196613987589> > 781425323511887799532353386645590231891087187113149337195731667group.public
        }' '{
        owner: aleo1yhvl35> > uwjcy72hmacufn4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk.private,
        bidder: aleo1m8mr8yq0nnfssfgyvcx48g7vd> ahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        amount:>  50u64.private,
        is_winner: false.private,
        _nonce: > > 1470786995366990318525485214894290178443403556897021385974898265083187105668group.public
        }'> 
            Leo ✅ Compiled 'gcsandab_leo_auction.aleo' into Aleo instructions

        ⛓  Constraints

        •  'gcsandab_leo_auction.aleo/resolve' - 2,161 constraints (called 1 time)

        ➡️  Output

        • {
        owner: aleo1yhvl35uwjcy72hmacufn4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk.private,
        bidder: aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        amount: 50u64.private,
        is_winner: false.private,
        _nonce: 1320646093331239823985603553966965645230521309170454603341482130970885656196group.public
        }

            Leo ✅ Finished 'gcsandab_leo_auction.aleo/resolve' 
        ~~~

        </blockquote></details>
       
        <img src="./Assets/bid-resolve.png">

#### Finish Bid:
- Keep same `.env` as we are running as owner.
- Run `finish` transaction:
    - Code Snippet:
        ```sh
        leo run finish <Resolve_Bid>
        ```
    
    - Command:
        ```sh
        leo run finish '{
        owner: aleo1yhvl35uwjcy72hmacufn4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk.private,
        bidder: aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        amount: 50u64.private,
        is_winner: false.private,
        _nonce: 1320646093331239823985603553966965645230521309170454603341482130970885656196group.public
        }'
        ```
        <details><summary> Detailed Output </summary><blockquote>

        ~~~
        $ leo run finish '{
        owner: aleo1yhvl35uwjcy72hmacufn> 4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk.private,
        bi> dder: aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        amount: 50u64.private> ,
        is_winner: false.private,
        _nonce: 13206460933312> > 39823985603553966965645230521309170454603341482130970885656196group.public
        }'> 
            Leo ✅ Compiled 'gcsandab_leo_auction.aleo' into Aleo instructions

        ⛓  Constraints

        •  'gcsandab_leo_auction.aleo/finish' - 2,026 constraints (called 1 time)

        ➡️  Output

        • {
        owner: aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        bidder: aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        amount: 50u64.private,
        is_winner: true.private,
        _nonce: 2886853821344537330355039323924468642572124950646717370882103731415592845104group.public
        }

            Leo ✅ Finished 'gcsandab_leo_auction.aleo/finish'
        ~~~

        </blockquote></details>
       
        <img src="./Assets/finish-bid.png">

    - According to above outputs, a/c with `Address`: `aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq` is the winner and is new owner.

#### Claim Bid:
- Now, the value of `PRIVATE_KEY` in `.env` must be replaced with the `PRIVATE_KEY` of the owner `aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq`.
- Run `claim` transaction:
    - Code Snippet:
        ```sh
        leo run claim <Finish_Bid>
        ```
    
    - Command:
        ```sh
        leo run claim '{
        owner: aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        bidder: aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        amount: 50u64.private,
        is_winner: true.private,
        _nonce: 2886853821344537330355039323924468642572124950646717370882103731415592845104group.public
        }'
        ```
        <details><summary> Detailed Output </summary><blockquote>

        ~~~
        $ leo run claim '{
        owner: aleo1m8mr8yq0nnfssfgyvcx48> g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        bid> der: aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        amount: 50u64.private,> 
        is_winner: true.private,
        _nonce: 2886853821344537> > 330355039323924468642572124950646717370882103731415592845104group.public
        }'> 
            Leo ✅ Compiled 'gcsandab_leo_auction.aleo' into Aleo instructions

        ⛓  Constraints

        •  'gcsandab_leo_auction.aleo/claim' - 2,021 constraints (called 1 time)

        ➡️  Output

        • {
        owner: aleo1m8mr8yq0nnfssfgyvcx48g7vdahprxc9rt5gt66e0dksvg7ktvxsr7euyq.private,
        amount: 50u64.private,
        _nonce: 2192410568513461525190833052818622574598429490901013028118806698994452895531group.public
        }

            Leo ✅ Finished 'gcsandab_leo_auction.aleo/claim'
        ~~~

        </blockquote></details>
       
        <img src="./Assets/claim-bid.png">

#### Deploy To Testnet:
- Command:
    ```sh
    leo deploy --network testnet
    ```


    <details><summary> Detailed Output </summary><blockquote>

    ~~~
    $ leo deploy --network testnet
        Leo ✅ Compiled 'gcsandab_leo_auction.aleo' into Aleo instructions
    📦 Creating deployment transaction for 'gcsandab_leo_auction.aleo'...


    Base deployment cost for 'gcsandab_leo_auction.aleo' is 11.4031 credits.

    +---------------------------+----------------+
    | gcsandab_leo_auction.aleo | Cost (credits) |
    +---------------------------+----------------+
    | Transaction Storage       | 2.782000       |
    +---------------------------+----------------+
    | Program Synthesis         | 7.621100       |
    +---------------------------+----------------+
    | Namespace                 | 1.000000       |
    +---------------------------+----------------+
    | Priority Fee              | 0.000000       |
    +---------------------------+----------------+
    | Total                     | 11.403100      |
    +---------------------------+----------------+

    Your current public balance is 12.25 credits.

    ? Do you want to submit deployment of program `gcsandab_leo_auction.aleo.aleo` to network testnet via endpoint https://api.explorer.provable.com/v1 using address aleo1yhvl35uwjcy72hmacufn4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk? (y
    ✔ Do you want to submit deployment of program `gcsandab_leo_auction.aleo.aleo` to network testnet via endpoint https://api.explorer.provable.com/v1 using address aleo1yhvl35uwjcy72hmacufn4h08ew7n7r97zq48vegqefuqtewh4grspu7gqk? · yes
    ✅ Created deployment transaction for 'gcsandab_leo_auction.aleo'

    Broadcasting transaction to https://api.explorer.provable.com/v1/testnet/transaction/broadcast...

    ⌛ Deployment at1wf6jldks0d3rqq5zgcc9sfglwfjafdst2t0fmu50elq5qg34zups5zkvl3 ('gcsandab_leo_auction.aleo') has been broadcast to https://api.explorer.provable.com/v1/testnet/transaction/broadcast.
    ~~~

    </blockquote></details>

    <img src="./Assets/deploy-to-testnet.png">


- Links: 
    - Aleo Program: [https://testnet.aleo.info/program/gcsandab_leo_auction.aleo](https://testnet.aleo.info/program/gcsandab_leo_auction.aleo)
    - Deploy Txn: [https://testnet.aleo.info/transition/at1wf6jldks0d3rqq5zgcc9sfglwfjafdst2t0fmu50elq5qg34zups5zkvl3](https://testnet.aleo.info/transition/at1wf6jldks0d3rqq5zgcc9sfglwfjafdst2t0fmu50elq5qg34zups5zkvl3) 


    <img src="./Assets/testnet-program.png">
package main

import (
        "bytes"
        "encoding/json"
        "flag"
        "fmt"
        "io"
        "log"
        "math/big"
        "os"
        "path/filepath"
        "text/template"

        "erc20/token"

        "github.com/consensys/gnark-crypto/ecc"
        "github.com/consensys/gnark/backend/groth16"
        groth16_bn254 "github.com/consensys/gnark/backend/groth16/bn254"
        "github.com/ethereum/go-ethereum/accounts/abi/bind"
        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/crypto"
        "github.com/ethereum/go-ethereum/ethclient"
)

type ProofPublicData struct {
        Proof struct {
                Ar struct {
                        X string
                        Y string
                }
                Krs struct {
                        X string
                        Y string
                }
                Bs struct {
                        X struct {
                                A0 string
                                A1 string
                        }
                        Y struct {
                                A0 string
                                A1 string
                        }
                }
                Commitments []struct {
                        X string
                        Y string
                }
        }
        PublicWitness []string
}

type Receipt struct {
        proof           token.Proof
        input           [65]*big.Int
        proofCommitment [2]*big.Int
}

var ChainId int64
var Network string
var HexPrivateKey string

func main() {
        // Initialize flags
        chainId := flag.Int64("chainId", 11155111, "chainId")
        network := flag.String("network", "https://eth-sepolia.g.alchemy.com/v2/RH793ZL_pQkZb7KttcWcTlOjPrN0BjOW", "network")
        privateKey := flag.String("privateKey", os.Getenv("PRIVATE_KEY"), "privateKey")
        erc20TokenAddr := flag.String("erc20TokenAddr", os.Getenv("ERC20_TOKEN_ADDR"), "erc20TokenAddr")
        outputDir := flag.String("outputDir", "hardhat/contracts", "outputDir")
        proofPath := flag.String("proofPath", "./hardhat/test/snark_proof_with_public_inputs.json", "proofPath")
        flag.Parse()

        // Set default values if necessary
        if *privateKey == "" {
                *privateKey = "df4bc5647fdb9600ceb4943d4adff3749956a8512e5707716357b13d5ee687d9"
        }
        if *erc20TokenAddr == "" {
                *erc20TokenAddr = "0xA234F9049720EDaDF3Ae697eE8bF762f2A03949A"
        }

        // Set global variables
        ChainId = *chainId
        Network = *network
        HexPrivateKey = *privateKey

        if len(os.Args) < 2 {
                log.Printf("expected subcommands")
                os.Exit(1)
        }

        switch os.Args[1] {
        case "generate":
                generateVerifierContract(*outputDir)
        case "mint":
                if len(os.Args) < 3 {
                        log.Fatalf("Expected contract address argument")
                }
                mintTokenFromProof(*erc20TokenAddr, *proofPath)
        }
}

func authInstance() (*bind.TransactOpts) {
        unlockedKey, err := crypto.HexToECDSA(HexPrivateKey)
        if err != nil {
                log.Fatalf("Failed to create authorized transactor: %v", err)
        }
        auth, err := bind.NewKeyedTransactorWithChainID(unlockedKey, big.NewInt(ChainId))
        if err != nil {
                log.Fatalf("Failed to create authorized transactor: %v", err)
        }
        auth.GasLimit = 1000000

        return auth
}

func clientInstance() (*ethclient.Client) {
        client, err := ethclient.Dial(Network)
        if err != nil {
                log.Fatalf("Failed to create eth client: %v", err)
        }

        return client
}

func generateVerifierContract(outputDir string) {
        tmpl, err := template.ParseFiles("verifier/verifier.sol.tmpl")
        if err != nil {
                log.Fatal(err)
        }

        type VerifyingKeyConfig struct {
                Alpha     string
                Beta      string
                Gamma     string
                Delta     string
                Gamma_abc string
        }

        var config VerifyingKeyConfig
        var vkBN254 = groth16.NewVerifyingKey(ecc.BN254)

        fVk, _ := os.Open("verifier/verifying.key")
        vkBN254.ReadFrom(fVk)
        defer fVk.Close()

        vk := vkBN254.(*groth16_bn254.VerifyingKey)

        config.Alpha = fmt.Sprint("Pairing.G1Point(uint256(", vk.G1.Alpha.X.String(), "), uint256(", vk.G1.Alpha.Y.String(), "))")
        config.Beta = fmt.Sprint("Pairing.G2Point([uint256(", vk.G2.Beta.X.A0.String(), "), uint256(", vk.G2.Beta.X.A1.String(), ")], [uint256(", vk.G2.Beta.Y.A0.String(), "), uint256(", vk.G2.Beta.Y.A1.String(), ")])")
        config.Gamma = fmt.Sprint("Pairing.G2Point([uint256(", vk.G2.Gamma.X.A0.String(), "), uint256(", vk.G2.Gamma.X.A1.String(), ")], [uint256(", vk.G2.Gamma.Y.A0.String(), "), uint256(", vk.G2.Gamma.Y.A1.String(), ")])")
        config.Delta = fmt.Sprint("Pairing.G2Point([uint256(", vk.G2.Delta.X.A0.String(), "), uint256(", vk.G2.Delta.X.A1.String(), ")], [uint256(", vk.G2.Delta.Y.A0.String(), "), uint256(", vk.G2.Delta.Y.A1.String(), ")])")
        config.Gamma_abc = fmt.Sprint("vk.gamma_abc = new Pairing.G1Point[](", len(vk.G1.K), ");\n")
        for k, v := range vk.G1.K {
                config.Gamma_abc += fmt.Sprint("        vk.gamma_abc[", k, "] = Pairing.G1Point(uint256(", v.X.String(), "), uint256(", v.Y.String(), "));\n")
        }
        var buf bytes.Buffer
        err = tmpl.Execute(&buf, config)
        if err != nil {
                log.Fatal(err)
        }
        fSol, _ := os.Create(filepath.Join(outputDir, "verifier.sol"))
        _, err = fSol.Write(buf.Bytes())
        if err != nil {
                log.Fatal(err)
        }
        fSol.Close()
        log.Println("success")
}

func generateReceipt(proofPath string) Receipt {
        jsonFile, err := os.Open(proofPath)
        if err != nil {
                log.Fatal(err)
        }
        defer jsonFile.Close()

        byteValue, err := io.ReadAll(jsonFile)
        if err != nil {
                log.Fatal(err)
        }

        proofPublicData := ProofPublicData{}
        err = json.Unmarshal(byteValue, &proofPublicData)
        if err != nil {
                log.Fatal(err)
        }

        var input [65]*big.Int
        for i := 0; i < len(proofPublicData.PublicWitness); i++ {
                input[i], _ = new(big.Int).SetString(proofPublicData.PublicWitness[i], 0)
        }

        var proof = token.Proof{}
        proof.A.X, _ = new(big.Int).SetString(proofPublicData.Proof.Ar.X, 0)
        proof.A.Y, _ = new(big.Int).SetString(proofPublicData.Proof.Ar.Y, 0)

        proof.B.X[0], _ = new(big.Int).SetString(proofPublicData.Proof.Bs.X.A0, 0)
        proof.B.X[1], _ = new(big.Int).SetString(proofPublicData.Proof.Bs.X.A1, 0)
        proof.B.Y[0], _ = new(big.Int).SetString(proofPublicData.Proof.Bs.Y.A0, 0)
        proof.B.Y[1], _ = new(big.Int).SetString(proofPublicData.Proof.Bs.Y.A1, 0)

        proof.C.X, _ = new(big.Int).SetString(proofPublicData.Proof.Krs.X, 0)
        proof.C.Y, _ = new(big.Int).SetString(proofPublicData.Proof.Krs.Y, 0)

        var proofCommitment [2]*big.Int
        proofCommitment[0], _ = new(big.Int).SetString(proofPublicData.Proof.Commitments[0].X, 0)
        proofCommitment[1], _ = new(big.Int).SetString(proofPublicData.Proof.Commitments[0].Y, 0)

        return Receipt{proof, input, proofCommitment}
}

func mintTokenFromProof(addr string, proofPath string) {
        receipt := generateReceipt(proofPath)

        erc20, err := token.NewTestToken(common.HexToAddress(addr), clientInstance())
        if err != nil {
                fmt.Printf("Failed to instantiate a Token contract: %v", err)
                return
        }

        name, err := erc20.Name(nil)
        if err != nil {
                fmt.Printf("Failed to get name: %v", err)
                return
        }
        fmt.Printf("name: %v\n", name)

        totalSupply, err := erc20.TotalSupply(nil)
        if err != nil {
                fmt.Printf("Failed to get total supply: %v", err)
                return
        }
        fmt.Printf("Total Supply: %v\n", totalSupply)

        // Ensure the second argument (recipient address) is passed and is valid
        if len(os.Args) < 3 {
                log.Fatalf("Recipient address not provided")
        }

        recipient := common.HexToAddress(os.Args[2])
        tx, err := erc20.MintWithProof(authInstance(), recipient, big.NewInt(14077144391565514), receipt.proof, receipt.input, receipt.proofCommitment)
        if err != nil {
                log.Fatalf("Failed to mint token with proof: %v", err)
        }
        log.Printf("Mint with proof txHash: %s\n", tx.Hash().Hex())
}
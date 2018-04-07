# CoderSchool Assigment Week03 - Salmon Supply Chain

# Commands

Generate channel artifacts

```
configtxgen -profile ThreeOrgsChannel -outputCreateChannelTx transfers.tx -channelID transfers
configtxgen -profile ThreeOrgsChannel -outputCreateChannelTx channel/fredrick-bob.tx -channelID fredrick-bob
configtxgen -profile ThreeOrgsChannel -outputCreateChannelTx channel/fredrick-alice.tx -channelID fredrick-alice
```

## Chain code install & instantiate:

Get inside cli container

```
docker exec -it cli.coderschool.vn bash
```

Get dependencies

```
cd $GOPATH/src/github.com/tiengtinh/salmon-supply-chain/chaincode && go get -d ./...
cd $GOPATH/src/github.com/tiengtinh/salmon-supply-chain/chaincode/salmon && go get github.com/Pallinder/go-randomdata
```

```
export SALMON=salmon
export SALMON_VERSION=v2
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/coderschool.vn/msp/tlscacerts/tlsca.coderschool.vn-cert.pem
```

Install salmon chaincode

```
peer chaincode install --name $SALMON --version $SALMON_VERSION --lang golang --path github.com/tiengtinh/salmon-supply-chain/chaincode/salmon
```

Instantiate the chain code

```
peer chaincode instantiate -o orderer.coderschool.vn:7050 --tls true --cafile $ORDERER_CA -C transfers -n $SALMON -v $SALMON_VERSION -c '{"Args":[]}'
```

Invoke the init function

```
peer chaincode invoke --tls true --cafile $ORDERER_CA -o orderer.coderschool.vn:7050  -C transfers -n $SALMON -v $SALMON_VERSION -c '{"Args":["initLedger"]}'
```

Query

```
peer chaincode query -C transfers -n $SALMON -v $SALMON_VERSION -c '{"Args":["querySalmon","1"]}'
peer chaincode query -C transfers -n $SALMON -v $SALMON_VERSION -c '{"Args":["queryAllSalmon"]}'
```

---

```
docker exec cli.coderschool.vn peer chaincode install --name salmon --version v1 --lang golang --path github.com/tiengtinh/salmon-supply-chain/chaincode
```

Instantiate the chain code

```
# with cafile
docker exec -e "CORE_PEER_LOCALMSPID=FredrickMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/fredrick.coderschool.vn/users/Admin@fredrick.coderschool.vn/msp" cli.coderschool.vn peer chaincode instantiate -o orderer.coderschool.vn:27050 --tls true --cafile "/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/coderschool.vn/msp/tlscacerts/tlsca.coderschool.vn-cert.pem" -C transfers -n salmon -v v1 -c '{"Args":[""]}'

# with cafile inside container
peer chaincode instantiate -o orderer.coderschool.vn:7050 --tls true --cafile "/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/coderschool.vn/msp/tlscacerts/tlsca.coderschool.vn-cert.pem" -C transfers -n salmon -v v1 -c '{"Args":["a", "b"]}'

```

Invoke it!

peer chaincode query -C transfers -n salmon -v v1 -c '{"Args":["get","a"]}'

```
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n marbles -c '{"function":"read_everything","Args":[""]}'
```

peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n marbles -c '{"function":"read_everything","Args":[""]}'

```
docker rm -f $(docker ps -aq)
docker network prune
```

node helper/init.js
# CoderSchool Assigment Week03 - Salmon Supply Chain

# Commands

Generate channel artifacts

```
configtxgen -profile ThreeOrgsChannel -outputCreateChannelTx transfers.tx -channelID transfers
configtxgen -profile ThreeOrgsChannel -outputCreateChannelTx channel/fredrick-bob.tx -channelID fredrick-bob
configtxgen -profile ThreeOrgsChannel -outputCreateChannelTx channel/fredrick-alice.tx -channelID fredrick-alice
```

Install chain code

```
docker exec -e "CORE_PEER_LOCALMSPID=FredrickMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/fredrick.coderschool.vn/users/Admin@fredrick.coderschool.vn/msp" cli.coderschool.vn peer chaincode install --name salmon --version v1 --lang golang --path github.com/tiengtinh/salmon-supply-chain/chaincode

# inside container
peer chaincode install --name salmon --version v1 --lang golang --path github.com/tiengtinh/salmon-supply-chain/chaincode
```

Instantiate the chain code

```
docker exec -e "CORE_PEER_LOCALMSPID=FredrickMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/fredrick.coderschool.vn/users/Admin@fredrick.coderschool.vn/msp" cli.coderschool.vn peer chaincode instantiate -o orderer.coderschool.vn:27050 -C transfers -n salmon -v v1 -c '{"Args":[""]}'

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
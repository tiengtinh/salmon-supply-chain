# CoderSchool Assigment Week03 - Salmon Supply Chain

See "Setting up and API testing guide"

# Setting up and API testing guide

For development purpose, artifacts are commited into the repo so one wouldn't need to go through the process of recreating them every time.

## setting up

`cd` into `network` folder

Grant executable permission for scripts:

```
chmod +x ./teardown.sh && chmod +x ./testAPIs.sh && chmod +x ./setup.sh
```

To be sure, we should start clean by clearing up pending containers and misc stuff:

```
./teardown.sh
```

Start the docker containers

```
docker-compose up -d
```

Setting up users, channels, joining users to channels and deploying chaincodes

```
./setup.sh
```

Start the api server

```
npm start
```

## API testing

To interact with data endpoints, first of all, you need to signin with the endpoint `http://localhost:8080/signin` by either `orgName` as `fredrick`, `alice`, `bob`, or `regulator`. No password is needed because I trust people are honest about who they are (actuall, it's just not implemented). Have a look at [testAPIs.sh](network/testAPIs.sh) to have an idea of how this works and what APIs are available.

You can also execute the file to run through all scenarios: Inside folder `network` run `./testAPIs.sh`

# Misc commands for development

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

Query

```
peer chaincode query -C transfers -n $SALMON -v $SALMON_VERSION -c '{"Args":["querySalmon","1"]}'
peer chaincode query -C transfers -n $SALMON -v $SALMON_VERSION -c '{"Args":["queryAllSalmon"]}'
```

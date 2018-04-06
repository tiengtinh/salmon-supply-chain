'use strict'
var util = require('util');
var path = require('path');
var hfc = require('fabric-client');

// More verbose debugging
var log4js = require('log4js');
var logger = log4js.getLogger('Helper');
logger.setLevel('DEBUG');
hfc.setLogger(logger);

// indicate to the application where the setup file is located so it able
// to have the hfc load it to initalize the fabric client instance
hfc.setConfigSetting('network-connection-profile-path',path.join(__dirname, 'config', 'network-config.yaml'));
hfc.setConfigSetting('fredrick-connection-profile-path',path.join(__dirname, 'config', 'fredrick.yaml'));
hfc.setConfigSetting('alice-connection-profile-path',path.join(__dirname, 'config', 'alice.yaml'));
hfc.setConfigSetting('bob-connection-profile-path',path.join(__dirname, 'config', 'bob.yaml'));

const getRegisteredUser = require('./getRegisteredUser')
const createChannel = require('./createChannel')
const joinChannel = require('./joinChannel')

async function start() {
  try {
    
    logger.info('--- Getting register users ---')
    
    const fredrickUser1 = await getRegisteredUser('user1', 'fredrick', true)
    logger.info('fredrickUser1: ', fredrickUser1)
    const aliceUser1 = await getRegisteredUser('user1', 'alice', true)
    logger.info('aliceUser1: ', aliceUser1)
    const bobUser1 = await getRegisteredUser('user1', 'bob', true)
    logger.info('bobUser1: ', bobUser1)
    
    logger.info('--- Creating channels ---')

    const channelTransfersResult = await createChannel('transfers','./channel/transfers.tx','user1','fredrick')
    logger.info('channelTransfersResult ', channelTransfersResult)
    const channelFredrickAliceResult = await createChannel('fredrick-alice','./channel/fredrick-alice.tx','user1','fredrick')
    logger.info('channelFredrickAliceResult ', channelFredrickAliceResult)
    const channelFredrickBobResult = await createChannel('fredrick-bob','./channel/fredrick-bob.tx','user1','fredrick')
    logger.info('channelFredrickBobResult ', channelFredrickBobResult)

    logger.info('--- Joining channels ---')

    logger.info('--- --- Joining transfers channel ---')
    const fredrickJoinTransferResult = await joinChannel("transfers",["peer0.fredrick.coderschool.vn"], "admin", "fredrick")
    logger.info('fredrickJoinTransferResult ', fredrickJoinTransferResult)
    const aliceJoinTransferResult = await joinChannel("transfers",["peer0.alice.coderschool.vn"], "admin", "alice")
    logger.info('aliceJoinTransferResult ', aliceJoinTransferResult)
    const bobJoinTransferResult = await joinChannel("transfers",["peer0.bob.coderschool.vn"], "admin", "bob")
    logger.info('bobJoinTransferResult ', bobJoinTransferResult)
    
    logger.info('--- --- Joining fredrick-bob channel ---')
    const fredrickJoinFredrickBobResult = await joinChannel("fredrick-bob",["peer0.fredrick.coderschool.vn"],"admin","fredrick")
    logger.info('fredrickJoinFredrickBobResult ', fredrickJoinFredrickBobResult)
    const bobJoinFredrickBobResult = await joinChannel("fredrick-bob",["peer0.bob.coderschool.vn"],"admin","bob")
    logger.info('bobJoinFredrickBobResult ', bobJoinFredrickBobResult)

    logger.info('--- --- Joining fredrick-alice channel ---')
    const fredrickJoinFredrickAliceResult = await joinChannel("fredrick-alice",["peer0.fredrick.coderschool.vn"],"admin","fredrick")
    logger.info('fredrickJoinFredrickAliceResult ', fredrickJoinFredrickAliceResult)
    const aliceJoinFredrickAliceResult = await joinChannel("fredrick-alice",["peer0.alice.coderschool.vn"],"admin","alice")
    logger.info('aliceJoinFredrickAliceResult ', aliceJoinFredrickAliceResult)

  } catch (err) {
    console.error(err)
  }
}

start()

const path = require('path')
const hfc = require('fabric-client');

// indicate to the application where the setup file is located so it able
// to have the hfc load it to initalize the fabric client instance
console.log('__dirname', __dirname)
const configPath = path.join(__dirname, '../..', 'network', 'config')
hfc.setConfigSetting('network-connection-profile-path', path.join(configPath, 'network-config.yaml'));
hfc.setConfigSetting('fredrick-connection-profile-path', path.join(configPath, 'fredrick.yaml'));
hfc.setConfigSetting('alice-connection-profile-path', path.join(configPath, 'alice.yaml'));
hfc.setConfigSetting('bob-connection-profile-path', path.join(configPath, 'bob.yaml'));

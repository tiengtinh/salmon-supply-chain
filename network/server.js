'use strict';

const express = require('express');

require('./context')
const invoke = require('./invoke-transaction')
const query = require('./query')
const helper = require('./helper')
var logger = helper.getLogger('server');

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/', (req, res) => {
  res.send('Home\n');
});

app.get('/profile', async (req, res) => {
  const user1 = await helper.getRegisteredUser('user1', 'fredrick', true)
  res.send(user1)
})

app.get('/salmons/:id', async (req, res) => {
  try {
    const channelName = 'transfers'
    const chaincodeName = 'salmon'
    const fcn = "querySalmon"
    const args = [req.params.id]

    let message = await query.queryChaincode([], channelName, chaincodeName, args, fcn, 'user1', 'fredrick');
    res.send(message);
  } catch (err) {
    console.error(err)
    res.status(500).send({error: err.message})
  }
})

app.get('/salmons', async (req, res) => {
  try {
    const channelName = 'transfers'
    const chaincodeName = 'salmon'
    const fcn = "queryAllSalmon"
    const args = []

    let message = await query.queryChaincode([], channelName, chaincodeName, args, fcn, 'user1', 'fredrick');
    res.send(message);
  } catch (err) {
    console.error(err)
    res.status(500).send({error: err.message})
  }
})

app.post('/salmons', async (req, res) => {
  logger.debug('username :' + req.username);
  logger.debug('orgname:' + req.orgname);
  
  const {
    id,
    vessel,
    datetime,
    location,
    holder,
  } = req.body

  const peers = [] // Optional. The peers that will receive this request, when not provided the list of peers added to this channel object will be used.
  const channelName = 'transfers'
  const chaincodeName = 'salmon'
  const fcn = "recordSalmon"
  const args = [
    id,
    vessel,
    datetime,
    location,
    holder,
  ]

  const message = await invoke.invokeChaincode(peers, channelName, chaincodeName, fcn, args, req.username, req.orgname);
	res.send(message);
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);

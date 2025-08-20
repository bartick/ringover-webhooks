// Sets up Express app, loads middleware, routes, and initializes the server.

const express = require("express");
const bodyParser = require('body-parser');
const morgan = require('morgan');

const { NODE_ENV } = require("./utils/config");

const app = express();

app.use(morgan(NODE_ENV === "production" ? "combined" : "dev")); 
app.use(bodyParser.json());

const router = require('./routes');
app.use('/', router);

module.exports = app;

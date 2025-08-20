// Router
// This files defines the main API routes
//
// so here is says that there are only one top level route /ringover and ringoverRoute handles all inner routes like /ringover/*

const router = require('express').Router();

const ringoverRoutes = require("./ringover.router");
router.use("/ringover", ringoverRoutes);

module.exports = router;
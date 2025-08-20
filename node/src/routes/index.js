const router = require('express').Router();

const ringoverRoutes = require("./ringover.router");
router.use("/ringover", ringoverRoutes);

module.exports = router;
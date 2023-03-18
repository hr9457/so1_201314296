const express = require('express');
const router = express.Router();
const { memory, cpu } = require('../controller/api.controller.js');


// routes controller
router.get('/memory', memory);
router.get('/cpu', cpu);

module.exports = router;

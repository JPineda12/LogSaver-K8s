var express = require('express');
var router = express.Router();
const gameController = require('../controllers/gameController')
router.post('/exec-game', gameController.execGame);

router.get('/getLogs', gameController.getLogs);
module.exports = router;
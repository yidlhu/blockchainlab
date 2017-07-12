var express = require('express');
var router = express.Router();
var ctrler = require('../controllers/controller.js');

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});

router.get('/about', function(req, res) {
    res.send('About birds');
});

router.route('/api/logistics/newinfo')
    .post(ctrler.logisticsAddOne);

router.route('/api/sensor/newinfo')
    .post(ctrler.measuresAddOne);

router.route('/api/logistics/list')
    .get(ctrler.logisticsList);

router.route('/api/sensor/list/:batch')
    .get(ctrler.measuresList);

router.route('/api/sensor/list')
    .get(ctrler.measuresList);

// router.route('/api/blockchain/retrieve/:recordId')
//     .get(ctrler.measuresRetrieveOne);
//
// router.route('/api/blockchain/retrieve_all')
//     .get(ctrler.measuresRetrieveAll);


module.exports = router;

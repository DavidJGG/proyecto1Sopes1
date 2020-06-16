const express = require('express');
const router = express.Router();
const procesosController = require('../controllers/procesosController')

// routes
//router.get('/children', procesosController.miFuncion);
router.get('/children/:pid', procesosController.childProcess);

router.get('/', (req, res) =>{
    res.render('index.ejs')
    //res.sendFile(path.join(__dirname,'views/index.html'));    
});

router.get('/cpu', (req, res) =>{
    res.render('cpu.ejs')
    //res.sendFile(path.join(__dirname,'views/cpu.html'));    
});

router.get('/memory', (req, res) =>{
    res.render('memoria.ejs')
    //res.sendFile(path.join(__dirname,'views/memory.html'));    
});

router.get('/process', (req, res) =>{
    res.render('process.ejs')
    //res.sendFile(path.join(__dirname,'views/memory.html'));    
});
router.get('/about', (req, res) =>{
    res.render('about.ejs')
});



module.exports = router;
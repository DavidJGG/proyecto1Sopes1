var http = require('http');
//var axios = require('axios')


const controller = {};



controller.childProcess = (req, res) =>{
    
    res.render('children.ejs', {data: req.params.pid});
    /*
    axios.get('http://localhost:3000/process/child/'+req.params.pid)
        .then(function (response) {
        //console.log(response.data);
        res.render('children.ejs', {data: response.data});
    })
    .catch(function (error) {
        console.log(error);
    })
        .then(function () {
    });*/
    

    //console.log("PARAMS: ")
    //console.log(req)
   /*var options = {
        host:'localhost',
        port: 3000,
        path: '/process/child/'+req.params.pid,
        method: 'GET'
    };
    console.log(options.path)

    //console.log("HOLA: ")
    //console.log(req.params)
    var dataJSON = '';
    http.get(options, function(res2){
        var body = '';    
        res2.on('data', function(chunk){
            body+=chunk;
        });
        res2.on('end', function(){
            var price =  JSON.parse(body);
            dataJSON = body;
            //console.log("BODY:")
            //console.log(body)
            //console.log("price.key:\n"+price.key)
            //console.log("price:\n"+price)
            res.render('children.ejs', {data: body});
        });
    }).end();
    */
};







controller.miFuncion = (req, res) =>{
    
    //setInterval( function(){
        var dataJSON = '';
        http.get(options, function(res2){
            var body = '';    
            res2.on('data', function(chunk){
                body+=chunk;
            });
            res2.on('end', function(){
                var price =  JSON.parse(body);
                dataJSON = body;
                console.log("aqui")
                res.render('children.ejs');//, {data: price});
            });
        }).end();
    //}, 3000); 

};



module.exports = controller;


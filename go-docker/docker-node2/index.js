//use path module
const path = require('path');
//use express module
const express = require('express');
//use hbs view engine
const hbs = require('hbs');
//use bodyParser middleware
const bodyParser = require('body-parser');
//use mysql database
const mysql = require('mysql');
const app = express();
 
//Create connection
const conn = mysql.createConnection({
  host: 'mysql-dev',
  user: 'root',
  password: 'root',
  database: 'node_crud'
});
 
//connect to database
conn.connect((err) =>{
  if(err) throw err;
  console.log('Mysql Connected...');
});
 
//set views file
app.set('views',path.join(__dirname,'views'));
//set view engine
app.set('view engine', 'hbs');
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
//set public folder as static folder for static file
app.use('/assets',express.static(__dirname + '/public'));
 
//route for homepage
app.get('/',(req, res) => {
  let sql = "SELECT * FROM ESTUDIANTES";
  let query = conn.query(sql, (err, results) => {
    if(err) throw err;
    res.render('product_view',{
      results: results
    });
  });
});
 
//route for insert data
app.post('/save',(req, res) => {
  let data = {carne: req.body.carne, nombre: req.body.nombre,apellido: req.body.apellido, carrera: req.body.carrera};
  let sql = 'INSERT INTO ESTUDIANTES ( carne, nombre , apellido ,carrera) values (';
      sql += ' "'+req.body.carne+'", ';
      sql += ' "'+req.body.nombre+'", ';
      sql += ' "'+req.body.apellido+'", ';
      sql += ' "'+req.body.carrera+'"';
      sql += ' )';
  let query = conn.query(sql,(err, results) => {
    if(err) throw err;
    res.redirect('/');
  });
});
 
//route for update data
app.post('/update',(req, res) => {
  let sql = "UPDATE ESTUDIANTES SET estudiantes_name='"+req.body.estudiantes_name+"', estudiantes_price='"+req.body.estudiantes_price+"' WHERE estudiantes_id="+req.body.id;
  let query = conn.query(sql, (err, results) => {
    if(err) throw err;
    res.redirect('/');
  });
});
 
//route for delete data
app.post('/delete',(req, res) => {
	console.log(req.body.carne)
  let sql = "DELETE FROM ESTUDIANTES WHERE carne="+req.body.carne+"";
  let query = conn.query(sql, (err, results) => {
    if(err) throw err;
      res.redirect('/');
  });
});
 
//server listening
app.listen(8080, () => {
  console.log('Server is running at port 8000');
});

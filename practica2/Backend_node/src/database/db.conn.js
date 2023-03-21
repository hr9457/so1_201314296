const mysql = require('mysql2');
require("dotenv").config();

// pool connection
const conn = mysql.createConnection({
    host: process.env.HOSTBD,
    port: process.env.PORTBD,
    database: process.env.BDNAME,
    user: process.env.USERBD,
    password: process.env.PASSBD,    
});
conn.connect();

conn.connect(function (err){
    if(err) {
        console.log(`error  ${err}`);
        throw err;
    }
    conn.query('use proyecto;');
    console.log('Bd connect with api_bakcend');
    return conn;
});

module.exports = conn;

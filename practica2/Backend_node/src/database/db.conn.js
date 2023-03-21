const mysql = require('mysql2');
require("dotenv").config();

// pool connection
const conn = mysql.createConnection({
    host: process.env.hostdb,
    port: process.env.portdb,
    database: process.env.namedb,
    user: process.env.userdb,
    password: process.env.passdb,    
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

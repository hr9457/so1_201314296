const conn = require('../database/db.conn.js');


const memory = async (req,res) => {
    const consult = 'select * from memory order by id_memory DESC  limit 1;';
    conn.query(consult, (err, results) => {
        if (err) res.status(400).json({message: "Not find results"});
        res.status(200).json(results[0]);
    });
};


const cpu = async (req,res) => {
    const consult = 'select * from cpu order by id_cpu DESC limit 1;';
    conn.query(consult, (err, results) => {
        if (err) res.status(400).json({message: "Not find results"});
        results[0].data = JSON.parse(results[0].data)
        res.status(200).json(results[0]);
    });
};



module.exports = {
    memory,
    cpu
};

const express = require('express');
const cors = require('cors');
const apiRoutes = require('./routes/api.routes.js');


const createApp = () => {
    const app = express();
    app.use(cors());
    app.use(express.json());
    app.use('/api',apiRoutes);
    return app;
};


const app = createApp();


module.exports = app;
const express = require('express');
const app = express.Router();

app.get('/', (req, res) => {
    res.render('faqs');
});

module.exports = { app };
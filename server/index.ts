import express from "express";
import path from "path";

const app = express();

// Serve static files from the 'dist' directory
app.use(express.static(path.join(__dirname, '../client/dist')));

// Serve index.html for all other requests
app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, '../client/dist/index.html'));
});

app.listen(3000, () => {
    console.log('Server started on port 3000');
});

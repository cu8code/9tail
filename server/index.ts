import express from "express";
import path from "path";
import morgan from 'morgan';
import helmet from "helmet";

const app = express();

// Middleware for logging
app.use(morgan('combined'));

// Security headers
app.use(
    helmet.contentSecurityPolicy({
        directives: {
            "default-src": ["'none'"],
            "img-src": ["'self'"],
            "script-src": ["'self'", "https://cdn.jsdelivr.net"],
            "style-src": ["'self'", "https://cdn.jsdelivr.net", "'unsafe-inline'"],
            "style-src-elem": ["'self'", "https://cdn.jsdelivr.net", "'unsafe-inline'"],
            "style-src-attr": ["'self'", "https://cdn.jsdelivr.net", "'unsafe-inline'"],
            "font-src": ["'self'", "https://cdn.jsdelivr.net"],
            "worker-src": ["'self'", "blob:"],
        },
    })
);

// Serve static files from the assets directory
app.use('/assets', express.static(path.join(__dirname, '../client/dist/assets')));

// API endpoint
app.get("/api", (req, res) => {
    res.json({ message: "Hello from server!" });
});

// Serve the React app for all other routes
app.get('*', (req, res) => {
    res.sendFile(path.join(__dirname, '../client/dist/index.html'));
});

// Error handling middleware
app.use((req, res, next) => {
    const error = new Error('Not Found');
    next(error); // Pass the error to the next middleware
});


const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server started on port ${PORT}`);
});

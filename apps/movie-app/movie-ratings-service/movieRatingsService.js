const express = require('express');
const app = express();
const port = 3000;

// Sample movie ratings data (in-memory database)
const movieRatings = {
    1: { rating: 4.5, reviews: ["Great movie!", "Must watch!"] },
    2: { rating: 3.8, reviews: ["Good movie.", "Enjoyed it."] },
};

app.get('/api/movieratings/:movieId', (req, res) => {
    const movieId = req.params.movieId;
    if (movieRatings[movieId]) {
        res.json(movieRatings[movieId]);
    } else {
        res.status(404).json({ error: 'Movie rating not found' });
    }
});

app.listen(port, () => {
    console.log(`Movie Ratings Service listening at http://localhost:${port}`);
});

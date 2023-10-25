import requests, os
from flask import Flask, jsonify

app = Flask(__name__)


# Endpoint URLs for Movie Info and Movie Ratings services
# Access the environment variables
movie_info_url = os.environ.get('MOVIE_INFO_URL', 'http://default-movie-info-url')
movie_ratings_url = os.environ.get('MOVIE_RATINGS_URL', 'http://default-movie-ratings-url')


# Function to get movie details
@app.route('/api/movie/<int:movie_id>', methods=['GET'])
def get_movie_details(movie_id):
    # Make a request to the Movie Info service
    response_info = requests.get(f"{movie_info_url}/{movie_id}")
    movie_info = response_info.json()

    # Make a request to the Movie Ratings service
    response_ratings = requests.get(f"{movie_ratings_url}/{movie_id}")
    movie_ratings = response_ratings.json()

    # Combine movie info and ratings
    movie_details = {
        "movie_info": movie_info,
        "movie_ratings": movie_ratings,
    }

    return movie_details

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5000)



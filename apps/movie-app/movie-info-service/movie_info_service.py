from flask import Flask, jsonify

app = Flask(__name__)

# Sample movie data (in-memory database)
movies = {
    1: {"title": "Movie A", "year": 2023},
    2: {"title": "Movie B", "year": 2022},
}


@app.route('/api/movieinfo/<int:movie_id>', methods=['GET'])
def get_movie_info(movie_id):
    if movie_id in movies:
        return jsonify(movies[movie_id])
    else:
        return jsonify({"error": "Movie not found"}), 404


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5000)

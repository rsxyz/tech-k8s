
mkdir movie-catalog-service
cd movie-catalog-service

python3 -m venv .venv
source .venv/bin/activate
pip install Flask
pip install requests
touch movie_catalog_service.py
touch requirements.txt
python movie_info_service.py
curl http://localhost:5000/api/movie/1
pip freeze > requirements.txt 
deactivate

touch Dockerfile
touch .dockerignore
touch .gitignore

docker build -t movie-catalog-service .
docker run -p 5000:5000 movie-catalog-service
curl http://localhost:5000/api/movie/2


docker run -p 5000:5000 -e MOVIE_INFO_URL="http://your-movie-info-url" -e MOVIE_RATINGS_URL="http://your-movie-ratings-url" your-image-name

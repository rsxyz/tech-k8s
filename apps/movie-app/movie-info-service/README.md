
mkdir movie-info-service
cd movie-info-service

python3 -m venv .venv
source .venv/bin/activate
pip install Flask
touch movie_info_service.py
requirements.txt
python movie_info_service.py
curl http://localhost:5000/api/movieinfo/1
pip freeze > requirements.txt 
deactivate

touch Dockerfile
touch .dockerignore
touch .gitignore

docker build -t movie-info-service .
docker run -p 5000:5000 movie-info-service
curl http://localhost:5000/api/movieinfo/2
mkdir movie-ratings-service
cd movie-ratings-service
npm init -y
npm install express
touch movieRatingsService.js
node movieRatingsService.js
curl http://localhost:3000/api/movieratings/1

docker build -t movie-ratings-service .
docker run -p 3000:3000 movie-ratings-service
curl http://localhost:3000/api/movieratings/2

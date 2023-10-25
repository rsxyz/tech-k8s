# Use an official Node.js runtime as a parent image
FROM node:18.18.1-alpine3.18

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install application dependencies
RUN npm install

# Copy the rest of your application's source code to the container
COPY . .

# Expose the port on which your Node.js application will run (adjust if needed)
EXPOSE 3000

# Define the command to start your Node.js application
CMD [ "node", "movieRatingsService.js" ]

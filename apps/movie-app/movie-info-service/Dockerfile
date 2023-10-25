# Use an official Python runtime as a parent image
FROM python:3.9-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the Python requirements file into the container
COPY requirements.txt .

# Install application dependencies
RUN pip install -r requirements.txt

# Copy the rest of your application's source code to the container
COPY . .

# Expose the port on which your Flask application will run (adjust if needed)
EXPOSE 5000

# Define the command to start your Flask application
CMD ["python", "movie_info_service.py"]

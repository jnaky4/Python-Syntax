# Use the official Python image as the base image
FROM python

# Set the working directory in the container
WORKDIR /app

# Copy the requirements.txt file into the container at /app
COPY requirements.txt .

# Install Flask and other dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the entire Flask application into the container at /app
COPY . .

# Expose port 5000 to allow communication to/from the Flask web server
EXPOSE 5000

# Set environment variables
ENV FLASK_APP flask_display.py
ENV FLASK_RUN_HOST 0.0.0.0

# Command to run the Flask application
CMD ["flask", "run"]

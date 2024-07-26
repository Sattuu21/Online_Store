# Use an official Node.js runtime as a parent image
FROM node:18.16.0-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the rest of the application code to the working directory
COPY . .

# Install serve
RUN npm install -g serve

# Install any needed dependencies specified in package.json
RUN npm install

# Create a Production Build
RUN npm run build

# Expose port 3000 for the application
EXPOSE 3000

# Start the application
CMD [ "serve", "-s", "build" ]
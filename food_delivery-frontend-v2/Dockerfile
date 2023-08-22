# Use the official Node.js 18 image as the base image
FROM node:18-alpine

# Set the working directory in the container
WORKDIR /srv/app

# Copy the contents of your Vue.js frontend project into the container
COPY . .

# Make the entrypoint script executable (if not needed)
# RUN chmod +x /srv/app/entrypoint.sh

# Expose port 8000 (if your Vue.js app is running on this port)
EXPOSE 8000

# Install Node packages, build the app, and start the preview
CMD npm install && npm run build && npm run preview
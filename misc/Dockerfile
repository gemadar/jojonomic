# Use the official PostgreSQL base image
FROM postgres:latest

# Set environment variables for PostgreSQL
ENV POSTGRES_DB=fitness
ENV POSTGRES_USER=adminpostgres
ENV POSTGRES_PASSWORD=87654321

# Copy the SQL script for initialization
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the default PostgreSQL port
EXPOSE 5432
# Use the official PostgreSQL base image
FROM postgres:latest

# Set environment variables for PostgreSQL
ENV POSTGRES_DB=postgres
ENV POSTGRES_USER=adminpostgres
ENV POSTGRES_USER=87654321

# Copy the SQL script for initialization
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the default PostgreSQL port
EXPOSE 5431
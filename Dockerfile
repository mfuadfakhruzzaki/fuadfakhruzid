# Gunakan base image Go 1.23
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Install dependencies dan build aplikasi
RUN go mod tidy
RUN go build -o main .

# Membuat file .env di dalam container
RUN echo "GOOGLE_APPLICATION_CREDENTIALS=./config/fuadfakhruzid-firebase-adminsdk-fbsvc-41a1b40d6d.json" > .env && \
    echo "PORT=8081" >> .env && \
    echo "FIREBASE_PROJECT_ID=fuadfakhruzid" >> .env && \
    echo "GCS_KEY_FILE=./config/fuadfakhruzid-1fb2c006393e.json" >> .env && \
    echo "GCS_BUCKET_NAME=fuadfakhruzid" >> .env && \
    echo "PROFILE_PICTURE_PATH=profile_pictures" >> .env && \
    echo "CV_PATH=cv_uploads" >> .env

# Expose port 8081
EXPOSE 8081

# Jalankan aplikasi
CMD ["./main"]

# ---------- Build Stage ----------
FROM ghcr.io/cirruslabs/flutter:3.13.6 AS builder

WORKDIR /app

# Create non-root user
RUN useradd -ms /bin/bash flutteruser

# Make Flutter SDK writable by the new user
RUN chown -R flutteruser:flutteruser /sdks/flutter /root/.pub-cache

# Copy the entire project first
COPY . .

# Change ownership of entire project so flutteruser can write
RUN chown -R flutteruser:flutteruser /app

# Switch to non-root user
USER flutteruser

# Disable analytics
RUN flutter config --no-analytics

# Install dependencies
RUN flutter pub get

# Build the web app
RUN flutter build web --release

# ---------- Serve Stage ----------
FROM nginx:alpine

# Copy built web app from builder
COPY --from=builder /app/build/web /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

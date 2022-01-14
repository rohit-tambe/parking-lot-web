FROM nginx:latest

# Expost port 80
EXPOSE 80

# Copy custom configuration file from the current directory
COPY nginx.conf /etc/nginx/nginx.conf:ro

# COPY ./*.html /web/index.html

# COPY /web/*.html /usr/share/nginx/html/

# WORKDIR /usr/share/nginx/html/js
# COPY /web/js/*.js /usr/share/nginx/html/js/

# WORKDIR /var/www/js
# COPY /web/js/*.js /var/www/js/

# WORKDIR /usr/share/nginx/html/css
# COPY /web/css/*.css /usr/share/nginx/html/css/

# WORKDIR /var/www/css
# COPY /web/css/*.css /var/www/css/


# COPY /web/favicon.ico /usr/share/nginx/html/
FROM php:8.2-fpm-alpine

# 设置时区
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 安装依赖
RUN apk add --no-cache \
    git \
    curl \
    wget \
    zip \
    unzip \
    libpng-dev \
    libjpeg-turbo-dev \
    freetype-dev \
    libzip-dev \
    oniguruma-dev \
    icu-dev \
    libxml2-dev \
    imap-dev \
    imagemagick-dev

# 安装PHP扩展
RUN docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install -j$(nproc) \
    pdo \
    pdo_mysql \
    mysqli \
    mbstring \
    gd \
    zip \
    bcmath \
    intl \
    opcache \
    exif \
    pcntl \
    posix

# 安装Redis扩展
RUN pecl install redis \
    && docker-php-ext-enable redis

# 安装Imagick扩展
RUN pecl install imagick \
    && docker-php-ext-enable imagick

# 安装Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# 设置工作目录
WORKDIR /var/www/html

# 复制配置
COPY . /var/www/html/

# 设置权限
RUN chown -R www-data:www-data /var/www/html \
    && chmod -R 755 /var/www/html

# 暴露端口
EXPOSE 9000

# 启动PHP-FPM
CMD ["php-fpm"]

-- Удаляем старые таблицы (если есть)
DROP TABLE IF EXISTS items, delivery, payment, orders;

-- Создаем таблицы с правильными именами (все в нижнем регистре)
CREATE TABLE orders (
    order_uid VARCHAR(255) PRIMARY KEY,
    track_number VARCHAR(255),
    entry VARCHAR(50),
    locale VARCHAR(10),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255),
    delivery_service VARCHAR(100),
    shardkey VARCHAR(50),
    sm_id INTEGER,
    date_created TIMESTAMP WITH TIME ZONE,
    oof_shard VARCHAR(50)
);

CREATE TABLE delivery (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(255) REFERENCES orders(order_uid) ON DELETE CASCADE,
    name VARCHAR(255),
    phone VARCHAR(50),
    zip VARCHAR(50),
    city VARCHAR(255),
    address TEXT,
    region VARCHAR(255),
    email VARCHAR(255)
);

CREATE TABLE payment (
    transaction VARCHAR(255) PRIMARY KEY,
    order_uid VARCHAR(255) REFERENCES orders(order_uid) ON DELETE CASCADE,
    request_id VARCHAR(255),
    currency VARCHAR(10),
    provider VARCHAR(100),
    amount INTEGER,
    payment_dt BIGINT,
    bank VARCHAR(100),
    delivery_cost INTEGER,
    goods_total INTEGER,
    custom_fee INTEGER
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(255) REFERENCES orders(order_uid) ON DELETE CASCADE,
    chrt_id BIGINT,
    track_number VARCHAR(255),
    price INTEGER,
    rid VARCHAR(255),
    name VARCHAR(255),
    sale INTEGER,
    size VARCHAR(50),
    total_price INTEGER,
    nm_id BIGINT,
    brand VARCHAR(255),
    status INTEGER
);
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO order_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO order_user;
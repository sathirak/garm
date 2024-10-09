CREATE TABLE auth_users (
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    verified_email BOOLEAN NOT NULL DEFAULT FALSE,
    locale VARCHAR(10) NOT NULL DEFAULT 'en',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE auth_methods ( 
    id SERIAL PRIMARY KEY,
    method_name VARCHAR(50) UNIQUE
);

CREATE TABLE auth_credentials (
    id SERIAL PRIMARY KEY,
    auth_user_id VARCHAR(255),
    auth_method_id INT,
    auth_identifier VARCHAR(255), 
    auth_secret VARCHAR(255),     
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (auth_user_id) REFERENCES auth_users(id),
    FOREIGN KEY (auth_method_id) REFERENCES auth_methods(id)
);

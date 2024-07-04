CREATE TABLE images (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    type_name VARCHAR(64) NOT NULL,
    resource_id CHAR(36) UNIQUE NOT NULL,
    mime_type VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL
);

CREATE TABLE image_types (
    name VARCHAR(64) PRIMARY KEY,
    available_resolutions JSON
);

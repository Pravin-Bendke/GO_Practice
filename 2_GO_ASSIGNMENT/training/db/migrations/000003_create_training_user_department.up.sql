CREATE TABLE training_user_department(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    department_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES training_user(id),

    CONSTRAINT fk_department
    FOREIGN KEY (department_id)
    REFERENCES training_department(id),

    CONSTRAINT unique_user_department
    UNIQUE (user_id,department_id)
);
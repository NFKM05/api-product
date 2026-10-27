CREATE TABLE IF NOT EXISTS users(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(150) NOT NULL,
    email VARCHAR(150) NOT NULL,
    password VARCHAR(90) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    inactive_dt DATE,

    CONSTRAINT pk_users PRIMARY KEY(id)
);
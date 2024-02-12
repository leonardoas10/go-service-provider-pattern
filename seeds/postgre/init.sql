DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    role VARCHAR(255) NOT NULL,
    hobby VARCHAR(255) NOT NULL
);

INSERT INTO users (username, password, country, age, role, hobby) VALUES
    ('postgresSQL', '12345', 'Venezuela', 25, 'Technical Leader | Full Stack Developer', 'Bicycle, Run, Volley, Soccer'),
    ('postgresSQL', '67890', 'United States', 18, 'Full Stack Developer', 'Run, GYM, Football'),
    ('postgresSQL', '67890', 'United States', 33, 'Killer', 'Sing, GYM, Racing'),
    ('postgresSQL', '67890', 'United States', 40, 'Scrum Master', 'Football');
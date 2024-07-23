CREATE TABLE demodb.users(
    id INT NOT NULL AUTO_INCREMENT,
    created_at TIME,
    updated_at TIME,
    deleted_at TIME,

    PRIMARY KEY(id)
) ENGINE = innoDB DEFAULT CHARSET = utf8;
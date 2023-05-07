CREATE TABLE IF NOT EXISTS users
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'user id',
    openid VARCHAR(255) NOT NULL COMMENT 'wechat openid',
    created_at DATETIME NOT NULL COMMENT 'user first collect time',
    updated_at DATETIME NOT NULL COMMENT 'user last collect time',
    PRIMARY KEY(id),
    UNIQUE KEY(openid),
    KEY (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS documents
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'document id',
    url VARCHAR(512) NOT NULL COMMENT 'document url',
    hash VARCHAR(255) NOT NULL COMMENT 'url hash',
    title VARCHAR(255) COMMENT 'document title',
    description VARCHAR(255) COMMENT 'document description',
    created_at DATETIME NOT NULL COMMENT 'document collect time',
    PRIMARY KEY (id),
    UNIQUE (url),
    KEY (hash)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS user_docs
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'id',
    uid BIGINT NOT NULL COMMENT 'user id',
    doc_id BIGINT NOT NULL COMMENT 'document id',
    created_at DATETIME NOT NULL COMMENT 'user collect time',
    PRIMARY KEY (id),
    KEY (uid),
    KEY (doc_id)
    # FOREIGN KEY (uid) REFERENCES users(id),
    # FOREIGN KEY (doc_id) REFERENCES documents(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS query_logs
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'log id',
    uid BIGINT NOT NULL COMMENT 'user id',
    content varchar(255) NOT NULL COMMENT 'query content',
    created_at DATETIME NOT NULL COMMENT 'query time',
    PRIMARY KEY (id),
    KEY (uid)
    # FOREIGN KEY (uid) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE thourus.company
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uid VARCHAR(12) DEFAULT (LEFT(MD5(UUID()),12)) UNIQUE,
    name NVARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE thourus.space
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uid VARCHAR(12) DEFAULT (LEFT(MD5(UUID()),12)) UNIQUE,
    name NVARCHAR(255) NOT NULL,
    company_id INT NOT NULL,
    FOREIGN KEY (company_id) REFERENCES thourus.company (id)
);

CREATE TABLE thourus.project
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uid VARCHAR(12) DEFAULT (LEFT(MD5(UUID()),12)) UNIQUE,
    name NVARCHAR(255) NOT NULL,
    space_id INT NOT NULL,
    FOREIGN KEY (space_id) REFERENCES thourus.space (id)
);

CREATE TABLE thourus.role
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uid VARCHAR(12) DEFAULT (LEFT(MD5(UUID()),12)) UNIQUE,
    name NVARCHAR(255) NOT NULL
);

CREATE TABLE thourus.status
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uid VARCHAR(12) DEFAULT (LEFT(MD5(UUID()),12)) UNIQUE,
    name NVARCHAR(255) NOT NULL
);

CREATE TABLE thourus.user
(
    id int PRIMARY KEY AUTO_INCREMENT,
    uid VARCHAR(12) DEFAULT (LEFT(MD5(UUID()),12)) UNIQUE,
    name NVARCHAR(255) NOT NULL,
    surname NVARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    login NVARCHAR(20) NOT NULL,
    password VARCHAR(100) NOT NULL,
    role_id INT NOT NULL,
    company_id INT NOT NULL,
    registration_date DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (role_id) REFERENCES thourus.role (id),
    FOREIGN KEY (company_id) REFERENCES thourus.company (id)
);

CREATE TABLE thourus.document
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uid VARCHAR(12) DEFAULT (LEFT(MD5(UUID()),12)) UNIQUE,
    name NVARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL,
    date_created DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    creator_id INT NOT NULL,
    status_id INT NOT NULL,
    project_id INT NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES thourus.user (id),
    FOREIGN KEY (status_id) REFERENCES thourus.status (id),
    FOREIGN KEY (project_id) REFERENCES thourus.project (id)
);

CREATE TABLE thourus.history
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uid VARCHAR(12) DEFAULT (LEFT(MD5(UUID()),12)) UNIQUE,
    document_id INT NOT NULL,
    hash TEXT NOT NULL,
    pow INT NOT NULL,
    initiator_id INT NOT NULL,
    date_changed DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (document_id) REFERENCES thourus.document (id),
    FOREIGN KEY (initiator_id) REFERENCES thourus.user (id)
);

CREATE TABLE thourus.project_user_relations
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uid VARCHAR(12) DEFAULT (LEFT(MD5(UUID()),12)) UNIQUE,
    project_id INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (project_id) REFERENCES thourus.project (id),
    FOREIGN KEY (user_id) REFERENCES thourus.user (id)
);

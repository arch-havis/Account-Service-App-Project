CREATE DATABASE account_service;

USE account_service;

CREATE TABLE users (
	user_id INT AUTO_INCREMENT,
	no_telepon VARCHAR(255),
	nama VARCHAR(50) NOT NULL,
	password VARCHAR(255) NOT NULL,
	alamat TEXT,
	gender ENUM('laki-laki', 'perempuan'),
	saldo DECIMAL,
	PRIMARY KEY (user_id)
);

CREATE TABLE topup (
	topup_id INT NOT NULL AUTO_INCREMENT,
	user_id INT,
	nominal_topup DECIMAL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP(),
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP(),
	PRIMARY KEY (topup_id),
	CONSTRAINT FK_user_topup FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE 
);

CREATE TABLE transfer (
	transfer_id INT NOT NULL AUTO_INCREMENT,
	user_id INT,
	user_id_penerima INT,
	nominal_transfer DECIMAL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP(),
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP(),
	PRIMARY KEY (transfer_id),
	CONSTRAINT FK_user_transfer FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT FK_user_penerima FOREIGN KEY (user_id_penerima) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

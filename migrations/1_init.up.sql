BEGIN;

CREATE TABLE IF NOT EXISTS customers ( 
  customer_id INT AUTO_INCREMENT, 
  first_name VARCHAR(255) NOT NULL, 
  last_name VARCHAR(255) NOT NULL, 
  age INT NOT NULL, 
  customer_type VARCHAR(20) NOT NULL, 
  street VARCHAR(255) NOT NULL, 
  city VARCHAR(50) NOT NULL, 
  state VARCHAR(50) NOT NULL,
  zip_code VARCHAR(5) NOT NULL, 
  phone_number VARCHAR(15) NOT NULL 
);

ALTER TABLE customers MODIFY COLUMN costomer_id INT UNSIGNED NOT NULL AUTO_INCREMENT;

CREATE TABLE IF NOT EXISTS account(
  account_id INT UNSIGNED NOT NULL, 
  customer_id INT UNSIGNED NOT NULL, 
  balance DECIMAL, 
  age INT NOT NULL, 
  account_type VARCHAR(20) NOT NULL, 
  PRIMARY KEY (account_id),
  FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

COMMIT;
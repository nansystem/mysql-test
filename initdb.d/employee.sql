CREATE TABLE employees (
    employee_id INT UNSIGNED NOT NULL,
    subsidiary_id INT UNSIGNED NOT NULL,
    first_name VARCHAR(1000) NOT NULL,
    last_name VARCHAR(1000) NOT NULL,
    date_of_birth Date NOT NULL,
    phone_number VARCHAR(1000) NOT NULL,
    CONSTRAINT employees_pk PRIMARY KEY (employee_id, subsidiary_id)
) ENGINE = InnoDB;

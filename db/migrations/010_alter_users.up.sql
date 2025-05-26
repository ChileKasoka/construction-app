ALTER TABLE users
ADD COLUMN project_id INT,
ADD COLUMN active BOOLEAN DEFAULT TRUE,
ADD CONSTRAINT fk_project
    FOREIGN KEY (project_id) REFERENCES projects(id)
    ON DELETE SET NULL;  -- or CASCADE or RESTRICT, based on your business rules

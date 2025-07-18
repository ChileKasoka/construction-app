INSERT INTO roles (name, description) VALUES
('admin', 'To manage the entire system, including user management, role assignments, and overall project oversight. The Administrator has full access to all features and functionalities.'),
('Client', 'To view and monitor project progress, communicate with the project team, and provide feedback. The Client has limited access to specific features related to their projects.'),
('Construction Site Manager', 'To oversee daily operations on construction sites, ensuring that projects are completed safely, on time, within budget, and to required quality standards.'),
('Quantity Surveyor', 'To manage all costs related to construction projects, ensuring value for money while achieving the required standards and quality. They play a critical role in budgeting, cost control, contract management, and financial reporting'),
('Site Engineer', 'To oversee technical aspects of construction work on-site, ensuring that structures are built accurately according to design specifications and engineering standards, while coordinating with supervisors, contractors, and suppliers.'),
('Foreman', 'To supervise and coordinate the work of construction crews on-site, ensuring that daily tasks are carried out efficiently, safely, and according to the project plan. The Foreman acts as the direct link between laborers and management.'),
('Safety Manager', 'To ensure the construction site operates in full compliance with local, national, and company-specific health and safety regulations, minimizing risks to workers and ensuring a safe work environment.'),
('Safety Officer', 'To monitor and enforce safety practices on-site, ensuring that all workers comply with safety protocols, and that hazards are identified and addressed immediately. The Safety Officer acts as the eyes and ears of the Safety Manager on the ground.'),
('Land Surveyor', 'To measure and map land boundaries, elevations, and layout data that support the planning, design, and execution of construction projects. The Land Surveyor ensures accuracy in site positioning and helps prevent costly errors due to layout issues.'),
('Site clerk/ Administration', 'To provide administrative support on the construction site, ensuring all records, documents, communications, and logistics are organized and up-to-date. The Site Clerk acts as the central point for documentation, supplies, and coordination between field and office staff.'),
('Stores Man', 'To manage the receipt, storage, issuance, and inventory of construction materials, tools, and supplies on-site, ensuring proper documentation and control to support smooth project operations.'),
('Labor Officer', 'To manage labor force activities on-site, including recruitment, attendance, payroll coordination, and resolving labor-related issues, ensuring the workforce operates efficiently and in compliance with labor laws.'),
('Accountant', 'To manage financial records, process invoices and payments, track project budgets, and ensure accurate accounting in compliance with company policies and accounting standards.')
ON CONFLICT DO NOTHING;

INSERT INTO users (name, email, password, role_id)
VALUES (
  'Admin',
  'admin@site.com',
  '$2a$10$ye6fCR92SCGqILkWYMxHqOmTayOLu5TriJ.LDZ9fiqWYDLWV9nmrS', -- bcrypt hash of Admin123!
  1
)

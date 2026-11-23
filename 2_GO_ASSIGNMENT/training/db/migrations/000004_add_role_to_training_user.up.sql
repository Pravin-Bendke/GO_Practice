ALTER TABLE training_user 
ADD COLUMN role VARCHAR(20) NOT NULL DEFAULT 'USER';

ALTER TABLE training_user
ADD CONSTRAINT check_trainning_user_role
CHECK (role IN ('ADMIN', 'MANAGER','USER'));
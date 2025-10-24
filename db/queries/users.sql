-- CreateUser
INSERT INTO users(name,dob) 
VALUES($1,$2)
RETURNING id,name,dob,created_at,updated_at;

-- GetUserByID
SELECT id,name,dob,created_at,updated_at
FROM users
WHERE id = $1;

-- ListUsers
SELECT id,name,dob,created_at,updated_at
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

-- UpdateUser
UPDATE users
SET name = $1, dob=$2 ,updated_at= CURRENT_TIMESTAMP
WHERE id = $3
RETURNING id,name,dob, created_at,updated_at;

-- DeleteUser
DELETE FROM users
WHERE id=$1;

-- CountUsers
SELECT COUNT(*) FROM users;
-- name: CreateHousehold :exec

INSERT INTO households (id,"name",billing_status,contact_phone,address_line_1,address_line_2,city,state,zip_code,country,latitude,longitude,belongs_to_user) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13);

package utils

// ro'yxatdan o'tish
var Insert = `
	insert into users (user_name, email, pass_word) 
	values ($1, $2, crypt($3, gen_salt('bf')))
	returning *;
`

// login qilish
var Select = `
	select user_name, email from users 
	where user_name = $1 and pass_word = crypt($2, pass_word);
`

// parolni o'zgartirish
var Update = `
	update users set pass_word = crypt($1, gen_salt('bf'))
	where user_name = $2 and pass_word = crypt($3, pass_word) 
	returning user_name, email
`

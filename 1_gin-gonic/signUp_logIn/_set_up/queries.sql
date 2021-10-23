-- for checking

select * from users;

select user_name, email from users where user_name = 'Jasurbek' and pass_word = crypt('4f358411', pass_word);

update 
	users
set 
	pass_word = '1001go'

where 
	user_name = 'Jasurbek' and pass_word = crypt('4f358411', pass_word)
returning 
	*
;

update 
	users set pass_word = '121212' 
where 
	user_name = 'Jasurbek' and pass_word = crypt('1001go', pass_word) 
returning 
	user_name, email
;

update 
	users set pass_word = crypt('salom', gen_salt('bf'))
where 
	user_name = 'Jasurbek' and pass_word = crypt('d1223f8b', pass_word) 
returning 
	user_name, email
;
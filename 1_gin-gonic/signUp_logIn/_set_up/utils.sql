-- bazaga birmarta yozilgan ma'lumot ikkincha marta bazaga yozilmasligi 
-- trigger ornatamiz!!!

create or replace function users_trigger_func() returns trigger language plpgsql as $$

  declare

     u_name varchar;
     u_email varchar;

  begin

    select u.user_name from users as u into u_name where u.user_name = new.user_name;
    
      if found then

          return null;

      end if;          

    select u.email from users as u into u_email where u.email= new.email;
     
      if found then

          return null;
          
      end if;

    return new;
    
  end;

$$;


create trigger user_trigger
before insert on users
for each row execute procedure users_trigger_func();

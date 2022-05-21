# apiaive

## Technical Test

## Here is request samples :

### Appointment:

###### create:

POST localhost:3000/api/appointment/

```yaml
{
  "Name": "Gaël",
  "last_name": "Duvinage",
  "Email": "gduvinage@gmail.com",
  "Date": "2022-05-22T08:15:00+01:00",
  "vcid": 2,
}
```

return the appointment + the link that we can send by email to the user

###### Get availables slots:

GET localhost:3000/api/appointment/2/2022-05-21T00:00
Instead of return a looooot of availables slots, return only taken slots over 5 days

###### Validate appointement

GET localhost:3000/api/token/:token
Check if token is still avaible, then valid the appoitnement related to
if not, delete the token and the appointment

###### List of appointements

GET /api/admin/appointment --> **secured route**
Need header: Authorization with token
and UserName: the email or the username of the admin

return daily list of appointments of the user's center

### Vaccination-center

GET /api/vaccination-center
return the list of vaccination centers

### User

###### Create an admin user

POST /api/user/signup

```yaml
{
  "username": "Gael",
  "email": "gduvinage@gmail.com",
  "password": "123456",
  "vcid": 2,
}
```

###### get token

POST /api/user/signin

```yaml
{ "email": "gduvinage@gmail.com", "password": "123456" }
```

return the bearer token for the secured route

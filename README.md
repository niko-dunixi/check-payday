# check-payday

Really simple and easy to use terminal application that tells me if it's payday or not.
Install by running `go install github.com/paul-nelson-baker/check-payday`
Running `check-payday` will print out how many days until the actual paydate, and
the exit code will actualy reflect the number of days until payday as well. This is useful
in the situation where you only want to execute automation or further script commands on payday.

When 0 days are left until payday, it is payday, hence the exit code will be 0 which is a successful application exit code!

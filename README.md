# check-payday

### Synopsis
Really simple and easy to use terminal application that tells me if it's payday or not. Running `check-payday` will print out how many days until the actual paydate, and
the exit code will actualy reflect the number of days until payday as well. This is useful
in the situation where you only want to execute automation or further script commands on payday.

When 0 days are left until payday, it is payday, hence the exit code will be 0 which is a successful application exit code!

### Installation
Dead simple because it's Go. Install by running:
```bash
go install github.com/paul-nelson-baker/check-payday
```

### Example
```bash
#!/usr/bin/env bash

# Note: I'm redirecting output for utility usage.
check-payday > /dev/null 2>&1
days_left="${?}"
if [ "${days_left}" -eq 0 ]; then
  echo "It is indeed PAYDAY!"
  echo "Time to use automation to move money into savings, pay bills, and use AI to gamble (EG: buy bitcoin/nasdaq)."
else
  echo "We have ${days_left} days(s) until payday. Let's stop execution here."
  exit 1
fi
```

### TODO
- [x] Calculate days left until payday
- [ ] Allow setting the initial date dynamically so other people can use it

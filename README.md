# check-payday

### Synopsis
Really simple and easy to use terminal application that tells me if it's payday or not. Running
`check-payday` will print out how many days until payday, and the exit code will actually reflect
the number of days until payday as well. This is useful in the situation where you only want to
execute automation or further script commands on payday.

When 0 days are left until payday, it is payday, hence the exit code will be 0 which is a successful
application exit code!

### Installation
Dead simple because it's Go. Install by running:
```bash
go get github.com/paul-nelson-baker/check-payday
go install github.com/paul-nelson-baker/check-payday
```

### Usage
`$ check-payday --help`
```
Usage:
  check-payday [OPTIONS]

Application Options:
  -d, --date= A given date that is a payday, all calculations will be based on this date.
              Must be in format YYYY-MM-DD with leading 0's. (First payday of the year is
              advisable, but the most recent will suffice.)

Help Options:
  -h, --help  Show this help message
```

### Example
```bash
#!/usr/bin/env bash

# Note: I'm redirecting output for utility usage.
check-payday > /dev/null 2>&1
days_left="${?}"
if [ "${days_left}" -eq 0 ]; then
  echo "It is indeed PAYDAY!"
  echo "Time to use automation to move money to:"
  echo "- Improve Savings"
  echo "- Paying Bills"
  echo "- AI to gamble (EG: buy bitcoin/nasdaq/horse-races)."
else
  echo "We have ${days_left} days(s) until payday. Let's stop execution here."
  exit 1
fi
```

### TODO
- [x] Calculate days left until payday
- [x] Allow setting the initial date dynamically so other people can use it
- [ ] Restful API for gambling on horses

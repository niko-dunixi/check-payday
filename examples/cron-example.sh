#!/usr/bin/env bash

# Put a daily cron entry that looks something like this:
# 0 7 * * * EMAIL="you@example.com" "${HOME}/go/src/github.com/paul-nelson-baker/check-payday/examples/cron-example.sh"
# See wikipedia for reference: https://en.wikipedia.org/wiki/Cron

if [ -z "${EMAIL}" ] ; then
  echo "No email set. Set the environment variable EMAIL."
  exit 1
fi

# Set GOPATH and include GOPATH/bin with PATH
if [ -z "${GOPATH}" ] ; then
  GOPATH="${HOME}/go"
  export GOPATH
fi
PATH="/usr/sbin/:${PATH}:${GOPATH}/bin"
export PATH

# Check if payday, proceed if it is
check-payday > /dev/null
days_left="${?}"
if [ "${days_left}" -ne 0 ] ; then
  exit 1
fi

# Build the email string
read -rd '' EMAIL_STR <<EOM
To: ${EMAIL}
Subject: $(check-payday)
From: ${EMAIL}

It is indeed PAYDAY!
Time to use automation to move money to:
- Improve Savings
- Paying Bills
- AI to gamble (EG: buy bitcoin/nasdaq/horse-races).
EOM

# Finaly send it out
echo "${EMAIL_STR}" | sendmail -f "${EMAIL}" "${EMAIL}"

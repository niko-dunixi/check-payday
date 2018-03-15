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

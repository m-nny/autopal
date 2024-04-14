#!/usr/bin/env jq -s -R --from-file
[
  split("\r\n")[] |
  split(",") |
  select(length >= 5) |
  {
    id: .[0],
    name: .[1],
    base_hp: .[2] | tonumber,
    base_attack: .[3] | tonumber,
    base_defence: .[4] | tonumber,
    types: ([
      .[5:][] |
      select(length > 0)
    ])
  }
]

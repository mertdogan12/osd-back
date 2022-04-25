set -e

mongo <<EOF

use $MONGO_DATABASE

db.users.insertOne({
  id: $OSD_ID,
  name: "$OSD_NAME",
  groups: ["admins"],
})

db.groups.insertOne({
  name: "admins",
  permissions: ["*"],
})

EOF

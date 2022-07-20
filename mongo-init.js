db = new Mongo().getDB("osd");

db.createCollection("users", { capped: false });
db.createCollection("groups", { capped: false });

db.users.insertOne({
  id: 16775794,
  name: "Mert Dogan",
  groups: ["admins"],
  uploadedReplays: [],
  replays: [],
});

db.groups.insertOne({
  name: "admins",
  permissions: ["*"],
});

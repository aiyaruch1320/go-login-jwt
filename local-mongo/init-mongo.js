// init user in mongo
db.createUser({
    user: "root",
    pwd: "password",
    roles: [{
        role: "readWrite",
        db: "test"
    }]
})

// createIndex 
db.users.createIndex({ username: 1 }, { unique: true });
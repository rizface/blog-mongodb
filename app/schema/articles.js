db.runCommand({
    collMod: "articles",
    validationAction: "error",
    validator: {
        $jsonSchema: {
            bsonType:"object",
            required:["body","title","author"],
            properties: {
                body: {
                    bsonType: "string",
                    description: "must be string"
                },
                title: {
                    bsonType: "string",
                    description: "must be string"
                },
                author: {
                    bsonType: "string",
                    description: "must be string"
                },
                comments: {
                    bsonType: "array",
                    items: {
                        bsonType: "object",
                        required: ["author","comment"],
                        properties: {
                            author: {
                                bsonType: "string",
                                description: "must be string"
                            },
                            comment: {
                                bsonType: "string",
                                description: "must be string"
                            },
                        }
                    }
                }
            }
        }
    }
})

db.createCollection("articles", {
    validationAction: "error",
    validator: {
        $jsonSchema: {
            bsonType:"object",
            required:["body","title","author"],
            properties: {
                body: {
                    bsonType: "string",
                    description: "must be string"
                },
                title: {
                    bsonType: "string",
                    description: "must be string"
                },
                author: {
                    bsonType: "string",
                    description: "must be string"
                },
                comments: {
                    bsonType: "array",
                    items: {
                        bsonType: "object",
                        required: ["author","comment"],
                        properties: {
                            author: {
                                bsonType: "string",
                                description: "must be string"
                            },
                            comment: {
                                bsonType: "string",
                                description: "must be string"
                            },
                        }
                    }
                }
            }
        }
    }
})
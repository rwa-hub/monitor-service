const MONGO_USER = _getEnv('MONGO_USER', 'admin');
const MONGO_PASSWORD = _getEnv('MONGO_PASSWORD', 'password');
const MONGO_DB_NAME = _getEnv('MONGO_DB_NAME', 'monitor-service');

function _getEnv(varName, defaultValue) {
    if (typeof process !== 'undefined' && process.env && process.env[varName]) {
        return process.env[varName];
    }
    return defaultValue;
}

// Primeiro criamos o usuário admin no banco admin
db = db.getSiblingDB('admin');

db.createUser({
    user: MONGO_USER,
    pwd: MONGO_PASSWORD,
    roles: [
        { role: "userAdminAnyDatabase", db: "admin" },
        { role: "readWriteAnyDatabase", db: "admin" },
        { role: "dbAdminAnyDatabase", db: "admin" },
        { role: "clusterAdmin", db: "admin" }
    ]
});

// Autenticamos com o usuário admin
db.auth(MONGO_USER, MONGO_PASSWORD);

// Criamos o banco de dados da aplicação
db = db.getSiblingDB(MONGO_DB_NAME);

// Criamos o usuário específico para o banco da aplicação
db.createUser({
    user: MONGO_USER,
    pwd: MONGO_PASSWORD,
    roles: [
        { role: "readWrite", db: MONGO_DB_NAME },
        { role: "dbAdmin", db: MONGO_DB_NAME }
    ]
});

// Criamos as coleções com validação de schema
db.createCollection("events", {
    validator: {
        $jsonSchema: {
            bsonType: "object",
            required: ["eventType", "timestamp", "contractAddress"],
            properties: {
                eventType: { bsonType: "string" },
                timestamp: { bsonType: "string" },
                contractAddress: { bsonType: "string" },
                transactionHash: { bsonType: "string" },
                blockNumber: { bsonType: "string" }
            }
        }
    }
});

db.createCollection("registry_md");
db.createCollection("identity_registry");
db.createCollection("modular_compliance");


db.events.createIndex({ "timestamp": 1 });
db.events.createIndex({ "eventType": 1 });
db.events.createIndex({ "contractAddress": 1 });
db.events.createIndex({ "transactionHash": 1 }, { unique: true });

db.registry_md.createIndex({ "timestamp": 1 });
db.registry_md.createIndex({ "transactionHash": 1 }, { unique: true });

db.identity_registry.createIndex({ "timestamp": 1 });
db.identity_registry.createIndex({ "transactionHash": 1 }, { unique: true });

db.modular_compliance.createIndex({ "timestamp": 1 });
db.modular_compliance.createIndex({ "transactionHash": 1 }, { unique: true });

print("✅ Inicialização do MongoDB concluída com sucesso!");
print(`📦 Banco de dados '${MONGO_DB_NAME}' configurado`);
print(`👤 Usuário '${MONGO_USER}' criado/atualizado`);
print("📑 Coleções e índices criados"); 
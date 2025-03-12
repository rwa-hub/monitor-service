Directory structure:
└── internal/
     ├── adapters/
     │    ├── rpc/
     │    │    └── rpc.go
     │    └── websocket/
     │         └── websocket.go
     ├── app/
     │    ├── model/
     │    └── service/
     │         └── event-service.go
     ├── modules/
     │    ├── IdentRegistryStorage/
     │    │    ├── IdentityRegistryStorage.json
     │    │    └── identity_registry_storage.go
     │    ├── IdentityRegistry/
     │    │    ├── IdentityRegistry.json
     │    │    └── identity_registry.go
     │    ├── ModularCompliance/
     │    │    ├── ModularCompliance.json
     │    │    └── modular_compliance.go
     │    └── compliance/
     │         ├── FinancialRWA/
     │         │    ├── FinancialCompliance.json
     │         │    └── financial_compliance.go
     │         └── RegistryMD/
     │              ├── RegistryMDCompliance.json
     │              └── registry_md_compliance.go
     └── ports/

     